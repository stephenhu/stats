package stats

import (
	"fmt"
  "testing"
)


func TestGetCurrentSeason(t *testing.T) {

	y := GetCurrentSeason()
	t.Log(y)

} // TestGetCurrentSeason


func TestIsFutureGame(t *testing.T) {

	t.Log(IsFutureGame("10/05/2023 00:00:00"))

} // TestIsFutureGame


func TestUtcToFolder(t *testing.T) {

	f := UtcToFolder("2023-10-24T23:30:00Z")

	if f != "20231024" {
		t.Error(fmt.Sprintf("Received %s, should have received '20231024", f))
	}

} // TestUtcToFolder
