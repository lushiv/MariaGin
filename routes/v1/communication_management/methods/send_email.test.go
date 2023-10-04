package communication_management

import (
	"fmt"
	communication_management_schemas "go-gin-api-boilerplate/routes/v1/communication_management/schemas"
	common_utils "go-gin-api-boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Send a test email
// @Tags Communication Management API TEST
// @Description Send a test email
// @Accept json
// @Produce json
// @Param request body SendEmailTestRequest true "Email data"
// @Success 200 {object} CommonResponse
// @Router /communication/send-test-email [post]
func SendEmailTest(c *gin.Context) {
	var request communication_management_schemas.SendEmailTestRequest

	// Parse email data from the request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Prepare email data
	emailData := &common_utils.EmailData{
		To:      request.SendTo,
		Subject: request.Subject,
		Message: request.Message,
	}
	fmt.Println(emailData)

	// Call the SendGrid helper function to send the email
	err := common_utils.SendEmailUsingSendGrid(emailData)
	if err == nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
