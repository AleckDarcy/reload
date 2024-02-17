package proto

// Test_PrerequisiteTree0 (EventA) && (EventB = 1)
var Test_PrerequisiteTree0 = &PrerequisiteTree{
	Nodes: []*PrerequisiteNode{
		NewPrerequisiteLogicNode(0, LogicType_And_, -1, []int64{1, 2}),
		NewPrerequisiteMessageNode(1, "EventA", nil, 0, nil),
		NewPrerequisiteMessageNode(2, "EventB", NewConditionTree([]*ConditionNode{Test_Condition_0_2_0}, nil), 0, nil),
	},
}

// Test_PrerequisiteTree1 ((EventA) && (EventB = 1)) || (1 < EventC < 4)
var Test_PrerequisiteTree1 = &PrerequisiteTree{
	Nodes: []*PrerequisiteNode{
		NewPrerequisiteLogicNode(0, LogicType_Or_, -1, []int64{1, 4}),
		NewPrerequisiteLogicNode(1, LogicType_And_, 0, []int64{2, 3}),
		NewPrerequisiteMessageNode(2, "EventA", nil, 1, nil),
		NewPrerequisiteMessageNode(3, "EventB", NewConditionTree([]*ConditionNode{Test_Condition_1_3_0}, nil), 1, nil),
		NewPrerequisiteMessageNode(4, "EventC", NewConditionTree([]*ConditionNode{Test_Condition_1_4_0, Test_Condition_1_4_0_0, Test_Condition_1_4_0_1}, nil), 0, nil),
	},
}
