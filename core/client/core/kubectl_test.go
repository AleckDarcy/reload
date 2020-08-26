package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/core/log"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/client/kubectl"
	"github.com/AleckDarcy/reload/core/tracer"
	rHtml "github.com/AleckDarcy/reload/runtime/html"
)

func TestReadJSON(t *testing.T) {
	file, _ := os.OpenFile("main/result.jsons", os.O_RDONLY, 0666)
	fr := bufio.NewReader(file)

	total := time.Duration(0)
	scaleInBlastCount := 0
	scaleOutBlastCount := 0
	scaleInBlastTime := time.Duration(0)
	scaleOutBlastTime := time.Duration(0)
	faults := 0

	scaleInTime := time.Duration(0)
	scaleOutTime := time.Duration(0)
	requestTime := time.Duration(0)

	for {
		line, _, err := fr.ReadLine()
		if err != nil {
			//t.Log(err)

			break
		}

		//t.Log(string(line))

		result := &ChaosResultJSON{}
		json.Unmarshal(line, result)

		total += result.TotalTime
		scaleInTime += result.ScaleIn.Time
		scaleOutTime += result.ScaleOut.Time
		requestTime += result.TestTime

		if result.SendRequestError != nil {
			faults++
		}

		if result.ScaleIn.Retry != 0 {
			//fmt.Println(string(line))
			scaleInBlastCount++
			scaleInBlastTime += result.ScaleIn.WaitReadyTime
		}

		if result.ScaleOut.Retry != 0 {
			scaleOutBlastCount++
			scaleOutBlastTime += result.ScaleOut.WaitReadyTime
		}
	}

	log.Logf("total time: %v", total)
	log.Logf("faults: %d", faults)
	log.Logf("scale in blast radius count: %d, time: %v", scaleInBlastCount, scaleInBlastTime)
	log.Logf("scale out blast radius count: %d, time: %v", scaleOutBlastCount, scaleOutBlastTime)
	log.Logf("scale in time: %v", scaleInTime)
	log.Logf("scale out time: %v", scaleOutTime)
	log.Logf("request time: %v", requestTime)
}

func TestChaos(t *testing.T) {
	names := []string{
		"adservice",
		"cartservice",
		"currencyservice",
		"productcatalogservice",

		//"checkoutservice",
		//"emailservice",
		//"paymentservice",
		//"recommendationservice",
		//"shippingservice",
	}
	g := kubectl.NewGenerator(names)

	resultChan := make(chan []int, 1)

	go func() {
		for i := 1; i <= len(names); i++ {
			g.SetLen(i)
			g.GenerateCombinartorial(resultChan)
		}
	}()

	expect := int64(math.Pow(2, float64(len(names))) - 1)
	start := int64(0)
	count := int64(0)

	services := make([]*kubectl.Service, len(names))
	for i, name := range names {
		services[i] = &kubectl.Service{
			Name:      name,
			NameSpace: "default",
			Replicas:  1,
		}
	}

	client := NewClient()
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

	var res *ChaosResult
	svc := make([]*kubectl.Service, 0, len(names))
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

			res = Chaos(count, client, reqs, svc, result)

			jsonBytes, _ := json.Marshal(res)
			fmt.Printf("%v\n", string(jsonBytes))
		}

		if count++; count == expect {
			break
		} else if count%100 == 0 {
			fmt.Println("=============", count, "=============")
		}
	}
}

func TestKubeCtl(t *testing.T) {
	client := NewClient()

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

	services := []*kubectl.Service{
		{Name: "currencyservice", NameSpace: "default", Replicas: 1},
		{Name: "adservice", NameSpace: "default", Replicas: 1},
	}

	Chaos(0, client, reqs, services, nil)
}
