package observation

import (
	"fmt"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"

	"sync"
)

var prometheusFakeGateway = httptest.NewServer(
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		lastBody, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("prometheus fake gateway received payload, len %d, body %v\n", len(lastBody), lastBody)
		w.WriteHeader(http.StatusOK)
	}),
)

var PrometheusPusher = push.New(prometheusFakeGateway.URL, "test")

type metricVecStore struct {
	lock sync.Mutex

	counters   map[int64]*counterVecWrap
	gauges     map[int64]*gaugeVecWrap
	histograms map[int64]*histogramVecWrap
	summaries  map[int64]*summaryVecWrap

	pushCounters   []int64
	pushGauges     []int64
	pushHistograms []int64
	pushSummaries  []int64
}

var MetricVecStore = &metricVecStore{
	counters:   map[int64]*counterVecWrap{},
	gauges:     map[int64]*gaugeVecWrap{},
	histograms: map[int64]*histogramVecWrap{},
	summaries:  map[int64]*summaryVecWrap{},
}

func (s *metricVecStore) Lock() {
	s.lock.Lock()
}

func (s *metricVecStore) Unlock() {
	s.lock.Unlock()
}

func (s *metricVecStore) Set(cfg *cb.PrometheusConfiguration) {
	s.lock.Lock()

	for _, opt := range cfg.Counters {
		vec := prometheus.NewCounterVec(opt.ToPrometheusCounterOpts())
		prometheus.Register(vec)
		s.counters[opt.Id] = &counterVecWrap{
			id:  opt.Id,
			vec: vec,
		}
	}

	for _, opt := range cfg.Gauges {
		vec := prometheus.NewGaugeVec(opt.ToPrometheusGaugeOpts())
		prometheus.Register(vec)
		s.gauges[opt.Id] = &gaugeVecWrap{
			id:  opt.Id,
			vec: vec,
		}
	}

	for _, opt := range cfg.Histograms {
		vec := prometheus.NewHistogramVec(opt.ToPrometheus())
		prometheus.Register(vec)
		s.histograms[opt.Id] = &histogramVecWrap{
			id:  opt.Id,
			vec: vec,
		}
	}

	for _, opt := range cfg.Summaries {
		vec := prometheus.NewSummaryVec(opt.ToPrometheus())
		prometheus.Register(vec)
		s.summaries[opt.Id] = &summaryVecWrap{
			id:  opt.Id,
			vec: vec,
		}
	}

	s.lock.Unlock()
}

func (s *metricVecStore) resetWrap() {
	for _, wrap := range s.counters {
		prometheus.Unregister(wrap.vec)
	}

	for _, wrap := range s.gauges {
		prometheus.Unregister(wrap.vec)
	}

	for _, wrap := range s.histograms {
		prometheus.Unregister(wrap.vec)
	}

	for _, wrap := range s.summaries {
		prometheus.Unregister(wrap.vec)
	}
}

func (s *metricVecStore) resetPush() {
	for _, id := range s.pushCounters {
		s.counters[id].Pushed()
	}

	for _, id := range s.pushGauges {
		s.gauges[id].Pushed()
	}

	for _, id := range s.pushHistograms {
		s.histograms[id].Pushed()
	}

	for _, id := range s.pushSummaries {
		s.summaries[id].Pushed()
	}

	s.pushCounters = s.pushCounters[:0]
	s.pushGauges = s.pushGauges[:0]
	s.pushHistograms = s.pushHistograms[:0]
	s.pushSummaries = s.pushSummaries[:0]
}

func (s *metricVecStore) Reset() {
	s.lock.Lock()

	s.resetWrap()
	s.resetPush()

	s.lock.Unlock()
}

