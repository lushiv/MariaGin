package main

import (
	 _ "github.com/lushiv/go-gin-api-boilerplate/docs"
	"go-gin-api-boilerplate/db"
	"go-gin-api-boilerplate/libs/restaurants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

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

	// Routes
	r.GET("/", HealthCheck)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Create a router group for restaurant routes
	restaurantRoutes := r.Group("/restaurants")
	// Initialize the database connection for the restaurants package
	restaurants.Initialize(database)
	// Define API routes for restaurants
	restaurants.RegisterRoutes(restaurantRoutes)

	r.Run(":3000")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
