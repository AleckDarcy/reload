package v3lockpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *LockRequest) GetFI_Name() string {
	return "LockRequest"
}

func (m *LockRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LockRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *LockResponse) GetFI_Name() string {
	return "LockResponse"
}

func (m *LockResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LockResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *UnlockRequest) GetFI_Name() string {
	return "UnlockRequest"
}

func (m *UnlockRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *UnlockRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *UnlockResponse) GetFI_Name() string {
	return "UnlockResponse"
}

func (m *UnlockResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *UnlockResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}
