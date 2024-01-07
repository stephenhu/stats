package stats


type StatsBase struct {
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


type StatsPlayerGame struct {
	GameID                string				`json:"gameId"`
	HomeCode              string				`json:"homeCode"`
	AwayCode              string				`json:"awayCode"`
	GameDate              string        `json:"gameDate"`
	Starter								string        `json:"starter"`
	Minutes								float64       `json:"minutes"`
	PlusMinus							float64       `json:"plusMinus"`
	StatsBase				`json:"stats"`
}


type StatsPlayerInfo struct {
	ID						int						`json:"personId"`
	First					string 				`json:"firstName"`
	Last					string        `json:"familyName"`
	Name          string        `json:"name"`
	NameShort     string        `json:"nameI"`
	Jersey				string				`json:"jerseyNum"`
	Position			string 				`json:"position"`
	//Order         int						`json:"order"` this is used in the game stats
}


type StatsSeasonPlayer struct {
	Games				map[int]map[string]StatsPlayerGame		`json:"games"`
  Info        StatsPlayerInfo 				`json:"info"`
}


type StatsTeamGame struct {
	ID                		int           `json:"id"`
	Abv                   string        `json:"abv"`
	Score                 int           `json:"score"`
	Players               map[int]StatsPlayerGame 	`json:"players"`
}


type StatsTeamInfo struct {
	ID                		int           `json:"id"`
	Name									string				`json:"name"`
	Code									string				`json:"last"`
	City									string				`json:"full"`
	Mascot								string				`json:"abv"`
	Conf									string				`json:"position"`
	Div										string				`json:"active"`
}


type StatsSeasonTeam struct {	
	Games     	map[string]StatsGame			`json:"games"`
	Info        StatsTeamInfo 				`json:"info"`
}


type StatsPeriod struct {
	Number								int           `json:"number"`
	PeriodType						string        `json:"periodType"`
	Away                  int           `json:"away"`
	Home                  int           `json:"home"`
}


type StatsPlay struct {
	ID										string				`json:"id"`
	Detail								string				`json:"detail"`
}


type StatsGame struct {
	GameID                string				`json:"gameId"`
	Date                  string        `json:"date"`
	Away									StatsTeamGame			`json:"away"`
	Home									StatsTeamGame			`json:"home"`
	Periods								map[int]StatsPeriod			`json:"periods"`
	Plays                 map[string]StatsPlay     `json:"plays"`
}


type StatsSeason struct {
	ID                    string        `json:"id"`
	Players				map[int]StatsSeasonPlayer		`json:"players"`
	Teams									int           `json:"teams"`
	Games					map[string]StatsGame           `json:"games"`
	Leaderboard           int           `json:"leaderboard"`
	Standings             int           `json:"standings"`
	Transactions          int           `json:"transactions"`
}


type StatsCache struct {
  Seasons								map[string]StatsSeason			`json:"seasons"`	
}
