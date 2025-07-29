package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"sports_api/internal/handlers"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// API v1 routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", handlers.HealthCheck)

		// Setup sport-specific routes
		SetupNFLRoutes(api, db)
		SetupNBARoutes(api, db)
		
		// Example of adding a new sport (MLB)
		// Uncomment the line below when MLB handlers are implemented
		// SetupMLBRoutes(api, db)
	}
} 