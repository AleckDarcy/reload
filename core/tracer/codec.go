package tracer

import (
	"context"
	"time"

	"github.com/AleckDarcy/reload/core/log"
)

type baseCodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

type codec struct {
	ctx   context.Context
	basic baseCodec
}

func NewCodec(ctx context.Context, basic baseCodec) *codec {
	return &codec{ctx: ctx, basic: basic}
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	if t, ok := v.(Tracer); ok {
		if idVal := c.ctx.Value(ThreadIDKey{}); idVal != nil {
			id := idVal.(int64)

			if Store.CheckByThreadID(id) {
				log.Logf("[RELOAD] Marshal, CheckByThreadID ok")
				record := &Record{
					Type:        RecordType_RecordSend,
					Timestamp:   time.Now().UnixNano(),
					MessageName: t.GetFI_Name(),
				}

				updateFunction := func(trace *Trace) {
					trace.Records = append(trace.Records, record)
				}
				if trace, ok := Store.UpdateFunctionByThreadID(id, updateFunction); ok {
					if t.GetMessageType() == MessageType_Message_Request {
						log.Logf("[RELOAD] Marshal send request")
						trace = &Trace{
							Id:      trace.Id,
							Records: []*Record{},
							Rlfi:    trace.Rlfi,
							Tfi:     trace.Tfi,
						}
					} else if t.GetMessageType() == MessageType_Message_Response {
						log.Logf("[RELOAD] Marshal send response")
						// todo: can not delete when the service is called concurrently
						Store.DeleteByTraceID(trace.Id)
						Store.DeleteByThreadID(id)
					}

					t.SetFI_Trace(trace)
					//log.Logf("thread %d trace %d start", id, trace.Id)
					//Store.IterateByThreadID(id, func(i int, record *Record) {
					//	log.Logf("thread %d trace %d, index %d: %s", id, trace.Id, i, record)
					//})
					//log.Logf("thread %d trace %d end", id, trace.Id)
				} else {
					log.Logf("[RELOAD] Marshal, UpdateFunctionByThreadID no")
				}
			} else {
				log.Logf("[RELOAD] Marshal, CheckByThreadID no")
			}
		} else {
			log.Logf("[RELOAD] Marshal, no ThreadIDKey")
		}
	}

	return c.basic.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	if err := c.basic.Unmarshal(data, v); err != nil {
		return err
	}

	if t, ok := v.(Tracer); ok {
		if idVal := c.ctx.Value(ThreadIDKey{}); idVal != nil {
			id := idVal.(int64)
			trace := t.GetFI_Trace()
			if trace != nil {
				trace.Records = append(trace.Records, &Record{
					Type:        RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: t.GetFI_Name(),
				})

				//log.Logf("thread %d trace %d start", id, trace.Id)
				//Store.IterateByThreadID(id, func(i int, record *Record) {
				//	log.Logf("thread %d trace %d, index %d: %s", id, trace.Id, i, record)
				//})
				//log.Logf("thread %d trace %d end", id, trace.Id)

				if err := trace.RLFI(); err != nil {
					return err
				} else if err = trace.TFI(); err != nil {
					return err
				}

				if t.GetMessageType() == MessageType_Message_Request {
					log.Logf("[RELOAD] Unmarshal receive request")
					Store.SetByThreadID(id, trace)
				} else if t.GetMessageType() == MessageType_Message_Response {
					log.Logf("[RELOAD] Unmarshal receive response")
					Store.UpdateFunctionByThreadID(id, func(oldTrace *Trace) {
						oldTrace.Records = append(oldTrace.Records, trace.Records...)
					})
				}
			} else {
				log.Logf("[RELOAD] Unmarshal, no tracing data")
			}
		} else {
			log.Logf("[RELOAD] Unmarshal, no ThreadIDKey")
		}
	}

	return nil
}

func (c *codec) Name() string {
	return "ProtoBuffer reloaded"
}
