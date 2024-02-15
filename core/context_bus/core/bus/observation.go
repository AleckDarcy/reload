package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/background"
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	"github.com/AleckDarcy/reload/core/context_bus/core/reaction"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"fmt"
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
	if reqCtx := ctx.GetRequestContext(); reqCtx != nil {
		er.What.WithLibrary(reqCtx.GetLib(), reqCtx.GetEventMessage())
	}
	// write code base info

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
		//if obs := cfg.GetObservationConfigure(ed.Event.Recorder.Name); obs != nil {
		//	obs.Do(ed.Event)
		//}

		if rac := cfg.GetReaction(who.Name); rac != nil {
			// todo update snapshot

			if ok, err := rac.PreTree.Check((*reaction.PrerequisiteSnapshot)(ctx.GetEventContext().GetPrerequisiteSnapshot())); err != nil {

			} else if !ok {

			} else {
				fmt.Println("bbbbbbbb")
			}
		}
	} // todo cfg == nil, default

	// todo put ed into bus
	Bus.OnSubmit(ctx.GetRequestContext().GetConfigureID(), ed)
}
