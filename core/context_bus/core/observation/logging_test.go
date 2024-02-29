package observation

import (
	cb_context "github.com/AleckDarcy/reload/core/context_bus/core/context"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"github.com/rs/zerolog/log"

	"context"
	"net/http"
	"testing"
	"time"
)

var path = cb.Test_Path_Rest_From
var pathNotFound = cb.Test_Path_Not_Found

var rest = cb.Test_EventMessage_Rest

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
	logCfg.Do(&cb.EventData{Event: &cb.EventRepresentation{When: &cb.EventWhen{Time: time.Now().UnixNano()}, What: what}})

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

func TestLoggingConfigure(t *testing.T) {
	what := new(cb.EventWhat)
	what.WithApplication(new(cb.EventMessage).SetMessage("received message from %s, a tag not found %s").SetPaths([]*cb.Path{path, pathNotFound}))
	what.WithLibrary("rest", rest)

	logCfg := &LoggingConfigure{}
	logCfg.Do(&cb.EventData{Event: &cb.EventRepresentation{When: &cb.EventWhen{Time: time.Now().UnixNano()}, What: what}})
}
