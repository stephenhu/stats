data.nba.net

https://github.com/kashav/nba.js/blob/master/docs/api/DATA.md

http://data.nba.net/data/10s/prod/v1/current/standings_all.json
http://data.nba.net/data/10s/prod/v1/2019/players/200746_profile.json
http://data.nba.net/data/10s/prod/v1/2019/players.json
https://data.nba.net/data/10s/prod/v1/2019/team_stats_rankings.json


*
* only supports data since the 2015 season, will need to use espn for older games
* players from different leagues are interspersed with nba players such
as cba, turkish league, etc, filtering out these players requires a flag
defined as `isActive` or to check if the `teamId` is an NBA team
* teams from different leagues are also interspersed, but there's a
flag called `isNBAFranchise` to identify the nba teams
* 2015, 2016 do not have the IsActive flag for players, will need to
filter based on team data
* player info is relative to a season, data is very inconsistent across seasons, an example is 2015 denzel valentine is missing a lot of fields, there's also a field for years which gets incremented after each season.  i think the best thing to do is to keep player info per season.
* for some reason 2016 player data team id has 2 values delimited by space, it's the same 2 team id's, something tells me the nba is messing around with this data to make it harder to parse the data.  so many data inconsistencies, this makes it even more important to gather and normalize this data.
* the pbp api is strange in that you're getting a quarter's worth of data, there's no index for how many periods so you need parse this from the boxscore or you can
keep incrementing the period until you get a plays array equal to 0, but that
wastes an http request.
* uber_stats don't seem to work
