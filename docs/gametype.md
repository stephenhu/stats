# game type

there are 6 game types for an nba season, for analytics and other purposes,
this information needs to be made clear and the nba data is not very clean
in this regard:

1. preseason
1. regular
1. in-season tournament
1. all star game
1. play-in
1. playoffs

## analytics

for the leaders, games, players, and standings, only the leaders
parquet needs to be differentiated at least for regular season and
playoffs.

1. timestamp.players.parquet # contains all game type data
1. timestamp.games.parquet # contains all game type data
1. timestamp.regular.leaders.parquet # contains regular season leader stats
1. timestamp.playoff.leaders.parquet # contains playoff leader stats, not counting play in

## preseason

```
"weekNumber": 0,
"weekName": "",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "Preseason",
"gameSubLabel": "",
"seriesText": "",
```

## in-season tournament

```
"weekNumber": 2,
"weekName": "Week 2",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "In-Season Tournament",
"gameSubLabel": "",
"seriesText": "West Group C",
```

## regular season

```
"weekNumber": 5,
"weekName": "Week 5",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "",
"gameSubLabel": "",
"seriesText": "",
```


## all star

```
"weekNumber": 17,
"weekName": "All-Star",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "All-Star Game",
"gameSubLabel": "",
"seriesText": "",
```

```
"weekNumber": 17,
"weekName": "All-Star",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "Rising Stars Final",
"gameSubLabel": "Championship",
"seriesText": "",
```

```
"weekNumber": 17,
"weekName": "All-Star",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "Rising Stars Semifinal",
"gameSubLabel": "Game 2",
"seriesText": "",
```

## play-in

```
"weekNumber": 26,
"weekName": "Play-In",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "Play-In",
"gameSubLabel": "",
"seriesText": "",
```

## playoffs

```
"weekNumber": 0,
"weekName": "",
"ifNecessary": false,
"seriesGameNumber": "Game 1",
"gameLabel": "East - First Round",
"gameSubLabel": "Game 1",
"seriesText": "Series tied 0-0",
```

## games played globally

~these should count as regular season games, note the gameLabel~

```
"weekNumber": 12,
"weekName": "Week 12",
"ifNecessary": false,
"seriesGameNumber": "",
"gameLabel": "NBA Paris Game",
"gameSubLabel": "",
"seriesText": "",
```
