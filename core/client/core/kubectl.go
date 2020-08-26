package core

import (
	"time"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/client/kubectl"
	"github.com/AleckDarcy/reload/core/log"
)

type Scale struct {
	Time          time.Duration
	Retry         int
	WaitReadyTime time.Duration
}

type ChaosResult struct {
	ID         int64
	ServiceIDs []int

	TotalTime time.Duration
	ScaleIn   Scale
	TestTime  time.Duration
	ScaleOut  Scale

	SendRequestError error
}

type ChaosResultJSON struct {
	ID         int64
	ServiceIDs []int

	TotalTime time.Duration
	ScaleIn   Scale
	TestTime  time.Duration
	ScaleOut  Scale

	SendRequestError map[string]interface{}
}

func Chaos(id int64, client *Client, reqs *data.Requests, services []*kubectl.Service, ids []int) *ChaosResult {
	log.Logf("test: %d, services: %d", id, ids)
	result := &ChaosResult{
		ID:         id,
		ServiceIDs: ids,
	}

	signals := make(chan struct{}, 1)
	t1 := time.Now()
	for i := range services {
		service := services[i]
		go func(signals chan struct{}) {
			kubectl.ScaleDeployment(service.Name, service.NameSpace, 0)

			signals <- struct{}{}
		}(signals)
	}
	for i := 0; i < len(services); i++ {
		<-signals
	}
	if result.ScaleIn.Retry, result.ScaleIn.WaitReadyTime = CheckAndWaitReady(); result.ScaleIn.Time != 0 {
		log.Logf("check and wait ready after scale in, retry: %d, time: %v", result.ScaleIn.Retry, result.ScaleIn.Time)
	}
	t2 := time.Now()

	rsp, err := client.SendRequests(reqs)
	t3 := time.Now()
	if err != nil {
		_ = rsp
		log.Logf("send request error: %s", err)
		result.SendRequestError = err
	}

	for i := range services {
		service := services[i]
		go func(signals chan struct{}) {
			kubectl.ScaleDeployment(service.Name, service.NameSpace, service.Replicas)

			signals <- struct{}{}
		}(signals)
	}
	for i := 0; i < len(services); i++ {
		<-signals
	}

	if result.ScaleOut.Retry, result.ScaleOut.WaitReadyTime = CheckAndWaitReady(); result.ScaleOut.WaitReadyTime != 0 {
		log.Logf("check and wait ready after scale out, retry: %d, time: %v", result.ScaleOut.Retry, result.ScaleOut.WaitReadyTime)
	}
	t4 := time.Now()

	//log.Logf("scale in time: %v", t2.Sub(t1))
	//log.Logf("test time: %v", t3.Sub(t2))
	//log.Logf("scale out time: %v", t4.Sub(t3))
	result.TotalTime = t4.Sub(t1)
	result.ScaleIn.Time = t2.Sub(t1)
	result.TestTime = t3.Sub(t2)
	result.ScaleOut.Time = t4.Sub(t3)

	return result
}

func CheckAndWaitReady() (int, time.Duration) {
	if !kubectl.AllReady() {
		retry := 0
		//log.Logf("wait ready")

		t1 := time.Now()
		for ; retry < 120; retry++ {
			if kubectl.AllReady() {
				break
			}

			if pods := kubectl.GetCrashLoopBackOff(); len(pods.Pods) != 0 {
				log.Logf("crash loop back off: %v", pods)

				signals := make(chan struct{}, 1)
				for i := range pods.Pods {
					pod := pods.Pods[i]
					go func(signals chan struct{}) {
						kubectl.ScaleDeployment(pod.Type, "default", 0)

						signals <- struct{}{}
					}(signals)
				}
				for i := 0; i < len(pods.Pods); i++ {
					<-signals
				}

				for i := range pods.Pods {
					pod := pods.Pods[i]
					go func(signals chan struct{}) {
						kubectl.ScaleDeployment(pod.Type, "default", 1)

						signals <- struct{}{}
					}(signals)
				}
				for i := 0; i < len(pods.Pods); i++ {
					<-signals
				}
			}

			if (retry+1)%10 == 0 {
				log.Logf("retry %d times", retry+1)
			}

			time.Sleep(time.Second - time.Duration(time.Now().UnixNano())%time.Second)
		}
		t2 := time.Now()
		//if retry == 60 {
		//	log.Logf("wait ready fail")
		//}

		//log.Logf("wait ready time: %v", t2.Sub(t1))
		return retry, t2.Sub(t1)
	}

	return 0, 0
}
