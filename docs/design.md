# design

crawl data primarily from data.nba.net and store to the filesystem.

_larry bird era 1978_

_1979 is the first year for the 3 point line_


## storage

in general, the idea is to save all non-transient data to the filesystem as
a set of normalized json files that can be redistributable or loaded into a 
system for analysis and calculation.  most files are directly relative to
a season, for example, lebron.james.json career stats are relative
to the current season, so previous seasons will not have his career stats
updated.

data is cleaned, aggregated, and stored to the local filesystem:

```
nba
  2019
    20200203
      rockers.blazers.json    # static or updated during game time
    players
      lebron.james.json       # should update after each game
      players.json            # updates if new players added from
    teams
      clippers.json
      teams.json
  teams    
    teams.json
```

## analysis

1.  last 5 games player statistics

