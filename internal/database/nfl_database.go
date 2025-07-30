package database

import (
	"database/sql"
	"fmt"
	"sports_api/internal/models"
)

// NFL Database operations

// GetPlayersByTeam retrieves all players for a given NFL team
func GetPlayersByTeam(db *sql.DB, teamName string) ([]models.NFLPlayer, error) {
	query := `
		SELECT player_name, position
		FROM nfl_data.nfl_roster_db 
		WHERE team_name = ? 
		ORDER BY player_name
	`

	rows, err := db.Query(query, teamName)
	if err != nil {
		return nil, fmt.Errorf("failed to query players: %w", err)
	}
	defer rows.Close()

	var players []models.NFLPlayer
	for rows.Next() {
		var player models.NFLPlayer
		err := rows.Scan(&player.PlayerName, &player.Position)
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

// GetAllTeams retrieves all unique NFL team names from the database
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

func GetPlayerRushingStats(db *sql.DB, playerName string) (models.NFLPlayerRushingStats, error) {
	query := `
		SELECT 
			avgGain, longRushing, netTotalYards, netYardsPerGame,
			rushingAttempts, rushingBigPlays, rushingFirstDowns, rushingFumbles,
			rushingFumblesLost, rushingTouchdowns, rushingYards, rushingYardsPerGame,
			stuffs, stuffYardsLost, teamGamesPlayed, totalOffensivePlays,
			totalPointsPerGame, totalTouchdowns, totalYards, totalYardsFromScrimmage,
			twoPointRushConvs, twoPtRush, twoPtRushAttempts,
			yardsFromScrimmagePerGame, yardsPerGame, yardsPerRushAttempt,
			player_name
		FROM nfl_data.nfl_rushing_db
		WHERE player_name = ?
	`

	var playerRushingStats models.NFLPlayerRushingStats
	err := db.QueryRow(query, playerName).Scan(
		&playerRushingStats.AvgGain,
		&playerRushingStats.LongRushing,
		&playerRushingStats.NetTotalYards,
		&playerRushingStats.NetYardsPerGame,
		&playerRushingStats.RushingAttempts,
		&playerRushingStats.RushingBigPlays,
		&playerRushingStats.RushingFirstDowns,
		&playerRushingStats.RushingFumbles,
		&playerRushingStats.RushingFumblesLost,
		&playerRushingStats.RushingTouchdowns,
		&playerRushingStats.RushingYards,
		&playerRushingStats.RushingYardsPerGame,
		&playerRushingStats.Stuffs,
		&playerRushingStats.StuffYardsLost,
		&playerRushingStats.TeamGamesPlayed,
		&playerRushingStats.TotalOffensivePlays,
		&playerRushingStats.TotalPointsPerGame,
		&playerRushingStats.TotalTouchdowns,
		&playerRushingStats.TotalYards,
		&playerRushingStats.TotalYardsFromScrimmage,
		&playerRushingStats.TwoPointRushConvs,
		&playerRushingStats.TwoPtRush,
		&playerRushingStats.TwoPtRushAttempts,
		&playerRushingStats.YardsFromScrimmagePerGame,
		&playerRushingStats.YardsPerGame,
		&playerRushingStats.YardsPerRushAttempt,
		&playerRushingStats.PlayerName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NFLPlayerRushingStats{}, fmt.Errorf("no rushing stats found for player: %s", playerName)
		}
		return models.NFLPlayerRushingStats{}, fmt.Errorf("failed to scan rushing stats row: %w", err)
	}

	return playerRushingStats, nil
}


func GetPlayerReceivingStats(db *sql.DB, playerName string) (models.NFLPlayerReceivingStats, error) {
	query := `
		SELECT 
			avgGain,
			longReception,
			netTotalYards,
			netYardsPerGame,
			receivingBigPlays,
			receivingFirstDowns,
			receivingFumbles,
			receivingFumblesLost,
			receivingTargets,
			receivingTouchdowns,
			receivingYards,
			receivingYardsAfterCatch,
			receivingYardsAtCatch,
			receivingYardsPerGame,
			receptions,
			teamGamesPlayed,
			totalOffensivePlays,
			totalPointsPerGame,
			totalTouchdowns,
			totalYards,
			totalYardsFromScrimmage,
			twoPointRecConvs,
			twoPtReception,
			twoPtReceptionAttempts,
			yardsFromScrimmagePerGame,
			yardsPerGame,
			yardsPerReception,
			player_name
		FROM nfl_data.nfl_receiving_db
		WHERE player_name = ?
	`

	var playerReceivingStats models.NFLPlayerReceivingStats
	err := db.QueryRow(query, playerName).Scan(
		&playerReceivingStats.AvgGain,
		&playerReceivingStats.LongReception,
		&playerReceivingStats.NetTotalYards,
		&playerReceivingStats.NetYardsPerGame,
		&playerReceivingStats.ReceivingBigPlays,
		&playerReceivingStats.ReceivingFirstDowns,
		&playerReceivingStats.ReceivingFumbles,
		&playerReceivingStats.ReceivingFumblesLost,
		&playerReceivingStats.ReceivingTargets,
		&playerReceivingStats.ReceivingTouchdowns,
		&playerReceivingStats.ReceivingYards,
		&playerReceivingStats.ReceivingYardsAfterCatch,
		&playerReceivingStats.ReceivingYardsAtCatch,
		&playerReceivingStats.ReceivingYardsPerGame,
		&playerReceivingStats.Receptions,
		&playerReceivingStats.TeamGamesPlayed,
		&playerReceivingStats.TotalOffensivePlays,
		&playerReceivingStats.TotalPointsPerGame,
		&playerReceivingStats.TotalTouchdowns,
		&playerReceivingStats.TotalYards,
		&playerReceivingStats.TotalYardsFromScrimmage,
		&playerReceivingStats.TwoPointRecConvs,
		&playerReceivingStats.TwoPtReception,
		&playerReceivingStats.TwoPtReceptionAttempts,
		&playerReceivingStats.YardsFromScrimmagePerGame,
		&playerReceivingStats.YardsPerGame,
		&playerReceivingStats.YardsPerReception,
		&playerReceivingStats.PlayerName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NFLPlayerReceivingStats{}, fmt.Errorf("no receiving stats found for player: %s", playerName)
		}
		return models.NFLPlayerReceivingStats{}, fmt.Errorf("failed to scan receiving stats row: %w", err)
	}

	return playerReceivingStats, nil
}