package observation

import (
	cb_context "github.com/AleckDarcy/reload/core/context_bus/core/context"
	"github.com/AleckDarcy/reload/core/context_bus/core/encoder"

	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"github.com/rs/zerolog/log"
	"net/http"

	"context"
	"fmt"
	"testing"
	"time"
)

var rest = &cb.EventMessage{
	Attrs: &cb.Attributes{
		Attrs: map[string]*cb.AttributeValue{
			"from": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "SenderB",
			},
			"key": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "This a string attribute",
			},
			"key_": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "This another string attribute",
			},
		},
	},
}

var path = &cb.Path{
	Type: cb.PathType_Library,
	Path: []string{"rest", "from"},
}

func Observe(what *cb.EventWhat) {
	ts := time.Now()
	e := newEvent()
	e.buf = encoder.JSONEncoder.BeginObject(e.buf)

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "level")
	e.buf = encoder.JSONEncoder.AppendString(e.buf, "info")

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "time")

	e.buf = encoder.JSONEncoder.BeginString(e.buf)
	e.buf = ts.AppendFormat(e.buf, time.RFC3339)
	e.buf = encoder.JSONEncoder.EndString(e.buf)

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "message")

	msg := what.Application.GetMessage()
	what.WithLibrary("rest", nil).Merge(rest)
	value, _ := what.GetValue(path)
	values := []interface{}{value}

	e.buf = encoder.JSONEncoder.AppendString(e.buf, fmt.Sprintf(msg, values...))
	e.buf = encoder.JSONEncoder.EndObject(e.buf)

	fmt.Println(string(e.buf))
	e.finalize()
}

// example of generated code for ServiceHandler

func ServiceHandlerA(ctx context.Context, req *http.Request) *http.Response {
	log.Info().Msgf("received message from %s", ctx.Value("from"))

	return &http.Response{}
}

/*
func ServiceHandlerAContextBus(ctx context.Context, req *http.Request) *http.Response {
	cb_extern.Observe("received message from {result.from}")

	return &http.Response{}
}
*/

func ServiceHandlerAContextBus(ctx *cb_context.Context, req *http.Request) *http.Response {
	what := new(cb.EventWhat)
	what.WithApplication(new(cb.EventMessage).SetMessage("received message from %s").SetPaths([]*cb.Path{path}))
	what.WithLibrary("rest", rest)

	logCfg := &LoggingConfigure{}
	str := logCfg.Do(&cb.EventRepresentation{When: &cb.EventWhen{Time: time.Now().UnixNano()}, What: what})

	fmt.Println(str)

	return &http.Response{}
}

func TestServiceHandlerA(t *testing.T) {
	ctx := context.WithValue(context.Background(), "from", "senderA")
	ServiceHandlerA(ctx, nil)

	cb_ctx := &cb_context.Context{}
	ServiceHandlerAContextBus(cb_ctx, nil)
}

func BenchmarkLogging(b *testing.B) {
	ctx := context.WithValue(context.Background(), "from", "senderA")
	cb_ctx := &cb_context.Context{}

	for i := 0; i < b.N; i++ {
		ServiceHandlerA(ctx, nil)

		//ServiceHandlerAContextBus(cb_ctx, nil)
	}

	_ = ctx
	_ = cb_ctx
}
