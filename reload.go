package reload

import "github.com/AleckDarcy/reload/injector"

// cheat gradle
func A() {
	injector.NewThreadID()
}