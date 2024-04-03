package stats

import (
	"testing"
)


func TestNbaGetLegacySchedule(t *testing.T) {

	years := []string{
		"2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022"}

	for _, y := range years {

		s := NbaGetLegacySchedule(y)
		t.Log(s)
	
	}

} // TestNbaGetLegacySchedule
