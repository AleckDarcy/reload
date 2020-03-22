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
		if metaVal := c.ctx.Value(ContextMetaKey{}); metaVal != nil {
			meta := metaVal.(*ContextMeta)

			if Store.CheckByContextMeta(meta) {
				log.Logf("[RELOAD] Marshal, %s, CheckByContextMeta ok", t.GetFI_Name())

				var uuid string
				if t.GetMessageType() == MessageType_Message_Request {
					uuid = NewUUID()
				} else if t.GetMessageType() == MessageType_Message_Response {
					uuid = meta.uuid
				}

				record := &Record{
					Type:        RecordType_RecordSend,
					Timestamp:   time.Now().UnixNano(),
					MessageName: t.GetFI_Name(),
					Uuid:        uuid,
				}

				updateFunction := func(trace *Trace) {
					trace.Records = append(trace.Records, record)
				}
				if trace, ok := Store.UpdateFunctionByContextMeta(meta, updateFunction); ok {
					if t.GetMessageType() == MessageType_Message_Request {
						log.Logf("[RELOAD] Marshal send request")
						trace = &Trace{
							Id:      trace.Id,
							Records: []*Record{record},
							Rlfi:    trace.Rlfi,
							Tfi:     trace.Tfi,
						}
					} else if t.GetMessageType() == MessageType_Message_Response {
						log.Logf("[RELOAD] Marshal send response")
						Store.DeleteByContextMeta(meta)
					}

					t.SetFI_Trace(trace)
				} else {
					log.Logf("[RELOAD] Marshal, UpdateFunctionByContextMeta fail")
				}
			} else {
				log.Logf("[RELOAD] Marshal, CheckByContextMeta fail")
			}
		} else {
			log.Logf("[RELOAD] Marshal, %s, no ContextMetaKey", t.GetFI_Name())
		}
	}

	return c.basic.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	if err := c.basic.Unmarshal(data, v); err != nil {
		return err
	}

	if t, ok := v.(Tracer); ok {
		if metaVal := c.ctx.Value(ContextMetaKey{}); metaVal != nil {
			meta := metaVal.(*ContextMeta)

			trace := t.GetFI_Trace()
			if trace != nil {
				if err := trace.DoFI(t.GetFI_Name()); err != nil {
					return err
				}

				if t.GetMessageType() == MessageType_Message_Request {
					log.Logf("[RELOAD] Unmarshal, receive request %s", t.GetFI_Name())
					if len(trace.Records) != 1 {
						log.Logf("[RELOAD] Unmarshal, receive invalid trace: %s", trace.JSONString())
					} else if uuid := trace.Records[0].Uuid; uuid == "" {
						log.Logf("[RELOAD] Unmarshal, receive invalid uuid: %s", uuid)
					} else {
						meta.traceID = trace.Id
						meta.uuid = uuid

						trace.Records[0] = &Record{
							Type:        RecordType_RecordReceive,
							Timestamp:   time.Now().UnixNano(),
							MessageName: t.GetFI_Name(),
							Uuid:        uuid,
						}
						Store.SetByContextMeta(meta, trace)
					}
				} else if t.GetMessageType() == MessageType_Message_Response {
					log.Logf("[RELOAD] Unmarshal, receive response %s", t.GetFI_Name())
					if len(trace.Records) == 0 {
						log.Logf("[RELOAD] Unmarshal, receive empty trace")
					} else if uuid := trace.Records[0].Uuid; uuid == "" {
						log.Logf("[RELOAD] Unmarshal, receive invalid uuid: %s", uuid)
					} else {
						Store.UpdateFunctionByContextMeta(meta, func(oldTrace *Trace) {
							length := len(trace.Records) + 1
							oldTrace.Records = append(oldTrace.Records, trace.Records...)
							oldTrace.Records = append(oldTrace.Records, &Record{
								Type:        RecordType_RecordReceive,
								Timestamp:   time.Now().UnixNano(),
								MessageName: t.GetFI_Name(),
								Uuid:        uuid,
							})
							oldTrace.CalFI(oldTrace.Records[len(oldTrace.Records)-length:])
						})

						t.SetFI_Trace(nil)
					}
				}
			} else {
				log.Logf("[RELOAD] Unmarshal, %s, no trace", t.GetFI_Name())
			}
		} else {
			log.Logf("[RELOAD] Unmarshal, no ContextMetaKey")
		}
	}

	return nil
}

func (c *codec) Name() string {
	return "ProtoBuffer reloaded"
}
