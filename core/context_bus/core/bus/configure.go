package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/observation"
	"github.com/AleckDarcy/reload/core/context_bus/core/reaction"
	"github.com/AleckDarcy/reload/core/context_bus/proto"

	"sync"
)

type configureStore struct {
	lock       sync.RWMutex
	configures map[int64]*Configure // int64: configure_id
}

var ConfigureStore = configureStore{configures: map[int64]*Configure{}}

func (s *configureStore) SetConfigure(id int64, configure *proto.Configure) {
	var racs map[string]*reaction.Configure
	if reactions := configure.Reactions; reactions != nil {
		racs = make(map[string]*reaction.Configure, len(reactions))
		for name, reaction_ := range reactions {
			pre := &reaction.PrerequisiteTree{
				PrerequisiteTree: reaction_.PreTree[0],
			}

			pre.Indexing()

			racs[name] = &reaction.Configure{
				Type:    reaction_.Type,
				Params:  reaction_.Params,
				PreTree: pre,
			}
		}
	}

	cfg := &Configure{
		Reactions:    racs,
		Observations: configure.Observations,
	}
	s.lock.Lock()
	s.configures[id] = cfg
	s.lock.Unlock()
}

func (s *configureStore) GetConfigure(id int64) *Configure {
	s.lock.RLock()
	cfg := s.configures[id]
	s.lock.RUnlock()

	return cfg
}

func (c *Configure) GetObservationConfigure(name string) *observation.Configure {
	if c.Observations == nil {
		return nil
	}

	return (*observation.Configure)(c.Observations[name])
}

func (c *Configure) GetReaction(name string) *reaction.Configure {
	if c.Reactions == nil {
		return nil
	}

	return c.Reactions[name]
}
