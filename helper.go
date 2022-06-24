package stats

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"net/url"
	"strconv"
	"strings"
	"time"
)


func logf(fname string, msg string) {
  log.Printf("%s(): %s", fname, msg)
} // logf


func atoi(s string) int {

	if s == STRING_EMPTY || s == STRING_MINUS {
		return 0
	}

	val, err := strconv.ParseInt(s, BASE10, BITS32)

	if err != nil {
		logf("atoi", fmt.Sprintf("Field: %s, %s", s, err.Error()))
		return 0
	} else {
		return int(val)
	}

} // atoi


func atof(s string) float32 {

	if s == STRING_EMPTY {
		return 0.0
	}

	val, err := strconv.ParseFloat(s, BITS32)

	if err != nil {
		logf("atof", fmt.Sprintf("Field: %s, %s", s, err.Error()))
		return 0.0
	} else {
		return float32(val)
	}

} // atof


func getYearsFrom(s string) []string {

	years := []string{}

	if s != "" {

		t, err := time.Parse(YEAR_FORMAT, s)

		if err != nil {
			logf("getYearsFrom", err.Error())
			return nil
		} else {

			now  := time.Now()

			endYear := now.Year()

			tn := t

			for {

				year := tn.Year()

				if year > endYear {
					break
				} else {

					current := tn.Format(YEAR_FORMAT)

					_, ok := OfficialSeasons[current]

					if ok {
						years = append(years, current)
					}

				}

				tn = tn.AddDate(1, 0, 0)

			}

			return years

		}

	} else {
		logf("getYearsFrom", "Cannot process empty string")
		return nil
	}

} // getYearsFrom


func seasonKey(t time.Time, current bool) string {

	if current {
		return t.Format(YEAR_FORMAT)
	} else {

		tp := t.AddDate(-1, 0, 0)

		return tp.Format(YEAR_FORMAT)

	}

} // seasonKey


func generatePath(d string) string {

	if d == "" || !SeasonCheck(d) {
		return filepath.Join(SEASON_DEFAULT, d)
	}

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("generatePath", err.Error())
		return filepath.Join(SEASON_DEFAULT, d)
	} else {

		cm := t.Month()

		if cm >= time.October && cm <= time.December {
			return filepath.Join(seasonKey(t, true), d)
		} else if cm >= time.January && cm <= time.June {
			return filepath.Join(seasonKey(t, false), d)
		} else {
			return filepath.Join(SEASON_DEFAULT, d)
		}

	}


} // generatePath


func dateCheck(d string) bool {

	if d == "" || !SeasonCheck(d) {
		return false
	} else {
		return true
	}

} // dateCheck


func mtoi(s string) (int, int) {

	toks := strings.Split(s, STRING_COLON)

	tokenLen := len(toks)

	if tokenLen == 2 {
		return atoi(toks[0]), atoi(toks[1])
	} else if tokenLen == 1 {
		return atoi(toks[0]), 0
	} else {
		logf("mtoi", fmt.Sprintf("Unknown minutes format: %s", s))
		return 0, 0
	}

} // mtoi


func filterId(id string) string {

	if id == STRING_EMPTY {
		return STRING_EMPTY
	}

	if strings.Contains(id, STRING_SPACE) {

		toks := strings.Split(id, STRING_SPACE)

		return strings.TrimSpace(toks[0])

	} else {
		return id
	}

} // filterId


func CurrentSeason() string {

	now := GetEstNow()

	m := now.Month()

	if m >= time.October && m <= time.December {
		return fmt.Sprintf("%d", now.Year())
	} else {

		tn := *now
		tn = tn.AddDate(-1, 0, 0)

		return fmt.Sprintf("%d", tn.Year())

	}

} // CurrentSeason


func GetGameType(d string) string {

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("GetGameType", err.Error())
		return SEASON_UNKNOWN
	} else {

		s := GetSeason(t)

		if d >= s[SEASON_BEGIN] && d <= s[SEASON_END] {
			return SEASON_REGULAR
		} else if d >= s[SEASON_PLAYOFFS_BEGIN] && d <= s[SEASON_PLAYOFFS_END] {
			return SEASON_PLAYOFF
		} else if d == s[SEASON_ALL_STAR_GAME] || d == s[SEASON_ALL_STAR_EMPTY] ||
		  d == s[SEASON_WORLD_GAME] {
			return SEASON_ALLSTAR
		} else {
			return SEASON_UNKNOWN
		}

	}

} // GetGameType


