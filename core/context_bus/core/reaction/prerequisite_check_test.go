package reaction

import "testing"

func TestPrerequisiteTree_Check_1(t *testing.T) {
	snapshot := Tree1.InitializeSnapshot()

	// only the 2nd and 3rd invocations of EventC lead to true values
	expected := []bool{false, true, true, false}
	for _, exp := range expected {
		Tree1.UpdateSnapshot("EventC", snapshot)
		acc, err := Tree1.Check(snapshot)
		if err != nil || acc != exp {
			t.Error("fail, err:", err, "acc:", acc)
		}
	}

	Tree1.UpdateSnapshot("EventA", snapshot)
	acc, err := Tree1.Check(snapshot)
	if err != nil || acc {
		t.Error("fail, err:", err, "acc:", acc)
	}
}

func BenchmarkPrerequisiteTree_Check_0(b *testing.B) {
	snapshot := Tree0.InitializeSnapshot()
	Tree0.UpdateSnapshot("EventA", snapshot)
	Tree0.UpdateSnapshot("EventB", snapshot)

	for i := 0; i < b.N; i++ {
		Tree0.Check(snapshot)
	}
}
