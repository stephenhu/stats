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
	Period            int             `json:"period"`
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
		play.Period         = p.Period
		play.Away						= atoi(p.AwayScore)
		play.Home						= atoi(p.HomeScore)
		play.EventType			= atoi(p.EventMsgType)
		play.ScoreChanged		= p.ScoreChange

		n, ok := OfficialTeams[p.TeamID]

		if ok {
			play.TeamName = n
		}

		out = append(out, play)

	}

	return out

} // convGameLog


func PlaysApi(d string, gid string, period int) string {

	if d == "" || gid == "" || period < 1 {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_PLAYS, d, gid, period))

} // PlaysApi


func NbaGetPeriodPlays(d string, gid string, period int) []NbaPlay {

	if d == "" || gid == "" || period < 1 || period > MAX_PERIODS {
		return nil
	}

	res, err := client.Get(PlaysApi(d, gid, period))

	if err != nil {
		logf("NbaGetPeriodPlays", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetPeriodPlays", err.Error())
			return nil
		} else {

			gl := NbaGameLog{}

			err := json.Unmarshal(buf, &gl)

			if err != nil {
				logf("NbaGetPeriodPlays", err.Error())
				return nil
			} else {

				for i, _ := range gl.Plays {
					gl.Plays[i].Period = period
				}

				return gl.Plays

			}

		}

	}

} // NbaGetPeriodPlays


func NbaGetGamePlays(d string, gid string, periods int) []NbaPlay {

	if d == "" || gid == "" || periods < 4 || periods > MAX_PERIODS {
		return nil
	}

	all := []NbaPlay{}

	for p := 1; p <= periods; p++ {

		period := NbaGetPeriodPlays(d, gid, p)

		if period != nil {
			all = append(all, period...)
		}

	}

	return all

} // NbaGetGamePlays
