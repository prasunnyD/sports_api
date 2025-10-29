package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"sports_api/internal/database"
	"sports_api/internal/routes"
)

func main() {
	// ------------------------------------------------------------------
	// Env
	// ------------------------------------------------------------------
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using system environment variables")
	}

	// Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// ------------------------------------------------------------------
	// DB
	// ------------------------------------------------------------------
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// ------------------------------------------------------------------
	// Router + middleware
	// ------------------------------------------------------------------
	// Use gin.New so we can choose the default middleware explicitly
	r := gin.New()

	r.Use(func(c *gin.Context) {
	log.Printf("[REQ] %s %s?%s params=%v",
		c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery, c.Params)
	c.Next()
})

	r.Use(gin.Logger(), gin.Recovery())

	// CORS (rs/cors wrapped for Gin)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",          // Vite dev UI
			"http://localhost:8080",          // Same-origin manual tests
			"https://www.sharpr-analytics.com",
			"https://sharpr-analytics.com",
			"https://api.sharpr-analytics.com",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})
	r.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Helpful NoRoute so 404s show exactly what path was hit
	r.NoRoute(func(c *gin.Context) {
		log.Printf("[NoRoute] %s %s", c.Request.Method, c.Request.URL.Path)
		c.String(404, "404 page not found")
	})

	// ------------------------------------------------------------------
	// Routes
	// ------------------------------------------------------------------
	routes.SetupRoutes(r, db)

	// Log all registered routes at startup
	for _, rt := range r.Routes() {
		log.Printf("ROUTE %-6s %-60s -> %s", rt.Method, rt.Path, rt.Handler)
	}

	// ------------------------------------------------------------------
	// Serve
	// ------------------------------------------------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
