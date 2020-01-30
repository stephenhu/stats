package stats

import (
	"testing"
)


func TestGetGames(t *testing.T) {

	games := GetGames("")

	if games == nil {
		t.Error("No games found.")
	}

	s := GetStats(games)

	if s == nil {
		t.Error("No stats found.")
	}

} // testGetGames
