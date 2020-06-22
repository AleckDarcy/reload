package storage

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/core/tracer"
	"github.com/gogo/protobuf/proto"
)

type baseCodec struct {
}

func (c *baseCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (c *baseCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (c *baseCodec) Name() string {
	return "test"
}

const optOn = false

type pool struct {
	trace *sync.Pool
}

var p = &pool{
	trace: &sync.Pool{
		New: func() interface{} {
			return &tracer.Trace{}
		},
	},
}

func getTrace() *tracer.Trace {
	if optOn {
		return p.trace.Get().(*tracer.Trace)
	}

	return &tracer.Trace{}
}

func putTrace(t *tracer.Trace) {
	if optOn {
		t.Records = nil
		p.trace.Put(t)
	}
}

func BenchmarkName(b *testing.B) {
	a := byte(0)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000000; j++ {
			a += byte(j)
		}
	}
}

func TestName(t *testing.T) {
	serviceA := tracer.NewUUID()
	serviceB := tracer.NewUUID()

	nThread := 1
	nTest := 100

	signal := make(chan struct{}, nThread)

	start := time.Now()

	traceID := int64(0)

	for i := 0; i < (1+nThread)*14; i++ {
		go func() {
			bytes := make([]byte, 5000)
			for {
				str := string(bytes)
				bytes = []byte(str)

				for j := 0; j < 10000000; j++ {
					bytes[0] += byte(j)
				}
				newBytes := make([]byte, len(bytes))
				newBytes = append(newBytes, bytes...)
				newBytes = append(newBytes, bytes...)

				runtime.Gosched()
			}
		}()
	}

	for i := 0; i < 4; i++ {
		go func() {
			bytes := make([]byte, 1000)
			for {
				str := string(bytes)
				bytes = []byte(str)
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	for i := 0; i < nThread; i++ {
		go func(signal chan struct{}) {
			for j := 0; j < nTest/nThread; j++ {
				trace := getTrace()
				trace.Id = atomic.AddInt64(&traceID, 1)

				for k := 0; k < 14; k++ {
					uuid := tracer.NewUUID()
					trace.Records = append(trace.Records, &tracer.Record{
						Type:        tracer.RecordType_RecordSend,
						Timestamp:   time.Now().UnixNano(),
						MessageName: "Request",
						Uuid:        uuid,
						Service:     serviceA,
					})

					tracer.Store.SetByContextMeta(tracer.NewContextMeta(traceID, uuid, "Request"), trace)

					requestT := &tracer.Trace{
						Records: []*tracer.Record{
							{
								Type:        tracer.RecordType_RecordSend,
								Timestamp:   time.Now().UnixNano(),
								MessageName: "Request",
								Uuid:        tracer.NewUUID(),
								Service:     serviceA,
							},
						},
					}
					bytes, _ := tracer.NewCodec(nil, &baseCodec{}).Marshal(requestT)

					requestT = &tracer.Trace{}
					tracer.NewCodec(nil, &baseCodec{}).Unmarshal(bytes, requestT)

					responseT := &tracer.Trace{
						Records: []*tracer.Record{
							{
								Type:        tracer.RecordType_RecordReceive,
								Timestamp:   time.Now().UnixNano(),
								MessageName: "Request",
								Uuid:        requestT.Records[0].Uuid,
								Service:     serviceB,
							},
							{
								Type:        tracer.RecordType_RecordSend,
								Timestamp:   time.Now().UnixNano(),
								MessageName: "Response",
								Uuid:        requestT.Records[0].Uuid,
								Service:     serviceB,
							},
						},
					}
					bytes, _ = proto.Marshal(responseT)

					responseT = &tracer.Trace{}
					tracer.NewCodec(nil, &baseCodec{}).Unmarshal(bytes, responseT)

					tracer.Store.UpdateFunctionByContextMeta(
						tracer.NewContextMeta(traceID, uuid, "Request"),
						func(trace *tracer.Trace) {
							trace.Records = append(trace.Records, responseT.Records...)

							trace.Records = append(trace.Records, &tracer.Record{
								Type:        tracer.RecordType_RecordReceive,
								Timestamp:   time.Now().UnixNano(),
								MessageName: "Response",
								Uuid:        tracer.NewUUID(),
								Service:     serviceA,
							})
						},
					)

					tracer.Store.DeleteByContextMeta(tracer.NewContextMeta(traceID, uuid, "Request"))

				}

				bytes, _ := proto.Marshal(trace)

				putTrace(trace)

				trace = getTrace()
				tracer.NewCodec(nil, &baseCodec{}).Unmarshal(bytes, trace)

				json.Marshal(trace)

				fmt.Println(trace)
				putTrace(trace)
			}

			signal <- struct{}{}
		}(signal)
	}

	for i := 0; i < nThread; i++ {
		<-signal
	}

	end := time.Now()
	t.Log(end.Sub(start))
}
