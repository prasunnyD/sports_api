package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns a simple health check response
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Sports API (NFL & NBA) is running",
		"version": "1.0.0",
	})
} 