package stats

import (
	"fmt"
	//"log"
)

var OfficialEras = map[string]string {
	PROFILE_SIMPLE_ERA: "2015",
	PROFILE_MODERN_ERA: "1979",
	PROFILE_BIRD_ERA: "1978",
	PROFILE_RELATIVE_ERA: "1998",
}

func GetEra(s string) {

	var start string

	switch s {
	case PROFILE_SIMPLE_ERA:
		start = OfficialEras[s]
	default:
		logf("GetEra", fmt.Sprintf("%s is not a supported era", s))
	}

	years := getYearsFrom(start)

	if years != nil {

		for _, y := range years {
			NbaStoreAll(y)
		}

	}

} // GetEra
