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

and more advanced data like play by play, synergy, and sportsvu etc data can
also enhance analysis.

gameid is different from espn and data.nba

## normalization

since there are different sources and these sources have different formats, it's
important to normalize the data.

## crawler

there needs to be a method to account for data download completion, for now,
let's make this single threaded and simple, messages will be printed out to
console.

## persistent storage

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

## redis

data should be stored to redis for caching and another level of persistence
in addition to the file system.  this allows data to be persisted in cache even
if the app is restarted and allows for quick start up times.

## caching

in general, all data should be placed into cache for highest performance.  the caching
happens by crawling the filesystem and loading into memory.

* do teams need to be cached per season or can there be a general teams map for all seasons?
* same question for players, in general i think there should be a generic players map

1. seattle supersonics -> oklahoma thunder
1. charlotte hornets -> new orleans pelicans
1. charlotte bobcats
1. new jersey nets -> brooklyn nets
1. vancouver grizzlies -> memphis grizzlies
1. washington bullets -> washington wizards

maps | key | value | objects estimated | notes
--- | --- | --- | --- | ---
players | steven.adams | PlayerInfo | 450+ | players names that are the same will cause conflict
teams | 2019.clippers | TeamInfo | 30 |
games | 20191022 | []Game | 82 x 30 = 2460/2 + playoffs (8*7 + 4*7 + 2*7 + 7 = 105)
team season stats | 2019.atl | Stats | 30 |
player season stats | 2019.steven.adams | Stats | 450 |
play logs | 20191022.atl.bos | Playlog |  1230  + 105 = 1335 |

## real time

the nba data api doesn't offer a streaming data api so we'll have to simulate this
by polling the api, this amount should be configurable such that we don't flood
the service.

on the scoreboard, there's a point at which you should show the latest games versus
the previous day's games, i guess this should be based on the first game's time,
should be EST 5pm.  let's say 1h before we show the latest scores.  weekends will start earlier, maybe 12pm EST. seems like this seasons, most games are from 15:30
on the weekends, i guess it's to provide enough time for pst fans to watch.

so here's the update algorithm:

1.  get current date, day of week in est
1.  get scoreboard for current date, if 1 hour before, show today's games
1.  if > 1h before then show yesterday's games

## analysis

1.  player season stats
1.  player career stats
1.  player team stats
1.  team season stats
1.  team career stats
1.  last 5 games player statistics
