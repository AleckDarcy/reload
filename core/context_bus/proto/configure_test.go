package proto

import "testing"

func BenchmarkName(b *testing.B) {
	var cfg = &Configure{
		Observations: map[string]*Observation{
			"event1": {
				Logging: &LoggingConfigure{
					Timestamp:  nil,
					Stacktrace: nil,
					Attrs: []*AttributeConfiguration{
						{
							Name: "app.key21",
							Path: &Path{
								Type: PathType_Application,
								Path: []string{"key2", "key21"},
							},
						},
						{
							Name: "app.message",
							Path: &Path{
								Type: PathType_Application,
								Path: []string{"__message__"},
							},
						},
						{
							Name: "lib1.key11",
							Path: &Path{
								Type: PathType_Library,
								Path: []string{"lib1", "key1", "key11"},
							},
						},
					},
				},
			},
		},
	}

	var what = &EventWhat{}
	what.WithApplication(nil).
		SetMessage("application message").SetString("key1", "value1").
		WithAttributes("key2", nil).
		SetString("key21", "value21")
	what.WithLibrary("lib1", nil).
		SetMessage("lib1 message").SetString("key2", "value2").
		WithAttributes("key1", nil).
		SetString("key11", "value11")

	var er = &EventRepresentation{}
	er.WithWhat(what)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cfg.Observations["event1"].Logging.Do(er)
	}
}
