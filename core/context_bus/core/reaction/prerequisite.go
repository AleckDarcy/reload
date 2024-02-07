package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

func NewCondition(typ cb.ConditionType, op cb.ConditionOperator, val int64) *cb.Condition {
	return &cb.Condition{Type: typ, Op: op, Value: val}
}

func NewPrerequisiteMessage(name string, conds []*cb.Condition, parent int64, list []int64) *cb.PrerequisiteMessage {
	return &cb.PrerequisiteMessage{MessageName: name, Conds: conds, Parent: parent, List: list}
}

func NewPrerequisiteLogic(typ cb.PrerequisiteLogicType, parent int64, list []int64) *cb.PrerequisiteLogic {
	return &cb.PrerequisiteLogic{Type: typ, Parent: parent, List: list}
}

func NewPrerequisiteMessageNode(id int64, name string, conds []*cb.Condition, parent int64, list []int64) *cb.PrerequisiteNode {
	return &cb.PrerequisiteNode{
		Id:      id,
		Type:    cb.PrerequisiteNodeType_Message,
		Message: NewPrerequisiteMessage(name, conds, parent, list)}
}

func NewPrerequisiteLogicNode(id int64, typ cb.PrerequisiteLogicType, parent int64, list []int64) *cb.PrerequisiteNode {
	return &cb.PrerequisiteNode{
		Id:    id,
		Type:  cb.PrerequisiteNodeType_Logic,
		Logic: NewPrerequisiteLogic(typ, parent, list),
	}
}

func (t *PrerequisiteTree) Indexing() {
	t.Index = map[string]*PrerequisiteNode{}
	for _, node := range t.Pres {
		if node.Type == cb.PrerequisiteNodeType_Message {
			t.Index[node.Message.MessageName] = (*PrerequisiteNode)(node)
		}
	}
}
