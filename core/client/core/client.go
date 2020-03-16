package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"
	"github.com/AleckDarcy/reload/runtime/html"
)

type Client struct {
	Client http.Client
}

func NewClient() *Client {
	cookieJar, _ := cookiejar.New(nil)

	return &Client{
		Client: http.Client{
			Jar: cookieJar,
		},
	}
}

func responseHandler(req *data.Request, httpRsp *http.Response) (*data.Response, error) {
	body, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}
	rsp := &data.Response{Body: body}

	log.Logf("[RELOAD] Content-Type: %s", httpRsp.Header.Get(html.ContentType))
	if httpRsp.Header.Get(html.ContentType) != html.ContentTypeJSON {
		log.Logf("[RELOAD] Url: %v", req.URL)
		log.Logf("[RELOAD] Body: %v", string(body))
		return rsp, nil
	}

	jsonData := map[string]json.RawMessage{}
	if err = json.Unmarshal(body, &jsonData); err != nil {
		return nil, err
	}

	if traceJSON, ok := jsonData["fi_trace"]; ok {
		rsp.Trace = &tracer.Trace{}
		if err = json.Unmarshal(traceJSON, rsp.Trace); err != nil {
			return nil, err
		}

		rsp.Trace.Records = append(rsp.Trace.Records, &tracer.Record{
			Type:        tracer.RecordType_RecordReceive,
			Timestamp:   time.Now().UnixNano(),
			MessageName: req.MessageName,
		})
	}

	return rsp, nil
}

func (c *Client) SendRequests(reqs *data.Requests) (*data.Response, error) {
	if reqs.Trace == nil {
		return nil, nil
	}

	trace := &tracer.Trace{
		Id:      reqs.Trace.Id,
		Records: []*tracer.Record{},
		Rlfi:    reqs.Trace.Rlfi,
		Tfi:     reqs.Trace.Tfi,
	}

	var rsp *data.Response
	var err error

	for _, req := range reqs.Requests {
		record := &tracer.Record{
			Type:        tracer.RecordType_RecordSend,
			Timestamp:   time.Now().UnixNano(),
			MessageName: req.MessageName,
		}

		req.Trace = &tracer.Trace{
			Id:      trace.Id,
			Records: []*tracer.Record{record},
			Rlfi:    trace.Rlfi,
			Tfi:     trace.Tfi,
		}

		rsp, err = c.sendRequest(&req)
		if err != nil {
			return nil, err
		}

		trace.Records = append(trace.Records, record)
		if rsp.Trace != nil {
			trace.Records = append(trace.Records, rsp.Trace.Records...)
		}
	}

	rsp.Trace = trace

	return rsp, nil
}

func (c *Client) sendRequest(req *data.Request) (*data.Response, error) {
	var httpReq *http.Request
	var err error
	var traceString string

	if req.Trace != nil {
		traceBytes, err := json.Marshal(req.Trace)
		if err != nil {
			return nil, err
		}

		traceString = string(traceBytes)
	}

	switch req.Method {
	case data.HTTPGet:
		httpReq, err = http.NewRequest(html.MethodGet, req.URL, nil)
	case data.HTTPPost:
		body := strings.NewReader(req.UrlValues.Encode())
		httpReq, err = http.NewRequest(html.MethodPost, req.URL, body)
		if err == nil {
			httpReq.Header.Set(html.ContentType, html.ContentTypeMIMEPostForm)
		}
	default:
		return nil, errors.New("unsupported http method")
	}

	if err != nil {
		return nil, err
	}

	if traceString != "" {
		log.Logf("[RELOAD] Fi-Trace: %d %s", len(traceString), traceString)
		httpReq.Header.Set("Fi-Trace", traceString)
	}

	httpRsp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if req.ResponseHandler != nil {
		return req.ResponseHandler(req, httpRsp)
	}

	return responseHandler(req, httpRsp)
}
