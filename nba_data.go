package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
  "log"
	//"net/http"
	"strings"
)

type NbaRankStat struct {
	Average				string					`json:"avg"`
	Rank					string					`json:"rank"`
}

type NbaInternal struct {
	PubDate				string						`json:"pubDateTime"`
	NumGames      int               `json:"numGames"`
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
	Tpg								string				`json:"topg"`			// turnovers per game
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
	Periods   		[]NbaScoreData    `json:"linescore"`
}

type NbaStats struct {
	Players			[]NbaPlayerData			`json:"activePlayers"`
	Away				NbaTeam							`json:"vTeam"`
	Home				NbaTeam							`json:"hTeam"`
}

type NbaGameData struct {
	ID						string            `json:"gameId"`
	SeasonID      string            `json:"seasonYear"`
	Date          string            `json:"homeStartDate"`
	StartUtc      string        		`json:"startTimeUTC"`
	EndUtc    		string        		`json:"endTimeUTC"`
	AwayScore			NbaTeamScore			`json:"vTeam"`
	HomeScore			NbaTeamScore			`json:"hTeam"`
	Plays         []NbaPlay					`json:"plays"`
}

type NbaBoxscore struct {
	NbaInternal   `json:"_internal"`
	NbaGameData		`json:"basicGameData"`
	NbaStats			`json:"stats"`
}

type NbaGame struct {
	ID        		string        `json:"gameId"`
	SeasonID			string				`json:"seasonYear"`
	Date          string        `json:"startDateEastern"`
	StartUtc      string        `json:"startTimeUTC"`
	EndUtc    		string        `json:"endTimeUTC"`
	Away          NbaTeamScore  `json:"vTeam"`
	Home          NbaTeamScore  `json:"hTeam"`
}

type NbaScoreboard struct {
	NbaInternal   `json:"_internal"`
	Games					[]NbaGame			`json:"games"`
	Date					string				`json:"date"`
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

		minutes, seconds := mtoi(p.Minutes)

		player.ID 					= p.ID
		player.Name 				= fmt.Sprintf("%s %s", p.First, p.Last)
		player.Points 			= atoi(p.Points)
		player.Minutes 			= minutes
		player.Seconds      = seconds
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
			player.Opponent = g.Home.Name
			awayPlayers = append(awayPlayers, player)
		} else {
			player.Opponent = g.Away.Name
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
	team.Name   				= strings.ToLower(t.ShortName)
	team.Score  				= atoi(t.Score)
	team.Periods  			= convScores(t.Periods)

	return team

} // convTeamScore


func ConvBoxscore(b *NbaBoxscore) *Game {

	game := Game{}

	if b == nil {
		return nil
	}

	game.ID					= b.ID
	game.SeasonID   = b.SeasonID
	game.Date   		= b.Date
	game.StartUtc   = b.StartUtc
	game.EndUtc   	= b.EndUtc
	game.PubDate   	= b.PubDate
	game.Away   		= convTeamScore(b.AwayScore)
	game.Home   		= convTeamScore(b.HomeScore)
	game.Plays      = convPlays(b.Plays)

	convTeam(&b.Away, &game.Away)
	convTeam(&b.Home, &game.Home)

	convPlayers(b.Players, &game)

	return &game

} // ConvBoxscore


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

	log.Println(ScoreboardApi(d))
  if len(d) == 0 {
		return nil
	}

	res, err := client.Get(ScoreboardApi(d))

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

				scoreboard.Date	= d

				if len(scoreboard.Games) == 0 {
					return nil
				}

				return &scoreboard

			}

		}

	}

} // NbaGetScoreboard


func NbaGetScoreboardFrom(d string) []NbaScoreboard {

	all := []NbaScoreboard{}

	days := GetDays(d)

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

		score := NbaGetBoxscore(g.Date, g.ID)

		if score != nil {
			scores = append(scores, *score)
		}

	}

	return scores

} // NbaGetBoxscores


func NbaGetBoxscore(d string, g string) *NbaBoxscore {

	res, err := client.Get(BoxscoreApi(d, g))

	if err != nil {
		logf("NbaGetBoxscores", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetBoxscores", err.Error())
			return nil
		} else {

			box := NbaBoxscore{}

			err := json.Unmarshal(buf, &box)

			if err != nil {
				logf("NbaGetBoxscores", err.Error())
				return nil
			} else {

				plays := NbaGetGamePlays(box.Date, box.ID,
					len(box.AwayScore.Periods))

				if plays != nil {
					box.Plays = plays
				}

				return &box

			}

		}

	}

} // NbaGetBoxscore
