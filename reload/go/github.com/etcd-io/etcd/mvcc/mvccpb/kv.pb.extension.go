package mvccpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *KeyValue) GetFI_Name() string {
	return "KeyValue"
}

func (m *KeyValue) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *KeyValue) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *KeyValue) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *Event) GetFI_Name() string {
	return "Event"
}

func (m *Event) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Event) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Event) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}
