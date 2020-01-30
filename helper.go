package stats

import (
	"errors"
	"fmt"
	"log"
	"path"
	"net/url"
	"strconv"
	"strings"
)


func logf(fname string, msg string) {
  log.Printf("%s(): %s", fname, msg)
} // logf


func atoi(s string) int {

	if s == "" {
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


func fieldGoals(p *Player, fn string, s string) {

	if s == "" {
		logf("fieldGoals", fmt.Sprintf("Empty string for %s", fn))
		return
	}

	tokens := strings.Split(s, MINUS_MATCH)

	if len(tokens) != 2 {
		logf("fieldsGoals", "field format is incorrect")
	} else {

		made     := atoi(tokens[INDEX_MADE])
		attempts := atoi(tokens[INDEX_ATTEMPTS])
		
		switch fn {
		case FIELD_FG:
			p.Fga = attempts
			p.Fgm = made

		case FIELD_FG3:
			p.Fg3a = attempts
			p.Fg3m = made

		case FIELD_FT:
			p.Fta = attempts
			p.Ftm = made
		
		default:
			logf("fieldGoals", "Unrecognized field name.")
		}

	}

} // fieldGoals
