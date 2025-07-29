# Project Structure

This document explains the modular structure of the Sports API, which is organized by sport for better maintainability and extensibility.

## Overview

The API is now organized into separate modules for each sport, making it easy to:
- Add new sports without modifying existing code
- Maintain sport-specific logic independently
- Scale the API as new sports are added
- Keep the main.go file clean and focused

## Directory Structure

```
sports_api/
├── main.go                          # Application entry point (simplified)
├── go.mod                           # Go module file
├── go.sum                           # Go module checksums
├── env.example                      # Environment variables template
├── README.md                        # Main documentation
├── NBA_API_README.md               # NBA-specific documentation
├── PROJECT_STRUCTURE.md            # This file
└── internal/
    ├── database/
    │   ├── database.go              # Database connection (shared)
    │   ├── nfl_database.go          # NFL-specific database operations
    │   └── nba_database.go          # NBA-specific database operations
    ├── handlers/
    │   ├── handlers.go              # Common handlers (health check)
    │   ├── nfl_handlers.go          # NFL-specific handlers
    │   └── nba_handlers.go          # NBA-specific handlers
    ├── models/
    │   └── models.go                # All data models (shared)
    └── routes/
        ├── routes.go                # Main route configuration
        ├── nfl_routes.go            # NFL route definitions
        ├── nba_routes.go            # NBA route definitions
        └── mlb_routes.go            # Example MLB routes (placeholder)
```

## File Responsibilities

### Main Application
- **main.go**: Application entry point, server setup, middleware configuration

### Database Layer
- **database.go**: Shared database connection logic
- **nfl_database.go**: NFL-specific database queries
- **nba_database.go**: NBA-specific database queries

### Handlers Layer
- **handlers.go**: Common handlers (health check, etc.)
- **nfl_handlers.go**: NFL API endpoint handlers
- **nba_handlers.go**: NBA API endpoint handlers

### Models Layer
- **models.go**: All data structures used across the API

### Routes Layer
- **routes.go**: Main route configuration and setup
- **nfl_routes.go**: NFL route definitions
- **nba_routes.go**: NBA route definitions
- **mlb_routes.go**: Example of how to add new sports

## Benefits of This Structure

### 1. Separation of Concerns
Each sport has its own dedicated files for:
- Database operations
- API handlers
- Route definitions

### 2. Easy to Add New Sports
To add a new sport (e.g., MLB), you only need to:
1. Create `internal/database/mlb_database.go`
2. Create `internal/handlers/mlb_handlers.go`
3. Create `internal/routes/mlb_routes.go`
4. Add the route setup to `internal/routes/routes.go`

### 3. Maintainability
- Changes to one sport don't affect others
- Clear file organization makes it easy to find code
- Each sport can have its own logic and requirements

### 4. Scalability
- New sports can be added without touching existing code
- Each sport can have different database schemas
- Independent versioning and updates per sport

## Adding a New Sport

Here's the step-by-step process to add a new sport (e.g., MLB):

### 1. Create Database Operations
```go
// internal/database/mlb_database.go
package database

import (
    "database/sql"
    "fmt"
    "sports_api/internal/models"
)

// GetMLBTeams retrieves all MLB teams
func GetMLBTeams(db *sql.DB) ([]models.Team, error) {
    // Implementation here
}

// GetMLBPlayersByTeam retrieves players for a specific MLB team
func GetMLBPlayersByTeam(db *sql.DB, teamName string) ([]models.Player, error) {
    // Implementation here
}
```

### 2. Create Handlers
```go
// internal/handlers/mlb_handlers.go
package handlers

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "sports_api/internal/database"
)

type MLBHandler struct {
    db *sql.DB
}

func NewMLBHandler(db *sql.DB) *MLBHandler {
    return &MLBHandler{db: db}
}

func (h *MLBHandler) GetMLBTeams(c *gin.Context) {
    // Implementation here
}

func (h *MLBHandler) GetMLBPlayersByTeam(c *gin.Context) {
    // Implementation here
}
```

### 3. Create Routes
```go
// internal/routes/mlb_routes.go
package routes

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "sports_api/internal/handlers"
)

func SetupMLBRoutes(router *gin.RouterGroup, db *sql.DB) {
    mlbHandler := handlers.NewMLBHandler(db)

    mlb := router.Group("/mlb")
    {
        mlb.GET("/teams", mlbHandler.GetMLBTeams)
        mlb.GET("/players/:team", mlbHandler.GetMLBPlayersByTeam)
    }
}
```

### 4. Register Routes
```go
// internal/routes/routes.go
func SetupRoutes(router *gin.Engine, db *sql.DB) {
    api := router.Group("/api/v1")
    {
        api.GET("/health", handlers.HealthCheck)
        
        SetupNFLRoutes(api, db)
        SetupNBARoutes(api, db)
        SetupMLBRoutes(api, db)  // Add this line
    }
}
```

## API Endpoints Structure

The API endpoints are organized by sport:

```
/api/v1/
├── health                    # Health check (shared)
├── nfl/                      # NFL endpoints
│   ├── teams
│   └── players/:team
├── nba/                      # NBA endpoints
│   ├── teams
│   ├── players/:city
│   ├── roster/:city
│   ├── player/:name/last/:X/games
│   ├── team/:city/last/:X/games
│   ├── :team_name/defense-stats
│   ├── :player_name/shooting-splits
│   ├── :player_name/headline-stats
│   ├── points-prediction/:player_name
│   ├── poisson-dist
│   └── scoreboard
└── mlb/                      # MLB endpoints (example)
    └── health
```

## Best Practices

### 1. Consistent Naming
- Use sport abbreviations in file names (nfl_, nba_, mlb_)
- Keep function names descriptive and consistent
- Use the same pattern for all sports

### 2. Error Handling
- Implement consistent error handling across all sports
- Use appropriate HTTP status codes
- Provide meaningful error messages

### 3. Database Operations
- Keep sport-specific queries in their own files
- Use prepared statements for security
- Handle database errors gracefully

### 4. API Design
- Follow RESTful conventions
- Use consistent URL patterns
- Maintain backward compatibility

### 5. Documentation
- Document each new sport's API endpoints
- Keep README files updated
- Include usage examples

## Future Enhancements

This structure makes it easy to add:

1. **Authentication**: Per-sport or global authentication
2. **Rate Limiting**: Different limits per sport
3. **Caching**: Sport-specific caching strategies
4. **Monitoring**: Per-sport metrics and logging
5. **Versioning**: Independent API versions per sport
6. **Testing**: Sport-specific test suites

## Conclusion

This modular structure provides a solid foundation for a scalable sports API that can easily accommodate new sports while maintaining clean, organized code. Each sport is self-contained but follows consistent patterns, making the codebase both maintainable and extensible. 