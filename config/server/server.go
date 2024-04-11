package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/the-feed/config/database"
	"github.com/the-feed/internal/routes"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database configuration
	config := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Set the listen address
	listenAddr := os.Getenv("HOST_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	// Initialize the client IP address
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	// Serve the static assets
	router.Static("assets", "./assets")

	// Load the HTML templates
	files := []string{
		"views/user/authentication/login.html", "views/user/authentication/signup.html",
		"views/user/authentication/logout.html", "views/user/account/update.html",
		"views/user/account/delete.html",
		"views/layout/footer.html", "views/layout/header.html",
		"views/app/dashboard.html", "views/app/index.html", "views/errors.html",
		"views/app/feed.html",
	}
	router.LoadHTMLFiles(files...)

	// Initialize the database
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Add CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true // Allow sending cookies from frontend to backend
	corsConfig.AllowedHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(corsConfig))

	// Initialize the routes and run the server
	routes.AuthRoutes(router, db)
	router.Run(listenAddr)
}
