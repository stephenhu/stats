package stats_test

import (
	"testing"

	"github.com/stephenhu/stats"
)


func init() {
	stats.ConnectRedis("tcp", ":6379")
} // init


func TestRedisStorePlayers(t *testing.T) {

	stats.RedisStorePlayers("2019")

} // TestRedisStorePlayers


func TestRedisStoreTeams(t *testing.T) {

	stats.RedisStoreTeams("2019")

} // TestRedisStoreTeams


func TestRedisStoreSeason(t *testing.T) {

	stats.RedisStoreSeason("2019")

} // TestRedisStoreSeason


func TestRedisLastGame(t *testing.T) {

	d := stats.RedisLastGame()

	t.Log(d)

} // TestRedisLastGame


func TestRedisGameDays(t *testing.T) {

	dates := stats.RedisGameDays("2019")

	t.Log(dates)

} // TestRedisGameDays


func TestRedisStoreGamesFrom(t *testing.T) {

	stats.RedisStoreGamesFrom("20200201")

} // TestRedisStoreGamesFrom


func TestRedisGames(t *testing.T) {

	keys := stats.RedisGames("2019")

	t.Log(keys)
	t.Log(len(keys))

} // TestRedisGames


func TestRedisSeasons(t *testing.T) {

	keys := stats.RedisSeasons()

	t.Log(keys)
	t.Log(len(keys))

} // TestRedisSeasons


func TestRedisGetTeamData(t *testing.T) {

	td := stats.RedisGetTeamData(stats.CurrentSeason(), "hou")

	t.Log(td)

} // TestRedisGetTeamData
