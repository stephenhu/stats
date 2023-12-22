package stats

import (
	"fmt"
)

type NbaFormatted struct {
	Description				string					`json:"description"`
}

type NbaAction struct {
	ActionNumber      int             `json:"actionNumber"`
	Clock							string					`json:"clock"`
  TimeActual        string          `json:"timeActual"`
	Description				string					`json:"description"`
	Descriptor				string					`json:"descriptor"`
	PersonID					int							`json:"personId"`
	TeamID						int							`json:"teamId"`
	TeamTriCode       string          `json:"teamTricode"`
	ActionType        string          `json:"actionType"`
	SubActionType     string          `json:"subType"`
	Qualifers         []string        `json:"qualifiers"`
	X                 float64					`json:"x"`
	Y                 float64					`json:"y"`
	Side              string					`json:"side"`
	ShotDistance      float64         `json:"shotDistance"`
	Possession        int             `json:"possession"`
	Period            int             `json:"period"`
	PeriodType        string          `json:"periodType"`
	AwayScore         string          `json:"scoreAway"`
	HomeScore         string          `json:"scoreHome"`
	Edited            string          `json:"edited"`
	OrderNumber       int							`json:"orderNumber"`
	XLegacy           int             `json:"xLegacy"`
	YLegacy           int             `json:"yLegacy"`
	IsFieldGoal       int             `json:"isFieldGoal"`
	ShotResult        string          `json:"shotResult"`
	Name        			string          `json:"playerName"`
	NameI       			string          `json:"playerNameI"`
	PersonIdsFilter   []int           `json:"personIdsFilter"`
	ScoreChange       bool            `json:"isScoreChange"`
	StealName         string          `json:"stealPlayerName"`
	StealPersonID     int             `json:"stealPersonId"`
	AssistName        string          `json:"assistPlayerName"`
	AssistPersonID    int             `json:"assistPersonId"`
	AssistNameI       string          `json:"assistPlayerNameInitial"`
	AssistTotal       int             `json:"assistTotal"`
	BlockName         string          `json:"blockPlayerName"`
	BlockPersonID     int             `json:"blockPersonId"`
	BlockTotal        int             `json:"blockTotal"`
	FoulDrawnPersonID int             `json:"foulDrawnPersonId"`
	FoulDrawnName     string          `json:"foulDrawnPlayerName"`
	FoulPersonTotal   int             `json:"foulPersonTotal"`
	FoulTechTotal     int        			`json:"foulTechnicalTotal"`
	JumpLostName      string          `json:"jumpBallLostPlayerName"`
	JumpLostPersonID  int             `json:"jumpBallLostPersonId"`
	JumpRecoveredName        string          `json:"jumpBallRecoveredPlayerName"`
	JumpRecoveredPersonID    int             `json:"jumpBallRecoveredPersonId"`
	JumpWonName       string          `json:"jumpBallWonPlayerName"`
	JumpWonPersonID   int             `json:"jumpBallWonPersonId"`
	OfficialID        int             `json:"officialId"`
	ShotNum       		int          		`json:"shotActionNumber"`
	TurnoverTotal     int             `json:"turnoverTotal"`
	Value       			string          `json:"value"`
}

type NbaGameActions struct {
	Actions							[]NbaAction			`json:"actions"`
}

type NbaPlayByPlay struct {
	Meta          		NbaMeta					`json:"meta"`
	Game          		NbaGameActions	`json:"game"`
}


func PlaysApi(d string) string {

	return fmt.Sprintf("%s%s%s", NBA_BASE_URL, NBA_LIVE,
	  fmt.Sprintf(NBA_API_PLAYS, d))

} // PlaysApi


func NbaGetPlays(id string) *NbaPlayByPlay {

	plays := NbaPlayByPlay{}

	apiInvoke(PlaysApi(id), &plays)

	return &plays

} // NbaGetPlays


func NbaGetPlaysJson(id string) []byte {
  return apiInvokeJson(PlaysApi(id))
} // NbaGetPlaysJson
