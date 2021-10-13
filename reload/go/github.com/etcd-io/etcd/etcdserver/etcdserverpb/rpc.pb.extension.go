package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *ResponseHeader) GetFI_Name() string {
	return "ResponseHeader"
}

func (m *ResponseHeader) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RangeRequest) GetFI_Name() string {
	return "RangeRequest"
}

func (m *RangeRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RangeResponse) GetFI_Name() string {
	return "RangeResponse"
}

func (m *RangeResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *PutRequest) GetFI_Name() string {
	return "PutRequest"
}

func (m *PutRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *PutResponse) GetFI_Name() string {
	return "PutResponse"
}

func (m *PutResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DeleteRangeRequest) GetFI_Name() string {
	return "DeleteRangeRequest"
}

func (m *DeleteRangeRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DeleteRangeResponse) GetFI_Name() string {
	return "DeleteRangeResponse"
}

func (m *DeleteRangeResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

// extended
type isRequestOp_Request interface {
	isRequestOp_Request()
	MarshalTo([]byte) (int, error)
	Size() int
	GetFI_Trace() *tracer.Trace
	SetFI_Trace(*tracer.Trace)
}

func (m *RequestOp) GetFI_Name() string {
	return "RequestOp"
}

func (m *RequestOp) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Request.GetFI_Trace()
	}

	return nil
}

func (m *RequestOp) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Request.SetFI_Trace(trace)
	}
}

func (m *RequestOp_RequestRange) GetFI_Name() string {
	return "RequestOp_RequestRange"
}

func (m *RequestOp_RequestRange) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.RequestRange.GetFI_Trace()
	}

	return nil
}

func (m *RequestOp_RequestRange) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.RequestRange.SetFI_Trace(trace)
	}
}

func (m *RequestOp_RequestPut) GetFI_Name() string {
	return "RequestOp_RequestPut"
}

func (m *RequestOp_RequestPut) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.RequestPut.GetFI_Trace()
	}

	return nil
}

func (m *RequestOp_RequestPut) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.RequestPut.SetFI_Trace(trace)
	}
}

func (m *RequestOp_RequestDeleteRange) GetFI_Name() string {
	return "RequestOp_RequestDeleteRange"
}

func (m *RequestOp_RequestDeleteRange) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.RequestDeleteRange.GetFI_Trace()
	}

	return nil
}

func (m *RequestOp_RequestDeleteRange) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.RequestDeleteRange.SetFI_Trace(trace)
	}
}

func (m *RequestOp_RequestTxn) GetFI_Name() string {
	return "RequestOp_RequestTxn"
}

func (m *RequestOp_RequestTxn) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.RequestTxn.GetFI_Trace()
	}

	return nil
}

func (m *RequestOp_RequestTxn) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.RequestTxn.SetFI_Trace(trace)
	}
}

// extended
type isResponseOp_Response interface {
	isResponseOp_Response()
	MarshalTo([]byte) (int, error)
	Size() int
	GetFI_Trace() *tracer.Trace
	SetFI_Trace(*tracer.Trace)
}

func (m *ResponseOp) GetFI_Name() string {
	return "ResponseOp"
}

func (m *ResponseOp) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Response.GetFI_Trace()
	}

	return nil
}

func (m *ResponseOp) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Response.SetFI_Trace(trace)
	}
}

func (m *ResponseOp_ResponseRange) GetFI_Name() string {
	return "ResponseOp_ResponseRange"
}

func (m *ResponseOp_ResponseRange) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.ResponseRange.GetFI_Trace()
	}

	return nil
}

func (m *ResponseOp_ResponseRange) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.ResponseRange.SetFI_Trace(trace)
	}
}

func (m *ResponseOp_ResponsePut) GetFI_Name() string {
	return "ResponseOp_ResponsePut"
}

func (m *ResponseOp_ResponsePut) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.ResponsePut.GetFI_Trace()
	}

	return nil
}

func (m *ResponseOp_ResponsePut) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.ResponsePut.SetFI_Trace(trace)
	}
}

