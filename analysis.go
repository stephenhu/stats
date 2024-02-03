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

	if game.Home.Statistics.Points > game.Away.Statistics.Points {

		s[game.Home.ID].Wins++
		s[game.Home.ID].HomeW++
		s[game.Away.ID].Losses++
		s[game.Away.ID].AwayL++
		
		if tm[game.Home.ID].Conference == tm[game.Away.ID].Conference {
			s[game.Home.ID].ConfW++
			s[game.Away.ID].ConfL++
		}

		if tm[game.Home.ID].Division == tm[game.Away.ID].Division {
			s[game.Home.ID].DivW++
			s[game.Away.ID].DivL++
		}

	} else {

		s[game.Away.ID].Wins++
		s[game.Away.ID].AwayW++
		s[game.Home.ID].Losses++
		s[game.Home.ID].HomeL++

		if tm[game.Home.ID].Conference == tm[game.Away.ID].Conference {
			s[game.Away.ID].ConfW++
			s[game.Home.ID].ConfL++
		}

		if tm[game.Home.ID].Division == tm[game.Away.ID].Division {
			s[game.Away.ID].DivW++
			s[game.Home.ID].DivL++
		}

	}

} // ParseWl
