package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/background"
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	"github.com/AleckDarcy/reload/core/context_bus/core/observation"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/AleckDarcy/reload/core/context_bus/public"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"

	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

var path = cb.Test_Path_Rest_From
var rest = cb.Test_EventMessage_Rest

var rest2 = &cb.EventMessage{
	Attrs: &cb.Attributes{
		Attrs: map[string]*cb.AttributeValue{
			"from": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "handler1",
			},
			"method": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "POST",
			},
			"handler": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "/handler2",
			},
			"key": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "This a string attribute",
			},
			"key_": {
				Type: cb.AttributeValueType_AttributeValueStr,
				Str:  "This another string attribute",
			},
		},
	},
}

var cfg1 = &cb.Configure{
	Reactions: nil,
	Observations: map[string]*cb.ObservationConfigure{
		"EventA": {
			Logging: &cb.LoggingConfigure{
				Timestamp: &cb.TimestampConfigure{Format: public.TIME_FORMAT_RFC3339},
				Attrs:     []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key},
				Out:       cb.LogOutType_Stdout,
			},
		},
	},
}

var cfg2 = &cb.Configure{
	Reactions: nil,
	Observations: map[string]*cb.ObservationConfigure{
		"EventA": {
			Logging: &cb.LoggingConfigure{
				Timestamp: &cb.TimestampConfigure{Format: public.TIME_FORMAT_RFC3339Nano},
				Attrs:     []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key},
				Out:       cb.LogOutType_LogOutType_, // omit print
			},
		},
	},
}

var cfg3 = &cb.Configure{
	Reactions: map[string]*cb.ReactionConfigure{
		"EventD": {
			Type: cb.ReactionType_FaultCrash,
			PreTree: &cb.PrerequisiteTree{
				Nodes: []*cb.PrerequisiteNode{
					cb.NewPrerequisiteMessageNode(0, "EventA", &cb.ConditionTree{}, -1, nil),
				},
			}},
		"EventC": {
			Type: cb.ReactionType_FaultCrash,
			PreTree: &cb.PrerequisiteTree{
				Nodes: []*cb.PrerequisiteNode{
					cb.NewPrerequisiteLogicNode(0, cb.LogicType_And_, -1, []int64{1, 2}),
					cb.NewPrerequisiteMessageNode(1, "EventA",
						cb.NewConditionTree([]*cb.ConditionNode{cb.Test_Condition_C_1_0}, nil), 0, nil),
					cb.NewPrerequisiteMessageNode(2, "EventB",
						cb.NewConditionTree([]*cb.ConditionNode{cb.Test_Condition_C_2_0}, nil), 0, nil),
				},
			},
		},
	},
}

type request struct {
	cbPayload *cb.Payload
}

type response struct {
	cbPayload *cb.Payload
}

type prerequisiteSnapshotsStore struct {
	lock sync.RWMutex

	map_ map[string]*cb.PrerequisiteSnapshots
}

var PrerequisiteSnapshotsStore = &prerequisiteSnapshotsStore{
	map_: map[string]*cb.PrerequisiteSnapshots{},
}

func (s *prerequisiteSnapshotsStore) Set(uuid string, store *cb.PrerequisiteSnapshots) {
	s.lock.Lock()
	s.map_[uuid] = store
	s.lock.Unlock()
}

func (s *prerequisiteSnapshotsStore) Get(uuid string) (*cb.PrerequisiteSnapshots, bool) {
	s.lock.RLock()
	ss, ok := s.map_[uuid]
	s.lock.RUnlock()

	return ss, ok
}

func (s *prerequisiteSnapshotsStore) Delete(uuid string) {
	s.lock.Lock()
	delete(s.map_, uuid)
	s.lock.Unlock()
}

var requestNetwork = make(chan *request, 1)
var responseNetwork = make(chan *response, 1)

// mocked blocked inter-service call
func sendRequest(ctx *context.Context, uid string, req *request) (*response, error) {
	PrerequisiteSnapshotsStore.Set(uid, ctx.GetEventContext().GetPrerequisiteSnapshots())

	req.cbPayload = &cb.Payload{
		RequestId: 0,
		ConfigId:  ctx.GetRequestContext().GetConfigureID(),
		Snapshots: ctx.GetEventContext().GetPrerequisiteSnapshots().Clone(),
		MType:     cb.MessageType_Message_Request,
		Uuid:      uid,
	}

	fmt.Println("uuid sent", req.cbPayload.Uuid)
	fmt.Println(req.cbPayload.Snapshots)
	// TODO do reaction

	requestNetwork <- req

	select {
	case rsp := <-responseNetwork:
		fmt.Println("uuid recv", rsp.cbPayload.Uuid)

		ss, ok := PrerequisiteSnapshotsStore.Get(rsp.cbPayload.Uuid)
		fmt.Println(ss)
		if ok {
			ss.MergeOffset(rsp.cbPayload.Snapshots)
		}

		// TODO do reaction

		ctx.GetEventContext().SetPrerequisiteSnapshots(ss)

		return rsp, nil
	case <-time.After(time.Second):
		return nil, errors.New("timeout")
	}
}

