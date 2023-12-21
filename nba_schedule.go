package stats

import (
  "fmt"
)


type NbaMeta struct {
	Version 					int					`json:"version"`
	Request 					string			`json:"request"`
	Time 							string			`json:"time"`
}


type NbaGameDates struct {
	GameDate          string      `json:"gameDate"`			// 10/05/2023 00:00:00
	Games          		[]NbaGame   `json:"games"`
}


type NbaWeek struct {
  WeekNumber				int					`json:"weekNumber"`
	StartDate					string			`json:"startDate"`		// UTC
	EndDate						string			`json:"endDate"`			// UTC
}


type NbaLeagueSchedule struct {
  SeasonYear 				string			`json:"seasonYear"`
	LeagueId          string      `json:"leagueId"`
	GameDates         []NbaGameDates     `json:"gameDates"`
	Weeks             []NbaWeek           `json:"weeks"`
}





type NbaSchedule struct {
	Meta   						NbaMeta							`json:"meta"`
  LeagueSchedule    NbaLeagueSchedule   `json:"leagueSchedule"`
}


func ScheduleApi() string {

	return fmt.Sprintf("%s%s%s",
    NBA_BASE_URL,
	  NBA_STATIC,
		NBA_SCHEDULE,
	)

} // ScheduleApi


func NbaGetSchedule() *NbaSchedule {

	schedule := NbaSchedule{}

	apiInvoke(ScheduleApi(), &schedule)

	return &schedule

} // NbaGetSchedule


func NbaGetScheduleJson() []byte {
	return apiInvokeJson(ScheduleApi())
} // NbaGetScheduleJson
