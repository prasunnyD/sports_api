# NBA API Documentation

This document describes the NBA API endpoints that have been added to the Sports API, replicating the functionality of the Python NBA API from [prasunnyD/NBA_Data](https://github.com/prasunnyD/NBA_Data).

## Overview

The NBA API provides endpoints for:
- Team and player data retrieval
- Game statistics and logs
- Player performance analytics
- Team defensive statistics
- Points prediction (placeholder)
- Poisson distribution calculations
- Live scoreboard (placeholder)

## Base URL

```
http://localhost:8080/api/v1/nba
```

## Endpoints

### 1. Get All NBA Teams

**GET** `/nba/teams`

Returns all NBA teams with their basic information.

**Response:**
```json
{
  "count": 30,
  "teams": [
    {
      "team_id": "1610612737",
      "team_name": "Atlanta Hawks",
      "city": "Atlanta",
      "abbr": "ATL"
    }
  ]
}
```

### 2. Get Players by Team

**GET** `/nba/players/{city}`

Returns all players for a specific team by city name.

**Parameters:**
- `city` (path): Team city (e.g., "Los Angeles", "Boston")

**Response:**
```json
{
  "team": "Los Angeles",
  "count": 15,
  "players": [
    {
      "player_id": "1628369",
      "player_name": "LeBron James",
      "position": "F",
      "team_name": "Los Angeles Lakers",
      "team_id": "1610612747",
      "number": "23"
    }
  ]
}
```

### 3. Get Team Roster

**GET** `/nba/roster/{city}`

Returns a team's roster in a simplified format.

**Parameters:**
- `city` (path): Team city

**Response:**
```json
{
  "Los Angeles": [
    {
      "PLAYER": "LeBron James",
      "NUM": "23",
      "POSITION": "F"
    }
  ]
}
```

### 4. Get Player's Last X Games

**GET** `/nba/player/{name}/last/{last_number_of_games}/games`

Returns a player's game statistics for their last X games.

**Parameters:**
- `name` (path): Player name
- `last_number_of_games` (path): Number of recent games to retrieve

**Response:**
```json
{
  "2024-01-15": {
    "points": 25.0,
    "assists": 8.0,
    "rebounds": 7.0,
    "threePointersMade": 3.0,
    "minutes": 35.0
  }
}
```

### 5. Get Team's Last X Games

**GET** `/nba/team/{city}/last/{number_of_days}/games`

Returns a team's game results for their last X games.

**Parameters:**
- `city` (path): Team city
- `number_of_days` (path): Number of recent games to retrieve

**Response:**
```json
{
  "2024-01-15": {
    "game_date": "2024-01-15",
    "points": 112.0
  }
}
```

### 6. Get Team Defense Stats

**GET** `/nba/{team_name}/defense-stats`

Returns comprehensive defensive statistics for a team.

**Parameters:**
- `team_name` (path): Full team name

**Response:**
```json
{
  "Los Angeles Lakers": {
    "team_name": "Los Angeles Lakers",
    "opp_fga_rank": 15,
    "opp_fga": 88.5,
    "opp_fg_pct_rank": 8,
    "opp_fg_pct": 0.445,
    "def_rating_rank": 12,
    "def_rating": 112.3,
    "pace_rank": 5,
    "pace": 101.2
  }
}
```

### 7. Get Player Shooting Splits

**GET** `/nba/{player_name}/shooting-splits`

Returns detailed shooting statistics for a player.

**Parameters:**
- `player_name` (path): Player name

**Response:**
```json
{
  "LeBron James": {
    "player_name": "LeBron James",
    "fg2a": 8.5,
    "fg2m": 5.2,
    "fg2_pct": 0.612,
    "fg3a": 6.8,
    "fg3m": 2.1,
    "fg3_pct": 0.309,
    "fga": 15.3,
    "fgm": 7.3,
    "fg_pct": 0.477,
    "efg_pct": 0.545,
    "fg2a_frequency": 0.556,
    "fg3a_frequency": 0.444
  }
}
```

### 8. Get Player Headline Stats

**GET** `/nba/{player_name}/headline-stats`

Returns basic headline statistics for a player.

**Parameters:**
- `player_name` (path): Player name

**Response:**
```json
{
  "LeBron James": {
    "player_name": "LeBron James",
    "points": 25.4,
    "assists": 7.8,
    "rebounds": 7.2
  }
}
```

### 9. Points Prediction

**POST** `/nba/points-prediction/{player_name}`

Predicts points for a player based on opponent and minutes.

**Parameters:**
- `player_name` (path): Player name

**Request Body:**
```json
{
  "opp_city": "Golden State",
  "minutes": 32.5
}
```

**Response:**
```json
{
  "projected_points": 16.25
}
```

### 10. Poisson Distribution

**POST** `/nba/poisson-dist`

Calculates Poisson distribution probabilities for betting odds.

**Request Body:**
```json
{
  "predictedPoints": 22.5,
  "bookLine": 21.5
}
```

**Response:**
```json
{
  "less": 0.4,
  "greater": 0.6
}
```

### 11. Get Scoreboard

**GET** `/nba/scoreboard`

Returns live game scores (placeholder implementation).

**Response:**
```json
{
  "game1": {
    "game_id": "game1",
    "home_team": "Los Angeles Lakers",
    "away_team": "Golden State Warriors",
    "home_score": 105,
    "away_score": 98,
    "status": "Final"
  }
}
```

## Database Schema

The NBA API expects the following database tables in the `nba_data` schema:

### Core Tables
- `nba_roster_db`: Player roster information
- `player_boxscores`: Individual player game statistics
- `team_boxscores`: Team game statistics

### Statistics Tables
- `teams_opponent_stats`: Team opponent statistics
- `teams_defense_stats`: Team defensive statistics
- `teams_advanced_stats`: Team advanced statistics
- `teams_four_factors_stats`: Team four factors statistics
- `player_shooting_splits`: Player shooting statistics
- `player_headline_stats`: Player headline statistics

## Environment Variables

Make sure to set the following environment variables:

```bash
MOTHERDUCK_TOKEN=your_motherduck_token
PORT=8080
GIN_MODE=release
```

## Error Handling

All endpoints return appropriate HTTP status codes:

- `200`: Success
- `400`: Bad Request (invalid parameters)
- `404`: Not Found (no data available)
- `500`: Internal Server Error (database or server error)

Error responses include:
```json
{
  "error": "Error description",
  "details": "Additional error details"
}
```

## Usage Examples

### Using curl

```bash
# Get all NBA teams
curl http://localhost:8080/api/v1/nba/teams

# Get Lakers roster
curl http://localhost:8080/api/v1/nba/players/Los%20Angeles

# Get LeBron's last 10 games
curl http://localhost:8080/api/v1/nba/player/LeBron%20James/last/10/games

# Get Lakers defense stats
curl http://localhost:8080/api/v1/nba/Los%20Angeles%20Lakers/defense-stats
```

### Using JavaScript/Fetch

```javascript
// Get player shooting splits
const response = await fetch('/api/v1/nba/LeBron James/shooting-splits');
const data = await response.json();

// Points prediction
const predictionResponse = await fetch('/api/v1/nba/points-prediction/LeBron James', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    opp_city: 'Golden State',
    minutes: 32.5
  })
});
const prediction = await predictionResponse.json();
```

## Notes

1. **Placeholder Implementations**: Some endpoints (points prediction, Poisson distribution, scoreboard) are placeholder implementations. In a production environment, these would need to be connected to actual ML models and live data sources.

2. **Database Requirements**: The API expects NBA data to be available in the MotherDuck database. You'll need to populate the required tables with NBA data.

3. **Player/Team Names**: Use exact player and team names as they appear in the database for best results.

4. **Rate Limiting**: Consider implementing rate limiting for production use.

5. **Authentication**: The current implementation doesn't include authentication. Add JWT or other authentication mechanisms for production use.

## Future Enhancements

- Add authentication and authorization
- Implement real ML models for predictions
- Add caching for frequently accessed data
- Integrate with live NBA API for real-time data
- Add more advanced analytics endpoints
- Implement proper Poisson distribution calculations
- Add player comparison endpoints
- Add season statistics endpoints 