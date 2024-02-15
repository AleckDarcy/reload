package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/AleckDarcy/reload/core/context_bus/public"

	"fmt"
	"runtime"
	"time"
)

// Payload is the package of event data inside LockFreeQueue
type Payload struct {
	cfgID int64
	ed    *cb.EventData
}

type observationBus struct {
	queue  *LockFreeQueue
	signal chan struct{}
}

var Bus = &observationBus{
	queue:  NewLockFreeQueue(),
	signal: make(chan struct{}, 1),
}

func (b *observationBus) OnSubmit(cfgID int64, ed *cb.EventData) {
	b.queue.Enqueue(&Payload{
		cfgID: cfgID,
		ed:    ed,
	})

	// try to invoke
	select {
	case b.signal <- struct{}{}:
		fmt.Println("notified")
		// message sent
	default:
		// fmt.Println("failed")
		// message dropped
	}
}

func (b *observationBus) doObservation() (cnt int) {
	for {
		v, ok := b.queue.Dequeue()
		if !ok {
			return
		}

		pay := v.(*Payload)
		if cfg := configure.ConfigureStore.GetConfigure(pay.cfgID); cfg != nil {
			if obs := cfg.GetObservationConfigure(pay.ed.Event.Recorder.Name); obs != nil {
				obs.Do(pay.ed.Event)
			}
		}

		cnt++
	}
}

func (b *observationBus) Run(sig chan struct{}) {
	for {
		cnt := 0
		select {
		case <-sig:
			return
		case <-b.signal:
			cnt = b.doObservation()
		case <-time.After(public.BUS_OBSERVATION_QUEUE_INTERVAL):
			cnt = b.doObservation()
		}

		fmt.Println("bus processed", cnt, "payloads")
		runtime.Gosched()
	}
}
