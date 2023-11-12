package stats

import (
  "fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
)


type NbaScheduleMeta struct {
	Version 					int					`json:"version"`
	Request 					string			`json:"request"`
	Time 							string			`json:"time"`
}


type NbaGameDates struct {
	GameDate          string      `json:"gameDate"`
	Games          		[]NbaGame   `json:"games"`
}


type NbaLeagueSchedule struct {
  SeasonYear 				string			`json:"seasonYear"`
	LeagueId          string      `json:"leagueId"`
	GameDates         []NbaGameDates     `json:"gameDates"`
}


type NbaSchedule struct {
	Meta   						NbaScheduleMeta			`json:"meta"`
  LeagueSchedule    NbaLeagueSchedule   `json:"leagueSchedule"`
}


func ScheduleApi(d string) string {

	return fmt.Sprintf("%s%s%s",
    NBA_BASE_URL,
	  NBA_STATIC,
		NBA_SCHEDULE,
	)

} // ScheduleApi


func NbaGetSchedule(d string) {

	if len(d) == 0 {
		return
	}

	s := fmt.Sprintf(NBA_SCHEDULE_FILE, d)

	if !fileExists(s) {

		res, err := client.Get(ScheduleApi(d))

		if err != nil {
			logf("NbaGetSchedule", err.Error())
			//return nil
		} else {

			defer res.Body.Close()

			buf, err := ioutil.ReadAll(res.Body)

			if err != nil {
				logf("NbaGetSchedule", err.Error())
				//return nil
			} else {

				f, err := os.Create(s)

				if err != nil {
					log.Println(err)
				} else {

					_, err := f.Write(buf)

					if err != nil {
						log.Println(err)
					}

				}

			}

		} 
		
	} else {

			buf, err := ioutil.ReadFile(s)

			if err != nil {
				log.Println(err)
			} else {

				schedule := NbaSchedule{}

				err := json.Unmarshal(buf, &schedule)

				if err != nil {
					log.Println(err)
				} else {
					log.Println(schedule)
				}

			}

		}

} // NbaGetSchedule
