package stats

import (
  "testing"
)


func TestBoxscoreApi(t *testing.T) {

	s := BoxscoreApi("20190620")

	t.Log(s)

} // TestBoxscoreApi


func TestNbaGetBoxscore(t *testing.T) {

	s := NbaGetBoxscore("0022300062")

	t.Log(s)

} // TestNbaGetBoxscore
