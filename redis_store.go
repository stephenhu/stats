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

			j, err := json.MarshalIndent(p, JSON_PREFIX, JSON_INDENT)

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

			j, err := json.MarshalIndent(t, JSON_PREFIX, JSON_INDENT)

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

	j, err := json.MarshalIndent(player, JSON_PREFIX, JSON_INDENT)

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

	j, err := json.MarshalIndent(game, JSON_PREFIX, JSON_INDENT)

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


func RedisStorePlayerInfo(s string) {

	players := NbaGetPlayers(s)

	all := convLeaguePlayers(players)

	rp := RP.Get()

	for _, player := range all.Players {

		j, err := json.MarshalIndent(player, JSON_PREFIX, JSON_INDENT)

		if err != nil {
			logf("RedisStorePlayerInfo", err.Error())
		} else {

			_, err := rp.Do(HSET, fmt.Sprintf("%s:players", s), KeyName(fmt.Sprintf("%s%s",
				player.First, player.Last)), j)

			if err != nil {
				logf("RedisStorePlayerInfo", err.Error())
			}

		}

	}

	rp.Close()

} // RedisStorePlayerInfo


func RedisStoreProfiles(s string) {

	players := NbaGetPlayers(s)

	all := NbaGetProfiles(s, players)

	rp := RP.Get()

	for _, player := range all {

		profile := convPlayerProfile(&player)

		j, err := json.MarshalIndent(profile, JSON_PREFIX, JSON_INDENT)

		if err != nil {
			logf("RedisStoreProfiles", err.Error())
		} else {

			_, err := rp.Do(HSET, fmt.Sprintf("%s:players:stats", s), KeyName(fmt.Sprintf(
				"%s%s", profile.First, profile.Last)), j)

			if err != nil {
				logf("RedisStoreProfiles", err.Error())
			}

		}

	}

	rp.Close()

} // RedisStoreProfiles


func RedisStoreTeamRosters(s string) {

	teams := NbaGetTeams(s)

	if teams != nil {

		rp := RP.Get()

		for _, team := range teams.Teams {

			mascot, ok := RosterTeams[strings.ToLower(team.Code)]

			if ok {

				roster := NbaGetTeamRoster(s, mascot)

				ros := convRoster(roster)

				if ros != nil {

					j, err := json.MarshalIndent(ros, JSON_PREFIX, JSON_INDENT)

					if err != nil {
						logf("RedisStoreTeamRosters", err.Error())
					} else {

						_, err := rp.Do(HSET, fmt.Sprintf("%s:teams:rosters", s), strings.ToLower(team.Code), j)

						if err != nil {
							logf("RedisStoreTeamRosters", err.Error())
						}

					}

				} else {
					logf("RedisStoreTeamRosters", "empty roster")
				}

			} else {
				logf("RedisStoreTeamRosters", fmt.Sprintf("Team %s not found.", team.Code))
			}

		}

		rp.Close()

	}

} // RedisStoreTeamRosters


func RedisStoreTeamInfo(s string) {

	teams := NbaGetTeams(s)

	all := convTeamInfo(teams)

	rp := RP.Get()

	for _, team := range all.Teams {

		j, err := json.MarshalIndent(team, JSON_PREFIX, JSON_INDENT)

		if err != nil {
			logf("RedisStoreTeamInfo", err.Error())
		} else {

			rp.Do(HSET, fmt.Sprintf("%s:teams", s), KeyName(team.Code), j)

		}

	}

	rp.Close()

} // RedisStoreTeamInfo


func RedisStoreTeamRanks(s string) {

	all := NbaGetTeamRanks(s)

	ranks := convTeamRanks(all)

	if ranks == nil {
		return
	}

	rp := RP.Get()

	for _, rank := range ranks.Teams {

		j, err := json.Marshal(rank)

		if err != nil {
			logf("RedisStoreTeamRanks", err.Error())
		} else {

			rp.Do(HSET, fmt.Sprintf("%s:teams:stats", s), KeyName(rank.Name), j)

		}

	}

	rp.Close()

} // RedisStoreTeamRanks


func RedisStoreTeamStandings(s string) {

	standings := NbaGetTeamStandings(s)

	if standings == nil {
		logf("RedisStoreTeamStandings", "Unable to retrieve standings.")
	} else {

		ns := convStandings(standings)

		rp := RP.Get()

		for _, record := range ns.Records {

			j, err := json.Marshal(record)

			if err != nil {
				logf("RedisStoreTeamStandings", err.Error())
			} else {

				_, err := rp.Do(HSET, fmt.Sprintf("%s:standings", s), record.Name, j)

				if err != nil {
					logf("RedisStoreTeamStandings", err.Error())
				}

			}

		}

		rp.Close()

	}

} // RedisStoreTeamStandings


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


func RedisGetPlayer(s string, name string) *PlayerCareer {

	if name == "" {
		return nil
	}

	rp := RP.Get()

	j, err := redis.String(rp.Do(HGET, fmt.Sprintf("%s:players:stats", s), name))

	rp.Close()

	if err != nil {
		logf("RedisGetPlayer", err.Error())
		return nil
	} else {

		p := PlayerCareer{}

		err := json.Unmarshal([]byte(j), &p)

		games := LastGamesPlayer(10, s, name)

		p.LastGames = games

		if err != nil {
			logf("RedisGetPlayer", err.Error())
			return nil
		} else {
			return &p
		}

	}

} // RedisGetPlayer


