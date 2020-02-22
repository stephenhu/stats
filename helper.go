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

					_, ok := official_seasons[current]

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


func seasonKeyByDate(d string) string {

	if d == "" {
		return SEASON_CURRENT
	}

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("seasonKeyByDate", err.Error())
		return SEASON_CURRENT
	} else {

		cm := t.Month()

		if cm >= time.October && cm <= time.December {
			return seasonKey(t, true)
		} else if cm >= time.January && cm <= time.June {
			return seasonKey(t, false)
		} else {
			return SEASON_CURRENT
		}

	}

} // seasonKeyByDate


func generatePath(d string) string {

	if d == "" || !seasonCheck(d) {
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


func getSeason(t time.Time) []string {

	now := time.Now()

	if t.After(now) {
		logf("getSeason", fmt.Sprintf("Date unsupported: %s", t.String()))
		return official_seasons[SEASON_CURRENT]
	}

	cm := t.Month()

	if cm >= time.October && cm <= time.December {
		return official_seasons[seasonKey(t, true)]
	} else if cm >= time.January && cm <= time.June {
		return official_seasons[seasonKey(t, false)]
	} else {
		return official_seasons[SEASON_CURRENT]
	}

} // getSeason


func seasonCheck(d string) bool {

	if d == "" {
		return false
	}

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("seasonCheck", err.Error())
		return false
	} else {

		season := getSeason(t)

		begin, err := time.Parse(DATE_FORMAT,
			season[SEASON_INDEX_BEGIN])

		if err != nil {
			logf("seasonCheck", err.Error())
			return false
		} else {

			end, err 	:= time.Parse(DATE_FORMAT,
				season[SEASON_INDEX_PLAYOFFS_END])

			if err != nil {

				logf("seasonCheck", err.Error())

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

} // seasonCheck


func dateCheck(d string) bool {

	if d == "" || !seasonCheck(d) {
		return false
	} else {
		return true
	}

} // dateCheck


func getDays(d string) []string {

	days := []string{}

	if !seasonCheck(d) {
		logf("getDays", "Invalid date, out of season schedule.")
		return days
	} else {

		t, err := time.Parse(DATE_FORMAT, d)

		if err != nil {
			logf("getDays", err.Error())
			return days
		} else {

			season := getSeason(t)

			end, err := time.Parse(DATE_FORMAT, season[SEASON_INDEX_PLAYOFFS_END])

			if err != nil {
				logf("getDays", err.Error())
				return days
			} else {

				now := time.Now()

				tn := t

				for {

					if tn.After(end) || tn.After(now) {
						break
					} else {
						days = append(days, tn.Format(DATE_FORMAT))
					}

					tn = tn.AddDate(0, 0, 1)

				}

				return days

			}

		}

	}

} // getDays


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


func moveDay(d string, n int) string {

	t, err := time.Parse(DATE_FORMAT, d)

	if err != nil {
		logf("moveDay", err.Error())
		return d
	} else {

		tn := t.AddDate(0, 0, n)

		return tn.Format(DATE_FORMAT)

	}

} // moveDay


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


func LatestScoreboardDate() string {

	now := GetEstNow()

	if now == nil {
		logf("LatestScoreboardDate", "Failed to get current time.")
		return ""
	} else {

		start := GetEstDate(now)

		if now.Before(*start) {

			yesterday := now.AddDate(0, 0, -1)

			return yesterday.Format(DATE_FORMAT)

		} else {
			return now.Format(DATE_FORMAT)
		}

	}

} // LatestScoreboardDate


func LastDownload() string {

  /*seasons := loadSeasons()

	last := seasons[len(seasons)-1]*/

	return ""

} // LastDownload
