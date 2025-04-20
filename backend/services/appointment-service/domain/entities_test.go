package domain

import "testing"

func TestAppointment_ValidTime(t *testing.T) {
	appt := Appointment{StartTime: 1000, EndTime: 2000}
	if appt.EndTime <= appt.StartTime {
		t.Error("EndTime should be after StartTime")
	}
}
