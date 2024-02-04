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
