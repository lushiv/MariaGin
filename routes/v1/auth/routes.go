package auth

import (
	"go-gin-api-boilerplate/middleware"
	auth_methods "go-gin-api-boilerplate/routes/v1/auth/methods"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", auth_methods.RegisterUser)
	router.POST("/login", auth_methods.LoginCustomer)
	router.POST("/logout", middleware.AuthMiddleware(), auth_methods.LogoutCustomer)
}
