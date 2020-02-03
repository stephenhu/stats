package stats

import (	
	"errors"
	"fmt"
	"log"
	"path"
	"net/url"
	"strconv"
	"time"
)


func logf(fname string, msg string) {
  log.Printf("%s(): %s", fname, msg)
} // logf


func atoi(s string) int {

	if s == STRING_EMPTY {
		logf("Atoi", "Empty string.")
		return 0
	}

	val, err := strconv.ParseInt(s, BASE10, BITS32)
	
	if err != nil {
		logf("Atoi", fmt.Sprintf("Field: %s, %s", s, err.Error()))
		return 0
	} else {
		return int(val)
	}

} // atoi


func seasonKey(t time.Time, current bool) string {

	ty := t.Format(YEAR_FORMAT)

	if current {

		tn := t.AddDate(1, 0, 0)

		return fmt.Sprintf("%s%s", ty,
	    tn.Format(YEAR_FORMAT))

	} else {

		tp := t.AddDate(-1, 0, 0)

		return fmt.Sprintf("%s%s", tp.Format(YEAR_FORMAT),
			ty)
			
	}

} // seasonKey


func getSeason(t time.Time) []string {

	cm := t.Month()

	if cm >= time.October && cm <= time.December {
		return seasons[seasonKey(t, true)]
	} else if cm >= time.January && cm <= time.June {
		return seasons[seasonKey(t, false)]
	} else {
		return seasons["1920"]
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

		begin, err := time.Parse(DATE_FORMAT, season[SEASON_INDEX_BEGIN])

		if err != nil {
			logf("seasonCheck", err.Error())
			return false
		} else {

			end, err 	:= time.Parse(DATE_FORMAT, season[SEASON_INDEX_PLAYOFFS_END])

			if err != nil {
				
				logf("seasonCheck", err.Error())
				
				end = time.Now()					
				
			}

			if (t.After(begin) || t.Equal(begin)) && (t.Equal(end) || t.Before(end)) {
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

			now := time.Now()

			tn := t

			for {
				
				if tn.After(now) {
					break
				} else {
					days = append(days, tn.Format(DATE_FORMAT))
				}

				tn = tn.AddDate(0, 0, 1)

			}
	
			return days

		}

	}

} // getDays


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
