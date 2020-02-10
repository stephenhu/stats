package stats

import (
	"testing"
	"time"
)


func TestGetSeason(t *testing.T) {

	season := getSeason(time.Now())

	if len(season) == 0 {
		t.Error("No season found.")
	}

	t.Log(season)

} // TestGetSeason


func TestGetDays(t *testing.T) {

	days := getDays("20151030")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days not returned.")
	}

}
