package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/background"
	"github.com/AleckDarcy/reload/core/context_bus/core/observation"
	"github.com/AleckDarcy/reload/core/context_bus/core/reaction"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/AleckDarcy/reload/core/context_bus/public"

	"testing"
	"time"
)

var ctx = &Context{}

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

var path = &cb.Path{
	Type: cb.PathType_Library,
	Path: []string{"rest", "from"},
}

func OnSubmission(ctx *Context, where *cb.EventWhere, who *cb.EventRecorder, app *cb.EventMessage) {
	er := &cb.EventRepresentation{
		When:     &cb.EventWhen{Time: time.Now().UnixNano()},
		Where:    where,
		Recorder: who,
		What:     &cb.EventWhat{Application: app},
	}

	// write network API attributes
	er.What.WithLibrary(ctx.GetRequestContext().GetLib(), nil).WithAttributes(ctx.GetRequestContext().GetAttrs())

	esp := background.EP.GetLatest()

	md := &cb.EventMetadata{
		Id:  0,
		Pcp: nil,
		Esp: esp.Timestamp,
	}

	ed := &cb.EventData{
		Event:    er,
		Metadata: md,
	}

	if cfg := ConfigureStore.GetConfigure(ctx.GetRequestContext().GetConfigureID()); cfg != nil {
		if obs := cfg.GetObservationConfigure(ed.Event.Recorder.Name); obs != nil {
			(*observation.Configure)(obs).Do(ed.Event)
		}

		if rac := cfg.GetReaction(ed.Event.Recorder.Name); rac != nil {
			_ = (*reaction.Configure)(rac)
		}
	}

	// todo update snapshot
	// todo put ed into bus
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

	ConfigureStore.SetConfigure(0, (*Configure)(cfg))

	ctx = ctx.SetRequestContext(NewRequestContext("rest", 0, rest.Attrs))

	app := new(cb.EventMessage).SetMessage("received message from %s")

	// func ServiceHandler(ctx, request) (response, error)
	// generated code
	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA"}, app)
	// application
}
