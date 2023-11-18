package stats

import (
)

type NbaTeamRanks struct {
	ID								int						    `json:"teamId"`
	Fgp								NbaRankStat			 	`json:"fgp"`
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
	ID            int             `json:"teamId"`
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
	TeamID      int                `json:"teamId"`
}

type NbaStandard2 struct {
	NbaLeague4 `json:"standard"`
}

type NbaStandard3 struct {
	NbaTeam2		`json:"standard"`
}

type NbaRanks struct {
	NbaStandard2		`json:"league"`
	SeasonID				string				`json:"seasonId"`
}

type NbaTeams struct {
	NbaLeague3			`json:"league"`
	SeasonID				string				`json:"seasonId"`
}

type NbaTeamRoster struct {
	NbaStandard3 		`json:"league"`
}

type TeamSites struct {
	Short						string				`json:"teamTricode"`
	Mascot					string				`json:"teamCode"`
	StreakText			string				`json:"streakText"`
}

type NbaTeam3 struct {
	TeamID					int			  		`json:"teamId"`
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
	NbaLeague6      `json:"league"`
}
