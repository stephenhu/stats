package stats

import (
  "fmt"
)


type NbaLegacyGame struct {
	ID 						string							`json:"gid"`
	Code          string              `json:"gcode"`
	State 				string							`json:"stt"`
}

type NbaLegacyMonthlyGames struct {
	Games 				[]NbaLegacyGame									`json:"g"`
}

type NbaLegacyMonthlySchedule struct {
	Mscd 					NbaLegacyMonthlyGames						`json:"mscd"`
}

type NbaLegacySchedule struct {
	Lscd 					[]NbaLegacyMonthlySchedule			`json:"lscd"`
}


func LegacyScheduleApi(y string) string {
	return fmt.Sprintf("%s%s", NBA_LEGACY_BASE_URL, fmt.Sprintf(
		NBA_LEGACY_API_SCHEDULE, y))
} // LegacyScheduleApi


func NbaGetLegacySchedule(y string) *NbaLegacySchedule {

	sched := NbaLegacySchedule{}

	apiInvoke(LegacyScheduleApi(y), &sched)

	return &sched

} // NbaGetLegacySchedule


func NbaGetLegacyScheduleJson(y string) []byte {
  return apiInvokeJson(LegacyScheduleApi(y))
} // NbaGetLegacyScheduleJson
