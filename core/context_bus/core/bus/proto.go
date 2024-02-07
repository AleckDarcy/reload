package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/reaction"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

type Configure struct {
	Reactions    map[string]*reaction.Configure
	Observations map[string]*cb.ObservationConfigure
}
