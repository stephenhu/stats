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


func CalculateStandings(s map[int]*Standing, scores []NbaBoxscore) {

	// last10
	Last10(scores)

	// gb
	// streak
	Streak(scores)

	for _, t := range s {
		
		t.WinPct = PctTotal(t.Wins, t.Losses)

	}

} // CalculateStandings
