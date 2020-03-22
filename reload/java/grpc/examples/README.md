grpc Examples
==============================================

The examples require grpc-java to already be built. You are strongly encouraged
to check out a git release tag, since there will already be a build of grpc
available. Otherwise you must follow [COMPILING](../COMPILING.md).

You may want to read through the
[Quick Start Guide](https://grpc.io/docs/quickstart/java.html)
before trying out the examples.

To build the examples, run in this directory:

```
$ ./gradlew installDist
```

This creates the scripts `hello-world-server`, `hello-world-client`,
`route-guide-server`, and `route-guide-client` in the
`build/install/examples/bin/` directory that run the examples. Each
example requires the server to be running before starting the client.

For example, to try the hello world example first run:

```
$ ./build/install/examples/bin/hello-world-server
```

And in a different terminal window run:

```
$ ./build/install/examples/bin/hello-world-client
```

That's it!

Please refer to gRPC Java's [README](../README.md) and
[tutorial](https://grpc.io/docs/tutorials/basic/java.html) for more
information.

## Maven

If you prefer to use Maven:
```
$ mvn verify
$ # Run the server
$ mvn exec:java -Dexec.mainClass=io.grpc.examples.helloworld.HelloWorldServer
$ # In another terminal run the client
$ mvn exec:java -Dexec.mainClass=io.grpc.examples.helloworld.HelloWorldClient
```

## Bazel

If you prefer to use Bazel:
```
(With Bazel v0.8.0 or above.)
$ bazel build :hello-world-server :hello-world-client
$ # Run the server:
$ bazel-bin/hello-world-server
$ # In another terminal run the client
$ bazel-bin/hello-world-client
```

Unit test examples
==============================================

Examples for unit testing gRPC clients and servers are located in [examples/src/test](src/test).

In general, we DO NOT allow overriding the client stub.
We encourage users to leverage `InProcessTransport` as demonstrated in the examples to
write unit tests. `InProcessTransport` is light-weight and runs the server
and client in the same process without any socket/TCP connection.

For testing a gRPC client, create the client with a real stub
using an
[InProcessChannel](../core/src/main/java/io/grpc/inprocess/InProcessChannelBuilder.java),
and test it against an
[InProcessServer](../core/src/main/java/io/grpc/inprocess/InProcessServerBuilder.java)
with a mock/fake service implementation.

For testing a gRPC server, create the server as an InProcessServer,
and test it against a real client stub with an InProcessChannel.

The gRPC-java library also provides a JUnit rule,
[GrpcServerRule](../testing/src/main/java/io/grpc/testing/GrpcServerRule.java), to do the starting
up and shutting down boilerplate for you.
