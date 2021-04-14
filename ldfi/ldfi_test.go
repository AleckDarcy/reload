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

	time.Sleep(1)
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
	interpreter.forwardStep(reqs, resp)
	fmt.Println("Making DAG Out of Trace to Feed to LDFI")

	/*
		This is where I need to pass information to
		LDFI interpreter, and get a fault injection
		to reason about injecting
	*/
}

func Test1(t *testing.T) {
	traces := []*tracer.Trace{
		{ // RLFI: crash CurrencyService, Fault: no
			Tfis: []*tracer.TFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"GetSupportedCurrenciesRequest"},
				},
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"CurrencyConversionRequest"},
				},
			},
		},
		{ // RLFI: crash AdService, Fault: no
			Tfis: []*tracer.TFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"AdRequest"},
				},
			},
		},
		{ // RLFI: crash CurrencyService and AdService, Fault: yes
			Tfis: []*tracer.TFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"GetSupportedCurrenciesRequest"},
				},
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"CurrencyConversionRequest"},
				},
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"AdRequest"},
				},
			},
		},
		{ // TFI: crash CurrencyService when receiving CurrencyConversionRequest two times, Fault: yes
			Tfis: []*tracer.TFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"CurrencyConversionRequest"},
					After: []*tracer.TFIMeta{
						{Name: "CurrencyConversionRequest", Times: 0},
					},
				},
			},
		},
	}

	interpreter := &Interpreter{}
	for i := 0; i < len(traces); i++ {
		if i != 2 {
			continue
		}

		client := core.NewClient()

		reqs := &data.Requests{
			CookieUrl: "localhost",
			Trace:     traces[i],
			Requests: []data.Request{
				{
					Method:      data.HTTPGet,
					URL:         "http://localhost/?render=json",
					MessageName: "home",
					Trace:       traces[i],
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeJSON,
						Action:      data.PrintResponse | data.DeserializeTrace,
					},
				},
			},
		}
		reqs.Trace.Id = time.Now().UnixNano()

		rsp, err := client.SendRequests(reqs)
		if err != nil {
			t.Errorf("%d err: %v", i, err)
		} else {
			t.Logf("body: %s", string(rsp.Body))
		}


		interpreter.forwardStep(reqs, rsp)
	}

}

func TestLDFI(t *testing.T) {
	/*
		Initial Test will be testing Home
	 */
	rounds := 3
	url := "localhost"
	addr := "http://localhost"
	client := core.NewClient()
	interpreter := &Interpreter{}

	/*
		The request we are testing
	 */
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
	interpreter.forwardStep(reqs, resp)
	fmt.Println("Making DAG Out of Trace to Feed to LDFI")

	/*
		This is where I need to pass information to
		LDFI interpreter, and get a fault injection
		to reason about injecting
	*/

	for i := 0; i < rounds; i++ {
		fmt.Printf("Round: %d\n", i)
	}

}