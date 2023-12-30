package stats


type CachePlayerData struct {
	ID                    string        `json:"id"`
	First									string        `json:"first"`
	Last									string        `json:"last"`
	Full									string        `json:"full"`
	Abv										string        `json:"abv"`
	Starter								string        `json:"starter"`
	Active                string        `json:"active"`
	Minutes               int           `json:"minutes"`
	Points								int           `json:"points"`
	Oreb									int           `json:"oreb"`
	Dreb									int           `json:"dreb"`
	Treb									int           `json:"treb"`
	Fga										int           `json:"fga"`
	Fgm										int           `json:"fgm"`
	Fta										int           `json:"fta"`
	Ftm										int           `json:"ftm"`
	Fg3a									int           `json:"fg3a"`
	Fg3m									int           `json:"fg3m"`
	Steals								int           `json:"steals"`
	Assists								int           `json:"assists"`
	Blocks								int           `json:"blocks"`
	Blocked								int           `json:"blocked"`
	Turnovers							int           `json:"turnovers"`
	Fouls									int           `json:"fouls"`
	Fouled								int           `json:"fouled"`
	Technicals						int           `json:"technicals"`
	Fragrants							int           `json:"fragrants"`
	Paint									int           `json:"paint"`
	Fastbreak							int           `json:"fastbreak"`
	SecondChance					int           `json:"secondChance"`
	PlusMinus							int           `json:"plusMinus"`
}


type CachePeriod struct {
	Points								int           `json:"points"`
	Away                  int           `json:"away"`
	Home                  int           `json:"home"`
}


type CacheTeamData struct {
	TeamID                int           `json:"teamId"`
	Score                 int           `json:"score"`
	Players               map[int]CachePlayerData 	`json:"players"`
}


type CachePlay struct {
	ID										string				`json:"id"`
	Detail								string				`json:"detail"`
}


type CacheGame struct {
	Away									CacheTeamData			`json:"away"`
	Home									CacheTeamData			`json:"home"`
	Periods								map[int]CachePeriod			`json:"periods"`
	Plays                 map[string]CachePlay     `json:"plays"`
}


type CacheRankings struct {

}


type CacheLeaders struct {

}


type CacheSeason struct {
	ID										string				`json:"id"`
	Games									map[int]CacheGame	`json:"games"`
	Rankings							CacheRankings			`json:"rankings"`
	Leaders								CacheLeaders				`json:"leaders"`
}


type CachePlayer struct {
	ID                    int           `json:"id"` 
	TeamID                int           `json:"teamId"`
	First									string				`json:"first"`
	Last									string				`json:"last"`
	Full									string				`json:"full"`
	Abv										string				`json:"abv"`
	Position							string				`json:"position"`
	Active								string				`json:"active"`
	Height								int						`json:"height"`
	Weight								int						`json:"weight"`
}


type CacheTeam struct {
	ID                		int           `json:"id"`
	Name									string				`json:"name"`
	Code									string				`json:"last"`
	City									string				`json:"full"`
	Mascot								string				`json:"abv"`
	Conf									string				`json:"position"`
	Div										string				`json:"active"`
}


type NbaCache struct {
  Seasons   						map[string]CacheSeason	`json:"seasons"`
	Players               map[int]CachePlayer			`json:"players"`
	Teams               	map[int]CacheTeam				`json:"teams"`
}
