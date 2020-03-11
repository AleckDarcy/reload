package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/AleckDarcy/reload/core/client/data"
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
	if httpRsp.Header.Get(html.ContentType) != html.ContentTypeJSON {
		return nil, nil
	}

	body, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	jsonData := map[string]json.RawMessage{}
	if err = json.Unmarshal(body, jsonData); err != nil {
		return nil, err
	}

	rsp := &data.Response{Body: body}
	if traceJSON, ok := jsonData["fi_trace"]; ok {
		rsp.Trace = &tracer.Trace{}
		if err = json.Unmarshal(traceJSON, rsp.Trace); err != nil {
			return nil, err
		}

		rsp.Trace.Records = append(rsp.Trace.Records, &tracer.Record{
			Type:        tracer.RecordType_RecordReceive,
			Timestamp:   time.Now().Unix(),
			MessageName: req.MessageName,
		})
	}

	return rsp, nil
}

func (c *Client) SendRequests(reqs []*data.Request) (*tracer.Trace, error) {
	trace := &tracer.Trace{}

	for _, req := range reqs {
		if trace != nil && req.Trace == nil {
			req.Trace = trace
		}

		if req.Trace != nil {
			req.Trace.Records = append(req.Trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordSend,
				Timestamp:   time.Now().Unix(),
				MessageName: req.MessageName,
			})
		}

		rsp, err := c.sendRequest(req)
		if err != nil {
			return nil, err
		}

		if rsp.Trace != nil {
			trace = rsp.Trace
		}
	}

	return trace, nil
}

func (c *Client) sendRequest(req *data.Request) (*data.Response, error) {
	var httpRsp *http.Response
	var err error

	switch req.Method {
	case data.HTTPGet:
		httpRsp, err = c.Client.Get(req.URL)
	case data.HTTPPost:
		if req.Trace != nil {
			traceBytes, err := json.Marshal(req.Trace)
			if err != nil {
				return nil, err
			}

			req.UrlValues["Fi-Trace"] = []string{string(traceBytes)}
		}

		httpRsp, err = c.Client.PostForm(req.URL, req.UrlValues)
	default:
		return nil, errors.New("unsupported http method")
	}

	if err != nil {
		return nil, err
	}

	if req.ResponseHandler != nil {
		return req.ResponseHandler(req, httpRsp)
	}

	return responseHandler(req, httpRsp)
}
