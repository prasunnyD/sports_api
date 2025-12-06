package handlers

import (
	"database/sql"
	"math"
	"net/http"
	"sports_api/internal/database"
	"sports_api/internal/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// NBAHandler handles NBA-related HTTP requests
type NBAHandler struct {
	db *sql.DB
}

// NewNBAHandler creates a new NBAHandler instance
func NewNBAHandler(db *sql.DB) *NBAHandler {
	return &NBAHandler{db: db}
}

// GetNBAPlayersByTeam retrieves all players for a given NBA team
func (h *NBAHandler) GetNBAPlayersByTeam(c *gin.Context) {
	teamCity := c.Param("city")

	// Validate team city
	if strings.TrimSpace(teamCity) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Team city is required",
		})
		return
	}

	// Get players from database
	players, err := database.GetNBAPlayersByTeam(h.db, teamCity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve NBA players",
			"details": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"team":    teamCity,
		"count":   len(players),
		"players": players,
	})
}

// GetNBATeams retrieves all NBA teams
func (h *NBAHandler) GetNBATeams(c *gin.Context) {
	teams, err := database.GetNBATeams(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve NBA teams",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(teams),
		"teams": teams,
	})
}

// GetPlayerLastXGames retrieves a player's last X games
func (h *NBAHandler) GetPlayerLastXGames(c *gin.Context) {
	playerName := c.Param("name")
	lastXGamesStr := c.Param("last_number_of_games")

	// Validate parameters
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	lastXGames, err := strconv.Atoi(lastXGamesStr)
	if err != nil || lastXGames <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid number of games",
		})
		return
	}

	// Get player game logs
	gameLogs, err := database.GetPlayerLastXGames(h.db, playerName, lastXGames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve player game logs",
			"details": err.Error(),
		})
		return
	}

	if len(gameLogs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No games found for player: " + playerName,
		})
		return
	}

	c.JSON(http.StatusOK, gameLogs)
}

// GetTeamLastXGames retrieves a team's last X games
func (h *NBAHandler) GetTeamLastXGames(c *gin.Context) {
	teamCity := c.Param("city")
	lastXGamesStr := c.Param("number_of_days")

	// Validate parameters
	if strings.TrimSpace(teamCity) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Team city is required",
		})
		return
	}

	lastXGames, err := strconv.Atoi(lastXGamesStr)
	if err != nil || lastXGames <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid number of games",
		})
		return
	}

	// Get team game logs
	gameLogs, err := database.GetTeamLastXGames(h.db, teamCity, lastXGames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve team game logs",
			"details": err.Error(),
		})
		return
	}

	if len(gameLogs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No games found for team: " + teamCity,
		})
		return
	}

	c.JSON(http.StatusOK, gameLogs)
}

// GetTeamRoster retrieves a team's roster
func (h *NBAHandler) GetTeamRoster(c *gin.Context) {
	teamCity := c.Param("city")

	// Validate team city
	if strings.TrimSpace(teamCity) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Team city is required",
		})
		return
	}

	// Get team roster
	players, err := database.GetNBAPlayersByTeam(h.db, teamCity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve team roster",
			"details": err.Error(),
		})
		return
	}

	if len(players) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No team members found for team: " + teamCity,
		})
		return
	}

	// Format response similar to Python API
	roster := make([]map[string]string, len(players))
	for i, player := range players {
		roster[i] = map[string]string{
			"PLAYER_ID": player.PlayerID,
			"PLAYER":    player.PlayerName,
			"NUM":       player.Number,
			"POSITION":  player.Position,
			"STATUS":    player.Status,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		teamCity: roster,
	})
}

// GetTeamDefenseStats retrieves team defensive statistics
func (h *NBAHandler) GetTeamDefenseStats(c *gin.Context) {
	teamName := c.Param("team_name")

	// Validate team name
	if strings.TrimSpace(teamName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Team name is required",
		})
		return
	}

	// Get team defense stats
	stats, err := database.GetTeamDefenseStats(h.db, teamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve team defense stats",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		teamName: stats,
	})
}

func (h *NBAHandler) GetTeamOffenseStats(c *gin.Context) {
	teamName := c.Param("team_name")

	// Validate team name
	if strings.TrimSpace(teamName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Team name is required",
		})
		return
	}

	// Get team defense stats
	stats, err := database.GetTeamOffenseStats(h.db, teamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve team defense stats",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		teamName: stats,
	})
}

// GetPlayerShootingSplits retrieves player shooting splits
func (h *NBAHandler) GetPlayerShootingSplits(c *gin.Context) {
	playerName := c.Param("player_name")

	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	// Get player shooting splits
	splits, err := database.GetPlayerShootingSplits(h.db, playerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve player shooting splits",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		playerName: splits,
	})
}

