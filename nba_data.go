package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

type NbaInternal struct {
	Date					string						`json:"pubDateTime"`
}

type NbaAdvStats struct {
	SeasonID					int						`json:"seasonYear"`
	TeamID						string        `json:"teamId"`
	Minutes						string				`json:"min"`
	Points						string				`json:"points"`
	Oreb							string				`json:"offReb"`
	Dreb							string				`json:"defReb"`
	Treb							string				`json:"totReb"`
	Assists						string				`json:"assists"`	
	Turnovers					string				`json:"turnovers"`
	Steals						string				`json:"steals"`
	Blocks						string				`json:"blocks"`
	Fouls							string				`json:"pFouls"`
	Fgm               string        `json:"fgm"`
	Fga               string        `json:"fga"`
	Fg3m              string        `json:"tpm"`
	Fg3a              string        `json:"tpa"`
	Ftm               string        `json:"ftm"`
	Fta               string        `json:"fta"`
	Played						string				`json:"gamesPlayed"`
	Started						string				`json:"gamesStarted"`
	PlusMinus					string				`json:"plusMinus"`
	Ppg								string				`json:"ppg"`			// points per game
	Rpg								string				`json:"rpg"`			// rebounds per game
	Apg								string				`json:"apg"`			// assists per game
	Mpg								string				`json:"mpg"`			// minutes per game
	Topg							string				`json:"topg"`			// turnovers per game
	Spg								string				`json:"spg"`			// steals per game
	Bpg								string				`json:"bpg"`			// blocks per game
	Fgp								string				`json:"fgp"`			// fg %
	Fg3p							string				`json:"tpp"`			// three point %
	Ftp								string				`json:"ftp"`			// free throw %
}

type NbaTeamData struct {
	Points        string           	`json:"points"`
	Minutes       string        		`json:"min"`
	Fgm        		string           	`json:"fgm"`
	Fga        		string           	`json:"fga"`
	Ftm        		string           	`json:"ftm"`
	Fta        		string           	`json:"fta"`
	Fg3m        	string           	`json:"tpm"`
	Fg3a        	string           	`json:"tpa"`
	Oreb        	string           	`json:"offReb"`
	Dreb        	string           	`json:"defReb"`
	Treb        	string           	`json:"totReb"`
	Assists       string           	`json:"assists"`
	Fouls        	string           	`json:"pFouls"`
	Steals        string           	`json:"steals"`
	Turnovers     string           	`json:"turnovers"`
	Blocks        string           	`json:"blocks"`
	PlusMinus     string           	`json:"plusMinus"`
}

type NbaPlayerData struct {
	ID						string						`json:"personId"`
	TeamID        string            `json:"teamId"`
	First         string        		`json:"firstName"`
	Last          string        		`json:"lastName"`
	Points        string           	`json:"points"`
	Minutes       string        		`json:"min"`
	Fgm        		string           	`json:"fgm"`
	Fga        		string           	`json:"fga"`
	Ftm        		string           	`json:"ftm"`
	Fta        		string           	`json:"fta"`
	Fg3m        	string          	`json:"tpm"`
	Fg3a        	string          	`json:"tpa"`
	Oreb        	string           	`json:"offReb"`
	Dreb        	string           	`json:"defReb"`
	Treb        	string           	`json:"totReb"`
	Assists       string           	`json:"assists"`
	Fouls        	string           	`json:"pFouls"`
	Steals        string           	`json:"steals"`
	Turnovers     string           	`json:"turnovers"`
	Blocks        string           	`json:"blocks"`
	PlusMinus     string           	`json:"plusMinus"`
	DnpReason     string           	`json:"dnp"`
	Position      string            `json:"pos"`
}

type NbaTeam struct {
	NbaTeamData				`json:"totals"`
}

type NbaTeamStats struct {	
	ShortName     string        		`json:"triCode"`	
	Players				[]NbaPlayerData   `json:"activePlayers"`
}

