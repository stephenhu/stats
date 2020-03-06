package stats

import (
	"encoding/json"
	"fmt"
	//"log"
	"sort"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	HEXISTS     = "HEXISTS"
	HGET				= "HGET"
	HSET				= "HSET"
	HGETALL     = "HGETALL"
	HKEYS       = "HKEYS"
)

const (
	KEY_SEASONS							= "seasons"
)

var RP *redis.Pool

var replacer = strings.NewReplacer(
	STRING_SPACE, STRING_EMPTY,
	STRING_SINGLE_QUOTE, STRING_EMPTY,
	STRING_PERIOD, STRING_EMPTY)


func KeyFilter(s string) string {
	return strings.ToLower(replacer.Replace(s))
} // KeyFilter


func KeyName(name string) string {
	return KeyFilter(name)
} // KeyName


func Key(a string, b string) string {

	if a == "" || b == "" {
		return ""
	}

	return fmt.Sprintf("%s:%s", a, b)

} // Key


func DeKey(k string) []string {
	return strings.Split(k, STRING_COLON)
} // DeKey


func ConnectRedis(protocol string, addr string) {

	RP = &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {return redis.Dial(protocol, addr)},
	}

} // ConnectRedis


func RedisStorePlayers(s string) {

	players := NbaGetPlayers(s)

	if players != nil {

		all := convLeaguePlayers(players)

		rp := RP.Get()

		for _, p := range all.Players {

			j, err := json.Marshal(p)

			if err != nil {
				logf("RedisStorePlayers", err.Error())
			} else {

				err := rp.Send(HSET, fmt.Sprintf("%s:players", s),
			    KeyFilter(fmt.Sprintf("%s%s", p.First, p.Last)), j)

				if err != nil {
					logf("RedisStorePlayers", err.Error())
				}

			}

		}

		err := rp.Flush()

		if err != nil {
			logf("RedisStorePlayers", err.Error())
		}

		rp.Close()

	}

} // RedisStorePlayers


func RedisStoreTeams(s string) {

	teams := NbaGetTeams(s)

	if teams != nil {

		all := convTeamInfo(teams)

		rp := RP.Get()

		for _, t := range all.Teams {

			j, err := json.Marshal(t)

			if err != nil {
				logf("RedisStoreTeams", err.Error())
			} else {

				err := rp.Send(HSET, fmt.Sprintf("%s:teams", s),
			    strings.ToLower(t.Short), j)

				if err != nil {
					logf("RedisStoreTeams", err.Error())
				}

			}

		}

		err := rp.Flush()

		if err != nil {
			logf("RedisStoreTeams", err.Error())
		}

		rp.Close()

	}

} // RedisStoreTeams


func RedisStorePlayer(s string, d string, gname string, player Player) {

	j, err := json.Marshal(player)

	if err != nil {
		logf("RedisStorePlayer", err.Error())
	} else {

		rp := RP.Get()

		_, err := rp.Do(HSET, Key(s, KeyName(player.Name)), Key(d, gname), j)

		if err != nil {
			logf("RedisStorePlayer", err.Error())
		}

		rp.Close()

	}

} // RedisStorePlayer


func RedisStoreTeam(s string, d string, gname string, team Team) {

	rp := RP.Get()

	_, err := rp.Do(HSET, fmt.Sprintf("%s:%s", s, strings.ToLower(team.Name)),
    d, gname)

	if err != nil {
		logf("RedisStoreTeam", err.Error())
	} else {

		for _, player := range team.Players {

			RedisStorePlayer(s, d, gname, player)

		}

	}

	rp.Close()

} // RedisStoreTeam


func RedisStoreIndex(game Game, d string) {

	rp := RP.Get()

	_, err := rp.Do(HSET, KEY_SEASONS, game.SeasonID, time.Now().String())

	if err != nil {
		logf("RedisStoreIndex", err.Error())
	} else {

		gt := GetGameType(d)

		_, err := rp.Do(HSET, game.SeasonID, d, gt)

		if err != nil {
			logf("RedisStoreIndex", err.Error())
			logf("RedisStoreIndex", fmt.Sprintf("Unable to store game index: %s", d))
		} else {

			name := fmt.Sprintf("%s%s", strings.ToLower(game.Away.Name),
				strings.ToLower(game.Home.Name))

			RedisStoreTeam(game.SeasonID, d, name, game.Away)
			RedisStoreTeam(game.SeasonID, d, name, game.Home)

		}

	}

	rp.Close()

} // RedisStoreIndex


func RedisStoreGame(d string, gid string, game *Game) {

	j, err := json.Marshal(game)

	if err != nil {
		logf("RedisStoreGame", err.Error())
	} else {

		name := fmt.Sprintf("%s%s", game.Away.Name, game.Home.Name)

		rp := RP.Get()

		_, err := rp.Do(HSET, d, name, j)

		if err != nil {
			logf("RedisStoreGame", err.Error())
		} else {

			RedisStoreIndex(*game, d)

		}

		rp.Close()

	}

} // RedisStoreGame


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

		RedisStoreGame(d, game.ID, game)

		count += 1

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

		rp := RP.Get()

		_, err := rp.Do(HSET, "seasons", s, "")

		rp.Close()

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

	rp := RP.Get()

	keys, err := redis.Strings(rp.Do(HKEYS, d))

	if err != nil {
		logf("RedisGetDay", err.Error())
	} else {

		for _, key := range keys {

			j, err := redis.String(rp.Do(HGET, d, key))

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

	rp.Close()

	return games

} // RedisGetDay


func RedisGetGame(d string, id string) *Game {

	rp := RP.Get()

	j, err := redis.String(rp.Do(HGET, d, id))

	rp.Close()

	if err != nil {
		logf("RedisGetGame", err.Error())
		return nil
	} else {

		g := Game{}

		err := json.Unmarshal([]byte(j), &g)

		if err != nil {
			logf("RedisGetGame", err.Error())
			return nil
		} else {
			return &g
		}

	}

} // RedisGetGame


func RedisGameDays(s string) []string {

	_, ok := OfficialSeasons[s]

	if ok {

		rp := RP.Get()

		keys, err := redis.Strings(rp.Do(HKEYS, s))

		rp.Close()

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

	if current == "" {
		return ""
	}

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

	rp := RP.Get()

	for _, day := range days {

		keys, err := redis.Strings(rp.Do(HKEYS, day))

		if err != nil {
			logf("RedisGames", err.Error())
		} else {

			for _, key := range keys {
				games = append(games, fmt.Sprintf("%s:%s", day, key))
			}

		}

	}

	rp.Close()

	return games

} // RedisGames


func RedisSeasons() []string {

	ret := []string{}

	rp := RP.Get()

	seasons, err := redis.Strings(rp.Do(HKEYS, "seasons"))

	rp.Close()

	if err != nil {
		logf("RedisSeasons", err.Error())
		return ret
	} else {
		return seasons
	}

} // RedisSeasons


func RedisCurrentSeason() string {

	seasons := RedisSeasons()

	if len(seasons) > 0 {
		sort.Strings(seasons)
		return seasons[len(seasons)-1]
	} else {
		return CurrentSeason()
	}

} // RedisCurrentSeason


func RedisGameDayExists(d string) bool {

	season := SeasonKeyByDate(d)

	if season == "" {
		return false
	} else {

		rp := RP.Get()

		ok, err := redis.Bool(rp.Do(HEXISTS, season, d))

		rp.Close()

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
