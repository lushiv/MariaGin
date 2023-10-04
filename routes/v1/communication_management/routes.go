package communication_management

import (
	communication_management_methods "go-gin-api-boilerplate/routes/v1/communication_management/methods"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/send-test-email", communication_management_methods.SendEmailTest)
	router.POST("/publish-message-test", communication_management_methods.PublishMessageTest)
	router.GET("/consume-message-test", communication_management_methods.ConsumeMessageTest)
}
