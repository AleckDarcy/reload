package reaction

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"

	"errors"
)

func (t *PrerequisiteTree) InitializeSnapshot() *PrerequisiteSnapshot {
	return &PrerequisiteSnapshot{
		Value: make([]int64, len(t.Pres)),
	}
}

// update PrerequisiteSnapshot:
// 1. number of occurrence

func (n *PrerequisiteNode) UpdateSnapshot(snapshot *PrerequisiteSnapshot) error {
	if n.Type != cb.PrerequisiteNodeType_Message {
		return errors.New("unexpected PrerequisiteNodeType")
	}

	snapshot.Value[n.Id]++

	return nil
}

func (t *PrerequisiteTree) UpdateSnapshot(name string, snapshot *PrerequisiteSnapshot) error {
	if snapshot == nil {
		return errors.New("nil pointer PrerequisiteSnapshot")
	} else if len(snapshot.Value) != len(t.Pres) {
		return errors.New("prerequisite length not match")
	}

	if node, ok := t.Index[name]; ok {
		return node.UpdateSnapshot(snapshot)
	}

	return nil
}
