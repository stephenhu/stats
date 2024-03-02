package stats

import (
	//"fmt"
  //"log"
)

type StreakData struct {
	Streak				int			`json:"streak"`
	IsWinStreak		bool		`json:"IsWinStreak"`
}


type L10Data struct {
	L10W          int			`json:"L10W"`
	L10L          int			`json:"L10L"`
}


var (
	LastM 				map[int]*L10Data
  StreakM    		map[int]*StreakData
)


var (
	GBCE 						= float32(0.5)
	WESTERN_CONF    = "Western"
	EASTERN_CONF    = "Eastern"
)


func Pct(dividend int, divisor int) float32 {

	if divisor == 0 {
		return 0.0
	} else {
		return float32(dividend)/float32(divisor)
	}

} // pct


func PctFp(dividend float32, divisor int) float32 {

	if divisor == 0 {
		return 0.0
	} else {
		return dividend/float32(divisor)
	}

} // PctFp


func GamePlayed(mins int) int {

	if mins == 0 {
		return 0
	} else {
		return 1
	}

} // GamePlayed


// a/(a+b)
func PctTotal(a int, b int) float32 {
  
	total := a+b

	if total == 0 {
		return 0.0
	} else {
		return float32(a)/float32(total)
	}

} // PctTotal


func ParseWl(s map[int]*Standing, game *NbaGame) {

	if game.Home.ID == 0 || game.Away.ID == 0 {
		return
	}

	if !CheckNbaTeam(game.Home.ID, game.Away.ID) {
		return
	}

	_, ok := s[game.Home.ID]

	if !ok {
		s[game.Home.ID] = &Standing{}
	}

	_, ok2 := s[game.Away.ID]

	if !ok2 {
		s[game.Away.ID] = &Standing{}
	}

	if game.Home.Score > game.Away.Score {

		s[game.Home.ID].Wins++
		s[game.Home.ID].HomeW++
		s[game.Away.ID].Losses++
		s[game.Away.ID].AwayL++

		if AllTeams[game.Home.ID].Conference == AllTeams[game.Away.ID].Conference {
			s[game.Home.ID].ConfW++
			s[game.Away.ID].ConfL++			
		}

		if AllTeams[game.Home.ID].Division == AllTeams[game.Away.ID].Division {
			s[game.Home.ID].DivW++
			s[game.Away.ID].DivL++
		}

	} else {

		s[game.Home.ID].Losses++
		s[game.Home.ID].HomeL++
		s[game.Away.ID].Wins++
		s[game.Away.ID].AwayW++
	

		if AllTeams[game.Home.ID].Conference == AllTeams[game.Away.ID].Conference {
			s[game.Home.ID].ConfL++
			s[game.Away.ID].ConfW++			
		}

		if AllTeams[game.Home.ID].Division == AllTeams[game.Away.ID].Division {
			s[game.Home.ID].DivL++
			s[game.Away.ID].DivW++
		}

	}

} // ParseWl


func UpdateStreak(s map[int]*Standing) {

	for _, t := range s {

		if t.TeamID == 0 {
			continue
		}
		
		s[t.TeamID].Streak 	= StreakM[t.TeamID].Streak
		s[t.TeamID].IsW 		= StreakM[t.TeamID].IsWinStreak

	}

} // UpdateStreak


func Streak(scores []NbaBoxscore) {

	StreakM = make(map[int]*StreakData)

	for _, s := range scores {

		_, ok := StreakM[s.Game.Home.ID]

		if !ok {
			StreakM[s.Game.Home.ID] = &StreakData{}
		}

		_, ok2 := StreakM[s.Game.Away.ID]

		if !ok2 {
			StreakM[s.Game.Away.ID] = &StreakData{}
		}

		if s.Game.Away.Score > s.Game.Home.Score {

			if StreakM[s.Game.Away.ID].IsWinStreak {
				StreakM[s.Game.Away.ID].Streak++
			} else {
				StreakM[s.Game.Away.ID].IsWinStreak = true
				StreakM[s.Game.Away.ID].Streak = 1
			}

			if StreakM[s.Game.Home.ID].IsWinStreak {
				StreakM[s.Game.Home.ID].IsWinStreak = false
				StreakM[s.Game.Home.ID].Streak = 1
			} else {
				StreakM[s.Game.Home.ID].Streak++
			}

		} else {

			if StreakM[s.Game.Home.ID].IsWinStreak {
				StreakM[s.Game.Home.ID].Streak++
			} else {
				StreakM[s.Game.Home.ID].IsWinStreak = true
				StreakM[s.Game.Home.ID].Streak = 1
			}

			if StreakM[s.Game.Away.ID].IsWinStreak {
				StreakM[s.Game.Away.ID].IsWinStreak = false
				StreakM[s.Game.Away.ID].Streak = 1
			} else {
				StreakM[s.Game.Away.ID].Streak++
			}

		}

	}

} // Streak


