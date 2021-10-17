package etcdserverpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *RequestHeader) GetFI_Name() string {
	return "RequestHeader"
}

func (m *RequestHeader) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *RequestHeader) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *RequestHeader) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *InternalRaftRequest) GetFI_Name() string {
	if m.V2 != nil {
		return m.V2.GetFI_Name()
	} else if m.Range != nil {
		return m.Range.GetFI_Name()
	} else if m.Put != nil {
		return m.Put.GetFI_Name()
	} else if m.DeleteRange != nil {
		return m.DeleteRange.GetFI_Name()
	} else if m.Txn != nil {
		return m.Txn.GetFI_Name()
	} else if m.Compaction != nil {
		return m.Compaction.GetFI_Name()
	} else if m.LeaseGrant != nil {
		return m.LeaseGrant.GetFI_Name()
	} else if m.LeaseRevoke != nil {
		return m.LeaseRevoke.GetFI_Name()
	} else if m.Alarm != nil {
		return m.Alarm.GetFI_Name()
	} else if m.LeaseCheckpoint != nil {
		return m.LeaseCheckpoint.GetFI_Name()
	} else if m.AuthEnable != nil {
		return m.AuthEnable.GetFI_Name()
	} else if m.AuthDisable != nil {
		return m.AuthDisable.GetFI_Name()
	} else if m.Authenticate != nil {
		return m.Authenticate.GetFI_Name()
	} else if m.AuthUserAdd != nil {
		return m.AuthUserAdd.GetFI_Name()
	} else if m.AuthUserDelete != nil {
		return m.AuthUserDelete.GetFI_Name()
	} else if m.AuthUserGet != nil {
		return m.AuthUserGet.GetFI_Name()
	} else if m.AuthUserChangePassword != nil {
		return m.AuthUserChangePassword.GetFI_Name()
	} else if m.AuthUserGrantRole != nil {
		return m.AuthUserGrantRole.GetFI_Name()
	} else if m.AuthUserRevokeRole != nil {
		return m.AuthUserRevokeRole.GetFI_Name()
	} else if m.AuthUserList != nil {
		return m.AuthUserList.GetFI_Name()
	} else if m.AuthRoleList != nil {
		return m.AuthRoleList.GetFI_Name()
	} else if m.AuthRoleAdd != nil {
		return m.AuthRoleAdd.GetFI_Name()
	} else if m.AuthRoleDelete != nil {
		return m.AuthRoleDelete.GetFI_Name()
	} else if m.AuthRoleGet != nil {
		return m.AuthRoleGet.GetFI_Name()
	} else if m.AuthRoleGrantPermission != nil {
		return m.AuthRoleGrantPermission.GetFI_Name()
	} else if m.AuthRoleRevokePermission != nil {
		return m.AuthRoleRevokePermission.GetFI_Name()
	} else {
		return "InternalRaftRequest"
	}
}

func (m *InternalRaftRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *InternalRaftRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *InternalRaftRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *EmptyResponse) GetFI_Name() string {
	return "EmptyResponse"
}

func (m *EmptyResponse) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *EmptyResponse) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *EmptyResponse) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *InternalAuthenticateRequest) GetFI_Name() string {
	return "InternalAuthenticateRequest"
}

func (m *InternalAuthenticateRequest) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *InternalAuthenticateRequest) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *InternalAuthenticateRequest) GetFI_MessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}
