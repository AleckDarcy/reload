package tracer

import (
	"sync"

	"github.com/AleckDarcy/reload/core/log"
)

type Storage struct {
	traces map[TraceID]*traces
	lock   sync.RWMutex
}

type traces struct {
	traces map[UUID]*wrapper
}

type Refer struct {
	UUID UUID
}

type wrapper struct {
	refer *Refer
	trace *Trace
}

var Store = &Storage{
	traces: map[TraceID]*traces{},
}

func NewStore() *Storage {
	return &Storage{
		traces: map[TraceID]*traces{},
	}
}

func (s *Storage) CheckByContextMeta(meta *ContextMeta) bool {
	s.lock.RLock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		_, ok = ts.traces[meta.uuid]
	}
	s.lock.RUnlock()

	return ok
}

func (s *Storage) GetByContextMeta(meta *ContextMeta) (*Trace, bool) {
	var t *Trace

	s.lock.RLock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		var w *wrapper
		if w, ok = ts.traces[meta.uuid]; ok {
			t = w.trace.Copy()
		}
	}
	s.lock.RUnlock()

	return t, ok
}

func (s *Storage) SetByContextMeta(meta *ContextMeta, trace *Trace) {
	s.lock.Lock()
	if ts, ok := s.traces[meta.traceID]; ok {
		if w, ok := ts.traces[meta.uuid]; ok {
			merge(w.trace, trace)
		} else {
			ts.traces[meta.uuid] = &wrapper{trace: trace.Copy()}
		}
	} else {
		s.traces[meta.traceID] = &traces{traces: map[UUID]*wrapper{meta.uuid: {trace: trace.Copy()}}}
	}
	s.lock.Unlock()
}

func (s *Storage) UpdateFunctionByContextMeta(meta *ContextMeta, function func(*Trace)) (*Trace, bool) {
	var t *Trace

	s.lock.Lock()
	ts, ok := s.traces[meta.traceID]
	if ok {
		var w *wrapper
		if w, ok = ts.traces[meta.uuid]; ok {
			function(w.trace)

			t = w.trace.Copy()
		}
	}
	s.lock.Unlock()

	return t, ok
}

func (s *Storage) DeleteByContextMeta(meta *ContextMeta) bool {
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

/**
 * 3milebeach note:
 * Use following methods with caution.
 * Parameter ref is used for helping build context-propagation mechanism.
 * Usually used under asynchronous communication patterns.
 */

// REQUEST
func (s *Storage) DoRequest(meta *ContextMeta, trace *Trace, record *Record, ref *Refer) {
	trace = trace.Copy()
	trace.Records = nil
	trace.AppendRecord(record)

	// 3milebeach note:
	// ref should be nil when receiving requests
	s.lock.Lock()
	if ts, ok := s.traces[meta.traceID]; ok {
		ts.traces[record.Uuid] = &wrapper{refer: ref, trace: trace}
	} else {
		s.traces[meta.traceID] = &traces{traces: map[UUID]*wrapper{meta.uuid: {refer: ref, trace: trace}}}
	}

	if ref == nil {
		log.Trace.PrintlnWithCaller("%s write trace %d:%s", meta.server, meta.traceID, record.Uuid)
	} else {
		log.Trace.PrintlnWithCaller("%s write trace with refer %d:%s -> %s", meta.server, meta.traceID, record.Uuid, ref.UUID)
	}

	s.lock.Unlock()
}

// RESPONSE
func (s *Storage) DoResponse(meta *ContextMeta, trace *Trace, record *Record) {
	s.lock.Lock()
	if ts, ok := s.traces[meta.traceID]; ok {
		if w, ok := ts.traces[meta.uuid]; ok {
			merge(w.trace, trace)

			// 3milebeach note:
			// ref should be nil when sending responses
			if ref := w.refer; ref != nil { // report to corresponding request's upstream request
				if wRef, ok := ts.traces[ref.UUID]; ok {
					merge(wRef.trace, w.trace)
					*trace = *wRef.trace
					trace.AppendRecord(record)

					log.Trace.PrintlnWithCaller("%s write trace to refer %d:%s", meta.server, meta.traceID, record.Uuid)
				} else {
					log.Trace.PrintlnWithCaller("%s no ref uuid %d:%s", meta.server, meta.traceID, meta.uuid)
				}
			} else {
				*trace = *w.trace
				trace.AppendRecord(record)

				log.Trace.PrintlnWithCaller("%s write trace %d:%s", meta.server, meta.traceID, record.Uuid)
			}

			delete(ts.traces, meta.uuid)
			if len(ts.traces) == 0 {
				delete(s.traces, meta.traceID)
			}
		} else {
			log.Trace.PrintlnWithCaller("%s no uuid %d:%s", meta.server, meta.traceID, meta.uuid)
		}
	} else {
		log.Trace.PrintlnWithCaller("%s no trace %d", meta.server, meta.traceID)
	}

	s.lock.Unlock()
}
