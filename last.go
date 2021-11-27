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

						if g.Home.Score > g.Away.Score {
							g.Home.Win = true
						} else {
							g.Home.Win = false
						}

						g.Home.Opponent = g.Away.Name
						g.Home.OpponentScore = g.Away.Score

						last = append(last, g.Home)

					} else {

						if g.Home.Score > g.Away.Score {
							g.Away.Win = false
						} else {
							g.Away.Win = true
						}

						g.Away.Opponent = g.Home.Name
						g.Away.OpponentScore = g.Home.Score

						last = append(last, g.Away)

					}

					if n != -1 && len(last) == n {
						break
					}

				}

			}

		}

	}

	rp.Close()

	return last

} // LastGamesTeam
