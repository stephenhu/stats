package stats

import (
	"fmt"
)


func Percentage(p *Player, fn string) string {

	switch fn {
	case FIELD_FG:

		if p.Stats.Fga == 0 {
			return STRING_ZERO_FLOAT
		}

		fmt.Sprintf("%.1f", (float64(p.Stats.Fgm)/float64(p.Stats.Fga)*FLOAT_TO_PERCENT))

	case FIELD_FG3:

		if p.Stats.Fg3a == 0 {
			return STRING_ZERO_FLOAT
		}

		fmt.Sprintf("%.1f", (float64(p.Stats.Fg3m)/float64(p.Stats.Fg3a)*FLOAT_TO_PERCENT))

	case FIELD_FT:

		if p.Stats.Fta == 0 {
			return STRING_ZERO_FLOAT
		}

		fmt.Sprintf("%.1f", (float64(p.Stats.Ftm)/float64(p.Stats.Fta)*FLOAT_TO_PERCENT))

	default:
		logf("percentage", fmt.Sprintf(
			"Field %s does not allow percentage calculation.", fn))

	}

	return STRING_ZERO_FLOAT

} // Percentage


func GameAnalysis() {

} // GameAnalysis


func PlayerAverages(games []Player) AdvStats {

	as := AdvStats{}

	for _, g := range games {

		as.Minutes 		+= g.Minutes
		as.Points  		+= g.Points
		as.Oreb    		+= g.Oreb
		as.Dreb    		+= g.Dreb
		as.Treb    		+= g.Treb
		as.Assists    += g.Assists
		as.Turnovers  += g.Turnovers
		as.Steals    	+= g.Steals
		as.Blocks     += g.Blocks
		as.Fouls      += g.Fouls
		as.Fgm				+= g.Fgm
		as.Fga				+= g.Fga
		as.Fg3m				+= g.Fg3m
		as.Fg3a				+= g.Fg3a
		as.Ftm				+= g.Ftm
		as.Fta				+= g.Fta
		as.PlusMinus  += g.PlusMinus

	}

	count := float32(len(games))

	if count > 0 {

		as.Ppg 	= float32(as.Points)/count
		as.Rpg 	= float32(as.Treb)/count
		as.Apg 	= float32(as.Assists)/count
		as.Mpg 	= float32(as.Minutes)/count
		as.Tpg = float32(as.Turnovers)/count
		as.Spg 	= float32(as.Steals)/count
		as.Bpg 	= float32(as.Blocks)/count

	}

	if as.Fga > 0 {
		as.Fgp 	= float32(as.Fgm)/float32(as.Fga)
	}

	if as.Fg3a > 0 {
		as.Fg3p = float32(as.Fg3m)/float32(as.Fg3a)
	}

	if as.Ftp > 0 {
		as.Ftp 	= float32(as.Ftm)/float32(as.Fta)
	}

	return as

} // PlayerAverages


func TeamAverages(games []Team) AdvStats {

	as := AdvStats{}

	for _, g := range games {

		as.Points  		+= g.Points
		as.Oreb    		+= g.Oreb
		as.Dreb    		+= g.Dreb
		as.Treb    		+= g.Treb
		as.Assists    += g.Assists
		as.Turnovers  += g.Turnovers
		as.Steals    	+= g.Steals
		as.Blocks     += g.Blocks
		as.Fouls      += g.Fouls
		as.Fgm				+= g.Fgm
		as.Fga				+= g.Fga
		as.Fg3m				+= g.Fg3m
		as.Fg3a				+= g.Fg3a
		as.Ftm				+= g.Ftm
		as.Fta				+= g.Fta
		as.PlusMinus  += g.PlusMinus

	}

	count := float32(len(games))

	if count > 0 {

		as.Ppg 	= float32(as.Points)/count
		as.Rpg 	= float32(as.Treb)/count
		as.Apg 	= float32(as.Assists)/count
		as.Mpg 	= float32(as.Minutes)/count
		as.Tpg = float32(as.Turnovers)/count
		as.Spg 	= float32(as.Steals)/count
		as.Bpg 	= float32(as.Blocks)/count

	}

	if as.Fga > 0 {
		as.Fgp 	= float32(as.Fgm)/float32(as.Fga)
	}

	if as.Fg3a > 0 {
		as.Fg3p = float32(as.Fg3m)/float32(as.Fg3a)
	}

	if as.Ftp > 0 {
		as.Ftp 	= float32(as.Ftm)/float32(as.Fta)
	}

	return as

} // TeamAverages
