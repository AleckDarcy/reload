package template

import (
	"context"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/tracer"
)

type Template struct {
	base *template.Template
}

func MarshalTracing(trace *tracer.Trace) string {
	return trace.JSONString()
}

func Must(t *template.Template, err error) *Template {
	if err != nil {
		panic(err)
	}

	return &Template{base: t}
}

func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.base.ExecuteTemplate(wr, name, data)
}

func (t *Template) ExecuteTemplateReload(ctx context.Context, w http.ResponseWriter, name string, data map[string]interface{}) error {
	//log.Logf("[RELOAD] ExecuteTemplateReload, called")
	if metaVal := ctx.Value(tracer.ContextMetaKey{}); metaVal != nil {
		meta := metaVal.(*tracer.ContextMeta)
		//log.Logf("[RELOAD] ExecuteTemplateReload, meta: %+v", meta)
		if trace, ok := tracer.Store.GetByContextMeta(meta); ok {
			//log.Logf("[RELOAD] ExecuteTemplateReload, trace found")

			trace.Records = append(trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordSend,
				Timestamp:   time.Now().UnixNano(),
				MessageName: meta.Url(),
				Uuid:        meta.UUID(),
				Service:     tracer.ServiceUUID,
			})

			trace.Rlfi = nil
			trace.Tfi = nil

			data["fi_trace"] = trace

			// delete trace from tracer.Store
			tracer.Store.DeleteByContextMeta(meta)

			// Content-Type: application/json instead of text/html
			w.Header().Set(html.ContentType, html.ContentTypeJSON)

			if err, ok := data["error"]; ok {
				if errStr, ok := err.(string); ok {
					if errStr != "" {
						return t.base.ExecuteTemplate(w, name, data)
					}
				}
			}

			return json.NewEncoder(w).Encode(data)
		}
	}

	return t.base.ExecuteTemplate(w, name, data)
}
