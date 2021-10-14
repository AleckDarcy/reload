package v3electionpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *CampaignRequest) GetFI_Name() string {
	return "CampaignRequest"
}

func (m *CampaignRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CampaignRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *CampaignResponse) GetFI_Name() string {
	return "CampaignResponse"
}

func (m *CampaignResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CampaignResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *LeaderKey) GetFI_Name() string {
	return "LeaderKey"
}

func (m *LeaderKey) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaderKey) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_ // 3milebeach todo
}

func (m *LeaderRequest) GetFI_Name() string {
	return "LeaderRequest"
}

func (m *LeaderRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaderRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *LeaderResponse) GetFI_Name() string {
	return "LeaderResponse"
}

func (m *LeaderResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaderResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *ResignRequest) GetFI_Name() string {
	return "ResignRequest"
}

func (m *ResignRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ResignRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *ResignResponse) GetFI_Name() string {
	return "ResignResponse"
}

func (m *ResignResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ResignResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *ProclaimRequest) GetFI_Name() string {
	return "ProclaimRequest"
}

func (m *ProclaimRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ProclaimRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *ProclaimResponse) GetFI_Name() string {
	return "ProclaimResponse"
}

func (m *ProclaimResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ProclaimResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}
