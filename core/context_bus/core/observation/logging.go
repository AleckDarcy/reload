package observation

import (
	"sync"
)

var ObserveHelper = &observeHelper{}

type observeHelper struct {
	bufPool sync.Pool
}

type event struct {
	buf []byte
}

const eventBufInitLen = 50

func newEvent() (e *event) {
	ep := ObserveHelper.bufPool.Get()
	if ep == nil {
		e = &event{buf: make([]byte, eventBufInitLen)}
	} else {
		e = ep.(*event)
	}

	e.buf = e.buf[:0]

	return
}

func (e *event) finalize() {
	ObserveHelper.bufPool.Put(e)
}
