package pkg

import (
	"testing"
)

func TestMostCriticalStatus(t *testing.T) {
	service := Service{
		Name:   "Test",
		Status: "Partial Outage",
	}
	services := []Service{service}

	mostCritical := MostCriticalStatus(services)

	if mostCritical != 2 {
		t.Error("The most critical status should be 2!")
	}
}
