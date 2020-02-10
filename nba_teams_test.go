package stats

import (
	"testing"
)


func TestNbaStoreTeams(t *testing.T) {

	teams := NbaGetTeams("2019")

	if teams == nil {
		t.Error("No teams returned")
	} else {
		NbaStoreTeams(teams)
	}

} // TestNbaStoreTeams


func TestNbaStoreTeamRanks(t *testing.T) {

	ranks := NbaGetTeamRanks("2019")

	if ranks == nil {
		t.Error("No teams returned")
	} else {
		NbaStoreTeamRanks(ranks)
	}

} // TestNbaStoreTeamRanks
