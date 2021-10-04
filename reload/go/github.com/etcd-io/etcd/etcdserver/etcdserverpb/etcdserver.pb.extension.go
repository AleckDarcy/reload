package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Request) MessageName() string {
	return "Request"
}

func (m *Request) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Request) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Metadata) MessageName() string {
	return "Metadata"
}

func (m *Metadata) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Metadata) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
