package reaction

import (
	"errors"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// Check if prerequisites are accomplished.
// Using greedy strategy: does not detect lineage among prerequisite nodes.

func (c *ConditionMessage) Check(val int64) (bool, error) {
	switch c.Op {
	case cb.ConditionOperator_LT:
		return val < c.Value, nil
	case cb.ConditionOperator_GT:
		return val > c.Value, nil
	case cb.ConditionOperator_LE:
		return val <= c.Value, nil
	case cb.ConditionOperator_GE:
		return val >= c.Value, nil
	case cb.ConditionOperator_EQ:
		return val == c.Value, nil
	case cb.ConditionOperator_NE:
		return val != c.Value, nil
	default:
		return false, errors.New("unsupported operation type")
	}
}

func (c *ConditionLogic) Check(nodes []*cb.ConditionNode, val int64) (bool, error) {
	switch c.Type {
	case cb.LogicType_And_:
		for _, nodeID := range c.List {
			if acc, err := (*ConditionNode)(nodes[nodeID]).Check(nodes, val); err != nil {
				return false, err
			} else if !acc {
				return false, nil
			}
		}

		return true, nil
	case cb.LogicType_Or_:
		for _, nodeID := range c.List {
			if acc, err := (*ConditionNode)(nodes[nodeID]).Check(nodes, val); err != nil {
				return false, err
			} else if acc {
				return true, nil
			}
		}

		return false, nil
	default:
		return false, errors.New("unsupported PrerequisiteLogicType")
	}
}

func (n *ConditionNode) Check(nodes []*cb.ConditionNode, val int64) (bool, error) {
	if n.Type == cb.ConditionNodeType_ConditionMessage_ {
		return (*ConditionMessage)(n.Message).Check(val)
	} else if n.Type == cb.ConditionNodeType_ConditionLogic_ {
		return (*ConditionLogic)(n.Logic).Check(nodes, val)
	}

	return false, errors.New("unsupported ConditionNodeType")
}

// Check
// Note that Check returns true if Conds is empty
func (m *PrerequisiteMessage) Check(snapshot *PrerequisiteSnapshot, id int64) (bool, error) {
	val := snapshot.Value[id]

	if len(m.Conds) != 0 {
		return (*ConditionNode)(m.Conds[0]).Check(m.Conds, val)
	}

	return true, nil
}

// Check
// Note that Check returns true if List is empty
func (l *PrerequisiteLogic) Check(tree *PrerequisiteTree, snapshot *PrerequisiteSnapshot, id int64) (bool, error) {
	switch l.Type {
	case cb.LogicType_And_:
		for _, nodeID := range l.List {
			if acc, err := (*PrerequisiteNode)(tree.Nodes[nodeID]).Check(tree, snapshot, nodeID); err != nil {
				return false, err
			} else if !acc {
				return false, nil
			}
		}

		return true, nil
	case cb.LogicType_Or_:
		for _, nodeID := range l.List {
			if acc, err := (*PrerequisiteNode)(tree.Nodes[nodeID]).Check(tree, snapshot, nodeID); err != nil {
				return false, err
			} else if acc {
				return true, nil
			}
		}

		return false, nil
	default:
		return false, errors.New("unsupported PrerequisiteLogicType")
	}
}

func (n *PrerequisiteNode) Check(tree *PrerequisiteTree, snapshot *PrerequisiteSnapshot, id int64) (bool, error) {
	if n.Type == cb.PrerequisiteNodeType_PrerequisiteMessage_ {
		return (*PrerequisiteMessage)(n.Message).Check(snapshot, id)
	} else if n.Type == cb.PrerequisiteNodeType_PrerequisiteLogic_ {
		return (*PrerequisiteLogic)(n.Logic).Check(tree, snapshot, id)
	}

	return false, errors.New("unsupported PrerequisiteNodeType")
}

func (t *PrerequisiteTree) Check(snapshot *PrerequisiteSnapshot) (bool, error) {
	if len(t.Nodes) != len(snapshot.Value) {
		return false, errors.New("prerequisite length not match")
	}

	// top-down
	if ok, err := (*PrerequisiteNode)(t.Nodes[0]).Check(t, snapshot, 0); err != nil {
		return false, err
	} else if !ok {
		return false, nil
	}

	return true, nil
}
