package leasepb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Lease) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Lease) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseInternalRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *LeaseInternalRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseInternalResponse) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *LeaseInternalResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
