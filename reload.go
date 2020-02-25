package reload

import (
	"github.com/AleckDarcy/reload/injector"
)

// cheat gradle

// cheat goDep
func A() {
	injector.NewThreadID()
}