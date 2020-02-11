package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"os"
	"path/filepath"
	"strings"
)


func fileExists(f string) bool {

	_, err := os.Stat(f)

	if err != nil {
		
		logf("fileExists", err.Error())

		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}

	} else {
		return true
	}

} // fileExists


func initStorage(dir string) string {	

	var f string

	if dir == "" {
		f = APP_STORAGE
	} else {
		f = filepath.Join(APP_STORAGE, dir)
	}
	
	if !fileExists(f) {
		os.MkdirAll(f, 0755)	
	}

	return f
	
} // initStorage


func checkStorage() bool {
	return fileExists(APP_STORAGE)
} // checkStorage


func loadFile(f string) []byte {	

	buf, err := ioutil.ReadFile(f)

	if err != nil {
		logf("loadFile", err.Error())
		return nil
	}

	return buf

} // loadFile


func mergeTeams(teams *AllTeams, f string) {

	buf := loadFile(f)

	if buf == nil {
		logf("mergeTeams", fmt.Sprintf("Unable to load %s", f))
	} else {

		oldTeams := AllTeams{}

		err := json.Unmarshal(buf, &oldTeams)

		if err != nil {
			logf("mergeTeams", err.Error())
		} else {

			newTeams := []TeamInfo{}

			for _, nt := range teams.Teams {

				found := false

				for _, ot := range oldTeams.Teams {

					if nt.ID == ot.ID {
						found = true
					}

				}

				if !found {
					newTeams = append(newTeams, nt)
				}

			}

			oldTeams.Teams = append(oldTeams.Teams, newTeams...)

			j, err := json.MarshalIndent(oldTeams, JSON_PREFIX, JSON_INDENT)

			if err != nil {
				logf("mergeTeams", err.Error())
			} else {
				put(f, j)
			}

		}

	}

} // mergeTeams


func mergePlayers(players *AllPlayers, f string) {

	buf := loadFile(f)

	if buf == nil {
		logf("mergePlayers", fmt.Sprintf("Unable to load %s", f))
	} else {

		oldPlayers := AllPlayers{}

		err := json.Unmarshal(buf, &oldPlayers)

		if err != nil {
			logf("mergePlayers", err.Error())
		} else {

			newPlayers := []PlayerInfo{}

			for _, np := range players.Players {

				found := false

				for _, op := range oldPlayers.Players {

					if np.ID == op.ID {
						found = true
					}

				}

				if !found {
					newPlayers = append(newPlayers, np)
				}

			}

			oldPlayers.Players = append(oldPlayers.Players, newPlayers...)

			j, err := json.MarshalIndent(oldPlayers, JSON_PREFIX, JSON_INDENT)

			if err != nil {
				logf("mergePlayers", err.Error())
			} else {
				put(f, j)
			}

		}

	}

} // mergePlayers


func put(f string, buf []byte) {
	
	lf := strings.ToLower(f)

	fh, err := os.Create(lf)

	if err != nil {
		logf("put", err.Error())
	} else {

		defer fh.Close()

		fh.Write(buf)

		fh.Sync()

		logf("INFO", fmt.Sprintf("%s created successfully.", lf))

	}

} // put


func putGame(g *Game) {

	dir := generatePath(g.Date)

	root := initStorage(dir)

	f := filepath.Join(root, fmt.Sprintf(
		GAME_FILE, g.Away.Name, g.Home.Name))

	j, err := json.MarshalIndent(g, JSON_PREFIX, JSON_INDENT)

	if err != nil {
		logf("putGame", err.Error())
	} else {
		put(f, j)	
	}

} // putGame


func putPlayers(all *AllPlayers) {

	if all != nil {

		root := initStorage("")
		
		f := filepath.Join(root, PLAYERS_FILE)

		if fileExists(f) {
			mergePlayers(all, f)
		} else {

			j, err := json.MarshalIndent(all, JSON_PREFIX, JSON_INDENT)

			if err != nil {
				logf("putPlayers", err.Error())
			} else {
				put(f, j)
			}
	
		}

	} else {
		logf("putPlayers", "Unable to store empty players")
	}

} // putPlayers


func putProfile(profile *PlayerCareer) {

	if profile != nil {
		
		root := initStorage(filepath.Join(profile.SeasonID, PLAYERS_DIR))		

		f := filepath.Join(root, fmt.Sprintf("%s.%s.json",
		  profile.First, profile.Last))

		j, err := json.MarshalIndent(profile, JSON_PREFIX, JSON_INDENT)

		if err != nil {
			logf("putPlayers", err.Error())
		} else {
			put(f, j)
		}

	} else {
		logf("putProfile", "Failed to store nil profile")
	}

} // putProfile


func putTeams(teams *AllTeams) {

	if teams == nil {
		logf("putTeams", "Unable to store empty teams")
		return
	}

	root := initStorage("")

	f := filepath.Join(root, TEAMS_FILE)

	if fileExists(f) {
		mergeTeams(teams, f)
	} else {

		j, err := json.MarshalIndent(teams, JSON_PREFIX, JSON_INDENT)

		if err != nil {
			logf("putTeams", err.Error())
		} else {
			put(f, j)
		}

	}

} // putTeams


func putTeamRanks(ranks *AllRanks) {

	if ranks != nil {

		if ranks.SeasonID == "" {
			logf("putTeamRanks", "Failed to store team ranks due to empty seasonId")
		} else {

			root := initStorage(ranks.SeasonID)
			
			f := filepath.Join(root, TEAM_RANKS_FILE)

			j, err := json.MarshalIndent(ranks, JSON_PREFIX, JSON_INDENT)

			if err != nil {
				logf("putTeamRanks", err.Error())
			} else {
				put(f, j)
			}

		}

	} else {
		logf("putTeamRanks", "Failed to store empty team ranks.")
	}

} // putTeamRanks
