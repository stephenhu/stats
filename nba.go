package stats

type Season struct {
	ID              string          `json:"id"`
	Name						string					`json:"name"`
	Games           []Game          `json:"games"`
}

type Game struct {
	ID              string          `json:"id"`
	SID        			string          `json:"sid"`
	Home						Team						`json:"home"`
	Away						Team						`json:"away"`
	Date            string          `json:"date"`
}

type Stats struct {
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

type AdvStats struct {
	SeasonID					string				`json:"seasonId"`
	TeamID						string        `json:"teamId"`
	Minutes						int						`json:"minutes"`
	Points						int						`json:"points"`
	Oreb							int						`json:"oreb"`
	Dreb							int						`json:"dreb"`
	Treb							int						`json:"treb"`
	Assists						int						`json:"assists"`	
	Turnovers					int						`json:"turnovers"`
	Steals						int						`json:"steals"`
	Blocks						int						`json:"blocks"`
	Fouls							int						`json:"fouls"`
	Fgm               int        		`json:"fgm"`
	Fga               int        		`json:"fga"`
	Fg3m              int        		`json:"fg3m"`
	Fg3a              int        		`json:"fg3a"`
	Ftm               int        		`json:"ftm"`
	Fta               int        		`json:"fta"`
	Played						int						`json:"played"`
	Started						int						`json:"started"`
	PlusMinus					int						`json:"plusMinus"`
	Ppg								string				`json:"ppg"`			// points per game
	Rpg								string				`json:"rpg"`			// rebounds per game
	Apg								string				`json:"apg"`			// assists per game
	Mpg								string				`json:"mpg"`			// minutes per game
	Topg							string				`json:"tpg"`			// turnovers per game
	Spg								string				`json:"spg"`			// steals per game
	Bpg								string				`json:"bpg"`			// blocks per game
	Fgp								string				`json:"fgp"`			// fg %
	Fg3p							string				`json:"fg3p"`			// three point %
	Ftp								string				`json:"ftp"`			// free throw %
}

type Player struct {
	ID              string          `json:"id"`
	Name						string					`json:"name"`
  Stats         	`json:"playerStats"`
	Minutes					int							`json:"minutes"`
	Seconds         int             `json:"seconds"`
	Starter         bool            `json:"starter"`
	DnpReason      	string          `json:"dnpReason"`
	Position        string          `json:"position"`
}

type PlayerInfo struct {
	ID						string				`json:"id"`
	First					string 				`json:"first"`
	Last					string        `json:"last"`
	TeamID				string				`json:"teamId"`
	Jersey				string				`json:"jersey"`
	Position			string				`json:"position"`
	Feet					int						`json:"feet"`
	Inches				int						`json:"inches"`
	Meters				string				`json:"meters"`
	Pounds				int						`json:"pounds"`
	Kilograms			string				`json:"kilograms"`
	Dob						string				`json:"dob"`
	Rookie				string				`json:"rookie"`
	Years      		int        		`json:"years"`
	College       string        `json:"college"`
	Active      	bool        	`json:"active"`
}

type SeasonStats struct {
	SeasonID			string        `json:"seasonId"`
	Teams					[]AdvStats		`json:"teams"`
	Summary       AdvStats      `json:"summary"`
}

type PlayerCareer struct {
	ID 						string				`json:"id"`				// player id
	SeasonID      string        `json:"seasonId"`
	First         string        `json:"first"`
	Last         	string        `json:"last"`
	Latest    		AdvStats      `json:"latest"`
	Career       	AdvStats      `json:"career"`
	Seasons       []SeasonStats `json:"seasons"`
}

type TeamInfo struct {
	ID              string          `json:"id"`
	Full            string          `json:"full"`
	Code       			string          `json:"code"`
	Short           string          `json:"short"`
	City            string          `json:"city"`
	Conference      string          `json:"conference"`
	Division      	string          `json:"division"`
}

type Team struct {
	ID              string          `json:"id"`
	SeasonID        string          `json:"seasonId"`
	Name						string					`json:"name"`
	Score           int             `json:"score"`
	ScoreDetail     []int           `json:"scoreDetail"`
	Players         []Player				`json:"players"`
	Stats           `json:"stats"`	
}

type AllTeams struct {
	SeasonID				string					`json:"seasonId"`
	Date						string					`json:"date"`
	Teams						[]TeamInfo			`json:"teams"`
}

type AllPlayers struct {
	SeasonID        string         	`json:"seasonId"`
	Date            string          `json:"date"`	
	Players					[]PlayerInfo		`json:"players"`
}
