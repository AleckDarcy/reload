/*
 * Copyright 2014, gRPC Authors All rights reserved.
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

package io.grpc.stub;

import java.util.logging.Level;
import java.util.logging.Logger;

import static com.google.common.base.Preconditions.checkNotNull;
import static com.google.common.base.Preconditions.checkState;

import com.google.common.annotations.VisibleForTesting;
import com.google.common.base.Preconditions;
import io.grpc.Metadata;
import io.grpc.MethodDescriptor;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.Status;

/**
 * Utility functions for adapting {@link ServerCallHandler}s to application service implementation,
 * meant to be used by the generated code.
 */
public final class ServerCalls {
  private static final Logger logger = Logger.getLogger(ServerCalls.class.getName());

  @VisibleForTesting
  static final String TOO_MANY_REQUESTS = "Too many requests";
  @VisibleForTesting
  static final String MISSING_REQUEST = "Half-closed without a request";

  private ServerCalls() {
  }

  /**
   * Creates a {@code ServerCallHandler} for a unary call method of the service.
   *
   * @param method an adaptor to the actual method on the service implementation.
   */
  public static <ReqT, RespT> ServerCallHandler<ReqT, RespT> asyncUnaryCall(
      final UnaryMethod<ReqT, RespT> method) {
    return asyncUnaryRequestCall(method);
  }

  /**
   * Creates a {@code ServerCallHandler} for a server streaming method of the service.
   *
   * @param method an adaptor to the actual method on the service implementation.
   */
  public static <ReqT, RespT> ServerCallHandler<ReqT, RespT> asyncServerStreamingCall(
      final ServerStreamingMethod<ReqT, RespT> method) {
    return asyncUnaryRequestCall(method);
  }

  /**
   * Creates a {@code ServerCallHandler} for a client streaming method of the service.
   *
   * @param method an adaptor to the actual method on the service implementation.
   */
  public static <ReqT, RespT> ServerCallHandler<ReqT, RespT> asyncClientStreamingCall(
      final ClientStreamingMethod<ReqT, RespT> method) {
    return asyncStreamingRequestCall(method);
  }

  /**
   * Creates a {@code ServerCallHandler} for a bidi streaming method of the service.
   *
   * @param method an adaptor to the actual method on the service implementation.
   */
  public static <ReqT, RespT> ServerCallHandler<ReqT, RespT> asyncBidiStreamingCall(
      final BidiStreamingMethod<ReqT, RespT> method) {
    return asyncStreamingRequestCall(method);
  }

  /**
   * Adaptor to a unary call method.
   */
  public interface UnaryMethod<ReqT, RespT> extends UnaryRequestMethod<ReqT, RespT> {}

  /**
   * Adaptor to a server streaming method.
   */
  public interface ServerStreamingMethod<ReqT, RespT> extends UnaryRequestMethod<ReqT, RespT> {}

  /**
   * Adaptor to a client streaming method.
   */
  public interface ClientStreamingMethod<ReqT, RespT> extends StreamingRequestMethod<ReqT, RespT> {}

  /**
   * Adaptor to a bi-directional streaming method.
   */
  public interface BidiStreamingMethod<ReqT, RespT> extends StreamingRequestMethod<ReqT, RespT> {}

