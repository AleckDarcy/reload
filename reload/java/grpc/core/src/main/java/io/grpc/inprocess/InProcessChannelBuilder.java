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

package io.grpc.inprocess;

import com.google.common.base.Preconditions;
import io.grpc.ExperimentalApi;
import io.grpc.Internal;
import io.grpc.internal.AbstractManagedChannelImplBuilder;
import io.grpc.internal.ClientTransportFactory;
import io.grpc.internal.ConnectionClientTransport;
import io.grpc.internal.GrpcUtil;
import io.grpc.internal.ProxyParameters;
import io.grpc.internal.SharedResourceHolder;
import java.net.SocketAddress;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

/**
 * Builder for a channel that issues in-process requests. Clients identify the in-process server by
 * its name.
 *
 * <p>The channel is intended to be fully-featured, high performance, and useful in testing.
 *
 * <p>For usage examples, see {@link InProcessServerBuilder}.
 */
@ExperimentalApi("https://github.com/grpc/grpc-java/issues/1783")
public final class InProcessChannelBuilder extends
        AbstractManagedChannelImplBuilder<InProcessChannelBuilder> {
  /**
   * Create a channel builder that will connect to the server with the given name.
   *
   * @param name the identity of the server to connect to
   * @return a new builder
   */
  public static InProcessChannelBuilder forName(String name) {
    return new InProcessChannelBuilder(name);
  }

  /**
   * Always fails.  Call {@link #forName} instead.
   */
  public static InProcessChannelBuilder forTarget(String target) {
    throw new UnsupportedOperationException("call forName() instead");
  }

  /**
   * Always fails.  Call {@link #forName} instead.
   */
  public static InProcessChannelBuilder forAddress(String name, int port) {
    throw new UnsupportedOperationException("call forName() instead");
  }

  private final String name;

  private InProcessChannelBuilder(String name) {
    super(new InProcessSocketAddress(name), "localhost");
    this.name = Preconditions.checkNotNull(name, "name");
    // In-process transport should not record its traffic to the stats module.
    // https://github.com/grpc/grpc-java/issues/2284
    setStatsRecordStartedRpcs(false);
    setStatsRecordFinishedRpcs(false);
  }

  @Override
  public final InProcessChannelBuilder maxInboundMessageSize(int max) {
    // TODO(carl-mastrangelo): maybe throw an exception since this not enforced?
    return super.maxInboundMessageSize(max);
  }

  /**
   * Does nothing.
   */
  @Override
  public InProcessChannelBuilder useTransportSecurity() {
    return this;
  }

  /**
   * Does nothing.
   */
  @Override
  public InProcessChannelBuilder usePlaintext(boolean skipNegotiation) {
    return this;
  }

  /** Does nothing. */
  @Override
  public InProcessChannelBuilder keepAliveTime(long keepAliveTime, TimeUnit timeUnit) {
    return this;
  }

  /** Does nothing. */
  @Override
  public InProcessChannelBuilder keepAliveTimeout(long keepAliveTimeout, TimeUnit timeUnit) {
    return this;
  }

  /** Does nothing. */
  @Override
  public InProcessChannelBuilder keepAliveWithoutCalls(boolean enable) {
    return this;
  }

  @Override
  @Internal
  protected ClientTransportFactory buildTransportFactory() {
    return new InProcessClientTransportFactory(name);
  }

  /**
   * Creates InProcess transports. Exposed for internal use, as it should be private.
   */
  @Internal
  static final class InProcessClientTransportFactory implements ClientTransportFactory {
    private final String name;
    private final ScheduledExecutorService timerService =
        SharedResourceHolder.get(GrpcUtil.TIMER_SERVICE);

    private boolean closed;

    private InProcessClientTransportFactory(String name) {
      this.name = name;
    }

    @Override
    public ConnectionClientTransport newClientTransport(
        SocketAddress addr, String authority, String userAgent, ProxyParameters proxy) {
      if (closed) {
        throw new IllegalStateException("The transport factory is closed.");
      }
      return new InProcessTransport(name, authority, userAgent);
    }

    @Override
    public ScheduledExecutorService getScheduledExecutorService() {
      return timerService;
    }

    @Override
    public void close() {
      if (closed) {
        return;
      }
      closed = true;
      SharedResourceHolder.release(GrpcUtil.TIMER_SERVICE, timerService);
    }
  }
}
