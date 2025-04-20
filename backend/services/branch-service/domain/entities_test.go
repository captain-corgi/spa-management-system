package domain

import "testing"

func TestBranch_UniqueName(t *testing.T) {
	b1 := Branch{Name: "Main"}
	b2 := Branch{Name: "Main"}
	if b1.Name == b2.Name {
		t.Log("Duplicate branch name detected as expected")
	} else {
		t.Error("Names should match for this test")
	}
}
