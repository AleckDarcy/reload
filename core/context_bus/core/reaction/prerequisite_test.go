package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// Preset PrerequisiteTree's for testing
// Condition naming: cond[tree_id]_[node_id]_[id]

var cond0_2_0 = NewCondition(cb.ConditionType_NumOfInvok, cb.ConditionOperator_EQ, 1)

// (EventA) && (EventB = 1)
var tree0 = &PrerequisiteTree{
	PrerequisiteTree: cb.PrerequisiteTree{
		Pres: []*cb.PrerequisiteNode{
			NewPrerequisiteLogicNode(0, cb.PrerequisiteLogicType_And, -1, []int64{1, 2}),
			NewPrerequisiteMessageNode(1, "EventA", nil, 0, nil),
			NewPrerequisiteMessageNode(2, "EventB", []*cb.Condition{cond0_2_0}, 0, nil),
		},
	},
}

var cond1_3_0 = cond0_2_0
var cond1_4_0 = NewCondition(cb.ConditionType_NumOfInvok, cb.ConditionOperator_GT, 1)
var cond1_4_1 = NewCondition(cb.ConditionType_NumOfInvok, cb.ConditionOperator_LT, 4)

// ((EventA) && (EventB = 1)) || (1 < EventC < 4)
var tree1 = &PrerequisiteTree{
	PrerequisiteTree: cb.PrerequisiteTree{
		Pres: []*cb.PrerequisiteNode{
			NewPrerequisiteLogicNode(0, cb.PrerequisiteLogicType_Or, -1, []int64{1, 4}),
			NewPrerequisiteLogicNode(1, cb.PrerequisiteLogicType_And, 0, []int64{2, 3}),
			NewPrerequisiteMessageNode(2, "EventA", nil, 1, nil),
			NewPrerequisiteMessageNode(3, "EventB", []*cb.Condition{cond1_3_0}, 1, nil),
			NewPrerequisiteMessageNode(4, "EventC", []*cb.Condition{cond1_4_0, cond1_4_1}, 0, nil),
		},
	},
}

func init() {
	tree0.Indexing()
	tree1.Indexing()
}
