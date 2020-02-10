package stats

import (
	"encoding/json"
	"fmt"
	//"log"
	"os"
	"path/filepath"
	"strings"
)


func initStorage(dir string) string {	

	f := filepath.Join(APP_STORAGE, dir)
	
	_, err := os.Stat(f)

	if err != nil {
		
		if os.IsNotExist(err) {
		
			logf("initStorage", err.Error())

			os.MkdirAll(f, 0755)

			return f

		} else {			
			return ""
		}

	}

	return f
	
} // initStorage


func put(f string, buf []byte) {
	
	fh, err := os.Create(strings.ToLower(f))

	if err != nil {
		logf("put", err.Error())
	} else {

		defer fh.Close()

		fh.Write(buf)

		fh.Sync()

	}

} // put


func putGame(g *Game) {

	dir := generatePath(g.Date)

	root := initStorage(dir)

	f := filepath.Join(root, fmt.Sprintf(
		"%s.%s.json", g.Away.Name, g.Home.Name))

	j, err := json.MarshalIndent(g, JSON_PREFIX, JSON_INDENT)

	if err != nil {
		logf("putGame", err.Error())
	} else {
		put(f, j)	
	}

} // putGame


func putPlayers(all *AllPlayers) {

	if all != nil {

		if all.SeasonID == "" {
			logf("putPlayers", "Failed to store players due to an empty season")
		} else {

			root := initStorage(all.SeasonID)
			
			f := filepath.Join(root, PLAYERS_FILE)

			j, err := json.MarshalIndent(all, JSON_PREFIX, JSON_INDENT)

			if err != nil {
				logf("putPlayers", err.Error())
			} else {
				put(f, j)
			}

		}

	} else {
		logf("putPlayers", "Failed to store nil players")
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

	if teams != nil {

		if teams.SeasonID == "" {
			logf("putTeams", "Failed to store teams due to empty seasonId")
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

	} else {
		logf("putTeams", "Failed to store teams, nil teams.")
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
		logf("putTeamRanks", "Failed to store team ranks, nil teams.")
	}

} // putTeamRanks
