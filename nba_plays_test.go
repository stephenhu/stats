package stats

import (
	"testing"
)


func TestNbaGetPlays(t *testing.T) {

	plays := NbaGetGamePlays("20190206", "0021800804", 4)

	if plays == nil {
		t.Error("Plays are nil")
	} else {
		t.Logf("%+v", plays)
	}

} // TestNbaGetPlays

func TestNbaGetPlaysWithOvertime(t *testing.T) {

	plays := NbaGetGamePlays("20190504", "0041800223", 5)

	if plays == nil {
		t.Error("Plays are nil")
	} else {
		t.Log(plays)
	}

} // TestNbaGetPlaysWithOvertime
