# design 2

* pull season schedule
  * pull boxscores per day from schedule, store to disk
  * if game not played yet then hold
  * check if game is "final" before storing to disk
* pull resume - basically starts pulling from where things left off
  * probably need to have a mechanism to figure out where things were stopped, some days could have partial amounts of days completed
* check current date, either pull single day or all days available


## in memory structure

Season
  Games
  Players (can see stats per game)
  Teams (see stats per game, per player)
  Leaderboard
  Standings


### players

* show player info
* show player per game


### teams

* show team info
* show per game
