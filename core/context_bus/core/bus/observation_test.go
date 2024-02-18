package bus

import (
	"github.com/AleckDarcy/reload/core/context_bus/background"
	"github.com/AleckDarcy/reload/core/context_bus/core/configure"
	"github.com/AleckDarcy/reload/core/context_bus/core/context"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/AleckDarcy/reload/core/context_bus/public"

	"testing"
	"time"
)

var path = cb.Test_Path_Rest_From
var rest = cb.Test_EventMessage_Rest

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

func TestObservation(t *testing.T) {
	background.Run()
	go Bus.Run(nil)

	var cfg4 = &cb.Configure{
		Observations: map[string]*cb.ObservationConfigure{
			"EventA-starts": {
				Type: cb.ObservationType_ObservationStart,
				Logging: &cb.LoggingConfigure{
					Attrs: []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key, cb.Test_AttributeConfigure_Rest_Key_},
					Out:   cb.LogOutType_Stdout,
				},
				Metrics: []*cb.MetricsConfigure{
					{Type: cb.MetricType_Counter, Name: "cnt_EventA"},
				},
			},
			"EventA-abcdef": {
				Type: cb.ObservationType_ObservationSingle,
				Logging: &cb.LoggingConfigure{
					Out: cb.LogOutType_Stdout,
				},
			},
			"EventA-ends": {
				Type: cb.ObservationType_ObservationEnd,
				Logging: &cb.LoggingConfigure{
					Attrs: []*cb.AttributeConfigure{cb.Test_AttributeConfigure_Rest_Key},
					Out:   cb.LogOutType_Stdout,
				},
				Metrics: []*cb.MetricsConfigure{
					{Type: cb.MetricType_Histogram, Name: "lat_EventA"},
				},
			},
		},
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

	id := int64(4)
	configure.ConfigureStore.SetConfigure(id, cfg4)

	ctx := context.NewContext(context.NewRequestContext("rest", id, rest), context.NewEventContext(nil, nil))

	// func ServiceHandler(ctx, request) (response, error)
	// generated code
	app1 := new(cb.EventMessage).SetMessage("received request from %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})
	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-starts"}, app1)
	//t.Logf("%+v", ctx.GetEventContext())

	// application
	app2 := new(cb.EventMessage).SetMessage("something happened")
	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-abcdef"}, app2)
	//t.Logf("%+v", ctx.GetEventContext())

	app3 := new(cb.EventMessage).SetMessage("sent response to %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})
	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA-ends"}, app3)
	//t.Logf("%+v", ctx.GetEventContext())

	time.Sleep(time.Second)
}
