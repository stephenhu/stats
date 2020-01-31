package stats

import (
	"fmt"
)


func Percentage(p *Player, fn string) string {

	switch fn {
	case FIELD_FG:

		if p.Fga == 0 {
			return STRING_ZERO_FLOAT
		}

		fmt.Sprintf("%.1f", (float64(p.Fgm)/float64(p.Fga)*FLOAT_TO_PERCENT))

	case FIELD_FG3:

		if p.Fga == 0 {
			return STRING_ZERO_FLOAT
		}

		fmt.Sprintf("%.1f", (float64(p.Fg3m)/float64(p.Fg3a)*FLOAT_TO_PERCENT))
		
	case FIELD_FT:

		if p.Fga == 0 {
			return STRING_ZERO_FLOAT
		}

		fmt.Sprintf("%.1f", (float64(p.Ftm)/float64(p.Fta)*FLOAT_TO_PERCENT))
		
	default:
		logf("percentage", fmt.Sprintf(
			"Field %s does not allow percentage calculation.", fn))
			
	}

	return STRING_ZERO_FLOAT

} // Percentage


func GameAnalysis() {

} // GameAnalysis
