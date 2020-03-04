package tracer

import (
	"sync"

	"github.com/AleckDarcy/reload/core/errors"
)

type store struct {
	traces  map[int64]*Trace
	threads map[int64]*Trace
	lock    sync.RWMutex
}

var Store = store{
	traces:  map[int64]*Trace{},
	threads: map[int64]*Trace{},
}

// each() must be a read-only function
func (s store) IterateByThreadID(id int64, each func(int, *Record)) (err error) {
	s.lock.RLock()
	if c, ok := s.threads[id]; ok {
		for i, record := range c.Records {
			each(i, record)
		}
	} else {
		err = errors.ErrorTracer_ThreadIDNotFound
	}
	s.lock.RUnlock()

	return err
}

// each() must be a read-only function
func (s store) IterateByTraceID(id int64, each func(int, *Record)) (err error) {
	s.lock.RLock()
	if c, ok := s.traces[id]; ok {
		for i, record := range c.Records {
			each(i, record)
		}
	} else {
		err = errors.ErrorTracer_TraceIDNotFound
	}
	s.lock.RUnlock()

	return err
}

func (s store) GetByTraceID(id int64) (*Trace, bool) {
	s.lock.RLock()
	c, ok := s.traces[id]
	if ok {
		c = c.Copy()
	}
	s.lock.RUnlock()

	return c, ok
}

func (s store) GetByThreadID(id int64) (*Trace, bool) {
	s.lock.RLock()
	c, ok := s.threads[id]
	if ok {
		c = c.Copy()
	}
	s.lock.RUnlock()

	return c, ok
}

func (s store) SetByThreadID(id int64, trace *Trace) {
	s.lock.Lock()
	t, ok := s.traces[trace.Id]
	if ok {
		s.merge(trace, t)
		s.threads[id] = trace
	} else {
		s.traces[trace.Id] = trace
		s.threads[id] = trace
	}
	s.lock.Unlock()
}

func (s store) UpdateFunctionByThreadID(id int64, function func(*Trace)) *Trace {
	s.lock.Lock()
	t, ok := s.threads[id]
	if ok {
		function(t)
		s.traces[t.Id] = t
		t = t.Copy()
	}
	s.lock.Unlock()

	return t
}

func (s store) DeleteByThraceID(id int64) bool {
	s.lock.Lock()
	_, ok := s.traces[id]
	if ok {
		delete(s.traces, id)
	}
	s.lock.Unlock()

	return ok
}

func (s store) DeleteByThreadID(id int64) bool {
	s.lock.Lock()
	_, ok := s.threads[id]
	if ok {
		delete(s.threads, id)
	}
	s.lock.Unlock()

	return ok
}

func (s store) merge(dst, src *Trace) {
	for dstI, srcI, dstLen, srcLen := 0, 0, len(dst.Records), len(src.Records); dstI < dstLen && srcI < srcLen; dstI++ {
		dstRecord, srcRecord := dst.Records[dstI], src.Records[srcI]
		// insert srcRecord between dst.Records[dstI-1] and dst.Records[dstI]
		if dstRecord.Timestamp > srcRecord.Timestamp {
			dst.Records = append(dst.Records[:dstI+1], dst.Records[dstI:]...)
			dst.Records[dstI] = srcRecord

			srcI++
			dstLen++
		} else if dstRecord.Timestamp == srcRecord.Timestamp {
			// deep equal
			if dstRecord.MessageName == srcRecord.MessageName && dstRecord.Type == srcRecord.Type {
				srcI++
			}
		} else {
			dstI++
		}
	}
}

func (s store) append(id int64, record *Record) {
	c, ok := s.threads[id]
	if !ok {
		c = &Trace{
			Id: id,
		}
	}

	c.Records = append(c.Records, record)
	c.Depth = int64(len(c.Records))
}

func (s store) Append(id int64, record *Record) {
	s.lock.Lock()
	s.append(id, record)
	s.lock.Unlock()
}
