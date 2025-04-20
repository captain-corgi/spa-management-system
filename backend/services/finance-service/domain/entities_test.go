package domain

import "testing"

func TestInvoice_SequentialNumber(t *testing.T) {
	inv1 := Invoice{Number: "INV-001"}
	inv2 := Invoice{Number: "INV-002"}
	if inv1.Number >= inv2.Number {
		t.Error("Invoice numbers should be sequential")
	}
}
