package tracer

import (
	"context"
	"time"
)

type messageNameIDMap struct {
	name2ID map[string]string
	id2Name map[string]string
}

var MessageNameIDMap *messageNameIDMap

type baseCodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

type compressCode interface {
	baseCodec

	MessageType() MessageType
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
			trace := Store.UpdateFunctionByThreadID(id, func(trace *Trace) {
				trace.Records = append(trace.Records, &Record{
					Type:        RecordType_RecordSend,
					Timestamp:   time.Now().UnixNano() - trace.BaseTimestamp,
					MessageName: t.GetFI_Name(),
				})
			})

			t.SetFI_Trace(trace)

			//log.Logf("thread %d trace %d start", id, trace.Id)
			//Store.IterateByThreadID(id, func(i int, record *Record) {
			//	log.Logf("thread %d trace %d, index %d: %s", id, trace.Id, i, record)
			//})
			//log.Logf("thread %d trace %d end", id, trace.Id)
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
					Timestamp:   time.Now().UnixNano() - trace.BaseTimestamp,
					MessageName: t.GetFI_Name(),
				})

				Store.SetByThreadID(id, trace)

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

				t.SetFI_Trace(nil)
			}
		}
	}

	return nil
}

func (c *codec) Name() string {
	return "ProtoBuffer reloaded"
}
