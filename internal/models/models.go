package models

import "time"

// Player represents an NBA player
type Player struct {
	PlayerID   string `json:"player_id"`
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
	TeamName   string `json:"team_name"`
	TeamID     string `json:"team_id"`
	Number     string `json:"number"`
	Status     string `json:"status"`
}

// Team represents an NBA team
type Team struct {
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	City     string `json:"city"`
	Abbr     string `json:"abbr"`
}

// GameStats represents player game statistics
type NBAGameStats struct {
	Points            float64 `json:"points"`
	Assists           float64 `json:"assists"`
	Rebounds          float64 `json:"rebounds"`
	ThreePointersMade float64 `json:"threePointersMade"`
	Minutes           float64 `json:"minutes"`
}

// PlayerGameLog represents a player's game log entry
type NBAPlayerGameLog struct {
	GameDate string       `json:"game_date"`
	Stats    NBAGameStats `json:"stats"`
}

// TeamGameLog represents a team's game log entry
type TeamGameLog struct {
	GameDate string  `json:"game_date"`
	Points   float64 `json:"points"`
}

// Game represents a live game
type Game struct {
	GameID   string `json:"game_id"`
	HomeCity string `json:"home_city"`
	HomeTeam string `json:"home_team"`
	AwayCity string `json:"away_city"`
	AwayTeam string `json:"away_team"`
}

// TeamDefenseStats represents team defensive statistics
type NBATeamDefenseStats struct {
	TeamName        string  `json:"team_name"`
	OppFgaRank      int     `json:"opp_fga_rank"`
	OppFga          float64 `json:"opp_fga"`
	OppFgPctRank    int     `json:"opp_fg_pct_rank"`
	OppFgPct        float64 `json:"opp_fg_pct"`
	OppFtaRank      int     `json:"opp_fta_rank"`
	OppFta          float64 `json:"opp_fta"`
	OppFtPctRank    int     `json:"opp_ft_pct_rank"`
	OppFtPct        float64 `json:"opp_ft_pct"`
	OppRebRank      int     `json:"opp_reb_rank"`
	OppReb          float64 `json:"opp_reb"`
	OppAstRank      int     `json:"opp_ast_rank"`
	OppAst          float64 `json:"opp_ast"`
	OppFg3aRank     int     `json:"opp_fg3a_rank"`
	OppFg3a         float64 `json:"opp_fg3a"`
	OppFg3PctRank   int     `json:"opp_fg3_pct_rank"`
	OppFg3Pct       float64 `json:"opp_fg3_pct"`
	DefRatingRank   int     `json:"def_rating_rank"`
	DefRating       float64 `json:"def_rating"`
	OppPtsPaintRank int     `json:"opp_pts_paint_rank"`
	OppPtsPaint     float64 `json:"opp_pts_paint"`
	PaceRank        int     `json:"pace_rank"`
	Pace            float64 `json:"pace"`
	OppEfgPctRank   int     `json:"opp_efg_pct_rank"`
	OppEfgPct       float64 `json:"opp_efg_pct"`
	OppFtaRateRank  int     `json:"opp_fta_rate_rank"`
	OppFtaRate      float64 `json:"opp_fta_rate"`
	OppOrebPctRank  int     `json:"opp_oreb_pct_rank"`
	OppOrebPct      float64 `json:"opp_oreb_pct"`
}

type NBATeamOffenseStats struct {
	TeamName      string  `json:"team_name"`
	OffRatingRank int     `json:"off_rating_rank"`
	OffRating     float64 `json:"off_rating"`
	RebPctRank    int     `json:"reb_pct_rank"`
	RebPct        float64 `json:"reb_pct"`
	AstPctRank    int     `json:"ast_pct_rank"`
	AstPct        float64 `json:"ast_pct"`
	PaceRank      int     `json:"pace_rank"`
	Pace          float64 `json:"pace"`
	EfgPctRank    int     `json:"efg_pct_rank"`
	EfgPct        float64 `json:"efg_pct"`
	FtaRateRank   int     `json:"fta_rate_rank"`
	FtaRate       float64 `json:"fta_rate"`
	TmTovPctRank  int     `json:"tm_tov_pct_rank"`
	TmTovPct      float64 `json:"tm_tov_pct"`
	OrebPctRank   int     `json:"oreb_pct_rank"`
	OrebPct       float64 `json:"oreb_pct"`
}

