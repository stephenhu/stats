package stats

import (
	"log"
	"testing"
)


func TestScheduleApi(t *testing.T) {

	api := ScheduleApi("2019")

	log.Println(api)

	if len(api) == 0 {
		t.Error("No schedule endpoint returned")
	}

} // TestScheduleApi


func TestGetSchedule(t *testing.T) {

	NbaGetSchedule("2019")

} // TestGetSchedule
