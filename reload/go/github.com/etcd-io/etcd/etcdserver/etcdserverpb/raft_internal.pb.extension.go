package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *RequestHeader) GetFI_Name() string {
	return "RequestHeader"
}

func (m *RequestHeader) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *RequestHeader) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *InternalRaftRequest) GetFI_Name() string {
	return "InternalRaftRequest"
}

func (m *InternalRaftRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *InternalRaftRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *EmptyResponse) GetFI_Name() string {
	return "EmptyResponse"
}

func (m *EmptyResponse) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *EmptyResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *InternalAuthenticateRequest) GetFI_Name() string {
	return "InternalAuthenticateRequest"
}

func (m *InternalAuthenticateRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *InternalAuthenticateRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
