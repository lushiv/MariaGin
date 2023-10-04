// routes/v1/routes.go
package v1_routes

import (
	"database/sql" // Import the database/sql package
	"go-gin-api-boilerplate/routes/v1/auth"
	auth_utils_db_con "go-gin-api-boilerplate/routes/v1/auth/utils"
	"go-gin-api-boilerplate/routes/v1/communication_management"
	"go-gin-api-boilerplate/routes/v1/file_uploads"
	restaurants "go-gin-api-boilerplate/routes/v1/restaurants"
	restaurants_utils_db_con "go-gin-api-boilerplate/routes/v1/restaurants/utils"
	common_utils "go-gin-api-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func SetupV1Routes(r *gin.RouterGroup, database *sql.DB) {
	// Create a router group for auth routes
	authRoutes := r.Group("/auth")
	auth_utils_db_con.Initialize(database)
	common_utils.Initialize(database)
	auth.RegisterRoutes(authRoutes)

	// Create a router group for communication management routes
	communicationManagementRoutes := r.Group("/communication")
	communication_management.RegisterRoutes(communicationManagementRoutes)

	// Create a router group for restaurant routes
	restaurantRoutes := r.Group("/restaurants")
	restaurants_utils_db_con.Initialize(database)
	restaurants.RegisterRoutes(restaurantRoutes)

	// Create a router group for communication management routes
	fileUploadsRoutes := r.Group("/files")
	file_uploads.RegisterRoutes(fileUploadsRoutes)
}