// GetPlayerHeadlineStats retrieves player headline statistics
func (h *NBAHandler) GetPlayerHeadlineStats(c *gin.Context) {
	playerName := c.Param("player_name")

	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	// Get player headline stats
	stats, err := database.GetPlayerHeadlineStats(h.db, playerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve player headline stats",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		playerName: stats,
	})
}

// PointsPrediction predicts player points (placeholder implementation)
func (h *NBAHandler) PointsPrediction(c *gin.Context) {
	_ = c.Param("player_name") // Using player name parameter but not implementing prediction logic yet

	var playerModel models.PlayerModel
	if err := c.ShouldBindJSON(&playerModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// This is a placeholder implementation
	// In a real scenario, you would load a trained model and make predictions
	// For now, we'll return a simple calculation based on minutes
	projectedPoints := playerModel.Minutes * 0.5 // Simple placeholder calculation

	c.JSON(http.StatusOK, gin.H{
		"projected_points": projectedPoints,
	})
}

// GetPoissonDistribution calculates Poisson distribution probabilities
func (h *NBAHandler) GetPoissonDistribution(c *gin.Context) {
	var poissonDist models.PoissonDist
	if err := c.ShouldBindJSON(&poissonDist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Calculate Poisson distribution
	_ = math.Round(poissonDist.PredictedPoints) // Using projected points but not implementing full calculation yet
	_ = poissonDist.BookLine                    // Using book line but not implementing full calculation yet

	// Simple Poisson calculation (placeholder)
	// In a real implementation, you would use proper Poisson distribution
	lessThan := 0.4    // Placeholder
	greaterThan := 0.6 // Placeholder

	c.JSON(http.StatusOK, gin.H{
		"less":    lessThan,
		"greater": greaterThan,
	})
}

// GetScoreboard retrieves live scoreboard (placeholder implementation)
func (h *NBAHandler) GetScoreboard(c *gin.Context) {
	scoreboard, err := database.GetScoreboard(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve player shot chart stats",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, scoreboard)
}

func (h *NBAHandler) GetPlayerShotChartStats(c *gin.Context) {
	playerName := c.Param("player_name")
	seasonID := c.Param("season_id")

	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	// Validate season ID
	if strings.TrimSpace(seasonID) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Season ID is required",
		})
		return
	}

	// Get player shot chart stats
	shots, err := database.GetPlayerShotChartStats(h.db, playerName, seasonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve player shot chart stats",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player": playerName,
		"season": seasonID,
		"shots":  shots,
	})
}

func (h *NBAHandler) GetPlayerAvgShotChartStats(c *gin.Context) {
	playerName := c.Param("player_name")
	seasonID := c.Param("season_id")

	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	// Get player avg shot chart stats
	stats, err := database.GetPlayerAvgShotChartStats(h.db, playerName, seasonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve player avg shot chart stats",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"player": playerName,
		"season": seasonID,
		"stats":  stats,
	})
}

// GetOpponentShootingByZone

func (h *NBAHandler) GetOpponentShootingByZone(c *gin.Context) {
	// 1) path params
	team := strings.TrimSpace(c.Param("team"))
	if team == "" {
		team = strings.TrimSpace(c.Param("opponent")) // <-- your routes use this
	}
	season := strings.TrimSpace(c.Param("season"))

	// 2) query params as fallback / alias
	if team == "" {
		team = strings.TrimSpace(c.Query("team"))
	}
	if team == "" {
		team = strings.TrimSpace(c.Query("opponent"))
	}
	if season == "" {
		season = strings.TrimSpace(c.Query("season"))
	}

	if season == "" {
		season = "2024-25"
	}
	if team == "" {
		c.JSON(400, models.APIResponse{
			Status:  "error",
			Message: "missing team/opponent (use path /:opponent or query ?team= / ?opponent=)",
		})
		return
	}

	resp, err := database.GetOpponentZonesByTeamSeason(h.db, team, season)
	if err != nil {
		c.JSON(500, models.APIResponse{
			Status:  "error",
			Message: "failed to load opponent zones",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, models.APIResponse{
		Status: "ok",
		Data:   resp,
	})
}

func (h *NBAHandler) GetPropOdds(c *gin.Context) {
	name := c.Param("name")
	market := c.Param("market")

	if strings.TrimSpace(name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	odds, err := database.GetPropOdds(h.db, name, market)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve odds",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, odds)
}

func (h *NBAHandler) GetMoneylineOdds(c *gin.Context) {
	team := c.Param("team")

	if strings.TrimSpace(team) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	odds, err := database.GetMoneylineOdds(h.db, team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve odds",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, odds)
}
