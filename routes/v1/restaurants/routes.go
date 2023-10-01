package restaurants

import (
	"go-gin-api-boilerplate/middleware"
	restaurants_methods "go-gin-api-boilerplate/routes/v1/restaurants/methods"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up restaurant-related API routes.
func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", middleware.AuthMiddleware(), restaurants_methods.GetRestaurants)
	router.POST("", restaurants_methods.AddRestaurant)
	router.PUT("/:id", restaurants_methods.UpdateRestaurant)
	router.DELETE("/:id", restaurants_methods.DeleteRestaurant)
}
