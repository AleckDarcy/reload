package v3electionpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *CampaignRequest) MessageName() string {
	return "CampaignRequest"
}

func (m *CampaignRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CampaignResponse) MessageName() string {
	return "CampaignResponse"
}

func (m *CampaignResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaderKey) MessageName() string {
	return "LeaderKey"
}

func (m *LeaderKey) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaderRequest) MessageName() string {
	return "LeaderRequest"
}

func (m *LeaderRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaderResponse) MessageName() string {
	return "LeaderResponse"
}

func (m *LeaderResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ResignRequest) MessageName() string {
	return "ResignRequest"
}

func (m *ResignRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ResignResponse) MessageName() string {
	return "ResignResponse"
}

func (m *ResignResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ProclaimRequest) MessageName() string {
	return "ProclaimRequest"
}

func (m *ProclaimRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *ProclaimResponse) MessageName() string {
	return "ProclaimResponse"
}

func (m *ProclaimResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
