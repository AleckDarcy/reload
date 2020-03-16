package html

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"
)

// called at the very beginning of handler()
func Init(r *http.Request) *http.Request {
	log.Logf("[RELOAD] Init, Fi-Trace: %s", r.Header.Get("Fi-Trace"))
	if traceStr := r.Header.Get("Fi-Trace"); traceStr != "" {
		trace := &tracer.Trace{}

		if err := json.Unmarshal([]byte(traceStr), trace); err != nil {
			log.Logf("[RELOAD] Init, unmarshal trace err: %s", err)
		} else {
			if len(trace.Records) != 1 {
				log.Logf("[RELOAD] Init, receive invalid trace: %v", trace.JSONString())
			} else {
				uuid := trace.Records[0].Uuid
				log.Logf("[RELOAD] Init, uuid: %s", uuid)

				meta := tracer.NewContextMeta(trace.Id, uuid)
				r = r.WithContext(tracer.NewContextWithContextMeta(r.Context(), meta))

				trace.Records[0] = &tracer.Record{
					Type:        tracer.RecordType_RecordReceive,
					Timestamp:   time.Now().UnixNano(),
					MessageName: r.URL.Path,
					Uuid:        uuid,
				}
				tracer.Store.SetByContextMeta(meta, trace)
			}
		}
	}

	return r
}
