package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	router.GET("/health", HealthCheck)

	// Create a test request
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
	assert.Contains(t, w.Body.String(), "NFL API is running")
}

func TestGetPlayersByTeam_EmptyTeamName(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	
	// Create a mock handler (we'll use nil for db since we're testing validation)
	handler := &PlayerHandler{db: nil}
	router.GET("/players/:team", handler.GetPlayersByTeam)

	// Create a test request with empty team name
	req, err := http.NewRequest("GET", "/players/", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Team name is required")
}

func TestGetPlayersByTeam_ValidTeamName(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	
	// Create a mock handler (we'll use nil for db since we're testing routing)
	handler := &PlayerHandler{db: nil}
	router.GET("/players/:team", handler.GetPlayersByTeam)

	// Create a test request with valid team name
	req, err := http.NewRequest("GET", "/players/Kansas%20City%20Chiefs", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Since we're using a nil database, this should return an internal server error
	// In a real test, you would mock the database
	assert.Equal(t, http.StatusInternalServerError, w.Code)
} 