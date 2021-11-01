// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3rpc

import (
	"context"
	"sync"
	"time"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"

	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft"

	"github.com/coreos/pkg/capnslog"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	maxNoLeaderCnt = 3
)

type streamsMap struct {
	mu      sync.Mutex
	streams map[grpc.ServerStream]struct{}
}

// 3milebeach todo: possible interposition point
func newUnaryInterceptor(s *etcdserver.EtcdServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !api.IsCapabilityEnabled(api.V3rpcCapability) {
			return nil, rpctypes.ErrGRPCNotCapable
		}

		if s.IsMemberExist(s.ID()) && s.IsLearner() && !isRPCSupportedForLearner(req) {
			return nil, rpctypes.ErrGPRCNotSupportedForLearner
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if ks := md[rpctypes.MetadataRequireLeaderKey]; len(ks) > 0 && ks[0] == rpctypes.MetadataHasLeader {
				if s.Leader() == types.ID(raft.None) {
					return nil, rpctypes.ErrGRPCNoLeader
				}
			}
		}

		if reqT, ok := req.(tracer.Tracer); ok { // 3milebeach begins
			if trace := reqT.GetFI_Trace(); trace != nil && trace.Id != 0 {
				records := trace.GetRecords()
				recordC := len(records)
				lastRecord := records[recordC-1]
				// log.Debug.PrintlnWithStackTrace(6, "%s last event %s", s.Cfg.ServerUUID, lastRecord) // 3milebeach: tracing

				// 1) read trace from request;
				event := &tracer.Record{
					Type:        tracer.RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: reqT.GetFI_Name(),
					Uuid:        lastRecord.GetUuid(),
					Service:     s.ServerID(),
				}
				// log.Debug.PrintlnWithStackTrace(6, "%s new event %s", s.Cfg.ServerUUID, event) // 3milebeach: tracing
				records = append(records, event)

				// 2) call handler();
				cm := tracer.NewContextMeta1(trace.Id, lastRecord.Uuid, reqT.GetFI_Name(), s.ServerID())
				ctx = tracer.NewContextWithContextMeta(ctx, cm)
				rsp, err := handler(ctx, req)

				// 3) append trace to rsp
				rspT := rsp.(tracer.Tracer)
				records = append(records, rspT.GetFI_Trace().GetRecords()...)

				event = &tracer.Record{
					Type:        tracer.RecordType_RecordSend,
					Timestamp:   time.Now().UnixNano(),
					MessageName: rsp.(tracer.Tracer).GetFI_Name(),
					Uuid:        lastRecord.GetUuid(),
					Service:     s.ServerID(),
				}
				records = append(records, event)
				// log.Debug.PrintlnWithStackTrace(6, "%s new event %s", s.Cfg.ServerUUID, event) // 3milebeach: tracing

				trace.Records = records
				rspT.SetFI_Trace(trace)

				log.Debug.PrintlnWithCaller("%s rsp: %s", s.ServerID(), log.Stringer.JSON(rsp)) // 3milebeach: tracing

				return rsp, err
			}
		} // 3milebeach ends

		return handler(ctx, req)
	}
}

func newLogUnaryInterceptor(s *etcdserver.EtcdServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		resp, err := handler(ctx, req)

		// log.Debug.PrintlnWithCaller("rsp: %s", resp) // 3milebeach: check server response

		lg := s.Logger()
		if (lg != nil && lg.Core().Enabled(zap.DebugLevel)) || // using zap logger and debug level is enabled
			(lg == nil && plog.LevelAt(capnslog.DEBUG)) { // or, using capnslog and debug level is enabled
			defer logUnaryRequestStats(ctx, lg, info, startTime, req, resp)
		}
		return resp, err
	}
}

// 3milebeach todo: possible interposition point
func logUnaryRequestStats(ctx context.Context, lg *zap.Logger, info *grpc.UnaryServerInfo, startTime time.Time, req interface{}, resp interface{}) {
	duration := time.Since(startTime)
	remote := "No remote client info."
	peerInfo, ok := peer.FromContext(ctx)
	if ok {
		remote = peerInfo.Addr.String()
	}
	responseType := info.FullMethod
	var reqCount, respCount int64
	var reqSize, respSize int
	var reqContent string
	switch _resp := resp.(type) {
	case *pb.RangeResponse:
		_req, ok := req.(*pb.RangeRequest)
		if ok {
			reqCount = 0
			reqSize = _req.Size()
			reqContent = _req.String()
		}
		if _resp != nil {
			respCount = _resp.GetCount()
			respSize = _resp.Size()
		}
	case *pb.PutResponse:
		_req, ok := req.(*pb.PutRequest)
		if ok {
			reqCount = 1
			reqSize = _req.Size()
			reqContent = pb.NewLoggablePutRequest(_req).String()
			// redact value field from request content, see PR #9821
		}
		if _resp != nil {
			respCount = 0
			respSize = _resp.Size()
		}
	case *pb.DeleteRangeResponse:
		_req, ok := req.(*pb.DeleteRangeRequest)
		if ok {
			reqCount = 0
			reqSize = _req.Size()
			reqContent = _req.String()
		}
		if _resp != nil {
			respCount = _resp.GetDeleted()
			respSize = _resp.Size()
		}
	case *pb.TxnResponse:
		_req, ok := req.(*pb.TxnRequest)
		if ok && _resp != nil {
			if _resp.GetSucceeded() { // determine the 'actual' count and size of request based on success or failure
				reqCount = int64(len(_req.GetSuccess()))
				reqSize = 0
				for _, r := range _req.GetSuccess() {
					reqSize += r.Size()
				}
			} else {
				reqCount = int64(len(_req.GetFailure()))
				reqSize = 0
				for _, r := range _req.GetFailure() {
					reqSize += r.Size()
				}
			}
			reqContent = pb.NewLoggableTxnRequest(_req).String()
			// redact value field from request content, see PR #9821
		}
		if _resp != nil {
			respCount = 0
			respSize = _resp.Size()
		}
	default:
		reqCount = -1
		reqSize = -1
		respCount = -1
		respSize = -1
	}

	logGenericRequestStats(lg, startTime, duration, remote, responseType, reqCount, reqSize, respCount, respSize, reqContent)
}

