package stats

import (
	"testing"
)


func TestTGameType(t *testing.T) {

	s := NbaGetSchedule()

	gt := TGameType(s.LeagueSchedule.GameDates)

	t.Log(gt)

} // TestTGameType
