package handlers

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"sports_api/internal/database"
	"sports_api/internal/models"
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
			"error": "Failed to retrieve NBA players",
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
			"error": "Failed to retrieve NBA teams",
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
			"error": "Failed to retrieve player game logs",
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
			"error": "Failed to retrieve team game logs",
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
			"error": "Failed to retrieve team roster",
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
			"PLAYER":   player.PlayerName,
			"NUM":      player.Number,
			"POSITION": player.Position,
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
			"error": "Failed to retrieve team defense stats",
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
			"error": "Failed to retrieve player shooting splits",
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
			"error": "Failed to retrieve player headline stats",
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
	_ = poissonDist.BookLine // Using book line but not implementing full calculation yet

	// Simple Poisson calculation (placeholder)
	// In a real implementation, you would use proper Poisson distribution
	lessThan := 0.4  // Placeholder
	greaterThan := 0.6 // Placeholder

	c.JSON(http.StatusOK, gin.H{
		"less":    lessThan,
		"greater": greaterThan,
	})
}

// GetScoreboard retrieves live scoreboard (placeholder implementation)
func (h *NBAHandler) GetScoreboard(c *gin.Context) {
	// This is a placeholder implementation
	// In a real scenario, you would fetch live data from NBA API
	scoreboard := map[string]models.Game{
		"game1": {
			GameID:    "game1",
			HomeTeam:  "Los Angeles Lakers",
			AwayTeam:  "Golden State Warriors",
			HomeScore: 105,
			AwayScore: 98,
			Status:    "Final",
		},
		"game2": {
			GameID:    "game2",
			HomeTeam:  "Boston Celtics",
			AwayTeam:  "Miami Heat",
			HomeScore: 112,
			AwayScore: 108,
			Status:    "Q4",
		},
	}

	c.JSON(http.StatusOK, scoreboard)
} 