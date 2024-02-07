package reaction

import (
	"reflect"
	"testing"
)

func TestPrerequisiteTree_InitializeSnapshot(t *testing.T) {
	snapshot := tree0.InitializeSnapshot()
	expect := &PrerequisiteSnapshot{Value: make([]int64, 3)}
	if !reflect.DeepEqual(snapshot, expect) {
		t.Error("fail")
	}
}

func TestPrerequisiteTree_UpdateSnapshot_0(t *testing.T) {
	snapshot := tree0.InitializeSnapshot()
	tree0.UpdateSnapshot("EventA", snapshot)

	acc, err := tree0.Check(snapshot)
	if err != nil || acc {
		t.Error("fail, err:", err)
	}

	tree0.UpdateSnapshot("EventB", snapshot)

	acc, err = tree0.Check(snapshot)
	if err != nil || !acc {
		t.Error("fail, err:", err)
	}
}

func BenchmarkPrerequisiteTree_UpdateSnapshot_0(b *testing.B) {
	snapshot := tree0.InitializeSnapshot()
	for i := 0; i < b.N; i++ {
		tree0.UpdateSnapshot("EventA", snapshot)
		tree0.UpdateSnapshot("EventB", snapshot)
	}
}
