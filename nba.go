package stats

type Season struct {
	Name						string					`json:"name"`
	Games           []Game          `json:"games"`
}

type Game struct {
	ID              string          `json:"id"`
	Home						Team						`json:"home"`
	Away						Team						`json:"away"`
	Date            string          `json:"date"`
}

type Data struct {
	Fga							int							`json:"fga"`
	Fgm							int							`json:"fgm"`
	Fg3a						int							`json:"fg3a"`
	Fg3m						int							`json:"fg3m"`
	Fta							int							`json:"fta"`
	Ftm							int							`json:"ftm"`
	Oreb						int							`json:"oreb"`
	Dreb						int							`json:"dreb"`
	Treb						int							`json:"treb"`
	Assists					int							`json:"assists"`
	Steals					int							`json:"steals"`
	Blocks					int							`json:"blocks"`
	Turnovers				int							`json:"turnovers"`
	Fouls						int							`json:"fouls"`
	PlusMinus				int							`json:"plusMinus"`
	Points					int							`json:"points"`
}

type Player struct {
	Name						string					`json:"name"`
  Stats         	Data         		`json:"data"`
	Minutes					int							`json:"minutes"`
	Starter         bool            `json:"starter"`
	DnpReason      	string          `json:"dnpReason"`
	Position        string          `json:"position"`
}

type Team struct {
	Name						string					`json:"name"`
	Score           []int           `json:"score"`
	Players         []Player				`json:"players"`
	Stats           Data            `json:"stats"`
}

type Info struct {
	Begin						int							`json:"begin"`
	End							int							`json:"end"`
}

type SeasonInfo struct {
	Regular					Info						`json:"regular"`
	Playoffs        Info            `json:"playoffs"`
}
