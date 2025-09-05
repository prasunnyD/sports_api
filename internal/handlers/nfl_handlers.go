package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"log/slog"

	"github.com/gin-gonic/gin"
	"sports_api/internal/database"
)

// PlayerHandler handles NFL player-related HTTP requests
type PlayerHandler struct {
	db *sql.DB
}

// NewPlayerHandler creates a new PlayerHandler instance
func NewPlayerHandler(db *sql.DB) *PlayerHandler {
	return &PlayerHandler{db: db}
}

// GetPlayersByTeam retrieves all players for a given NFL team
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

// GetAllTeams retrieves all available NFL team names
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

func (h *PlayerHandler) GetPlayerRushingStats(c *gin.Context) {
	playerName := c.Param("player")
	
	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	// Gin automatically URL-decodes the parameter, so "James%20Connor" becomes "James Connor"
	stats, err := database.GetPlayerRushingStats(h.db, playerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve player rushing stats",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"player": playerName,
		"stats":  stats,
	})
}

func (h *PlayerHandler) GetPlayerReceivingStats(c *gin.Context) {
	playerName := c.Param("player")
	
	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
		return
	}

	// Gin automatically URL-decodes the parameter, so "James%20Connor" becomes "James Connor"
	stats, err := database.GetPlayerReceivingStats(h.db, playerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve player receiving stats",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"player": playerName,
		"stats":  stats,
	})
}

func (h *PlayerHandler) GetRushingGameStats(c *gin.Context) {
	playerName := c.Param("player")
	slog.Info("Getting rushing game stats for player: %s", playerName)
	// Validate player name
	if strings.TrimSpace(playerName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Player name is required",
		})
	}

	stats, err := database.GetRushingGameStats(h.db, playerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve player rushing game stats",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"player": playerName,
		"stats":  stats,
	})
}

func (h *PlayerHandler) GetTeamDefenseStats(c *gin.Context) {
	teamName := c.Param("team")
	stats, err := database.GetNFLTeamDefenseStats(h.db, teamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve team defense stats",
			"details": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"stats": stats,
	})
}