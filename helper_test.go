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


func TestGetYearsFrom(t *testing.T) {

	years := getYearsFrom("2015")

	if years == nil {
		t.Error("Years are nil")
	} else {

		for _, y := range years {

			_, ok := official_seasons[y]

			if !ok {
				t.Errorf("%s non-official season", y)
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


func TestLatestScoreboardDate(t *testing.T) {

  d := LatestScoreboardDate()

	if d == "" {
		t.Error("Empty date returned")
	}

	t.Log(d)

} // TestLatestScoreboardDate


func TestLastDownload(t *testing.T) {

	d := LastDownload()

	t.Log(d)

} // TestLastDownload


func TestGetDays(t *testing.T) {

	days := getDays("20200220")

	t.Log(days)

	if len(days) == 0 {
		t.Error("Days should not be empty.")
	}

} // TestGetDays
