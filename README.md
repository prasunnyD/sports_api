# Sports API (NFL & NBA)

A Go-based REST API framework that provides access to NFL and NBA data stored in MotherDuck. The API allows you to retrieve player information, team statistics, and game data for both NFL and NBA.

## Features

### NFL Features
- 🏈 Retrieve all players for a specific NFL team
- 📋 Get a list of all available NFL teams
- 🔍 Case-insensitive team name search

### NBA Features
- 🏀 Retrieve all players for a specific NBA team
- 📊 Get comprehensive team and player statistics
- 🎯 Player game logs and performance analytics
- 🛡️ Team defensive statistics
- 📈 Player shooting splits and headline stats
- 🔮 Points prediction (placeholder)
- 📊 Poisson distribution calculations
- 🏆 Live scoreboard (placeholder)

### General Features
- 🚀 Fast and efficient with DuckDB/MotherDuck
- 🔒 Secure connection with MotherDuck token authentication
- 🌐 CORS-enabled for web applications
- 📊 JSON API responses

## Prerequisites

- Go 1.21 or higher
- MotherDuck account and token
- Access to the `nfl_data.nfl_roster_db` table in MotherDuck
- Access to NBA data tables in MotherDuck (for NBA features)

## Database Schema

The API connects to a MotherDuck database with the following schemas:

### NFL Schema
```sql
CREATE TABLE nfl_roster_db(
  player_id VARCHAR,
  player_name VARCHAR,
  "position" VARCHAR,
  team_name VARCHAR,
  team_id VARCHAR
);
```

### NBA Schema
The NBA API expects the following tables in the `nba_data` schema:
- `nba_roster_db`: Player roster information
- `player_boxscores`: Individual player game statistics
- `team_boxscores`: Team game statistics
- `teams_opponent_stats`: Team opponent statistics
- `teams_defense_stats`: Team defensive statistics
- `teams_advanced_stats`: Team advanced statistics
- `teams_four_factors_stats`: Team four factors statistics
- `player_shooting_splits`: Player shooting statistics
- `player_headline_stats`: Player headline statistics

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd sports_api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up environment variables:
```bash
cp env.example .env
```

4. Edit `.env` file and add your MotherDuck token:
```bash
MOTHERDUCK_TOKEN=your_actual_motherduck_token_here
PORT=8080
GIN_MODE=debug
```

## Running the API

### Development Mode
```bash
go run main.go
```

### Production Mode
```bash
go build -o nfl-api main.go
./nfl-api
```

The API will start on port 8080 (or the port specified in your environment variables).

## API Endpoints

### Health Check
```
GET /api/v1/health
```

**Response:**
```json
{
  "status": "healthy",
  "message": "Sports API (NFL & NBA) is running",
  "version": "1.0.0"
}
```

## API Endpoints

### NFL Endpoints

#### Get Players by Team
```
GET /api/v1/nfl/players/{team_name}
```

**Parameters:**
- `team_name` (path parameter): The name of the NFL team (case-insensitive)

**Response:**
```json
{
  "team": "Kansas City Chiefs",
  "count": 53,
  "players": [
    {
      "player_id": "12345",
      "player_name": "Patrick Mahomes",
      "position": "QB",
      "team_name": "Kansas City Chiefs",
      "team_id": "KC"
    }
  ]
}
```

### Get All Teams
```
GET /api/v1/nfl/teams
```

**Response:**
```json
{
  "count": 32,
  "teams": [
    "Arizona Cardinals",
    "Atlanta Falcons",
    "Baltimore Ravens",
    "Buffalo Bills"
  ]
}
```

### NBA Endpoints

For detailed NBA API documentation, see [NBA_API_README.md](NBA_API_README.md).

#### Get All NBA Teams
```
GET /api/v1/nba/teams
```

#### Get NBA Players by Team
```
GET /api/v1/nba/players/{city}
```

#### Get Team Roster
```
GET /api/v1/nba/roster/{city}
```

#### Get Player's Last X Games
```
GET /api/v1/nba/player/{name}/last/{last_number_of_games}/games
```

#### Get Team's Last X Games
```
GET /api/v1/nba/team/{city}/last/{number_of_days}/games
```

#### Get Team Defense Stats
```
GET /api/v1/nba/{team_name}/defense-stats
```

