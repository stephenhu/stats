package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NbaTeamStandard struct {
	ID            string          `json:"teamId"`
	IsNba 				bool						`json:"isNBAFranchise"`
	City          string          `json:"city"`
	Code     			string          `json:"tricode"`
	Full          string          `json:"fullName"`
	Short        	string          `json:"nickname"`
	Conference    string          `json:"confName"`
	Division      string          `json:"divName"`
}

type NbaLeague3 struct {
	Teams 		[]NbaTeamStandard			`json:"standard"`
}

type NbaTeams struct {
	NbaInternal			`json:"_internal"`
	NbaLeague3			`json:"league"`
	SeasonID				string				`json:"seasonId"`
}


func convTeamInfo(teams *NbaTeams) *AllTeams {

	at := AllTeams{}

	at.SeasonID	= teams.SeasonID

	for _, t := range teams.Teams {

		if t.IsNba {

			ti := TeamInfo{}

			ti.ID						= t.ID
			ti.Full     		= t.Full
			ti.Short    		= t.Short
			ti.City     		= t.City
			ti.Code     		= t.Code
			ti.Conference		= t.Conference
			ti.Division			= t.Division
	
			at.Teams = append(at.Teams, ti)

		}

	}

	return &at

} // convTeamInfo


func TeamsApi(s string) string {
	
	if s == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_TEAMS, s))

} // TeamsApi


func NbaGetTeams(d string) *NbaTeams {

	teams := NbaTeams{}

	teams.SeasonID = d

	res, err := http.Get(TeamsApi(d))

	if err != nil {
		logf("NbaGetTeams", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetTeams", err.Error())
			return nil
		} else {

			err := json.Unmarshal(buf, &teams)

			if err != nil {
				logf("NbaGetTeams", err.Error())
				return nil
			} else {				
				return &teams
			}

		}
				
	}
	
} // NbaGetTeams


func NbaStoreTeams(t *NbaTeams) {

	if t != nil {
		
		all := convTeamInfo(t)

		putTeams(all)

	}

} // NbaStoreTeams
