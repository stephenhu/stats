package stats

import (
	"testing"
)


func TestNbaGetTeams(t *testing.T) {

	teams := NbaGetTeams("2019")

	if teams == nil {
		t.Error("No teams returned")
	}

} // TestNbaGetTeams


func TestNbaGetTeamRanks(t *testing.T) {

	ranks := NbaGetTeamRanks("2019")

	if ranks == nil {
		t.Error("No teams returned")
	}

} // TestNbaGetTeamRanks
