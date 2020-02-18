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
		
		//logf("fileExists", err.Error())

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

		root := initStorage(all.SeasonID)
		
		f := filepath.Join(root, PLAYERS_FILE)

		j, err := json.MarshalIndent(all, JSON_PREFIX, JSON_INDENT)

		if err != nil {
			logf("putPlayers", err.Error())
		} else {
			put(f, j)
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
	} else {

		root := initStorage(teams.SeasonID)

		f := filepath.Join(root, TEAMS_FILE)

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
