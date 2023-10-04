package communication_management

import (
	common_utils "go-gin-api-boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Publish a message to RabbitMQ
// @Tags Communication Management API TEST
// @Description Publish a message to RabbitMQ queue
// @Accept json
// @Produce json
// @Param request body PublishRequest true "Message data"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/communication/publish-message-test [post]
func PublishMessageTest(c *gin.Context) {
	// Parse the request body to get the message
	var PublishRequest struct {
		Message string `json:"message"`
	}

	if err := c.BindJSON(&PublishRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Publish the message to RabbitMQ queue
	queueName := "TEST_JANAK" // Replace with your queue name
	err := common_utils.PublishToQueue(queueName, []byte(PublishRequest.Message))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message published successfully"})
}
