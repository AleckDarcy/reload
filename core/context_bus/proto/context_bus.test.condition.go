package proto

// global variables for testing only

// Condition naming: cond[tree_id]_[parent_node_id]_..._[node_id]_[id]

var Test_Condition_0_2_0 = NewConditionMessageNode(ConditionType_NumOfInvok, ConditionOperator_EQ, 1)

var Test_Condition_1_3_0 = Test_Condition_0_2_0
var Test_Condition_1_4_0_0 = NewConditionMessageNode(ConditionType_NumOfInvok, ConditionOperator_GT, 1)
var Test_Condition_1_4_0_1 = NewConditionMessageNode(ConditionType_NumOfInvok, ConditionOperator_LT, 4)
var Test_Condition_1_4_0 = NewConditionLogicNode(LogicType_And_, -1, []int64{1, 2})
