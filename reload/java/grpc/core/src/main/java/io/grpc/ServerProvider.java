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

import com.google.common.annotations.VisibleForTesting;
import io.grpc.ManagedChannelProvider.ProviderNotFoundException;
import java.util.ServiceLoader;

/**
 * Provider of servers for transport agnostic consumption.
 *
 * <p>Implementations <em>should not</em> throw. If they do, it may interrupt class loading. If
 * exceptions may reasonably occur for implementation-specific reasons, implementations should
 * generally handle the exception gracefully and return {@code false} from {@link #isAvailable()}.
 */
@Internal
public abstract class ServerProvider {
  private static final ServerProvider provider =
      load(ServerProvider.class.getClassLoader());

  @VisibleForTesting
  static final ServerProvider load(ClassLoader cl) {
    ServiceLoader<ServerProvider> providers = ServiceLoader.load(ServerProvider.class, cl);

    // Attempt to load using the context class loader and ServiceLoader.
    // This allows frameworks like http://aries.apache.org/modules/spi-fly.html to plug in.
    if (!providers.iterator().hasNext()) {
      providers = ServiceLoader.load(ServerProvider.class);
    }

    ServerProvider best = null;

    for (ServerProvider current : providers) {
      if (!current.isAvailable()) {
        continue;
      } else if (best == null) {
        best = current;
      } else if (current.priority() > best.priority()) {
        best = current;
      }
    }
    return best;
  }

  /**
   * Returns the ClassLoader-wide default server.
   *
   * @throws ProviderNotFoundException if no provider is available
   */
  public static ServerProvider provider() {
    if (provider == null) {
      throw new ProviderNotFoundException("No functional server found. "
          + "Try adding a dependency on the grpc-netty artifact");
    }
    return provider;
  }

  /**
   * Whether this provider is available for use, taking the current environment into consideration.
   * If {@code false}, no other methods are safe to be called.
   */
  protected abstract boolean isAvailable();

  /**
   * A priority, from 0 to 10 that this provider should be used, taking the current environment into
   * consideration. 5 should be considered the default, and then tweaked based on environment
   * detection. A priority of 0 does not imply that the provider wouldn't work; just that it should
   * be last in line.
   */
  protected abstract int priority();

  /**
   * Creates a new builder with the given port.
   */
  protected abstract ServerBuilder<?> builderForPort(int port);
}
