package database

import (
	"database/sql"
	"fmt"
	"os"
	"sports_api/internal/models"
)

// NBA Database operations

func GetScoreboard(db *sql.DB) ([]models.Game, error) {
	query := `SELECT game_id, home_team_city, home_team_name, away_team_city, away_team_name FROM nba_data.scoreboard`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Scoreboard: %w", err)
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		err := rows.Scan(&game.GameID, &game.HomeCity, &game.HomeTeam, &game.AwayCity, &game.AwayTeam)
		if err != nil {
			return nil, fmt.Errorf("failed to game: %w", err)
		}
		games = append(games, game)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over NBA game rows: %w", err)
	}
	return games, nil
}

// GetNBAPlayersByTeam retrieves all players for a given NBA team
func GetNBAPlayersByTeam(db *sql.DB, teamCity string) ([]models.Player, error) {
	query := `
		SELECT PLAYER_ID, PLAYER, "POSITION", TEAM, NUM 
		FROM nba_data.team_roster 
		WHERE TEAM = ? 
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
		SELECT DISTINCT TeamID, TEAM
		FROM nba_data.team_roster 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query NBA teams: %w", err)
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.TeamID, &team.City)
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
func GetPlayerLastXGames(db *sql.DB, playerName string, lastXGames int) (map[string]models.NBAGameStats, error) {
	query := `
		SELECT 
			game_date, 
			points, 
			assists, 
			reboundsTotal, 
			threePointersMade, 
			minutes_per_game
		FROM 
			nba_data.player_boxscores bx
			JOIN nba_data.team_roster tr ON bx.player_id = tr.player_id
		WHERE 
			tr.PLAYER = ?
		ORDER BY 
			GAME_ID DESC
		LIMIT ?
	`

	rows, err := db.Query(query, playerName, lastXGames)
	if err != nil {
		return nil, fmt.Errorf("failed to query player game logs: %w", err)
	}
	defer rows.Close()

	gameLogs := make(map[string]models.NBAGameStats)
	for rows.Next() {
		var gameDate string
		var stats models.NBAGameStats
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
func GetTeamDefenseStats(db *sql.DB, teamName string) (*models.NBATeamDefenseStats, error) {
	query := `
		SELECT 
			opp.OPP_FGA_RANK, 
			opp.OPP_FGA, 
			opp.OPP_FG_PCT_RANK, 
			opp.OPP_FG_PCT,
			opp.OPP_FTA_RANK, 
			opp.OPP_FTA, 
			opp.OPP_FT_PCT_RANK, 
			opp.OPP_FT_PCT,
			opp.OPP_REB_RANK, 
			opp.OPP_REB, 
			opp.OPP_AST_RANK, 
			opp.OPP_AST,
			opp.OPP_FG3A_RANK, 
			opp.OPP_FG3A, 
			opp.OPP_FG3_PCT_RANK,
			opp.OPP_FG3_PCT,
			def.DEF_RATING_RANK, 
			def.DEF_RATING, 
			def.OPP_PTS_PAINT_RANK, 
			def.OPP_PTS_PAINT,
			adv.PACE_RANK, 
			adv.PACE,
			ff.OPP_EFG_PCT_RANK, 
			ff.OPP_EFG_PCT, 
			ff.OPP_FTA_RATE_RANK, 
			ff.OPP_FTA_RATE,
			ff.OPP_OREB_PCT_RANK, 
			ff.OPP_OREB_PCT, 
			opp.TEAM_NAME
		FROM 
			nba_data.teams_opponent_stats opp
			JOIN nba_data.teams_defense_stats def ON opp.TEAM_ID = def.TEAM_ID
			JOIN nba_data.teams_advanced_stats adv ON opp.TEAM_ID = adv.TEAM_ID
			JOIN nba_data.teams_four_factors_stats ff ON opp.TEAM_ID = ff.TEAM_ID
		WHERE 
			opp.TEAM_NAME = ?
		LIMIT 1
	`

	var stats models.NBATeamDefenseStats

	err := db.QueryRow(query, teamName).Scan(
		&stats.OppFgaRank, &stats.OppFga, &stats.OppFgPctRank, &stats.OppFgPct,
		&stats.OppFtaRank, &stats.OppFta, &stats.OppFtPctRank, &stats.OppFtPct,
		&stats.OppRebRank, &stats.OppReb, &stats.OppAstRank, &stats.OppAst,
		&stats.OppFg3aRank, &stats.OppFg3a, &stats.OppFg3PctRank, &stats.OppFg3Pct,
		&stats.DefRatingRank, &stats.DefRating, &stats.OppPtsPaintRank, &stats.OppPtsPaint,
		&stats.PaceRank, &stats.Pace,
		&stats.OppEfgPctRank, &stats.OppEfgPct, &stats.OppFtaRateRank, &stats.OppFtaRate,
		&stats.OppOrebPctRank, &stats.OppOrebPct, &stats.TeamName,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query team defense stats: %w", err)
	}

	return &stats, nil
}

func GetTeamOffenseStats(db *sql.DB, teamName string) (*models.NBATeamOffenseStats, error) {
	query := `
	SELECT 
			adv.OFF_RATING_RANK, 
			adv.OFF_RATING, 
			adv.REB_PCT_RANK, 
			adv.REB_PCT,
      		adv.AST_PCT_RANK, 
			adv.AST_PCT,
			adv.PACE_RANK, 
			adv.PACE,
			ff.EFG_PCT_RANK, 
			ff.EFG_PCT, 
			ff.FTA_RATE_RANK, 
			ff.FTA_RATE,
      		ff.TM_TOV_PCT_RANK, 
			ff.TM_TOV_PCT,
			adv.OREB_PCT_RANK, 
			adv.OREB_PCT, 
			adv.TEAM_NAME
		FROM 
			nba_data.teams_advanced_stats adv
			JOIN nba_data.teams_four_factors_stats ff ON adv.TEAM_ID = ff.TEAM_ID
		WHERE 
			adv.TEAM_NAME = ?
		LIMIT 1
	`

	var stats models.NBATeamOffenseStats

	err := db.QueryRow(query, teamName).Scan(
		&stats.OffRatingRank, &stats.OffRating, &stats.RebPctRank, &stats.RebPct,
		&stats.AstPctRank, &stats.AstPct, &stats.PaceRank, &stats.Pace,
		&stats.EfgPctRank, &stats.EfgPct, &stats.FtaRateRank, &stats.FtaRate,
		&stats.TmTovPctRank, &stats.TmTovPct, &stats.OrebPctRank, &stats.OrebPct, &stats.TeamName,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query team defense stats: %w", err)
	}

	return &stats, nil
}

// GetPlayerShootingSplits retrieves player shooting splits
func GetPlayerShootingSplits(db *sql.DB, playerName string) (*models.NBAPlayerShootingSplits, error) {
	query := `
		SELECT 
			FG2A, 
			FG2M, 
			FG2_PCT, 
			FG3A, 
			FG3M, 
			FG3_PCT, 
			FGA, 
			FGM, 
			FG_PCT, 
			EFG_PCT, 
			FG2A_FREQUENCY, 
			FG3A_FREQUENCY
		FROM nba_data.player_shooting_splits ssp
		JOIN nba_data.team_roster tr on ssp.player_id = tr.player_id
		WHERE tr.PLAYER = ?
		LIMIT 1
	`

	var splits models.NBAPlayerShootingSplits
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
func GetPlayerHeadlineStats(db *sql.DB, playerName string) (*models.NBAPlayerHeadlineStats, error) {
	query := `
		SELECT 
			PTS, 
			AST, 
			REB
		FROM 
			nba_data.player_headline_stats phs
		JOIN 
			nba_data.team_roster tr 
			ON phs.player_id = tr.player_id
		WHERE 
			tr.PLAYER = ?
		LIMIT 1
	`

	var stats models.NBAPlayerHeadlineStats
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

func GetPlayerShotChartStats(db *sql.DB, playerName string, seasonID string) ([]models.NBAPlayerShotChartStats, error) {
	query := `
		SELECT
          psr.LOC_X AS x,
          psr.LOC_Y AS y,
          psr.SHOT_MADE_FLAG::INT AS made,
		  pb.OPPONENT as opponent
        FROM nba_data.player_shotchart_raw psr
		JOIN nba_data.player_boxscores pb on psr.player_id = pb.player_id and psr.game_id = pb.GAME_ID
		JOIN nba_data.team_roster tr on psr.player_id = tr.player_id
        WHERE tr.PLAYER = ?
          AND psr.SEASON_ID = ?
          AND psr.LOC_X BETWEEN -250 AND 250
          AND psr.LOC_Y BETWEEN -50 AND 470
	`

	rows, err := db.Query(query, playerName, seasonID)
	if err != nil {
		return nil, fmt.Errorf("failed to query player shot chart stats: %w", err)
	}
	defer rows.Close()

	var stats []models.NBAPlayerShotChartStats

	for rows.Next() {
		var shot models.NBAPlayerShotChartStats
		err := rows.Scan(
			&shot.LocX,
			&shot.LocY,
			&shot.ShotMadeFlag,
			&shot.Opponent,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan shot chart row: %w", err)
		}
		stats = append(stats, shot)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating shot chart rows: %w", err)
	}

	return stats, nil
}

func GetPlayerAvgShotChartStats(db *sql.DB, playerName string, seasonID string) ([]models.NBAPlayerAvgShotChartStats, error) {
	query := `
		SELECT
          SHOT_ZONE_BASIC,
          SHOT_ZONE_AREA,
          COUNT(*)::BIGINT AS attempts,
          SUM(SHOT_MADE_FLAG)::BIGINT AS made,
          (made/attempts) *100 ::DOUBLE AS fg_pct
        FROM nba_data.player_shotchart_raw psr
		JOIN nba_data.player_boxscores pb on psr.player_id = pb.player_id and psr.game_id = pb.GAME_ID
		JOIN nba_data.team_roster tr on psr.player_id = tr.player_id
        WHERE tr.PLAYER = ?
          AND psr.SEASON_ID = ?
        GROUP BY SHOT_ZONE_BASIC,SHOT_ZONE_AREA
	`
	rows, err := db.Query(query, playerName, seasonID)
	if err != nil {
		return nil, fmt.Errorf("failed to query player avg shot chart stats: %w", err)
	}
	defer rows.Close()

	var stats []models.NBAPlayerAvgShotChartStats
	for rows.Next() {
		var shotZone models.NBAPlayerAvgShotChartStats
		err := rows.Scan(&shotZone.ShotZoneBasic, &shotZone.ShotZoneArea, &shotZone.Attempts, &shotZone.Made, &shotZone.FgPct)
		if err != nil {
			return nil, fmt.Errorf("failed to scan player avg shot chart stats row: %w", err)
		}
		stats = append(stats, shotZone)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over player avg shot chart stats rows: %w", err)
	}

	return stats, nil
}

func opponentZonesTable() string {
	if t := os.Getenv("OPP_ZONES_TABLE"); t != "" {
		return t
	}
	// Keep this consistent with your other queries (schema prefix: nba_data.*).
	// If your table lives under nba_data.main.team_opponent_zones in MotherDuck,
	// you can set OPP_ZONES_TABLE to "nba_data.main.team_opponent_zones" at runtime.
	return "nba_data.team_opponent_zones"
}

// GetOpponentZonesByTeamSeason fetches opponent overall shooting by zone
// for a given team abbreviation and season.
// Returns a map keyed by region name with FGM/FGA/FG_PCT (FG_PCT is 0..1).
func GetOpponentZonesByTeamSeason(db *sql.DB, teamAbbr, season string) (*models.OpponentZonesResponse, error) {
	// FG_RANK and OUT_OF are now persisted in the table by the Python pipeline.
	query := fmt.Sprintf(`
        SELECT REGION, FGM, FGA, FG_PCT, FG_RANK, OUT_OF
        FROM %s
        WHERE SEASON = ?
          AND UPPER(TEAM_ABBR) = UPPER(?)
    `, opponentZonesTable())

	// only two args – season, teamAbbr
	rows, err := db.Query(query, season, teamAbbr)
	if err != nil {
		return nil, fmt.Errorf("failed to query opponent zones: %w", err)
	}
	defer rows.Close()

	zones := make(map[string]models.ZoneValue, 8)

	for rows.Next() {
		var (
			region      string
			fgm, fga    float64
			fgp         float64
			rank, outOf int
		)

		if err := rows.Scan(&region, &fgm, &fga, &fgp, &rank, &outOf); err != nil {
			return nil, fmt.Errorf("failed to scan opponent zone row: %w", err)
		}

		// take addresses of local vars so struct gets *float64 / *int
		fgmCopy := fgm
		fgaCopy := fga
		fgpCopy := fgp
		rankCopy := rank
		outOfCopy := outOf
		zones[region] = models.ZoneValue{
			FgPct:  &fgpCopy,
			Fgm:    &fgmCopy,
			Fga:    &fgaCopy,
			FgRank: &rankCopy,
			OutOf:  &outOfCopy,
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating opponent zone rows: %w", err)
	}

	return &models.OpponentZonesResponse{
		Team:   teamAbbr,
		Season: season,
		Zones:  zones,
	}, nil
}
