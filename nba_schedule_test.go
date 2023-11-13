package stats

import (
	"testing"
)


func TestScheduleApi(t *testing.T) {

	api := ScheduleApi()

	if len(api) == 0 {
		t.Error("No schedule endpoint returned")
	}

} // TestScheduleApi


func TestNbaGetSchedule(t *testing.T) {

	s := NbaGetSchedule()

	t.Log(s)

	if s == nil {
		t.Error()
	}

} // TestNbaGetSchedule