// PlayerShootingSplits represents player shooting statistics
type NBAPlayerShootingSplits struct {
	PlayerName    string  `json:"player_name"`
	Fg2a          float64 `json:"fg2a"`
	Fg2m          float64 `json:"fg2m"`
	Fg2Pct        float64 `json:"fg2_pct"`
	Fg3a          float64 `json:"fg3a"`
	Fg3m          float64 `json:"fg3m"`
	Fg3Pct        float64 `json:"fg3_pct"`
	Fga           float64 `json:"fga"`
	Fgm           float64 `json:"fgm"`
	FgPct         float64 `json:"fg_pct"`
	EfgPct        float64 `json:"efg_pct"`
	Fg2aFrequency float64 `json:"fg2a_frequency"`
	Fg3aFrequency float64 `json:"fg3a_frequency"`
}

// PlayerHeadlineStats represents player headline statistics
type NBAPlayerHeadlineStats struct {
	PlayerName string  `json:"player_name"`
	Points     float64 `json:"points"`
	Assists    float64 `json:"assists"`
	Rebounds   float64 `json:"rebounds"`
}

type NBAPlayerShotChartStats struct {
	GameDate     time.Time `json:game_date`
	LocX         int       `json:"loc_x"`
	LocY         int       `json:"loc_y"`
	ShotMadeFlag int       `json:"shot_made_flag"`
	Opponent     string    `json:"opponent"`
}

type NBAPlayerAvgShotChartStats struct {
	ShotZoneBasic string  `json:"shot_zone_basic"`
	ShotZoneArea  string  `json:"shot_zone_area"`
	Attempts      int     `json:"attempts"`
	Made          int     `json:"made"`
	FgPct         float64 `json:"fg_pct"`
}

type ZoneValue struct {
	Zone      string  `json:"zone"`
	FgPct     float64 `json:"fg_pct,omitempty"`
	Fgm       float64 `json:"fgm,omitempty"`
	Fga       float64 `json:"fga,omitempty"`
	FgPctRank int     `json:"fg_pct_rank"`
	FgmRank   int     `json:"fgm_rank"`
	FgaRank   int     `json:"fga_pct_rank"`
}

type OpponentZonesResponse struct {
	Team   string               `json:"team"`
	Season string               `json:"season"`
	Zones  map[string]ZoneValue `json:"zones"`
}

// PlayerModel represents the input for points prediction
type PlayerModel struct {
	OppCity string  `json:"opp_city"`
	Minutes float64 `json:"minutes"`
}

// PoissonDist represents Poisson distribution input
type PoissonDist struct {
	PredictedPoints float64 `json:"predictedPoints"`
	BookLine        float64 `json:"bookLine"`
}

// PoissonResponse represents Poisson distribution response
type PoissonResponse struct {
	Less    float64 `json:"less"`
	Greater float64 `json:"greater"`
}

// RegisterItem represents user registration input
type RegisterItem struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginItem represents user login input
type LoginItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type NFLPlayer struct {
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
}

