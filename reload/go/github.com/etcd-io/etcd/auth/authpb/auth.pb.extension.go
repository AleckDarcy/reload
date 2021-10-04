package authpb

import "github.com/AleckDarcy/reload/core/tracer"

func (m *UserAddOptions) MessageName() string {
	return "UserAddOptions"
}

func (m *UserAddOptions) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *UserAddOptions) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *User) MessageName() string {
	return "User"
}

func (m *User) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *User) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Permission) MessageName() string {
	return "Permission"
}

func (m *Permission) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Permission) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}

func (m *Role) MessageName() string {
	return "Role"
}

func (m *Role) GetTrace() *tracer.Trace {
	if m != nil {
		return m.Trace
	}

	return nil
}

func (m *Role) SetTrace(trace *tracer.Trace) {
	if m != nil {
		m.Trace = trace
	}
}
