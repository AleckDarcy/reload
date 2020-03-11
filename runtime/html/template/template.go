package template

import (
	"context"
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/log"
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
	if idVal := ctx.Value(tracer.ThreadIDKey{}); idVal != nil {
		id := idVal.(int64)
		log.Logf("[RELOAD] ExecuteTemplateReload, thread id: %d", id)
		if trace, ok := tracer.Store.GetByThreadID(id); ok {
			log.Logf("[RELOAD] ExecuteTemplateReload, trace found")
			data["fi_trace"] = trace

			// delete trace from tracer.Store
			tracer.Store.DeleteByTraceID(trace.Id)
			tracer.Store.DeleteByThreadID(id)

			// Content-Type: application/json instead of text/html
			w.Header().Set(html.ContentType, html.ContentTypeJSON)

			return json.NewEncoder(w).Encode(data)
		}
	}

	return t.base.ExecuteTemplate(w, name, data)
}
