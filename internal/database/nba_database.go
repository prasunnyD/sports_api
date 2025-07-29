package database

import (
	"database/sql"
	"fmt"
	"sports_api/internal/models"
)

// NBA Database operations

// GetNBAPlayersByTeam retrieves all players for a given NBA team
func GetNBAPlayersByTeam(db *sql.DB, teamCity string) ([]models.Player, error) {
	query := `
		SELECT PLAYER_ID, PLAYER, "POSITION", TeamID, NUM 
		FROM nba_data.team_roster 
		WHERE TeamID = ? 
		ORDER BY PLAYER
	`

	rows, err := db.Query(query, teamCity)
	if err != nil {
		return nil, fmt.Errorf("failed to query NBA players: %w", err)
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.PlayerID, &player.PlayerName, &player.Position, &player.TeamID, &player.Number)
		if err != nil {
			return nil, fmt.Errorf("failed to scan NBA player row: %w", err)
		}
		// Set TeamName to empty since we don't have it in the query
		player.TeamName = ""
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over NBA player rows: %w", err)
	}

	return players, nil
}

// GetNBATeams retrieves all NBA teams
func GetNBATeams(db *sql.DB) ([]models.Team, error) {
	query := `
		SELECT DISTINCT TeamID, NICKNAME, NICKNAME, TeamID 
		FROM nba_data.team_roster 
		ORDER BY NICKNAME
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query NBA teams: %w", err)
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.TeamID, &team.TeamName, &team.City, &team.Abbr)
		if err != nil {
			return nil, fmt.Errorf("failed to scan NBA team row: %w", err)
		}
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over NBA team rows: %w", err)
	}

	return teams, nil
}

// GetPlayerLastXGames retrieves a player's last X games
func GetPlayerLastXGames(db *sql.DB, playerName string, lastXGames int) (map[string]models.GameStats, error) {
	query := `
		SELECT GAME_DATE, PTS, AST, REB, FG3M, MIN 
		FROM nba_data.player_boxscores 
		WHERE PLAYER_NAME = ? 
		ORDER BY GAME_ID DESC 
		LIMIT ?
	`

	rows, err := db.Query(query, playerName, lastXGames)
	if err != nil {
		return nil, fmt.Errorf("failed to query player game logs: %w", err)
	}
	defer rows.Close()

	gameLogs := make(map[string]models.GameStats)
	for rows.Next() {
		var gameDate string
		var stats models.GameStats
		err := rows.Scan(&gameDate, &stats.Points, &stats.Assists, &stats.Rebounds, &stats.ThreePointersMade, &stats.Minutes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan player game log row: %w", err)
		}
		gameLogs[gameDate] = stats
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over player game log rows: %w", err)
	}

	return gameLogs, nil
}

// GetTeamLastXGames retrieves a team's last X games
func GetTeamLastXGames(db *sql.DB, teamCity string, lastXGames int) (map[string]models.TeamGameLog, error) {
	query := `
		SELECT GAME_DATE, PTS 
		FROM nba_data.team_boxscores 
		WHERE TEAM_CITY = ? 
		ORDER BY GAME_ID DESC 
		LIMIT ?
	`

	rows, err := db.Query(query, teamCity, lastXGames)
	if err != nil {
		return nil, fmt.Errorf("failed to query team game logs: %w", err)
	}
	defer rows.Close()

	gameLogs := make(map[string]models.TeamGameLog)
	for rows.Next() {
		var gameLog models.TeamGameLog
		err := rows.Scan(&gameLog.GameDate, &gameLog.Points)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team game log row: %w", err)
		}
		gameLogs[gameLog.GameDate] = gameLog
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over team game log rows: %w", err)
	}

	return gameLogs, nil
}

// GetTeamDefenseStats retrieves team defensive statistics
func GetTeamDefenseStats(db *sql.DB, teamName string) (*models.TeamDefenseStats, error) {
	query := `
		SELECT 
			opp.OPP_FGA_RANK, opp.OPP_FGA, opp.OPP_FG_PCT_RANK, opp.OPP_FG_PCT,
			opp.OPP_FTA_RANK, opp.OPP_FTA, opp.OPP_FT_PCT_RANK, opp.OPP_FT_PCT,
			opp.OPP_REB_RANK, opp.OPP_REB, opp.OPP_AST_RANK, opp.OPP_AST,
			opp.OPP_FG3A_RANK, opp.OPP_FG3A,
			def.DEF_RATING_RANK, def.DEF_RATING, def.OPP_PTS_PAINT_RANK, def.OPP_PTS_PAINT,
			adv.PACE_RANK, adv.PACE,
			ff.OPP_EFG_PCT_RANK, ff.OPP_EFG_PCT, ff.OPP_FTA_RATE_RANK, ff.OPP_FTA_RATE,
			ff.OPP_OREB_PCT_RANK, ff.OPP_OREB_PCT
		FROM nba_data.teams_opponent_stats opp
		JOIN nba_data.teams_defense_stats def ON opp.TEAM_ID = def.TEAM_ID
		JOIN nba_data.teams_advanced_stats adv ON opp.TEAM_ID = adv.TEAM_ID
		JOIN nba_data.teams_four_factors_stats ff ON opp.TEAM_ID = ff.TEAM_ID
		WHERE opp.TEAM_NAME = ?
		LIMIT 1
	`

	var stats models.TeamDefenseStats
	stats.TeamName = teamName

	err := db.QueryRow(query, teamName).Scan(
		&stats.OppFgaRank, &stats.OppFga, &stats.OppFgPctRank, &stats.OppFgPct,
		&stats.OppFtaRank, &stats.OppFta, &stats.OppFtPctRank, &stats.OppFtPct,
		&stats.OppRebRank, &stats.OppReb, &stats.OppAstRank, &stats.OppAst,
		&stats.OppFg3aRank, &stats.OppFg3a,
		&stats.DefRatingRank, &stats.DefRating, &stats.OppPtsPaintRank, &stats.OppPtsPaint,
		&stats.PaceRank, &stats.Pace,
		&stats.OppEfgPctRank, &stats.OppEfgPct, &stats.OppFtaRateRank, &stats.OppFtaRate,
		&stats.OppOrebPctRank, &stats.OppOrebPct,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query team defense stats: %w", err)
	}

	return &stats, nil
}

// GetPlayerShootingSplits retrieves player shooting splits
func GetPlayerShootingSplits(db *sql.DB, playerName string) (*models.PlayerShootingSplits, error) {
	query := `
		SELECT FG2A, FG2M, FG2_PCT, FG3A, FG3M, FG3_PCT, FGA, FGM, FG_PCT, EFG_PCT, FG2A_FREQUENCY, FG3A_FREQUENCY
		FROM nba_data.player_shooting_splits 
		WHERE PLAYER_NAME = ?
		LIMIT 1
	`

	var splits models.PlayerShootingSplits
	splits.PlayerName = playerName

	err := db.QueryRow(query, playerName).Scan(
		&splits.Fg2a, &splits.Fg2m, &splits.Fg2Pct, &splits.Fg3a, &splits.Fg3m, &splits.Fg3Pct,
		&splits.Fga, &splits.Fgm, &splits.FgPct, &splits.EfgPct, &splits.Fg2aFrequency, &splits.Fg3aFrequency,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query player shooting splits: %w", err)
	}

	return &splits, nil
}

// GetPlayerHeadlineStats retrieves player headline statistics
func GetPlayerHeadlineStats(db *sql.DB, playerName string) (*models.PlayerHeadlineStats, error) {
	query := `
		SELECT PTS, AST, REB
		FROM nba_data.player_headline_stats 
		WHERE PLAYER_NAME = ?
		LIMIT 1
	`

	var stats models.PlayerHeadlineStats
	stats.PlayerName = playerName

	err := db.QueryRow(query, playerName).Scan(&stats.Points, &stats.Assists, &stats.Rebounds)
	if err != nil {
		return nil, fmt.Errorf("failed to query player headline stats: %w", err)
	}

	return &stats, nil
}

// GetPlayerIDByName retrieves player ID by name
func GetPlayerIDByName(db *sql.DB, playerName string) (string, error) {
	query := `SELECT PLAYER_ID FROM nba_data.nba_roster_db WHERE PLAYER_NAME = ? LIMIT 1`
	
	var playerID string
	err := db.QueryRow(query, playerName).Scan(&playerID)
	if err != nil {
		return "", fmt.Errorf("failed to get player ID: %w", err)
	}

	return playerID, nil
}

// GetTeamIDByName retrieves team ID by name
func GetTeamIDByName(db *sql.DB, teamName string) (string, error) {
	query := `SELECT TEAM_ID FROM nba_data.nba_roster_db WHERE TEAM_NAME = ? LIMIT 1`
	
	var teamID string
	err := db.QueryRow(query, teamName).Scan(&teamID)
	if err != nil {
		return "", fmt.Errorf("failed to get team ID: %w", err)
	}

	return teamID, nil
} 