package domain

import "testing"

func TestMetric_Timestamped(t *testing.T) {
	m := Metric{Timestamp: 1234567890}
	if m.Timestamp == 0 {
		t.Error("Metric must be timestamped")
	}
}
