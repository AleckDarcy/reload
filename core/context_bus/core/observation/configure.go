package observation

import (
	"fmt"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	"github.com/AleckDarcy/reload/core/context_bus/core/encoder"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"os"
	"time"
)

// implementations

func (c *Configure) Do(ctx *context.Context, er *cb.EventRepresentation) {
	(*LoggingConfigure)(c.Logging).Do(ctx, er)
	(*TracingConfigure)(c.Tracing).Do(er)
	for _, metric := range c.Metrics {
		(*MetricsConfigure)(metric).Do(er)
	}
}

func (c *TimestampConfigure) Do(ctx *context.Context) {
	if c != nil {
		return
	}
}

func (c *StackTraceConfigure) Do(ctx *context.Context) {

}

func (c *LoggingConfigure) Do(ctx *context.Context, er *cb.EventRepresentation) interface{} {
	if c == nil {
		return nil
	}

	e := newEvent()
	e.buf = encoder.JSONEncoder.BeginObject(e.buf)

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "level")
	e.buf = encoder.JSONEncoder.AppendString(e.buf, "info")

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "time")

	e.buf = encoder.JSONEncoder.BeginString(e.buf)
	e.buf = time.Unix(0, er.When.Time).AppendFormat(e.buf, time.RFC3339)
	e.buf = encoder.JSONEncoder.EndString(e.buf)

	if reqCtx := ctx.GetRequestContext(); reqCtx != nil {
		er.What.WithLibrary(reqCtx.GetLib(), reqCtx.GetEventMessage())
	}

	// do message
	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "message")
	msg := er.What.Application.GetMessage()
	paths := er.What.Application.GetPaths()
	values := make([]interface{}, len(paths))
	for i, path := range paths {
		values[i], _ = er.What.GetValue(path)
	}
	e.buf = encoder.JSONEncoder.AppendString(e.buf, fmt.Sprintf(msg, values...))

	// do tag
	if len(c.Attrs) != 0 {
		e.buf = encoder.JSONEncoder.AppendKey(e.buf, "tags")
		tags := DoTag(c.Attrs, er)
		e.buf = encoder.JSONEncoder.AppendTags(e.buf, tags)
	}

	e.buf = encoder.JSONEncoder.EndObject(e.buf)
	str := string(e.buf)
	e.finalize()

	(*TimestampConfigure)(c.Timestamp).Do(ctx)
	(*StackTraceConfigure)(c.Stacktrace).Do(ctx)

	switch c.Out {
	case cb.LogOutType_Stdout:
		fmt.Fprintln(os.Stdout, str)
	case cb.LogOutType_Stderr:
		fmt.Fprintln(os.Stderr, str)
	case cb.LogOutType_File:

	default:
		fmt.Fprintln(os.Stdout, str)
	}

	return str
}

func (c *TracingConfigure) Do(er *cb.EventRepresentation) {
	if c == nil {
		return
	}
}

func (c *MetricsConfigure) Do(er *cb.EventRepresentation) {
	if c == nil {
		return
	}
}

func DoTag(cfg []*cb.AttributeConfigure, er *cb.EventRepresentation) map[string]string {
	tags := map[string]string{}

	for _, path := range cfg {
		str, err := er.What.GetValue(path.Path)

		if err == nil {
			tags[path.Name] = str
		}
	}

	return tags
}