func (m *ResponseOp_ResponseDeleteRange) GetFI_Name() string {
	return "ResponseOp_ResponseDeleteRange"
}

func (m *ResponseOp_ResponseDeleteRange) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.ResponseDeleteRange.GetFI_Trace()
	}

	return nil
}

func (m *ResponseOp_ResponseDeleteRange) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.ResponseDeleteRange.SetFI_Trace(trace)
	}
}

func (m *ResponseOp_ResponseTxn) GetFI_Name() string {
	return "ResponseOp_ResponseTxn"
}

func (m *ResponseOp_ResponseTxn) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.ResponseTxn.GetFI_Trace()
	}

	return nil
}

func (m *ResponseOp_ResponseTxn) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.ResponseTxn.SetFI_Trace(trace)
	}
}

func (m *Compare) GetFI_Name() string {
	return "Compare"
}

func (m *Compare) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *TxnRequest) GetFI_Name() string {
	return "TxnRequest"
}

func (m *TxnRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *TxnResponse) GetFI_Name() string {
	return "TxnResponse"
}

func (m *TxnResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CompactionRequest) GetFI_Name() string {
	return "CompactionRequest"
}

func (m *CompactionRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CompactionResponse) GetFI_Name() string {
	return "CompactionResponse"
}

func (m *CompactionResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashRequest) GetFI_Name() string {
	return "HashRequest"
}

func (m *HashRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashKVRequest) GetFI_Name() string {
	return "HashKVRequest"
}

func (m *HashKVRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashKVResponse) GetFI_Name() string {
	return "HashKVResponse"
}

func (m *HashKVResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashResponse) GetFI_Name() string {
	return "HashResponse"
}

func (m *HashResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotRequest) GetFI_Name() string {
	return "SnapshotRequest"
}

func (m *SnapshotRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotResponse) GetFI_Name() string {
	return "SnapshotResponse"
}

func (m *SnapshotResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

// extended
type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
	MarshalTo([]byte) (int, error)
	Size() int
	GetFI_Trace() *tracer.Trace
	SetFI_Trace(*tracer.Trace)
}

func (m *WatchRequest_CreateRequest) GetFI_Name() string {
	return "WatchRequest_CreateRequest"
}

func (m *WatchRequest_CreateRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.CreateRequest.GetFI_Trace()
	}

	return nil
}

func (m *WatchRequest_CreateRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.CreateRequest.SetFI_Trace(trace)
	}
}

func (m *WatchRequest_CancelRequest) GetFI_Name() string {
	return "WatchRequest_CancelRequest"
}

func (m *WatchRequest_CancelRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.CancelRequest.GetFI_Trace()
	}

	return nil
}

func (m *WatchRequest_CancelRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.CancelRequest.SetFI_Trace(trace)
	}
}

func (m *WatchRequest_ProgressRequest) GetFI_Name() string {
	return "WatchRequest_ProgressRequest"
}

func (m *WatchRequest_ProgressRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.ProgressRequest.GetFI_Trace()
	}

	return nil
}

func (m *WatchRequest_ProgressRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.ProgressRequest.SetFI_Trace(trace)
	}
}

func (m *WatchRequest) GetFI_Name() string {
	return "WatchRequest"
}

func (m *WatchRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.RequestUnion.GetFI_Trace()
	}

	return nil
}

func (m *WatchRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.RequestUnion.SetFI_Trace(trace)
	}
}

func (m *WatchCreateRequest) GetFI_Name() string {
	return "WatchCreateRequest"
}

func (m *WatchCreateRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchCancelRequest) GetFI_Name() string {
	return "WatchCancelRequest"
}

func (m *WatchCancelRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchProgressRequest) GetFI_Name() string {
	return "WatchProgressRequest"
}

func (m *WatchProgressRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchResponse) GetFI_Name() string {
	return "WatchResponse"
}

func (m *WatchResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseGrantRequest) GetFI_Name() string {
	return "LeaseGrantRequest"
}

func (m *LeaseGrantRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseGrantResponse) GetFI_Name() string {
	return "LeaseGrantResponse"
}

