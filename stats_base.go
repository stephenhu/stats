package stats

// Basic data structs normalized for better in memory usage

type Base struct {
	Points								int           `json:"points"`
	Oreb									int           `json:"oreb"`
	Dreb									int           `json:"dreb"`
	Treb									int           `json:"treb"`
	Fta										int           `json:"fta"`
	Ftm										int           `json:"ftm"`
	Ftp										float32       `json:"ftp"`
	Fg2a									int           `json:"fg2a"`
	Fg2m									int           `json:"fg2m"`
	Fg2p                  float32				`json:"fg2p"`
	Fg3a									int           `json:"fg3a"`
	Fg3m									int           `json:"fg3m"`
	Fg3p									float32       `json:"fg3p"`
	Fgta                  int						`json:"fgta"`
	Fgtm                  int						`json:"fgtm"`
	Fgtp                  float32				`json:"fgtp"`
	Steals								int           `json:"steals"`
	Assists								int           `json:"assists"`
	Blocks								int           `json:"blocks"`
	Blocked								int           `json:"blocked"`
	Turnovers							int           `json:"turnovers"`
	Fouls									int           `json:"fouls"`
	Fouled								int           `json:"fouled"`
	FoulsO								int           `json:"foulsOffensive"`
	Technicals						int           `json:"technicals"`
	Paint									int           `json:"paint"`
	Fastbreak							int           `json:"fastbreak"`
	SecondChance					int           `json:"secondChance"`
}


type Advanced struct {
	Points								float32       `json:"points"`
	Oreb									float32       `json:"oreb"`
	Dreb									float32       `json:"dreb"`
	Treb									float32       `json:"treb"`
	Fta										float32       `json:"fta"`
	Ftm										float32       `json:"ftm"`
	Ftp										float32       `json:"ftp"`
	Fg2a									float32       `json:"fg2a"`
	Fg2m									float32       `json:"fg2m"`
	Fg2p                  float32				`json:"fg2p"`
	Fg3a									float32       `json:"fg3a"`
	Fg3m									float32       `json:"fg3m"`
	Fg3p									float32       `json:"fg3p"`
	Fgta                  float32				`json:"fgta"`
	Fgtm                  float32				`json:"fgtm"`
	Fgtp                  float32				`json:"fgtp"`
	Steals								float32       `json:"steals"`
	Assists								float32       `json:"assists"`
	Blocks								float32       `json:"blocks"`
	Blocked								float32       `json:"blocked"`
	Turnovers							float32       `json:"turnovers"`
	Fouls									float32       `json:"fouls"`
	Fouled								float32       `json:"fouled"`
	FoulsO								float32       `json:"foulsOffensive"`
	Technicals					  float32       `json:"technicals"`
	Paint									float32       `json:"paint"`
	Fastbreak							float32       `json:"fastbreak"`
	SecondChance					float32       `json:"secondChance"`
	Minutes								int           `json:"minutes"`
	PlusMinus             float32       `json:"plusMinus"`
}


type Leaders struct {
	ID										int						`json:"id"`
	First									string 				`json:"first"`
	Last									string        `json:"last"`
	Full          				string        `json:"full"`
	Abv     							string        `json:"abv"`
	Minutes								int           `json:"minutes"`
	Games									int           `json:"games"`
	PlusMinus             float32       `json:"plusMinus"`
	Base 				`json:"base"`
	Advanced    `json:"advanced"`
}


type Standings struct {
  Wins  								int						`json:"wins"`
	Losses  							int						`json:"losses"`
	GB  									float32				`json:"gamesBack"`
	WinPct  							float32				`json:"winPct"`
	HomeW  								int						`json:"homeWins"`
	HomeL  								int						`json:"homeLosses"`
	AwayW  								int						`json:"awayWins"`
	AwayL  								int						`json:"awayLosses"`
	DivW  								int						`json:"divWins"`
	DivL  								int						`json:"divLosses"`
	ConfW  								int						`json:"confWins"`
	ConfL  								int						`json:"confLosses"`
	Ppg  									float32				`json:"pointsPerGame"`
	Oppg  								float32				`json:"oppPointsPerGame"`
	Streak  							int						`json:"streak"`
	IsW                   bool          `json:"isW"`
	Last10W  							int						`json:"last10W"`
	Last10L  							int						`json:"last10L"`
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
	GameType              int           `json:"gameType"`
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
