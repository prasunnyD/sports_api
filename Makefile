.PHONY: help build run test clean deps dev prod

# Default target
help:
	@echo "Available commands:"
	@echo "  deps     - Install dependencies"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  dev      - Run in development mode"
	@echo "  prod     - Run in production mode"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  docker   - Build Docker image"

# Install dependencies
deps:
	go mod tidy
	go mod download

# Build the application
build:
	go build -o bin/nfl-api main.go

# Run the application
run: build
	./bin/nfl-api

# Development mode
dev:
	go run main.go

# Production mode
prod:
	GIN_MODE=release go run main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Build Docker image
docker:
	docker build -t nfl-api .

# Docker run
docker-run:
	docker run -p 8080:8080 --env-file .env nfl-api

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Check for security vulnerabilities
security:
	gosec ./...

# Generate API documentation
docs:
	@echo "API Documentation:"
	@echo "=================="
	@echo ""
	@echo "Health Check:"
	@echo "  GET /api/v1/health"
	@echo ""
	@echo "Get Players by Team:"
	@echo "  GET /api/v1/players/{team_name}"
	@echo ""
	@echo "Get All Teams:"
	@echo "  GET /api/v1/teams"
	@echo ""
	@echo "For detailed documentation, see README.md" 