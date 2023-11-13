package stats

import (
  "fmt"
	"log"
)


type NbaScoreboardToday struct {
	Meta 							NbaMeta									`json:"meta"`
	Scoreboard        NbaScoreboard						`json:"scoreboard"`
}


func ScoreboardTodayApi() string {

	return fmt.Sprintf("%s%s%s",
    NBA_BASE_URL,
		NBA_LIVE,
		NBA_API_TODAYS_SCOREBOARD,
	)

} // ScoreboardTodayApi


func NbaGetScoreboardToday() *NbaScoreboardToday {

	today := NbaScoreboardToday{}

	log.Println(ScoreboardTodayApi())
	
	apiInvoke(ScoreboardTodayApi(), &today)

	return &today

} // NbaGetScoreboardToday
