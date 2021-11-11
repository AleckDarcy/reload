package raftpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Entry) GetFI_Name() string {
	return "Entry"
}

func (m *Entry) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Entry) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Entry) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *SnapshotMetadata) GetFI_Name() string {
	return "SnapshotMetadata"
}

func (m *SnapshotMetadata) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *SnapshotMetadata) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotMetadata) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

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

func (m *Snapshot) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *Message) GetFI_Name() string {
	return "Message"
}

func (m *Message) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Message) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Message) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

// 3milebeach todo: multiple entries
func (m *Message) PrepareTrace() *Message {
	if m.Trace == nil {
		for _, ent := range m.Entries {
			if ent.Trace != nil {
				m.Trace = ent.Trace
			}

			break
		}
	}

	return m
}

func (m *HardState) GetFI_Name() string {
	return "HardState"
}

func (m *HardState) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *HardState) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HardState) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *ConfState) GetFI_Name() string {
	return "ConfState"
}

func (m *ConfState) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfState) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfState) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *ConfChange) GetFI_Name() string {
	return "ConfChange"
}

func (m *ConfChange) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfChange) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfChange) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *ConfChangeSingle) GetFI_Name() string {
	return "ConfChangeSingle"
}

func (m *ConfChangeSingle) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfChangeSingle) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfChangeSingle) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *ConfChangeV2) GetFI_Name() string {
	return "ConfChangeV2"
}

func (m *ConfChangeV2) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *ConfChangeV2) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ConfChangeV2) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}
