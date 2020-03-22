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

import java.io.File;
import java.util.concurrent.Executor;
import java.util.concurrent.TimeUnit;
import javax.annotation.Nullable;

/**
 * A builder for {@link Server} instances.
 *
 * @param <T> The concrete type of this builder.
 * @since 1.0.0
 */
public abstract class ServerBuilder<T extends ServerBuilder<T>> {

  /**
   * Static factory for creating a new ServerBuilder.
   *
   * @param port the port to listen on
   * @since 1.0.0
   */
  public static ServerBuilder<?> forPort(int port) {
    return ServerProvider.provider().builderForPort(port);
  }

  /**
   * Execute application code directly in the transport thread.
   *
   * <p>Depending on the underlying transport, using a direct executor may lead to substantial
   * performance improvements. However, it also requires the application to not block under
   * any circumstances.
   *
   * <p>Calling this method is semantically equivalent to calling {@link #executor(Executor)} and
   * passing in a direct executor. However, this is the preferred way as it may allow the transport
   * to perform special optimizations.
   *
   * @return this
   * @since 1.0.0
   */
  public abstract T directExecutor();

  /**
   * Provides a custom executor.
   *
   * <p>It's an optional parameter. If the user has not provided an executor when the server is
   * built, the builder will use a static cached thread pool.
   *
   * <p>The server won't take ownership of the given executor. It's caller's responsibility to
   * shut down the executor when it's desired.
   *
   * @return this
   * @since 1.0.0
   */
  public abstract T executor(@Nullable Executor executor);

  /**
   * Adds a service implementation to the handler registry.
   *
   * @param service ServerServiceDefinition object
   * @return this
   * @since 1.0.0
   */
  public abstract T addService(ServerServiceDefinition service);

  /**
   * Adds a service implementation to the handler registry. If bindableService implements
   * {@link InternalNotifyOnServerBuild}, the service will receive a reference to the generated
   * server instance upon build().
   *
   * @param bindableService BindableService object
   * @return this
   * @since 1.0.0
   */
  public abstract T addService(BindableService bindableService);

  /**
   * Adds a {@link ServerInterceptor} that is run for all services on the server.  Interceptors
   * added through this method always run before per-service interceptors added through {@link
   * ServerInterceptors}.  Interceptors run in the reverse order in which they are added.
   *
   * @param interceptor the all-service interceptor
   * @return this
   * @since 1.5.0
   */
  @ExperimentalApi("https://github.com/grpc/grpc-java/issues/3117")
  public T intercept(ServerInterceptor interceptor) {
    throw new UnsupportedOperationException();
  }

  /**
   * Adds a {@link ServerTransportFilter}. The order of filters being added is the order they will
   * be executed.
   *
   * @return this
   * @since 1.2.0
   */
  @ExperimentalApi("https://github.com/grpc/grpc-java/issues/2132")
  public T addTransportFilter(ServerTransportFilter filter) {
    throw new UnsupportedOperationException();
  }

  /**
   * Adds a {@link ServerStreamTracer.Factory} to measure server-side traffic.  The order of
   * factories being added is the order they will be executed.
   *
   * @return this
   * @since 1.3.0
   */
  @ExperimentalApi("https://github.com/grpc/grpc-java/issues/2861")
  public T addStreamTracerFactory(ServerStreamTracer.Factory factory) {
    throw new UnsupportedOperationException();
  }

  /**
   * Sets a fallback handler registry that will be looked up in if a method is not found in the
   * primary registry. The primary registry (configured via {@code addService()}) is faster but
   * immutable. The fallback registry is more flexible and allows implementations to mutate over
   * time and load services on-demand.
   *
   * @return this
   * @since 1.0.0
   */
  public abstract T fallbackHandlerRegistry(@Nullable HandlerRegistry fallbackRegistry);

  /**
   * Makes the server use TLS.
   *
   * @param certChain file containing the full certificate chain
   * @param privateKey file containing the private key
   *
   * @return this
   * @throws UnsupportedOperationException if the server does not support TLS.
   * @since 1.0.0
   */
  public abstract T useTransportSecurity(File certChain, File privateKey);

  /**
   * Set the decompression registry for use in the channel.  This is an advanced API call and
   * shouldn't be used unless you are using custom message encoding.   The default supported
   * decompressors are in {@code DecompressorRegistry.getDefaultInstance}.
   *
   * @return this
   * @since 1.0.0
   */
  @ExperimentalApi("https://github.com/grpc/grpc-java/issues/1704")
  public abstract T decompressorRegistry(@Nullable DecompressorRegistry registry);

  /**
   * Set the compression registry for use in the channel.  This is an advanced API call and
   * shouldn't be used unless you are using custom message encoding.   The default supported
   * compressors are in {@code CompressorRegistry.getDefaultInstance}.
   *
   * @return this
   * @since 1.0.0
   */
  @ExperimentalApi("https://github.com/grpc/grpc-java/issues/1704")
  public abstract T compressorRegistry(@Nullable CompressorRegistry registry);

  /**
   * Sets the permitted time for new connections to complete negotiation handshakes before being
   * killed.
   *
   * @return this
   * @throws IllegalArgumentException if timeout is negative
   * @throws UnsupportedOperationException if unsupported
   * @since 1.8.0
   */
  @ExperimentalApi("https://github.com/grpc/grpc-java/issues/3706")
  public T handshakeTimeout(long timeout, TimeUnit unit) {
    throw new UnsupportedOperationException();
  }

  /**
   * Builds a server using the given parameters.
   *
   * <p>The returned service will not been started or be bound a port. You will need to start it
   * with {@link Server#start()}.
   *
   * @return a new Server
   * @since 1.0.0
   */
  public abstract Server build();
}
