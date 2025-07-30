package models

// Player represents an NBA player
type Player struct {
	PlayerID   string `json:"player_id"`
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
	TeamName   string `json:"team_name"`
	TeamID     string `json:"team_id"`
	Number     string `json:"number"`
}

// Team represents an NBA team
type Team struct {
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	City     string `json:"city"`
	Abbr     string `json:"abbr"`
}

// GameStats represents player game statistics
type GameStats struct {
	Points           float64 `json:"points"`
	Assists          float64 `json:"assists"`
	Rebounds         float64 `json:"rebounds"`
	ThreePointersMade float64 `json:"threePointersMade"`
	Minutes          float64 `json:"minutes"`
}

// PlayerGameLog represents a player's game log entry
type PlayerGameLog struct {
	GameDate string     `json:"game_date"`
	Stats    GameStats  `json:"stats"`
}

// TeamGameLog represents a team's game log entry
type TeamGameLog struct {
	GameDate string  `json:"game_date"`
	Points   float64 `json:"points"`
}

// Game represents a live game
type Game struct {
	GameID    string `json:"game_id"`
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	HomeScore int    `json:"home_score"`
	AwayScore int    `json:"away_score"`
	Status    string `json:"status"`
}

// TeamDefenseStats represents team defensive statistics
type TeamDefenseStats struct {
	TeamName           string  `json:"team_name"`
	OppFgaRank         int     `json:"opp_fga_rank"`
	OppFga             float64 `json:"opp_fga"`
	OppFgPctRank       int     `json:"opp_fg_pct_rank"`
	OppFgPct           float64 `json:"opp_fg_pct"`
	OppFtaRank         int     `json:"opp_fta_rank"`
	OppFta             float64 `json:"opp_fta"`
	OppFtPctRank       int     `json:"opp_ft_pct_rank"`
	OppFtPct           float64 `json:"opp_ft_pct"`
	OppRebRank         int     `json:"opp_reb_rank"`
	OppReb             float64 `json:"opp_reb"`
	OppAstRank         int     `json:"opp_ast_rank"`
	OppAst             float64 `json:"opp_ast"`
	OppFg3aRank        int     `json:"opp_fg3a_rank"`
	OppFg3a            float64 `json:"opp_fg3a"`
	DefRatingRank      int     `json:"def_rating_rank"`
	DefRating          float64 `json:"def_rating"`
	OppPtsPaintRank    int     `json:"opp_pts_paint_rank"`
	OppPtsPaint        float64 `json:"opp_pts_paint"`
	PaceRank           int     `json:"pace_rank"`
	Pace               float64 `json:"pace"`
	OppEfgPctRank      int     `json:"opp_efg_pct_rank"`
	OppEfgPct          float64 `json:"opp_efg_pct"`
	OppFtaRateRank     int     `json:"opp_fta_rate_rank"`
	OppFtaRate         float64 `json:"opp_fta_rate"`
	OppOrebPctRank     int     `json:"opp_oreb_pct_rank"`
	OppOrebPct         float64 `json:"opp_oreb_pct"`
}

// PlayerShootingSplits represents player shooting statistics
type PlayerShootingSplits struct {
	PlayerName     string  `json:"player_name"`
	Fg2a           float64 `json:"fg2a"`
	Fg2m           float64 `json:"fg2m"`
	Fg2Pct         float64 `json:"fg2_pct"`
	Fg3a           float64 `json:"fg3a"`
	Fg3m           float64 `json:"fg3m"`
	Fg3Pct         float64 `json:"fg3_pct"`
	Fga            float64 `json:"fga"`
	Fgm            float64 `json:"fgm"`
	FgPct          float64 `json:"fg_pct"`
	EfgPct         float64 `json:"efg_pct"`
	Fg2aFrequency  float64 `json:"fg2a_frequency"`
	Fg3aFrequency  float64 `json:"fg3a_frequency"`
}

// PlayerHeadlineStats represents player headline statistics
type PlayerHeadlineStats struct {
	PlayerName string  `json:"player_name"`
	Points     float64 `json:"points"`
	Assists    float64 `json:"assists"`
	Rebounds   float64 `json:"rebounds"`
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
	AvgGain                  float64 `json:"avgGain"`
	LongRushing              int     `json:"longRushing"`
	NetTotalYards            int     `json:"netTotalYards"`
	NetYardsPerGame          float64 `json:"netYardsPerGame"`
	RushingAttempts          int     `json:"rushingAttempts"`
	RushingBigPlays          int     `json:"rushingBigPlays"`
	RushingFirstDowns        int     `json:"rushingFirstDowns"`
	RushingFumbles           int     `json:"rushingFumbles"`
	RushingFumblesLost       int     `json:"rushingFumblesLost"`
	RushingTouchdowns        int     `json:"rushingTouchdowns"`
	RushingYards             int     `json:"rushingYards"`
	RushingYardsPerGame      float64 `json:"rushingYardsPerGame"`
	Stuffs                   int     `json:"stuffs"`
	StuffYardsLost           int     `json:"stuffYardsLost"`
	TeamGamesPlayed          int     `json:"teamGamesPlayed"`
	TotalOffensivePlays      int     `json:"totalOffensivePlays"`
	TotalPointsPerGame       float64 `json:"totalPointsPerGame"`
	TotalTouchdowns          int     `json:"totalTouchdowns"`
	TotalYards               int     `json:"totalYards"`
	TotalYardsFromScrimmage  int     `json:"totalYardsFromScrimmage"`
	TwoPointRushConvs        int     `json:"twoPointRushConvs"`
	TwoPtRush                int     `json:"twoPtRush"`
	TwoPtRushAttempts        int     `json:"twoPtRushAttempts"`
	YardsFromScrimmagePerGame float64 `json:"yardsFromScrimmagePerGame"`
	YardsPerGame             float64 `json:"yardsPerGame"`
	YardsPerRushAttempt      float64 `json:"yardsPerRushAttempt"`
	PlayerName               string  `json:"player_name"`
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
