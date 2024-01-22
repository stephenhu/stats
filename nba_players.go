package stats


type NbaPlayerStatistics struct {
	Assists       int           		`json:"assists"`
	Blocks        int           		`json:"blocks"`
	Blocked       int           		`json:"blocksReceived"`
	Fga        		int           		`json:"fieldGoalsAttempted"`
	Fgm        		int           		`json:"fieldGoalsMade"`
	Fgp        		float32           `json:"fieldGoalsPercentage"`
	FoulsOff      int           		`json:"foulsOffensive"`
	FoulsDrawn    int        				`json:"foulsDrawn"`
	Fouls        	int        				`json:"foulsPersonal"`
	Technicals    int  	         		`json:"foulsTechnical"`
	Fta        		int           		`json:"freeThrowsAttempted"`
	Ftm        		int           		`json:"freeThrowsMade"`
	Ftp           float32           `json:"freeThrowsPercentage"`
	Minutes       string        		`json:"minutesCalculated"`
	Plus          float32           `json:"plus"`
	PlusMinus     float32           `json:"plusMinusPoints"`
	Points        int   	        	`json:"points"`
  PointsFast    int               `json:"pointsFastBreak"`
	PointsPaint   int               `json:"pointsInThePaint"`
	PointsSecond  int               `json:"pointsSecondChance"`
	Dreb        	int           		`json:"reboundsDefensive"`
	Oreb        	int           		`json:"reboundsOffensive"`
	Treb        	int           		`json:"reboundsTotal"`
	Steals        int           		`json:"steals"`
	Fg3a        	int          			`json:"threePointersAttempted"`
	Fg3m        	int          			`json:"threePointersMade"`
  Fg3p          float32           `json:"threePointersPercentage"`
	Turnovers     int           		`json:"turnovers"`
	Fg2a          int           		`json:"twoPointersAttempted"`
	Fg2m          int           		`json:"twoPointersMade"`
	Fg2p          float32       		`json:"twoPointersPercentage"`
}


type NbaPlayer struct {
	ID						int						`json:"personId"`
	First					string 				`json:"firstName"`
	Last					string        `json:"familyName"`
	Name          string        `json:"name"`
	NameShort     string        `json:"nameI"`
	Jersey				string				`json:"jerseyNum"`
	Position			string 				`json:"position"`
	Starter				string        `json:"starter"`
	Order         int						`json:"order"`
	Statistics    NbaPlayerStatistics		`json:"statistics"`
}


type NbaPlayerInfo struct {
	ID						string				`json:"personId"`
	First					string 				`json:"firstName"`
	Last					string        `json:"lastName"`
	TeamID				int						`json:"teamId"`
	Jersey				string				`json:"jerseyNum"`
	Position			string				`json:"position"`
	Starter       string        `json:"starter"`
	Feet					string				`json:"heightFeet"`
	Inches				string				`json:"heightInches"`
	Meters				string				`json:"heightMeters"`
	Pounds				string				`json:"weightPounds"`
	Kilograms			string				`json:"weightKilograms"`
	Dob						string				`json:"dateOfBirthUTC"`
	Rookie				string				`json:"nbaDebutYear"`
	Years      		string        `json:"yearsPro"`
	College       string        `json:"collegeName"`
	Active      	bool        	`json:"isActive"`
}

type NbaSeasonStats struct {
	SeasonID			int							`json:"seasonYear"`
	Teams					[]NbaAdvStats		`json:"teams"`
	Total					NbaAdvStats			`json:"total"`
}

type NbaRegularSeason struct {
	SeasonInfo 		[]NbaSeasonStats 		`json:"season"`
}

type NbaPlayerStats struct {
	Latest 				NbaAdvStats		`json:"latest"`
	Career 				NbaAdvStats		`json:"careerSummary"`
	NbaRegularSeason		`json:"regularSeason"`

}

type NbaStats2 struct {
	NbaPlayerStats		`json:"stats"`
}

type NbaLeague2 struct {
	NbaStats2				`json:"standard"`
}

type NbaLeague struct {
	Players 	[]NbaPlayerInfo			`json:"standard"`
}

type NbaLeaguePlayers struct {
	NbaLeague				`json:"league"`
	SeasonID        string				`json:"season"`
}

type NbaPlayerProfile struct {
	NbaLeague2				`json:"league"`
	ID						string        `json:"id"`
	SeasonID			string        `json:"seasonId"`
	First         string        `json:"first"`
	Last         	string        `json:"last"`
}