func UpdateLast10(s map[int]*Standing) {

	for _, t := range s {

		if t.TeamID == 0 {
			continue
		}

		s[t.TeamID].Last10W	= LastM[t.TeamID].L10W
		s[t.TeamID].Last10L = LastM[t.TeamID].L10L
		
	}

} // UpdateLast10


// assumes boxscores in order by date
func Last10(scores []NbaBoxscore) {

	LastM = make(map[int]*L10Data)

	games := len(scores)

	for i, _ := range scores {

		latestGame := scores[games - 1 - i]

		_, ok := LastM[latestGame.Game.Home.ID]

		if !ok {
			LastM[latestGame.Game.Home.ID] = &L10Data{}
		}

		_, ok2 := LastM[latestGame.Game.Away.ID]

		if !ok2 {
			LastM[latestGame.Game.Away.ID] = &L10Data{}
		}

		if latestGame.Game.Home.Score > latestGame.Game.Away.Score {

			if LastM[latestGame.Game.Home.ID].L10W +
		  	LastM[latestGame.Game.Home.ID].L10L < 10 {

				LastM[latestGame.Game.Home.ID].L10W++
						
			}

			if LastM[latestGame.Game.Away.ID].L10W +
			  LastM[latestGame.Game.Away.ID].L10L < 10 {

				LastM[latestGame.Game.Away.ID].L10L++
			
			}

		} else {

			if LastM[latestGame.Game.Away.ID].L10W +
			  LastM[latestGame.Game.Away.ID].L10L < 10 {

				LastM[latestGame.Game.Away.ID].L10W++

			}

			if LastM[latestGame.Game.Home.ID].L10W +
			  LastM[latestGame.Game.Home.ID].L10L < 10 {

				LastM[latestGame.Game.Home.ID].L10L++
				
			}
				
		}

	}

} // Last10


func TopTeam(s map[int]*Standing) (*Standing, *Standing) {

	var east, west *Standing

	for k, v := range s {

		if AllTeams[k].Conference == WESTERN_CONF {

			if west == nil {
				west = v
			} else {
	
				if v.WinPct > west.WinPct {
					west = v
				}
			
			}
	
		} else if AllTeams[k].Conference == EASTERN_CONF {

			if east == nil {
				east = v
			} else {
	
				if v.WinPct > east.WinPct {
					east = v
				}
			
			}
	
		}

	}

	return east, west

} // TopTeam


func GamesBack(s map[int]*Standing) {

	east, west := TopTeam(s)

	for _, team := range s {

		if team.TeamID != east.TeamID && team.TeamID != west.TeamID {

			if AllTeams[team.TeamID].Conference == WESTERN_CONF {

				s[team.TeamID].Gb = float32(west.Wins - team.Wins) * GBCE +
			  	float32(team.Losses - west.Losses) * GBCE

			} else {

				s[team.TeamID].Gb = float32(east.Wins - team.Wins) * GBCE +
			  	float32(team.Losses - east.Losses) * GBCE

			}

		}
	}

} // GamesBack


func WinPct(s map[int]*Standing) {

	for _, v := range s {
		
		v.WinPct = PctTotal(v.Wins, v.Losses)

	}

} // WinPct


func initStanding(s map[int]*Standing, id int) {

	if id == 0 {
		return
	}

	_, ok := s[id]

	if !ok {
		s[id] = &Standing{}
	}

} // initStanding


