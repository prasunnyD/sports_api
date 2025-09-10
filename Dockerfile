# Build stage
FROM golang:1.24-alpine AS builder

# Install git, ca-certificates, and build dependencies for DuckDB
RUN apk add --no-cache git ca-certificates gcc musl-dev libstdc++

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./
COPY go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate go.sum and build the application
RUN go mod tidy && CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o sports-api main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and libstdc++ for runtime
RUN apk --no-cache add ca-certificates libstdc++

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/sports-api .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1

# Run the application
CMD ["./sports-api"]
