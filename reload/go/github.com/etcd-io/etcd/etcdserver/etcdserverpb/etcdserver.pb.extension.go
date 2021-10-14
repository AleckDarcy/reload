package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Request) GetFI_Name() string {
	return "Request"
}

func (m *Request) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Request) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Request) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *Metadata) GetFI_Name() string {
	return "Metadata"
}

func (m *Metadata) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Metadata) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Metadata) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}