  private static final class UnaryServerCallHandler<ReqT, RespT>
      implements ServerCallHandler<ReqT, RespT> {

    private final UnaryRequestMethod<ReqT, RespT> method;

    // Non private to avoid synthetic class
    UnaryServerCallHandler(UnaryRequestMethod<ReqT, RespT> method) {
      this.method = method;
    }

    @Override
    public ServerCall.Listener<ReqT> startCall(ServerCall<ReqT, RespT> call, Metadata headers) {
      Preconditions.checkArgument(
          call.getMethodDescriptor().getType().clientSendsOneMessage(),
          "asyncUnaryRequestCall is only for clientSendsOneMessage methods");
      ServerCallStreamObserverImpl<ReqT, RespT> responseObserver =
          new ServerCallStreamObserverImpl<ReqT, RespT>(call);
      // We expect only 1 request, but we ask for 2 requests here so that if a misbehaving client
      // sends more than 1 requests, ServerCall will catch it. Note that disabling auto
      // inbound flow control has no effect on unary calls.
      call.request(2);
      return new UnaryServerCallListener(responseObserver, call);
    }

    private final class UnaryServerCallListener extends ServerCall.Listener<ReqT> {
      private final ServerCall<ReqT, RespT> call;
      private final ServerCallStreamObserverImpl<ReqT, RespT> responseObserver;
      private boolean canInvoke = true;
      private ReqT request;

      // Non private to avoid synthetic class
      UnaryServerCallListener(
          ServerCallStreamObserverImpl<ReqT, RespT> responseObserver,
          ServerCall<ReqT, RespT> call) {
        this.call = call;
        this.responseObserver = responseObserver;
      }

      @Override
      public void onMessage(ReqT request) {
        if (this.request != null) {
          // Safe to close the call, because the application has not yet been invoked
          call.close(
              Status.INTERNAL.withDescription(TOO_MANY_REQUESTS),
              new Metadata());
          canInvoke = false;
          return;
        }

        // We delay calling method.invoke() until onHalfClose() to make sure the client
        // half-closes.
        this.request = request;
      }

      @Override
      public void onHalfClose() {
        if (!canInvoke) {
          return;
        }
        if (request == null) {
          // Safe to close the call, because the application has not yet been invoked
          call.close(
              Status.INTERNAL.withDescription(MISSING_REQUEST),
              new Metadata());
          return;
        }

        boolean crashed = false;
        if (request instanceof io.grpc.tracer.Tracer) {
          long threadID = Thread.currentThread().getId();
//          logger.log(Level.INFO, "[RELOAD] invoke, threadID: " + threadID);

          io.grpc.tracer.Tracer tracer = (io.grpc.tracer.Tracer) request;
          io.grpc.tracer.Message.Trace trace = tracer.GetFI_Trace();
          if (trace != null) {
            java.util.List<io.grpc.tracer.Message.Record> records = trace.getRecordsList();
            if (records.size() == 1) {
              logger.log(Level.INFO, "[RELOAD] onHalfClose, trace: " + trace + ", name: " + tracer.GetFI_Name());
              crashed = trace.DoFI(tracer.GetFI_Name());
              logger.log(Level.INFO, "[RELOAD] onHalfClose, crashed: " + crashed);

              io.grpc.tracer.Message.Trace trace1 = trace.copy();
              trace1.replaceRecord(0,
                new io.grpc.tracer.Message.Record(
                  io.grpc.tracer.Message.RecordType.RecordReceive_VALUE,
                  tracer.GetFI_Name(),
                  records.get(0).getUuid(),
                  io.grpc.tracer.Store.uuid));
              logger.log(Level.INFO, "[RELOAD] onHalfClose, set trace, thread: " + threadID + ", trace:" + trace1);
              io.grpc.tracer.Store.SetTrace(threadID, trace1);
            } else {
              logger.log(Level.INFO, "[RELOAD] request is an invalid tracer:" + request.getClass() + ", records: " + records.size());
            }
          } else {
//            logger.log(Level.INFO, "[RELOAD] request is an empty tracer:" + request.getClass() + ", name: " + tracer.GetFI_Name());
          }
        } else {
//          logger.log(Level.INFO, "[RELOAD] request is not a tracer:" + request.getClass());
        }

        if (crashed) {
          logger.log(Level.INFO, "[RELOAD] service crashed");

          return;
        }

        method.invoke(request, responseObserver);
        responseObserver.freeze();
        if (call.isReady()) {
          // Since we are calling invoke in halfClose we have missed the onReady
          // event from the transport so recover it here.
          onReady();
        }
      }

      @Override
      public void onCancel() {
        responseObserver.cancelled = true;
        if (responseObserver.onCancelHandler != null) {
          responseObserver.onCancelHandler.run();
        }
      }

      @Override
      public void onReady() {
        if (responseObserver.onReadyHandler != null) {
          responseObserver.onReadyHandler.run();
        }
      }
    }
  }

