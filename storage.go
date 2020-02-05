package stats

import (
	"fmt"	
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


func put(g *Game, buf []byte) {

	dir := generatePath(g.Date)

	root := initStorage(dir)

	f := filepath.Join(root, fmt.Sprintf(
		"%s.%s.json", g.Away.Name, g.Home.Name))

	fh, err := os.Create(strings.ToLower(f))

	if err != nil {
		logf("put", err.Error())
	} else {

		defer fh.Close()

		fh.Write(buf)

	}

} // put
