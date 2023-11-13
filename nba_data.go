package stats

import (
	"fmt"
	"strings"
)




func convScores(scores []NbaScoreData) []int {

	ret := []int{}

	for _, s := range scores {
		ret = append(ret, s.Score)
	}

	return ret

} // convScores


func convPlayers(players []NbaPlayerData, g *Game) {

	awayPlayers := []Player{}
	homePlayers := []Player{}

	for _, p := range players {

		player := Player{}

		minutes, seconds := mtoi(p.Minutes)

		player.ID 					= p.ID
		player.Name 				= fmt.Sprintf("%s %s", p.First, p.Last)
		player.Points 			= atoi(p.Points)
		player.Minutes 			= minutes
		player.Seconds      = seconds
		player.Fgm 					= atoi(p.Fgm)
		player.Fga 					= atoi(p.Fga)
		player.Fg3m 				= atoi(p.Fg3m)
		player.Fg3a 				= atoi(p.Fg3a)
		player.Ftm 					= atoi(p.Ftm)
		player.Fta 					= atoi(p.Fta)
		player.Oreb					= atoi(p.Oreb)
		player.Dreb 				= atoi(p.Dreb)
		player.Treb 				= atoi(p.Treb)
		player.Assists 			= atoi(p.Assists)
		player.Steals 			= atoi(p.Steals)
		player.Blocks 			= atoi(p.Blocks)
		player.Turnovers 		= atoi(p.Turnovers)
		player.Fouls 				= atoi(p.Fouls)
		player.PlusMinus 		= atoi(p.PlusMinus)
		player.DnpReason 		= p.DnpReason
		player.Position     = p.Position

		if p.TeamID == g.Away.ID {
			player.Opponent = g.Home.Name
			awayPlayers = append(awayPlayers, player)
		} else {
			player.Opponent = g.Away.Name
			homePlayers = append(homePlayers, player)
		}

	}

	g.Away.Players = awayPlayers
	g.Home.Players = homePlayers

} // convPlayers


func convTeam(ts *NbaTeam, td *Team) {

	td.Points				= atoi(ts.Points)
	td.Fgm					= atoi(ts.Fgm)
	td.Fga					= atoi(ts.Fga)
	td.Fg3m					= atoi(ts.Fg3m)
	td.Fg3a					= atoi(ts.Fg3a)
	td.Ftm					= atoi(ts.Ftm)
	td.Fta					= atoi(ts.Fta)
	td.Oreb					= atoi(ts.Oreb)
	td.Dreb					= atoi(ts.Dreb)
	td.Treb					= atoi(ts.Treb)
	td.Assists			= atoi(ts.Assists)
	td.Steals				= atoi(ts.Steals)
	td.Blocks				= atoi(ts.Blocks)
	td.Fouls				= atoi(ts.Fouls)
	td.Turnovers		= atoi(ts.Turnovers)
	td.PlusMinus    = atoi(ts.PlusMinus)

} // convTeam


func convTeamScore(t NbaTeamScore) Team {

	team := Team{}

	team.ID							= t.ID
	team.Name   				= strings.ToLower(t.ShortName)
	team.Score  				= t.Score
	team.Periods  			= convScores(t.Periods)

	return team

} // convTeamScore


func ConvBoxscore(b *NbaBoxscore) *Game {

	game := Game{}

	if b == nil {
		return nil
	}

	game.ID					= b.ID
	game.SeasonID   = b.SeasonID
	game.Date   		= b.Date
	game.StartUtc   = b.StartUtc
	game.EndUtc   	= b.EndUtc
	game.Away   		= convTeamScore(b.AwayScore)
	game.Home   		= convTeamScore(b.HomeScore)
	game.Plays      = convPlays(b.Plays)

	convTeam(&b.Away, &game.Away)
	convTeam(&b.Home, &game.Home)

	convPlayers(b.Players, &game)

	return &game

} // ConvBoxscore
