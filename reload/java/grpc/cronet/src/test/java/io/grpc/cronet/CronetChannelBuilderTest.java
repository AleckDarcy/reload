/*
 * Copyright 2016, gRPC Authors All rights reserved.
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

package io.grpc.cronet;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

import io.grpc.CallOptions;
import io.grpc.Metadata;
import io.grpc.MethodDescriptor;
import io.grpc.cronet.CronetChannelBuilder.CronetTransportFactory;
import io.grpc.testing.TestMethodDescriptors;
import java.net.InetSocketAddress;
import org.chromium.net.CronetEngine;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.robolectric.RobolectricTestRunner;

@RunWith(RobolectricTestRunner.class)
public final class CronetChannelBuilderTest {

  @Mock private CronetEngine mockEngine;

  private MethodDescriptor<?, ?> method = TestMethodDescriptors.voidMethod();

  @Before
  public void setUp() {
    MockitoAnnotations.initMocks(this);
  }

  @Test
  public void alwaysUsePutTrue_cronetStreamIsIdempotent() throws Exception {
    CronetChannelBuilder builder =
        CronetChannelBuilder.forAddress("address", 1234, mockEngine).alwaysUsePut(true);
    CronetTransportFactory transportFactory =
        (CronetTransportFactory) builder.buildTransportFactory();
    CronetClientTransport transport =
        (CronetClientTransport)
            transportFactory.newClientTransport(
                new InetSocketAddress("localhost", 443), "", null, null);
    CronetClientStream stream = transport.newStream(method, new Metadata(), CallOptions.DEFAULT);

    assertTrue(stream.idempotent);
  }

  @Test
  public void alwaysUsePut_defaultsToFalse() throws Exception {
    CronetChannelBuilder builder = CronetChannelBuilder.forAddress("address", 1234, mockEngine);
    CronetTransportFactory transportFactory =
        (CronetTransportFactory) builder.buildTransportFactory();
    CronetClientTransport transport =
        (CronetClientTransport)
            transportFactory.newClientTransport(
                new InetSocketAddress("localhost", 443), "", null, null);
    CronetClientStream stream = transport.newStream(method, new Metadata(), CallOptions.DEFAULT);

    assertFalse(stream.idempotent);
  }
}