  /**
   * Creates a {@code ServerCallHandler} for a unary request call method of the service.
   *
   * @param method an adaptor to the actual method on the service implementation.
   */
  private static <ReqT, RespT> ServerCallHandler<ReqT, RespT> asyncUnaryRequestCall(
      UnaryRequestMethod<ReqT, RespT> method) {
    return new UnaryServerCallHandler<ReqT, RespT>(method);
  }

  private static final class StreamingServerCallHandler<ReqT, RespT>
      implements ServerCallHandler<ReqT, RespT> {

    private final StreamingRequestMethod<ReqT, RespT> method;

    // Non private to avoid synthetic class
    StreamingServerCallHandler(StreamingRequestMethod<ReqT, RespT> method) {
      this.method = method;
    }

    @Override
    public ServerCall.Listener<ReqT> startCall(ServerCall<ReqT, RespT> call, Metadata headers) {
      ServerCallStreamObserverImpl<ReqT, RespT> responseObserver =
          new ServerCallStreamObserverImpl<ReqT, RespT>(call);
      StreamObserver<ReqT> requestObserver = method.invoke(responseObserver);
      responseObserver.freeze();
      if (responseObserver.autoFlowControlEnabled) {
        call.request(1);
      }
      return new StreamingServerCallListener(requestObserver, responseObserver, call);
    }

    private final class StreamingServerCallListener extends ServerCall.Listener<ReqT> {

      private final StreamObserver<ReqT> requestObserver;
      private final ServerCallStreamObserverImpl<ReqT, RespT> responseObserver;
      private final ServerCall<ReqT, RespT> call;
      private boolean halfClosed = false;

      // Non private to avoid synthetic class
      StreamingServerCallListener(
          StreamObserver<ReqT> requestObserver,
          ServerCallStreamObserverImpl<ReqT, RespT> responseObserver,
          ServerCall<ReqT, RespT> call) {
        this.requestObserver = requestObserver;
        this.responseObserver = responseObserver;
        this.call = call;
      }

      @Override
      public void onMessage(ReqT request) {
        requestObserver.onNext(request);

        // Request delivery of the next inbound message.
        if (responseObserver.autoFlowControlEnabled) {
          call.request(1);
        }
      }

      @Override
      public void onHalfClose() {
        halfClosed = true;
        requestObserver.onCompleted();
      }

      @Override
      public void onCancel() {
        responseObserver.cancelled = true;
        if (responseObserver.onCancelHandler != null) {
          responseObserver.onCancelHandler.run();
        }
        if (!halfClosed) {
          requestObserver.onError(
              Status.CANCELLED
                  .withDescription("cancelled before receiving half close")
                  .asRuntimeException());
        }
      }

      @Override
      public void onReady() {
        if (responseObserver.onReadyHandler != null) {
          responseObserver.onReadyHandler.run();
        }
      }
    }
  }

  /**
   * Creates a {@code ServerCallHandler} for a streaming request call method of the service.
   *
   * @param method an adaptor to the actual method on the service implementation.
   */
  private static <ReqT, RespT> ServerCallHandler<ReqT, RespT> asyncStreamingRequestCall(
      StreamingRequestMethod<ReqT, RespT> method) {
    return new StreamingServerCallHandler<ReqT, RespT>(method);
  }

  private interface UnaryRequestMethod<ReqT, RespT> {
    void invoke(ReqT request, StreamObserver<RespT> responseObserver);
  }

  private interface StreamingRequestMethod<ReqT, RespT> {
    StreamObserver<ReqT> invoke(StreamObserver<RespT> responseObserver);
  }

