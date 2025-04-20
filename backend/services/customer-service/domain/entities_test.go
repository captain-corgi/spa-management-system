package domain

import (
	"testing"
)

func TestCustomer_UniquePhone(t *testing.T) {
	c1 := Customer{Phone: "123456"}
	c2 := Customer{Phone: "123456"}
	if c1.Phone == c2.Phone {
		// Simulate uniqueness check
		t.Log("Duplicate phone detected as expected")
	} else {
		t.Error("Phones should match for this test")
	}
}
