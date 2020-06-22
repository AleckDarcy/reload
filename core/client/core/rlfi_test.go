package core

import (
	"encoding/json"
	"math"
	"sync/atomic"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/core/client/data"
	rHtml "github.com/AleckDarcy/reload/runtime/html"

	"github.com/AleckDarcy/reload/core/tracer"
)

var getSupportedCurrenciesRequest = &tracer.NameMeta{
	Service: "currencyservice",
	Name:    "GetSupportedCurrenciesRequest",
}

var currencyConversionRequest = &tracer.NameMeta{
	Service: "currencyservice",
	Name:    "CurrencyConversionRequest",
}

var service2requests = map[string][]string{
	"cartservice":           {"AddItemRequest", "EmptyCartRequest", "GetCartRequest"},
	"recommendationservice": {"ListRecommendationsRequest"},
	"productcatalogservice": {"ListProductsRequest", "GetProductRequest", "SearchProductsRequest"},
	"shippingservice":       {"GetQuoteRequest", "ShipOrderRequest"},
	"currencyservice":       {"GetSupportedCurrenciesRequest", "CurrencyConversionRequest"},
	"paymentservice":        {"ChargeRequest"},
	"emailservice":          {"EmailService", "SendOrderConfirmationRequest"},
	"checkoutservice":       {"PlaceOrderRequest"},
	"adservice":             {"AdRequest"},
}

func chaosNames() (cNames []*tracer.NameMeta) {
	for service, names := range service2requests {
		for _, name := range names {
			cNames = append(cNames, &tracer.NameMeta{
				Service: service,
				Name:    name,
			})
		}
	}

	return cNames
}

var homeNames = []*tracer.NameMeta{
	getSupportedCurrenciesRequest,
	{
		Service: "productcatalogservice",
		Name:    "ListProductsRequest",
	},
	{
		Service: "cartservice",
		Name:    "GetCartRequest",
	},
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	currencyConversionRequest,
	{
		Service: "adservice",
		Name:    "AdRequest",
	},
	currencyConversionRequest,
}

func TestName(t *testing.T) {
	c := tracer.NewGenerator(homeNames)

	idsList := [][]int{
		{1, 2, 3, 4, 6, 7, 8, 9, 10, 13},

		//{1, 3, 6, 11, 12},
		//{1, 2, 3, 5, 11, 12},
		//{1, 3, 4, 6, 8, 12},
		//{1, 2, 3, 4, 7, 9, 12},
		//{1, 2, 3, 5, 9, 10, 12},
		//{1, 2, 3, 5, 10, 11, 12},
		//{1, 2, 3, 6, 8, 10, 12},
		//{1, 2, 3, 6, 8, 11, 12},
		//{1, 2, 3, 6, 10, 11, 12},
		//{1, 3, 5, 6, 8, 11, 12},
		//{1, 3, 5, 6, 9, 10, 12},
		//{1, 2, 3, 4, 5, 7, 10, 12},
		//{1, 2, 3, 4, 5, 8, 11, 12},
		//{1, 2, 3, 4, 5, 9, 11, 12},
		//{1, 3, 6, 7, 9, 10, 11, 12},
		//{1, 3, 6, 8, 9, 10, 11, 12},
		//{1, 3, 5, 6, 7, 8, 9, 10, 12},
		//{1, 2, 3, 4, 5, 6, 7, 8, 11, 12},
		//{1, 2, 3, 4, 5, 6, 7, 9, 11, 12},
		//{1, 2, 3, 4, 5, 6, 8, 9, 10, 12},
		//{1, 2, 3, 4, 5, 6, 8, 10, 11, 12},
	}

	client := NewClient()
	_ = client

	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeHTML,
					//Action:      data.PrintResponse,
				},
			},
		},
	}

	traceIDOffset := time.Now().UnixNano()
	for i, ids := range idsList {
		result := c.GetFaults(ids)

		reqs.Trace.Id = int64(i) + traceIDOffset
		reqs.Trace.Tfis = result.Faults

		jsonBytes, _ := json.Marshal(result)
		//fmt.Println(string(jsonBytes))
		_ = jsonBytes

		if result.FIType == tracer.FI_TFI {
			//fmt.Println(result.Faults)
			rsp, err := client.SendRequests(reqs)
			_ = rsp
			if err != nil {
				t.Log(result)
				//if reflect.DeepEqual(result.IDs, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}) && rsp != nil {
				//	t.Log(string(rsp.Body))
				//}
				t.Error(err)
				//errCount++
			}
		}
	}
}

func TestChaos(t *testing.T) {
	//var names = chaosNames()
	var names = homeNames

	resultChan := make(chan *tracer.Faults, 1)

	go func() {
		c := tracer.NewGenerator(names)
		for i := 1; i <= len(names); i++ {
			c.SetLen(i)
			c.GenerateCombinartorial(resultChan)
		}
	}()

	expect := int64(math.Pow(2, float64(len(names))) - 1)
	count := int64(0)
	traceID := time.Now().UnixNano()

	countSignal := make(chan struct{}, 1)

	traceIDOffset := time.Now().UnixNano()

	//expect = 10
	errorCount := 0

	for i := 0; i < 4; i++ {
		go func(signal chan struct{}) {
			client := NewClient()
			_ = client

			reqs := &data.Requests{
				CookieUrl: "localhost",
				Trace:     &tracer.Trace{},
				Requests: []data.Request{
					{
						Method:      data.HTTPGet,
						URL:         "http://localhost",
						MessageName: "home",
						Expect: &data.ExpectedResponse{
							ContentType: rHtml.ContentTypeHTML,
							//Action:      data.PrintResponse,
						},
					},
				},
			}

			for {
				select {
				case result := <-resultChan:
					traceID := atomic.AddInt64(&traceID, 1)

					reqs.Trace.Id = traceID + traceIDOffset
					reqs.Trace.Tfis = result.Faults

					jsonBytes, _ := json.Marshal(result)
					//fmt.Println(string(jsonBytes))
					_ = jsonBytes

					if result.FIType == tracer.FI_TFI || result.FIType == tracer.FI_RLFI {
						//fmt.Println(result.Faults)
						rsp, err := client.SendRequests(reqs)
						_ = rsp
						if err != nil {
							errorCount++
							t.Log(result)
							//if reflect.DeepEqual(result.IDs, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}) && rsp != nil {
							//	t.Log(string(rsp.Body))
							//}
							t.Error(err)
							//errCount++
						}
					}

					if atomic.AddInt64(&count, 1) == expect {
						signal <- struct{}{}

						return
					}
				}
			}
		}(countSignal)
	}

	<-countSignal

	t.Log(errorCount)
}