type NFLPlayerRushingStats struct {
	AvgGain                   float64 `json:"avgGain"`
	LongRushing               int     `json:"longRushing"`
	NetTotalYards             int     `json:"netTotalYards"`
	NetYardsPerGame           float64 `json:"netYardsPerGame"`
	RushingAttempts           int     `json:"rushingAttempts"`
	RushingBigPlays           int     `json:"rushingBigPlays"`
	RushingFirstDowns         int     `json:"rushingFirstDowns"`
	RushingFumbles            int     `json:"rushingFumbles"`
	RushingFumblesLost        int     `json:"rushingFumblesLost"`
	RushingTouchdowns         int     `json:"rushingTouchdowns"`
	RushingYards              int     `json:"rushingYards"`
	RushingYardsPerGame       float64 `json:"rushingYardsPerGame"`
	Stuffs                    int     `json:"stuffs"`
	StuffYardsLost            int     `json:"stuffYardsLost"`
	TeamGamesPlayed           int     `json:"teamGamesPlayed"`
	TotalOffensivePlays       int     `json:"totalOffensivePlays"`
	TotalPointsPerGame        float64 `json:"totalPointsPerGame"`
	TotalTouchdowns           int     `json:"totalTouchdowns"`
	TotalYards                int     `json:"totalYards"`
	TotalYardsFromScrimmage   int     `json:"totalYardsFromScrimmage"`
	TwoPointRushConvs         int     `json:"twoPointRushConvs"`
	TwoPtRush                 int     `json:"twoPtRush"`
	TwoPtRushAttempts         int     `json:"twoPtRushAttempts"`
	YardsFromScrimmagePerGame float64 `json:"yardsFromScrimmagePerGame"`
	YardsPerGame              float64 `json:"yardsPerGame"`
	YardsPerRushAttempt       float64 `json:"yardsPerRushAttempt"`
	PlayerName                string  `json:"player_name"`
}

type NFLPlayerReceivingStats struct {
	AvgGain                   float64 `json:"avgGain"`
	LongReception             int     `json:"longReception"`
	NetTotalYards             int     `json:"netTotalYards"`
	NetYardsPerGame           float64 `json:"netYardsPerGame"`
	ReceivingBigPlays         int     `json:"receivingBigPlays"`
	ReceivingFirstDowns       int     `json:"receivingFirstDowns"`
	ReceivingFumbles          int     `json:"receivingFumbles"`
	ReceivingFumblesLost      int     `json:"receivingFumblesLost"`
	ReceivingTargets          int     `json:"receivingTargets"`
	ReceivingTouchdowns       int     `json:"receivingTouchdowns"`
	ReceivingYards            int     `json:"receivingYards"`
	ReceivingYardsAfterCatch  int     `json:"receivingYardsAfterCatch"`
	ReceivingYardsAtCatch     int     `json:"receivingYardsAtCatch"`
	ReceivingYardsPerGame     float64 `json:"receivingYardsPerGame"`
	Receptions                int     `json:"receptions"`
	TeamGamesPlayed           int     `json:"teamGamesPlayed"`
	TotalOffensivePlays       int     `json:"totalOffensivePlays"`
	TotalPointsPerGame        float64 `json:"totalPointsPerGame"`
	TotalTouchdowns           int     `json:"totalTouchdowns"`
	TotalYards                int     `json:"totalYards"`
	TotalYardsFromScrimmage   int     `json:"totalYardsFromScrimmage"`
	TwoPointRecConvs          int     `json:"twoPointRecConvs"`
	TwoPtReception            int     `json:"twoPtReception"`
	TwoPtReceptionAttempts    int     `json:"twoPtReceptionAttempts"`
	YardsFromScrimmagePerGame float64 `json:"yardsFromScrimmagePerGame"`
	YardsPerGame              float64 `json:"yardsPerGame"`
	YardsPerReception         float64 `json:"yardsPerReception"`
	PlayerName                string  `json:"player_name"`
}

type NFLPlayerPassingStats struct {
	AvgGain                float64 `json:"avgGain"`
	CompletionPct          float64 `json:"completionPct"`
	Completions            int     `json:"completions"`
	InterceptionPct        float64 `json:"interceptionPct"`
	Interceptions          int     `json:"interceptions"`
	LongPassing            int     `json:"longPassing"`
	NetPassingYards        int     `json:"netPassingYards"`
	NetPassingYardsPerGame float64 `json:"netPassingYardsPerGame"`
	NetTotalYards          int     `json:"netTotalYards"`
	NetYardsPerGame        float64 `json:"netYardsPerGame"`
	PassingAttempts        int     `json:"passingAttempts"`
	PassingYards           int     `json:"passingYards"`
	TotalOffensivePlays    int     `json:"totalOffensivePlays"`
	PlayerName             string  `json:"player_name"`
}

type NFLEvent struct {
	EventID   string `json:"event_id"`
	EventDate string `json:"event_date"`
	EventWeek int    `json:"event_week"`
}

