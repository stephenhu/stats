package stats

type RankStat struct {
	Val						float32					`json:"val"`
	Rank					int							`json:"rank"`
}

type Season struct {
	ID              string          		`json:"id"`
	Name						string							`json:"name"`
	Schedule        map[string][]Game   `json:"schedule"`
}

type Play struct {
	Clock							string					`json:"clock"`
	Description				string					`json:"description"`
	PersonID					string					`json:"personId"`
	TeamID						string					`json:"teamId"`
	TeamName          string          `json:"teamName"`
	Period            int             `json:"period"`
	Away         			int          		`json:"away"`
	Home         			int          		`json:"home"`
	EventType      		int          		`json:"eventType"`
	ScoreChanged      bool            `json:"scoreChanged"`
	Formatted					string					`json:"formatted"`
}

type GameLog struct {
	Date            string          `json:"date"`
	PubDate					string					`json:"pubDate"`
	GameID          string          `json:"gameId"`
	Plays           []Play          `json:"plays"`
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
	Ppg								float32				`json:"ppg"`			// points per game
	Rpg								float32				`json:"rpg"`			// rebounds per game
	Apg								float32				`json:"apg"`			// assists per game
	Mpg								float32				`json:"mpg"`			// minutes per game
	Tpg								float32				`json:"tpg"`			// turnovers per game
	Spg								float32				`json:"spg"`			// steals per game
	Bpg								float32				`json:"bpg"`			// blocks per game
	Fpg               float32				`json:"fpg"`			// fouls per game
	Fgp								float32				`json:"fgp"`			// fg %
	Fg3p							float32				`json:"fg3p"`			// three point %
	Ftp								float32				`json:"ftp"`			// free throw %
}

type Player struct {
	ID              string          `json:"id"`
	Opponent        string          `json:"opponent"`
	Name						string					`json:"name"`
	Minutes					int							`json:"minutes"`
	Seconds         int             `json:"seconds"`
	Starter         bool            `json:"starter"`
	DnpReason      	string          `json:"dnpReason"`
	Position        string          `json:"position"`
	Stats
}

type Player2 struct {
	OpponentID      string          `json:"opponentId"`
	Opponent				string					`json:"opponentName"`
	Player
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

type TeamSeasonStats struct {
	Name              string    `json:"name"`
	AdvStats
	OpponentPoints		int				`json:"opponentPoints"`
	Efficiency				string		`json:"efficiency"`
}

type TeamRanks struct {
	ID								string						`json:"id"`
	Fgp								RankStat					`json:"fgp"`
	Fg3p							RankStat					`json:"fg3p"`
	Ftp								RankStat					`json:"ftp"`
	Oreb							RankStat					`json:"orpg"`
	Dreb							RankStat					`json:"drpg"`
	Treb							RankStat					`json:"trpg"`
	Assists						RankStat					`json:"apg"`
	Turnovers					RankStat					`json:"tpg"`
	Steals						RankStat					`json:"spg"`
	Blocks						RankStat					`json:"bpg"`
	Fouls             RankStat          `json:"fpg"`
	Points						RankStat					`json:"ppg"`
	OpponentPoints		RankStat					`json:"oppg"`
	Efficiency				RankStat					`json:"efficiency"`
	Name      				string            `json:"name"`
}

type TeamData struct {
	Name          string    						`json:"name"`
	SeasonID      string                `json:"seasonId"`
	Ranks					TeamRanks							`json:"ranks"`
	Players				[]PlayerSeason				`json:"players"`
	Standings     `json:"standings"`
	Games         []Team               	`json:"games"`
}

type PlayerSeason struct {
	ID 						string				`json:"id"`				// player id
	PubDate				string				`json:"pubDate"`
	SeasonID      string        `json:"seasonId"`
	First         string        `json:"first"`
	Last         	string        `json:"last"`
	AdvStats
}

type PlayerCareer struct {
	ID 						string				`json:"id"`				// player id
	PubDate				string				`json:"pubDate"`
	SeasonID      string        `json:"seasonId"`
	First         string        `json:"first"`
	Last         	string        `json:"last"`
	TeamName      string        `json:"teamName"`
	Latest    		AdvStats      `json:"latest"`
	Career       	AdvStats      `json:"career"`
	Seasons       []SeasonStats `json:"seasons"`
	LastGames     []Player      `json:"lastGames"`
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
	SeasonID        string          `json:"seasonId"`   // discard
	Name						string					`json:"name"`
	Opponent        string          `json:"opponent"`
	Win             bool            `json:"win"`
	Score           int             `json:"score"`
	OpponentScore   int             `json:"opponentScore"`
	Periods     		[]int           `json:"periods"`
	Players         []Player				`json:"players"`
	Stats           `json:"summary"`
}

type Game struct {
	ID              string          `json:"id"`
	SeasonID        string          `json:"seasonId"`
	Date            string          `json:"date"`
	PubDate					string					`json:"pubDate"`
	StartUtc        string          `json:"startUtc"`
	EndUtc          string          `json:"endUtc"`
	Source          string          `json:"source"`
	Home						Team						`json:"home"`
	Away						Team						`json:"away"`
	Plays           []Play					`json:"plays"`
}

type AllRanks struct {
	SeasonID				string          `json:"seasonId"`
	PubDate					string					`json:"pubDate"`
	Teams 					[]TeamRanks			`json:"teams"`
}

type AllTeams struct {
	SeasonID				string					`json:"seasonId"`
	PubDate					string					`json:"pubDate"`
	Teams						[]TeamInfo			`json:"teams"`
}

type AllPlayers struct {
	SeasonID        string         	`json:"seasonId"`
	PubDate					string					`json:"pubDate"`
	Players					[]PlayerInfo		`json:"players"`
}

type Roster struct {
	TeamID          string          `json:"teamId"`
	Players					[]string				`json:"players"`
}

type TeamRecord struct {
	TeamID					string				`json:"teamId"`
	Name           	string        `json:"name"`
	Mascot          string        `json:"mascot"`
	Rank            int           `json:"rank"`
	W								int						`json:"w"`
	L								int						`json:"l"`
	Pct							float32				`json:"pct"`
	Gb							float32				`json:"gb"`
	Cw							int						`json:"cw"`
	Cl							int						`json:"cl"`
	Dw							int						`json:"dw"`
	Dl							int						`json:"dl"`
	Hw							int						`json:"hw"`
	Hl							int						`json:"hl"`
	Aw							int						`json:"aw"`
	Al							int						`json:"al"`
	L10w						int						`json:"l10w"`
	L10l						int						`json:"l10l"`
	Streak					int						`json:"streak"`
	WinStreak				bool					`json:"isWinStreak"`
	StreakText      string        `json:"streakText"`
}

type Standings struct {
	SeasonID				string										`json:"seasonId"`
	PubDate         string        						`json:"pubDate"`
	Records					map[string]TeamRecord			`json:"records"`
}
