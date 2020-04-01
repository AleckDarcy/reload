/*
 * Copyright 2015, gRPC Authors All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.grpc;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertSame;
import static org.junit.Assert.assertTrue;
import static org.junit.Assert.fail;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.ServiceConfigurationError;
import java.util.regex.Pattern;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

/** Unit tests for {@link ManagedChannelProvider}. */
@RunWith(JUnit4.class)
public class ManagedChannelProviderTest {
  private final String serviceFile = "META-INF/services/io.grpc.ManagedChannelProvider";

  @Test(expected = ManagedChannelProvider.ProviderNotFoundException.class)
  public void noProvider() {
    ManagedChannelProvider.provider();
  }

  @Test
  public void multipleProvider() {
    ClassLoader cl = new ReplacingClassLoader(getClass().getClassLoader(), serviceFile,
        "io/grpc/ManagedChannelProviderTest-multipleProvider.txt");
    assertSame(Available7Provider.class, ManagedChannelProvider.load(cl).getClass());
  }

  @Test
  public void unavailableProvider() {
    ClassLoader cl = new ReplacingClassLoader(getClass().getClassLoader(), serviceFile,
        "io/grpc/ManagedChannelProviderTest-unavailableProvider.txt");
    assertNull(ManagedChannelProvider.load(cl));
  }

  @Test
  public void getCandidatesViaHardCoded_triesToLoadClasses() throws Exception {
    ClassLoader cl = getClass().getClassLoader();
    final RuntimeException toThrow = new RuntimeException();
    cl = new ClassLoader(cl) {
      @Override
      public Class<?> loadClass(String name, boolean resolve) throws ClassNotFoundException {
        if (name.startsWith("io.grpc.netty.") || name.startsWith("io.grpc.okhttp.")) {
          throw toThrow;
        } else {
          return super.loadClass(name, resolve);
        }
      }
    };
    cl = new StaticTestingClassLoader(cl, Pattern.compile("io\\.grpc\\.[^.]*"));
    try {
      invokeGetCandidatesViaHardCoded(cl);
      fail("Expected exception");
    } catch (RuntimeException ex) {
      assertSame(toThrow, ex);
    }
  }

  @Test
  public void getCandidatesViaHardCoded_ignoresMissingClasses() throws Exception {
    ClassLoader cl = getClass().getClassLoader();
    cl = new ClassLoader(cl) {
      @Override
      public Class<?> loadClass(String name, boolean resolve) throws ClassNotFoundException {
        if (name.startsWith("io.grpc.netty.") || name.startsWith("io.grpc.okhttp.")) {
          throw new ClassNotFoundException();
        } else {
          return super.loadClass(name, resolve);
        }
      }
    };
    cl = new StaticTestingClassLoader(cl, Pattern.compile("io\\.grpc\\.[^.]*"));
    Iterable<?> i = invokeGetCandidatesViaHardCoded(cl);
    assertFalse("Iterator should be empty", i.iterator().hasNext());
  }

  @Test
  public void create_throwsErrorOnMisconfiguration() throws Exception {
    class PrivateClass {}

    try {
      ManagedChannelProvider.create(PrivateClass.class);
      fail("Expected exception");
    } catch (ServiceConfigurationError e) {
      assertTrue("Expected ClassCastException cause: " + e.getCause(),
          e.getCause() instanceof ClassCastException);
    }
  }

  private static Iterable<?> invokeGetCandidatesViaHardCoded(ClassLoader cl) throws Exception {
    // An error before the invoke likely means there is a bug in the test
    Class<?> klass = Class.forName(ManagedChannelProvider.class.getName(), true, cl);
    Method getCandidatesViaHardCoded = klass.getMethod("getCandidatesViaHardCoded");
    try {
      return (Iterable<?>) getCandidatesViaHardCoded.invoke(null);
    } catch (InvocationTargetException ex) {
      if (ex.getCause() instanceof Exception) {
        throw (Exception) ex.getCause();
      }
      throw ex;
    }
  }

  private static class BaseProvider extends ManagedChannelProvider {
    private final boolean isAvailable;
    private final int priority;

    public BaseProvider(boolean isAvailable, int priority) {
      this.isAvailable = isAvailable;
      this.priority = priority;
    }

    @Override
    protected boolean isAvailable() {
      return isAvailable;
    }

    @Override
    protected int priority() {
      return priority;
    }

    @Override
    protected ManagedChannelBuilder<?> builderForAddress(String host, int port) {
      throw new UnsupportedOperationException();
    }

    @Override
    protected ManagedChannelBuilder<?> builderForTarget(String target) {
      throw new UnsupportedOperationException();
    }
  }

  public static class Available0Provider extends BaseProvider {
    public Available0Provider() {
      super(true, 0);
    }
  }

  public static class Available5Provider extends BaseProvider {
    public Available5Provider() {
      super(true, 5);
    }
  }

  public static class Available7Provider extends BaseProvider {
    public Available7Provider() {
      super(true, 7);
    }
  }

  public static class UnavailableProvider extends BaseProvider {
    public UnavailableProvider() {
      super(false, 10);
    }

    @Override
    protected int priority() {
      throw new RuntimeException("purposefully broken");
    }
  }
}
