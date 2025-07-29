#!/bin/bash

# NFL API Setup Script
echo "🏈 Setting up NFL Roster API..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✅ Go version: $GO_VERSION"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy
go mod download

# Check if .env file exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp env.example .env
    echo "⚠️  Please edit .env file and add your MotherDuck token!"
    echo "   MOTHERDUCK_TOKEN=your_actual_token_here"
else
    echo "✅ .env file already exists"
fi

# Build the application
echo "🔨 Building application..."
go build -o bin/nfl-api main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo ""
    echo "🚀 To run the API:"
    echo "   ./bin/nfl-api"
    echo ""
    echo "🔧 To run in development mode:"
    echo "   go run main.go"
    echo ""
    echo "🐳 To run with Docker:"
    echo "   docker-compose up"
    echo ""
    echo "📚 For more information, see README.md"
else
    echo "❌ Build failed!"
    exit 1
fi 