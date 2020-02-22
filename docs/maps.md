# maps

internal representation of stats in memory for the current season.

map | key | example | description |
--- | --- | --- | --- | ---
GamesMap | date | 20200219 | contains an array of games for that day
PlayersMap | player name, date | Kawhi Leonard, 20200219 | player game stats
TeamsMap | team name, date | lac, 20200219 | team game stats

maybe since games already contain all the player info, there only needs to be pointers to the games.

games data is the most empirical set of data, no need to store duplicate data.

this is how i'd like to access data:

GamesMap["20200219"]["lac.lal"]["lac"]["Kawhi Leonard"]
PlayersMap["Kawhi Leonard"]["20200219"]

GamesMap["20200219"] []Games