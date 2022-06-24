package stats

import (
	"testing"
	"time"
)


func TestGetSeason(t *testing.T) {

	season := GetSeason(time.Now())

	if len(season) == 0 {
		t.Error("No season found.")
	}

	t.Log(season)

} // TestGetSeason


func TestGetDaysOldSeason(t *testing.T) {

	days := GetDays("20151030")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days not returned.")
	}

} // TestGetDaysOldSeason


func TestGetDaysPastSeason(t *testing.T) {

	days := GetDays("20170701")

	t.Log(days)

	if len(days) > 0 {
		t.Error("Days should be empty.")
	}

} // TestGetDaysPastSeason


func TestGetDaysBeforeSeason(t *testing.T) {

	days := GetDays("20191022")

	t.Log(days)

	if len(days) > 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDaysBeforeSeason


func TestGetDaysFuture(t *testing.T) {

	now := time.Now()

	future := now.AddDate(1, 0, 0)

	d := future.Format(DATE_FORMAT)

	days := GetDays(d)

	t.Log(days)

	if len(days) > 0 {
		t.Error("Days should be empty.")
	}

} // TestGetDaysFuture


func TestGetYearsFrom(t *testing.T) {

	years := GetYearsFrom(2015)

	if years == nil {
		t.Error("Years are nil")
	} else {

		for _, y := range years {

			_, ok := OfficialSeasons[y]

			if !ok {
				t.Errorf("%d non-official season", y)
			}

		}

	}

} // TestGetYearsFrom


func TestGetEstNow(t *testing.T) {

	now := GetEstNow()

	if now == nil {
		t.Error("Return nil")
	}

	t.Log(now)

} // TestGetEstNow


func TestGetDays(t *testing.T) {

	days := GetDays("20200220")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDays


func TestGetDays2(t *testing.T) {

	days := GetDays("20191120")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDays2


func TestGetDays3(t *testing.T) {

	days := GetDays("20181120")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDays3
