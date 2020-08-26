package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"sync/atomic"
	"time"

	"github.com/AleckDarcy/reload/core/log"

	"github.com/AleckDarcy/reload/core/client/kubectl"

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
	start := int64(0)
	end := int64(0)

	flag.Int64Var(&start, "start", 0, "start id")
	flag.Int64Var(&start, "end", 0, "end id")
	flag.Parse()

	chaos(start, end)
}

func chaos(start, end int64) {
	names := []string{
		"adservice",
		"cartservice",
		"currencyservice",
		"productcatalogservice",
		"checkoutservice",
		"emailservice",
		"paymentservice",
		"recommendationservice",
		"shippingservice",
	}
	g := kubectl.NewGenerator(names)

	file, _ := os.OpenFile("result.jsons", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
	fw := bufio.NewWriter(file)

	resultChan := make(chan []int, 1)

	go func() {
		for i := 1; i <= len(names); i++ {
			g.SetLen(i)
			g.GenerateCombinartorial(resultChan)
		}
	}()

	expect := int64(math.Pow(2, float64(len(names))) - 1)
	if end <= start {
		end = expect
	}

	count := int64(0)

	services := make([]*kubectl.Service, len(names))
	for i, name := range names {
		services[i] = &kubectl.Service{
			Name:      name,
			NameSpace: "default",
			Replicas:  1,
		}
	}

	client := core.NewClient()
	reqs := &data.Requests{
		CookieUrl: "localhost",
		Trace:     &tracer.Trace{Id: time.Now().UnixNano()},
		Requests: []data.Request{
			{
				Method:      data.HTTPGet,
				URL:         "http://localhost",
				MessageName: "home",
				Expect: &data.ExpectedResponse{
					ContentType: rHtml.ContentTypeHTML,
					Action:      data.DeserializeTrace,
				},
			},
		},
	}

	var res *core.ChaosResult
	svc := make([]*kubectl.Service, 0, len(names))

	log.Logf("start task, range: [%d, %d]", start, end)
	for {
		select {
		case result := <-resultChan:
			if count < start {
				break
			}

			svc = svc[:len(result)]
			for i, id := range result {
				svc[i] = services[id]
			}

			res = core.Chaos(count, client, reqs, svc, result)

			jsonBytes, _ := json.Marshal(res)
			jsonStr := string(jsonBytes)
			fw.WriteString(jsonStr)
			fw.WriteByte('\n')
			fw.Flush()
			log.Logf("%v\n", jsonStr)
		}

		if count++; count == expect || count == end {
			break
		} else if count%100 == 0 {
			fmt.Println("=============", count, "=============")
		}
	}
}

func chaosStyle() {
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
