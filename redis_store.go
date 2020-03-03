package stats

import (
	"encoding/json"
	"fmt"
	//"log"
	"sort"
	"strings"

	"github.com/gomodule/redigo/redis"
)

const (
	HEXISTS     = "HEXISTS"
	HGET				= "HGET"
	HSET				= "HSET"
	HGETALL     = "HGETALL"
	HKEYS       = "HKEYS"
)

var Red redis.Conn

var replacer = strings.NewReplacer(
	STRING_SPACE, STRING_EMPTY,
	STRING_SINGLE_QUOTE, STRING_EMPTY,
	STRING_PERIOD, STRING_EMPTY)


func KeyFilter(s string) string {
	return strings.ToLower(replacer.Replace(s))
} // KeyFilter


func ConnectRedis(protocol string, addr string) {

	c, err := redis.Dial(protocol, addr)

	if err != nil {
		logf("ConnectRedis", err.Error())
	} else {

		Red = c

		_, err := Red.Do("PING")

		if err != nil {
			logf("DEBUG", err.Error())
		}

	}

} // ConnectRedis


func RedisStorePlayers(s string) {

	players := NbaGetPlayers(s)

	if players != nil {

		all := convLeaguePlayers(players)

		for _, p := range all.Players {

			j, err := json.Marshal(p)

			if err != nil {
				logf("RedisStorePlayers", err.Error())
			} else {

				err := Red.Send(HSET, fmt.Sprintf("%s:players", s),
			    KeyFilter(fmt.Sprintf("%s%s", p.First, p.Last)), j)

				if err != nil {
					logf("RedisStorePlayers", err.Error())
				}

			}

		}

		err := Red.Flush()

		if err != nil {
			logf("RedisStorePlayers", err.Error())
		}

	}

} // RedisStorePlayers


func RedisStoreTeams(s string) {

	teams := NbaGetTeams(s)

	if teams != nil {

		all := convTeamInfo(teams)

		for _, t := range all.Teams {

			j, err := json.Marshal(t)

			if err != nil {
				logf("RedisStoreTeams", err.Error())
			} else {

				err := Red.Send(HSET, fmt.Sprintf("%s:teams", s),
			    strings.ToLower(t.Short), j)

				if err != nil {
					logf("RedisStoreTeams", err.Error())
				}

			}

		}

		err := Red.Flush()

		if err != nil {
			logf("RedisStoreTeams", err.Error())
		}

	}

} // RedisStoreTeams


func RedisStoreDay(d string) int {

	s := NbaGetScoreboard(d)

	if s == nil {
		return -1
	}

	scores := NbaGetBoxscores(s)

	if scores == nil {
		return -1
	}

	count := 0

	for _, score := range scores {

		game := ConvBoxscore(&score)

		j, err := json.Marshal(game)

		if err != nil {
			logf("RedisStoreDay", err.Error())
		} else {

			name := fmt.Sprintf("%s%s", game.Away.Name, game.Home.Name)

			err := Red.Send(HSET, d, name, j)

			if err != nil {
				logf("RedisStoreDay", err.Error())
			} else {

				gt := GetGameType(d)

				err := Red.Send(HSET, game.SeasonID, d, gt)

				if err != nil {
					logf("RedisStoreDay", err.Error())
					logf("RedisStoreDay", fmt.Sprintf("Unable to store %s %s", d, name))
				} else {
					count += 1
				}

			}

		}

	}

	err := Red.Flush()

	if err != nil {
		logf("RedisStoreDay", err.Error())
		count = 0
	}

	return count

} // RedisStoreDay


func RedisStoreGamesFrom(d string) int {

	days := GetDays(d)

	count := 0

	for _, day := range days {
		count += RedisStoreDay(day)
	}

	return count

} // RedisStoreGamesFrom


func RedisStoreSeason(s string) int {

	season, ok := OfficialSeasons[s]

	if ok {

		_, err := Red.Do(HSET, "seasons", s, "")

		if err != nil {
			return 0
		} else {
			return RedisStoreGamesFrom(season[SEASON_BEGIN])
		}

	} else {
		logf("RedisStoreSeason", fmt.Sprintf("Unknown season: %s", s))
		return 0
	}

} // RedisStoreSeason


func RedisGetDay(d string) []Game {

	games := []Game{}

	keys, err := redis.Strings(Red.Do(HKEYS, d))

	if err != nil {
		logf("RedisGetDay", err.Error())
	} else {

		for _, key := range keys {

			j, err := redis.String(Red.Do(HGET, d, key))

			if err != nil {
				logf("RedisGetDay", err.Error())
			} else {

				game := Game{}

				err := json.Unmarshal([]byte(j), &game)

				if err != nil {
					logf("RedisGetDay", err.Error())
				} else {
					games = append(games, game)
				}

			}

		}

	}

	return games

} // RedisGetDay


func RedisGameDays(s string) []string {

	_, ok := OfficialSeasons[s]

	if ok {

		keys, err := redis.Strings(Red.Do(HKEYS, s))

		if err != nil {
			logf("RedisGameDays", err.Error())
			return []string{}
		} else {
			return keys
		}

	} else {
		logf("RedisGameDays", fmt.Sprintf("Unknown season: %s", s))
		return []string{}
	}

} // RedisGameDays


func RedisLastGame() string {

	current := RedisCurrentSeason()

	keys := RedisGameDays(current)

	sort.Strings(keys)

	if len(keys) > 0 {
		return keys[len(keys)-1]
	} else {
		return ""
	}

} // RedisLastGame


func RedisGames(s string) []string {

	games := []string{}

	days := RedisGameDays(s)

	for _, day := range days {

		keys, err := redis.Strings(Red.Do(HKEYS, day))

		if err != nil {
			logf("RedisGames", err.Error())
		} else {

			for _, key := range keys {
				games = append(games, fmt.Sprintf("%s:%s", day, key))
			}

		}

	}

	return games

} // RedisGames


func RedisSeasons() []string {

	seasons, err := redis.Strings(Red.Do(HKEYS, "seasons"))

	if err != nil {
		logf("RedisSeasons", err.Error())
		return nil
	} else {
		return seasons
	}

} // RedisSeasons


func RedisCurrentSeason() string {

	seasons := RedisSeasons()

	sort.Strings(seasons)

	return seasons[len(seasons)-1]

} // RedisCurrentSeason


func RedisGameDayExists(d string) bool {

	season := SeasonKeyByDate(d)

	if season == "" {
		return false
	} else {

		ok, err := redis.Bool(Red.Do(HEXISTS, season, d))

		if err != nil {
			logf("RedisGameDayExists", err.Error())
			return false
		} else {

			if ok {
				return true
			} else {
				return false
			}

		}

	}

} // RedisGameDayExists


func RedisSeasonCheck(s string) []string {

	download := []string{}

	season, ok := OfficialSeasons[s]

	if ok {

		days := GetDays(season[SEASON_BEGIN])

		for _, d := range days {

			if !RedisGameDayExists(d) {
				download = append(download, d)
			}

		}

	}

	return download

} // RedisSeasonCheck


func RedisSeasonSync(s string) int {

	count := 0

	season, ok := OfficialSeasons[s]

	if ok {

		days := GetDays(season[SEASON_BEGIN])

		for _, d := range days {

			if !RedisGameDayExists(d) {
				count += RedisStoreDay(d)
			}

		}

	}

	return count

} // RedisSeasonSync
