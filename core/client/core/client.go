package core

import (
	"encoding/json"
	"errors"
	"fmt"
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
	if expect := req.Expect; expect == nil || expect.ContentType == html.ContentTypeJSON { // json
		jsonData := map[string]json.RawMessage{}
		if err = json.Unmarshal(body, &jsonData); err != nil {
			fmt.Println(string(body))
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
				Uuid:        "generated by client",
				Service:     "client",
			})
		}

	} else if expect.ContentType == html.ContentTypeHTML {
		if expect.Action == data.PrintResponse {
			log.Logf("[RELOAD] Url: %v", req.URL)
			log.Logf("[RELOAD] Body: %v", string(body))
		}
	}

	return rsp, nil
}

func (c *Client) SendRequests(reqs *data.Requests) (*data.Response, error) {
	if reqs.Trace == nil {
		for _, req := range reqs.Requests {
			return c.sendRequest(&req)
		}
	}

	trace := &tracer.Trace{
		Id:      reqs.Trace.Id,
		Records: []*tracer.Record{},
		Rlfis:   reqs.Trace.Rlfis,
		Tfis:    reqs.Trace.Tfis,
	}

	var rsp *data.Response
	var err error

	for _, req := range reqs.Requests {
		record := &tracer.Record{
			Type:        tracer.RecordType_RecordSend,
			Timestamp:   time.Now().UnixNano(),
			MessageName: req.MessageName,
			Uuid:        "generated by client",
			Service:     "client",
		}

		req.Trace = &tracer.Trace{
			Id:      trace.Id,
			Records: []*tracer.Record{record},
			Rlfis:   trace.Rlfis,
			Tfis:    trace.Tfis,
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
		//log.Logf("[RELOAD] Fi-Trace: %s", traceString)
		httpReq.Header.Set("Fi-Trace", traceString)
	}

	httpRsp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return responseHandler(req, httpRsp)
}
