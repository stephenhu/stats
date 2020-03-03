package stats

import (
	//"fmt"
	//"log"
)


func LastPlayedGames(n int, p string) []Player {

	last := []Player{}

	now := GetEstNow()

	s := GetSeason(*now)

	d := *now

	for {

		date := d.Format(DATE_FORMAT)

		if len(last) == n || s[SEASON_BEGIN] > date {
			break
		}

		g, ok := PlayersMap[p][date]

		if ok {

			if date != s[SEASON_ALL_STAR_GAME] && g.DnpReason == "" {
				last = append(last, g)
			}

		}

		d = d.AddDate(0, 0, -1)

	}

	return last

} // LastPlayedGames


func LastTeamGames(n int, t string) []Team {

	last := []Team{}

	now := GetEstNow()

	s := GetSeason(*now)

	d := *now

	for {

		date := d.Format(DATE_FORMAT)

		if len(last) == n || s[SEASON_BEGIN] > date {
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


func LastScores() []Game {

	now := GetEstNow()

	s := GetSeason(*now)

	d := *now

	for {

		date := d.Format(DATE_FORMAT)

		if s[SEASON_BEGIN] > date {
			break
		}

		games, ok := GamesMap[date]

		if ok {
			return games
		}

		d = d.AddDate(0, 0, -1)

  }

	return GamesMap[s[SEASON_BEGIN]]

} // LastScores
