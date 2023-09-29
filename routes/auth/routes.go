package auth

import (
	"go-gin-api-boilerplate/middleware"

	"github.com/gin-gonic/gin"
)

// API routes.
func RegisterRoutes(router *gin.RouterGroup) {

	router.POST("/register", RegisterUser)
	router.POST("/login", LoginCustomer)
	router.POST("/logout", middleware.AuthMiddleware(), LogoutCustomer)
}
