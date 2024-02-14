package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/background"
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/AleckDarcy/reload/core/context_bus/public"

	"testing"
	"time"
)

var path = &cb.Path{
	Type: cb.PathType_Library,
	Path: []string{"rest", "from"},
}

var rest = &cb.EventMessage{
	Attrs: &cb.Attributes{
		Attrs: map[string]*cb.AttributeValue{
			"from": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "SenderA",
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

func TestObservation(t *testing.T) {
	background.Run()
	time.Sleep(public.ENV_PROFILE_INTERVAL)

	cfg := &cb.Configure{
		Reactions: nil,
		Observations: map[string]*cb.ObservationConfigure{
			"EventA": {
				Logging: &cb.LoggingConfigure{
					Timestamp: &cb.TimestampConfigure{Format: public.TIME_FORMAT_RFC3339},
					Attrs: []*cb.AttributeConfigure{
						{
							Name: "rest.key",
							Path: &cb.Path{
								Type: cb.PathType_Library,
								Path: []string{"rest", "key"},
							},
						},
					},
				},
			},
		},
	}

	id := int64(1)
	configure.ConfigureStore.SetConfigure(id, cfg)

	ctx := context.NewContext(context.NewRequestContext("rest", id, rest), nil)

	app := new(cb.EventMessage).SetMessage("received message from %s").SetPaths([]*cb.Path{path})

	// func ServiceHandler(ctx, request) (response, error)
	// generated code
	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA"}, app)
	// application
}
