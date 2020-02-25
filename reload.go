package reload

import (
	"github.com/AleckDarcy/reload/google_golang_org/grpc"
	"github.com/AleckDarcy/reload/injector"
)

// cheat gradle

// cheat goDep
func A() {
	injector.NewThreadID()

	_ = grpc.INJECTOR_CHEATER
}