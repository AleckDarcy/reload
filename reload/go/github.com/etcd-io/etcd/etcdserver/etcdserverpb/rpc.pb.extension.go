package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *ResponseHeader) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RangeRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RangeResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *PutRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *PutResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DeleteRangeRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DeleteRangeResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

type isRequestOp_Request interface {
	isRequestOp_Request()
	MarshalTo([]byte) (int, error)
	Size() int
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func (m *RequestOp) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Request.SetTrace(trace)
	}
}

func (m *RequestOp_RequestRange) GetTrace() *tracer.Trace {
	if m != nil {
		return m.RequestRange.GetTrace()
	}

	return nil
}

func (m *RequestOp_RequestRange) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.RequestRange.SetTrace(trace)
	}
}

func (m *RequestOp_RequestPut) GetTrace() *tracer.Trace {
	if m != nil {
		return m.RequestPut.GetTrace()
	}

	return nil
}

func (m *RequestOp_RequestPut) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.RequestPut.SetTrace(trace)
	}
}

func (m *RequestOp_RequestDeleteRange) GetTrace() *tracer.Trace {
	if m != nil {
		return m.RequestDeleteRange.GetTrace()
	}

	return nil
}

func (m *RequestOp_RequestDeleteRange) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.RequestDeleteRange.SetTrace(trace)
	}
}

func (m *RequestOp_RequestTxn) GetTrace() *tracer.Trace {
	if m != nil {
		return m.RequestTxn.GetTrace()
	}

	return nil
}

func (m *RequestOp_RequestTxn) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.RequestTxn.SetTrace(trace)
	}
}

type isResponseOp_Response interface {
	isResponseOp_Response()
	MarshalTo([]byte) (int, error)
	Size() int
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func (m *ResponseOp) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Response.SetTrace(trace)
	}
}

func (m *ResponseOp_ResponseRange) GetTrace() *tracer.Trace {
	if m != nil {
		return m.ResponseRange.GetTrace()
	}

	return nil
}

func (m *ResponseOp_ResponseRange) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.ResponseRange.SetTrace(trace)
	}
}

func (m *ResponseOp_ResponsePut) GetTrace() *tracer.Trace {
	if m != nil {
		return m.ResponsePut.GetTrace()
	}

	return nil
}

func (m *ResponseOp_ResponsePut) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.ResponsePut.SetTrace(trace)
	}
}

func (m *ResponseOp_ResponseDeleteRange) GetTrace() *tracer.Trace {
	if m != nil {
		return m.ResponseDeleteRange.GetTrace()
	}

	return nil
}

func (m *ResponseOp_ResponseDeleteRange) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.ResponseDeleteRange.SetTrace(trace)
	}
}

func (m *ResponseOp_ResponseTxn) GetTrace() *tracer.Trace {
	if m != nil {
		return m.ResponseTxn.GetTrace()
	}

	return nil
}

func (m *ResponseOp_ResponseTxn) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.ResponseTxn.SetTrace(trace)
	}
}

func (m *Compare) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *TxnRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *TxnResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CompactionRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CompactionResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashKVRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashKVResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
	MarshalTo([]byte) (int, error)
	Size() int
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func (m *WatchRequest_CreateRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.CreateRequest.GetTrace()
	}

	return nil
}

func (m *WatchRequest_CreateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.CreateRequest.SetTrace(trace)
	}
}

func (m *WatchRequest_CancelRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.CancelRequest.GetTrace()
	}

	return nil
}

func (m *WatchRequest_CancelRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.CancelRequest.SetTrace(trace)
	}
}

func (m *WatchRequest_ProgressRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.ProgressRequest.GetTrace()
	}

	return nil
}

func (m *WatchRequest_ProgressRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.ProgressRequest.SetTrace(trace)
	}
}

func (m *WatchRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.RequestUnion.SetTrace(trace)
	}
}

func (m *WatchCreateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchCancelRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchProgressRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseGrantRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseGrantResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseRevokeRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseRevokeResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpoint) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpointRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpointResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseKeepAliveRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseKeepAliveResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseTimeToLiveRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseTimeToLiveResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseLeasesRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseStatus) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseLeasesResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Member) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberAddRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberAddResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberRemoveRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberRemoveResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberUpdateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberUpdateResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberListRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberListResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberPromoteRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberPromoteResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DefragmentRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DefragmentResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MoveLeaderRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MoveLeaderResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmMember) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *StatusRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *StatusResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthEnableRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthDisableRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthenticateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserAddRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGetRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserDeleteRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserChangePasswordRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGrantRoleRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserRevokeRoleRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleAddRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGetRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserListRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleListRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleDeleteRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGrantPermissionRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleRevokePermissionRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthEnableResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthDisableResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthenticateResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserAddResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGetResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserDeleteResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserChangePasswordResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGrantRoleResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserRevokeRoleResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleAddResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGetResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleListResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserListResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleDeleteResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGrantPermissionResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleRevokePermissionResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
