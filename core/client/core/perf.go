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
	StdErr   float64
}

type ThroughputAvg struct {
	Max, Min float64
	Mean     float64
	StdDev   float64
	StdErr   float64
}

func RunPerf(nTests int64, nRound int64, nClients []int, caseConfs []CaseConf, status *Status, rspFunc func(rsp *data.Response)) *Perf {
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

								if case_.Request.Expect.Action&data.ServiceLatency != 0 {
									if rspFunc != nil {
										go rspFunc(rsp)
									}
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

				e2eLatencyAvg := &perfRound.RequestsAvg.E2ELatencyAvg
				feLatencyAvg := &perfRound.RequestsAvg.FELatencyAvg

				e2eLatencyAvg.Min = math.MaxInt64
				feLatencyAvg.Min = math.MaxInt64

				e2eLatencies := make([]float64, len(perfRound.Requests))
				feLatencies := make([]float64, len(perfRound.Requests))

				for requestI, perfRequest := range perfRound.Requests {
					e2eLatencies[requestI] = float64(perfRequest.E2ELatency)
					feLatencies[requestI] = float64(perfRequest.FELatency)

					if perfRequest.E2ELatency > e2eLatencyAvg.Max {
						e2eLatencyAvg.Max = perfRequest.E2ELatency
					}

					if perfRequest.E2ELatency < e2eLatencyAvg.Min {
						e2eLatencyAvg.Min = perfRequest.E2ELatency
					}

					if perfRequest.FELatency > feLatencyAvg.Max {
						feLatencyAvg.Max = perfRequest.FELatency
					}

					if perfRequest.FELatency < feLatencyAvg.Min {
						feLatencyAvg.Min = perfRequest.FELatency
					}
				}

				e2eLatencyAvg.Mean, e2eLatencyAvg.StdDev, e2eLatencyAvg.StdErr = MeanAndStdDevAndStdErr(e2eLatencies)
				feLatencyAvg.Mean, feLatencyAvg.StdDev, feLatencyAvg.StdErr = MeanAndStdDevAndStdErr(feLatencies)
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

			e2eLatencyAvg := &perfRoundsAvg.RequestsAvg.E2ELatencyAvg
			feLatencyAvg := &perfRoundsAvg.RequestsAvg.FELatencyAvg
			throughputAvg := &perfRoundsAvg.ThroughputAvg

			e2eLatencyAvg.Min = math.MaxInt64
			feLatencyAvg.Min = math.MaxInt64
			throughputAvg.Min = math.MaxFloat64

			e2eLatencies := make([]float64, len(perfNClient.Rounds))
			feLatencies := make([]float64, len(perfNClient.Rounds))
			throughputs := make([]float64, len(perfNClient.Rounds))

			for roundI := range perfNClient.Rounds {
				perfRound := &perfNClient.Rounds[roundI]

				perfRoundsAvg.ErrorCount += float64(perfRound.ErrorCount) / fRound

				e2eLatencies[roundI] = perfRound.RequestsAvg.E2ELatencyAvg.Mean
				feLatencies[roundI] = perfRound.RequestsAvg.FELatencyAvg.Mean
				throughputs[roundI] = perfRound.Throughput

				if perfRound.RequestsAvg.E2ELatencyAvg.Max > e2eLatencyAvg.Max {
					e2eLatencyAvg.Max = perfRound.RequestsAvg.E2ELatencyAvg.Max
				}

				if perfRound.RequestsAvg.E2ELatencyAvg.Min < e2eLatencyAvg.Min {
					e2eLatencyAvg.Min = perfRound.RequestsAvg.E2ELatencyAvg.Min
				}

				if perfRound.RequestsAvg.FELatencyAvg.Max > feLatencyAvg.Max {
					feLatencyAvg.Max = perfRound.RequestsAvg.FELatencyAvg.Max
				}

				if perfRound.RequestsAvg.FELatencyAvg.Min < feLatencyAvg.Min {
					feLatencyAvg.Min = perfRound.RequestsAvg.FELatencyAvg.Min
				}

				if perfRound.Throughput > perfRoundsAvg.ThroughputAvg.Max {
					perfRoundsAvg.ThroughputAvg.Max = perfRound.Throughput
				}

				if perfRound.Throughput < perfRoundsAvg.ThroughputAvg.Min {
					perfRoundsAvg.ThroughputAvg.Min = perfRound.Throughput
				}
			}

			e2eLatencyAvg.Mean, e2eLatencyAvg.StdDev, e2eLatencyAvg.StdErr = MeanAndStdDevAndStdErr(e2eLatencies)
			feLatencyAvg.Mean, feLatencyAvg.StdDev, feLatencyAvg.StdErr = MeanAndStdDevAndStdErr(feLatencies)
			throughputAvg.Mean, throughputAvg.StdDev, throughputAvg.StdErr = MeanAndStdDevAndStdErr(throughputs)
		}
	}

	return p
}

