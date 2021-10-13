package snappb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Snapshot) GetFI_Name() string {
	return "Snapshot"
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
