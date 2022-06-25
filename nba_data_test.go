package stats

import (
	"testing"
)


func TestNbaGetScoreboardInvalidDate(t *testing.T) {

	s := NbaGetScoreboard("20190620")

	t.Log(s)
	if s != nil {
		t.Error("Invalid date, this day should have no games")
	}

} // TestNbaGetScoreboardInvalidDate


func TestNbaGetScoreboardEmptyDate(t *testing.T) {

	s := NbaGetScoreboard("")

	if s != nil {
		t.Error("Invalid date should not return valid json.")
	}

} // TestNbaGetScoreboardEmptyDate


func TestNbaGetBoxscores(t *testing.T) {

	s := NbaGetScoreboard("20181017")

	if s == nil {
		t.Error("This date has game data.")
	}

	scores := NbaGetBoxscores(s)

	if len(scores) == 0 {
		t.Error("No scores found.")
	} else {
		t.Log(len(scores))
		t.Log(scores)
	}

} // TestNbaGetBoxscores
