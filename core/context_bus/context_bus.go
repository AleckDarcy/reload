package context_bus

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"time"
)

var Configures = map[int64]*cb.Configure{}

func Submit(data *cb.EventRepresentation) {
	er := &cb.EventRepresentation{
		When:     &cb.EventWhen{Time: time.Now().UnixNano()},
		Where:    nil,
		Recorder: nil,
		What:     &cb.EventWhat{},
	}

	_ = er
}

func EnvironmentProfileProcessor(signal chan struct{}) {
	for {
		select {
		case <- signal:
			return
		case <- time.After(ENV_PROFILE_INTERVAL):
			EP.GetEnvironmentProfile()
		}
	}
}
