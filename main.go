package main

import (
	"fmt"
	"log"
	"os"
	"resto_app_api/config"
	"resto_app_api/migrations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration successful")

	server := gin.Default()

	server.GET("/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the server",
			})
		},
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
