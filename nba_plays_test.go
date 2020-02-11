package stats

import (
	"testing"
)


func TestNbaStorePlays(t *testing.T) {

	plays := NbaGetPlays("20190206", "0021800804")

	if plays == nil {
		t.Error("Plays are nil")
	} else {

		t.Logf("%+v", plays)

	}

} // TestNbaStorePlays
