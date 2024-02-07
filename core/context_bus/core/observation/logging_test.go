package observation

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/encoder"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/delimitrou/DeathStarBench/hotelreservation/vendor/github.com/AleckDarcy/reload/core/context_bus/core/bus"

	"github.com/rs/zerolog/log"
	"net/http"

	"context"
	"fmt"
	"testing"
	"time"
)

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
	what.WithLibrary("rest", nil).Merge(bus.rest)
	value, _ := what.GetValue(bus.path)
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
func ServiceHandlerA(ctx context.Context, req *http.Request) *http.Response {
	cb_extern.Observe("received message from {result.from}")

	return &http.Response{}
}
*/

func ServiceHandlerAContextBus(ctx context.Context, req *http.Request) *http.Response {
	what := new(cb.EventWhat)
	what.WithApplication(new(cb.EventMessage).SetMessage("received message from %s"))
	Observe(what)

	return &http.Response{}
}

func TestServiceHandlerA(t *testing.T) {
	ctx := context.WithValue(context.Background(), "from", "sender")

	ServiceHandlerA(ctx, nil)

	ServiceHandlerAContextBus(ctx, nil)
}

func BenchmarkLogging(b *testing.B) {
	ctx := context.WithValue(context.Background(), "from", "sender")

	for i := 0; i < b.N; i++ {
		//ServiceHandlerA(ctx, nil)

		ServiceHandlerAContextBus(ctx, nil)
	}

	_ = ctx
}
