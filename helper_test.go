package stats

import (
	"fmt"
  "testing"
)


func TestGetCurrentSeason(t *testing.T) {

	y := GetCurrentSeason()
	t.Log(y)

} // TestGetCurrentSeason


func TestIsFutureGameTrue(t *testing.T) {

	if !IsFutureGame("10/24/2050 00:00:00") {
		t.Error("10/24/2050 00:00:00 game has not been played")
	}

} // TestIsFutureGameTrue


func TestIsFutureGameFalse(t *testing.T) {

	if IsFutureGame("10/24/2023 00:00:00") {
		t.Error("10/24/2023 00:00:00 game has already been played")
	}

} // TestIsFutureGameFalse


func TestUtcToFolder(t *testing.T) {

	f := UtcToFolder("2023-10-24T23:30:00Z")

	if f != "20231024" {
		t.Error(fmt.Sprintf("Received %s, should have received '20231024", f))
	}

} // TestUtcToFolder
