package stats

import (
	"context"
	"fmt"
	"log"
	"net/http"	
	//"regexp"
	//"strconv"
	"strings"
	//"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/cdp"
)


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
		log.Println(err)
	} else {
		
		for _, n := range nodes {

			u := n.AttributeValue(HTML_ATTR_HREF)

			if strings.Contains(u, BOXSCORE_MATCH) {
				
				x := strings.Split(u, EQUAL_MATCH)
				games[strings.TrimSpace(x[len(x)-1])] = 1

			}

		}

	}

	return games

} // getGames


func GetStats(games map[string] int) map[string] Stats {

	all := map[string] Stats{}

	for g, _ := range games {

		_, err := StringUrlJoin(ESPN_BASE_URL, g)

		if err != nil {
			logf("getStats", err.Error())
		} else {

			res, err := http.Get(ESPN_BOXSCORE_URL + g)

			if err != nil {
				logf("getStats", err.Error())
			} else {
		
				defer res.Body.Close()
		
				doc, err := goquery.NewDocumentFromReader(res.Body)
		
				if err != nil {
					logf("getStats", err.Error())
				} else {

					doc.Find(HTML_DIV).Each(func(index int, div *goquery.Selection) {
		
						s, _ := div.Attr(HTML_ATTR_ID)						

						if s == ESPN_BOXSCORE_ID {

							players := []Player{}
						
							div.Find(HTML_TR).Each(func(itr int, tr *goquery.Selection) {
																
								p := Player{}

								tr.Find(HTML_TD).Each(func(itd int, td *goquery.Selection) {								
									
									process
									value := td.Text()
									
									if value != "" {

										attr, ok := td.Attr("class")

										if ok {
											
											if attr != STATUS_DNP && attr != HIGHLIGHT {

												switch itd {
												case 0:													
													p.Name = value
													// TODO: clean up the name
													
												case 1:	
												
													if value == STARTERS {
														p.Starter = true
													} else if value == BENCH {
														p.Starter = false
													} else {
														p.Minutes = atoi(value)
													}
													
												case 2:
													fieldGoals(&p, FIELD_FG, value)
												case 3:
													fieldGoals(&p, FIELD_FG3, value)
												case 4:
													fieldGoals(&p, FIELD_FT, value)
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
													logf("getStats", fmt.Sprintf("Unrecognized table field #%d %s",
														itd, value))
												}
	
											} else {
												p.DnpReason = value
											}

										}																		

									}

								})

								players = append(players, p)

							})

							log.Println(players)

						}

						
		
					})
		
				}
		
			}

		}
	
	}

	return all

} // getStats


func getStatsww(games map[string] int) map[string] Stats {

	all := map[string] Stats{}

	for g, _ := range games {

		_, err := StringUrlJoin(ESPN_BASE_URL, g)

		if err != nil {
			logf("getStats", err.Error())
		} else {

			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", false),
				chromedp.Flag("disable-gpu", false),
		
			)
		
			ea, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
			defer cancel()
		
			ctx, cancel := chromedp.NewContext(ea)
			defer cancel()
	
			var nodes []*cdp.Node
	
			err := chromedp.Run(ctx,
				chromedp.Navigate(ESPN_BASE_URL + g),
				chromedp.WaitVisible(ESPN_BOXSCORE_ID, chromedp.ByID),
				chromedp.Nodes(`//table`, &nodes, chromedp.BySearch))
		
			if err != nil {
				log.Println(err)
			} else {
				
				for _, n := range nodes {

					td := n.AttributeValue(HTML_TD)
		
					log.Println(td)
		
				}

			}
	
		}

	}

	return all

} // getStats
