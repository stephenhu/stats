package stats

import (
	"testing"
)


func TestFindSeasons(t *testing.T) {

	seasons := findSeasons()

	t.Log(seasons)

	if len(seasons) == 0 {
		t.Error("Seasons should not be empty.")
	}

} // TestFindSeasons


func TestLoadCache(t *testing.T) {

	LoadCache()

	LastPlayerGames(5, "Kawhi Leonard")
	//t.Log(SeasonsMap)

} // TestLoadCache


func TestLoadTeams(t *testing.T) {

	LoadTeams()

/*
	t.Log(teams_map)

	if len(teams_map) == 0 {
		t.Error("teams_map should not be empty.")
	}
*/
} // TestLoadTeams
