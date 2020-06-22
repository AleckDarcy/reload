package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"github.com/AleckDarcy/reload/core/client/core"
	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
	rHtml "github.com/AleckDarcy/reload/runtime/html"
)

var getSupportedCurrenciesRequest = &tracer.NameMeta{
	Service: "currencyservice",
	Name:    "GetSupportedCurrenciesRequest",
}

var currencyConversionRequest = &tracer.NameMeta{
	Service: "currencyservice",
	Name:    "CurrencyConversionRequest",
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

func main() {
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

	start := time.Now()

	for i := 0; i < 4; i++ {
		go func(signal chan struct{}) {
			client := core.NewClient()
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

					//jsonBytes, _ := json.Marshal(result)
					//fmt.Println(string(jsonBytes))
					//_ = jsonBytes

					if result.FIType == tracer.FI_TFI || result.FIType == tracer.FI_RLFI {
						//fmt.Println(result.Faults)
						rsp, err := client.SendRequests(reqs)
						_ = rsp
						if err != nil {
							errorCount++

							fmt.Println(result)
							//if reflect.DeepEqual(result.IDs, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}) && rsp != nil {
							//	t.Log(string(rsp.Body))
							//}
							fmt.Println(err)
							//errCount++
						}
					}

					if tmp := atomic.AddInt64(&count, 1); tmp == expect {
						signal <- struct{}{}

						return
					} else if tmp%100 == 0 {
						fmt.Println("=============", tmp, "=============")
					}
				}
			}
		}(countSignal)
	}

	<-countSignal

	end := time.Now()

	fmt.Println(errorCount)
	fmt.Println(end.Sub(start))
}
