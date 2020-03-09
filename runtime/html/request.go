package html

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AleckDarcy/reload/core/log"

	"github.com/AleckDarcy/reload/core/tracer"
)

// called at the very beginning of handler()
func Init(r *http.Request) *http.Request {
	log.Logf("[RELOAD] Init, Request: %v", r)
	log.Logf("[RELOAD] Init, Fi-Trace: %s", r.Header.Get("Fi-Trace"))
	if traceStr := r.Header.Get("Fi-Trace"); traceStr != "" {
		trace := &tracer.Trace{}
		if err := json.Unmarshal([]byte(traceStr), trace); err != nil {

		} else {
			id := tracer.NewThreadID()
			log.Logf("[RELOAD] Init, thread id: %d", id)
			r = r.WithContext(context.WithValue(r.Context(), tracer.ThreadIDKey{}, id))
			trace.Records = append(trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordReceive,
				Timestamp:   time.Now().UnixNano(),
				MessageName: r.URL.Path,
			})
		}
	}

	return r
}
