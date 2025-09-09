package database

import (
	"database/sql"
	"fmt"
	"sports_api/internal/models"
	"log/slog"
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
	fmt.Printf("Getting players by team: %s\n", teamName)
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


func GetPlayerPassingStats(db *sql.DB, playerName string) (models.NFLPlayerPassingStats, error) {
	query := `
		SELECT 
			avgGain,
			completionPct,
			completions,
			interceptionPct,
			interceptions,
			longPassing,
			netPassingYards,
			netPassingYardsPerGame,
			netTotalYards,
			netYardsPerGame,
			passingAttempts,
			passingYards,
			totalOffensivePlays,
			player_name
		FROM nfl_data.nfl_passing_db
		WHERE player_name = ?
	`

	var playerPassingStats models.NFLPlayerPassingStats
	err := db.QueryRow(query, playerName).Scan(
		&playerPassingStats.AvgGain,               // avgGain
		&playerPassingStats.CompletionPct,         // completionPct
		&playerPassingStats.Completions,           // completions
		&playerPassingStats.InterceptionPct,       // interceptionPct
		&playerPassingStats.Interceptions,         // interceptions
		&playerPassingStats.LongPassing,           // longPassing
		&playerPassingStats.NetPassingYards,       // netPassingYards
		&playerPassingStats.NetPassingYardsPerGame,// netPassingYardsPerGame
		&playerPassingStats.NetTotalYards,         // netTotalYards
		&playerPassingStats.NetYardsPerGame,       // netYardsPerGame
		&playerPassingStats.PassingAttempts,       // passingAttempts
		&playerPassingStats.PassingYards,          // passingYards
		&playerPassingStats.TotalOffensivePlays,   // totalOffensivePlays
		&playerPassingStats.PlayerName,            // player_name
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NFLPlayerPassingStats{}, fmt.Errorf("no passing stats found for player: %s", playerName)
		}
		return models.NFLPlayerPassingStats{}, fmt.Errorf("failed to scan passing stats row: %w", err)
	}

	return playerPassingStats, nil
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

func GetEvents(db *sql.DB, eventType string) (models.NFLEvent, error) {
	query := `
		SELECT 
			event_id,
			event_date,
			event_week
		FROM nfl_data.nfl_game_events_db
	`
	var events models.NFLEvent
	err := db.QueryRow(query).Scan(
		&events.EventID,
		&events.EventDate,
		&events.EventWeek,
	)
	if err != nil {
		return models.NFLEvent{}, fmt.Errorf("failed to scan event row: %w", err)
	}
	return events, nil
}

func GetRushingGameStats(db *sql.DB, playerName string) (models.NFLPlayerGamelogCollection[models.NFLPlayerRushingReceivingGamelogStats], error) {
	slog.Debug("Getting rushing game stats for player: %s", playerName)
	query := `
		SELECT DISTINCT
			gl.game_id,
			gl.player_name,
			gl.rushingAttempts,
			gl.rushingYards,
			gl.rushingTouchdowns,
			gl.longRushing,
			gl.receptions,
			gl.receivingTargets,
			gl.receivingYards,
			gl.yardsPerReception,
			gl.receivingTouchdowns,
			gl.longReception,
			gl.fumbles,
			gl.fumblesLost,
			e.game_date,
			e.game_week
		FROM nfl_data.nfl_player_gamelog gl
		JOIN nfl_data.nfl_games e ON e.game_id = gl.game_id
		WHERE gl.player_name = ?
	`

	rows, err := db.Query(query, playerName)
	if err != nil {
		return models.NFLPlayerGamelogCollection[models.NFLPlayerRushingReceivingGamelogStats]{}, fmt.Errorf("failed to query gamelog stats: %w", err)
	}
	defer rows.Close()
	var games []models.NFLPlayerRushingReceivingGamelogStats
	for rows.Next() {
		var game models.NFLPlayerRushingReceivingGamelogStats
		err := rows.Scan(
			&game.GameID,
			&game.PlayerName,
			&game.RushingAttempts,
			&game.RushingYards,
			&game.RushingTouchdowns,
			&game.LongRushing,
			&game.Receptions,
			&game.ReceivingTargets,
			&game.ReceivingYards,
			&game.YardsPerReception,
			&game.ReceivingTouchdowns,
			&game.LongReception,
			&game.Fumbles,
			&game.FumblesLost,
			&game.GameDate,
			&game.GameWeek,
		)
		if err != nil {
			return models.NFLPlayerGamelogCollection[models.NFLPlayerRushingReceivingGamelogStats]{}, fmt.Errorf("failed to scan gamelog stats row: %w", err)
		}
		games = append(games, game)
	}
	if err = rows.Err(); err != nil {
		return models.NFLPlayerGamelogCollection[models.NFLPlayerRushingReceivingGamelogStats]{}, fmt.Errorf("error iterating over gamelog stats rows: %w", err)
	}
	return models.NFLPlayerGamelogCollection[models.NFLPlayerRushingReceivingGamelogStats] {
		Games: games,
	}, nil
}

func GetPassingGameStats(db *sql.DB, playerName string) (models.NFLPlayerGamelogCollection[models.NFLPlayerPassingGamelogStats], error) {
	slog.Debug("Getting passing game stats for player: %s", playerName)
	query := `
		SELECT DISTINCT
			gl.game_id,
			gl.player_name,
			e.game_date,
			e.game_week,
			gl.rushingAttempts,
			gl.yardsPerRushAttempt,
			gl.rushingYards,
			gl.rushingTouchdowns,
			gl.longRushing,
			gl.passingAttempts,
			gl.completions,
			gl.passingYards,
			gl.passingTouchdowns,
			gl.interceptions,
			gl.QBRating,
			gl.yardsPerPassAttempt
		FROM nfl_data.nfl_player_gamelog gl
		JOIN nfl_data.nfl_games e ON e.game_id = gl.game_id
		WHERE gl.player_name = ?
	`

	rows, err := db.Query(query, playerName)
	if err != nil {
		return models.NFLPlayerGamelogCollection[models.NFLPlayerPassingGamelogStats]{}, fmt.Errorf("failed to query gamelog stats: %w", err)
	}
	defer rows.Close()
	var games []models.NFLPlayerPassingGamelogStats
	for rows.Next() {
		var game models.NFLPlayerPassingGamelogStats
		err := rows.Scan(
			&game.GameID,
			&game.PlayerName,
			&game.GameDate,
			&game.GameWeek,
			&game.RushingAttempts,
			&game.YardsPerRushAttempt,
			&game.RushingYards,
			&game.RushingTouchdowns,
			&game.LongRushing,
			&game.PassingAttempts,
			&game.PassingCompletions,
			&game.PassingYards,
			&game.PassingTouchdowns,
			&game.Interceptions,
			&game.QBRating,
			&game.YardsPerPassAttempt,
		)
		if err != nil {
			return models.NFLPlayerGamelogCollection[models.NFLPlayerPassingGamelogStats]{}, fmt.Errorf("failed to scan gamelog stats row: %w", err)
		}
		games = append(games, game)
	}
	if err = rows.Err(); err != nil {
		return models.NFLPlayerGamelogCollection[models.NFLPlayerPassingGamelogStats]{}, fmt.Errorf("error iterating over gamelog stats rows: %w", err)
	}
	return models.NFLPlayerGamelogCollection[models.NFLPlayerPassingGamelogStats] {
		Games: games,
	}, nil
}

func GetNFLTeamDefenseStats(db *sql.DB, teamName string) (models.NFLTeamDefenseStats, error) {
	query := `
		select 
			team_name,
			yardsAllowed,
			pointsAllowed,
			totalTackles,
			tacklesForLoss,
			stuffs,
			stuffYards,
			avgStuffYards,
			sacks,
			sackYards,
			avgSackYards,
			passesDefended,
			passesBattedDown,
			hurries,
			defensiveTouchdowns
		from nfl_data.nfl_team_defensive_stats_db
		where team_name = ?
	`

	var teamDefenseStats models.NFLTeamDefenseStats
	err := db.QueryRow(query, teamName).Scan(
		&teamDefenseStats.TeamName,
		&teamDefenseStats.YardsAllowed,
		&teamDefenseStats.PointsAllowed,
		&teamDefenseStats.TotalTackles,
		&teamDefenseStats.TacklesForLoss,
		&teamDefenseStats.Stuffs,
		&teamDefenseStats.StuffYards,
		&teamDefenseStats.AvgStuffYards,
		&teamDefenseStats.Sacks,
		&teamDefenseStats.SackYards,
		&teamDefenseStats.AvgSackYards,
		&teamDefenseStats.PassesDefended,
		&teamDefenseStats.PassesBattedDown,
		&teamDefenseStats.Hurries,
		&teamDefenseStats.DefensiveTouchdowns,
	)
	if err != nil {
		return models.NFLTeamDefenseStats{}, fmt.Errorf("failed to scan team defense stats row: %w", err)
	}
	return teamDefenseStats, nil
}