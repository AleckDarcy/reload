package mvccpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *KeyValue) MessageName() string {
	return "KeyValue"
}

func (m *KeyValue) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *KeyValue) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Event) MessageName() string {
	return "Event"
}

func (m *Event) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Event) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
