package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"sync"
	"testing"
	"time"
)

func TestObservationBus_Observation(t *testing.T) {
	go Bus.Run(nil)

	id := int64(2)
	configure.ConfigureStore.SetConfigure(id, cfg2)

	ctx := context.NewContext(context.NewRequestContext("rest", id, rest), nil)
	app := new(cb.EventMessage).SetMessage("received message from %s").SetPaths([]*cb.Path{path})

	n := 100
	wg := sync.WaitGroup{}
	worker := func() {
		for i := 0; i < n; i++ {
			OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA"}, app)

			time.Sleep(time.Millisecond * 500)
		}
		wg.Done()
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()

	time.Sleep(time.Second * 2)
}

func TestObservationBus_Reaction(t *testing.T) {
	go Bus.Run(nil)

	id := int64(3)
	configure.ConfigureStore.SetConfigure(id, cfg3)

	cfg := configure.ConfigureStore.GetConfigure(id)
	t.Log(cfg.ReactionIndex)

	ss := cfg.InitializeSnapshots()
	t.Log(ss)

	cfg.UpdateSnapshots("EventA", ss)
	t.Log(ss)

	cfg.UpdateSnapshots("EventB", ss)
	t.Log(ss)

	cfg.UpdateSnapshots("EventC", ss)
	t.Log(ss)

	reqCtx := context.NewRequestContext("rest", id, rest)
	eveCtx := context.NewEventContext(nil, &cb.PrerequisiteSnapshot{
		Value: []int64{0, 0, 0},
		Acc:   false,
	})
	ctx := context.NewContext(reqCtx, eveCtx)
	app := new(cb.EventMessage).SetMessage("received message from %s").SetPaths([]*cb.Path{path})

	n := 1
	for i := 0; i < n; i++ {
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA"}, app)

		// initialize PrerequisiteSnapshot for each submission

		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventC"}, app)

		time.Sleep(time.Millisecond * 500)
	}

	time.Sleep(time.Second * 2)
}
