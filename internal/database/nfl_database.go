package database

import (
	"database/sql"
	"fmt"
	"log/slog"
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
		&playerPassingStats.AvgGain,                // avgGain
		&playerPassingStats.CompletionPct,          // completionPct
		&playerPassingStats.Completions,            // completions
		&playerPassingStats.InterceptionPct,        // interceptionPct
		&playerPassingStats.Interceptions,          // interceptions
		&playerPassingStats.LongPassing,            // longPassing
		&playerPassingStats.NetPassingYards,        // netPassingYards
		&playerPassingStats.NetPassingYardsPerGame, // netPassingYardsPerGame
		&playerPassingStats.NetTotalYards,          // netTotalYards
		&playerPassingStats.NetYardsPerGame,        // netYardsPerGame
		&playerPassingStats.PassingAttempts,        // passingAttempts
		&playerPassingStats.PassingYards,           // passingYards
		&playerPassingStats.TotalOffensivePlays,    // totalOffensivePlays
		&playerPassingStats.PlayerName,             // player_name
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
			COALESCE(ps.offense_snaps::int, 0) as offense_snaps,
			COALESCE(ps.offense_snap_pct, 0) as offense_snap_pct,
		FROM nfl_data.nfl_player_gamelog gl
		JOIN nfl_data.nfl_player_snap_counts ps 
				ON gl.player_id = ps.player_id 
				AND gl.season = ps.season 
				AND gl.game_week = ps.game_week
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
	return models.NFLPlayerGamelogCollection[models.NFLPlayerRushingReceivingGamelogStats]{
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
			COALESCE(ps.offense_snaps::int, 0)            AS offense_snaps,
			COALESCE(ps.offense_snap_pct, 0)              AS offense_snap_pct,
			COALESCE(gl.rushingAttempts::int, 0)          AS rushingAttempts,
			COALESCE(gl.yardsPerRushAttempt::int, 0)      AS yardsPerRushAttempt,
			COALESCE(gl.rushingYards::int, 0)             AS rushingYards,
			COALESCE(gl.rushingTouchdowns::int, 0)        AS rushingTouchdowns,
			COALESCE(gl.longRushing::int, 0)              AS longRushing,
			COALESCE(gl.passingAttempts::int, 0)          AS passingAttempts,
			COALESCE(gl.completions::int, 0)              AS completions,
			COALESCE(gl.passingYards::int, 0)             AS passingYards,
			COALESCE(gl.passingTouchdowns::int, 0)        AS passingTouchdowns,
			COALESCE(gl.interceptions::int, 0)            AS interceptions,
			COALESCE(gl.QBRating::double, 0)              AS QBRating,
			COALESCE(gl.yardsPerPassAttempt::double, 0)   AS yardsPerPassAttempt
		FROM 
			nfl_data.nfl_qb_gamelog gl
			JOIN nfl_data.nfl_player_snap_counts ps 
				ON gl.player_id = ps.player_id 
				AND gl.season = ps.season 
				AND gl.game_week = ps.game_week
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
	return models.NFLPlayerGamelogCollection[models.NFLPlayerPassingGamelogStats]{
		Games: games,
	}, nil
}

func GetNFLTeamDefenseStats(db *sql.DB, teamName string) (models.NFLTeamDefenseStats, error) {
	query := `
		SELECT
			tds.team_name,
			sumer."sack_%",
			sumer."sack_%_rank",
			sumer."epa/play",
			sumer."success_%",
			tds.rush_epa_allowed,
			tds.rush_epa_allowed_rank,
			tds.rush_success_rate_allowed,
			tds.dropback_epa_allowed,
			tds.dropback_epa_allowed_rank,
			tds.dropback_success_rate_allowed,
			sumer."epa/play_rank",
			sumer."success_%_rank",
			tds.rush_success_rate_allowed_rank,
			tds.dropback_success_rate_allowed_rank,
			sds.explosive_play_rate_allowed,
			sds.explosive_play_rate_allowed_rank,
			sds.pressure_rate,
			sds.pressure_rate_rank,
			sds.blitz_rate,
			sds.blitz_rate_rank,
			sds.man_rate,
			sds.man_rate_rank,
			sds.zone_rate,
			sds.zone_rate_rank,
			sds.rush_stuff_rate,
			sds.rush_stuff_rate_rank,
			sds.yards_before_contact_per_rb_rush,
			sds.yards_before_contact_per_rb_rush_rank,
			sds.down_conversion_rate_allowed,
			sds.down_conversion_rate_allowed_rank,
			sds.yards_per_play_allowed,
			sds.yards_per_play_allowed_rank,
			sumer.adot,
			sumer.adot_rank,
			sumer."scramble_%",
			sumer."scramble_%_rank",
			sumer."int_%",
			sumer."int_%_rank"
		FROM
			nfl_data.nfl_team_defensive_stats_db tds
			JOIN nfl_data.nfl_sharp_defense_stats sds
				ON tds.team_name = sds.team
			JOIN nfl_data.nfl_sumer_defense_stats sumer
				ON sumer.team = tds.team_name
		WHERE
			tds.team_name = ?
	`

	var teamDefenseStats models.NFLTeamDefenseStats
	err := db.QueryRow(query, teamName).Scan(
		&teamDefenseStats.TeamName,
		&teamDefenseStats.SacksRate,
		&teamDefenseStats.SacksRateRank,
		&teamDefenseStats.EPAperPlayAllowed,
		&teamDefenseStats.SuccessRateAllowed,
		&teamDefenseStats.RushEPAAllowed,
		&teamDefenseStats.RushEPAAllowedRank,
		&teamDefenseStats.RushSuccessRateAllowed,
		&teamDefenseStats.DropbackEPAAllowed,
		&teamDefenseStats.DropbackEPAAllowedRank,
		&teamDefenseStats.DropbackSuccessRateAllowed,
		&teamDefenseStats.EPAperPlayAllowedRank,
		&teamDefenseStats.SuccessRateAllowedRank,
		&teamDefenseStats.RushSuccessRateAllowedRank,
		&teamDefenseStats.DropbackSuccessRateAllowedRank,
		&teamDefenseStats.ExplosivePlayRateAllowed,
		&teamDefenseStats.ExplosivePlayRateAllowedRank,
		&teamDefenseStats.PressureRate,
		&teamDefenseStats.PressureRateRank,
		&teamDefenseStats.BlitzRate,
		&teamDefenseStats.BlitzRateRank,
		&teamDefenseStats.ManRate,
		&teamDefenseStats.ManRateRank,
		&teamDefenseStats.ZoneRate,
		&teamDefenseStats.ZoneRateRank,
		&teamDefenseStats.RushStuffRate,
		&teamDefenseStats.RushStuffRateRank,
		&teamDefenseStats.YardsBeforeContactPerRbRush,
		&teamDefenseStats.YardsBeforeContactPerRbRushRank,
		&teamDefenseStats.DownConversionRateAllowed,
		&teamDefenseStats.DownConversionRateAllowedRank,
		&teamDefenseStats.YardsPerPlayAllowed,
		&teamDefenseStats.YardsPerPlayAllowedRank,
		&teamDefenseStats.Adot,
		&teamDefenseStats.AdotRank,
		&teamDefenseStats.ScrambleRate,
		&teamDefenseStats.ScrambleRateRank,
		&teamDefenseStats.IntRate,
		&teamDefenseStats.IntRateRank,
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
			sumer."epa/play",
			oa."dropback_epa",
			oa."dropback_epa_rank",
			oa."rush_epa",
			oa."rush_epa_rank",
			sumer."success_%",
			rush_success_rate,
			dropback_success_rate,
			sumer."epa/play_rank",
			sumer."success_%_rank",
			rush_success_rate_rank,
			dropback_success_rate_rank,
			ps.passingYardsPerGame,
			ps.passingYardsPerGame_rank,
			ps.yardsPerCompletion,
			ps.yardsPerCompletion_rank,
			sumer."sack_%",
			sumer."sack_%_rank",
			rs.rushingAttempts,
			rs.rushingAttempts_rank,
			rs.yardsPerRushAttempt,
			rs.yardsPerRushAttempt_rank,
			ps.passingAttempts,
			ps.passingAttempts_rank,
			sumer.adot,
			sumer.adot_rank,
			sumer."scramble_%",
			sumer."scramble_%_rank",
			sumer."int_%",
			sumer."int_%_rank"
		from nfl_data.nfl_team_offense_advanced_stats oa
		join nfl_data.nfl_team_passing_stats_db ps on oa.team_name = ps.team_name
		join nfl_data.nfl_team_rushing_stats_db rs on oa.team_name = rs.team_name
		JOIN nfl_data.nfl_sumer_offense_stats sumer
				ON sumer.team = oa.team_name
		where oa.team_name = ?
	`

	var teamOffenseStats models.NFLTeamOffenseStats
	err := db.QueryRow(query, teamName).Scan(
		&teamOffenseStats.TeamName,
		&teamOffenseStats.EPAperPlay,
		&teamOffenseStats.DropbackEPA,
		&teamOffenseStats.DropbackEPARank,
		&teamOffenseStats.RushEPA,
		&teamOffenseStats.RushEPARank,
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
		&teamOffenseStats.Adot,
		&teamOffenseStats.AdotRank,
		&teamOffenseStats.ScrambleRate,
		&teamOffenseStats.ScrambleRateRank,
		&teamOffenseStats.IntRate,
		&teamOffenseStats.IntRateRank,
	)
	if err != nil {
		return models.NFLTeamOffenseStats{}, fmt.Errorf("failed to scan team offense stats row: %w", err)
	}
	return teamOffenseStats, nil
}

func GetNFLPassingPBPStats(db *sql.DB, playerName string, season int) ([]models.NFLPassingPBPStats, error) {
	query := `
		SELECT 
			week,
			opponent,
			complete_pass, 
			interception, 
			air_yards, 
			pass_location, 
			pass_length 
		FROM nfl_data.nfl_pbp_qb_data 
		WHERE passer = ?
		AND season = ?
	`

	var passingPBPStats []models.NFLPassingPBPStats
	rows, err := db.Query(query, playerName, season)
	if err != nil {
		return []models.NFLPassingPBPStats{}, fmt.Errorf("failed to query passing PBP stats: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var passingPBPStat models.NFLPassingPBPStats
		err := rows.Scan(
			&passingPBPStat.Week,
			&passingPBPStat.Opponent,
			&passingPBPStat.CompletePass,
			&passingPBPStat.Interception,
			&passingPBPStat.AirYards,
			&passingPBPStat.PassLocation,
			&passingPBPStat.PassLength,
		)
		if err != nil {
			return []models.NFLPassingPBPStats{}, fmt.Errorf("failed to scan passing PBP stats row: %w", err)
		}
		passingPBPStats = append(passingPBPStats, passingPBPStat)
	}
	if err = rows.Err(); err != nil {
		return []models.NFLPassingPBPStats{}, fmt.Errorf("error iterating over passing PBP stats rows: %w", err)
	}
	return passingPBPStats, nil
}

func GetNFLPropOdds(db *sql.DB, name string, market string) ([]models.Odds, error) {
	query := `SELECT 
				player,
				sport_book,
				market,
				line,
				over_odds,
				under_odds
			FROM nba_data.nfl_prop_odds t1
			WHERE "timestamp" = (
				SELECT MAX("timestamp")
				FROM nba_data.nfl_prop_odds t2
				WHERE t2.sport_book = t1.sport_book 
				AND t2.player = t1.player
			) and sport_book IN ('FanDuel', 'DraftKings', 'BetMGM') and player = ? and market = ?`
	rows, err := db.Query(query, name, market)
	if err != nil {
		return nil, fmt.Errorf("error querying odds: %w", err)
	}
	defer rows.Close()
	var odds []models.Odds
	for rows.Next() {
		var odd models.Odds
		err := rows.Scan(&odd.Name, &odd.Sportbook, &odd.Market, &odd.Line, &odd.Over, &odd.Under)
		if err != nil {
			return nil, fmt.Errorf("error in odds: %w", err)
		}
		odds = append(odds, odd)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over player odds rows: %w", err)
	}

	return odds, nil
}
