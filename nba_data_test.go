package stats

import (
	"testing"
)


func TestNbaGetScoreboard(t *testing.T) {

	s := NbaGetScoreboard("20200201")

	if s == nil {
		t.Error("No scoreboard")
	} else {
		t.Log(*s)
	}

} // TestNbaGetScoreboard


func TestNbaGetScoreboardInvalidDate(t *testing.T) {

	s := NbaGetScoreboard("")

	if s != nil {
		t.Error("Invalid date should not return valid json.")
	}

} // TestNbaGetScoreboardInvalidDate


func TestNbaGetBoxscores(t *testing.T) {

	s := NbaGetScoreboard("20200201")

	if s == nil {
		t.Error("Invalid date should not return valid json.")
	}

	scores := NbaGetBoxscores(s)

	if len(scores) == 0 {
		t.Error("No scores found.")
	} else {
		t.Log(len(scores))
		t.Log(scores)
	}

} // TestNbaGetBoxscores
