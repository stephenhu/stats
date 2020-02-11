package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type NbaFormatted struct {
	Description				string					`json:"description"`
}

type NbaPlay struct {
	Clock							string					`json:"clock"`
	Description				string					`json:"description"`
	PersonID					string					`json:"personId"`
	TeamID						string					`json:"teamId"`
	AwayScore         string          `json:"vTeamScore"`
	HomeScore         string          `json:"hTeamScore"`
	EventMsgType      string          `json:"eventMsgType"`
	ScoreChange       bool            `json:"isScoreChange"`	
	FmtDesc    				NbaFormatted		`json:"formatted"`
}

type NbaGameLog struct {
	NbaInternal				`json:"_internal"`
	Plays							[]NbaPlay				`json:"plays"`
	Date          		string					`json:"date"`
	GameID          	string					`json:"gameId"`
}


func convPlays(plays []NbaPlay) []Play {

	if plays == nil {
		return nil
	}	

	out := []Play{}

	for _, p := range plays {

		play := Play{}

		play.Clock					= p.Clock
		play.Description		= p.Description
		play.Formatted			= p.FmtDesc.Description
		play.PersonID				= p.PersonID
		play.TeamID					= p.TeamID
		play.Away						= atoi(p.AwayScore)
		play.Home						= atoi(p.HomeScore)
		play.EventType			= atoi(p.EventMsgType)
		play.ScoreChanged		= p.ScoreChange

		out = append(out, play)

	}

	return out

} // convGameLog


func PlaysApi(d string, gid string) string {

	if d == "" || gid == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_PLAYS, d, gid, 1))

} // PlaysApi


func NbaGetPlays(d string, gid string) []NbaPlay {

	if d == "" || gid == "" {
		return nil
	}

	gl := NbaGameLog{}

	res, err := client.Get(PlaysApi(d, gid))

	if err != nil {
		logf("NbaGetPlays", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetPlays", err.Error())
			return nil
		} else {

			err := json.Unmarshal(buf, &gl)

			if err != nil {
				logf("NbaGetPlays", err.Error())
				return nil
			} else {								

				return gl.Plays

			}

		}
				
	}

} // NbaGetPlays
