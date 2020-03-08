package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	//"net/http"
)

type NbaPlayerInfo struct {
	ID						string				`json:"personId"`
	First					string 				`json:"firstName"`
	Last					string        `json:"lastName"`
	TeamID				string				`json:"teamId"`
	Jersey				string				`json:"jersey"`
	Position			string				`json:"pos"`
	Feet					string				`json:"heightFeet"`
	Inches				string				`json:"heightInches"`
	Meters				string				`json:"heightMeters"`
	Pounds				string				`json:"weightPounds"`
	Kilograms			string				`json:"weightKilograms"`
	Dob						string				`json:"dateOfBirthUTC"`
	Rookie				string				`json:"nbaDebutYear"`
	Years      		string        `json:"yearsPro"`
	College       string        `json:"collegeName"`
	Active      	bool        	`json:"isActive"`
}

type NbaSeasonStats struct {
	SeasonID			int							`json:"seasonYear"`
	Teams					[]NbaAdvStats		`json:"teams"`
	Total					NbaAdvStats			`json:"total"`
}

type NbaRegularSeason struct {
	SeasonInfo 		[]NbaSeasonStats 		`json:"season"`
}

type NbaPlayerStats struct {
	Latest 				NbaAdvStats		`json:"latest"`
	Career 				NbaAdvStats		`json:"careerSummary"`
	NbaRegularSeason		`json:"regularSeason"`

}

type NbaStats2 struct {
	NbaPlayerStats		`json:"stats"`
}

type NbaLeague2 struct {
	NbaStats2				`json:"standard"`
}

type NbaLeague struct {
	Players 	[]NbaPlayerInfo			`json:"standard"`
}

type NbaLeaguePlayers struct {
	NbaLeague				`json:"league"`
	NbaInternal  		`json:"_internal"`
	SeasonID        string				`json:"season"`
}

type NbaPlayerProfile struct {
	NbaLeague2				`json:"league"`
	NbaInternal       `json:"_internal"`
	ID						string        `json:"id"`
	SeasonID			string        `json:"seasonId"`
	First         string        `json:"first"`
	Last         	string        `json:"last"`
}


func copyAdvStats(src *NbaAdvStats, dst *AdvStats) {

	dst.SeasonID			= fmt.Sprintf("%d", src.SeasonID)
	dst.TeamID				= src.TeamID
	dst.Minutes     	= atoi(src.Minutes)
	dst.Points     		= atoi(src.Points)
	dst.Oreb     			= atoi(src.Oreb)
	dst.Dreb     			= atoi(src.Dreb)
	dst.Treb     			= atoi(src.Treb)
	dst.Assists     	= atoi(src.Assists)
	dst.Turnovers   	= atoi(src.Turnovers)
	dst.Steals     		= atoi(src.Steals)
	dst.Blocks     		= atoi(src.Blocks)
	dst.Fouls     		= atoi(src.Fouls)
	dst.Fgm     			= atoi(src.Fgm)
	dst.Fga     			= atoi(src.Fga)
	dst.Fg3m     			= atoi(src.Fg3m)
	dst.Fg3a     			= atoi(src.Fg3a)
	dst.Ftm     			= atoi(src.Ftm)
	dst.Fta     			= atoi(src.Fta)
	dst.Played     		= atoi(src.Played)
	dst.Started     	= atoi(src.Started)
	dst.PlusMinus   	= atoi(src.PlusMinus)
	dst.Ppg     			= atof(src.Ppg)
	dst.Rpg     			= atof(src.Rpg)
	dst.Apg     			= atof(src.Apg)
	dst.Mpg     			= atof(src.Mpg)
	dst.Tpg     			= atof(src.Tpg)
	dst.Spg     			= atof(src.Spg)
	dst.Bpg     			= atof(src.Bpg)
	dst.Fgp     			= atof(src.Fgp)
	dst.Fg3p     			= atof(src.Fg3p)
	dst.Ftp     			= atof(src.Ftp)

} // copyAdvStats


func convLeaguePlayers(lp *NbaLeaguePlayers) *AllPlayers {

	if lp != nil {

		all := AllPlayers{}

		all.SeasonID  = lp.SeasonID
		all.PubDate   = lp.PubDate

		for _, p := range lp.Players {

			tid := filterId(p.TeamID)

			_, ok := OfficialTeams[tid]

			if ok {

				pi := PlayerInfo{}

				pi.ID							= p.ID
				pi.First					= p.First
				pi.Last						= p.Last
				pi.TeamID					= tid
				pi.Jersey					= p.Jersey
				pi.Position				= p.Position
				pi.Feet						= atoi(p.Feet)
				pi.Inches					= atoi(p.Inches)
				pi.Meters					= p.Meters
				pi.Pounds					= atoi(p.Pounds)
				pi.Kilograms			= p.Kilograms
				pi.Dob						= p.Dob
				pi.Rookie					= p.Rookie
				pi.Years					= atoi(p.Years)
				pi.College				= p.College
				pi.Active					= p.Active

				all.Players = append(all.Players, pi)

			}

		}

		return &all

	} else {
		return nil
	}

} // convLeaguePlayers


