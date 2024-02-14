package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

func NewConditionMessageNode(typ cb.ConditionType, op cb.ConditionOperator, val int64) *cb.ConditionNode {
	return &cb.ConditionNode{
		Type:    cb.ConditionNodeType_ConditionMessage_,
		Message: &cb.ConditionMessage{Type: typ, Op: op, Value: val},
	}
}

func NewConditionLogicNode(typ cb.LogicType, parent int64, list []int64) *cb.ConditionNode {
	return &cb.ConditionNode{
		Type:  cb.ConditionNodeType_ConditionLogic_,
		Logic: &cb.ConditionLogic{Type: typ, Parent: parent, List: list},
	}
}

func NewConditionTree(nodes []*cb.ConditionNode, leafIDs []int64) *cb.ConditionTree {
	return &cb.ConditionTree{
		Nodes:   nodes,
		LeafIDs: leafIDs,
	}
}

func NewPrerequisiteMessage(name string, condTree *cb.ConditionTree, parent int64, list []int64) *cb.PrerequisiteMessage {
	return &cb.PrerequisiteMessage{Name: name, CondTree: condTree, Parent: parent, List: list}
}

func NewPrerequisiteLogic(typ cb.LogicType, parent int64, list []int64) *cb.PrerequisiteLogic {
	return &cb.PrerequisiteLogic{Type: typ, Parent: parent, List: list}
}

func NewPrerequisiteMessageNode(id int64, name string, condTree *cb.ConditionTree, parent int64, list []int64) *cb.PrerequisiteNode {
	return &cb.PrerequisiteNode{
		Id:      id,
		Type:    cb.PrerequisiteNodeType_PrerequisiteMessage_,
		Message: NewPrerequisiteMessage(name, condTree, parent, list)}
}

func NewPrerequisiteLogicNode(id int64, typ cb.LogicType, parent int64, list []int64) *cb.PrerequisiteNode {
	return &cb.PrerequisiteNode{
		Id:    id,
		Type:  cb.PrerequisiteNodeType_PrerequisiteLogic_,
		Logic: NewPrerequisiteLogic(typ, parent, list),
	}
}

func (t *PrerequisiteTree) Indexing() {
	t.Index = map[string]*PrerequisiteNode{}
	for _, node := range t.Nodes {
		if node.Type == cb.PrerequisiteNodeType_PrerequisiteMessage_ {
			t.Index[node.Message.Name] = (*PrerequisiteNode)(node)
		}
	}
}
