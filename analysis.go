package stats

import (
	//"log"
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


func ParseWl(s map[int]*Standing, game *NbaGame) {

	home, ok := s[game.Home.ID]

	if !ok {
		s[game.Home.ID] = &Standing{}
	}

	away, ok2 := s[game.Away.ID]

	if !ok2 {
	  s[game.Away.ID] = &Standing{}
	}

	if game.Home.Score > game.Away.Score {

		home.Wins++
		home.HomeW++
		away.Losses++
		away.AwayL++

		if AllTeams[game.Home.ID].Conference == AllTeams[game.Away.ID].Conference {
			home.ConfW++
			away.ConfL++			
		}

		if AllTeams[game.Home.ID].Division == AllTeams[game.Away.ID].Division {
			home.DivW++
			away.DivL++
		}

	} else {

		home.Losses++
		home.HomeL++
		away.Wins++
		away.AwayW++
	
		if AllTeams[game.Home.ID].Conference == AllTeams[game.Away.ID].Conference {
			home.ConfL++
			away.ConfW++			
		}

		if AllTeams[game.Home.ID].Division == AllTeams[game.Away.ID].Division {
			home.DivL++
			away.DivW++
		}

	}

} // ParseWl
