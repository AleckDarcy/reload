package configure

import (
	"github.com/AleckDarcy/reload/core/context_bus/core/observation"
	"github.com/AleckDarcy/reload/core/context_bus/core/reaction"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"sync"
)

type configureStore struct {
	lock       sync.RWMutex
	configures map[int64]*Configure // int64: configure_id
}

var ConfigureStore = configureStore{configures: map[int64]*Configure{}}

func (s *configureStore) SetConfigure(id int64, configure *cb.Configure) {
	racs := map[string]*reaction.Configure(nil)
	racMapMap := map[string]map[string]*reaction.Configure{} // <event (observation), <event (reaction), cfg>>

	if reactions := configure.Reactions; reactions != nil {
		racs = make(map[string]*reaction.Configure, len(reactions))
		for name, reaction_ := range reactions {
			pre := &reaction.PrerequisiteTree{
				PrerequisiteTree: *reaction_.PreTree,
			}

			pre.Indexing()

			rac := &reaction.Configure{
				Name:    name,
				Type:    reaction_.Type,
				Params:  reaction_.Params,
				PreTree: pre,
			}
			racs[name] = rac

			for _, node := range reaction_.PreTree.Nodes {
				if node.Type == cb.PrerequisiteNodeType_PrerequisiteMessage_ {
					racMap, ok := racMapMap[node.Message.Name]
					if !ok {
						racMap = map[string]*reaction.Configure{}
						racMapMap[node.Message.Name] = racMap
					}
					racMap[name] = rac
				}
			}
		}
	}

	racIndex := map[string][]*reaction.Configure{}
	for name, racMap := range racMapMap {
		racList := make([]*reaction.Configure, 0, len(racMap))
		for _, rac := range racMap {
			racList = append(racList, rac)
		}

		racIndex[name] = racList
	}

	cfg := &Configure{
		Reactions:     racs,
		Observations:  configure.Observations,
		ReactionIndex: racIndex,
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

func (c *Configure) InitializeSnapshots() *cb.PrerequisiteSnapshots {
	ss := make(map[string]*cb.PrerequisiteSnapshot, len(c.Reactions))

	for name, rac := range c.Reactions {
		ss[name] = (*cb.PrerequisiteSnapshot)(rac.PreTree.InitializeSnapshot())
	}

	return &cb.PrerequisiteSnapshots{Snapshots: ss}
}

func (c *Configure) UpdateSnapshots(name string, ss *cb.PrerequisiteSnapshots) {
	racs, ok := c.ReactionIndex[name]
	if !ok {
		return
	}

	for _, rac := range racs {
		rac.PreTree.UpdateSnapshot(name, (*reaction.PrerequisiteSnapshot)(ss.Snapshots[rac.Name]))
	}
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
