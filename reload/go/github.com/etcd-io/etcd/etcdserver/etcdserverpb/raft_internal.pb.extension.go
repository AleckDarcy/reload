package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *RequestHeader) MessageName() string {
	return "RequestHeader"
}

func (m *RequestHeader) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *RequestHeader) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *InternalRaftRequest) MessageName() string {
	return "InternalRaftRequest"
}

func (m *InternalRaftRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *InternalRaftRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *EmptyResponse) MessageName() string {
	return "EmptyResponse"
}

func (m *EmptyResponse) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *EmptyResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *InternalAuthenticateRequest) MessageName() string {
	return "InternalAuthenticateRequest"
}

func (m *InternalAuthenticateRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *InternalAuthenticateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
