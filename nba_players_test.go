package stats

import (
	"testing"
)


func TestNbaStorePlayers(t *testing.T) {

	lp := NbaGetPlayers("2016")

	if lp == nil {
		t.Error("No players returned")
	} else {
		NbaStorePlayers(lp)
	}

} // TestNbaStorePlayers


func TestNbaStoreProfiles(t *testing.T) {

	lp := NbaGetPlayers("2019")

	if lp == nil {
		t.Error("No players returned")
	} else {
		
		profiles := NbaGetProfiles("2019", lp)

		NbaStoreProfiles(profiles)

	}

} // TestNbaStoreProfiles