type NbaScoreData struct {
	Score					string						`json:"score"`
}

type NbaTeamScore struct {
	ID            string        		`json:"teamId"`
	ShortName     string        		`json:"triCode"`
	Score         string						`json:"score"`
	ScoreDetail   []NbaScoreData    `json:"linescore"`
}

type NbaStats struct {
	Players			[]NbaPlayerData			`json:"activePlayers"`
	Away				NbaTeam							`json:"vTeam"`
	Home				NbaTeam							`json:"hTeam"`
}

type NbaGameData struct {
	ID						string            `json:"gameId"`
	SID           string            `json:"seasonYear"`
	Date          string            `json:"homeStartDate"`
	AwayScore			NbaTeamScore			`json:"vTeam"`
	HomeScore			NbaTeamScore			`json:"hTeam"`
}

type NbaBoxscore struct {
	NbaGameData		`json:"basicGameData"`
	NbaStats			`json:"stats"`
}

type NbaGame struct {
	ID        		string        `json:"gameId"`
	SID						string				`json:"seasonYear"`
	Date          string        `json:"homeStartDate"`
	Away          NbaTeamScore  `json:"away"`
	Home          NbaTeamScore  `json:"home"`
}

type NbaScoreboard struct {
	Games					[]NbaGame			`json:"games"`
}


func convScores(scores []NbaScoreData) []int {

	ret := []int{}

	for _, s := range scores {
		ret = append(ret, atoi(s.Score))
	}

	return ret

} // convScores


func convPlayers(players []NbaPlayerData, g *Game) {

	awayPlayers := []Player{}
	homePlayers := []Player{}

	for _, p := range players {

		player := Player{}

		player.ID 					= p.ID
		player.Name 				= fmt.Sprintf("%s %s", p.First, p.Last)
		player.Points 			= atoi(p.Points)
		player.Minutes 			= mtoi(p.Minutes)
		player.Fgm 					= atoi(p.Fgm)
		player.Fga 					= atoi(p.Fga)
		player.Fg3m 				= atoi(p.Fg3m)
		player.Fg3a 				= atoi(p.Fg3a)
		player.Ftm 					= atoi(p.Ftm)
		player.Fta 					= atoi(p.Fta)
		player.Oreb					= atoi(p.Oreb)
		player.Dreb 				= atoi(p.Dreb)
		player.Treb 				= atoi(p.Treb)
		player.Assists 			= atoi(p.Assists)
		player.Steals 			= atoi(p.Steals)
		player.Blocks 			= atoi(p.Blocks)
		player.Turnovers 		= atoi(p.Turnovers)
		player.Fouls 				= atoi(p.Fouls)
		player.PlusMinus 		= atoi(p.PlusMinus)
		player.DnpReason 		= p.DnpReason
		player.Position     = p.Position

		if p.TeamID == g.Away.ID {
			awayPlayers = append(awayPlayers, player)
		} else {
			homePlayers = append(homePlayers, player)
		}

	}

	g.Away.Players = awayPlayers
	g.Home.Players = homePlayers

} // convPlayers


func convTeam(ts *NbaTeam, td *Team) {

	td.Points				= atoi(ts.Points)
	td.Fgm					= atoi(ts.Fgm)
	td.Fga					= atoi(ts.Fga)
	td.Fg3m					= atoi(ts.Fg3m)
	td.Fg3a					= atoi(ts.Fg3a)
	td.Ftm					= atoi(ts.Ftm)
	td.Fta					= atoi(ts.Fta)
	td.Oreb					= atoi(ts.Oreb)
	td.Dreb					= atoi(ts.Dreb)
	td.Treb					= atoi(ts.Treb)
	td.Assists			= atoi(ts.Assists)
	td.Steals				= atoi(ts.Steals)
	td.Blocks				= atoi(ts.Blocks)
	td.Fouls				= atoi(ts.Fouls)
	td.Turnovers		= atoi(ts.Turnovers)
	td.PlusMinus    = atoi(ts.PlusMinus)

} // convTeam


