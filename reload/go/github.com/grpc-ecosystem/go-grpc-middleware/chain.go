// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

// gRPC Server Interceptor chaining middleware.

package grpc_middleware

import (
	"time"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ChainUnaryServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
// will see context changes of one and two.
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			var (
				chainHandler grpc.UnaryHandler
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				// log.CriticalPath.PrintlnWithCaller("stub") // 3milebeach begins

				var trace *tracer.Trace
				var lastEvent *tracer.Record
				cm := ctx.Value(tracer.ContextMetaKey{}).(*tracer.ContextMeta)

				if reqT, ok := req.(tracer.Tracer); ok {
					trace, ok = tracer.Assertion.GetTrace(req)
					if ok {
						//log.Debug.PrintlnWithStackTrace(3, "hhhhhhh")
						//log.Debug.PrintlnWithStackTrace(4, "hhhhhhh")
						//log.Debug.PrintlnWithStackTrace(5, "hhhhhhh")
						//log.Debug.PrintlnWithStackTrace(6, "hhhhhhh")
						//log.Debug.PrintlnWithStackTrace(7, "hhhhhhh")
						//log.Debug.PrintlnWithStackTrace(8, "hhhhhhh")

						trace = reqT.GetFI_Trace()
						if lastEvent, ok = tracer.Assertion.GetLastEvent(reqT); ok {
							log.Debug.PrintlnWithCaller("%s processing %s %s", cm.ServerUUID(), reqT.GetFI_Name(), lastEvent.GetMessageName())
						} else {
							log.Error.PrintlnWithCaller("%s processing %s without events", cm.ServerUUID(), reqT.GetFI_Name())
						}
					}
				} // 3milebeach ends

				if curI == lastI {
					resp, err := handler(currentCtx, currentReq) // 3milebeach begins
					if err == nil && trace != nil {
						if respT, ok := resp.(tracer.Tracer); ok {
							if trace2 := respT.GetFI_Trace(); trace2 != nil {
								trace.AppendRecords(trace2.GetRecords())

								event := &tracer.Record{
									Type:        tracer.RecordType_RecordReceive,
									Timestamp:   time.Now().UnixNano(),
									MessageName: respT.GetFI_Name(),
									Uuid:        lastEvent.Uuid,
									Service:     cm.ServerUUID(),
								}

								trace.AppendRecord(event)
							}

							respT.SetFI_Trace(trace)

							log.Debug.PrintlnWithCaller("ChainUnaryServer() rsp: %s", log.Stringer.JSON(resp))
						}
					}

					return resp, err // 3milebeach ends

					// return handler(currentCtx, currentReq)
				}
				curI++
				resp, err := interceptors[curI](currentCtx, currentReq, info, chainHandler)
				curI--

				if err == nil && trace != nil { // 3milebeach begins
					if respT, ok := resp.(tracer.Tracer); ok {
						if trace2 := respT.GetFI_Trace(); trace2 != nil {
							trace.AppendRecords(trace2.GetRecords())

							event := &tracer.Record{
								Type:        tracer.RecordType_RecordReceive,
								Timestamp:   time.Now().UnixNano(),
								MessageName: respT.GetFI_Name(),
								Uuid:        lastEvent.Uuid,
								Service:     cm.ServerUUID(),
							}

							trace.AppendRecord(event)
						}

						respT.SetFI_Trace(trace)

						log.Debug.PrintlnWithCaller("ChainUnaryServer() rsp: %s", resp)
					}
				} // 3milebeach ends

				return resp, err
			}

			return interceptors[0](ctx, req, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}

// ChainStreamServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three.
// If you want to pass context between interceptors, use WrapServerStream.
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			var (
				chainHandler grpc.StreamHandler
				curI         int
			)

			chainHandler = func(currentSrv interface{}, currentStream grpc.ServerStream) error {
				if curI == lastI {
					return handler(currentSrv, currentStream)
				}
				curI++
				err := interceptors[curI](currentSrv, currentStream, info, chainHandler)
				curI--
				return err
			}

			return interceptors[0](srv, stream, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, stream)
	}
}

// ChainUnaryClient creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryClient(one, two, three) will execute one before two before three.
func ChainUnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			var (
				chainHandler grpc.UnaryInvoker
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
				if curI == lastI {
					return invoker(currentCtx, currentMethod, currentReq, currentRepl, currentConn, currentOpts...)
				}
				curI++
				err := interceptors[curI](currentCtx, currentMethod, currentReq, currentRepl, currentConn, chainHandler, currentOpts...)
				curI--
				return err
			}

			return interceptors[0](ctx, method, req, reply, cc, chainHandler, opts...)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// ChainStreamClient creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainStreamClient(one, two, three) will execute one before two before three.
func ChainStreamClient(interceptors ...grpc.StreamClientInterceptor) grpc.StreamClientInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			var (
				chainHandler grpc.Streamer
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentDesc *grpc.StreamDesc, currentConn *grpc.ClientConn, currentMethod string, currentOpts ...grpc.CallOption) (grpc.ClientStream, error) {
				if curI == lastI {
					return streamer(currentCtx, currentDesc, currentConn, currentMethod, currentOpts...)
				}
				curI++
				stream, err := interceptors[curI](currentCtx, currentDesc, currentConn, currentMethod, chainHandler, currentOpts...)
				curI--
				return stream, err
			}

			return interceptors[0](ctx, desc, cc, method, chainHandler, opts...)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(ctx, desc, cc, method, opts...)
	}
}

// Chain creates a single interceptor out of a chain of many interceptors.
//
// WithUnaryServerChain is a grpc.Server config option that accepts multiple unary interceptors.
// Basically syntactic sugar.
func WithUnaryServerChain(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(ChainUnaryServer(interceptors...))
}

// WithStreamServerChain is a grpc.Server config option that accepts multiple stream interceptors.
// Basically syntactic sugar.
func WithStreamServerChain(interceptors ...grpc.StreamServerInterceptor) grpc.ServerOption {
	return grpc.StreamInterceptor(ChainStreamServer(interceptors...))
}
