package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// Preset PrerequisiteTree's for testing
// Condition naming: cond[tree_id]_[node_id]+_[id]

var cond0_2_0 = NewConditionMessageNode(cb.ConditionType_NumOfInvok, cb.ConditionOperator_EQ, 1)

// (EventA) && (EventB = 1)
var tree0 = &PrerequisiteTree{
	PrerequisiteTree: cb.PrerequisiteTree{
		Nodes: []*cb.PrerequisiteNode{
			NewPrerequisiteLogicNode(0, cb.LogicType_And_, -1, []int64{1, 2}),
			NewPrerequisiteMessageNode(1, "EventA", nil, 0, nil),
			NewPrerequisiteMessageNode(2, "EventB", NewConditionTree([]*cb.ConditionNode{cond0_2_0}, nil), 0, nil),
		},
	},
}

var cond1_3_0 = cond0_2_0
var cond1_4_0_0 = NewConditionMessageNode(cb.ConditionType_NumOfInvok, cb.ConditionOperator_GT, 1)
var cond1_4_0_1 = NewConditionMessageNode(cb.ConditionType_NumOfInvok, cb.ConditionOperator_LT, 4)
var cond1_4_0 = NewConditionLogicNode(cb.LogicType_And_, -1, []int64{1, 2})

// ((EventA) && (EventB = 1)) || (1 < EventC < 4)
var tree1 = &PrerequisiteTree{
	PrerequisiteTree: cb.PrerequisiteTree{
		Nodes: []*cb.PrerequisiteNode{
			NewPrerequisiteLogicNode(0, cb.LogicType_Or_, -1, []int64{1, 4}),
			NewPrerequisiteLogicNode(1, cb.LogicType_And_, 0, []int64{2, 3}),
			NewPrerequisiteMessageNode(2, "EventA", nil, 1, nil),
			NewPrerequisiteMessageNode(3, "EventB", NewConditionTree([]*cb.ConditionNode{cond1_3_0}, nil), 1, nil),
			NewPrerequisiteMessageNode(4, "EventC", NewConditionTree([]*cb.ConditionNode{cond1_4_0, cond1_4_0_0, cond1_4_0_1}, nil), 0, nil),
		},
	},
}

func init() {
	tree0.Indexing()
	tree1.Indexing()
}
