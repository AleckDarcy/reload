package tracer

import (
	"sync"
)

type store struct {
	traces map[TraceID]*traces
	lock   sync.RWMutex
}

type traces struct {
	traces map[UUID]*Trace
}

var Store = &store{
	traces: map[TraceID]*traces{},
}

func (s *store) CheckByContextMeta(meta *ContextMeta) bool {
	s.lock.RLock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		_, ok = ts.traces[meta.uuid]
	}
	s.lock.RUnlock()

	return ok
}

func (s *store) GetByContextMeta(meta *ContextMeta) (*Trace, bool) {
	var t *Trace

	s.lock.RLock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		t, ok = ts.traces[meta.uuid]
		t = t.Copy()
	}
	s.lock.RUnlock()

	return t, ok
}

func (s *store) SetByContextMeta(meta *ContextMeta, trace *Trace) {
	s.lock.Lock()
	if ts, ok := s.traces[meta.traceID]; ok {
		if t, ok := ts.traces[meta.uuid]; ok {
			merge(t, trace)
		} else {
			ts.traces[meta.uuid] = trace
		}
	} else {
		s.traces[meta.traceID] = &traces{traces: map[UUID]*Trace{meta.uuid: trace}}
	}
	s.lock.Unlock()
}

func (s *store) UpdateFunctionByContextMeta(meta *ContextMeta, function func(*Trace)) (*Trace, bool) {
	var t *Trace

	s.lock.Lock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		if t, ok = ts.traces[meta.uuid]; ok {
			function(t)

			t = t.Copy()
		}
	}
	s.lock.Unlock()

	return t, ok
}

func (s *store) DeleteByContextMeta(meta *ContextMeta) bool {
	s.lock.Lock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		if _, ok = ts.traces[meta.uuid]; ok {
			delete(ts.traces, meta.uuid)

			if len(ts.traces) == 0 {
				delete(s.traces, meta.traceID)
			}
		}
	}
	s.lock.Unlock()

	return ok
}

func merge(dst, src *Trace) {
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