func initBase(s map[int]*Standing, id int) {

	if id == 0 {
		return
	}

	_, ok := s[id]

	if !ok {
		s[id] = &Standing{
			Base: Base{},
		}
	}

} // initBase


func CalcTeamPct(s map[int]*Standing) {

	for tid, team := range s {

		s[tid].Base.Ftp = Pct(team.Base.Ftm, team.Base.Fta)
		s[tid].Base.Fg2p = Pct(team.Base.Fg2m, team.Base.Fg2a)
		s[tid].Base.Fg3p = Pct(team.Base.Fg3m, team.Base.Fg3a)
		s[tid].Base.Fgtp = Pct(team.Base.Fgtm, team.Base.Fgta)
		s[tid].Base.At = Pct(team.Base.Assists, team.Base.Turnovers)

	}

} // CalcTeamPct


func CalcTeamStats(s map[int]*Standing, scores []NbaBoxscore) {

	for _, score := range scores {

		initBase(s, score.Game.Home.ID)
		initBase(s, score.Game.Away.ID)

		if score.Game.ID == "" {
			continue
		}

		s[score.Game.Home.ID].Points 	+= score.Game.Home.Score
		s[score.Game.Home.ID].Oreb 		+= score.Game.Home.Statistics.Oreb
		s[score.Game.Home.ID].Dreb 		+= score.Game.Home.Statistics.Dreb
		s[score.Game.Home.ID].Treb		+= score.Game.Home.Statistics.Treb
		s[score.Game.Home.ID].Fta 		+= score.Game.Home.Statistics.Fta
		s[score.Game.Home.ID].Ftm 		+= score.Game.Home.Statistics.Ftm
		s[score.Game.Home.ID].Fg2a 		+= score.Game.Home.Statistics.Fg2a
		s[score.Game.Home.ID].Fg2m 		+= score.Game.Home.Statistics.Fg2m
		s[score.Game.Home.ID].Fg3a 		+= score.Game.Home.Statistics.Fg3a
		s[score.Game.Home.ID].Fg3m 		+= score.Game.Home.Statistics.Fg3m
		s[score.Game.Home.ID].Fgta 		+= score.Game.Home.Statistics.Fga
		s[score.Game.Home.ID].Fgtm 		+= score.Game.Home.Statistics.Fgm
		s[score.Game.Home.ID].Steals  += score.Game.Home.Statistics.Steals
		s[score.Game.Home.ID].Assists 		+= score.Game.Home.Statistics.Assists
		s[score.Game.Home.ID].Blocks 			+= score.Game.Home.Statistics.Blocks
		s[score.Game.Home.ID].Blocked 		+= score.Game.Home.Statistics.Blocked
		s[score.Game.Home.ID].Turnovers 	+= score.Game.Home.Statistics.Turnovers
		s[score.Game.Home.ID].Fouls 			+= score.Game.Home.Statistics.Fouls
		s[score.Game.Home.ID].Fouled 			+= score.Game.Home.Statistics.Fouled
		s[score.Game.Home.ID].FoulsO 			+= score.Game.Home.Statistics.FoulsO
		s[score.Game.Home.ID].Technicals 	+= score.Game.Home.Statistics.Technicals
		s[score.Game.Home.ID].Paint 			+= score.Game.Home.Statistics.Paint
		s[score.Game.Home.ID].Fastbreak 			+= score.Game.Home.Statistics.Fastbreak
		s[score.Game.Home.ID].SecondChance 		+= score.Game.Home.Statistics.PointsSecond

		s[score.Game.Home.ID].TeamAdvanced.Bench 	+= score.Game.Home.Statistics.Bench
		s[score.Game.Home.ID].TeamAdvanced.Pft 		+= score.Game.Home.Statistics.PointsTurnovers
		s[score.Game.Home.ID].TeamAdvanced.Pf 		+= score.Game.Home.Statistics.Points
		s[score.Game.Home.ID].TeamAdvanced.Pa 		+= score.Game.Home.Statistics.Against

		for _, period := range score.Game.Home.Periods {
			
			if period.Period == 1 {
				s[score.Game.Home.ID].TeamAdvanced.P1 		+= period.Score
			} else if period.Period == 2 {
				s[score.Game.Home.ID].TeamAdvanced.P2 		+= period.Score
			} else if period.Period == 3 {
				s[score.Game.Home.ID].TeamAdvanced.P3 		+= period.Score
			} else if period.Period == 4 {
				s[score.Game.Home.ID].TeamAdvanced.P4 		+= period.Score
			} else {
				s[score.Game.Home.ID].TeamAdvanced.Pot 		+= period.Score
			}

		}

		s[score.Game.Away.ID].Points 	+= score.Game.Away.Score
		s[score.Game.Away.ID].Oreb 		+= score.Game.Away.Statistics.Oreb
		s[score.Game.Away.ID].Dreb 		+= score.Game.Away.Statistics.Dreb
		s[score.Game.Away.ID].Treb		+= score.Game.Away.Statistics.Treb
		s[score.Game.Away.ID].Fta 		+= score.Game.Away.Statistics.Fta
		s[score.Game.Away.ID].Ftm 		+= score.Game.Away.Statistics.Ftm
		s[score.Game.Away.ID].Fg2a 		+= score.Game.Away.Statistics.Fg2a
		s[score.Game.Away.ID].Fg2m 		+= score.Game.Away.Statistics.Fg2m
		s[score.Game.Away.ID].Fg3a 		+= score.Game.Away.Statistics.Fg3a
		s[score.Game.Away.ID].Fg3m 		+= score.Game.Away.Statistics.Fg3m
		s[score.Game.Away.ID].Fgta 		+= score.Game.Away.Statistics.Fga
		s[score.Game.Away.ID].Fgtm 		+= score.Game.Away.Statistics.Fgm
		s[score.Game.Away.ID].Steals  += score.Game.Away.Statistics.Steals
		s[score.Game.Away.ID].Assists 		+= score.Game.Away.Statistics.Assists
		s[score.Game.Away.ID].Blocks 			+= score.Game.Away.Statistics.Blocks
		s[score.Game.Away.ID].Blocked 		+= score.Game.Away.Statistics.Blocked
		s[score.Game.Away.ID].Turnovers 	+= score.Game.Away.Statistics.Turnovers
		s[score.Game.Away.ID].Fouls 			+= score.Game.Away.Statistics.Fouls
		s[score.Game.Away.ID].Fouled 			+= score.Game.Away.Statistics.Fouled
		s[score.Game.Away.ID].FoulsO 			+= score.Game.Away.Statistics.FoulsO
		s[score.Game.Away.ID].Technicals 	+= score.Game.Away.Statistics.Technicals
		s[score.Game.Away.ID].Paint 			+= score.Game.Away.Statistics.Paint
		s[score.Game.Away.ID].Fastbreak 			+= score.Game.Away.Statistics.Fastbreak
		s[score.Game.Away.ID].SecondChance 		+= score.Game.Away.Statistics.PointsSecond

		for _, period := range score.Game.Away.Periods {
			
			if period.Period == 1 {
				s[score.Game.Away.ID].TeamAdvanced.P1 		+= period.Score
			} else if period.Period == 2 {
				s[score.Game.Away.ID].TeamAdvanced.P2 		+= period.Score
			} else if period.Period == 3 {
				s[score.Game.Away.ID].TeamAdvanced.P3 		+= period.Score
			} else if period.Period == 4 {
				s[score.Game.Away.ID].TeamAdvanced.P4 		+= period.Score
			} else {
				s[score.Game.Away.ID].TeamAdvanced.Pot 		+= period.Score
			}

		}

		s[score.Game.Away.ID].TeamAdvanced.Bench 	+= score.Game.Away.Statistics.Bench
		s[score.Game.Away.ID].TeamAdvanced.Pft 		+= score.Game.Away.Statistics.PointsTurnovers
		s[score.Game.Away.ID].TeamAdvanced.Pf 		+= score.Game.Away.Statistics.Points
		s[score.Game.Away.ID].TeamAdvanced.Pa 		+= score.Game.Away.Statistics.Against

	}

	CalcTeamPct(s)

} // CalcTeamStats


func CalculateStandings(s map[int]*Standing, scores []NbaBoxscore) {

	Last10(scores)

	Streak(scores)

	WinPct(s)

	GamesBack(s)

	CalcTeamStats(s, scores)

} // CalculateStandings
