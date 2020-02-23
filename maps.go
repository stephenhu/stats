package stats

import (
	"encoding/json"
	"fmt"
	//"log"
	"os"
	"path/filepath"
	"strings"
)

var GamesMap					= map[string] []Game{}
var TeamInfoMap 			= map[string] TeamInfo{}

// player name, date
var PlayersMap  			= map[string] map[string]Player{}

// team name, date
var TeamsMap          = map[string] map[string]Team{}


func teamKey(s string, id string) string {
	return fmt.Sprintf("%s.%s", s, id)
} // teamKey


func gameKey(d string, away string, home string) string {
	return fmt.Sprintf("%s.%s.%s", d, away, home)
} // gameKey


func loadPlayers(game Game, home bool) {

	t := game.Home

	if !home {
		t = game.Away
	}

	for _, p := range t.Players {

		if home {
			p.Opponent = game.Away.Name
		} else {
			p.Opponent = game.Home.Name
		}


		if PlayersMap[p.Name] == nil {
			PlayersMap[p.Name] = map[string]Player{}
		}

		PlayersMap[p.Name][game.Date] = p

	}

} // loadPlayers


func loadGames(game Game) {

	loadPlayers(game, true)
	loadPlayers(game, false)
	loadTeams(game)

} // loadPlayers


func findSeasons() []os.FileInfo {
	return getDirs("")
} // findSeasons


func parseSeason(s string) {

	days := getDirs(s)

	for _, day := range days {

		files := getFiles(s, day)

		games := []Game{}

		for _, f := range files {

			buf := loadFile(filepath.Join(APP_STORAGE, s, day.Name(), f.Name()))

			g := Game{}

			err := json.Unmarshal(buf, &g)

			if err != nil {
				logf("parseSeason", err.Error())
			} else {
				games = append(games, g)
				loadGames(g)
			}

		}

		GamesMap[day.Name()] = games

	}

} // parseSeason


func loadTeams(game Game) {

	if TeamsMap[strings.ToLower(game.Away.Name)] == nil {
		TeamsMap[strings.ToLower(game.Away.Name)] = map[string]Team{}
	}

	if TeamsMap[strings.ToLower(game.Home.Name)] == nil {
		TeamsMap[strings.ToLower(game.Home.Name)] = map[string]Team{}
	}

	TeamsMap[strings.ToLower(game.Away.Name)][game.Date] = game.Away
	TeamsMap[strings.ToLower(game.Home.Name)][game.Date] = game.Home

} // LoadTeams


func LoadCache() {

	if checkStorage() {

		now := GetEstNow()

		sk := SeasonKeyByDate(now.Format(DATE_FORMAT))

		parseSeason(sk)

	}

} // LoadCache
