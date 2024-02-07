package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"testing"
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

	var what = &cb.EventWhat{}
	what.WithApplication(nil).
		SetMessage("application message").GetAttributes().SetString("key1", "value1").
		WithAttributes("key2", nil).
		SetString("key21", "value21")
	what.WithLibrary("lib1", nil).
		SetMessage("lib1 message").GetAttributes().SetString("key2", "value2").
		WithAttributes("key1", nil).
		SetString("key11", "value11")

	var er = &cb.EventRepresentation{}
	er.WithWhat(what)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		(*LoggingConfigure)(cfg).Do(er)
	}
}
