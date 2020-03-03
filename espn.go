package stats

import (
	"context"
	//"encoding/json"
	"fmt"
	//"log"
	//"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/cdp"
)


func fieldGoals(p *Player, fn string, s string) {

	if s == STRING_EMPTY {
		logf("fieldGoals", fmt.Sprintf("Empty string for %s", fn))
		return
	}

	tokens := strings.Split(s, STRING_MINUS)

	if len(tokens) != 2 {
		logf("fieldsGoals", fmt.Sprintf(
			"Field format incorrect for %s, should be of the format x-y", s))
	} else {

		made     := atoi(tokens[INDEX_MADE])
		attempts := atoi(tokens[INDEX_ATTEMPTS])

		switch fn {
		case FIELD_FG:
			p.Stats.Fga = attempts
			p.Stats.Fgm = made

		case FIELD_FG3:
			p.Stats.Fg3a = attempts
			p.Stats.Fg3m = made

		case FIELD_FT:
			p.Stats.Fta = attempts
			p.Stats.Ftm = made

		default:
			logf("fieldGoals", "Unrecognized field name.")
		}

	}

} // fieldGoals


func parsePosition(n *goquery.Selection) string {

	pos := ""

	n.Find(HTML_SPAN).Each(func(index int, span *goquery.Selection) {

		if index == INDEX_SPAN_POSITION {
			pos = span.Text()
		}

	})

	return pos

} // parsePosition


func parseName(n *goquery.Selection) string {

	name := n.Text()

	n.Find(HTML_SPAN).Each(func(index int, span *goquery.Selection) {

		if index == INDEX_SPAN_NAME {
			name = span.Text()
		}

	})

	return name

} // parseName


func parsePlayer(index int, value string, p *Player) {

	switch index {
		case 1:
			p.Minutes = atoi(value)
		case 2:
			fieldGoals(p, FIELD_FG, value)
		case 3:
			fieldGoals(p, FIELD_FG3, value)
		case 4:
			fieldGoals(p, FIELD_FT, value)
		case 5:
			p.Stats.Oreb = atoi(value)
		case 6:
			p.Stats.Dreb = atoi(value)
		case 7:
			p.Stats.Treb = atoi(value)
		case 8:
			p.Stats.Assists = atoi(value)
		case 9:
			p.Stats.Steals = atoi(value)
		case 10:
			p.Stats.Blocks = atoi(value)
		case 11:
			p.Stats.Turnovers = atoi(value)
		case 12:
			p.Stats.Fouls = atoi(value)
		case 13:
			p.Stats.PlusMinus = atoi(value)
		case 14:
			p.Stats.Points = atoi(value)
		default:
			logf("parsePlayer", fmt.Sprintf("Unrecognized table field #%d %s",
				index, value))
		}

} // parsePlayer


func parsePlayers(tbody *goquery.Selection, starting bool) []Player {

	players := []Player{}

	tbody.Find(HTML_TR).Each(func(itr int, tr *goquery.Selection) {

		if !tr.HasClass(HIGHLIGHT) {

			p := Player{}

			if starting {
				p.Starter = true
			}

			tr.Find(HTML_TD).Each(func(itd int, td *goquery.Selection) {

				value := td.Text()

				if value != STRING_EMPTY {

					if td.HasClass(DNP) {
						p.DnpReason = value
					} else {

						if itd == INDEX_FIELD_NAME {

							p.Name 			= parseName(td)
							p.Position 	= parsePosition(td)

						} else {
							parsePlayer(itd, value, &p)
						}

					}

				}

			})

			players = append(players, p)

		} else {


		}

	})

	return players

} // parsePlayers


