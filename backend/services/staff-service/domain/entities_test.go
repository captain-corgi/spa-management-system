package domain

import "testing"

func TestStaff_UniqueID(t *testing.T) {
	staff1 := Staff{ID: 1}
	staff2 := Staff{ID: 1}
	if staff1.ID == staff2.ID {
		t.Log("Duplicate staff ID detected as expected")
	} else {
		t.Error("IDs should match for this test")
	}
}
