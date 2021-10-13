package tracer

type Tracer interface {
	GetFI_Name() string

	GetFI_Trace() *Trace
	SetFI_Trace(trace *Trace)

	GetFI_MessageType() MessageType
}
