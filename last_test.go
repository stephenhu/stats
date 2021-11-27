package stats_test

import (
	"testing"

	"github.com/stephenhu/stats"
)


func init() {
	stats.ConnectRedis("tcp", ":6379")
}

func TestLastGames(t *testing.T) {

	t.Skip("changed to redis")
	//x := stats.LastGames(5, "2019", "lebronjames")

	//t.Log(x)

} // TestLastGames
