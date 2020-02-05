package stats

import (
	"testing"
)


func TestGetGameIDsEmptyDate(t *testing.T) {

	t.Skip("deprecate in favor of data.nba.net")

	ids := GetGameIDs("")

	if ids == nil {
		t.Error("No ids found.")
	}

	g := GetGames(ids)

	if g == nil {
		t.Error("No games found.")
	}

} // TestGetGameIDsEmptyDate


func TestGetGameIDsWithDate(t *testing.T) {

	t.Skip("deprecate in favor of data.nba.net")

  ids := GetGameIDs("20200128")

	if ids == nil {
		t.Error("No ids found.")
	}

	g := GetGames(ids)

	if g == nil {
		t.Error("No games found.")
	}

} // TestGetGameIDsWithDate


func TestGetGameIDsFrom(t *testing.T) {

	t.Skip("deprecate in favor of data.nba.net")

	ids := GetGameIDsFrom("20200128")

	if len(ids) == 0 {
		t.Error("No ids found.")
	}

	t.Log(ids)

	g := GetGames(ids)

	if g == nil {
		t.Error("No games found.")
	}

	t.Log(g)
	
} // TestGetGameIDsFrom


func TestGetSeasonGameIDs(t *testing.T) {

	t.Skip("deprecate in favor of data.nba.net")
	
	ids := GetGameIDsBySeason("1920")

	if len(ids) == 0 {
		t.Error("No ids found.")
	}

	t.Log(ids)

	/*
	g := GetGames(ids)

	if g == nil {
		t.Error("No games found.")
	}

	t.Log(g)
	*/
	
} // TestGetSeasonGameIDs
