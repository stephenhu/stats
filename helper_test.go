package stats

import (
	"fmt"
  "testing"
	"time"
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


func TestUtcToString(t *testing.T) {

	f := UtcToString("2023-10-24T23:30:00Z")

	if f != "20231024" {
		t.Error(fmt.Sprintf("Received %s, should have received '20231024", f))
	}

} // TestUtcToString


func TestGameDateToString(t *testing.T) {

	f := GameDateToString("10/24/2023 23:30:00")

	if f != "20231024" {
		t.Error(fmt.Sprintf("Received %s, should have received '20231024", f))
	}

} // TestGameDateToString


func TestInvalidApiInvoke(t *testing.T) {

	apiInvoke(BoxscoreApi("0032300327"), nil)

} // TestInvalidApiInvoke


func TestAfterDate(t *testing.T) {

	if !IsAfterDate("20231004", "20231005") {
		t.Error("20231004 should be after 20231005")
	}

} // TestBeforeDate


func TestSameDate(t *testing.T) {

	if !IsAfterDate("20231004", "20231004") {
		t.Error("20231004 should be same 20231004")
	}
	
} // TestSameDate


func TestBeforeDate(t *testing.T) {

	if IsAfterDate("20231005", "20231004") {
		t.Error("20231005 should be before 20231004")
	}
	
} // TestBeforeDate

func TestS(t *testing.T) {

	a, err := time.Parse(DATE_FORMAT, "")

	if err != nil {
		t.Error(err)
	} else {
		t.Log(a)
	}
	
} // TestS


func TestPtm0(t *testing.T) {

	t.Log(PtmToMin("PT00M"))

} // TestPtm0
