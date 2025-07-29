package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"sports_api/internal/handlers"
)

// SetupNBARoutes configures all NBA-related routes
func SetupNBARoutes(router *gin.RouterGroup, db *sql.DB) {
	nbaHandler := handlers.NewNBAHandler(db)

	// NBA routes
	nba := router.Group("/nba")
	{
		nba.GET("/teams", nbaHandler.GetNBATeams)
		nba.GET("/players/:city", nbaHandler.GetNBAPlayersByTeam)
		nba.GET("/roster/:city", nbaHandler.GetTeamRoster)
		nba.GET("/player/:name/last/:last_number_of_games/games", nbaHandler.GetPlayerLastXGames)
		nba.GET("/team/:city/last/:number_of_days/games", nbaHandler.GetTeamLastXGames)
		nba.GET("/defense-stats/:team_name", nbaHandler.GetTeamDefenseStats)
		nba.GET("/shooting-splits/:player_name", nbaHandler.GetPlayerShootingSplits)
		nba.GET("/headline-stats/:player_name", nbaHandler.GetPlayerHeadlineStats)
		nba.POST("/points-prediction/:player_name", nbaHandler.PointsPrediction)
		nba.POST("/poisson-dist", nbaHandler.GetPoissonDistribution)
		nba.GET("/scoreboard", nbaHandler.GetScoreboard)
	}
} 