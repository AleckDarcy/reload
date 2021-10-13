package walpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Record) GetFI_Name() string {
	return "RequestHeader"
}

func (m *Record) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Record) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Snapshot) GetFI_Name() string {
	return "RequestHeader"
}

func (m *Snapshot) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Snapshot) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
