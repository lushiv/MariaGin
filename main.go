package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go-gin-api-boilerplate/db"                  // Import your database package
	_ "go-gin-api-boilerplate/docs"              // Import generated docs package
	v1_routes "go-gin-api-boilerplate/routes/v1" // Import your API package

	// Import your API package
	common_utils "go-gin-api-boilerplate/utils" // Import your database package
	logger "go-gin-api-boilerplate/utils"       // Import your database package

	// Import your API package
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type AppConfig struct {
	DatabaseEnabled bool `json:"database_enabled"`
	RedisEnabled    bool `json:"redis_enabled"`
	RabbitMQEnabled bool `json:"rabbitmq_enabled"`
	FirebaseEnabled bool `json:"firebase_enabled"`
}

var AppSettings AppConfig

func LoadConfig() error {
	// Load and parse the configuration file
	data, err := ioutil.ReadFile("configs/app_config.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &AppSettings); err != nil {
		return err
	}

	return nil
}

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
	// Initialize the logger
	logger.InitLogger()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file")
	}

	// Load configuration settings
	if err := LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := gin.Default()

	// Initialize the database connection if enabled

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close() // Close the database connection when the application exits
	fmt.Println("Database connected...")

	// Initialize the redis connection if enabled
	if AppSettings.RedisEnabled {
		err := common_utils.InitializeRedisConnection(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), 0)
		if err != nil {
			log.Fatalf("Failed to initialize the redis connection: %v", err)
		}
		defer common_utils.RedisClient.Close() // Defer closing the Redis connection
		fmt.Println("Redis server connected...")
	}

	// Initialize RabbitMQ connection if enabled
	if AppSettings.RabbitMQEnabled {
		rabbitMQURI := os.Getenv("RABBITMQ_URI")
		if rabbitMQURI == "" {
			log.Fatal("RABBITMQ_URI environment variable not set")
		}
		err := common_utils.InitializeRabbitMQConnection(rabbitMQURI)
		if err != nil {
			log.Fatalf("Failed to initialize the RabbitMQ connection: %v", err)
		}
		defer common_utils.CloseRabbitMQConnection() // Defer closing the RabbitMQ connection
		fmt.Println("RabbitMQ connected...")
	}

	// Initialize Firebase if enabled
	if AppSettings.FirebaseEnabled {
		if err := common_utils.InitializeFirebaseApp("configs/local_firebase_configs"); err != nil {
			log.Fatalf("Failed to initialize Firebase: %v", err)
		}
		fmt.Println("Firebase connected...")
	}

	port := os.Getenv("PORT")

	// Swagger documentation setup
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// HealthCheck route
	r.GET("/api/health-check", v1_routes.HealthCheck)

	// Create a router group for v1 routes
	v1Routes := r.Group("/api/v1/")
	v1_routes.SetupV1Routes(v1Routes, database) // Pass the database if it's enabled

	// Attach Logger and Recovery middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	currentTime := time.Now().Format(time.RFC3339)
	fmt.Println("                                                      ")
	log.Println("To shut down MariaGin, press <CTRL> + C at any time.")
	log.Println("Read more at https://github.com/lushiv/MariaGin.")
	log.Println("                 MariaGin                                 ")
	log.Println("---------------   v1.0.1   -----------------------------")
	log.Println("Time ::::::> ", currentTime)
	log.Println("Server is running on ::::::> ", port)
	log.Println("Final API docs are running on ::::::> http://localhost:3000/docs/index.html#")
	r.Run(":" + port)
}
