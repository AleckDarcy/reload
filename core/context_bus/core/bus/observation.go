package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/background"
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"time"
)

// user interface

func OnSubmission(ctx *context.Context, where *cb.EventWhere, who *cb.EventRecorder, app *cb.EventMessage) {
	er := &cb.EventRepresentation{
		When:     &cb.EventWhen{Time: time.Now().UnixNano()},
		Where:    where,
		Recorder: who,
		What:     &cb.EventWhat{Application: app},
	}

	// write network API attributes
	er.What.WithLibrary(ctx.GetRequestContext().GetLib(), ctx.GetRequestContext().GetEventMessage())

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

	if cfg := configure.ConfigureStore.GetConfigure(ctx.GetRequestContext().GetConfigureID()); cfg != nil {
		if obs := cfg.GetObservationConfigure(ed.Event.Recorder.Name); obs != nil {
			obs.Do(ctx, ed.Event)
		}

		if rac := cfg.GetReaction(ed.Event.Recorder.Name); rac != nil {
			_ = rac
		}
	} // todo cfg == nil, default

	// todo update snapshot
	// todo put ed into bus
}
