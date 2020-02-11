package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"strings"
)


var teams_map = map[string] TeamInfo {}


func teamKey(s string, id string) string {
	return fmt.Sprintf("teams.%s.%s", s, id)
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


func LoadTeams() bool {

	seasons := loadSeasons()

	for _, season := range seasons {

		p := filepath.Join(APP_STORAGE, season, TEAMS_FILE)

		buf, err := ioutil.ReadFile(p)

		if err != nil {
			logf("LoadTeams", err.Error())
		} else {

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
		
	}

	return true

} // LoadTeams


func LoadGames() bool {
	return false
} // LoadGames


func LoadCache() {

	if checkStorage() {

		LoadTeams()
		
	}

} // LoadCache
