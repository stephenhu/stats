package stats

import(
	"net/http"
	"time"
)

const (
	APP_NAME										= "stats"
	APP_STORAGE                 = "nba"
	APP_VERSION             		= "1.1"
)

const (
	ESPN_BASE_URL               = `https://www.espn.com`
	ESPN_SCOREBOARD_URL					= `https://espn.com/nba/scoreboard`
	ESPN_SCOREBOARD_DATE_URL		= `https://espn.com/nba/scoreboard/_/date/%s`
	ESPN_BOXSCORE_URL           = `https://www.espn.com/nba/boxscore?gameId=`
	ESPN_PLAYER_ICON_URL        = `https://a.espncdn.com/combiner/i?img=/i/headshots/nba/players/full/%s.png&h=80&w=110&scale=crop`
	ESPN_TEAM_ICON_URL          = `https://a.espncdn.com/i/teamlogos/nba/500/%s.png`
)

const (
	ESPN_SCOREBOARD_EVENTS_ID     = "events"
	ESPN_BOXSCORE_ID              = "gamepackage-box-score"
	ESPN_TEAM_NAME_CLASS          = "team-name"
)

const (
	//NBA_BASE_URL									= "https://data.nba.net/data/10s/prod/v1"
	NBA_BASE_URL                  = "https://cdn.nba.com/static/json"
	NBA_STATIC                    = "/staticData"
	NBA_SCHEDULE                  = "/scheduleLeagueV2_9.json"
	NBA_LIVE                      = "/liveData"
	NBA_API_BOXSCORE							= "/boxscore/boxscore_%s.json"
	NBA_API_PLAYS                 = "/playbyplay/playbyplay_%s.json"
	//NBA_API_BOXSCORE							= "/%s/%s_boxscore.json"
	NBA_API_PLAYERS								= "/%s/players.json"
	NBA_API_PLAYER_PROFILE        = "/%s/players/%s_profile.json"
	//NBA_API_PLAYS            			= "/%s/%s_pbp_%d.json"
	NBA_API_ROSTER            		= "/%s/teams/%s/roster.json"
	//NBA_API_SCOREBOARD            = "/%s/scoreboard.json"
	NBA_API_TODAYS_SCOREBOARD     = "/scoreboard/todaysScoreboard_00.json"
	NBA_API_STANDINGS             = "/current/standings_all.json"
	NBA_API_TEAMS                 = "/%s/teams.json"
	NBA_API_TEAM_RANKS            = "/%s/team_stats_rankings.json"
)

const (
	HTML_ANCHOR                 = "a"
	HTML_ARTICLE								= "article"
	HTML_ATTR_CLASS							= "class"
	HTML_ATTR_HREF              = "href"
	HTML_ATTR_ID                = "id"
	HTML_DIV										= "div"
	HTML_SCRIPT                 = "script"
	HTML_SPAN                   = "span"
	HTML_TABLE                  = "table"
	HTML_TBODY                  = "tbody"
	HTML_TD                     = "td"
	HTML_TR                     = "tr"
)

const (
	MATCH_BOXSCORE        			= `/nba/game?gameId=`
	MATCH_BOXSCORE_RE						= "gameId=[0-9]+}"
)

const (
	STRING_COLON                = ":"
	STRING_EMPTY                = ""
	STRING_EQUAL                = "="
	STRING_MINUS                = "-"
	STRING_NA                   = "--"
	STRING_PERIOD               = "."
	STRING_SINGLE_QUOTE         = "'"
	STRING_SPACE                = " "
	STRING_TWO_SPACES           = "  "
	STRING_TAB                  = "    "
	STRING_ZERO                 = "0"
	STRING_ZERO_FLOAT           = "0.0"
)

const (
	FIELD_FG										= "fg"
	FIELD_FG3										= "fg3"
	FIELD_FT                    = "ft"
)

const (
	INDEX_MADE                  = 0
	INDEX_ATTEMPTS							= 1
	INDEX_SPAN_NAME             = 0
	INDEX_SPAN_NAME_ABBR        = 1
	INDEX_SPAN_POSITION         = 2
	INDEX_AWAY_STARTERS					= 0
	INDEX_AWAY_BENCH            = 1
	INDEX_HOME_STARTERS         = 2
	INDEX_HOME_BENCH            = 3
)

const (
	INDEX_FIELD_NAME            = 0
)

const (
	BASE10											= 10
	BITS32											= 32
	BY2                 				= 2
)

const (
	FLOAT_TO_PERCENT            = 100.0
)

const (
	DNP													= "dnp"
	HIGHLIGHT                   = "highlight"
)

const (
	STARTERS										= "starters"
	BENCH												= "bench"
)

const (
	EST_FORMAT									= "20060102 3:04 PM MST"
	DATE_FORMAT                 = "20060102"
	YEAR_FORMAT                 = "2006"
	NBA_DATETIME_FORMAT         = "01/02/2006 15:04:05"
	EST													= "America/New_York"			// Eastern Standard Time
)

const (
	START_TIME_WEEKDAY          = 17
	START_TIME_WEEKEND          = 15
	SCOREBOARD_MAX_RETRY        = 5
)

const (
	MAX_SEASON_END							= "0630"			// add buffer to be safe
	MAX_PERIODS                 = 10					// 6 overtimes
)

const (
	JSON_INDENT									= "  "
	JSON_PREFIX									= ""
)

const (
	GAME_FILE                 	= "%s%s.json"
	PLAYERS_DIR                 = "players"
	PLAYERS_FILE								= "players.json"
	PLAYS_FILE                  = "%s.%s.plays.json"
	//TEAMS_DIR                   = "teams"
	TEAMS_FILE                  = "teams.json"
	TEAM_RANKS_FILE             = "team.ranks.json"
	NBA_SCHEDULE_FILE           = "schedule.%s.json"
)

const (
	PROFILE_BIRD_ERA         		= "bird.era"					// 1978
	PROFILE_MODERN_ERA					= "modern.era"				// 1979
	PROFILE_RELATIVE_ERA     		= "relative.era" 			// 1998
	PROFILE_SIMPLE_ERA					= "simple.era"				// 2015
)

const (
	YEAR_BIRD_ERA               = 1978
	YEAR_MODERN_ERA             = 1979
)

const (
	HTTP_TIMEOUT                = 45			// seconds
)

var client = http.Client{
	Timeout: HTTP_TIMEOUT * time.Second,
}
