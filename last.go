package stats

import (
	"encoding/json"
	//"fmt"
	//"log"
	"sort"

	"github.com/gomodule/redigo/redis"
)


func LastGamesPlayer(n int, y string, p string) []Player {

	last := []Player{}

	rp := RP.Get()

	games, err := redis.StringMap(rp.Do(HGETALL, Key(y, KeyFilter(p))))

	if err != nil {
		logf("LastGamesPlayer", err.Error())
	} else {

		keys := []string{}

		for key, _ := range games {
			keys = append(keys, key)
		}

		sort.Sort(sort.Reverse(sort.StringSlice(keys)))

		for _, key := range keys {

			player := Player{}

			err := json.Unmarshal([]byte(games[key]), &player)

			if err != nil {
				logf("LastGamesPlayer", err.Error())
			} else {

				if player.DnpReason == "" {
					last = append(last, player)
				}

				if len(last) == n {
					break
				}

			}

		}

	}

	rp.Close()

	return last

} // LastGamesPlayer


func LastGamesTeam(n int, y string, t string) []Team {

	// TODO: make sure n is not astronomical
	last := []Team{}

	rp := RP.Get()

	games, err := redis.StringMap(rp.Do(HGETALL, Key(y, KeyFilter(t))))

	if err != nil {
		logf("LastGamesTeam", err.Error())
	} else {

		keys := []string{}

		for key, _ := range games {
			keys = append(keys, key)
		}

		sort.Sort(sort.Reverse(sort.StringSlice(keys)))

		for _, key := range keys {

			game, err := redis.String(rp.Do(HGET, key, games[key]))

			if err != nil {
				logf("LastGamesTeam", err.Error())
			} else {

				g := Game{}

				err := json.Unmarshal([]byte(game), &g)

				if err != nil {
					logf("LastGamesTeam", err.Error())
				} else {

					if g.Home.Name == t {
						last = append(last, g.Home)
					} else {
						last = append(last, g.Away)
					}

					if len(last) == n {
						break
					}

				}

			}

		}

	}

	rp.Close()

	return last

} // LastGamesTeam


/*
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
*/

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
