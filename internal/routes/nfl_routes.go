package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"sports_api/internal/handlers"
)

// SetupNFLRoutes configures all NFL-related routes
func SetupNFLRoutes(router *gin.RouterGroup, db *sql.DB) {
	playerHandler := handlers.NewPlayerHandler(db)

	// NFL routes
	nfl := router.Group("/nfl")
	{
		nfl.GET("/teams", playerHandler.GetAllTeams)
		nfl.GET("/team-roster/:team", playerHandler.GetPlayersByTeam)
		nfl.GET("/players/:player/rushing-stats", playerHandler.GetPlayerRushingStats)
		nfl.GET("/players/:player/receiving-stats", playerHandler.GetPlayerReceivingStats)
		nfl.GET("/players/:player/game-stats", playerHandler.GetRushingGameStats)
	}
} 