package tracer

import (
	"context"
	"time"
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
			trace := Store.UpdateFunctionByThreadID(id, func(trace *Trace) {
				trace.Records = append(trace.Records, &Record{
					Type:        RecordType_RecordSend,
					Timestamp:   time.Now().UnixNano(),
					MessageName: t.GetFI_Name(),
				})
				trace.Depth = int64(len(trace.Records))
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
					Timestamp:   time.Now().UnixNano(),
					MessageName: t.GetFI_Name(),
				})
				trace.Depth = int64(len(trace.Records))

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
			}
		}
	}

	return nil
}

func (c *codec) Name() string {
	return "ProtoBuffer reloaded"
}
