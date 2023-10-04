package communication_management

import (
	common_utils "go-gin-api-boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Consume a message from RabbitMQ
// @Tags Communication Management API TEST
// @Description Consume a message from a RabbitMQ queue
// @Accept json
// @Produce json
// @Success 200 {object} CommonResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/communication/consume-message-test [get]
func ConsumeMessageTest(c *gin.Context) {
	// Specify the queue name to consume from
	queueName := "TEST_JANAK" // Replace with your queue name

	// Consume a message from the queue
	msgs, err := common_utils.ConsumeFromQueue(queueName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to consume message"})
		return
	}

	// Wait for a message from the queue
	msg := <-msgs

	// Process the consumed message
	consumedMessage := string(msg.Body)

	// You can perform additional processing or return the message as needed
	c.JSON(http.StatusOK, gin.H{"message": consumedMessage})
}
