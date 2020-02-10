# design

crawl data from various sources to provide a comprehensive set of normalized 
nba data from the oldest player in the current nba season (e.g. 2019-20 vince 
carter).  this data should be stored in json format which is easily readable
and distributable.

a secondary, but not absolutely must have goal would be to get data from the 
larry bird era which is defined as the season starting from his rookie year in 
1978.  Important to note that ideally 1979 would be the better year to start 
from because this was the first year that the 3 point line was introduced, but 
for personal reasons, larry bird was one of the greatest players that i grew 
up watching.

a third goal would be to get the entire nba/aba history, this would be quite 
painful as there were many changes in terms of teams, cities, leagues, and
possibly rule changes.

the ultimate goal of this framework is to provide a set of analysis libraries
on the collected data which could be used for predictive analysis, trending,

## sources

the most rudimentary data set are game stats, with game stats, all data such as
season, team, or career stats could be calculated from this data.  this is the 
most empirical set of data.  implicitly this set of data includes roster 
information especially after trades, team standings and stats could also be 
derived.  game data doesn't change after all data is confirmed, data such 
as season player stats are rolling and updated after each game so doesn't
need a snapshot or point in time of the data, just the most current
information.

there's an ancillary set of data that also needs to be maintained such as
player and team information.

and more advanced data like play by play and sportsvu, etc data can also
enhance analysis.

gameid is different from espn and data.nba

## crawler

there needs to be a method to account for data download completion


## storage

in general, the idea is to save all non-transient data to the filesystem as
a set of normalized json files that can be redistributable or loaded into a 
system for analysis and calculation.  most files are directly relative to
a season, for example, lebron.james.json career stats are relative
to the current season, so previous seasons will not have his career stats
updated.  _suggest only storing non-transient data_ 

data is cleaned, aggregated, and stored to the local filesystem:

```
nba
  players.json
  teams.json
  2019   # 2019-20 season
    20200203
      rockets.blazers.json    # static or updated during game time
    players
      lebron.james.json       # should update after each game      
    teams
      clippers.json
      teams.json
  teams
    teams.json
```

## analysis

1.  player season stats
1.  player career stats
1.  player team stats
1.  team season stats
1.  team career stats
1.  last 5 games player statistics
