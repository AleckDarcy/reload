package leasepb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *Lease) GetFI_Name() string {
	return "Lease"
}

func (m *Lease) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Lease) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Lease) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *LeaseInternalRequest) GetFI_Name() string {
	return "LeaseInternalRequest"
}

func (m *LeaseInternalRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *LeaseInternalRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseInternalRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *LeaseInternalResponse) GetFI_Name() string {
	return "LeaseInternalResponse"
}

func (m *LeaseInternalResponse) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *LeaseInternalResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseInternalResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}
