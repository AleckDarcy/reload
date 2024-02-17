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
					{
						Id:   0,
						Type: cb.PrerequisiteNodeType_PrerequisiteMessage_,
						Message: &cb.PrerequisiteMessage{
							Name:     "EventA",
							CondTree: &cb.ConditionTree{},
							Parent:   -1,
						},
					},
				},
			}},
		"EventC": {
			Type: cb.ReactionType_FaultCrash,
			PreTree: &cb.PrerequisiteTree{
				Nodes: []*cb.PrerequisiteNode{
					{
						Id:   0,
						Type: cb.PrerequisiteNodeType_PrerequisiteLogic_,
						Logic: &cb.PrerequisiteLogic{
							Type:   cb.LogicType_And_,
							Parent: -1,
							List:   []int64{1, 2},
						},
					}, {
						Id:   1,
						Type: cb.PrerequisiteNodeType_PrerequisiteMessage_,
						Message: &cb.PrerequisiteMessage{
							Name: "EventA",
							CondTree: &cb.ConditionTree{
								Nodes: []*cb.ConditionNode{
									{
										Type: cb.ConditionNodeType_ConditionMessage_,
										Message: &cb.ConditionMessage{
											Type:  cb.ConditionType_NumOfInvok,
											Op:    cb.ConditionOperator_GE,
											Value: 1,
										},
									},
								},
							},
							Parent: 0,
						},
					}, {
						Id:   2,
						Type: cb.PrerequisiteNodeType_PrerequisiteMessage_,
						Message: &cb.PrerequisiteMessage{
							Name: "EventB",
							CondTree: &cb.ConditionTree{
								Nodes: []*cb.ConditionNode{
									{
										Type: cb.ConditionNodeType_ConditionMessage_,
										Message: &cb.ConditionMessage{
											Type:  cb.ConditionType_NumOfInvok,
											Op:    cb.ConditionOperator_GE,
											Value: 0,
										},
									},
								},
							},
							Parent: 0,
						},
					},
				},
			},
		},
	},
}

func TestObservation(t *testing.T) {
	background.Run()
	go Bus.Run(nil)

	id := int64(1)
	configure.ConfigureStore.SetConfigure(id, cfg1)

	ctx := context.NewContext(context.NewRequestContext("rest", id, rest), nil)

	app := new(cb.EventMessage).SetMessage("received message from %s").SetPaths([]*cb.Path{cb.Test_Path_Rest_From})

	// func ServiceHandler(ctx, request) (response, error)
	// generated code
	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA"}, app)
	// application

	time.Sleep(public.ENV_PROFILE_INTERVAL)

	OnSubmission(ctx, &cb.EventWhere{}, &cb.EventRecorder{Name: "EventA"}, app)

	time.Sleep(time.Second)
}
