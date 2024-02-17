package reaction

import (
	"reflect"
	"testing"
)

func TestPrerequisiteTree_InitializeSnapshot(t *testing.T) {
	snapshot := Tree0.InitializeSnapshot()
	expect := &PrerequisiteSnapshot{Value: make([]int64, 3)}
	if !reflect.DeepEqual(snapshot, expect) {
		t.Error("fail")
	}
}

func TestPrerequisiteTree_UpdateSnapshot_0(t *testing.T) {
	snapshot := Tree0.InitializeSnapshot()
	Tree0.UpdateSnapshot("EventA", snapshot)

	acc, err := Tree0.Check(snapshot)
	if err != nil || acc {
		t.Error("fail, err:", err)
	}

	Tree0.UpdateSnapshot("EventB", snapshot)

	acc, err = Tree0.Check(snapshot)
	if err != nil || !acc {
		t.Error("fail, err:", err)
	}
}

func BenchmarkPrerequisiteTree_UpdateSnapshot_0(b *testing.B) {
	snapshot := Tree0.InitializeSnapshot()
	for i := 0; i < b.N; i++ {
		Tree0.UpdateSnapshot("EventA", snapshot)
		Tree0.UpdateSnapshot("EventB", snapshot)
	}
}
