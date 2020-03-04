package errors

import "errors"

// ProtoBuffer
var ()

// Tracer
var (
	ErrorTracer_ThreadIDNotFound = errors.New("error Tracer: thread id not found")
	ErrorTracer_TraceIDNotFound  = errors.New("error Tracer: trace id not found")
)

// FI
var (
	StringFI_      = "error FI"
	StringFI_Delay = "FI delay triggered"

	ErrorFI_RLFI_      = errors.New("error FI: RLFI unknown triggered")
	ErrorFI_RLFI_Crash = errors.New("error FI: RLFI crash triggered")
	ErrorFI_RLFI_Delay = error(nil)
	ErrorFI_TFI_       = errors.New("error FI: TFI unknown triggered")
	ErrorFI_TFI_Crash  = errors.New("error FI: TFI crash triggered")
	ErrorFI_TFI_Delay  = error(nil)
)
