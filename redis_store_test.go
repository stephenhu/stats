package stats

import (
	"fmt"
	"testing"
)


func init() {
	ConnectRedis("tcp", ":6379")
} // init


func TestRedisStorePlayers(t *testing.T) {

	t.Skip("removing redis")

	RedisStorePlayers("2019")

} // TestRedisStorePlayers


func TestRedisStoreTeams(t *testing.T) {

	t.Skip("removing redis")

	RedisStoreTeams("2019")

} // TestRedisStoreTeams


func TestRedisStoreSeason(t *testing.T) {

	t.Skip("removing redis")
	RedisStoreSeason(2019)

} // TestRedisStoreSeason


func TestRedisLastGame(t *testing.T) {

	t.Skip("removing redis")

	d := RedisLastGame()

	t.Log(d)

} // TestRedisLastGame


func TestRedisGameDays(t *testing.T) {

	t.Skip("removing redis")

	dates := RedisGameDays(2019)

	t.Log(dates)

} // TestRedisGameDays


func TestRedisStoreGamesFrom(t *testing.T) {

	t.Skip("removing redis")
	RedisStoreGamesFrom("20200201")

} // TestRedisStoreGamesFrom


func TestRedisGames(t *testing.T) {

	t.Skip("removing redis")

	keys := RedisGames(2019)

	t.Log(keys)
	t.Log(len(keys))

} // TestRedisGames


func TestRedisSeasons(t *testing.T) {

	t.Skip("removing redis")

	keys := RedisSeasons()

	t.Log(keys)
	t.Log(len(keys))

} // TestRedisSeasons


func TestRedisGetTeamData(t *testing.T) {

	t.Skip("removing redis")
	td := RedisGetTeamData(fmt.Sprintf("%d", CurrentSeason()), "hou")

	t.Log(td)

} // TestRedisGetTeamData
