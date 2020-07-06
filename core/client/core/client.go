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

	rHtml "github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/client/parser"
	"github.com/AleckDarcy/reload/core/log"
	"github.com/AleckDarcy/reload/core/tracer"
	"golang.org/x/net/html"
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

func (c *Client) responseHandler(req *data.Request, httpReq *http.Request) (*data.Response, error) {
	start := time.Now().UnixNano()
	httpRsp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(httpRsp.Body)
	end := time.Now().UnixNano()
	if err != nil {
		return nil, err
	}
	rsp := &data.Response{Latency: end - start, Body: body}

	//log.Logf("[RELOAD] Content-Type: %s", httpRsp.Header.Get(rHtml.ContentType))
	if req.Trace == nil || req.Expect == nil {

	} else if expect := req.Expect; expect.ContentType == rHtml.ContentTypeHTML {
		print := expect.Action&data.PrintResponse != 0
		deTrace := expect.Action&data.DeserializeTrace != 0

		if print {
			log.Logf("[RELOAD] Url: %v", req.URL)
			log.Logf("[RELOAD] Body: %v", string(body))
		}

		if deTrace {
			node, _ := html.Parse(strings.NewReader(string(body)))
			if traceNode := parser.GetElementByClass(node, "trace"); traceNode != nil {
				traceString := parser.GetJSON(traceNode)
				if print {
					log.Logf("[RELOAD] Trace: %v", traceString)
				}

				rsp.Trace = &tracer.Trace{}
				if err = json.Unmarshal([]byte(traceString), rsp.Trace); err != nil {
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
		}
	} else if expect.ContentType == rHtml.ContentTypeJSON { // json
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
		httpReq, err = http.NewRequest(rHtml.MethodGet, req.URL, nil)
	case data.HTTPPost:
		body := strings.NewReader(req.UrlValues.Encode())
		httpReq, err = http.NewRequest(rHtml.MethodPost, req.URL, body)
		if err == nil {
			httpReq.Header.Set(rHtml.ContentType, rHtml.ContentTypeMIMEPostForm)
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

	return c.responseHandler(req, httpReq)
}
