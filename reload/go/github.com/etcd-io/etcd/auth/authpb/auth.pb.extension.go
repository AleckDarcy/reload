package authpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *UserAddOptions) GetFI_Name() string {
	return "UserAddOptions"
}

func (m *UserAddOptions) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *UserAddOptions) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *User) GetFI_Name() string {
	return "User"
}

func (m *User) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *User) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Permission) GetFI_Name() string {
	return "Permission"
}

func (m *Permission) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Permission) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Role) GetFI_Name() string {
	return "Role"
}

func (m *Role) GetFI_Trace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Role) SetFI_Trace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
