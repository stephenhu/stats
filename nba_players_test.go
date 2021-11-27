package stats

import (
	"testing"
)


func TestNbaGetPlayers(t *testing.T) {

	lp := NbaGetPlayers("2016")

	if lp == nil {
		t.Error("No players returned")
	}

} // TestNbaGetPlayers


func TestNbaGetProfiles(t *testing.T) {

	lp := NbaGetPlayers("2019")

	if lp == nil {
		t.Error("No players returned")
	} else {

		profiles := NbaGetProfiles("2019", lp)

		if len(profiles) == 0 {
			t.Error("No profiles returned")
		}
	}

} // TestNbaGetProfiles