func logGenericRequestStats(lg *zap.Logger, startTime time.Time, duration time.Duration, remote string, responseType string,
	reqCount int64, reqSize int, respCount int64, respSize int, reqContent string) {
	if lg == nil {
		plog.Debugf("start time = %v, "+
			"time spent = %v, "+
			"remote = %s, "+
			"response type = %s, "+
			"request count = %d, "+
			"request size = %d, "+
			"response count = %d, "+
			"response size = %d, "+
			"request content = %s",
			startTime, duration, remote, responseType, reqCount, reqSize, respCount, respSize, reqContent,
		)
	} else {
		lg.Debug("request stats",
			zap.Time("start time", startTime),
			zap.Duration("time spent", duration),
			zap.String("remote", remote),
			zap.String("response type", responseType),
			zap.Int64("request count", reqCount),
			zap.Int("request size", reqSize),
			zap.Int64("response count", respCount),
			zap.Int("response size", respSize),
			zap.String("request content", reqContent),
		)
	}
}

func newStreamInterceptor(s *etcdserver.EtcdServer) grpc.StreamServerInterceptor {
	smap := monitorLeader(s)

	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !api.IsCapabilityEnabled(api.V3rpcCapability) {
			return rpctypes.ErrGRPCNotCapable
		}

		if s.IsMemberExist(s.ID()) && s.IsLearner() { // learner does not support stream RPC
			return rpctypes.ErrGPRCNotSupportedForLearner
		}

		md, ok := metadata.FromIncomingContext(ss.Context())
		if ok {
			if ks := md[rpctypes.MetadataRequireLeaderKey]; len(ks) > 0 && ks[0] == rpctypes.MetadataHasLeader {
				if s.Leader() == types.ID(raft.None) {
					return rpctypes.ErrGRPCNoLeader
				}

				cctx, cancel := context.WithCancel(ss.Context())
				ss = serverStreamWithCtx{ctx: cctx, cancel: &cancel, ServerStream: ss}

				smap.mu.Lock()
				smap.streams[ss] = struct{}{}
				smap.mu.Unlock()

				defer func() {
					smap.mu.Lock()
					delete(smap.streams, ss)
					smap.mu.Unlock()
					cancel()
				}()

			}
		}

		return handler(srv, ss)
	}
}

type serverStreamWithCtx struct {
	grpc.ServerStream
	ctx    context.Context
	cancel *context.CancelFunc
}

func (ssc serverStreamWithCtx) Context() context.Context { return ssc.ctx }

func monitorLeader(s *etcdserver.EtcdServer) *streamsMap {
	smap := &streamsMap{
		streams: make(map[grpc.ServerStream]struct{}),
	}

	go func() {
		election := time.Duration(s.Cfg.TickMs) * time.Duration(s.Cfg.ElectionTicks) * time.Millisecond
		noLeaderCnt := 0

		for {
			select {
			case <-s.StopNotify():
				return
			case <-time.After(election):
				if s.Leader() == types.ID(raft.None) {
					noLeaderCnt++
				} else {
					noLeaderCnt = 0
				}

				// We are more conservative on canceling existing streams. Reconnecting streams
				// cost much more than just rejecting new requests. So we wait until the member
				// cannot find a leader for maxNoLeaderCnt election timeouts to cancel existing streams.
				if noLeaderCnt >= maxNoLeaderCnt {
					smap.mu.Lock()
					for ss := range smap.streams {
						if ssWithCtx, ok := ss.(serverStreamWithCtx); ok {
							(*ssWithCtx.cancel)()
							<-ss.Context().Done()
						}
					}
					smap.streams = make(map[grpc.ServerStream]struct{})
					smap.mu.Unlock()
				}
			}
		}
	}()

	return smap
}