func convPlayerProfile(pp *NbaPlayerProfile) *PlayerCareer {

	if pp == nil {
		return nil
	}

	career := PlayerCareer{}

	career.ID 				= pp.ID
	career.PubDate    = pp.PubDate
	career.SeasonID   = pp.SeasonID
	career.First			= pp.First
	career.Last				= pp.Last

	copyAdvStats(&pp.Latest, &career.Latest)
	copyAdvStats(&pp.Career, &career.Career)

	for _, player := range pp.SeasonInfo {

		ss := SeasonStats{}

		ss.SeasonID	= fmt.Sprintf("%d", player.SeasonID)

		copyAdvStats(&player.Total, &ss.Summary)

		for _, team := range player.Teams {

			as := AdvStats{}

			copyAdvStats(&team, &as)

			ss.Teams = append(ss.Teams, as)

		}

		career.Seasons = append(career.Seasons, ss)

	}

	return &career

} // convPlayerProfile


func convCareerSeason(pc *PlayerCareer) *PlayerSeason {

	if pc != nil {

		ps := PlayerSeason{}

		ps.Mpg 					= pc.Latest.Mpg
		ps.Ppg 					= pc.Latest.Ppg
		ps.Rpg 					= pc.Latest.Rpg
		ps.Apg 					= pc.Latest.Apg
		ps.Spg 					= pc.Latest.Spg
		ps.Bpg 					= pc.Latest.Bpg
		ps.Tpg          = pc.Latest.Tpg
		ps.Fgp 					= pc.Latest.Fgp
		ps.Fg3p 				= pc.Latest.Fg3p
		ps.Ftp 					= pc.Latest.Ftp
		ps.Mpg 					= pc.Latest.Mpg
		ps.Points				= pc.Latest.Points
		ps.Minutes			= pc.Latest.Minutes
		ps.Oreb 				= pc.Latest.Oreb
		ps.Dreb 				= pc.Latest.Dreb
		ps.Treb 				= pc.Latest.Treb
		ps.Fgm 					= pc.Latest.Fgm
		ps.Fga 					= pc.Latest.Fga
		ps.Fg3m 				= pc.Latest.Fg3m
		ps.Fg3a 				= pc.Latest.Fg3a
		ps.Fta 					= pc.Latest.Ftm
		ps.Ftm 					= pc.Latest.Fta
		ps.Fouls 				= pc.Latest.Fouls
		ps.Assists 			= pc.Latest.Assists
		ps.Steals 			= pc.Latest.Steals
		ps.Blocks 			= pc.Latest.Blocks
		ps.Turnovers 		= pc.Latest.Turnovers
		ps.Played       = pc.Latest.Played
		ps.Started      = pc.Latest.Started
		ps.PlusMinus    = pc.Latest.PlusMinus

		return &ps

	} else {
		return nil
	}

} // convCareerSeason


func PlayersApi(s string) string {

	if s == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_PLAYERS, s))

} // PlayersApi


func PlayerProfileApi(s string, pid string) string {

	if s == "" || pid == "" {
		return ""
	}

	return fmt.Sprintf("%s%s",
		NBA_BASE_URL,
		fmt.Sprintf(NBA_API_PLAYER_PROFILE, s, pid))

} // PlayerProfileApi


func NbaGetPlayers(s string) *NbaLeaguePlayers {

	if s == "" {
		return nil
	}

	lp := NbaLeaguePlayers{}

	res, err := client.Get(PlayersApi(s))

	if err != nil {
		logf("NbaGetPlayers", err.Error())
		return nil
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			logf("NbaGetPlayers", err.Error())
			return nil
		} else {

			err := json.Unmarshal(buf, &lp)

			if err != nil {
				logf("NbaGetPlayers", err.Error())
				return nil
			} else {

				lp.SeasonID	= s

				return &lp

			}

		}

	}

} // NbaGetPlayers


func NbaGetProfiles(s string, lp *NbaLeaguePlayers) []NbaPlayerProfile {

	if lp == nil {
		return nil
	} else {

		all := []NbaPlayerProfile{}

		for _, p := range lp.Players {

			if p.Active {

				res, err := client.Get(PlayerProfileApi(s, p.ID))

				if err != nil {
					logf("NbaGetProfiles", err.Error())
				} else {

					buf, err := ioutil.ReadAll(res.Body)

					if err != nil {
						logf("NbaGetProfiles", err.Error())
						return nil
					} else {

						player := NbaPlayerProfile{}

						player.SeasonID		= lp.SeasonID

						err := json.Unmarshal(buf, &player)

						if err != nil {
							logf("NbaGetProfiles", err.Error())
							return nil
						} else {

							player.ID     		= p.ID
							player.First 			= p.First
							player.Last 			= p.Last

							all = append(all, player)

						}

					}

				}

			}

		}

		return all

	}

} // NbaGetProfiles


func NbaStorePlayers(lp *NbaLeaguePlayers) {

	if lp != nil {

		all := convLeaguePlayers(lp)

		putPlayers(all)

	}

} // NbaStorePlayers


func NbaStoreProfiles(profiles []NbaPlayerProfile) {

	for _, profile := range profiles {

		career := convPlayerProfile(&profile)

		putProfile(career)

	}

} // NbaStoreProfiles