func convTeamScore(t NbaTeamScore) Team {

	team := Team{}

	team.ID							= t.ID
	team.Name   				= t.ShortName
	team.Score  				= atoi(t.Score)
	team.ScoreDetail  	= convScores(t.ScoreDetail)

	return team

} // convTeamScore


func convBoxscore(b *NbaBoxscore) *Game {

	game := Game{}

	if b == nil {
		return nil
	}

	game.ID			= b.ID
	game.SID    = b.SID
	game.Date   = b.Date
	game.Away   = convTeamScore(b.AwayScore)
	game.Home   = convTeamScore(b.HomeScore)

	convTeam(&b.Away, &game.Away)
	convTeam(&b.Home, &game.Home)

	convPlayers(b.Players, &game)
	
	return &game

} // convBoxscore


func ScoreboardApi(d string) string {

	if d == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_SCOREBOARD, d))

} // ScoreboardApi


func BoxscoreApi(d string, gid string) string {

	if d == "" || gid == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_BOXSCORE, d, gid))

} // BoxscoreApi


func NbaGetScoreboard(d string) *NbaScoreboard {

	scoreboard := NbaScoreboard{}

	res, err := http.Get(ScoreboardApi(d))

	if err != nil {
		logf("NbaGetScoreboard", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetScoreboard", err.Error())
			return nil
		} else {

			err := json.Unmarshal(buf, &scoreboard)

			if err != nil {
				logf("NbaGetScoreboard", err.Error())
				return nil
			} else {
				return &scoreboard
			}

		}
				
	}

} // NbaGetScoreboard


func NbaGetScoreboardFrom(d string) []NbaScoreboard {

	all := []NbaScoreboard{}

	days := getDays(d)

	for _, day := range days {

		s := NbaGetScoreboard(day)

		if s != nil {						
			all = append(all, *s)			
		}
		
	}

	return all

} // NbaGetScoreboardFrom


func NbaGetBoxscores(s *NbaScoreboard) []NbaBoxscore {

	scores := []NbaBoxscore{}

	if s == nil {
		return nil
	}

	for _, g := range s.Games {

		res, err := http.Get(BoxscoreApi(g.Date, g.ID))

		if err != nil {
			logf("NbaGetBoxscores", err.Error())
		} else {

			defer res.Body.Close()

			buf, err := ioutil.ReadAll(res.Body)

			if err != nil {
				logf("NbaGetBoxscores", err.Error())
			} else {

				box := NbaBoxscore{}

				err := json.Unmarshal(buf, &box)

				if err != nil {
					logf("NbaGetBoxscores", err.Error())
				} else {
					scores = append(scores, box)					
				}

			}

		}

	}

	return scores
	
} // NbaGetBoxscores


func NbaStoreDay(d string) {

	s := NbaGetScoreboard(d)

	if s == nil {
		logf("NbaStoreDay", fmt.Sprintf("No scores found for %s", d))
	} else {

		scores := NbaGetBoxscores(s)
	
		if len(scores) == 0 {
			logf("NbaStoreDay", fmt.Sprintf("No scores found for %s", d))
		} else {
	
			for _, b := range scores {
	
				game := convBoxscore(&b)
	
				putGame(game)				
	
			}
	
		}
	
	}

} // NbaStoreDay


func NbaStoreFromDay(d string) {

	days := getDays(d)

	for _, day := range days {
		NbaStoreDay(day)
	}

} // NbaStoreFromDay


func NbaStoreSeason(s string) {

	season, ok := Seasons[s]

	if ok {
		NbaStoreFromDay(season[SEASON_INDEX_BEGIN])
	} else {
		logf("NbaStoreSeason", fmt.Sprintf("Season not found %s", s))
	}

} // NbaStoreSeason