func MeanAndStdDevAndStdErr(nums []float64) (mean, stdDev, stdErr float64) {
	length := float64(len(nums))

	for _, num := range nums {
		mean += num / length
	}

	for _, num := range nums {
		stdDev += math.Pow(num-mean, 2)
	}

	stdDev = math.Sqrt(stdDev / 10)
	stdErr = stdDev / math.Sqrt(10)

	return
}

type Report struct {
	Cases []ReportCase
}

type ReportCase struct {
	E2ELatenciesMean string
	ThroughputsMean  string
	FELatenciesMean  string

	E2ELatenciesErrorBar string
	ThroughputsErrorBar  string
	FELatenciesErrorBar  string
}

func GetReport(perf *Perf) *Report {
	r := &Report{Cases: make([]ReportCase, len(perf.Cases))}

	for caseI, perfCase := range perf.Cases {
		e2eLatenciesMean := ""
		throughputsMean := ""
		feLatenciesMean := ""

		e2eLatenciesErrorBar := ""
		throughputsErrorBar := ""
		feLatenciesErrorBar := ""

		for nClientsI, perfNClients := range perfCase.NClients {
			perfRoundsAvg := &perfNClients.RoundsAvg
			if nClientsI == 0 {
				e2eLatenciesMean += fmt.Sprintf("%d", int(perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/1e6))
				throughputsMean += fmt.Sprintf("%d", int(perfRoundsAvg.ThroughputAvg.Mean))
				feLatenciesMean += fmt.Sprintf("%d", int(perfRoundsAvg.RequestsAvg.FELatencyAvg.Mean/1e6))

				e2eLatenciesErrorBar += fmt.Sprintf("%f", perfRoundsAvg.RequestsAvg.E2ELatencyAvg.StdErr/1e6)
				throughputsErrorBar += fmt.Sprintf("%f", perfRoundsAvg.ThroughputAvg.StdErr)
				feLatenciesErrorBar += fmt.Sprintf("%f", perfRoundsAvg.RequestsAvg.FELatencyAvg.StdErr/1e6)
			} else {
				e2eLatenciesMean += fmt.Sprintf(",%d", int(perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/1e6))
				throughputsMean += fmt.Sprintf(",%d", int(perfRoundsAvg.ThroughputAvg.Mean))
				feLatenciesMean += fmt.Sprintf(",%d", int(perfRoundsAvg.RequestsAvg.FELatencyAvg.Mean/1e6))

				e2eLatenciesErrorBar += fmt.Sprintf(",%f", perfRoundsAvg.RequestsAvg.E2ELatencyAvg.StdErr/1e6)
				throughputsErrorBar += fmt.Sprintf(",%f", perfRoundsAvg.ThroughputAvg.StdErr)
				feLatenciesErrorBar += fmt.Sprintf(",%f", perfRoundsAvg.RequestsAvg.FELatencyAvg.StdErr/1e6)
			}
		}

		r.Cases[caseI] = ReportCase{
			E2ELatenciesMean: e2eLatenciesMean,
			ThroughputsMean:  throughputsMean,
			FELatenciesMean:  feLatenciesMean,

			E2ELatenciesErrorBar: e2eLatenciesErrorBar,
			ThroughputsErrorBar:  throughputsErrorBar,
			FELatenciesErrorBar:  feLatenciesErrorBar,
		}
	}

	return r
}

type Overhead struct {
	Cases []OverheadCase
}

type OverheadCase struct {
	E2ELatencies string
	Throughputs  string
}

