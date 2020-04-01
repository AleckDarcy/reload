package io.grpc.testing.integration;

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
 * <pre>
 * A service used to control reconnect server.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler",
    comments = "Source: io/grpc/testing/integration/test.proto")
public final class ReconnectServiceGrpc {

  private ReconnectServiceGrpc() {}

  public static final String SERVICE_NAME = "grpc.testing.ReconnectService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getStartMethod()} instead. 
  public static final io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      com.google.protobuf.EmptyProtos.Empty> METHOD_START = getStartMethodHelper();

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      com.google.protobuf.EmptyProtos.Empty> getStartMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      com.google.protobuf.EmptyProtos.Empty> getStartMethod() {
    return getStartMethodHelper();
  }

  private static io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      com.google.protobuf.EmptyProtos.Empty> getStartMethodHelper() {
    io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty, com.google.protobuf.EmptyProtos.Empty> getStartMethod;
    if ((getStartMethod = ReconnectServiceGrpc.getStartMethod) == null) {
      synchronized (ReconnectServiceGrpc.class) {
        if ((getStartMethod = ReconnectServiceGrpc.getStartMethod) == null) {
          ReconnectServiceGrpc.getStartMethod = getStartMethod = 
              io.grpc.MethodDescriptor.<com.google.protobuf.EmptyProtos.Empty, com.google.protobuf.EmptyProtos.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "grpc.testing.ReconnectService", "Start"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.EmptyProtos.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.EmptyProtos.Empty.getDefaultInstance()))
                  .setSchemaDescriptor(new ReconnectServiceMethodDescriptorSupplier("Start"))
                  .build();
          }
        }
     }
     return getStartMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getStopMethod()} instead. 
  public static final io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      io.grpc.testing.integration.Messages.ReconnectInfo> METHOD_STOP = getStopMethodHelper();

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      io.grpc.testing.integration.Messages.ReconnectInfo> getStopMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      io.grpc.testing.integration.Messages.ReconnectInfo> getStopMethod() {
    return getStopMethodHelper();
  }

  private static io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty,
      io.grpc.testing.integration.Messages.ReconnectInfo> getStopMethodHelper() {
    io.grpc.MethodDescriptor<com.google.protobuf.EmptyProtos.Empty, io.grpc.testing.integration.Messages.ReconnectInfo> getStopMethod;
    if ((getStopMethod = ReconnectServiceGrpc.getStopMethod) == null) {
      synchronized (ReconnectServiceGrpc.class) {
        if ((getStopMethod = ReconnectServiceGrpc.getStopMethod) == null) {
          ReconnectServiceGrpc.getStopMethod = getStopMethod = 
              io.grpc.MethodDescriptor.<com.google.protobuf.EmptyProtos.Empty, io.grpc.testing.integration.Messages.ReconnectInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "grpc.testing.ReconnectService", "Stop"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.EmptyProtos.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.testing.integration.Messages.ReconnectInfo.getDefaultInstance()))
                  .setSchemaDescriptor(new ReconnectServiceMethodDescriptorSupplier("Stop"))
                  .build();
          }
        }
     }
     return getStopMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ReconnectServiceStub newStub(io.grpc.Channel channel) {
    return new ReconnectServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ReconnectServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new ReconnectServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ReconnectServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new ReconnectServiceFutureStub(channel);
  }

  /**
   * <pre>
   * A service used to control reconnect server.
   * </pre>
   */
  public static abstract class ReconnectServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void start(com.google.protobuf.EmptyProtos.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.EmptyProtos.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(getStartMethodHelper(), responseObserver);
    }

    /**
     */
    public void stop(com.google.protobuf.EmptyProtos.Empty request,
        io.grpc.stub.StreamObserver<io.grpc.testing.integration.Messages.ReconnectInfo> responseObserver) {
      asyncUnimplementedUnaryCall(getStopMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getStartMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.EmptyProtos.Empty,
                com.google.protobuf.EmptyProtos.Empty>(
                  this, METHODID_START)))
          .addMethod(
            getStopMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.EmptyProtos.Empty,
                io.grpc.testing.integration.Messages.ReconnectInfo>(
                  this, METHODID_STOP)))
          .build();
    }
  }

  /**
   * <pre>
   * A service used to control reconnect server.
   * </pre>
   */
  public static final class ReconnectServiceStub extends io.grpc.stub.AbstractStub<ReconnectServiceStub> {
    private ReconnectServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ReconnectServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReconnectServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ReconnectServiceStub(channel, callOptions);
    }

    /**
     */
    public void start(com.google.protobuf.EmptyProtos.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.EmptyProtos.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getStartMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void stop(com.google.protobuf.EmptyProtos.Empty request,
        io.grpc.stub.StreamObserver<io.grpc.testing.integration.Messages.ReconnectInfo> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getStopMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * A service used to control reconnect server.
   * </pre>
   */
  public static final class ReconnectServiceBlockingStub extends io.grpc.stub.AbstractStub<ReconnectServiceBlockingStub> {
    private ReconnectServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ReconnectServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReconnectServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ReconnectServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.google.protobuf.EmptyProtos.Empty start(com.google.protobuf.EmptyProtos.Empty request) {
      return blockingUnaryCall(
          getChannel(), getStartMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public io.grpc.testing.integration.Messages.ReconnectInfo stop(com.google.protobuf.EmptyProtos.Empty request) {
      return blockingUnaryCall(
          getChannel(), getStopMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * A service used to control reconnect server.
   * </pre>
   */
  public static final class ReconnectServiceFutureStub extends io.grpc.stub.AbstractStub<ReconnectServiceFutureStub> {
    private ReconnectServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ReconnectServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReconnectServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ReconnectServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.EmptyProtos.Empty> start(
        com.google.protobuf.EmptyProtos.Empty request) {
      return futureUnaryCall(
          getChannel().newCall(getStartMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.grpc.testing.integration.Messages.ReconnectInfo> stop(
        com.google.protobuf.EmptyProtos.Empty request) {
      return futureUnaryCall(
          getChannel().newCall(getStopMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_START = 0;
  private static final int METHODID_STOP = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ReconnectServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ReconnectServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_START:
          serviceImpl.start((com.google.protobuf.EmptyProtos.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.EmptyProtos.Empty>) responseObserver);
          break;
        case METHODID_STOP:
          serviceImpl.stop((com.google.protobuf.EmptyProtos.Empty) request,
              (io.grpc.stub.StreamObserver<io.grpc.testing.integration.Messages.ReconnectInfo>) responseObserver);
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
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class ReconnectServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ReconnectServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.grpc.testing.integration.Test.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ReconnectService");
    }
  }

  private static final class ReconnectServiceFileDescriptorSupplier
      extends ReconnectServiceBaseDescriptorSupplier {
    ReconnectServiceFileDescriptorSupplier() {}
  }

  private static final class ReconnectServiceMethodDescriptorSupplier
      extends ReconnectServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ReconnectServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (ReconnectServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ReconnectServiceFileDescriptorSupplier())
              .addMethod(getStartMethodHelper())
              .addMethod(getStopMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