func (s *metricVecStore) Push(pusher *push.Pusher) {
	s.lock.Lock()

	for _, id := range s.pushCounters {
		wrap := s.counters[id]
		pusher.Collector(wrap.vec)
	}

	for _, id := range s.pushGauges {
		wrap := s.gauges[id]
		pusher.Collector(wrap.vec)
	}

	for _, id := range s.pushHistograms {
		wrap := s.histograms[id]
		pusher.Collector(wrap.vec)
	}

	pusher.Push()

	// todo summary won't push, don't know why
	for _, id := range s.pushSummaries {
		wrap := s.summaries[id]
		pusher.Collector(wrap.vec)
	}

	s.resetPush()

	s.lock.Unlock()
}

func (s *metricVecStore) setCounter(id int64, vec *counterVecWrap) {
	s.counters[id] = vec
}

func (s *metricVecStore) getCounter(id int64) *prometheus.CounterVec {
	if wrap := s.counters[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushCounters = append(s.pushCounters, id)
		}

		return vec
	}

	return nil
}

func (s *metricVecStore) deleteCounter(id int64) {
	delete(s.counters, id)
}

func (s *metricVecStore) setGauge(id int64, vec *gaugeVecWrap) {
	s.gauges[id] = vec
}

func (s *metricVecStore) getGauge(id int64) *prometheus.GaugeVec {
	if wrap := s.gauges[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushGauges = append(s.pushGauges, id)
		}

		return vec
	}

	return nil
}

func (s *metricVecStore) deleteGauge(id int64) {
	delete(s.gauges, id)
}

func (s *metricVecStore) setHistogram(id int64, vec *histogramVecWrap) {
	s.histograms[id] = vec
}

func (s *metricVecStore) getHistogram(id int64) *prometheus.HistogramVec {
	if wrap := s.histograms[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushHistograms = append(s.pushHistograms, id)
		}

		return vec
	}

	return nil
}

func (s *metricVecStore) deleteHistogram(id int64) {
	delete(s.histograms, id)
}

func (s *metricVecStore) setSummary(id int64, vec *summaryVecWrap) {
	s.summaries[id] = vec
}

func (s *metricVecStore) getSummary(id int64) *prometheus.SummaryVec {
	if wrap := s.summaries[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushSummaries = append(s.pushSummaries, id)
		}

		return vec
	}

	return nil
}

func (s *metricVecStore) deleteSummary(id int64) {
	delete(s.summaries, id)
}

type counterVecWrap struct {
	id   int64
	vec  *prometheus.CounterVec
	push bool
}

type gaugeVecWrap struct {
	id   int64
	vec  *prometheus.GaugeVec
	push bool
}

type histogramVecWrap struct {
	id   int64
	vec  *prometheus.HistogramVec
	push bool
}

type summaryVecWrap struct {
	id   int64
	vec  *prometheus.SummaryVec
	push bool
}

func (w *counterVecWrap) GetVec() (vec *prometheus.CounterVec, pushFlag bool) {
	if pushFlag = !w.push; pushFlag {
		w.push = true
	}

	vec = w.vec

	return
}

func (w *counterVecWrap) Pushed() {
	w.vec.Reset()
	w.push = false
}

func (w *gaugeVecWrap) GetVec() (vec *prometheus.GaugeVec, pushFlag bool) {
	if pushFlag = !w.push; pushFlag {
		w.push = true
	}

	vec = w.vec

	return
}

func (w *gaugeVecWrap) Pushed() {
	w.vec.Reset()
	w.push = false
}

func (w *histogramVecWrap) GetVec() (vec *prometheus.HistogramVec, pushFlag bool) {
	if pushFlag = !w.push; pushFlag {
		w.push = true
	}

	vec = w.vec

	return
}

func (w *histogramVecWrap) Pushed() {
	w.vec.Reset()
	w.push = false
}

func (w *summaryVecWrap) GetVec() (vec *prometheus.SummaryVec, pushFlag bool) {
	if pushFlag = !w.push; pushFlag {
		w.push = true
	}

	vec = w.vec

	return
}

func (w *summaryVecWrap) Pushed() {
	w.vec.Reset()
	w.push = false
}
