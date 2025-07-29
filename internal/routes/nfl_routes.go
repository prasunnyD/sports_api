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
		nfl.GET("/players/:team", playerHandler.GetPlayersByTeam)
		nfl.GET("/teams", playerHandler.GetAllTeams)
	}
} 