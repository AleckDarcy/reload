package reaction

import "testing"

func TestPrerequisiteTree_Check_1(t *testing.T) {
	snapshot := tree1.InitializeSnapshot()

	// only the 2nd and 3rd invocations of EventC lead to true values
	expected := []bool{false, true, true, false}
	for _, exp := range expected {
		tree1.UpdateSnapshot("EventC", snapshot)
		acc, err := tree1.Check(snapshot)
		if err != nil || acc != exp {
			t.Error("fail, err:", err)
		}
	}

	tree1.UpdateSnapshot("EventA", snapshot)
	acc, err := tree1.Check(snapshot)
	if err != nil || acc {
		t.Error("fail, err:", err)
	}
}

func BenchmarkPrerequisiteTree_Check_0(b *testing.B) {
	snapshot := tree0.InitializeSnapshot()
	tree0.UpdateSnapshot("EventA", snapshot)
	tree0.UpdateSnapshot("EventB", snapshot)

	for i := 0; i < b.N; i++ {
		tree0.Check(snapshot)
	}
}