func GetOverhead(base *Perf, perf *Perf) *Overhead {
	o := &Overhead{Cases: make([]OverheadCase, len(base.Cases))}

	for caseI, baseCase := range base.Cases {
		perfCase := perf.Cases[caseI]

		e2eLatencies := ""
		throughputs := ""

		for nClientsI, baseNClients := range baseCase.NClients {
			perfNClients := perfCase.NClients[nClientsI]

			baseRoundsAvg := &baseNClients.RoundsAvg
			perfRoundsAvg := &perfNClients.RoundsAvg
			if nClientsI == 0 {
				e2eLatencies += fmt.Sprintf("%f", perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/baseRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean*100-100)
				throughputs += fmt.Sprintf("%f", perfRoundsAvg.ThroughputAvg.Mean/baseRoundsAvg.ThroughputAvg.Mean*100-100)
			} else {
				e2eLatencies += fmt.Sprintf(",%f", perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/baseRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean*100-100)
				throughputs += fmt.Sprintf(",%f", perfRoundsAvg.ThroughputAvg.Mean/baseRoundsAvg.ThroughputAvg.Mean*100-100)
			}
		}

		o.Cases[caseI] = OverheadCase{
			E2ELatencies: e2eLatencies,
			Throughputs:  throughputs,
		}
	}

	return o
}

type Table struct {
	Cases []TableCase
}

type TableCase struct {
	Throughput string
	Latency    string
}

func GetTable(base, _3MileBeach, jaeger *Perf) *Table {
	t := &Table{Cases: make([]TableCase, len(base.Cases))}

	for caseI, baseCase := range base.Cases {
		_3MileBeachCase := _3MileBeach.Cases[caseI]
		jaegerCase := jaeger.Cases[caseI]

		throughput := ""
		latency := ""

		for nClientsI, baseNClients := range baseCase.NClients {
			_3MileBeachNClients := _3MileBeachCase.NClients[nClientsI]
			jaegerNClients := jaegerCase.NClients[nClientsI]

			baseRoundsAvg := &baseNClients.RoundsAvg
			_3MileBeachRoundsAvg := &_3MileBeachNClients.RoundsAvg
			jaegerRoundsAvg := &jaegerNClients.RoundsAvg

			throughput += fmt.Sprintf(""+
				"%d & %d & %d(%0.2fx) & %d(%0.2fx) \\\\\n",
				baseRoundsAvg.NClient, int(baseRoundsAvg.ThroughputAvg.Mean),
				int(_3MileBeachRoundsAvg.ThroughputAvg.Mean), _3MileBeachRoundsAvg.ThroughputAvg.Mean/baseRoundsAvg.ThroughputAvg.Mean,
				int(jaegerRoundsAvg.ThroughputAvg.Mean), jaegerRoundsAvg.ThroughputAvg.Mean/baseRoundsAvg.ThroughputAvg.Mean,
			)
			latency += fmt.Sprintf(""+
				"%d & %d & %d(%0.2fx) & %d(%0.2fx) \\\\\n",
				baseRoundsAvg.NClient, int(baseRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean)/1e6,
				int(_3MileBeachRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean)/1e6, _3MileBeachRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/baseRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean,
				int(jaegerRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean)/1e6, jaegerRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/baseRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean,
			)
		}

		t.Cases[caseI].Throughput = throughput
		t.Cases[caseI].Latency = latency
	}

	return t
}

func GetProcessLatency(perf *Perf) string {
	perfCase := perf.Cases[0]

	result := ""

	for _, perfNClients := range perfCase.NClients {

		perfRoundsAvg := &perfNClients.RoundsAvg

		result += fmt.Sprintf(""+
			"%d & %d & %d(%0.2fx) & %d \\\\\n",
			perfRoundsAvg.NClient, int(perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/1e6),
			int(perfRoundsAvg.RequestsAvg.FELatencyAvg.Mean/1e6),
			perfRoundsAvg.RequestsAvg.FELatencyAvg.Mean/perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean,
			int(perfRoundsAvg.RequestsAvg.E2ELatencyAvg.Mean/1e6)-int(perfRoundsAvg.RequestsAvg.FELatencyAvg.Mean/1e6),
		)
	}

	return result
}
