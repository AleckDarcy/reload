package tracer

type Tracer interface {
	GetFI_Name() string

	GetFI_Trace() *Trace
	SetFI_Trace(trace *Trace)

	GetFI_MessageType() MessageType
}

type assertion struct{}

var Assertion assertion

func (a *assertion) GetTrace(msg interface{}) (*Trace, bool) {
	if t, ok := msg.(Tracer); ok {
		trace := t.GetFI_Trace()
		return trace, trace != nil
	}

	return nil, false
}

func (a *assertion) GetLastEvent(t Tracer) (*Record, bool) {
	if trace := t.GetFI_Trace(); trace != nil {
		return trace.GetLastEvent()
	}

	return nil, false
}
