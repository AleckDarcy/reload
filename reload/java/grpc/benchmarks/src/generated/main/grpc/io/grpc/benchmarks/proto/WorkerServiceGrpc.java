package io.grpc.benchmarks.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler",
    comments = "Source: services.proto")
public final class WorkerServiceGrpc {

  private WorkerServiceGrpc() {}

  public static final String SERVICE_NAME = "grpc.testing.WorkerService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getRunServerMethod()} instead. 
  public static final io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ServerArgs,
      io.grpc.benchmarks.proto.Control.ServerStatus> METHOD_RUN_SERVER = getRunServerMethodHelper();

  private static volatile io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ServerArgs,
      io.grpc.benchmarks.proto.Control.ServerStatus> getRunServerMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ServerArgs,
      io.grpc.benchmarks.proto.Control.ServerStatus> getRunServerMethod() {
    return getRunServerMethodHelper();
  }

  private static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ServerArgs,
      io.grpc.benchmarks.proto.Control.ServerStatus> getRunServerMethodHelper() {
    io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ServerArgs, io.grpc.benchmarks.proto.Control.ServerStatus> getRunServerMethod;
    if ((getRunServerMethod = WorkerServiceGrpc.getRunServerMethod) == null) {
      synchronized (WorkerServiceGrpc.class) {
        if ((getRunServerMethod = WorkerServiceGrpc.getRunServerMethod) == null) {
          WorkerServiceGrpc.getRunServerMethod = getRunServerMethod = 
              io.grpc.MethodDescriptor.<io.grpc.benchmarks.proto.Control.ServerArgs, io.grpc.benchmarks.proto.Control.ServerStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(
                  "grpc.testing.WorkerService", "RunServer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.ServerArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.ServerStatus.getDefaultInstance()))
                  .setSchemaDescriptor(new WorkerServiceMethodDescriptorSupplier("RunServer"))
                  .build();
          }
        }
     }
     return getRunServerMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getRunClientMethod()} instead. 
  public static final io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ClientArgs,
      io.grpc.benchmarks.proto.Control.ClientStatus> METHOD_RUN_CLIENT = getRunClientMethodHelper();

  private static volatile io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ClientArgs,
      io.grpc.benchmarks.proto.Control.ClientStatus> getRunClientMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ClientArgs,
      io.grpc.benchmarks.proto.Control.ClientStatus> getRunClientMethod() {
    return getRunClientMethodHelper();
  }

  private static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ClientArgs,
      io.grpc.benchmarks.proto.Control.ClientStatus> getRunClientMethodHelper() {
    io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.ClientArgs, io.grpc.benchmarks.proto.Control.ClientStatus> getRunClientMethod;
    if ((getRunClientMethod = WorkerServiceGrpc.getRunClientMethod) == null) {
      synchronized (WorkerServiceGrpc.class) {
        if ((getRunClientMethod = WorkerServiceGrpc.getRunClientMethod) == null) {
          WorkerServiceGrpc.getRunClientMethod = getRunClientMethod = 
              io.grpc.MethodDescriptor.<io.grpc.benchmarks.proto.Control.ClientArgs, io.grpc.benchmarks.proto.Control.ClientStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(
                  "grpc.testing.WorkerService", "RunClient"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.ClientArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.ClientStatus.getDefaultInstance()))
                  .setSchemaDescriptor(new WorkerServiceMethodDescriptorSupplier("RunClient"))
                  .build();
          }
        }
     }
     return getRunClientMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getCoreCountMethod()} instead. 
  public static final io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.CoreRequest,
      io.grpc.benchmarks.proto.Control.CoreResponse> METHOD_CORE_COUNT = getCoreCountMethodHelper();

  private static volatile io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.CoreRequest,
      io.grpc.benchmarks.proto.Control.CoreResponse> getCoreCountMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.CoreRequest,
      io.grpc.benchmarks.proto.Control.CoreResponse> getCoreCountMethod() {
    return getCoreCountMethodHelper();
  }

  private static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.CoreRequest,
      io.grpc.benchmarks.proto.Control.CoreResponse> getCoreCountMethodHelper() {
    io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.CoreRequest, io.grpc.benchmarks.proto.Control.CoreResponse> getCoreCountMethod;
    if ((getCoreCountMethod = WorkerServiceGrpc.getCoreCountMethod) == null) {
      synchronized (WorkerServiceGrpc.class) {
        if ((getCoreCountMethod = WorkerServiceGrpc.getCoreCountMethod) == null) {
          WorkerServiceGrpc.getCoreCountMethod = getCoreCountMethod = 
              io.grpc.MethodDescriptor.<io.grpc.benchmarks.proto.Control.CoreRequest, io.grpc.benchmarks.proto.Control.CoreResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "grpc.testing.WorkerService", "CoreCount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.CoreRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.CoreResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new WorkerServiceMethodDescriptorSupplier("CoreCount"))
                  .build();
          }
        }
     }
     return getCoreCountMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getQuitWorkerMethod()} instead. 
  public static final io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.Void,
      io.grpc.benchmarks.proto.Control.Void> METHOD_QUIT_WORKER = getQuitWorkerMethodHelper();

