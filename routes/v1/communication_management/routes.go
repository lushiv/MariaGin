package communication_management

import (
	communication_management_methods "go-gin-api-boilerplate/routes/v1/communication_management/methods"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up restaurant-related API routes.
func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/send-test-email", communication_management_methods.SendEmailTest)
}
