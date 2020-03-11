package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/golang/protobuf/proto"

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
		log.Logf("%v", httpRsp)
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
			Timestamp:   time.Now().Unix(),
			MessageName: req.MessageName,
		})
	}

	return rsp, nil
}

func (c *Client) SendRequests(reqs *data.Requests) (*data.Response, error) {
	var trace *tracer.Trace
	var rsp *data.Response
	var err error

	for _, req := range reqs.Requests {
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

		rsp, err = c.sendRequest(&req)
		if err != nil {
			return nil, err
		}

		if rsp.Trace != nil {
			trace = rsp.Trace
		}
	}

	return rsp, nil
}

func (c *Client) sendRequest(req *data.Request) (*data.Response, error) {
	var httpReq *http.Request
	var err error
	var traceString string

	if req.Trace != nil {
		traceBytes, err := proto.Marshal(req.Trace)
		if err != nil {
			return nil, err
		}

		jjj, _ := json.Marshal(req.Trace)

		trace := req.Trace
		if trace.BaseTimestamp != 0 {
			trace.BaseTimestamp = req.Trace.Records[0].Timestamp

			for _, record := range req.Trace.Records {
				record.Timestamp -= trace.BaseTimestamp
			}
		}

		log.Logf("[RELOAD] %d vs %d", len(traceBytes), len(jjj))
		traceString = string(traceBytes)
	}

	switch req.Method {
	case data.HTTPGet:
		httpReq, err = http.NewRequest("GET", req.URL, nil)
	case data.HTTPPost:
		httpReq, err = http.NewRequest("POST", req.URL, nil)
		httpReq.PostForm = req.UrlValues
	default:
		return nil, errors.New("unsupported http method")
	}

	if err != nil {
		return nil, err
	}

	if traceString != "" {
		//log.Logf("[RELOAD] Fi-Trace: %d %s", len(traceString), traceString)
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
