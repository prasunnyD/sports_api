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
			COALESCE(gl.rushingAttempts::int, 0) as rushingAttempts,
			COALESCE(gl.rushingYards::int, 0) as rushingYards,
			COALESCE(gl.rushingTouchdowns::int, 0) as rushingTouchdowns,
			COALESCE(gl.longRushing::int, 0) as longRushing,
			COALESCE(gl.receptions::int, 0) as receptions,
			COALESCE(gl.receivingTargets::int, 0) as receivingTargets,
			COALESCE(gl.receivingYards::int, 0) as receivingYards,
			COALESCE(gl.yardsPerReception::double, 0) as yardsPerReception,
			COALESCE(gl.receivingTouchdowns::int, 0) as receivingTouchdowns,
			COALESCE(gl.longReception::int, 0) as longReception,
			COALESCE(gl.fumbles::int, 0) as fumbles,
			COALESCE(gl.fumblesLost::int, 0) as fumblesLost,
			gl.game_date,
			gl.game_week,
			COALESCE(gl.offense_snaps::int, 0) as offense_snaps,
			COALESCE(gl.offense_snap_pct, 0) as offense_snap_pct,
		FROM nfl_data.nfl_player_gamelog gl
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
			&game.OffenseSnaps,
			&game.OffenseSnapPct,
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
			gl.game_date,
			gl.game_week,
			COALESCE(gl.offense_snaps::int, 0) as offense_snaps,
			COALESCE(gl.offense_snap_pct, 0) as offense_snap_pct,
			COALESCE(gl.rushingAttempts::int, 0) as rushingAttempts,
			COALESCE(gl.yardsPerRushAttempt::int, 0) as yardsPerRushAttempt,
			COALESCE(gl.rushingYards::int, 0) as rushingYards,
			COALESCE(gl.rushingTouchdowns::int, 0) as rushingTouchdowns,
			COALESCE(gl.longRushing::int, 0) as longRushing,
			COALESCE(gl.passingAttempts::int, 0) as passingAttempts,
			COALESCE(gl.completions::int, 0) as completions,
			COALESCE(gl.passingYards::int, 0) as passingYards,
			COALESCE(gl.passingTouchdowns::int, 0) as passingTouchdowns,
			COALESCE(gl.interceptions::int, 0) as interceptions,
			COALESCE(gl.QBRating::double, 0) as QBRating,
			COALESCE(gl.yardsPerPassAttempt::double, 0) as yardsPerPassAttempt
		FROM nfl_data.nfl_player_gamelog gl
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
			&game.OffenseSnaps,
			&game.OffenseSnapPct,
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
			totalTackles,
			tacklesForLoss,
			tacklesForLoss_rank,
			stuffs,
			stuffs_rank,
			stuffYards,
			avgStuffYards,
			sacks,
			sacks_rank,
			sackYards,
			avgSackYards,
			passesDefended,
			passesDefended_rank,
			hurries,
			epa_per_play_allowed,
			success_rate_allowed,
			rush_success_rate_allowed,
			dropback_success_rate_allowed,
			epa_per_play_allowed_rank,
			success_rate_allowed_rank,
			rush_success_rate_allowed_rank,
			dropback_success_rate_allowed_rank
		from nfl_data.nfl_team_defensive_stats_db
		where team_name = ?
	`

	var teamDefenseStats models.NFLTeamDefenseStats
	err := db.QueryRow(query, teamName).Scan(
		&teamDefenseStats.TeamName,
		&teamDefenseStats.TotalTackles,
		&teamDefenseStats.TacklesForLoss,
		&teamDefenseStats.TacklesForLossRank,
		&teamDefenseStats.Stuffs,
		&teamDefenseStats.StuffsRank,
		&teamDefenseStats.StuffYards,
		&teamDefenseStats.AvgStuffYards,
		&teamDefenseStats.Sacks,
		&teamDefenseStats.SacksRank,
		&teamDefenseStats.SackYards,
		&teamDefenseStats.AvgSackYards,
		&teamDefenseStats.PassesDefended,
		&teamDefenseStats.PassesDefendedRank,
		&teamDefenseStats.Hurries,
		&teamDefenseStats.EPAperPlayAllowed,
		&teamDefenseStats.SuccessRateAllowed,
		&teamDefenseStats.RushSuccessRateAllowed,
		&teamDefenseStats.DropbackSuccessRateAllowed,
		&teamDefenseStats.EPAperPlayAllowedRank,
		&teamDefenseStats.SuccessRateAllowedRank,
		&teamDefenseStats.RushSuccessRateAllowedRank,
		&teamDefenseStats.DropbackSuccessRateAllowedRank,
	)
	if err != nil {
		return models.NFLTeamDefenseStats{}, fmt.Errorf("failed to scan team defense stats row: %w", err)
	}
	return teamDefenseStats, nil
}

func GetNFLTeamOffenseStats(db *sql.DB, teamName string) (models.NFLTeamOffenseStats, error) {
	query := `
		select 
			oa.team_name,
			epa_per_play,
			success_rate,
			rush_success_rate,
			dropback_success_rate,
			epa_per_play_rank,
			success_rate_rank,
			rush_success_rate_rank,
			dropback_success_rate_rank,
			ps.passingYardsPerGame,
			ps.passingYardsPerGame_rank,
			ps.yardsPerCompletion,
			ps.yardsPerCompletion_rank,
			ps.sacks,
			ps.sacks_rank,
			rs.rushingAttempts,
			rs.rushingAttempts_rank,
			rs.yardsPerRushAttempt,
			rs.yardsPerRushAttempt_rank,
			ps.passingAttempts,
			ps.passingAttempts_rank,
		from nfl_data.nfl_team_offense_advanced_stats oa
		join nfl_data.nfl_team_passing_stats_db ps on oa.team_name = ps.team_name
		join nfl_data.nfl_team_rushing_stats_db rs on oa.team_name = rs.team_name
		where oa.team_name = ?
	`

	var teamOffenseStats models.NFLTeamOffenseStats
	err := db.QueryRow(query, teamName).Scan(
		&teamOffenseStats.TeamName,
		&teamOffenseStats.EPAperPlay,
		&teamOffenseStats.SuccessRate,
		&teamOffenseStats.RushSuccessRate,
		&teamOffenseStats.DropbackSuccessRate,
		&teamOffenseStats.EPAperPlayRank,
		&teamOffenseStats.SuccessRateRank,
		&teamOffenseStats.RushSuccessRateRank,
		&teamOffenseStats.DropbackSuccessRateRank,
		&teamOffenseStats.PassingYardsPerGame,
		&teamOffenseStats.PassingYardsPerGameRank,
		&teamOffenseStats.YardsPerCompletion,
		&teamOffenseStats.YardsPerCompletionRank,
		&teamOffenseStats.Sacks,
		&teamOffenseStats.SacksRank,
		&teamOffenseStats.RushingAttempts,
		&teamOffenseStats.RushingAttemptsRank,
		&teamOffenseStats.YardsPerRushAttempt,
		&teamOffenseStats.YardsPerRushAttemptRank,
		&teamOffenseStats.PassingAttempts,
		&teamOffenseStats.PassingAttemptsRank,
	)
	if err != nil {
		return models.NFLTeamOffenseStats{}, fmt.Errorf("failed to scan team offense stats row: %w", err)
	}
	return teamOffenseStats, nil
}