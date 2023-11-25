package stats

import (
	"fmt"
)

type NbaFormatted struct {
	Description				string					`json:"description"`
}

type NbaPlay struct {
	Clock							string					`json:"clock"`
	Description				string					`json:"description"`
	PersonID					string					`json:"personId"`
	TeamID						int							`json:"teamId"`
	Period            int             `json:"period"`
	AwayScore         string          `json:"vTeamScore"`
	HomeScore         string          `json:"hTeamScore"`
	EventMsgType      string          `json:"eventMsgType"`
	ScoreChange       bool            `json:"isScoreChange"`
	FmtDesc    				NbaFormatted		`json:"formatted"`
}

type NbaGameLog struct {
	Plays							[]NbaPlay				`json:"plays"`
	Date          		string					`json:"date"`
	GameID          	string					`json:"gameId"`
}


func PlaysApi(d string, gid string, period int) string {

	if d == "" || gid == "" || period < 1 {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_PLAYS, gid))

} // PlaysApi
