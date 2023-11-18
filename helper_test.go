package stats

import (
  "testing"
)


func TestGetCurrentSeason(t *testing.T) {

	y := GetCurrentSeason()
	t.Log(y)

} // TestGetCurrentSeason
