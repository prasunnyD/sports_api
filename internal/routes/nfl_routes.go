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
		nfl.GET("/players/:player/passing-stats", playerHandler.GetPlayerPassingStats)
		nfl.GET("/players/:player/rushing-receiving-game-stats", playerHandler.GetRushingGameStats)
		nfl.GET("/players/:player/passing-game-stats", playerHandler.GetPassingGameStats)
		nfl.GET("/team-defense-stats/:team", playerHandler.GetTeamDefenseStats)
		nfl.GET("/team-offense-stats/:team", playerHandler.GetTeamOffenseStats)
		nfl.GET("/players/:player/passing-pbp-stats/:season", playerHandler.GetNFLPassingPBPStats)
	}
} 