package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

func NewConditionMessage(typ cb.ConditionType, op cb.ConditionOperator, val int64) *cb.ConditionMessage {
	return &cb.ConditionMessage{Type: typ, Op: op, Value: val}
}

func NewConditionLogic(typ cb.LogicType, parent int64, list []int64) *cb.ConditionLogic {
	return &cb.ConditionLogic{
		Type:   typ,
		Parent: parent,
		List:   list,
	}
}

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

func NewPrerequisiteMessage(name string, conds []*cb.ConditionNode, parent int64, list []int64) *cb.PrerequisiteMessage {
	return &cb.PrerequisiteMessage{Name: name, Conds: conds, Parent: parent, List: list}
}

func NewPrerequisiteLogic(typ cb.LogicType, parent int64, list []int64) *cb.PrerequisiteLogic {
	return &cb.PrerequisiteLogic{Type: typ, Parent: parent, List: list}
}

func NewPrerequisiteMessageNode(id int64, name string, conds []*cb.ConditionNode, parent int64, list []int64) *cb.PrerequisiteNode {
	return &cb.PrerequisiteNode{
		Id:      id,
		Type:    cb.PrerequisiteNodeType_PrerequisiteMessage_,
		Message: NewPrerequisiteMessage(name, conds, parent, list)}
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
