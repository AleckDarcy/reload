package core

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/tracer"

	"github.com/AleckDarcy/reload/core/client/data"
)

const NTests = 1000

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
						ContentType: html.ContentTypeHTML,
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

func TestConcurrency(t *testing.T) {
	nClients := []int{1, 2, 4, 8}

	traces := []*tracer.Trace{
		nil,
		{},
		{

			Rlfis: []*tracer.RLFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: "GetSupportedCurrenciesRequest",
				},
			},
		},
		{

			Rlfis: []*tracer.RLFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: "CurrencyConversionRequest",
				},
			},
		},
		{
			Tfis: []*tracer.TFI{
				{
					Type: tracer.FaultType_FaultCrash,
					Name: []string{"CurrencyConversionRequest"},
					After: []*tracer.TFIMeta{
						{Name: "CurrencyConversionRequest", Times: 2},
					},
				},
			},
		},
	}

	for j, trace := range traces {
		print(fmt.Sprintf("| %d |", j))
		for _, nCLient := range nClients {
			nTest := NTests / nCLient
			clients := make([]*Client, nCLient)
			signals := make(chan struct{}, nCLient)
			reqss := make([]*data.Requests, nCLient)

			for i := 0; i < nCLient; i++ {
				clients[i] = NewClient()

				reqss[i] = &data.Requests{
					CookieUrl: "localhost",
					Trace:     trace,
					Requests: []data.Request{
						{
							Method:      data.HTTPGet,
							URL:         "http://localhost",
							MessageName: "home",
							Expect: &data.ExpectedResponse{
								ContentType: html.ContentTypeHTML,
								//Action:      data.PrintResponse,
							},
						},
					},
				}

				if reqss[i].Trace != nil {
					reqss[i].Trace.Id = int64(i*NTests + 1)
				}
			}

			t.Log(nCLient, nTest)
			start := time.Now()

			for i := 0; i < nCLient; i++ {
				go func(i int, signals chan struct{}) {
					client := clients[i]
					reqs := reqss[i]

					for k := 0; k < nTest; k++ {
						_, err := client.SendRequests(reqs)
						if err != nil {
							t.Error(err)
						}

						if reqs.Trace != nil {
							reqs.Trace.Id++
						}
					}

					signals <- struct{}{}
				}(i, signals)
			}

			for i := 0; i < nCLient; i++ {
				<-signals
			}

			end := time.Now()

			print(fmt.Sprintf(" %v |", end.Sub(start).Seconds()))
		}
		print("\n")
	}
}

func TestHome(t *testing.T) {
	client := NewClient()

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: 1},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
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
					ContentType: html.ContentTypeHTML,
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
					ContentType: html.ContentTypeHTML,
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
				URL:         "http://localhost",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: html.ContentTypeHTML,
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
					ContentType: html.ContentTypeHTML,
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
				Expect:      &data.ExpectedResponse{ContentType: html.ContentTypeHTML},
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
