data.nba.net

https://github.com/kashav/nba.js/blob/master/docs/api/DATA.md

* 
* only supports data since the 2015 season, will need to use espn for older games
* players from different leagues are interspersed with nba players such
as cba, turkish league, etc, filtering out these players requires a flag
defined as `isActive` or to check if the `teamId` is an NBA team
* teams from different leagues are also interspersed, but there's a
flag called `isNBAFranchise` to identify the nba teams
* 2015, 2016 do not have the IsActive flag for players, will need to
filter based on team data
* 