  private static final class ServerCallStreamObserverImpl<ReqT, RespT>
      extends ServerCallStreamObserver<RespT> {
    final ServerCall<ReqT, RespT> call;
    volatile boolean cancelled;
    private boolean frozen;
    private boolean autoFlowControlEnabled = true;
    private boolean sentHeaders;
    private Runnable onReadyHandler;
    private Runnable onCancelHandler;

    // Non private to avoid synthetic class
    ServerCallStreamObserverImpl(ServerCall<ReqT, RespT> call) {
      this.call = call;
    }

    private void freeze() {
      this.frozen = true;
    }

    @Override
    public void setMessageCompression(boolean enable) {
      call.setMessageCompression(enable);
    }

    @Override
    public void setCompression(String compression) {
      call.setCompression(compression);
    }

    @Override
    public void onNext(RespT response) {
//      StackTraceElement[] st = Thread.currentThread().getStackTrace();
//      for (int i = 0; i < st.length; i++) {
//        logger.info("[RELOAD] onNext, stacktrace " + i+ ", " + st[i].toString());
//      }

      if (cancelled) {
        throw Status.CANCELLED.withDescription("call already cancelled").asRuntimeException();
      }
      if (!sentHeaders) {
        call.sendHeaders(new Metadata());
        sentHeaders = true;
      }
      call.sendMessage(response);
    }

    @Override
    public void onError(Throwable t) {
      Metadata metadata = Status.trailersFromThrowable(t);
      if (metadata == null) {
        metadata = new Metadata();
      }
      call.close(Status.fromThrowable(t), metadata);
    }

    @Override
    public void onCompleted() {
      if (cancelled) {
        throw Status.CANCELLED.withDescription("call already cancelled").asRuntimeException();
      } else {
        call.close(Status.OK, new Metadata());
      }
    }

    @Override
    public boolean isReady() {
      return call.isReady();
    }

    @Override
    public void setOnReadyHandler(Runnable r) {
      checkState(!frozen, "Cannot alter onReadyHandler after initialization");
      this.onReadyHandler = r;
    }

    @Override
    public boolean isCancelled() {
      return call.isCancelled();
    }

    @Override
    public void setOnCancelHandler(Runnable onCancelHandler) {
      checkState(!frozen, "Cannot alter onCancelHandler after initialization");
      this.onCancelHandler = onCancelHandler;
    }

    @Override
    public void disableAutoInboundFlowControl() {
      checkState(!frozen, "Cannot disable auto flow control after initialization");
      autoFlowControlEnabled = false;
    }

    @Override
    public void request(int count) {
      call.request(count);
    }
  }

  /**
   * Sets unimplemented status for method on given response stream for unary call.
   *
   * @param methodDescriptor of method for which error will be thrown.
   * @param responseObserver on which error will be set.
   */
  public static void asyncUnimplementedUnaryCall(MethodDescriptor<?, ?> methodDescriptor,
      StreamObserver<?> responseObserver) {
    checkNotNull(methodDescriptor, "methodDescriptor");
    checkNotNull(responseObserver, "responseObserver");
    responseObserver.onError(Status.UNIMPLEMENTED
        .withDescription(String.format("Method %s is unimplemented",
            methodDescriptor.getFullMethodName()))
        .asRuntimeException());
  }

  /**
   * Sets unimplemented status for streaming call.
   *
   * @param methodDescriptor of method for which error will be thrown.
   * @param responseObserver on which error will be set.
   */
  public static <T> StreamObserver<T> asyncUnimplementedStreamingCall(
      MethodDescriptor<?, ?> methodDescriptor, StreamObserver<?> responseObserver) {
    // NB: For streaming call we want to do the same as for unary call. Fail-fast by setting error
    // on responseObserver and then return no-op observer.
    asyncUnimplementedUnaryCall(methodDescriptor, responseObserver);
    return new NoopStreamObserver<T>();
  }

  /**
   * No-op implementation of StreamObserver. Used in abstract stubs for default implementations of
   * methods which throws UNIMPLEMENTED error and tests.
   */
  static class NoopStreamObserver<V> implements StreamObserver<V> {
    @Override
    public void onNext(V value) {
    }

    @Override
    public void onError(Throwable t) {
    }

    @Override
    public void onCompleted() {
    }
  }
}
