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
