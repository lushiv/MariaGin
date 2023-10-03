package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-gin-api-boilerplate/db"             // Import your database package
	_ "go-gin-api-boilerplate/docs"         // Import generated docs package
	"go-gin-api-boilerplate/routes/v1/auth" // Import your API package
	auth_utils "go-gin-api-boilerplate/routes/v1/auth/utils"
	"go-gin-api-boilerplate/routes/v1/communication_management" // Import your API package
	"go-gin-api-boilerplate/routes/v1/restaurants"              // Import your API package
	restaurants_utils "go-gin-api-boilerplate/routes/v1/restaurants/utils"
	common_utils "go-gin-api-boilerplate/utils" // Import your database package

	// Import your API package
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MariaGin API Docs
// @version 1.0
// @description This is a sample docs.
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
	fmt.Println("Database connected...")

	err = common_utils.InitializeRedisConnection(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), 0)
	if err != nil {
		log.Fatalf("Failed to initialize the redis connection: %v", err)
	}
	fmt.Println("Redis server connected...")
	defer common_utils.RedisClient.Close() // Defer closing the Redis connection

	// Read the PORT environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	// Swagger documentation setup
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// HealthCheck route
	r.GET("/abcd/health-check", HealthCheck)

	// Create a router group for auth routes
	authRoutes := r.Group("/auth")
	auth_utils.Initialize(database)
	common_utils.Initialize(database)
	auth.RegisterRoutes(authRoutes)

	// Create a router group for communication management routes
	communicationManagementRoutes := r.Group("/communication")
	communication_management.RegisterRoutes(communicationManagementRoutes)

	// Create a router group for restaurant routes
	restaurantRoutes := r.Group("/restaurants")
	restaurants_utils.Initialize(database)
	restaurants.RegisterRoutes(restaurantRoutes)
	// Attach Logger and Recovery middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	r.Run(":" + port)
}

// HealthCheck godoc
// @Summary HealthCheck
// @Description get the status of server.
// @Tags Health Check
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /abcd/health-check [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}
