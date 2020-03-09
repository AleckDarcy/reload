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
	log.Logf("[RELOAD] Request: %v", r)
	if traceStr := r.Header.Get("Fi-Trace"); traceStr != "" {
		trace := &tracer.Trace{}
		if err := json.Unmarshal([]byte(traceStr), trace); err != nil {

		} else {
			threadID := tracer.NewThreadID()
			r = r.WithContext(context.WithValue(r.Context(), tracer.ThreadIDKey{}, threadID))
			trace.Records = append(trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordReceive,
				Timestamp:   time.Now().UnixNano(),
				MessageName: r.URL.Path,
			})
		}
	}

	return r
}