type NFLPlayerRushingReceivingGamelogStats struct {
	GameID              string    `json:"game_id"`
	PlayerName          string    `json:"player_name"`
	GameDate            time.Time `json:"game_date"`
	GameWeek            int       `json:"game_week"`
	RushingAttempts     int       `json:"rushingAttempts"`
	YardsPerRushAttempt float64   `json:"yardsPerRushAttempt"`
	RushingYards        int       `json:"rushingYards"`
	RushingTouchdowns   int       `json:"rushingTouchdowns"`
	LongRushing         int       `json:"longRushing"`
	Receptions          int       `json:"receptions"`
	ReceivingTargets    int       `json:"receivingTargets"`
	ReceivingYards      int       `json:"receivingYards"`
	YardsPerReception   float64   `json:"yardsPerReception"`
	ReceivingTouchdowns int       `json:"receivingTouchdowns"`
	LongReception       int       `json:"longReception"`
	Fumbles             int       `json:"fumbles"`
	FumblesLost         int       `json:"fumblesLost"`
	OffenseSnaps        int       `json:"offenseSnaps"`
	OffenseSnapPct      float64   `json:"offenseSnapPct"`
}

type NFLPlayerPassingGamelogStats struct {
	GameID              string    `json:"game_id"`
	PlayerName          string    `json:"player_name"`
	GameDate            time.Time `json:"game_date"`
	GameWeek            int       `json:"game_week"`
	RushingAttempts     int       `json:"rushingAttempts"`
	YardsPerRushAttempt float64   `json:"yardsPerRushAttempt"`
	RushingYards        int       `json:"rushingYards"`
	RushingTouchdowns   int       `json:"rushingTouchdowns"`
	LongRushing         int       `json:"longRushing"`
	PassingAttempts     int       `json:"passingAttempts"`
	PassingCompletions  int       `json:"passingCompletions"`
	PassingYards        int       `json:"passingYards"`
	PassingTouchdowns   int       `json:"passingTouchdowns"`
	Interceptions       int       `json:"interceptions"`
	QBRating            float64   `json:"QBRating"`
	YardsPerPassAttempt float64   `json:"yardsPerPassAttempt"`
	OffenseSnaps        int       `json:"offenseSnaps"`
	OffenseSnapPct      float64   `json:"offenseSnapPct"`
}

type NFLPlayerGamelogCollection[T any] struct {
	Games []T `json:"games"`
}

type NFLTeamDefenseStats struct {
	TeamName                        string  `json:"team_name"`
	SacksRate                       string  `json:"sacks_rate"`
	SacksRateRank                   int     `json:"sacks_rate_rank"`
	EPAperPlayAllowed               float64 `json:"epa_per_play_allowed"`
	SuccessRateAllowed              string  `json:"success_rate_allowed"`
	RushEPAAllowed                  float64 `json:"rush_epa_allowed"`
	RushEPAAllowedRank              int     `json:"rush_epa_allowed_rank"`
	RushSuccessRateAllowed          string  `json:"rush_success_rate_allowed"`
	DropbackEPAAllowed              string  `json:"dropback_epa_allowed"`
	DropbackEPAAllowedRank          int     `json:"dropback_epa_allowed_rank"`
	DropbackSuccessRateAllowed      string  `json:"dropback_success_rate_allowed"`
	EPAperPlayAllowedRank           string  `json:"epa_per_play_allowed_rank"`
	SuccessRateAllowedRank          string  `json:"success_rate_allowed_rank"`
	RushSuccessRateAllowedRank      string  `json:"rush_success_rate_allowed_rank"`
	DropbackSuccessRateAllowedRank  string  `json:"dropback_success_rate_allowed_rank"`
	ExplosivePlayRateAllowed        float64 `json:"explosive_play_rate_allowed"`
	ExplosivePlayRateAllowedRank    string  `json:"explosive_play_rate_allowed_rank"`
	PressureRate                    float64 `json:"pressure_rate"`
	PressureRateRank                string  `json:"pressure_rate_rank"`
	BlitzRate                       float64 `json:"blitz_rate"`
	BlitzRateRank                   string  `json:"blitz_rate_rank"`
	ManRate                         float64 `json:"man_rate"`
	ManRateRank                     string  `json:"man_rate_rank"`
	ZoneRate                        float64 `json:"zone_rate"`
	ZoneRateRank                    string  `json:"zone_rate_rank"`
	RushStuffRate                   float64 `json:"rush_stuff_rate"`
	RushStuffRateRank               string  `json:"rush_stuff_rate_rank"`
	YardsBeforeContactPerRbRush     float64 `json:"yards_before_contact_per_rb_rush"`
	YardsBeforeContactPerRbRushRank string  `json:"yards_before_contact_per_rb_rush_rank"`
	DownConversionRateAllowed       float64 `json:"down_conversion_rate_allowed"`
	DownConversionRateAllowedRank   string  `json:"down_conversion_rate_allowed_rank"`
	YardsPerPlayAllowed             float64 `json:"yards_per_play_allowed"`
	YardsPerPlayAllowedRank         string  `json:"yards_per_play_allowed_rank"`
	Adot                            float64 `json:"adot"`
	AdotRank                        string  `json:"adot_rank"`
	ScrambleRate                    string  `json:"scramble_rate"`
	ScrambleRateRank                string  `json:"scramble_rate_rank"`
	IntRate                         string  `json:"int_rate"`
	IntRateRank                     string  `json:"int_rate_rank"`
}

