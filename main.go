package main

import (
	"log"
	"net/http"
	"os"

	"go-gin-api-boilerplate/db"     // Import your database package
	_ "go-gin-api-boilerplate/docs" // Import generated docs package
	"go-gin-api-boilerplate/routes/auth"
	"go-gin-api-boilerplate/routes/restaurants" // Import your API package

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Gin API Boilerplate
// @version 2.0
// @description This is a sample boilerplate server.
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

	// Read the PORT environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	// HealthCheck route
	r.GET("/health-check", HealthCheck)
	// Swagger documentation setup
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := r.Group("/auth")

	// Define authentication routes
	authGroup.POST("/login", auth.LoginCustomer)
	authGroup.POST("/logout", auth.LogoutCustomer)

	// Create a router group for restaurant routes
	restaurantRoutes := r.Group("/restaurants")
	restaurants.RegisterRoutes(restaurantRoutes)

	r.Run(":" + port)
}

// HealthCheck godoc
// @Summary HealthCheck
// @Description get the status of server.
// @Tags Health Check
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}
