package v3lockpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *LockRequest) MessageName() string {
	return "LockRequest"
}

func (m *LockRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LockResponse) MessageName() string {
	return "LockResponse"
}

func (m *LockResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *UnlockRequest) MessageName() string {
	return "UnlockRequest"
}

func (m *UnlockRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *UnlockResponse) MessageName() string {
	return "UnlockResponse"
}

func (m *UnlockResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
