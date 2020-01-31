package stats

type Season struct {
	Name						string					`json:"name"`
	Games           []Stats         `json:"games"`
}

type Stats struct {
	Home						Team						`json:"homeTeam"`
	Away						Team						`json:"awayTeam"`
	DateOf          string          `json:"dateOf"`
}

type Team struct {
	Name						string					`json:"name"`
	Players         []Player				`json:"players"`
}

type Player struct {
	Name						string					`json:"name"`
	Minutes					int							`json:"minutes"`
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
	Starter         bool            `json:"starter"`
	DnpReason      	string          `json:"dnpReason"`
	Position        string          `json:"position"`
}
 