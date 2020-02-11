package stats

import (
	"testing"
)


func TestLoadSeasons(t *testing.T) {

	seasons := loadSeasons()

	t.Log(seasons)

	if len(seasons) == 0 {
		t.Error("Seasons should not be empty.")
	}

} // TestLoadSeasons


func TestLoadTeams(t *testing.T) {

	LoadTeams()

	t.Log(teams_map)

	if len(teams_map) == 0 {
		t.Error("teams_map should not be empty.")
	}

} // TestLoadTeams
