package tracer

import (
	"context"
	"errors"
	"time"

	"github.com/AleckDarcy/reload/core/log"
)

type baseCodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

type codec struct {
	serverUUID UUID
	ctx        context.Context
	basic      baseCodec
}

func NewCodec(ctx context.Context, basic baseCodec) baseCodec {
	if obj := ctx.Value(ContextMetaKey{}); obj != nil {
		return &codec{ctx: ctx, basic: basic, serverUUID: obj.(*ContextMeta).server}
	}

	return basic
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	if t, ok := v.(Tracer); ok {
		if metaVal := c.ctx.Value(ContextMetaKey{}); metaVal != nil {
			meta := metaVal.(*ContextMeta)

			if Store.CheckByContextMeta(meta) {
				//log.Logf("[RELOAD] Marshal, %s, CheckByContextMeta ok", t.GetFI_Name())

				var uuid string
				if t.GetFI_MessageType() == MessageType_Message_Request {
					uuid = NewUUID()
				} else if t.GetFI_MessageType() == MessageType_Message_Response {
					uuid = meta.uuid
				}

				record := &Record{
					Type:        RecordType_RecordSend,
					Timestamp:   time.Now().UnixNano(),
					MessageName: t.GetFI_Name(),
					Uuid:        uuid,
					Service:     c.serverUUID,
				}

				updateFunction := func(trace *Trace) {
					trace.Records = append(trace.Records, record)
				}
				if trace, ok := Store.UpdateFunctionByContextMeta(meta, updateFunction); ok {
					if t.GetFI_MessageType() == MessageType_Message_Request {
						if tfis := trace.Tfis; tfis != nil {
							crash, found := true, false
							for _, tfi := range tfis {
								if tfi.Type == FaultType_FaultCrash {
									if found = tfi.Name[0] == t.GetFI_Name(); found {
										for _, after := range tfi.After {
											if after.Times != -1 && after.Already != after.Times {
												crash = false
												break
											}
										}

										break
									}
								}
							}

							if crash && found {
								//log.Logf("[RELOAD] Marshal tfi crash triggered")
								for _, tfi := range tfis {
									for _, after := range tfi.After {
										if after.Name == t.GetFI_Name() {
											after.Already++
										}
									}
								}

								return nil, errors.New("transport is closing")
							}
						}

						//log.Logf("[RELOAD] Marshal send request")

						trace = &Trace{
							Id:      trace.Id,
							Records: []*Record{record},
							Rlfis:   trace.Rlfis,
							Tfis:    trace.Tfis,
						}
					} else if t.GetFI_MessageType() == MessageType_Message_Response {
						//log.Logf("[RELOAD] Marshal send response")
						Store.DeleteByContextMeta(meta)
					}

					t.SetFI_Trace(trace)
				} else {
					log.Logf("[RELOAD] Marshal, UpdateFunctionByContextMeta fail")
				}
			} else {
				//log.Logf("[RELOAD] Marshal, CheckByContextMeta fail")
			}
		} else {
			//log.Logf("[RELOAD] Marshal, %s, no ContextMetaKey", t.GetFI_Name())
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

				if t.GetFI_MessageType() == MessageType_Message_Request {
					//log.Logf("[RELOAD] Unmarshal, receive request %s", t.GetFI_Name())
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
							Service:     c.serverUUID,
						}
						Store.SetByContextMeta(meta, trace)
					}
				} else if t.GetFI_MessageType() == MessageType_Message_Response {
					//log.Logf("[RELOAD] Unmarshal, receive response %s", t.GetFI_Name())
					if len(trace.Records) == 0 {
						//log.Logf("[RELOAD] Unmarshal, receive empty trace")
					} else if uuid := trace.Records[0].Uuid; uuid == "" {
						log.Logf("[RELOAD] Unmarshal, receive invalid uuid: %s", uuid)
					} else {
						//Store.UpdateFunctionByContextMeta(meta, func(oldTrace *Trace) {
						//	length := len(trace.Records) + 1
						//	oldTrace.Records = append(oldTrace.Records, trace.Records...)
						//	oldTrace.Records = append(oldTrace.Records, &Record{
						//		Type:        RecordType_RecordReceive,
						//		Timestamp:   time.Now().UnixNano(),
						//		MessageName: t.GetFI_Name(),
						//		Uuid:        uuid,
						//		Service:     ServiceUUID,
						//	})
						//	oldTrace.CalFI(oldTrace.Records[len(oldTrace.Records)-length:])
						//})

						trace.Records = append(trace.Records, &Record{
							Type:        RecordType_RecordReceive,
							Timestamp:   time.Now().UnixNano(),
							MessageName: t.GetFI_Name(),
							Uuid:        uuid,
							Service:     c.serverUUID,
						})

						t.SetFI_Trace(trace)
					}
				}
			} else {
				//log.Logf("[RELOAD] Unmarshal, %s, no trace", t.GetFI_Name())
			}
		} else {
			//log.Logf("[RELOAD] Unmarshal, no ContextMetaKey")
		}
	}

	return nil
}

func (c *codec) Name() string {
	return "ProtoBuffer reloaded"
}