func RedisGetPlayerInfo(s string, name string) *PlayerInfo {

	if name == "" {
		return nil
	}

	rp := RP.Get()

	j, err := redis.String(rp.Do(HGET, fmt.Sprintf("%s:players", s), KeyName(name)))

	rp.Close()

	if err != nil {
		logf("RedisGetPlayer", err.Error())
		return nil
	} else {

		p := PlayerInfo{}

		err := json.Unmarshal([]byte(j), &p)

		if err != nil {
			logf("RedisGetPlayer", err.Error())
			return nil
		} else {
			return &p
		}

	}

} // RedisGetPlayerInfo


func RedisGetAllPlayerInfo(s string) map[string]PlayerInfo {

	mpi := map[string] PlayerInfo{}

	rp := RP.Get()

	all, err := redis.StringMap(rp.Do(HGETALL, fmt.Sprintf("%s:players", s)))

	if err != nil {
		logf("RedisGetAllPlayerInfo", err.Error())
	} else {

		for _, v := range all {

			pi := PlayerInfo{}

			err := json.Unmarshal([]byte(v), &pi)

			if err != nil {
				logf("RedisGetAllPlayerInfo", err.Error())
			} else {
				mpi[pi.ID] = pi
			}

		}

	}

	rp.Close()

	return mpi

} // RedisGetAllPlayerInfo


func RedisGetTeamRanks(s string, name string) *TeamRanks {

	if name == "" {
		return nil
	}

	rp := RP.Get()

	j, err := redis.String(rp.Do(HGET, fmt.Sprintf("%s:teams:stats", s), name))

	rp.Close()

	if err != nil {
		logf("RedisGetTeamRanks", err.Error())
		return nil
	} else {

		tr := TeamRanks{}

		err := json.Unmarshal([]byte(j), &tr)

		if err != nil {
			logf("RedisGetTeamRanks", err.Error())
			return nil
		} else {
			return &tr
		}

	}

} // RedisGetTeamRanks


func RedisGetTeamRoster(s string, name string) *Roster {

	if s == "" || name == "" {
		return nil
	}

	rp := RP.Get()

	j, err := redis.String(rp.Do(HGET, fmt.Sprintf("%s:teams:rosters", s), name))

	rp.Close()

	if err != nil {
		logf("RedisGetTeamRoster", err.Error())
		return nil
	} else {

		r := Roster{}

		err := json.Unmarshal([]byte(j), &r)

		if err != nil {
			logf("RedisGetTeamRoster", err.Error())
			return nil
		} else {
			return &r
		}

	}

} // RedisGetTeamRoster


func RedisGetTeamRosterStats(s string, name string) []PlayerSeason {

	all := []PlayerSeason{}

	roster := RedisGetTeamRoster(s, name)

	if roster != nil {

		mpi := RedisGetAllPlayerInfo(s)

		rp := RP.Get()

		for _, player := range roster.Players {

			pi, ok := mpi[player]

			if ok {

				val, err := redis.String(rp.Do(HGET, fmt.Sprintf("%s:players:stats", s), KeyName(fmt.Sprintf(
					"%s%s", pi.First, pi.Last))))

					if err != nil {
						logf("RedisGetTeamRosterStats", err.Error())
					} else {

						pc := PlayerCareer{}

						err := json.Unmarshal([]byte(val), &pc)

						if err != nil {
							logf("RedisGetTeamRosterStats", err.Error())
						} else {

							ps := convCareerSeason(&pc)

							ps.ID 			= pi.ID
							ps.First		= pi.First
							ps.Last			= pi.Last

							all = append(all, *ps)

						}

					}

			} else {
				logf("RedisGetTeamRosterStats", fmt.Sprintf("player id %s not found", player))
			}

		}

		rp.Close()

	}

	return all

} // RedisGetTeamRosterStats


func RedisGetTeamStandings(s string) *Standings {

	if s == "" {
		return nil
	}

	rp := RP.Get()

	m, err := redis.StringMap(rp.Do(HGETALL, fmt.Sprintf("%s:standings", s)))

	rp.Close()

	if err != nil {
		logf("RedisGetTeamStandings", err.Error())
		return nil
	} else {

		ns := Standings{
			Records: make(map[string]TeamRecord),
		}

		for key, val := range m {

			tr := TeamRecord{}

			err := json.Unmarshal([]byte(val), &tr)

			if err != nil {
				logf("RedisGetTeamStandings", err.Error())
			} else {
				ns.Records[key] = tr
			}

		}

		return &ns

	}

} // RedisGetTeamStandings


func RedisGetTeamData(s string, n string) *TeamData {

	td := TeamData{}

	ranks 			:= RedisGetTeamRanks(s, n)
	players 		:= RedisGetTeamRosterStats(s, n)
	standings 	:= RedisGetTeamStandings(s)
	last        := LastGamesTeam(10, s, n)


	td.SeasonID   = s
	td.Name				= n
	td.Ranks 			= *ranks
	td.Players		= players
	td.Standings  = *standings
	td.Games      = last

	return &td

} // RedisGetTeamDta


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