func parseBoxScore(d *goquery.Document) *Game {

	game := Game{}

	d.Find(HTML_DIV).Each(func(index int, div *goquery.Selection) {

		id, _ := div.Attr("id")

		if id == ESPN_BOXSCORE_ID {

			div.Find(HTML_TBODY).Each(func(itb int, tbody *goquery.Selection) {

				players := parsePlayers(tbody, true)

				if itb == INDEX_AWAY_STARTERS || itb == INDEX_AWAY_BENCH {
					game.Away.Players = append(game.Away.Players, players...)
				} else if itb == INDEX_HOME_STARTERS || itb == INDEX_HOME_BENCH {
					game.Home.Players = append(game.Home.Players, players...)
				}

			})

			homeTeam := false

			div.Find(HTML_DIV).Each(func(index2 int, div2 *goquery.Selection) {

				if div2.HasClass(ESPN_TEAM_NAME_CLASS) {

					name := div2.Text()

					if !homeTeam {
						game.Away.Name = name
						homeTeam = true
					} else {
						game.Home.Name = name
					}

				}

			})

		}

	})

	return &game

} // parseBoxScore


func GetGameIDs(d string) map[string] string {

	var location string

	if dateCheck(d) {
		location = fmt.Sprintf(ESPN_SCOREBOARD_DATE_URL, d)
	} else {
		location = ESPN_BOXSCORE_URL
	}

	ids := map[string] string{}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", false),
	)

	ea, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(ea)
	defer cancel()

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(location),
		chromedp.WaitVisible(ESPN_SCOREBOARD_EVENTS_ID, chromedp.ByID),
		chromedp.Nodes(`//a`, &nodes, chromedp.BySearch))

	if err != nil {
		logf("GetGameIDs", err.Error())
	} else {

		for _, n := range nodes {

			u := n.AttributeValue(HTML_ATTR_HREF)

			if strings.Contains(u, MATCH_BOXSCORE) {

				x := strings.Split(u, STRING_EQUAL)
				ids[strings.TrimSpace(x[len(x)-1])] = d

			}

		}

	}

	return ids

} // GetGameIDs


func GetGameIDsFrom(d string) map[string]string {

	all := map[string]string{}

	days := GetDays(d)

	for _, day := range days {

		ids := GetGameIDs(day)

		for k, _ := range ids {
			all[k] = day
		}

	}

	return all

} // GetGameIDsFrom


func GetGameIDsBySeason(d string) map[string]string {

	all := map[string]string{}

	season, ok := OfficialSeasons[d]

	if ok {

		days := GetDays(season[SEASON_BEGIN])

		for _, d := range days {

			ids := GetGameIDs(d)

			for k, _ := range ids {
				all[k] = d
			}

		}

	}

	return all

} // GetGameIDsBySeason


func GetGames(gameIDs map[string] string) []Game {

	if gameIDs == nil || len(gameIDs) == 0 {
		return nil
	}

	all := []Game{}

	for id, gameDate := range gameIDs {

		res, err := client.Get(fmt.Sprintf("%s%s", ESPN_BOXSCORE_URL, id))

		if err != nil {
			logf("GetGames", err.Error())
		} else {

			defer res.Body.Close()

			doc, err := goquery.NewDocumentFromReader(res.Body)

			if err != nil {
				logf("GetGames", err.Error())
			} else {

				game := parseBoxScore(doc)

				game.ID 		= id
				game.Date   = gameDate

				all = append(all, *game)

			}

		}

	}

	//log.Printf("%+v", all)

	return all

} // GetGames


func StoreSeason(s string) {

	ids := GetGameIDsBySeason(s)

	games := GetGames(ids)

	StoreGames(games)

} // StoreSeason


func StoreGameDay(d string) {

	ids := GetGameIDs(d)

	games := GetGames(ids)

  StoreGames(games)

} // StoreGameDay


func StoreFromDay(d string) {

	ids := GetGameIDsFrom(d)

	games := GetGames(ids)

	StoreGames(games)

} // StoreFromDay


func StoreGames(games []Game) {

	for _, g := range games {
		putGame(&g)
	}

} // StoreGames
