package stats

import (
	//"log"
)


func TPlayerInfo(p NbaPlayer) PlayerInfo {
 
  return PlayerInfo{
		ID:	p.ID,
		First: p.First,
		Last: p.Last,
		Jersey: p.Jersey,
	}

} // TPlayerInfo


func TPlayerGame(p NbaPlayer) PlayerGame {
 
  return PlayerGame{
		Minutes: PtmToMin(p.Statistics.Minutes),
		PlusMinus: p.Statistics.PlusMinus,
		Starter: p.Starter,
		Base: Base{
			Points: p.Statistics.Points,
			Oreb: p.Statistics.Oreb,
			Dreb: p.Statistics.Dreb,
			Treb: p.Statistics.Treb,
			Assists: p.Statistics.Assists,
			Steals: p.Statistics.Steals,
			Blocks: p.Statistics.Blocks,
			Blocked: p.Statistics.Blocked,
			Turnovers: p.Statistics.Turnovers,
			Fouls: p.Statistics.Fouls,
			Fouled: p.Statistics.FoulsDrawn,
			FoulsO: p.Statistics.FoulsOff,
			Technicals: p.Statistics.Technicals,
			Fta: p.Statistics.Fta,
			Ftm: p.Statistics.Ftm,
			Fgta: p.Statistics.Fga,
		  Fgtm: p.Statistics.Fgm,
			Fg2a: p.Statistics.Fg2a,
			Fg2m: p.Statistics.Fg2m,
			Fg3a: p.Statistics.Fg3a,
		  Fg3m: p.Statistics.Fg3m,
			Paint: p.Statistics.PointsPaint,
			Fastbreak: p.Statistics.PointsFast,
			SecondChance: p.Statistics.PointsSecond,
		},
	}

} // TPlayerGame


func TGameType(days []NbaGameDates) map[string]int {

	ret := make(map[string]int)

	for _, d := range days {

		for _, g := range d.Games {

			if g.WeekNumber == 0 {
				ret[g.ID] = PRESEASON
			} else if g.WeekNumber > 0 {
				ret[g.ID] = REGULAR
			} else if g.WeekNumber > 25 {
				ret[g.ID] = PLAYOFFS
			}
	
		}

	}

	return ret

} // TGameType
