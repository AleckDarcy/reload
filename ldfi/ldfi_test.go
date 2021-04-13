package ldfi

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/core/client/core"
	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
	rHtml "github.com/AleckDarcy/reload/runtime/html"
)

func TestLdfiOne(t *testing.T) {
	t.Log(time.Now().UnixNano())
}

func TestHome(t *testing.T) {
	url := "localhost"
	addr := "http://localhost"
	client := core.NewClient()
	interpreter := &Interpreter{}

	reqs := &data.Requests{
		CookieUrl: url,
		Trace: &tracer.Trace{
			Id: time.Now().UnixNano(),
			Tfis: []*tracer.TFI{

			},
		},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         addr + "/?render=json",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeJSON,
					Action:      data.DeserializeTrace,
				},
			},
		},
	}

	resp, err := client.SendRequests(reqs)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(resp.Body))

	bytes, _ := json.Marshal(resp.Trace)

	t.Log(string(bytes))
	/*
		1. Make A DAG Representation of the trace
	*/
	interpreter.handleTrace(reqs, resp)
	fmt.Println("Making DAG Out of Trace to Feed to LDFI")

	/*
		This is where I need to pass information to
		LDFI interpreter, and get a fault injection
		to reason about injecting
	*/
}