package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


func fileExists(fpath string) bool {

	if len(fpath) != 0 {

		_, err := os.Stat(fpath)

		if err != nil || os.IsNotExist(err) {
			log.Println(err)
			return false
		} else {
			return true
		}

	} else {
		return false
	}
	
} // fileExists


func apiInvoke(u string, data interface{}) {

	res, err := client.Get(u)

	if err != nil {
		logf("NbaGetSchedule", err.Error())
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Println(err)
		} else {

			err := json.Unmarshal(buf, &data)

			if err != nil {
				log.Println(err)
			}

		}

	}

} // apiInvoke


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


func getSeasonNow() string {

	current := time.Now()

	month := current.Month()
	year 	:= current.Year()

	if month >= time.October {
		return fmt.Sprintf("%d", year)
	} else {
		return fmt.Sprintf("%d", year - 1)
	}

} // getSeasonNow
