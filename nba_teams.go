package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	//"net/http"
	"strings"
)

type NbaTeamRanks struct {
	ID								string						`json:"teamId"`
	Fgp								NbaRankStat				`json:"fgp"`
	Fg3p							NbaRankStat				`json:"tpp"`
	Ftp								NbaRankStat				`json:"ftp"`
	Oreb							NbaRankStat				`json:"orpg"`
	Dreb							NbaRankStat				`json:"drpg"`
	Treb							NbaRankStat				`json:"trpg"`
	Assists						NbaRankStat				`json:"apg"`
	Turnovers					NbaRankStat				`json:"tpg"`
	Steals						NbaRankStat				`json:"spg"`
	Blocks						NbaRankStat				`json:"bpg"`
	Fouls							NbaRankStat				`json:"pfpg"`
	Points						NbaRankStat				`json:"ppg"`
	OpponentPoints		NbaRankStat				`json:"oppg"`
	Efficiency				NbaRankStat				`json:"eff"`
	Minutes						NbaRankStat				`json:"min"`
	Abbreviation      string            `json:"abbreviation"`
}

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

type NbaRegularSeason2 struct {
	Teams				[]NbaTeamRanks		`json:"teams"`
}

type NbaLeague3 struct {
	Teams 		[]NbaTeamStandard			`json:"standard"`
}

type NbaLeague4 struct {
	NbaRegularSeason2 `json:"regularSeason"`
}

type NbaLeague5 struct {
	NbaRegularSeason2 `json:"regularSeason"`
}

type NbaTeamPlayer struct {
	ID						string							`json:"personId"`
}

type NbaTeam2 struct {
	Players			[]NbaTeamPlayer				`json:"players"`
	TeamID      string                `json:"teamId"`
}

type NbaStandard2 struct {
	NbaLeague4 `json:"standard"`
}

type NbaStandard3 struct {
	NbaTeam2		`json:"standard"`
}

type NbaRanks struct {
	NbaInternal			`json:"_internal"`
	NbaStandard2		`json:"league"`
	SeasonID				string				`json:"seasonId"`
}

type NbaTeams struct {
	NbaInternal			`json:"_internal"`
	NbaLeague3			`json:"league"`
	SeasonID				string				`json:"seasonId"`
}

type NbaTeamRoster struct {
	NbaInternal			`json:"_internal"`
	NbaStandard3 		`json:"league"`
}


func convTeamRanks(ranks *NbaRanks) *AllRanks {

	all := AllRanks{}

	if ranks != nil {

		all.SeasonID = ranks.SeasonID
		all.PubDate  = ranks.PubDate

		for _, rank := range ranks.Teams {

			_, ok := OfficialTeams[rank.ID]

			if ok {

				tr := TeamRanks{}

				tr.ID								= rank.ID
				tr.Fgp							= atof(rank.Fgp.Average)
				tr.Fg3p							= atof(rank.Fg3p.Average)
				tr.Ftp							= atof(rank.Ftp.Average)
				tr.Oreb							= atof(rank.Oreb.Average)
				tr.Dreb							= atof(rank.Dreb.Average)
				tr.Treb							= atof(rank.Treb.Average)
				tr.Assists					= atof(rank.Assists.Average)
				tr.Turnovers				= atof(rank.Turnovers.Average)
				tr.Steals						= atof(rank.Steals.Average)
				tr.Blocks						= atof(rank.Blocks.Average)
				tr.Fouls						= atof(rank.Fouls.Average)
				tr.Points						= atof(rank.Points.Average)
				tr.OpponentPoints		= atof(rank.OpponentPoints.Average)
				tr.Efficiency       = atof(rank.Efficiency.Average)
				tr.Name     				= strings.ToLower(rank.Abbreviation)

				all.Teams = append(all.Teams, tr)

			}

		}

		return &all

	} else {
		return nil
	}


} // convTeamRanks


func convTeamInfo(teams *NbaTeams) *AllTeams {

	at := AllTeams{}

	at.PubDate  = teams.PubDate
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


func convRoster(roster *NbaTeamRoster) *Roster {

	if roster == nil {
		return nil
	}

	r := Roster{}

	r.TeamID = roster.TeamID

	for _, player := range roster.Players {
		r.Players = append(r.Players, player.ID)
	}

	return &r

} // convRoster


func TeamRanksApi(s string) string {

	if s == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_TEAM_RANKS, s))

} // TeamRanksApi


func TeamRosterApi(s string, n string) string {

	if s == "" || n == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
	  NBA_BASE_URL,
		fmt.Sprintf(NBA_API_ROSTER, s, n))

} // TeamRosterApi


func TeamsApi(s string) string {

	if s == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_TEAMS, s))

} // TeamsApi


func NbaGetTeams(s string) *NbaTeams {

	teams := NbaTeams{}

	teams.SeasonID = s

	res, err := client.Get(TeamsApi(s))

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


func NbaGetTeamRanks(s string) *NbaRanks {

	ranks := NbaRanks{}

	ranks.SeasonID = s

	res, err := client.Get(TeamRanksApi(s))

	if err != nil {
		logf("NbaGetTeamRanks", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetTeamRanks", err.Error())
			return nil
		} else {

			err := json.Unmarshal(buf, &ranks)

			if err != nil {
				logf("NbaGetTeamRanks", err.Error())
				return nil
			} else {
				return &ranks
			}

		}

	}

} // NbaGetTeamRanks


func NbaGetTeamRoster(s string, n string) *NbaTeamRoster {

	if s == "" || n == "" {
		return nil
	}

	tr := TeamRosterApi(s, n)

	res, err := client.Get(tr)

	if err != nil {
		logf("NbaGetTeamRoster", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetTeamRoster", err.Error())
			return nil
		} else {

			roster := NbaTeamRoster{}

			err := json.Unmarshal(buf, &roster)

			if err != nil {
				logf("NbaGetTeamRoster", err.Error())
				return nil
			} else {
				return &roster
			}

		}

	}

} // NbaGetTeamRoster


func NbaStoreTeams(t *NbaTeams) {

	if t != nil {

		all := convTeamInfo(t)

		putTeams(all)

	}

} // NbaStoreTeams


func NbaStoreTeamRanks(r *NbaRanks) {

	if r != nil {

		all := convTeamRanks(r)

		putTeamRanks(all)

	}

} // NbaStoreTeamRanks