// mocked blocked inter-service call
func sendResponse(rsp *response) error {
	responseNetwork <- rsp

	return nil
}

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

func TestObservation(t *testing.T) {
	background.Run()
	go Bus.Run(nil)

	// set MetricVecStore
	observation.MetricVecStore.Set(prometheusCfg)

	var cfg4 = &cb.Configure{
		Observations: map[string]*cb.ObservationConfigure{
			"EventA-starts": {
				Type: cb.ObservationType_ObservationStart,
				Logging: &cb.LoggingConfigure{
					Attrs: []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key, cb.Test_AttributeConfigure_Rest_Key_},
					Out:   cb.LogOutType_Stdout,
				},
				Metrics: []*cb.MetricsConfigure{
					{
						Type: cb.MetricType_Counter,
						Name: "cnt_EventA",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
					{
						Type: cb.MetricType_Counter,
						Name: "api_restful_request_total",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
				},
			},
			"EventA-abcdef": {
				Type: cb.ObservationType_ObservationSingle,
				Logging: &cb.LoggingConfigure{
					Out: cb.LogOutType_Stdout,
				},
			},
			"EventA-bcdefg": {
				Type: cb.ObservationType_ObservationInter,
				Logging: &cb.LoggingConfigure{
					Out: cb.LogOutType_Stdout,
				},
			},
			"EventA-cdefgh": {
				Type: cb.ObservationType_ObservationInter,
				Logging: &cb.LoggingConfigure{
					Out: cb.LogOutType_Stdout,
				},
				Metrics: []*cb.MetricsConfigure{
					{
						Type:     cb.MetricType_Histogram,
						Name:     "lat_RequestC",
						PrevName: "EventA-bcdefg",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
				},
			},
			"EventA-ends": {
				Type: cb.ObservationType_ObservationEnd,
				Logging: &cb.LoggingConfigure{
					Attrs: []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key},
					Out:   cb.LogOutType_Stdout,
				},
				Tracing: &cb.TracingConfigure{
					Name:     "EventA",
					PrevName: "EventA-starts",
					Attrs: []*cb.AttributeConfigure{
						cb.Test_AttributeConfigure_Rest_Method,
						cb.Test_AttributeConfigure_Rest_Handler,
					},
				},
				Metrics: []*cb.MetricsConfigure{
					{
						Type:     cb.MetricType_Histogram,
						Name:     "lat_HandlerA",
						PrevName: "EventA-starts",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
				},
			},
			"EventB-starts": {
				Type: cb.ObservationType_ObservationStart,
				Logging: &cb.LoggingConfigure{
					Attrs: []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key, cb.Test_AttributeConfigure_Rest_Key_},
					Out:   cb.LogOutType_Stdout,
				},
				Metrics: []*cb.MetricsConfigure{
					{
						Type: cb.MetricType_Counter,
						Name: "cnt_EventB",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
					{
						Type: cb.MetricType_Counter,
						Name: "api_restful_request_total",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
				},
			},
			"EventB-ends": {
				Type: cb.ObservationType_ObservationEnd,
				Logging: &cb.LoggingConfigure{
					Attrs: []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key},
					Out:   cb.LogOutType_Stdout,
				},
				Tracing: &cb.TracingConfigure{
					Name:     "EventB",
					PrevName: "EventB-starts",
					Attrs: []*cb.AttributeConfigure{
						cb.Test_AttributeConfigure_Rest_Method,
						cb.Test_AttributeConfigure_Rest_Handler,
					},
				},
				Metrics: []*cb.MetricsConfigure{
					{
						Type:     cb.MetricType_Histogram,
						Name:     "lat_HandlerB",
						PrevName: "EventB-starts",
						Attrs: []*cb.AttributeConfigure{
							cb.Test_AttributeConfigure_Rest_Method,
							cb.Test_AttributeConfigure_Rest_Handler,
						},
					},
				},
			},
		},
		Reactions: map[string]*cb.ReactionConfigure{
			"EventA-cdefgh": {
				Type:   cb.ReactionType_FaultDelay,
				Params: &cb.ReactionConfigure_FaultDelay{FaultDelay: &cb.FaultDelayParam{Ms: 500}},
				PreTree: &cb.PrerequisiteTree{
					Nodes: []*cb.PrerequisiteNode{
						{
							Id:   0,
							Type: cb.PrerequisiteNodeType_PrerequisiteMessage_,
							Message: &cb.PrerequisiteMessage{
								Name:     "EventB-starts",
								CondTree: nil,
								Parent:   -1,
							},
						},
					},
				},
			},
		},
	}

	id := int64(4)
	configure.ConfigureStore.SetConfigure(id, cfg4)
	cfg4_ := configure.ConfigureStore.GetConfigure(id)

	handler2 := func(ctx *context.Context, req *request) (rsp *response, err error) {
		t.Log("handler2 invoked")
		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler2", File: "path/file2.go", Line: 200})
		app1 := new(cb.EventMessage).SetMessage("recv request from %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventB-starts"}, app1)

		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler2", File: "path/file2.go", Line: 240})
		app5 := new(cb.EventMessage).SetMessage("send response to %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventB-ends"}, app5)

		rsp = &response{
			cbPayload: &cb.Payload{
				RequestId: req.cbPayload.RequestId,
				ConfigId:  req.cbPayload.ConfigId,
				Snapshots: ctx.GetEventContext().GetOffsetSnapshots(),
				MType:     cb.MessageType_Message_Response,
				Uuid:      req.cbPayload.Uuid,
			},
		}

		return
	}

	// mocked framework for handler2
	go func() {
		// handler2 framework inbound
		req := <-requestNetwork
		id := req.cbPayload.ConfigId
		cfg := configure.ConfigureStore.GetConfigure(id)
		ctx1 := new(context.Context).
			SetRequestContext(context.NewRequestContext("rest", id, rest2)).
			SetEventContext(new(context.EventContext).SetPrerequisiteSnapshots(req.cbPayload.Snapshots).SetOffsetSnapshots(cfg.InitializeSnapshots()))
		rsp, err := handler2(ctx1, req)
		_ = err
		sendResponse(rsp)
		// handler2 framework outbound
	}()

	handler1 := func(ctx *context.Context, req *request) (rsp *response, err error) {
		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler1", File: "path/file1.go", Line: 140})
		app1 := new(cb.EventMessage).SetMessage("recv request from %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-starts"}, app1)
		//ctx.PrintPrevEventData(t)

		// ---------- start handler1 logic ----------
		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler1", File: "path/file1.go", Line: 145})
		app2 := new(cb.EventMessage).SetMessage("something happened")
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-abcdef"}, app2)
		//ctx.PrintPrevEventData(t)

		// handler1 outbound: calling handler2
		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler1", File: "path/file1.go", Line: 150})
		app3 := new(cb.EventMessage).SetMessage("send request to handler2")
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-bcdefg"}, app3)
		//ctx.PrintPrevEventData(t)

		uid := uuid.New().String()
		t.Logf("send request to handler2, snapshots: %+v", ctx.GetEventContext().GetPrerequisiteSnapshots())
		rsp2, err2 := sendRequest(ctx, uid, &request{})
		t.Logf("receive response from handler2, snapshots: %+v (offset), %+v (updated)", rsp2.cbPayload.Snapshots, ctx.GetEventContext().GetPrerequisiteSnapshots())
		_, _ = rsp2, err2

		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler1", File: "path/file1.go", Line: 155})
		app4 := new(cb.EventMessage).SetMessage("recv response from handler2")
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-cdefgh"}, app4)
		//ctx4.PrintPrevEventData(t)

		// ---------- end handler1 logic ----------
		ctx.GetEventContext().SetCodeInfoBasic(&cb.CodeBaseInfo{Name: "handler1", File: "path/file1.go", Line: 160})
		app5 := new(cb.EventMessage).SetMessage("send response to %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})
		OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-ends"}, app5)
		//ctx.PrintPrevEventData(t)

		return
	}

	start := time.Now().UnixNano()
	ctx := new(context.Context).
		SetRequestContext(context.NewRequestContext("rest", id, rest)).
		SetEventContext(context.NewEventContext(nil, cfg4_.InitializeSnapshots()))
	handler1(ctx, nil)
	end := time.Now().UnixNano()

	t.Logf("duration %d", end-start)

	t.Logf("snapshots: %+v\n", ctx.GetEventContext().GetPrerequisiteSnapshots())

	time.Sleep(time.Second)
}