func GetDays(d string) []string {

	days := []string{}

	if !SeasonCheck(d) {
		logf("GetDays", "Invalid date, out of season schedule.")
	} else {

		fromDate := ParseEstTime(d)

		if fromDate != nil {

			season := GetSeason(*fromDate)

			end := ParseEstTime(season[SEASON_PLAYOFFS_END])

			if end != nil {

				now := DayStartTime()

				tmp := fromDate.AddDate(0, 0, 1)

				tn := &tmp

				for {

					if tn.After(*end) || tn.Equal(*end) || tn.After(*now) ||
					  tn.Equal(*now) {
						break
					} else {

						if tn.Year() == now.Year() && tn.Month() == now.Month() && tn.Day() == now.Day() {
							break
						} else {
							days = append(days, tn.Format(DATE_FORMAT))
						}

					}

					next := tn.AddDate(0, 0, 1)

					tn = &next

				}

				return days

			}

		}

	}

	return days

} // GetDays


func GetDaysBySeason(y string) []string {

	days := []string{}

	if len(y) == 0 {
		logf("GetDaysBySeason", "Invalid season provided " + y)
	} else {

		s := OfficialSeasons[y]

		days = GetDays(s[SEASON_BEGIN])
		
	}

	return days

} // GetDaysBySeason


func MoveDay(d string, n int) string {

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("MoveDay", err.Error())
		return d
	} else {

		tn := t.AddDate(0, 0, n)

		return tn.Format(DATE_FORMAT)

	}

} // MoveDay


func StringUrlJoin(base string, p string) (string, error) {

	u, err := url.Parse(base)

	if err != nil {

		logf("StringUrlJoin", err.Error())
		return "", errors.New(fmt.Sprintf("Unable to join strings %s and %s.",
	    base, p))

	} else {

		u.Path = path.Join(u.Path, p)

		return u.String(), nil

	}

} // StringUrlJoin


func DayStartTime() *time.Time {

	now := GetEstNow()

	day := now.Weekday()

	var t time.Time

	if day == time.Saturday || day == time.Sunday {
		// 3pm est
		t = time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, now.Location())
	} else {
		// 5pm est
		t = time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())
	}

	return &t

} // DayStartTime


func GetEstNow() *time.Time {

	est, err := time.LoadLocation(EST)

	if err != nil {
		logf("GetEstNow", err.Error())
		return nil
	} else {

		now := time.Now().In(est)

		return &now

	}

} // GetEstNow


func ParseEstTime(d string) *time.Time {

	if d == "" || len(d) != 8 {
		return nil
	}

	year := atoi(d[0:4])
	mon  := atoi(d[4:6])
	day  := atoi(d[6:8])

	est, err := time.LoadLocation(EST)

	if err != nil {
		logf("ParseEstTime", err.Error())
		return nil
	} else {

		t := time.Date(year, time.Month(mon), day, 0, 0, 0, 0, est)

		return &t

	}

} // ParseEstTime


func GetEstDate(now *time.Time) *time.Time {

	day := now.Weekday()

	st := START_TIME_WEEKDAY

	if day == time.Saturday || day == time.Sunday {
		st = START_TIME_WEEKEND
	}

	est, err := time.LoadLocation(EST)

	if err != nil {
		logf("GetEstDate", err.Error())
		return now
	} else {

		start := time.Date(now.Year(), now.Month(), now.Day(), st, 0, 0, 0, est)

		return &start

	}

} // GetEstDate


func SeasonKeyByDate(d string) string {

	if d == "" {
		return CurrentSeason()
	}

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("SeasonKeyByDate", err.Error())
		return CurrentSeason()
	} else {

		cm := t.Month()

		if cm >= time.October && cm <= time.December {
			return seasonKey(t, true)
		} else if cm >= time.January && cm <= time.June {
			return seasonKey(t, false)
		} else {
			return CurrentSeason()
		}

	}

} // SeasonKeyByDate


func SeasonCheck(d string) bool {

	if d == "" {
		return false
	}

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("SeasonCheck", err.Error())
		return false
	} else {

		season := GetSeason(t)

		begin, err := time.Parse(DATE_FORMAT,
			season[SEASON_BEGIN])

		if err != nil {
			logf("SeasonCheck", err.Error())
			return false
		} else {

			end, err 	:= time.Parse(DATE_FORMAT,
				season[SEASON_PLAYOFFS_END])

			if err != nil {

				logf("SeasonCheck", err.Error())

				end = time.Now()

			}

			if (t.After(begin) || t.Equal(begin)) &&
			  (t.Equal(end) || t.Before(end)) {
				return true
			} else {
				return false
			}

		}

	}

} // SeasonCheck


func GetSeason(t time.Time) []string {

	now := time.Now()

	if t.After(now) {
		logf("GetSeason", fmt.Sprintf("Date unsupported: %s", t.String()))
		return OfficialSeasons[CurrentSeason()]
	}

	cm := t.Month()

	if cm >= time.October && cm <= time.December {
		return OfficialSeasons[seasonKey(t, true)]
	} else if cm >= time.January && cm <= time.June {
		return OfficialSeasons[seasonKey(t, false)]
	} else {
		return OfficialSeasons[CurrentSeason()]
	}

} // GetSeason
