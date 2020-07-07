package core

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	rHtml "github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
)

func TestName1(t *testing.T) {
	t.Log(time.Now().UnixNano())
}

const NTests = 100
const NRound = 1

var addr = "http://34.83.167.255"

//var addr = "http://localhost"

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
						{Name: "CurrencyConversionRequest", Times: 1},
					},
				},
			},
		},
	}

	for i := 0; i < len(traces); i++ {
		if i != 3 {
			continue
		}

		client := NewClient()

		reqs := &data.Requests{
			CookieUrl: "localhost",
			Trace:     traces[i],
			Requests: []data.Request{
				{
					Method:      data.HTTPGet,
					URL:         "http://localhost",
					MessageName: "home",
					Trace:       traces[i],
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeHTML,
						//Action:      data.PrintResponse,
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
	}
}

type perf struct {
	PerfCases []perfCase
}

type perfCase struct {
	PerfNClients []perfNClient
}

type perfNClient struct {
	PerfRoundsAvg perfRoundsAvg

	PerfRounds []perfRound `json:"-"`
}

type perfRound struct {
	PerfRequestsAvg perfRequestsAvg
	PerfRequests    []perfRequest

	ErrorCount int64
	Throughput float64
}

type perfRoundsAvg struct {
	NClient         int
	PerfRequestsAvg perfRequestsAvg

	ErrorCount float64
	Throughput float64
}

type perfRequest struct {
	Latency   int64
	FELatency int64
}

type perfRequestsAvg struct {
	Max, Min int64
	Avg      float64

	FEMax, FEMin int64
	FEAvg        float64
}

func TestConcurrency(t *testing.T) {
	nClients := []int{1, 2, 4, 8, 16, 32, 64, 128}

	cases := []*tracer.Trace{
		nil,
		//{},
		//{
		//	Tfis: []*tracer.TFI{
		//		{
		//			Type: tracer.FaultType_FaultCrash,
		//			Name: []string{"GetSupportedCurrenciesRequest"},
		//		},
		//		{
		//			Type: tracer.FaultType_FaultCrash,
		//			Name: []string{"CurrencyConversionRequest"},
		//		},
		//		//{
		//		//	Type: tracer.FaultType_FaultCrash,
		//		//	Name: []string{"AdRequest"},
		//		//},
		//	},
		//},
		//{
		//	Tfis: []*tracer.TFI{
		//		{
		//			Type: tracer.FaultType_FaultCrash,
		//			Name: []string{"CurrencyConversionRequest"},
		//			After: []*tracer.TFIMeta{
		//				{Name: "CurrencyConversionRequest", Times: 2},
		//			},
		//		},
		//	},
		//},
	}

	p := &perf{
		PerfCases: make([]perfCase, len(cases)),
	}

	for caseI, case_ := range cases {
		perfCase := &p.PerfCases[caseI]
		perfCase.PerfNClients = make([]perfNClient, len(nClients))

		for nClientI, nCLient := range nClients {
			fmt.Printf("case %d, nClient %d\n", caseI, nCLient)

			perfNClient := &perfCase.PerfNClients[nClientI]
			perfNClient.PerfRounds = make([]perfRound, NRound)

			clients := make([]*Client, nCLient)
			signals := make(chan struct{}, nCLient)
			requests := make([]*data.Requests, nCLient)

			for i := 0; i < nCLient; i++ {
				clients[i] = NewClient()

				request := &data.Requests{
					CookieUrl: "localhost",
					Trace:     case_,
					Requests: []data.Request{
						{
							Method:      data.HTTPGet,
							URL:         addr,
							MessageName: "home",
							Expect: &data.ExpectedResponse{
								ContentType: rHtml.ContentTypeHTML,
								Action:      data.DeserializeTrace,
							},
						},
					},
				}

				if case_ != nil {
					trace := *case_
					request.Trace = &trace
				}

				requests[i] = request
			}

			for roundI := 0; roundI < NRound; roundI++ {
				perfRound := &perfNClient.PerfRounds[roundI]
				perfRound.PerfRequests = make([]perfRequest, NTests)

				traceIDOffset := time.Now().UnixNano()
				traceID := int64(0)

				start := time.Now()

				for i := 0; i < nCLient; i++ {
					go func(i int, signals chan struct{}) {
						client := clients[i]
						reqs := requests[i]

						for {
							traceID := atomic.AddInt64(&traceID, 1)
							if traceID > NTests {
								break
							}

							if reqs.Trace != nil {
								reqs.Trace.Id = traceID + traceIDOffset
							}

							perfRequest := &perfRound.PerfRequests[traceID-1]

							rsp, err := client.SendRequests(reqs)
							perfRequest.Latency = rsp.Latency

							if err != nil {
								t.Error(err)
								atomic.AddInt64(&perfRound.ErrorCount, 1)
							} else if trace := rsp.Trace; trace != nil {
								if entryCount := len(trace.Records); entryCount >= 4 {
									perfRequest.FELatency = trace.Records[entryCount-2].Timestamp - trace.Records[1].Timestamp
								}
							}
						}

						signals <- struct{}{}
					}(i, signals)
				}

				for i := 0; i < nCLient; i++ {
					<-signals
				}

				end := time.Now()

				perfRequestsAvg := &perfRound.PerfRequestsAvg
				perfRequestsAvg.Min = math.MaxInt64
				perfRequestsAvg.FEMin = math.MaxInt64

				FECount := 1.0
				for _, perfRequest := range perfRound.PerfRequests {
					perfRequestsAvg.Avg += float64(perfRequest.Latency) / NTests
					perfRequestsAvg.FEAvg += float64(perfRequest.FELatency) / NTests

					if perfRequest.Latency > perfRequestsAvg.Max {
						perfRequestsAvg.Max = perfRequest.Latency
					}
					if perfRequest.Latency < perfRequestsAvg.Min {
						perfRequestsAvg.Min = perfRequest.Latency
					}

					if perfRequest.FELatency != 0 {
						FECount++
						if perfRequest.FELatency > perfRequestsAvg.FEMax {
							perfRequestsAvg.FEMax = perfRequest.FELatency
						}

						if perfRequest.FELatency < perfRequestsAvg.FEMin {
							perfRequestsAvg.FEMin = perfRequest.FELatency
						}
					}
				}

				perfRequestsAvg.FEAvg *= NTests / FECount

				perfRound.Throughput = NTests * 1e9 / float64(end.Sub(start).Nanoseconds())
			}
		}
	}

	for caseI := range p.PerfCases {
		perfCase := &p.PerfCases[caseI]
		for nClientI := range perfCase.PerfNClients {
			perfNClient := &perfCase.PerfNClients[nClientI]

			PerfRoundsAvg := &perfNClient.PerfRoundsAvg
			PerfRoundsAvg.NClient = nClients[nClientI]
			perfRequestsAvg := &PerfRoundsAvg.PerfRequestsAvg

			perfRequestsAvg.Min = math.MaxInt64
			perfRequestsAvg.FEMin = math.MaxInt64

			for roundI := range perfNClient.PerfRounds {
				perfRound := &perfNClient.PerfRounds[roundI]

				PerfRoundsAvg.ErrorCount += float64(perfRound.ErrorCount) / NRound
				PerfRoundsAvg.Throughput += perfRound.Throughput / NRound
				PerfRoundsAvg.PerfRequestsAvg.Avg += perfRound.PerfRequestsAvg.Avg / NRound
				PerfRoundsAvg.PerfRequestsAvg.FEAvg += perfRound.PerfRequestsAvg.FEAvg / NRound

				if perfRound.PerfRequestsAvg.Max > perfRequestsAvg.Max {
					perfRequestsAvg.Max = perfRound.PerfRequestsAvg.Max
				}

				if perfRound.PerfRequestsAvg.Min < perfRequestsAvg.Min {
					perfRequestsAvg.Min = perfRound.PerfRequestsAvg.Min
				}

				if perfRound.PerfRequestsAvg.FEMax > perfRequestsAvg.FEMax {
					perfRequestsAvg.FEMax = perfRound.PerfRequestsAvg.FEMax
				}

				if perfRound.PerfRequestsAvg.FEMin < perfRequestsAvg.FEMin {
					perfRequestsAvg.FEMin = perfRound.PerfRequestsAvg.FEMin
				}
			}
		}
	}

	jsonBytes, _ := json.Marshal(p)
	t.Log(string(jsonBytes))

	for caseI, perfCase := range p.PerfCases {
		latencies := ""
		throughputs := ""
		feLatencies := ""

		for nClientsI, perfNClients := range perfCase.PerfNClients {
			perfRoundsAvg := &perfNClients.PerfRoundsAvg
			if nClientsI == 0 {
				latencies += fmt.Sprintf("%d", int(perfRoundsAvg.PerfRequestsAvg.Avg/1e6))
				throughputs += fmt.Sprintf("%d", int(perfRoundsAvg.Throughput))
				feLatencies += fmt.Sprintf("%d", int(perfRoundsAvg.PerfRequestsAvg.FEAvg/1e6))
			} else {
				latencies += fmt.Sprintf(",%d", int(perfRoundsAvg.PerfRequestsAvg.Avg/1e6))
				throughputs += fmt.Sprintf(",%d", int(perfRoundsAvg.Throughput))
				feLatencies += fmt.Sprintf(",%d", int(perfRoundsAvg.PerfRequestsAvg.FEAvg/1e6))
			}
		}

		fmt.Printf("case %d\n", caseI)
		fmt.Printf("x=[%s]\n", throughputs)
		fmt.Printf("y=[%s]\n", latencies)
		fmt.Printf("z=[%s]\n", feLatencies)
	}
}

func TestHome(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: time.Now().UnixNano()},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         addr,
				MessageName: "home",
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(rsp.Body))
	//t.Log(len(rsp.Trace.Records))
	//t.Log(rsp.Trace)
	bytes, _ := json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}