#### Get Player Shooting Splits
```
GET /api/v1/nba/{player_name}/shooting-splits
```

#### Get Player Headline Stats
```
GET /api/v1/nba/{player_name}/headline-stats
```

#### Points Prediction
```
POST /api/v1/nba/points-prediction/{player_name}
```

#### Poisson Distribution
```
POST /api/v1/nba/poisson-dist
```

#### Get Scoreboard
```
GET /api/v1/nba/scoreboard
```

## Usage Examples

### Using curl

#### NFL Examples
1. Get players for a specific team:
```bash
curl http://localhost:8080/api/v1/nfl/players/Kansas%20City%20Chiefs
```

2. Get all available teams:
```bash
curl http://localhost:8080/api/v1/nfl/teams
```

#### NBA Examples
3. Get all NBA teams:
```bash
curl http://localhost:8080/api/v1/nba/teams
```

4. Get Lakers roster:
```bash
curl http://localhost:8080/api/v1/nba/players/Los%20Angeles
```

5. Get LeBron's last 10 games:
```bash
curl http://localhost:8080/api/v1/nba/player/LeBron%20James/last/10/games
```

6. Health check:
```bash
curl http://localhost:8080/api/v1/health
```

### Using JavaScript/Fetch

```javascript
// Get NFL players for a team
fetch('http://localhost:8080/api/v1/nfl/players/Kansas%20City%20Chiefs')
  .then(response => response.json())
  .then(data => {
    console.log(`Found ${data.count} players for ${data.team}`);
    data.players.forEach(player => {
      console.log(`${player.player_name} - ${player.position}`);
    });
  });

// Get all NFL teams
fetch('http://localhost:8080/api/v1/nfl/teams')
  .then(response => response.json())
  .then(data => {
    console.log(`Available teams: ${data.teams.join(', ')}`);
  });

// Get NBA player shooting splits
fetch('http://localhost:8080/api/v1/nba/LeBron%20James/shooting-splits')
  .then(response => response.json())
  .then(data => {
    console.log('Player shooting stats:', data);
  });

// Points prediction
fetch('http://localhost:8080/api/v1/nba/points-prediction/LeBron%20James', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    opp_city: 'Golden State',
    minutes: 32.5
  })
})
.then(response => response.json())
.then(data => {
  console.log('Projected points:', data.projected_points);
});
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- `400 Bad Request`: Invalid team name or missing parameters
- `500 Internal Server Error`: Database connection issues or query errors

Error responses include:
```json
{
  "error": "Error description",
  "details": "Additional error details"
}
```

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `MOTHERDUCK_TOKEN` | Your MotherDuck authentication token | - | Yes |
| `PORT` | Server port | 8080 | No |
| `GIN_MODE` | Gin framework mode (debug/release) | debug | No |

### CORS Configuration

The API includes CORS middleware configured to allow:
- All origins (`*`)
- Common HTTP methods (GET, POST, PUT, DELETE, OPTIONS)
- Standard headers (Origin, Content-Type, Accept, Authorization)

## Project Structure

```
sports_api/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── go.sum                  # Go module checksums
├── env.example            # Environment variables template
├── README.md              # This file
└── internal/
    ├── database/
    │   └── database.go     # Database connection and queries
    └── handlers/
        └── handlers.go     # HTTP request handlers
```

## Development

### Adding New Endpoints

1. Add the route in `main.go`
2. Create the handler function in `internal/handlers/handlers.go`
3. Add any necessary database functions in `internal/database/database.go`

### Testing

To test the API locally:

1. Ensure you have a valid MotherDuck token
2. Set up your `.env` file
3. Run the server: `go run main.go`
4. Test endpoints using curl or a tool like Postman

## Troubleshooting

### Common Issues

1. **"MOTHERDUCK_TOKEN environment variable is required"**
   - Ensure you've set the `MOTHERDUCK_TOKEN` in your `.env` file

2. **"Failed to ping database"**
   - Check your MotherDuck token is valid
   - Verify network connectivity to MotherDuck

3. **"No players found"**
   - Verify the team name exists in your database
   - Check the team name spelling and case

### Logs

The application logs connection status and errors. Check the console output for debugging information.

## License

This project is licensed under the MIT License.