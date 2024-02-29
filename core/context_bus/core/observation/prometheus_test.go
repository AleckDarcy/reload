package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"

	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var prometheusCfg = &cb.PrometheusConfiguration{
	Counters: nil,
	Gauges:   nil,
	Histograms: []*cb.PrometheusHistogramOpts{
		{
			Id:         0,
			Namespace:  "test_application",
			Subsystem:  "test_service",
			Name:       "http_response_latency",
			Buckets:    []float64{1, 10, 100, 1000, 10000},
			LabelNames: []string{"handler", "method"},
		},
	},
	Summaries: nil,
}

func TestMetrics(t *testing.T) {
	VecStore.Set(prometheusCfg)

	vec := VecStore.getHistogram(0)
	vec.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Observe(1)
	vec.With(prometheus.Labels{"handler": "handler1", "method": "POST"}).Observe(10)
	vec.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(100)
	vec.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(1000)
	vec.With(prometheus.Labels{"handler": "handler2", "method": "POST"}).Observe(10000)

	pgwOK := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			lastBody, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(lastBody)
			w.WriteHeader(http.StatusOK)
		}),
	)
	defer pgwOK.Close()

	_ = push.New(pgwOK.URL, "test").Collector(vec).Push()
}
