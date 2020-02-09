package stats

import (
	"testing"
)


func TestNbaStoreTeams(t *testing.T) {

	teams := NbaGetTeams("2019")

	if teams == nil {
		t.Error("No players returned")
	} else {
		NbaStoreTeams(teams)
	}

} // TestNbaStoreTeams
