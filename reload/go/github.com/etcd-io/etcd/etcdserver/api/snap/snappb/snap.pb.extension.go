package snappb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Snapshot) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Snapshot) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
