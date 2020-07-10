package core

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"github.com/AleckDarcy/reload/core/client/data"
)

const (
	Idle int64 = iota
	Running
)

type Status struct {
	Status int64
	TaskID int64

	ID      int
	CaseID  int
	NClient int
	Round   int
}

type CaseConf struct {
	Request *data.Request
}

type Perf struct {
	Cases []Case
}

type Case struct {
	NClients []NClient
}

type NClient struct {
	RoundsAvg RoundsAvg
	Rounds    []Round `json:"-"`
}

type Round struct {
	RequestsAvg RequestsAvg
	Requests    []Request

	ErrorCount int64
	Throughput float64
}

type RoundsAvg struct {
	NClient     int
	RequestsAvg RequestsAvg

	ErrorCount    float64
	ThroughputAvg ThroughputAvg
}

type Request struct {
	E2ELatency int64
	FELatency  int64
}

type RequestsAvg struct {
	E2ELatencyAvg LatencyAvg
	FELatencyAvg  LatencyAvg
}

type LatencyAvg struct {
	Max, Min int64
	Mean     float64
	StdDev   float64
}

type ThroughputAvg struct {
	Max, Min float64
	Mean     float64
	StdDev   float64
}

func RunPerf(nTests int64, nRound int64, nClients []int, caseConfs []CaseConf, status *Status) *Perf {
	fTests, fRound := float64(nTests), float64(nRound)

	p := &Perf{
		Cases: make([]Case, len(caseConfs)),
	}

	for caseI, case_ := range caseConfs {
		perfCase := &p.Cases[caseI]
		perfCase.NClients = make([]NClient, len(nClients))

		for nClientI, nClient := range nClients {
			status.CaseID, status.NClient = caseI, nClient
			//fmt.Printf("case %d, nClient %d\n", caseI, nClient)

			perfNClient := &perfCase.NClients[nClientI]
			perfNClient.Rounds = make([]Round, nRound)

			clients := make([]*Client, nClient)
			signals := make(chan struct{}, nClient)
			requests := make([]*data.Requests, nClient)

			for i := 0; i < nClient; i++ {
				clients[i] = NewClient()

				request := &data.Requests{
					CookieUrl: "localhost",
					Requests: []data.Request{
						*case_.Request,
					},
				}

				if case_.Request.Trace != nil {
					trace := *case_.Request.Trace
					request.Trace = &trace
				}

				requests[i] = request
			}

			for roundI := int64(0); roundI < nRound; roundI++ {
				status.Round = int(roundI)
				perfRound := &perfNClient.Rounds[roundI]
				perfRound.Requests = make([]Request, nTests)

				traceIDOffset := time.Now().UnixNano()
				traceID := int64(0)

				start := time.Now()

				for i := 0; i < nClient; i++ {
					go func(i int, signals chan struct{}) {
						client := clients[i]
						reqs := requests[i]

						for {
							traceID := atomic.AddInt64(&traceID, 1)
							if traceID > nTests {
								break
							}

							if reqs.Trace != nil {
								reqs.Trace.Id = traceID + traceIDOffset
							}

							perfRequest := &perfRound.Requests[traceID-1]

							rsp, err := client.SendRequests(reqs)
							perfRequest.E2ELatency = rsp.Latency

							if err != nil {
								fmt.Printf("Send Requests err: %v\n", err)
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

				for i := 0; i < nClient; i++ {
					<-signals
				}

				end := time.Now()

				perfRequestsAvg := &perfRound.RequestsAvg
				perfRequestsAvg.E2ELatencyAvg.Min = math.MaxInt64
				perfRequestsAvg.FELatencyAvg.Min = math.MaxInt64

				e2eLatencies := make([]float64, len(perfRound.Requests))
				feLatencies := make([]float64, len(perfRound.Requests))

				for requestI, perfRequest := range perfRound.Requests {
					e2eLatencies[requestI] = float64(perfRequest.E2ELatency)
					feLatencies[requestI] = float64(perfRequest.FELatency)

					if perfRequest.E2ELatency > perfRequestsAvg.E2ELatencyAvg.Max {
						perfRequestsAvg.E2ELatencyAvg.Max = perfRequest.E2ELatency
					}

					if perfRequest.E2ELatency < perfRequestsAvg.E2ELatencyAvg.Min {
						perfRequestsAvg.E2ELatencyAvg.Min = perfRequest.E2ELatency
					}

					if perfRequest.FELatency > perfRequestsAvg.FELatencyAvg.Max {
						perfRequestsAvg.FELatencyAvg.Max = perfRequest.FELatency
					}

					if perfRequest.FELatency < perfRequestsAvg.FELatencyAvg.Min {
						perfRequestsAvg.FELatencyAvg.Min = perfRequest.FELatency
					}
				}

				perfRequestsAvg.E2ELatencyAvg.Mean, perfRequestsAvg.E2ELatencyAvg.StdDev = MeanAndStdDev(e2eLatencies)
				perfRequestsAvg.FELatencyAvg.Mean, perfRequestsAvg.FELatencyAvg.StdDev = MeanAndStdDev(feLatencies)

				perfRound.Throughput = fTests * 1e9 / float64(end.Sub(start).Nanoseconds())
			}
		}
	}

	for caseI := range p.Cases {
		perfCase := &p.Cases[caseI]
		for nClientI := range perfCase.NClients {
			perfNClient := &perfCase.NClients[nClientI]

			perfRoundsAvg := &perfNClient.RoundsAvg
			perfRoundsAvg.NClient = nClients[nClientI]
			perfRequestsAvg := &perfRoundsAvg.RequestsAvg
			perfRequestsAvg.E2ELatencyAvg.Min = math.MaxInt64
			perfRequestsAvg.FELatencyAvg.Min = math.MaxInt64

			throughputAvg := &perfRoundsAvg.ThroughputAvg
			throughputAvg.Min = math.MaxFloat64

			e2eLatencies := make([]float64, len(perfNClient.Rounds))
			feLatencies := make([]float64, len(perfNClient.Rounds))

			roundThroughputs := make([]float64, len(perfNClient.Rounds))
			for roundI := range perfNClient.Rounds {
				perfRound := &perfNClient.Rounds[roundI]

				perfRoundsAvg.ErrorCount += float64(perfRound.ErrorCount) / fRound
				e2eLatencies[roundI] = perfRound.RequestsAvg.E2ELatencyAvg.Mean
				feLatencies[roundI] = perfRound.RequestsAvg.FELatencyAvg.Mean

				if perfRound.RequestsAvg.E2ELatencyAvg.Max > perfRequestsAvg.E2ELatencyAvg.Max {
					perfRequestsAvg.E2ELatencyAvg.Max = perfRound.RequestsAvg.E2ELatencyAvg.Max
				}

				if perfRound.RequestsAvg.E2ELatencyAvg.Min < perfRequestsAvg.E2ELatencyAvg.Min {
					perfRequestsAvg.E2ELatencyAvg.Min = perfRound.RequestsAvg.E2ELatencyAvg.Min
				}

				if perfRound.RequestsAvg.FELatencyAvg.Max > perfRequestsAvg.FELatencyAvg.Max {
					perfRequestsAvg.FELatencyAvg.Max = perfRound.RequestsAvg.FELatencyAvg.Max
				}

				if perfRound.RequestsAvg.FELatencyAvg.Min < perfRequestsAvg.FELatencyAvg.Min {
					perfRequestsAvg.FELatencyAvg.Min = perfRound.RequestsAvg.FELatencyAvg.Min
				}

				roundThroughputs[roundI] = perfRound.Throughput

				if perfRound.Throughput > perfRoundsAvg.ThroughputAvg.Max {
					perfRoundsAvg.ThroughputAvg.Max = perfRound.Throughput
				}

				if perfRound.Throughput < perfRoundsAvg.ThroughputAvg.Min {
					perfRoundsAvg.ThroughputAvg.Min = perfRound.Throughput
				}
			}

			perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean, perfRoundsAvg.RequestsAvg.E2ELatencyAvg.StdDev = MeanAndStdDev(e2eLatencies)
			perfRoundsAvg.RequestsAvg.FELatencyAvg.Mean, perfRoundsAvg.RequestsAvg.FELatencyAvg.StdDev = MeanAndStdDev(feLatencies)

			throughputAvg.Mean, throughputAvg.StdDev = MeanAndStdDev(roundThroughputs)
		}
	}

	return p
}

func MeanAndStdDev(nums []float64) (mean float64, sd float64) {
	length := float64(len(nums))

	for _, num := range nums {
		mean += num / length
	}

	for _, num := range nums {
		sd += math.Pow(num-mean, 2)
	}

	sd = math.Sqrt(sd / 10)

	return
}