  private static volatile io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.Void,
      io.grpc.benchmarks.proto.Control.Void> getQuitWorkerMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.Void,
      io.grpc.benchmarks.proto.Control.Void> getQuitWorkerMethod() {
    return getQuitWorkerMethodHelper();
  }

  private static io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.Void,
      io.grpc.benchmarks.proto.Control.Void> getQuitWorkerMethodHelper() {
    io.grpc.MethodDescriptor<io.grpc.benchmarks.proto.Control.Void, io.grpc.benchmarks.proto.Control.Void> getQuitWorkerMethod;
    if ((getQuitWorkerMethod = WorkerServiceGrpc.getQuitWorkerMethod) == null) {
      synchronized (WorkerServiceGrpc.class) {
        if ((getQuitWorkerMethod = WorkerServiceGrpc.getQuitWorkerMethod) == null) {
          WorkerServiceGrpc.getQuitWorkerMethod = getQuitWorkerMethod = 
              io.grpc.MethodDescriptor.<io.grpc.benchmarks.proto.Control.Void, io.grpc.benchmarks.proto.Control.Void>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "grpc.testing.WorkerService", "QuitWorker"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.Void.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.benchmarks.proto.Control.Void.getDefaultInstance()))
                  .setSchemaDescriptor(new WorkerServiceMethodDescriptorSupplier("QuitWorker"))
                  .build();
          }
        }
     }
     return getQuitWorkerMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static WorkerServiceStub newStub(io.grpc.Channel channel) {
    return new WorkerServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static WorkerServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new WorkerServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static WorkerServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new WorkerServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class WorkerServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * Start server with specified workload.
     * First request sent specifies the ServerConfig followed by ServerStatus
     * response. After that, a "Mark" can be sent anytime to request the latest
     * stats. Closing the stream will initiate shutdown of the test server
     * and once the shutdown has finished, the OK status is sent to terminate
     * this RPC.
     * </pre>
     */
    public io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ServerArgs> runServer(
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ServerStatus> responseObserver) {
      return asyncUnimplementedStreamingCall(getRunServerMethodHelper(), responseObserver);
    }

    /**
     * <pre>
     * Start client with specified workload.
     * First request sent specifies the ClientConfig followed by ClientStatus
     * response. After that, a "Mark" can be sent anytime to request the latest
     * stats. Closing the stream will initiate shutdown of the test client
     * and once the shutdown has finished, the OK status is sent to terminate
     * this RPC.
     * </pre>
     */
    public io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ClientArgs> runClient(
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ClientStatus> responseObserver) {
      return asyncUnimplementedStreamingCall(getRunClientMethodHelper(), responseObserver);
    }

    /**
     * <pre>
     * Just return the core count - unary call
     * </pre>
     */
    public void coreCount(io.grpc.benchmarks.proto.Control.CoreRequest request,
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.CoreResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getCoreCountMethodHelper(), responseObserver);
    }

    /**
     * <pre>
     * Quit this worker
     * </pre>
     */
    public void quitWorker(io.grpc.benchmarks.proto.Control.Void request,
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.Void> responseObserver) {
      asyncUnimplementedUnaryCall(getQuitWorkerMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getRunServerMethodHelper(),
            asyncBidiStreamingCall(
              new MethodHandlers<
                io.grpc.benchmarks.proto.Control.ServerArgs,
                io.grpc.benchmarks.proto.Control.ServerStatus>(
                  this, METHODID_RUN_SERVER)))
          .addMethod(
            getRunClientMethodHelper(),
            asyncBidiStreamingCall(
              new MethodHandlers<
                io.grpc.benchmarks.proto.Control.ClientArgs,
                io.grpc.benchmarks.proto.Control.ClientStatus>(
                  this, METHODID_RUN_CLIENT)))
          .addMethod(
            getCoreCountMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                io.grpc.benchmarks.proto.Control.CoreRequest,
                io.grpc.benchmarks.proto.Control.CoreResponse>(
                  this, METHODID_CORE_COUNT)))
          .addMethod(
            getQuitWorkerMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                io.grpc.benchmarks.proto.Control.Void,
                io.grpc.benchmarks.proto.Control.Void>(
                  this, METHODID_QUIT_WORKER)))
          .build();
    }
  }

  /**
   */
  public static final class WorkerServiceStub extends io.grpc.stub.AbstractStub<WorkerServiceStub> {
    private WorkerServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private WorkerServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WorkerServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new WorkerServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Start server with specified workload.
     * First request sent specifies the ServerConfig followed by ServerStatus
     * response. After that, a "Mark" can be sent anytime to request the latest
     * stats. Closing the stream will initiate shutdown of the test server
     * and once the shutdown has finished, the OK status is sent to terminate
     * this RPC.
     * </pre>
     */
    public io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ServerArgs> runServer(
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ServerStatus> responseObserver) {
      return asyncBidiStreamingCall(
          getChannel().newCall(getRunServerMethodHelper(), getCallOptions()), responseObserver);
    }

    /**
     * <pre>
     * Start client with specified workload.
     * First request sent specifies the ClientConfig followed by ClientStatus
     * response. After that, a "Mark" can be sent anytime to request the latest
     * stats. Closing the stream will initiate shutdown of the test client
     * and once the shutdown has finished, the OK status is sent to terminate
     * this RPC.
     * </pre>
     */
    public io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ClientArgs> runClient(
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ClientStatus> responseObserver) {
      return asyncBidiStreamingCall(
          getChannel().newCall(getRunClientMethodHelper(), getCallOptions()), responseObserver);
    }

    /**
     * <pre>
     * Just return the core count - unary call
     * </pre>
     */
    public void coreCount(io.grpc.benchmarks.proto.Control.CoreRequest request,
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.CoreResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCoreCountMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Quit this worker
     * </pre>
     */
    public void quitWorker(io.grpc.benchmarks.proto.Control.Void request,
        io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.Void> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getQuitWorkerMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class WorkerServiceBlockingStub extends io.grpc.stub.AbstractStub<WorkerServiceBlockingStub> {
    private WorkerServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private WorkerServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WorkerServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new WorkerServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Just return the core count - unary call
     * </pre>
     */
    public io.grpc.benchmarks.proto.Control.CoreResponse coreCount(io.grpc.benchmarks.proto.Control.CoreRequest request) {
      return blockingUnaryCall(
          getChannel(), getCoreCountMethodHelper(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Quit this worker
     * </pre>
     */
    public io.grpc.benchmarks.proto.Control.Void quitWorker(io.grpc.benchmarks.proto.Control.Void request) {
      return blockingUnaryCall(
          getChannel(), getQuitWorkerMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class WorkerServiceFutureStub extends io.grpc.stub.AbstractStub<WorkerServiceFutureStub> {
    private WorkerServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private WorkerServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WorkerServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new WorkerServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Just return the core count - unary call
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.grpc.benchmarks.proto.Control.CoreResponse> coreCount(
        io.grpc.benchmarks.proto.Control.CoreRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getCoreCountMethodHelper(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Quit this worker
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.grpc.benchmarks.proto.Control.Void> quitWorker(
        io.grpc.benchmarks.proto.Control.Void request) {
      return futureUnaryCall(
          getChannel().newCall(getQuitWorkerMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CORE_COUNT = 0;
  private static final int METHODID_QUIT_WORKER = 1;
  private static final int METHODID_RUN_SERVER = 2;
  private static final int METHODID_RUN_CLIENT = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final WorkerServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(WorkerServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CORE_COUNT:
          serviceImpl.coreCount((io.grpc.benchmarks.proto.Control.CoreRequest) request,
              (io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.CoreResponse>) responseObserver);
          break;
        case METHODID_QUIT_WORKER:
          serviceImpl.quitWorker((io.grpc.benchmarks.proto.Control.Void) request,
              (io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.Void>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_RUN_SERVER:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.runServer(
              (io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ServerStatus>) responseObserver);
        case METHODID_RUN_CLIENT:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.runClient(
              (io.grpc.stub.StreamObserver<io.grpc.benchmarks.proto.Control.ClientStatus>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class WorkerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    WorkerServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.grpc.benchmarks.proto.Services.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("WorkerService");
    }
  }

  private static final class WorkerServiceFileDescriptorSupplier
      extends WorkerServiceBaseDescriptorSupplier {
    WorkerServiceFileDescriptorSupplier() {}
  }

  private static final class WorkerServiceMethodDescriptorSupplier
      extends WorkerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    WorkerServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (WorkerServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new WorkerServiceFileDescriptorSupplier())
              .addMethod(getRunServerMethodHelper())
              .addMethod(getRunClientMethodHelper())
              .addMethod(getCoreCountMethodHelper())
              .addMethod(getQuitWorkerMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
