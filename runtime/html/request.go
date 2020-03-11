package html

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"
)

// called at the very beginning of handler()
func Init(r *http.Request) *http.Request {
	log.Logf("[RELOAD] Init, Fi-Trace: %s", r.Header.Get("Fi-Trace"))
	if traceUrlStr := r.Header.Get("Fi-Trace"); traceUrlStr != "" {
		trace := &tracer.Trace{}

		if traceStr, err := url.QueryUnescape(traceUrlStr); err != nil {
			log.Logf("[RELOAD] Decode trace err: %s", err)
		} else if err = proto.Unmarshal([]byte(traceStr), trace); err != nil {
			log.Logf("[RELOAD] Unmarshal trace err: %s", err)
		} else {
			id := tracer.NewThreadID()
			log.Logf("[RELOAD] Init, thread id: %d", id)
			r = r.WithContext(context.WithValue(r.Context(), tracer.ThreadIDKey{}, id))
			trace.Records = append(trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordReceive,
				Timestamp:   time.Now().UnixNano() - trace.BaseTimestamp,
				MessageName: r.URL.Path,
			})

			tracer.Store.SetByThreadID(id, trace)
		}
	}

	return r
}