func TestHomeCrashCurrency0(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace: &tracer.Trace{
			Id: 2,
			Tfis: []*tracer.TFI{
				{
					Type:  tracer.FaultType_FaultCrash,
					Name:  []string{"CurrencyConversionRequest"},
					Delay: 0,
					After: []*tracer.TFIMeta{{Name: "CurrencyConversionRequest", Times: 0}},
				},
			},
		},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeHTML,
					Action:      data.PrintResponse,
				},
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	bytes, _ := json.Marshal(rsp)
	t.Log(string(bytes))
	bytes, _ = json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}

func TestHomeCrashCurrency1(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace: &tracer.Trace{
			Id: 3,
			Tfis: []*tracer.TFI{
				{
					Type:  tracer.FaultType_FaultCrash,
					Name:  []string{"CurrencyConversionRequest"},
					Delay: 0,
					After: []*tracer.TFIMeta{{Name: "CurrencyConversionRequest", Times: 1}},
				},
			},
		},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeHTML,
					Action:      data.PrintResponse,
				},
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	bytes, _ := json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}

func TestHome_RLFI_Java_AdRequest(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace: &tracer.Trace{
			Id: 1,
			Rlfis: []*tracer.RLFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: "AdRequest",
				},
			},
		},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         addr,
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeHTML,
					Action:      data.PrintResponse,
				},
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	//t.Log(string(rsp.Body))
	//t.Log(len(rsp.Trace.Records))
	//t.Log(rsp.Trace)
	bytes, _ := json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}

