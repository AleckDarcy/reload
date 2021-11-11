package tracer

import "sync"

type Plugin struct {
	ServerID UUID

	Store *Storage
}

// Only prints serverID
func (p *Plugin) String() string {
	return p.ServerID
}

var pluginLock = sync.RWMutex{}
var plugins = map[UUID]*Plugin{}

func GetPlugin(id UUID) *Plugin {
	if id == "" {
		panic("server id is empty")
	}

	pluginLock.RLock()
	p, ok := plugins[id]
	pluginLock.RUnlock()

	if !ok {
		pluginLock.Lock()
		p, ok = plugins[id]
		if !ok {
			p = &Plugin{
				ServerID: id,
				Store:    NewStore(),
			}
			plugins[id] = p
		}
		pluginLock.Unlock()
	}

	return p
}
