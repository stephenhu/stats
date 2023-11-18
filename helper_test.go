package stats

import (
  "testing"
)


func TestGetCurrentSeason(t *testing.T) {

	y := GetCurrentSeason()
	t.Log(y)

} // TestGetCurrentSeason


func TestIsFutureGame(t *testing.T) {

	t.Log(IsFutureGame("10/05/2023 00:00:00"))

} // TestIsFutureGame
