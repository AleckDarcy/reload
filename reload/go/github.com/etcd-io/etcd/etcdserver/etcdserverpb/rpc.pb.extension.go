package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *ResponseHeader) MessageName() string {
	return "ResponseHeader"
}

func (m *ResponseHeader) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RangeRequest) MessageName() string {
	return "RangeRequest"
}

func (m *RangeRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RangeResponse) MessageName() string {
	return "RangeResponse"
}

func (m *RangeResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *PutRequest) MessageName() string {
	return "PutRequest"
}

func (m *PutRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *PutResponse) MessageName() string {
	return "PutResponse"
}

func (m *PutResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DeleteRangeRequest) MessageName() string {
	return "DeleteRangeRequest"
}

func (m *DeleteRangeRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DeleteRangeResponse) MessageName() string {
	return "DeleteRangeResponse"
}

func (m *DeleteRangeResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

// extended
type isRequestOp_Request interface {
	isRequestOp_Request()
	MarshalTo([]byte) (int, error)
	Size() int
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func (m *RequestOp) MessageName() string {
	return "RequestOp"
}

func (m *RequestOp) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Request.GetTrace()
	}

	return nil
}

func (m *RequestOp) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Request.SetTrace(trace)
	}
}

func (m *RequestOp_RequestRange) MessageName() string {
	return "RequestOp_RequestRange"
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

func (m *RequestOp_RequestPut) MessageName() string {
	return "RequestOp_RequestPut"
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

func (m *RequestOp_RequestDeleteRange) MessageName() string {
	return "RequestOp_RequestDeleteRange"
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

func (m *RequestOp_RequestTxn) MessageName() string {
	return "RequestOp_RequestTxn"
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

// extended
type isResponseOp_Response interface {
	isResponseOp_Response()
	MarshalTo([]byte) (int, error)
	Size() int
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func (m *ResponseOp) MessageName() string {
	return "ResponseOp"
}

func (m *ResponseOp) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Response.GetTrace()
	}

	return nil
}

func (m *ResponseOp) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Response.SetTrace(trace)
	}
}

func (m *ResponseOp_ResponseRange) MessageName() string {
	return "ResponseOp_ResponseRange"
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

func (m *ResponseOp_ResponsePut) MessageName() string {
	return "ResponseOp_ResponsePut"
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

func (m *ResponseOp_ResponseDeleteRange) MessageName() string {
	return "ResponseOp_ResponseDeleteRange"
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

func (m *ResponseOp_ResponseTxn) MessageName() string {
	return "ResponseOp_ResponseTxn"
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

func (m *Compare) MessageName() string {
	return "Compare"
}

func (m *Compare) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *TxnRequest) MessageName() string {
	return "TxnRequest"
}

func (m *TxnRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *TxnResponse) MessageName() string {
	return "TxnResponse"
}

func (m *TxnResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CompactionRequest) MessageName() string {
	return "CompactionRequest"
}

func (m *CompactionRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *CompactionResponse) MessageName() string {
	return "CompactionResponse"
}

func (m *CompactionResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashRequest) MessageName() string {
	return "HashRequest"
}

func (m *HashRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashKVRequest) MessageName() string {
	return "HashKVRequest"
}

func (m *HashKVRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashKVResponse) MessageName() string {
	return "HashKVResponse"
}

func (m *HashKVResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *HashResponse) MessageName() string {
	return "HashResponse"
}

func (m *HashResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotRequest) MessageName() string {
	return "SnapshotRequest"
}

func (m *SnapshotRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *SnapshotResponse) MessageName() string {
	return "SnapshotResponse"
}

func (m *SnapshotResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

// extended
type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
	MarshalTo([]byte) (int, error)
	Size() int
	GetTrace() *tracer.Trace
	SetTrace(*tracer.Trace)
}

func (m *WatchRequest_CreateRequest) MessageName() string {
	return "WatchRequest_CreateRequest"
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

func (m *WatchRequest_CancelRequest) MessageName() string {
	return "WatchRequest_CancelRequest"
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

func (m *WatchRequest_ProgressRequest) MessageName() string {
	return "WatchRequest_ProgressRequest"
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

func (m *WatchRequest) MessageName() string {
	return "WatchRequest"
}

func (m *WatchRequest) GetTrace() *tracer.Trace {
	if m != nil {
		return m.RequestUnion.GetTrace()
	}

	return nil
}

func (m *WatchRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.RequestUnion.SetTrace(trace)
	}
}

func (m *WatchCreateRequest) MessageName() string {
	return "WatchCreateRequest"
}

func (m *WatchCreateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchCancelRequest) MessageName() string {
	return "WatchCancelRequest"
}

func (m *WatchCancelRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchProgressRequest) MessageName() string {
	return "WatchProgressRequest"
}

func (m *WatchProgressRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *WatchResponse) MessageName() string {
	return "WatchResponse"
}

func (m *WatchResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseGrantRequest) MessageName() string {
	return "LeaseGrantRequest"
}

func (m *LeaseGrantRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseGrantResponse) MessageName() string {
	return "LeaseGrantResponse"
}

func (m *LeaseGrantResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseRevokeRequest) MessageName() string {
	return "LeaseRevokeRequest"
}

func (m *LeaseRevokeRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseRevokeResponse) MessageName() string {
	return "LeaseRevokeResponse"
}

func (m *LeaseRevokeResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpoint) MessageName() string {
	return "LeaseCheckpoint"
}

func (m *LeaseCheckpoint) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpointRequest) MessageName() string {
	return "LeaseCheckpointRequest"
}

func (m *LeaseCheckpointRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseCheckpointResponse) MessageName() string {
	return "LeaseCheckpointResponse"
}

func (m *LeaseCheckpointResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseKeepAliveRequest) MessageName() string {
	return "LeaseKeepAliveRequest"
}

func (m *LeaseKeepAliveRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseKeepAliveResponse) MessageName() string {
	return "LeaseKeepAliveResponse"
}

func (m *LeaseKeepAliveResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseTimeToLiveRequest) MessageName() string {
	return "LeaseTimeToLiveRequest"
}

func (m *LeaseTimeToLiveRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseTimeToLiveResponse) MessageName() string {
	return "LeaseTimeToLiveResponse"
}

func (m *LeaseTimeToLiveResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseLeasesRequest) MessageName() string {
	return "LeaseLeasesRequest"
}

func (m *LeaseLeasesRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseStatus) MessageName() string {
	return "LeaseStatus"
}

func (m *LeaseStatus) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *LeaseLeasesResponse) MessageName() string {
	return "LeaseLeasesResponse"
}

func (m *LeaseLeasesResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Member) MessageName() string {
	return "Member"
}

func (m *Member) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberAddRequest) MessageName() string {
	return "MemberAddRequest"
}

func (m *MemberAddRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberAddResponse) MessageName() string {
	return "MemberAddResponse"
}

func (m *MemberAddResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberRemoveRequest) MessageName() string {
	return "MemberRemoveRequest"
}

func (m *MemberRemoveRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberRemoveResponse) MessageName() string {
	return "MemberRemoveResponse"
}

func (m *MemberRemoveResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberUpdateRequest) MessageName() string {
	return "MemberUpdateRequest"
}

func (m *MemberUpdateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberUpdateResponse) MessageName() string {
	return "MemberUpdateResponse"
}

func (m *MemberUpdateResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberListRequest) MessageName() string {
	return "MemberListRequest"
}

func (m *MemberListRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberListResponse) MessageName() string {
	return "MemberListResponse"
}

func (m *MemberListResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberPromoteRequest) MessageName() string {
	return "MemberPromoteRequest"
}

func (m *MemberPromoteRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MemberPromoteResponse) MessageName() string {
	return "MemberPromoteResponse"
}

func (m *MemberPromoteResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DefragmentRequest) MessageName() string {
	return "DefragmentRequest"
}

func (m *DefragmentRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *DefragmentResponse) MessageName() string {
	return "DefragmentResponse"
}

func (m *DefragmentResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MoveLeaderRequest) MessageName() string {
	return "MoveLeaderRequest"
}

func (m *MoveLeaderRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *MoveLeaderResponse) MessageName() string {
	return "MoveLeaderResponse"
}

func (m *MoveLeaderResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmRequest) MessageName() string {
	return "AlarmRequest"
}

func (m *AlarmRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmMember) MessageName() string {
	return "AlarmMember"
}

func (m *AlarmMember) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AlarmResponse) MessageName() string {
	return "AlarmResponse"
}

func (m *AlarmResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *StatusRequest) MessageName() string {
	return "StatusRequest"
}

func (m *StatusRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *StatusResponse) MessageName() string {
	return "StatusResponse"
}

func (m *StatusResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthEnableRequest) MessageName() string {
	return "AuthEnableRequest"
}

func (m *AuthEnableRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthDisableRequest) MessageName() string {
	return "AuthDisableRequest"
}

func (m *AuthDisableRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthenticateRequest) MessageName() string {
	return "AuthenticateRequest"
}

func (m *AuthenticateRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserAddRequest) MessageName() string {
	return "AuthUserAddRequest"
}

func (m *AuthUserAddRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGetRequest) MessageName() string {
	return "AuthUserGetRequest"
}

func (m *AuthUserGetRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserDeleteRequest) MessageName() string {
	return "AuthUserDeleteRequest"
}

func (m *AuthUserDeleteRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserChangePasswordRequest) MessageName() string {
	return "AuthUserChangePasswordRequest"
}

func (m *AuthUserChangePasswordRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGrantRoleRequest) MessageName() string {
	return "AuthUserGrantRoleRequest"
}

func (m *AuthUserGrantRoleRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserRevokeRoleRequest) MessageName() string {
	return "AuthUserRevokeRoleRequest"
}

func (m *AuthUserRevokeRoleRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleAddRequest) MessageName() string {
	return "AuthRoleAddRequest"
}

func (m *AuthRoleAddRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGetRequest) MessageName() string {
	return "AuthRoleGetRequest"
}

func (m *AuthRoleGetRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserListRequest) MessageName() string {
	return "AuthUserListRequest"
}

func (m *AuthUserListRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleListRequest) MessageName() string {
	return "AuthRoleListRequest"
}

func (m *AuthRoleListRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleDeleteRequest) MessageName() string {
	return "AuthRoleDeleteRequest"
}

func (m *AuthRoleDeleteRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGrantPermissionRequest) MessageName() string {
	return "AuthRoleGrantPermissionRequest"
}

func (m *AuthRoleGrantPermissionRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleRevokePermissionRequest) MessageName() string {
	return "AuthRoleRevokePermissionRequest"
}

func (m *AuthRoleRevokePermissionRequest) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthEnableResponse) MessageName() string {
	return "AuthEnableResponse"
}

func (m *AuthEnableResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthDisableResponse) MessageName() string {
	return "AuthDisableResponse"
}

func (m *AuthDisableResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthenticateResponse) MessageName() string {
	return "AuthenticateResponse"
}

func (m *AuthenticateResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserAddResponse) MessageName() string {
	return "AuthUserAddResponse"
}

func (m *AuthUserAddResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGetResponse) MessageName() string {
	return "AuthUserGetResponse"
}

func (m *AuthUserGetResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserDeleteResponse) MessageName() string {
	return "AuthUserDeleteResponse"
}

func (m *AuthUserDeleteResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserChangePasswordResponse) MessageName() string {
	return "AuthUserChangePasswordResponse"
}

func (m *AuthUserChangePasswordResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserGrantRoleResponse) MessageName() string {
	return "AuthUserGrantRoleResponse"
}

func (m *AuthUserGrantRoleResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserRevokeRoleResponse) MessageName() string {
	return "AuthUserRevokeRoleResponse"
}

func (m *AuthUserRevokeRoleResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleAddResponse) MessageName() string {
	return "AuthRoleAddResponse"
}

func (m *AuthRoleAddResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGetResponse) MessageName() string {
	return "AuthRoleGetResponse"
}

func (m *AuthRoleGetResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleListResponse) MessageName() string {
	return "AuthRoleListResponse"
}

func (m *AuthRoleListResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthUserListResponse) MessageName() string {
	return "AuthUserListResponse"
}

func (m *AuthUserListResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleDeleteResponse) MessageName() string {
	return "AuthRoleDeleteResponse"
}

func (m *AuthRoleDeleteResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleGrantPermissionResponse) MessageName() string {
	return "AuthRoleGrantPermissionResponse"
}

func (m *AuthRoleGrantPermissionResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *AuthRoleRevokePermissionResponse) MessageName() string {
	return "AuthRoleRevokePermissionResponse"
}

func (m *AuthRoleRevokePermissionResponse) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
