package stats

import (
  "testing"
)

func TestNbaGetScoreboardToday(t *testing.T) {

	s := NbaGetScoreboardToday()

	t.Log(s)

	if s == nil {
		t.Error("Empty")
	}

} // TestNbaGetScoreboardToday
