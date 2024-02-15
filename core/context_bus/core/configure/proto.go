package configure

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/reaction"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

type Configure struct {
	Reactions    map[string]*reaction.Configure
	Observations map[string]*cb.ObservationConfigure

	ReactionIndex map[string][]*reaction.Configure // <event name, reaction.Configure where use this event as a prerequisite>
}
