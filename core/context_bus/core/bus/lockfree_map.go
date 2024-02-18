package bus

import cb "github.com/AleckDarcy/reload/core/context_bus/proto"

type eventDataMap struct {
	mp map[uint64]*cb.EventData
}

func (b *eventDataMap) Set(eveID uint64, ed *cb.EventData) {

}

func (b *eventDataMap) Get(eveID uint64) {

}

func (b *eventDataMap) Delete(eveID uint64) {

}

func (b *eventDataMap) Run(sig chan struct{}) {
	for {
		select {}
	}
}
