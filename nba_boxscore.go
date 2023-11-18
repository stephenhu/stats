package stats

import (
  "fmt"
)


type NbaRankStat struct {
	Average				string					`json:"avg"`
	Rank					string					`json:"rank"`
}


type NbaAdvStats struct {
	SeasonID					int						`json:"seasonYear"`
	TeamID						int           `json:"teamId"`
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
	TeamID        int            		`json:"teamId"`
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
	Score					int								`json:"score"`
	PeriodType    string            `json:"periodType"`
	Period    		int            		`json:"period"`
}

type NbaTeamScore struct {
	ID            int        				`json:"teamId"`
	ShortName     string        		`json:"teamTriCode"`
	Score         int						    `json:"score"`
	Periods   		[]NbaScoreData    `json:"periods"`
	Players       []NbaPlayer       `json:"players"`
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


type NbaGame struct {
	ID        		string        `json:"gameId"`
	GameCode      string      	`json:"gameCode"`
	WeekNumber    int      			`json:"weekNumber"`
	Away          NbaTeamScore  `json:"awayTeam"`
	Home          NbaTeamScore  `json:"homeTeam"`
}


type NbaBoxscore struct {
	Meta          NbaMeta						`json:"meta"`
	Game          NbaGame						`json:"game"`
}


type NbaScoreboard struct {
	GameDate      string        `json:"gameDate"`
	LeagueId      string        `json:"leagueId"`
	LeagueName    string        `json:"leagueName"`
	Games					[]NbaGame			`json:"games"`
}


func BoxscoreApi(id string) string {
	return fmt.Sprintf("%s%s%s", NBA_BASE_URL, NBA_LIVE,
	  fmt.Sprintf(NBA_API_BOXSCORE, id))
} // BoxscoreApi


func NbaGetBoxscore(id string) *NbaBoxscore {

	box := NbaBoxscore{}

	apiInvoke(BoxscoreApi(id), &box)

	return &box

} // NbaGetBoxscore
