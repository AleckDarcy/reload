package raftpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Entry) MessageName() string {
	return "Entry"
}

func (m *Entry) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Entry) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotMetadata) MessageName() string {
	return "SnapshotMetadata"
}

func (m *SnapshotMetadata) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *SnapshotMetadata) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Snapshot) MessageName() string {
	return "Snapshot"
}

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

func (m *Message) MessageName() string {
	return "Message"
}

func (m *Message) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Message) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HardState) MessageName() string {
	return "HardState"
}

func (m *HardState) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *HardState) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfState) MessageName() string {
	return "ConfState"
}

func (m *ConfState) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfState) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfChange) MessageName() string {
	return "ConfChange"
}

func (m *ConfChange) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfChange) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfChangeSingle) MessageName() string {
	return "ConfChangeSingle"
}

func (m *ConfChangeSingle) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfChangeSingle) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfChangeV2) MessageName() string {
	return "ConfChangeV2"
}

func (m *ConfChangeV2) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfChangeV2) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
