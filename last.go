package stats

import (
	//"fmt"
	//"log"
)


func LastPlayedGames(n int, p string) []Player {

	last := []Player{}

	now := GetEstNow()

	s := getSeason(*now)

	d := *now

	for {

		date := d.Format(DATE_FORMAT)

		if len(last) == n || s[SEASON_INDEX_BEGIN] > date {
			break
		}

		g, ok := PlayersMap[p][date]

		if ok {
			last = append(last, g)
		}

		d = d.AddDate(0, 0, -1)

	}

	return last

} // LastPlayedGames


func LastTeamGames(n int, t string) []Team {

	last := []Team{}

	now := GetEstNow()

	s := getSeason(*now)

	d := *now

	for {

		date := d.Format(DATE_FORMAT)

		if len(last) == n || s[SEASON_INDEX_BEGIN] > date {
			break
		}

		g, ok := TeamsMap[t][date]

		if ok {
			last = append(last, g)
		}

		d = d.AddDate(0, 0, -1)

  }

	return last

} // LastTeamGames
