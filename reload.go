package reload

import "log"

//go:generate protoc --go_out=. core/tracer/message.proto

func init() {
	log.Printf("[RELOAD] reload initialized for tracing and fault injection")
}
