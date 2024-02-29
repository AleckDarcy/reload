package proto

func NewConditionMessageNode(typ ConditionType, op ConditionOperator, val int64) *ConditionNode {
	return &ConditionNode{
		Type:    ConditionNodeType_ConditionMessage_,
		Message: &ConditionMessage{Type: typ, Op: op, Value: val},
	}
}

func NewConditionLogicNode(typ LogicType, parent int64, list []int64) *ConditionNode {
	return &ConditionNode{
		Type:  ConditionNodeType_ConditionLogic_,
		Logic: &ConditionLogic{Type: typ, Parent: parent, List: list},
	}
}

func NewConditionTree(nodes []*ConditionNode, leafIDs []int64) *ConditionTree {
	return &ConditionTree{Nodes: nodes, LeafIDs: leafIDs}
}

func NewPrerequisiteMessage(name string, condTree *ConditionTree, parent int64, list []int64) *PrerequisiteMessage {
	return &PrerequisiteMessage{Name: name, CondTree: condTree, Parent: parent}
}

func NewPrerequisiteLogic(typ LogicType, parent int64, list []int64) *PrerequisiteLogic {
	return &PrerequisiteLogic{Type: typ, Parent: parent, List: list}
}

func NewPrerequisiteMessageNode(id int64, name string, condTree *ConditionTree, parent int64, list []int64) *PrerequisiteNode {
	return &PrerequisiteNode{
		Id:      id,
		Type:    PrerequisiteNodeType_PrerequisiteMessage_,
		Message: NewPrerequisiteMessage(name, condTree, parent, list)}
}

func NewPrerequisiteLogicNode(id int64, typ LogicType, parent int64, list []int64) *PrerequisiteNode {
	return &PrerequisiteNode{
		Id:    id,
		Type:  PrerequisiteNodeType_PrerequisiteLogic_,
		Logic: NewPrerequisiteLogic(typ, parent, list),
	}
}

func NewPath(typ PathType, path []string) *Path {
	return &Path{Type: typ, Path: path}
}

func NewAttributeConfigure(name string, path *Path) *AttributeConfigure {
	return &AttributeConfigure{Name: name, Path: path}
}

func NewLoggingConfigure(ts *TimestampConfigure, st *StackTraceConfigure, attrs []*AttributeConfigure, out LogOutType) *LoggingConfigure {
	return &LoggingConfigure{Timestamp: ts, Stacktrace: st, Attrs: attrs, Out: out}
}
