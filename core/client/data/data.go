package data

import (
	"net/url"

	"github.com/AleckDarcy/reload/core/tracer"
)

type HTTPMethod = int64

const (
	_ HTTPMethod = iota
	HTTPGet
	HTTPPost
)

type Requests struct {
	CookieUrl string
	Trace     *tracer.Trace

	Requests []Request
}

type Request struct {
	Method    HTTPMethod
	URL       string
	UrlValues url.Values // HTTPPost

	// tracing
	MessageName string
	Trace       *tracer.Trace

	Expect *ExpectedResponse
}

type ActionResponse int64

const (
	PrintResponse ActionResponse = iota << 1
	DeserializeTrace
	CustomizedRspFunc
)

type ExpectedResponse struct {
	ContentType string

	Action ActionResponse
}

type Response struct {
	Latency int64
	Body    []byte
	Trace   *tracer.Trace
}
