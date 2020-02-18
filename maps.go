package stats

import (
	"fmt"
	"os"
	//"path/filepath"
	//"strings"
)

var seasons_map				= map[string] Season{}						// seasons
//var teams_map 				= map[string] TeamInfo {}					// team list
//var team_stats_map		= map[string] TeamSeasonStats{}		// team season stats


func teamKey(s string, id string) string {
	return fmt.Sprintf("%s.%s", s, id)
} // teamKey


func loadSeasons() []string {

	f, err := os.Open(APP_STORAGE)

	if err != nil {
		logf("loadSeasons", err.Error())
		return []string{}
	} else {

		dirs, err := f.Readdirnames(0)

		if err != nil {
			logf("loadSeasons", err.Error())
			return []string{}
		} else {
			return dirs
		}

	}

} // loadSeasons


func LoadTeams() {
/*
	seasons := loadSeasons()

	for _, season := range seasons {

		p := filepath.Join(APP_STORAGE, season, TEAMS_FILE)

		buf := loadFile(p)

		teams := AllTeams{}

		err := json.Unmarshal(buf, &teams)

		if err != nil {
			logf("LoadTeams", err.Error())
		} else {

			for _, t := range teams.Teams {
				teams_map[teamKey(season, t.ID)] = t
			}
			
		}
		
	}
*/
} // LoadTeams


func LoadTeamStats() {
} // LoadTeamStats


func LoadGames() {
} // LoadGames


func LoadCache() {

	if checkStorage() {

		LoadTeams()
		
	}

} // LoadCache
