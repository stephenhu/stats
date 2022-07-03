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

type TeamSites struct {
	Short						string				`json:"teamTricode"`
	Mascot					string				`json:"teamCode"`
	StreakText			string				`json:"streakText"`
}

type NbaTeam3 struct {
	TeamID					string				`json:"teamId"`
	W								string				`json:"win"`
	L								string				`json:"loss"`
	Pct							string				`json:"winPct"`
	Gb							string				`json:"gamesBehind"`
	Rank            string        `json:"confRank"`
	Cw							string				`json:"confWin"`
	Cl							string				`json:"confLoss"`
	Dw							string				`json:"divWin"`
	Dl							string				`json:"divLoss"`
	Hw							string				`json:"homeWin"`
	Hl							string				`json:"homeLoss"`
	Aw							string				`json:"awayWin"`
	Al							string				`json:"awayLoss"`
	L10w						string				`json:"lastTenWin"`
	L10l						string				`json:"lastTenLoss"`
	Streak					string				`json:"streak"`
	WinStreak				bool					`json:"isWinStreak"`

	TeamSites				`json:"teamSitesOnly"`
}

type NbaStandard4 struct {
	SeasonID				int						`json:"seasonYear"`
	Teams           []NbaTeam3		`json:"teams"`
}

type NbaLeague6 struct {
	NbaStandard4    `json:"standard"`
}

type NbaTeamStandings struct {
	NbaInternal     `json:"_internal"`
	NbaLeague6      `json:"league"`
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

				tr.ID										= rank.ID
				tr.Fgp.Val							= atof(rank.Fgp.Average)
				tr.Fgp.Rank							= atoi(rank.Fgp.Rank)
				tr.Fg3p.Val							= atof(rank.Fg3p.Average)
				tr.Fg3p.Rank						= atoi(rank.Fg3p.Rank)
				tr.Ftp.Val							= atof(rank.Ftp.Average)
				tr.Ftp.Rank							= atoi(rank.Ftp.Rank)
				tr.Oreb.Val							= atof(rank.Oreb.Average)
				tr.Oreb.Rank						= atoi(rank.Oreb.Rank)
				tr.Dreb.Val							= atof(rank.Dreb.Average)
				tr.Dreb.Rank						= atoi(rank.Dreb.Rank)
				tr.Treb.Val							= atof(rank.Treb.Average)
				tr.Treb.Rank						= atoi(rank.Treb.Rank)
				tr.Assists.Val					= atof(rank.Assists.Average)
				tr.Assists.Rank					= atoi(rank.Assists.Rank)
				tr.Turnovers.Val				= atof(rank.Turnovers.Average)
				tr.Turnovers.Rank				= atoi(rank.Turnovers.Rank)
				tr.Steals.Val						= atof(rank.Steals.Average)
				tr.Steals.Rank					= atoi(rank.Steals.Rank)
				tr.Blocks.Val						= atof(rank.Blocks.Average)
				tr.Blocks.Rank					= atoi(rank.Blocks.Rank)
				tr.Fouls.Val						= atof(rank.Fouls.Average)
				tr.Fouls.Rank						= atoi(rank.Fouls.Rank)
				tr.Points.Val						= atof(rank.Points.Average)
				tr.Points.Rank					= atoi(rank.Points.Rank)
				tr.OpponentPoints.Val		= atof(rank.OpponentPoints.Average)
				tr.OpponentPoints.Rank	= atoi(rank.OpponentPoints.Rank)
				tr.Efficiency.Val       = atof(rank.Efficiency.Average)
				tr.Efficiency.Rank			= atoi(rank.Efficiency.Rank)
				tr.Name     						= strings.ToLower(rank.Abbreviation)

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


func convStandings(standings *NbaTeamStandings) *Standings {

	if standings == nil {
		return nil
	}

	s := Standings{
		Records: make(map[string]TeamRecord),
	}

	s.PubDate 	= standings.PubDate
	s.SeasonID	= fmt.Sprintf("%d", standings.SeasonID)

	for _, team := range standings.Teams {

		tr := TeamRecord{}

		tr.TeamID				= team.TeamID
		tr.Name					= strings.ToLower(team.Short)
		tr.Mascot				= strings.ToLower(team.Mascot)
		tr.Rank       	= atoi(team.Rank)
		tr.W						= atoi(team.W)
		tr.L						= atoi(team.L)
		tr.Pct					= atof(team.Pct)
		tr.Gb						= atof(team.Gb)
		tr.Cw						= atoi(team.Cw)
		tr.Cl						= atoi(team.Cl)
		tr.Dw						= atoi(team.Dw)
		tr.Dl						= atoi(team.Dl)
		tr.Hw						= atoi(team.Hw)
		tr.Hl						= atoi(team.Hl)
		tr.Aw						= atoi(team.Aw)
		tr.Al						= atoi(team.Al)
		tr.L10w					= atoi(team.L10w)
		tr.L10l					= atoi(team.L10l)
		tr.Streak				= atoi(team.Streak)
		tr.WinStreak		= team.WinStreak
		tr.StreakText 	= team.StreakText

		s.Records[tr.Name] = tr

	}

	return &s

} // convStandings


func TeamRanksApi(y int) string {

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_TEAM_RANKS, y))

} // TeamRanksApi


func TeamRosterApi(y int, n string) string {

	if n == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
	  NBA_BASE_URL,
		fmt.Sprintf(NBA_API_ROSTER, y, n))

} // TeamRosterApi


func TeamsApi(y int) string {

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_TEAMS, y))

} // TeamsApi


func TeamStandingsApi(y int) string {

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		NBA_API_STANDINGS)

} // TeamStandingsApi


func NbaGetTeams(y int) *NbaTeams {

	teams := NbaTeams{}

	teams.SeasonID = fmt.Sprintf("%d", y)

	res, err := client.Get(TeamsApi(y))

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


func NbaGetTeamRanks(y int) *NbaRanks {

	ranks := NbaRanks{}

	ranks.SeasonID = fmt.Sprintf("%d", y)

	res, err := client.Get(TeamRanksApi(y))

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


func NbaGetTeamRoster(y int, n string) *NbaTeamRoster {

	if n == "" {
		return nil
	}

	tr := TeamRosterApi(y, n)

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


func NbaGetTeamStandings(y int) *NbaTeamStandings {

	standings := NbaTeamStandings{}

	res, err := client.Get(TeamStandingsApi(y))

	if err != nil {
		logf("NbaGetTeamStandings", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetTeamStandings", err.Error())
			return nil
		} else {

			err := json.Unmarshal(buf, &standings)

			if err != nil {
				logf("NbaGetTeamStandings", err.Error())
				return nil
			} else {
				return &standings
			}

		}

	}

} // NbaGetTeamStandings
