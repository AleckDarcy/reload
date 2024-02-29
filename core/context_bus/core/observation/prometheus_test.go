package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"

	"testing"
)

var prometheusCfg = &cb.PrometheusConfiguration{
	Counters: []*cb.PrometheusOpts{
		{
			Id:          0,
			Namespace:   "test_application",
			Subsystem:   "test_service",
			Name:        "http_request_count",
			Help:        "",
			ConstLabels: nil,
			LabelNames:  []string{"handler", "method"},
		},
	},
	Gauges: []*cb.PrometheusOpts{
		{
			Id:          0,
			Namespace:   "test_application",
			Subsystem:   "test_service",
			Name:        "cpu_usage",
			Help:        "",
			ConstLabels: nil,
			LabelNames:  nil,
		},
	},
	Histograms: []*cb.PrometheusHistogramOpts{
		{
			Id:          0,
			Namespace:   "test_application",
			Subsystem:   "test_service",
			Name:        "http_request_latency",
			Help:        "",
			ConstLabels: nil,
			Buckets:     []float64{1, 10, 100, 1000, 10000},
			LabelNames:  []string{"handler", "method"},
		},
	},
	Summaries: []*cb.PrometheusSummaryOpts{
		{
			Id:          0,
			Namespace:   "test_application",
			Subsystem:   "test_service",
			Name:        "http_request_latency",
			Help:        "",
			ConstLabels: nil,
			Objectives: []*cb.PrometheusSummaryObjective{
				{Key: 0.5, Value: 0.05},
				{Key: 0.9, Value: 0.01},
				{Key: 0.99, Value: 0.001},
			},
			MaxAge:     int64(prometheus.DefMaxAge),
			AgeBuckets: prometheus.DefAgeBuckets,
			BufCap:     prometheus.DefBufCap,
			LabelNames: []string{"handler", "method"},
		},
	},
}

func TestMetrics(t *testing.T) {
	MetricVecStore.Set(prometheusCfg)

	cnt := MetricVecStore.getCounter(0)
	cnt.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Inc()
	cnt.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Inc()
	cnt.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Inc()
	cnt.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Inc()
	cnt.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Inc()

	gau := MetricVecStore.getGauge(0)
	gau.With(nil).Set(99.9)

	his := MetricVecStore.getHistogram(0)
	his.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Observe(1)
	his.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Observe(10)
	his.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(100)
	his.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(1000)
	his.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(10000)

	sum := MetricVecStore.getSummary(0)
	sum.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Observe(1)
	sum.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Observe(10)
	sum.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(100)
	sum.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(1000)
	sum.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(10000)

	pgwOK := prometheusFakeGateway
	defer pgwOK.Close()

	pusher := push.New(pgwOK.URL, "test")
	//_ = pusher.Collector(cnt).Collector(gau).Collector(his).Push()
	//
	//_ = pusher.Collector(sum).Push()

	MetricVecStore.Push(pusher)
}
