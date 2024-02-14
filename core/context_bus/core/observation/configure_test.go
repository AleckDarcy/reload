package observation

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"testing"
	"time"
)

func BenchmarkName(b *testing.B) {
	var cfg = &cb.LoggingConfigure{
		Timestamp:  nil,
		Stacktrace: nil,
		Attrs: []*cb.AttributeConfigure{
			{
				Name: "app.key21",
				Path: &cb.Path{
					Type: cb.PathType_Application,
					Path: []string{"key2", "key21"},
				},
			},
			{
				Name: "app.message",
				Path: &cb.Path{
					Type: cb.PathType_Application,
					Path: []string{"__message__"},
				},
			},
			{
				Name: "lib1.key11",
				Path: &cb.Path{
					Type: cb.PathType_Library,
					Path: []string{"lib1", "key1", "key11"},
				},
			},
		},
	}

	what := new(cb.EventWhat)
	what.WithApplication(nil).
		SetMessage("application message").GetAttributes().SetString("key1", "value1").
		WithAttributes("key2", nil).
		SetString("key21", "value21")

	ctx := context.NewContext(context.NewRequestContext("rest", 0, rest), nil)

	er := new(cb.EventRepresentation)
	er.WithWhen(&cb.EventWhen{Time: time.Now().Unix()})
	er.WithWhat(what)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.Log((*LoggingConfigure)(cfg).Do(ctx, er))
	}
}
