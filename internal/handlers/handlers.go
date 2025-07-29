package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"sports_api/internal/database"
)

// PlayerHandler handles player-related HTTP requests
type PlayerHandler struct {
	db *sql.DB
}

// NewPlayerHandler creates a new PlayerHandler instance
func NewPlayerHandler(db *sql.DB) *PlayerHandler {
	return &PlayerHandler{db: db}
}

// HealthCheck returns a simple health check response
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "NFL API is running",
		"version": "1.0.0",
	})
}

// GetPlayersByTeam retrieves all players for a given team
func (h *PlayerHandler) GetPlayersByTeam(c *gin.Context) {
	teamName := c.Param("team")
	
	// Validate team name
	if strings.TrimSpace(teamName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Team name is required",
		})
		return
	}

	// Get players from database
	players, err := database.GetPlayersByTeam(h.db, teamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve players",
			"details": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"team":    teamName,
		"count":   len(players),
		"players": players,
	})
}

// GetAllTeams retrieves all available team names
func (h *PlayerHandler) GetAllTeams(c *gin.Context) {
	// Get teams from database
	teams, err := database.GetAllTeams(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve teams",
			"details": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"count": len(teams),
		"teams": teams,
	})
} 