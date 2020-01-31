package stats

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
			p.Fga = attempts
			p.Fgm = made

		case FIELD_FG3:
			p.Fg3a = attempts
			p.Fg3m = made

		case FIELD_FT:
			p.Fta = attempts
			p.Ftm = made
		
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
			p.Oreb = atoi(value)
		case 6:
			p.Dreb = atoi(value)
		case 7:
			p.Treb = atoi(value)
		case 8:
			p.Assists = atoi(value)
		case 9:
			p.Steals = atoi(value)
		case 10:
			p.Blocks = atoi(value)
		case 11:
			p.Turnovers = atoi(value)
		case 12:
			p.Fouls = atoi(value)
		case 13:
			p.PlusMinus = atoi(value)
		case 14:
			p.Points = atoi(value)
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

		}

	})

	return players

} // parsePlayers


func parseBoxScore(d *goquery.Document) *Stats {

	stats := Stats{}

	d.Find(HTML_DIV).Each(func(index int, div *goquery.Selection) {
		
		id, _ := div.Attr("id")

		if id == ESPN_BOXSCORE_ID {

			div.Find(HTML_TBODY).Each(func(itb int, tbody *goquery.Selection) {
				
				players := parsePlayers(tbody, true)				

				if itb == INDEX_AWAY_STARTERS || itb == INDEX_AWAY_BENCH {
					stats.Away.Players = append(stats.Away.Players, players...)
				} else if itb == INDEX_HOME_STARTERS || itb == INDEX_HOME_BENCH {
					stats.Home.Players = append(stats.Home.Players, players...)
				}

			})

			homeTeam := false

			div.Find(HTML_DIV).Each(func(index2 int, div2 *goquery.Selection) {

				if div2.HasClass(ESPN_TEAM_NAME_CLASS) {

					name := div2.Text()

					if !homeTeam {
						stats.Away.Name = name
						homeTeam = true
					} else {
						stats.Home.Name = name
					}
					
				}

			})

		}

	})

	return &stats

} // parseBoxScore


func GetGames(d string) map[string] int {

	games := map[string] int{}

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
		chromedp.Navigate(ESPN_SCOREBOARD_URL),
		chromedp.WaitVisible(ESPN_SCOREBOARD_EVENTS_ID, chromedp.ByID),
		chromedp.Nodes(`//a`, &nodes, chromedp.BySearch))

	if err != nil {
		logf("GetGames", err.Error())
	} else {
		
		for _, n := range nodes {

			u := n.AttributeValue(HTML_ATTR_HREF)

			if strings.Contains(u, MATCH_BOXSCORE) {
				
				x := strings.Split(u, STRING_EQUAL)
				games[strings.TrimSpace(x[len(x)-1])] = 1

			}

		}

	}

	return games

} // getGames


func GetStats(games map[string] int) []Stats {

	all := []Stats{}

	for g, _ := range games {

		res, err := http.Get(ESPN_BOXSCORE_URL + g)

		if err != nil {
			logf("GetStats", err.Error())
		} else {
	
			defer res.Body.Close()
	
			doc, err := goquery.NewDocumentFromReader(res.Body)
	
			if err != nil {
				logf("GetStats", err.Error())
			} else {

				stats := parseBoxScore(doc)

				all = append(all, *stats)

			}
	
		}
	
	}

	log.Printf("%+v", all)	

	return all

} // getStats
