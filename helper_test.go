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


func TestGetDaysOldSeason(t *testing.T) {

	days := getDays("20151030")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days not returned.")
	}

} // TestGetDaysOldSeason


func TestGetDaysNow(t *testing.T) {

	now := time.Now()

	d := now.Format(DATE_FORMAT)

	days := getDays(d)

	t.Log(days)
	
	if len(days) == 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDaysNow


func TestGetDaysCurrentSeason(t *testing.T) {

	days := getDays("20191022")

	t.Log(days)
	
	if len(days) == 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDaysCurrentSeason



func TestGetDaysFuture(t *testing.T) {

	now := time.Now()

	future := now.AddDate(1, 0, 0)

	d := future.Format(DATE_FORMAT)

	days := getDays(d)

	t.Log(days)

	if len(days) > 0 {
		t.Error("Days should be empty.")
	}

} // TestGetDaysFuture


