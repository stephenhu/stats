package stats

import (
	"testing"
)


func TestParseWl(t *testing.T) {

	s := make(map[int]*Standing)

	s[1610612738] = &Standing{}
	s[1610612750] = &Standing{}

	g := NbaGame{
		Home: NbaTeamScore{
			ID: 1610612738,
			Score: 100,
			Statistics: NbaTeamData{
				Points: 100,
			},
		},
		Away: NbaTeamScore{
			ID: 1610612750,
			Score: 83,
			Statistics: NbaTeamData{
				Points: 83,
			},
		},
	}
	ParseWl(s, &g)

	t.Log(s[1610612750])

} // TestParseWl
