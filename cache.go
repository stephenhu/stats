package stats

import (
	//"log"
)


type Base struct {
	Points								int           `json:"points"`
	Oreb									int           `json:"oreb"`
	Dreb									int           `json:"dreb"`
	Treb									int           `json:"treb"`
	Fta										int           `json:"fta"`
	Ftm										int           `json:"ftm"`
	Ftp										float64       `json:"ftp"`
	Fg2a									int           `json:"fg2a"`
	Fg2m									int           `json:"fg2m"`
	Fg2p                  float64				`json:"fg2p"`
	Fg3a									int           `json:"fg3a"`
	Fg3m									int           `json:"fg3m"`
	Fg3p									float64       `json:"fg3p"`
	Fgta                  int						`json:"fgta"`
	Fgtm                  int						`json:"fgtm"`
	Fgtp                  float64				`json:"fgtp"`
	Steals								int           `json:"steals"`
	Assists								int           `json:"assists"`
	Blocks								int           `json:"blocks"`
	Blocked								int           `json:"blocked"`
	Turnovers							int           `json:"turnovers"`
	Fouls									int           `json:"fouls"`
	Fouled								int           `json:"fouled"`
	FoulsOffensive				int           `json:"foulsOffensive"`
	Technicals						int           `json:"technicals"`
	Paint									int           `json:"paint"`
	Fastbreak							int           `json:"fastbreak"`
	SecondChance					int           `json:"secondChance"`
}


type PlayerGame struct {
	GameID                string				`json:"gameId"`
	TeamID                int           `json:"teamId"`
	HomeTeamID            int        		`json:"homeTeamId"`
	AwayTeamID            int        		`json:"awayTeamId"`
	HomeCode              string				`json:"homeCode"`
	AwayCode              string				`json:"awayCode"`
	GameDate              string        `json:"gameDate"`
	Starter								string        `json:"starter"`
	Minutes								int       		`json:"minutes"`
	PlusMinus							float64       `json:"plusMinus"`
	Base				`json:"stats"`
}


type PlayerInfo struct {
	ID						int						`json:"personId"`
	First					string 				`json:"firstName"`
	Last					string        `json:"familyName"`
	Name          string        `json:"name"`
	NameShort     string        `json:"nameI"`
	Jersey				string				`json:"jerseyNum"`
	Position			string 				`json:"position"`
	//Order         int						`json:"order"` this is used in the game stats
}


type PlayerSeason struct {
	Games				map[string]PlayerGame		`json:"games"`
  Info        PlayerInfo 				`json:"info"`
}


type TeamGame struct {
	ID                		int           `json:"id"`
	Abv                   string        `json:"abv"`
	Score                 int           `json:"score"`
	Players               map[int]PlayerGame 	`json:"players"`
}


type TeamInfo struct {
	ID                		int           `json:"id"`
	Name									string				`json:"name"`
	Code									string				`json:"last"`
	City									string				`json:"full"`
	Mascot								string				`json:"abv"`
	Conf									string				`json:"position"`
	Div										string				`json:"active"`
}


type SeasonTeam struct {	
	Games     	map[string]Game			`json:"games"`
	Info        TeamInfo 				`json:"info"`
}


type Period struct {
	Number								int           `json:"number"`
	PeriodType						string        `json:"periodType"`
	Away                  int           `json:"away"`
	Home                  int           `json:"home"`
}


type Play struct {
	ID										string				`json:"id"`
	Detail								string				`json:"detail"`
}


type Game struct {
	GameID                string				`json:"gameId"`
	Date                  string        `json:"date"`
	Away									TeamGame			`json:"away"`
	Home									TeamGame			`json:"home"`
	Periods								map[int]Period			`json:"periods"`
	Plays                 map[string]Play     `json:"plays"`
}


type Season struct {
	ID                    string        `json:"id"`
	Players				map[int]PlayerSeason		`json:"players"`
	Teams									int           `json:"teams"`
	Games					map[string]Game           `json:"games"`
	Leaderboard           int           `json:"leaderboard"`
	Standings             int           `json:"standings"`
	Transactions          int           `json:"transactions"`
}


type Cache struct {
  Seasons								map[string]Season			`json:"seasons"`	
}


func TPlayerInfo(p NbaPlayer) PlayerInfo {
 
  return PlayerInfo{
		ID:	p.ID,
		First: p.First,
		Last: p.Last,
		Jersey: p.Jersey,
	}

} // TPlayerInfo


func TPlayerGame(p NbaPlayer) PlayerGame {
 
  return PlayerGame{
		Minutes: PtmToMin(p.Statistics.Minutes),
		PlusMinus: p.Statistics.PlusMinus,
		Starter: p.Starter,
		Base: Base{
			Points: p.Statistics.Points,
			Oreb: p.Statistics.Oreb,
			Dreb: p.Statistics.Dreb,
			Treb: p.Statistics.Treb,
			Assists: p.Statistics.Assists,
			Steals: p.Statistics.Steals,
			Blocks: p.Statistics.Blocks,
			Blocked: p.Statistics.Blocked,
			Turnovers: p.Statistics.Turnovers,
			Fouls: p.Statistics.Fouls,
			Fouled: p.Statistics.FoulsDrawn,
			FoulsOffensive: p.Statistics.FoulsOff,
			Technicals: p.Statistics.Technicals,
			Fta: p.Statistics.Fta,
			Ftm: p.Statistics.Ftm,
			Fgta: p.Statistics.Fga,
		  Fgtm: p.Statistics.Fgm,
			Fg2a: p.Statistics.Fg2a,
			Fg2m: p.Statistics.Fg2m,
			Fg3a: p.Statistics.Fg3a,
		  Fg3m: p.Statistics.Fg3m,
			Paint: p.Statistics.PointsPaint,
			Fastbreak: p.Statistics.PointsFast,
			SecondChance: p.Statistics.PointsSecond,
		},
	}

} // TPlayerGame