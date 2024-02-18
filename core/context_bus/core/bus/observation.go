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

// OnSubmission user interface
func OnSubmission(ctx *context.Context, where *cb.EventWhere, who *cb.EventRecorder, app *cb.EventMessage) {
	er := &cb.EventRepresentation{
		When:     &cb.EventWhen{Time: time.Now().UnixNano()},
		Where:    where,
		Recorder: who,
		What:     &cb.EventWhat{Application: app},
	}

	reqCtx := ctx.GetRequestContext()
	eveCtx := ctx.GetEventContext()

	// write network API attributes
	if reqCtx != nil {
		er.What.WithLibrary(reqCtx.GetLib(), reqCtx.GetEventMessage())
	}
	// write code base info

	esp := background.EP.GetLatest()

	md := &cb.EventMetadata{
		Id:  Bus.NewEventID(),
		Pcp: nil,
		Esp: esp.Timestamp,
	}

	ed := &cb.EventData{
		Event:    er,
		Metadata: md,
	}

	cfg := configure.ConfigureStore.GetConfigure(reqCtx.GetConfigureID())
	snapshots := cfg.UpdateSnapshots(who.GetName(), eveCtx.GetPrerequisiteSnapshots())

	if obs := cfg.GetObservationConfigure(who.GetName()); obs != nil {
		// update EventMetadata
		switch obs.Type {
		case cb.ObservationType_ObservationSingle:
			// by pass PrevEvent

		case cb.ObservationType_ObservationStart:
			// initialize event pair
			newEveCtx := new(context.EventContext).SetPrerequisiteSnapshots(snapshots).SetPrevEvent(eveCtx, ed)
			ctx.SetEventContext(newEveCtx)
		case cb.ObservationType_ObservationInter:
			newEveCtx := new(context.EventContext).SetPrerequisiteSnapshots(snapshots).SetPrevEvent(eveCtx, ed)
			ctx.SetEventContext(newEveCtx)
			fallthrough
		case cb.ObservationType_ObservationEnd:
			// finalize event pair
			_, prevED := eveCtx.GetPrevEvent()
			if prevED == nil {
				fmt.Errorf("eveCtx.GetPrevEvent() get nil")
			}
			ed.PrevEventData = prevED
		}
	}

	if rac := cfg.GetReaction(who.Name); rac != nil {
		if snapshot := snapshots.GetPrerequisiteSnapshot(who.Name); snapshot != nil {
			if ok, err := rac.PreTree.Check((*reaction.PrerequisiteSnapshot)(snapshot)); err != nil {

			} else if !ok {

			} else {
				fmt.Println("prerequisites accomplished")
			}
		}
	}

	// push EventData to bus
	Bus.OnSubmit(reqCtx.GetConfigureID(), ed)
}
