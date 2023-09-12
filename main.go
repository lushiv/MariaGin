package main

import (
	"go-gin-api-boilerplate/db"
	"go-gin-api-boilerplate/libs/restaurants"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection (after loading environment variables)
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close() // Close the database connection when the application exits

	r := gin.Default()

	// Create a router group for restaurant routes
	restaurantRoutes := r.Group("/restaurants")
	// Initialize the database connection for the restaurants package
	restaurants.Initialize(database)
	// Define API routes for restaurants
	restaurants.RegisterRoutes(restaurantRoutes)

	r.Run(":8080")
}