func (m *LeaseGrantResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseRevokeRequest) GetFI_Name() string {
	return "LeaseRevokeRequest"
}

func (m *LeaseRevokeRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseRevokeResponse) GetFI_Name() string {
	return "LeaseRevokeResponse"
}

func (m *LeaseRevokeResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpoint) GetFI_Name() string {
	return "LeaseCheckpoint"
}

func (m *LeaseCheckpoint) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpointRequest) GetFI_Name() string {
	return "LeaseCheckpointRequest"
}

func (m *LeaseCheckpointRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpointResponse) GetFI_Name() string {
	return "LeaseCheckpointResponse"
}

func (m *LeaseCheckpointResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseKeepAliveRequest) GetFI_Name() string {
	return "LeaseKeepAliveRequest"
}

func (m *LeaseKeepAliveRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseKeepAliveResponse) GetFI_Name() string {
	return "LeaseKeepAliveResponse"
}

func (m *LeaseKeepAliveResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseTimeToLiveRequest) GetFI_Name() string {
	return "LeaseTimeToLiveRequest"
}

func (m *LeaseTimeToLiveRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseTimeToLiveResponse) GetFI_Name() string {
	return "LeaseTimeToLiveResponse"
}

func (m *LeaseTimeToLiveResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseLeasesRequest) GetFI_Name() string {
	return "LeaseLeasesRequest"
}

func (m *LeaseLeasesRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseStatus) GetFI_Name() string {
	return "LeaseStatus"
}

func (m *LeaseStatus) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseLeasesResponse) GetFI_Name() string {
	return "LeaseLeasesResponse"
}

func (m *LeaseLeasesResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Member) GetFI_Name() string {
	return "Member"
}

func (m *Member) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberAddRequest) GetFI_Name() string {
	return "MemberAddRequest"
}

func (m *MemberAddRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberAddResponse) GetFI_Name() string {
	return "MemberAddResponse"
}

func (m *MemberAddResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberRemoveRequest) GetFI_Name() string {
	return "MemberRemoveRequest"
}

func (m *MemberRemoveRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberRemoveResponse) GetFI_Name() string {
	return "MemberRemoveResponse"
}

func (m *MemberRemoveResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberUpdateRequest) GetFI_Name() string {
	return "MemberUpdateRequest"
}

func (m *MemberUpdateRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberUpdateResponse) GetFI_Name() string {
	return "MemberUpdateResponse"
}

func (m *MemberUpdateResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberListRequest) GetFI_Name() string {
	return "MemberListRequest"
}

func (m *MemberListRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberListResponse) GetFI_Name() string {
	return "MemberListResponse"
}

func (m *MemberListResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberPromoteRequest) GetFI_Name() string {
	return "MemberPromoteRequest"
}

func (m *MemberPromoteRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberPromoteResponse) GetFI_Name() string {
	return "MemberPromoteResponse"
}

func (m *MemberPromoteResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DefragmentRequest) GetFI_Name() string {
	return "DefragmentRequest"
}

func (m *DefragmentRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DefragmentResponse) GetFI_Name() string {
	return "DefragmentResponse"
}

func (m *DefragmentResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MoveLeaderRequest) GetFI_Name() string {
	return "MoveLeaderRequest"
}

func (m *MoveLeaderRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MoveLeaderResponse) GetFI_Name() string {
	return "MoveLeaderResponse"
}

func (m *MoveLeaderResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmRequest) GetFI_Name() string {
	return "AlarmRequest"
}

func (m *AlarmRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmMember) GetFI_Name() string {
	return "AlarmMember"
}

func (m *AlarmMember) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmResponse) GetFI_Name() string {
	return "AlarmResponse"
}

func (m *AlarmResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *StatusRequest) GetFI_Name() string {
	return "StatusRequest"
}

func (m *StatusRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *StatusResponse) GetFI_Name() string {
	return "StatusResponse"
}

func (m *StatusResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthEnableRequest) GetFI_Name() string {
	return "AuthEnableRequest"
}

func (m *AuthEnableRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthDisableRequest) GetFI_Name() string {
	return "AuthDisableRequest"
}

