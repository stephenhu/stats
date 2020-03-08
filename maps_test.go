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

	t.Skip("no more file system storage")

	//LoadCache()

	//LastPlayedGames(5, "Kawhi Leonard")
	//t.Log(SeasonsMap)

} // TestLoadCache


func TestLoadTeams(t *testing.T) {

	t.Skip("load teams changed")
	//LoadTeams()

/*
	t.Log(teams_map)

	if len(teams_map) == 0 {
		t.Error("teams_map should not be empty.")
	}
*/
} // TestLoadTeams
