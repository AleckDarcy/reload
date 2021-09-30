package v3lockpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *LockRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LockResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *UnlockRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *UnlockResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
