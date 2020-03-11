package data

import (
	"net/http"
	"net/url"

	"github.com/AleckDarcy/reload/core/tracer"
)

type HTTPMethod int64

const (
	_ HTTPMethod = iota
	HTTPPost
	HTTPGet
)

type ResponseHandler func(req *Request, rsp *http.Response) (*Response, error)

type Request struct {
	Method    HTTPMethod
	URL       string
	UrlValues url.Values // HTTPPost

	// tracing
	MessageName string
	Trace       *tracer.Trace

	ResponseHandler ResponseHandler
}

type Response struct {
	Body  []byte
	Trace *tracer.Trace
}
