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


func TestInvalidApiInvoke(t *testing.T) {

	apiInvoke(BoxscoreApi("0032300327"), nil)

} // TestInvalidApiInvoke


func TestBeforeDate(t *testing.T) {

	if !compareDates("20231004", "20231005") {
		t.Error("20231004 should be before 20231005")
	}

} // TestBeforeDate


func TestSameDate(t *testing.T) {

	if !compareDates("20231004", "20231004") {
		t.Error("20231004 should be same 20231004")
	}
	
} // TestSameDate


func TestAfterDate(t *testing.T) {

	if compareDates("20231005", "20231004") {
		t.Error("20231005 should be before 20231004")
	}
	
} // TestAfterDate
