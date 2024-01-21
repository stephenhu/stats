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
	Fgm               string        `json:"fg2m"`
	Fga               string        `json:"fg2a"`
	Fg3m              string        `json:"fg3m"`
	Fg3a              string        `json:"fg3a"`
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
	Fg2p								string				`json:"fg2p"`			// fg %
	Fg3p							string				`json:"tpp"`			// three point %
	Ftp								string				`json:"ftp"`			// free throw %
}

type NbaTeamData struct {
	Assists       int           		`json:"assists"`
	Atr       		float64           `json:"assistsTurnoverRatio"`
	Bench       	int           		`json:"benchPoints"`
	Bl   					int           		`json:"biggestLead"`
	Bls       		string           	`json:"biggestLeadScore"`
	Bsr       		int           		`json:"biggestScoringRun"`
	Bsrs       		string           	`json:"biggestScoringRunScore"`
	Blocks        int           		`json:"blocks"`
	Blocked       int           	`json:"blocksReceived"`
	FastbreakA    int           	`json:"fastBreakPointsAttempted"`
	FastbreakM    int           	`json:"fastBreaKPointsMade"`
	FastbreakP    float64         `json:"fastBreakPointsPercentage"`
	Fga        		int           	`json:"fieldGoalsAttempted"`
	EffectiveFG   float64         `json:"fieldGoalsEffectiveAdjusted"`
	Fgm        		int           	`json:"fieldGoalsMade"`
	Fgp       		float64         `json:"fieldGoalsPercentage"`
	FoulsO        int           	`json:"foulsOffensive"`
	Fouled       	int           	`json:"foulsDrawn"`
	Fouls        	int           	`json:"foulsPersonal"`
	Technicals    int           	`json:"foulsTechnical"`
	Fta        		int           	`json:"freeThrowsAttempted"`
	Ftm       		int           	`json:"freeThrowsMade"`
	Ftp        		float64         `json:"freeThrowsPercentage"`
	LeadChanges   int           	`json:"leadChanges"`
	Minutes       string          `json:"minutesCalculated"`
	Points       	int           	`json:"points"`
	Against       int           	`json:"pointsAgainst"`
	Fastbreak     int           	`json:"pointsFastBreak"`
	PointsTurnovers     int       `json:"pointsFromTurnovers"`
	Paint       	int           	`json:"pointsInThePaint"`
	PaintA     		int           	`json:"pointsInThePaintAttempted"`
	PaintM        int           	`json:"pointsInThePaintMade"`
	PaintP       	float64         `json:"pointsInThePaintPercentage"`
	PointsSecond  int           	`json:"pointsSecondChance"`
	Dreb       		int           	`json:"reboundsDefensive"`
	Oreb     			int           	`json:"reboundsOffensive"`
	Treb        	int           	`json:"reboundsPersonal"`
	SecondA       int           	`json:"secondChancePointsAttempted"`
	SecondM       int           	`json:"secondChancePointsMade"`
	SecondP       float64         `json:"secondChancePointsPercentage"`
	Steals       	int           	`json:"steals"`
	Fg3a     			int           	`json:"threePointersAttempted"`
	Fg3m       		int           	`json:"threePointersMade"`
	Fg3p        	float64         `json:"threePointersPercentage"`
	Leading       string          `json:"timeLeading"`
	Ties     			int           	`json:"timesTied"`
	TrueA       	float64         		`json:"trueShootingAttempts"`
	TrueP        	float64         `json:"trueShootingPercentage"`
	Turnovers     int           	`json:"turnovers"`
	Fg2a        	int           	`json:"twoPointersAttempted"`
	Fg2m       		int           	`json:"twoPointersMade"`
	Fg2p     			float64         `json:"twoPointersPercentage"`
}

type NbaPlayerData struct {
	ID						string						`json:"personId"`
	TeamID        int            		`json:"teamId"`
	First         string        		`json:"firstName"`
	Last          string        		`json:"familyName"`
	Points        int           	`json:"points"`
	Minutes       string        		`json:"minutesCalculated"`
	Fgm        		int           	`json:"fieldGoalsMade"`
	Fga        		int           	`json:"fieldGoalsAttempted"`
	Ftm        		int           	`json:"freeThrowsMade"`
	Fta        		int           	`json:"freeThrowsAttempted"`
	Fg2m        	int          	`json:"twoPointersMade"`
	Fg2a        	int          	`json:"twoPointersAttempted"`
	Fg3m        	int          	`json:"threePointersMade"`
	Fg3a        	int          	`json:"threePointersAttempted"`
	Oreb        	int           	`json:"reboundsOffensive"`
	Dreb        	int           	`json:"reboundsDefensive"`
	Treb        	int           	`json:"reboundsTotal"`
	Assists       int           		`json:"assists"`
	Fouls        	int           	`json:"foulsPersonal"`
	FoulsOffensive    int           	`json:"foulsOffensive"`
	FoulsDrawn        int           	`json:"foulsDrawn"`
	FoulsTechnical    int           	`json:"foulsTechnical"`
	Steals        int           	`json:"steals"`
	Turnovers     int           	`json:"turnovers"`
	Blocks        int           	`json:"blocks"`
	Blocked       int            `json:"blocksReceived"`
	PlusMinus     float32           	`json:"plusMinusPoints"`
	Position      string            `json:"position"`
	Fastbreak     int               `json:"pointsFastBreak"`
	Paint     		int               `json:"pointsInThePaint"`
	SecondChance  int               `json:"pointsSecondChance"`

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
	Statistics    NbaTeamData       `json:"statistics"`
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
	//Plays         []NbaPlay					`json:"plays"`
}


type NbaGame struct {
	ID        		string        `json:"gameId"`
	GameCode      string      	`json:"gameCode"`
	WeekNumber    int      			`json:"weekNumber"`
	GameTime      string        `json:"gameDateUTC"`		// UTC
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


func NbaGetBoxscoreJson(id string) []byte {
	return apiInvokeJson(BoxscoreApi(id))
} // NbaGetBoxscoreJson