func TestHome_TFI_Java_AdRequest(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace: &tracer.Trace{
			Id: 1,
			Tfis: []*tracer.TFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"AdRequest"},
					After: []*tracer.TFIMeta{
						{Name: "GetSupportedCurrenciesRequest", Times: 1},
					},
				},
			},
		},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeHTML,
					Action:      data.PrintResponse,
				},
			},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	//t.Log(string(rsp.Body))
	//t.Log(len(rsp.Trace.Records))
	//t.Log(rsp.Trace)
	bytes, _ := json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}

func TestHipsterShop(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: 2},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost/product/OLJCESPC7Z",
				MessageName: "product",
			},
			{
				Method: data.HTTPPost,
				URL:    "http://localhost/cart",
				UrlValues: url.Values{
					"product_id": {"OLJCESPC7Z"},
					"quantity":   {"1"},
				},
				MessageName: "cart",
				Expect:      &data.ExpectedResponse{ContentType: rHtml.ContentTypeHTML},
			},
			{
				Method: data.HTTPPost,
				URL:    "http://localhost/cart",
				UrlValues: url.Values{
					"product_id": {"L9ECAV7KIM"},
					"quantity":   {"1"},
				},
				MessageName: "cart",
			},
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost/product/L9ECAV7KIM",
				MessageName: "product",
			},
			//{
			//	Method: data.HTTPPost,
			//	URL:    "http://localhost/cart/checkout",
			//	UrlValues: url.Values{
			//		"email":                        {"someone@example.com"},
			//		"street_address":               {"1600 Amphitheatre Parkway"},
			//		"zip_code":                     {"94043"},
			//		"city":                         {"Mountain View"},
			//		"state":                        {"CA"},
			//		"country":                      {"United States"},
			//		"credit_card_number":           {"4432-8015-6152-0454"},
			//		"credit_card_expiration_month": {"1"},
			//		"credit_card_expiration_year":  {"2021"},
			//		"credit_card_cvv":              {"672"},
			//	},
			//	MessageName: "checkout",
			//},
		},
	}

	rsp, err := client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	reqs = &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: 3},
		Requests: []data.Request{
			{
				Method: data.HTTPPost,
				URL:    "http://localhost/cart/checkout",
				UrlValues: url.Values{
					"email":                        {"someone@example.com"},
					"street_address":               {"1600 Amphitheatre Parkway"},
					"zip_code":                     {"94043"},
					"city":                         {"Mountain View"},
					"state":                        {"CA"},
					"country":                      {"United States"},
					"credit_card_number":           {"4432-8015-6152-0454"},
					"credit_card_expiration_month": {"1"},
					"credit_card_expiration_year":  {"2021"},
					"credit_card_cvv":              {"672"},
				},
				MessageName: "checkout",
			},
		},
	}

	rsp, err = client.SendRequests(reqs)
	if err != nil {
		t.Error(err)
	}

	t.Log(len(rsp.Trace.Records))
	t.Log(rsp.Trace.JSONString())
	bytes, _ := json.Marshal(rsp.Trace)
	t.Log(string(bytes))
}
