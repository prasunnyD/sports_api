package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/marcboeker/go-duckdb"
)

// Player represents a player in the NFL roster
type Player struct {
	PlayerID   string `json:"player_id"`
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
	TeamName   string `json:"team_name"`
	TeamID     string `json:"team_id"`
}

// InitDB initializes the connection to MotherDuck
func InitDB() (*sql.DB, error) {
	// Get MotherDuck token from environment
	token := os.Getenv("MOTHERDUCK_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("MOTHERDUCK_TOKEN environment variable is required")
	}

	// Create connection string for MotherDuck
	connStr := fmt.Sprintf("md:?motherduck_token=%s", token)
	
	db, err := sql.Open("duckdb", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to MotherDuck")
	return db, nil
}

// GetPlayersByTeam retrieves all players for a given team
func GetPlayersByTeam(db *sql.DB, teamName string) ([]Player, error) {
	query := `
		SELECT player_id, player_name, position, team_name, team_id 
		FROM nfl_data.nfl_roster_db 
		WHERE LOWER(team_name) = LOWER(?) 
		ORDER BY player_name
	`

	rows, err := db.Query(query, teamName)
	if err != nil {
		return nil, fmt.Errorf("failed to query players: %w", err)
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.PlayerID, &player.PlayerName, &player.Position, &player.TeamName, &player.TeamID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan player row: %w", err)
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return players, nil
}

// GetAllTeams retrieves all unique team names from the database
func GetAllTeams(db *sql.DB) ([]string, error) {
	query := `
		SELECT DISTINCT team_name 
		FROM nfl_data.nfl_roster_db 
		ORDER BY team_name
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query teams: %w", err)
	}
	defer rows.Close()

	var teams []string
	for rows.Next() {
		var teamName string
		err := rows.Scan(&teamName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team row: %w", err)
		}
		teams = append(teams, teamName)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return teams, nil
} 