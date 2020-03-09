package template

import (
	"context"
	"html/template"
	"io"

	"github.com/AleckDarcy/reload/core/log"

	"github.com/AleckDarcy/reload/core/tracer"
)

type Template struct {
	base *template.Template
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

func (t *Template) ExecuteTemplateReload(ctx context.Context, wr io.Writer, name string, data map[string]interface{}) error {
	log.Logf("[RELOAD] ExecuteTemplateReload, called")
	if idVal := ctx.Value(tracer.ThreadIDKey{}); idVal != nil {
		id := idVal.(int64)
		log.Logf("[RELOAD] ExecuteTemplateReload, thread id: %d", id)
		if trace, ok := tracer.Store.GetByThreadID(id); ok {
			log.Logf("[RELOAD] ExecuteTemplateReload, trace found")
			data["fi_trace"] = trace

			// delete trace from tracer.Store
			tracer.Store.DeleteByThraceID(trace.Id)
			tracer.Store.DeleteByThreadID(id)
		}
	}

	return t.base.ExecuteTemplate(wr, name, data)
}
