package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"

	"sync"
)

type vecStore struct {
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

var VecStore = &vecStore{
	counters:   map[int64]*counterVecWrap{},
	gauges:     map[int64]*gaugeVecWrap{},
	histograms: map[int64]*histogramVecWrap{},
	summaries:  map[int64]*summaryVecWrap{},
}

func (s *vecStore) Lock() {
	s.lock.Lock()
}

func (s *vecStore) Unlock() {
	s.lock.Unlock()
}

func (s *vecStore) Set(cfg *cb.PrometheusConfiguration) {
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

func (s *vecStore) resetWrap() {
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

func (s *vecStore) resetPush() {
	s.pushCounters = s.pushCounters[:0]
	s.pushGauges = s.pushGauges[:0]
	s.pushHistograms = s.pushHistograms[:0]
	s.pushSummaries = s.pushSummaries[:0]
}

func (s *vecStore) Reset() {
	s.lock.Lock()

	s.resetWrap()
	s.resetPush()

	s.lock.Unlock()
}

func (s *vecStore) Push(pusher push.Pusher) {
	s.lock.Lock()

	for _, id := range s.pushCounters {
		wrap := s.counters[id]
		pusher.Collector(wrap.vec)
		wrap.Pushed()
	}

	for _, id := range s.pushGauges {
		wrap := s.gauges[id]
		pusher.Collector(wrap.vec)
		wrap.Pushed()
	}

	for _, id := range s.pushHistograms {
		wrap := s.histograms[id]
		pusher.Collector(wrap.vec)
		wrap.Pushed()
	}

	for _, id := range s.pushSummaries {
		wrap := s.summaries[id]
		pusher.Collector(wrap.vec)
		wrap.Pushed()
	}

	s.resetPush()

	s.lock.Unlock()
}

func (s *vecStore) setCounter(id int64, vec *counterVecWrap) {
	s.counters[id] = vec
}

func (s *vecStore) getCounter(id int64) *prometheus.CounterVec {
	if wrap := s.counters[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushCounters = append(s.pushCounters, id)
		}

		return vec
	}

	return nil
}

func (s *vecStore) deleteCounter(id int64) {
	delete(s.counters, id)
}

func (s *vecStore) setGauge(id int64, vec *gaugeVecWrap) {
	s.gauges[id] = vec
}

func (s *vecStore) getGauge(id int64) *prometheus.GaugeVec {
	if wrap := s.gauges[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushGauges = append(s.pushGauges, id)
		}

		return vec
	}

	return nil
}

func (s *vecStore) deleteGauge(id int64) {
	delete(s.gauges, id)
}

func (s *vecStore) setHistogram(id int64, vec *histogramVecWrap) {
	s.histograms[id] = vec
}

func (s *vecStore) getHistogram(id int64) *prometheus.HistogramVec {
	if wrap := s.histograms[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushHistograms = append(s.pushHistograms, id)
		}

		return vec
	}

	return nil
}

func (s *vecStore) deleteHistogram(id int64) {
	delete(s.histograms, id)
}

func (s *vecStore) setSummary(id int64, vec *summaryVecWrap) {
	s.summaries[id] = vec
}

func (s *vecStore) getSummary(id int64) *prometheus.SummaryVec {
	if wrap := s.summaries[id]; wrap != nil {
		vec, pushFlag := wrap.GetVec()

		if pushFlag {
			s.pushSummaries = append(s.pushSummaries, id)
		}

		return vec
	}

	return nil
}

func (s *vecStore) deleteSummary(id int64) {
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
