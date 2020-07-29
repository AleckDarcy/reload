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
	//currencyConversionRequest,
	//currencyConversionRequest,
	//currencyConversionRequest,
	//currencyConversionRequest,
	//currencyConversionRequest,
	//currencyConversionRequest,
	//currencyConversionRequest,
	//currencyConversionRequest,
	{
		Service: "adservice",
		Name:    "AdRequest",
	},
	{"1", "1"},
	{"2", "2"},
	{"3", "3"},
	{"4", "4"},
	//currencyConversionRequest,
}

func main() {
	var names = homeNames

	resultChan := make(chan *tracer.Faults, 1)
	expect := int64(math.Pow(2, float64(len(names))) - 1)
	fmt.Println("expect:", expect)
	c := tracer.NewGenerator(names)
	go func() {
		//resultChan <- c.GetFaults([]int{0, 3, 12})
		//expect = 1

		for i := 1; i <= len(names); i++ {
			c.SetLen(i)
			c.GenerateCombinartorial(resultChan)
		}
	}()

	count := int64(0)
	traceID := time.Now().UnixNano()

	countSignal := make(chan struct{}, 1)

	traceIDOffset := time.Now().UnixNano()

	//expect = 10
	errorCount := int64(0)
	okCount := int64(0)
	start := time.Now()

	tmp := []int{0, 3, 10, 11, 12}
	_ = tmp

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
						//fmt.Print(".")
						//if reflect.DeepEqual(result.IDs, []int{0, 3, 12}) {
						//	reqs.Trace.Id = 110
						//}

						//fmt.Println(result.Faults)
						rsp, err := client.SendRequests(reqs)
						_ = rsp

						if err != nil {
							atomic.AddInt64(&errorCount, 1)

							fmt.Println(result.IDs)
							//if reflect.DeepEqual(result.IDs, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}) && rsp != nil {
							//	t.Log(string(rsp.Body))
							//}
							fmt.Println(err)
							//errCount++
						} else {
							atomic.AddInt64(&okCount, 1)
							//fmt.Println("hahahah")
						}

						//if reflect.DeepEqual(result.IDs, tmp) {
						//	if err != nil {
						//		fmt.Println("guale")
						//	}
						//	if rsp != nil {
						//		fmt.Println(result)
						//		fmt.Println(string(rsp.Body))
						//		//	os.Exit(0)
						//	}
						//}
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