func (m *AuthDisableRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthenticateRequest) GetFI_Name() string {
	return "AuthenticateRequest"
}

func (m *AuthenticateRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserAddRequest) GetFI_Name() string {
	return "AuthUserAddRequest"
}

func (m *AuthUserAddRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGetRequest) GetFI_Name() string {
	return "AuthUserGetRequest"
}

func (m *AuthUserGetRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserDeleteRequest) GetFI_Name() string {
	return "AuthUserDeleteRequest"
}

func (m *AuthUserDeleteRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserChangePasswordRequest) GetFI_Name() string {
	return "AuthUserChangePasswordRequest"
}

func (m *AuthUserChangePasswordRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGrantRoleRequest) GetFI_Name() string {
	return "AuthUserGrantRoleRequest"
}

func (m *AuthUserGrantRoleRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserRevokeRoleRequest) GetFI_Name() string {
	return "AuthUserRevokeRoleRequest"
}

func (m *AuthUserRevokeRoleRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleAddRequest) GetFI_Name() string {
	return "AuthRoleAddRequest"
}

func (m *AuthRoleAddRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGetRequest) GetFI_Name() string {
	return "AuthRoleGetRequest"
}

func (m *AuthRoleGetRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserListRequest) GetFI_Name() string {
	return "AuthUserListRequest"
}

func (m *AuthUserListRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleListRequest) GetFI_Name() string {
	return "AuthRoleListRequest"
}

func (m *AuthRoleListRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleDeleteRequest) GetFI_Name() string {
	return "AuthRoleDeleteRequest"
}

func (m *AuthRoleDeleteRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGrantPermissionRequest) GetFI_Name() string {
	return "AuthRoleGrantPermissionRequest"
}

func (m *AuthRoleGrantPermissionRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleRevokePermissionRequest) GetFI_Name() string {
	return "AuthRoleRevokePermissionRequest"
}

func (m *AuthRoleRevokePermissionRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthEnableResponse) GetFI_Name() string {
	return "AuthEnableResponse"
}

func (m *AuthEnableResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthDisableResponse) GetFI_Name() string {
	return "AuthDisableResponse"
}

func (m *AuthDisableResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthenticateResponse) GetFI_Name() string {
	return "AuthenticateResponse"
}

func (m *AuthenticateResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserAddResponse) GetFI_Name() string {
	return "AuthUserAddResponse"
}

func (m *AuthUserAddResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGetResponse) GetFI_Name() string {
	return "AuthUserGetResponse"
}

func (m *AuthUserGetResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserDeleteResponse) GetFI_Name() string {
	return "AuthUserDeleteResponse"
}

func (m *AuthUserDeleteResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserChangePasswordResponse) GetFI_Name() string {
	return "AuthUserChangePasswordResponse"
}

func (m *AuthUserChangePasswordResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGrantRoleResponse) GetFI_Name() string {
	return "AuthUserGrantRoleResponse"
}

func (m *AuthUserGrantRoleResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserRevokeRoleResponse) GetFI_Name() string {
	return "AuthUserRevokeRoleResponse"
}

func (m *AuthUserRevokeRoleResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleAddResponse) GetFI_Name() string {
	return "AuthRoleAddResponse"
}

func (m *AuthRoleAddResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGetResponse) GetFI_Name() string {
	return "AuthRoleGetResponse"
}

func (m *AuthRoleGetResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleListResponse) GetFI_Name() string {
	return "AuthRoleListResponse"
}

func (m *AuthRoleListResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserListResponse) GetFI_Name() string {
	return "AuthUserListResponse"
}

func (m *AuthUserListResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleDeleteResponse) GetFI_Name() string {
	return "AuthRoleDeleteResponse"
}

func (m *AuthRoleDeleteResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGrantPermissionResponse) GetFI_Name() string {
	return "AuthRoleGrantPermissionResponse"
}

func (m *AuthRoleGrantPermissionResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleRevokePermissionResponse) GetFI_Name() string {
	return "AuthRoleRevokePermissionResponse"
}

func (m *AuthRoleRevokePermissionResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