type NFLTeamOffenseStats struct {
	TeamName                string  `json:"team_name"`
	EPAperPlay              float64 `json:"epa_per_play"`
	DropbackEPA             float64 `json:"dropback_epa"`
	DropbackEPARank         int     `json:"dropback_epa_rank"`
	RushEPA                 float64 `json:"rush_epa"`
	RushEPARank             int     `json:"rush_epa_rank"`
	SuccessRate             string  `json:"success_rate"`
	RushSuccessRate         string  `json:"rush_success_rate"`
	DropbackSuccessRate     string  `json:"dropback_success_rate"`
	EPAperPlayRank          int     `json:"epa_per_play_rank"`
	SuccessRateRank         int     `json:"success_rate_rank"`
	RushSuccessRateRank     int     `json:"rush_success_rate_rank"`
	DropbackSuccessRateRank int     `json:"dropback_success_rate_rank"`
	PassingYardsPerGame     float64 `json:"passingYardsPerGame"`
	PassingYardsPerGameRank int     `json:"passingYardsPerGame_rank"`
	YardsPerCompletion      float64 `json:"yardsPerCompletion"`
	YardsPerCompletionRank  int     `json:"yardsPerCompletion_rank"`
	Sacks                   string  `json:"sacks"`
	SacksRank               int     `json:"sacks_rank"`
	RushingAttempts         float64 `json:"rushingAttempts"`
	RushingAttemptsRank     int     `json:"rushingAttempts_rank"`
	YardsPerRushAttempt     float64 `json:"yardsPerRushAttempt"`
	YardsPerRushAttemptRank int     `json:"yardsPerRushAttempt_rank"`
	PassingAttempts         float64 `json:"passingAttempts"`
	PassingAttemptsRank     int     `json:"passingAttempts_rank"`
	Adot                    float64 `json:"adot"`
	AdotRank                int     `json:"adot_rank"`
	ScrambleRate            string  `json:"scramble_rate"`
	ScrambleRateRank        int     `json:"scramble_rate_rank"`
	IntRate                 string  `json:"int_rate"`
	IntRateRank             int     `json:"int_rate_rank"`
}

type NFLPassingPBPStats struct {
	Week         int    `json:"week"`
	Opponent     string `json:"opponent"`
	CompletePass int    `json:"complete_pass"`
	Interception int    `json:"interception"`
	AirYards     int    `json:"air_yards"`
	PassLocation string `json:"pass_location"`
	PassLength   string `json:"pass_length"`
}

type Odds struct {
	Name      string  `json:"name"`
	Market    string  `json:"market"`
	Sportbook string  `json:"sportbook"`
	Line      float32 `json:"line"`
	Over      int     `json:"over"`
	Under     int     `json:"under"`
}

type MoneylineOdds struct {
	Team      string `json:"team"`
	Sportbook string `json:"sportbook"`
	Price     string `json:price`
}
