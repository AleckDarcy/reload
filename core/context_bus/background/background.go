package background

import (
	"github.com/AleckDarcy/reload/core/context_bus/public"

	"time"
)

// List of background tasks:
// 1. environmental profiling
// 2. (?) configuration
// Initialized during deployment.

type Configure struct {
	EnvironmentProfiler bool
}

type signal struct {
	environmentProfiler chan struct{}
}

var Signal = signal{
	environmentProfiler: make(chan struct{}),
}

func Run() {
	go EnvironmentProfileProcessor(Signal.environmentProfiler)
}

func Stop() {

}

func EnvironmentProfileProcessor(sig chan struct{}) {
	EP.GetEnvironmentProfile()

	for {
		select {
		case <-sig:
			return
		case <-time.After(public.ENV_PROFILE_INTERVAL):
			EP.GetEnvironmentProfile()
		}
	}
}
