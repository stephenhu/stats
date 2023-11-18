package stats

import (
  "testing"
)


func TestGetSeasonNow(t *testing.T) {

	y := getSeasonNow()
	t.Log(y)

} // TestGetSeasonNow
