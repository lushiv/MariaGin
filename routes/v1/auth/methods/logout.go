package auth

import (
	"fmt"
	auth_utils "go-gin-api-boilerplate/routes/v1/auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Log out a customer
// @Description Log out a customer
// @Tags Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} CommonResponse
// @Router /api/v1/auth/logout [post]
func LogoutCustomer(c *gin.Context) {
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	fmt.Println(userID)
	// Get the session token from the JWT token
	sessionToken := c.GetHeader("Authorization")

	// Invalidate the session by marking it as deleted in the database
	err := auth_utils.InvalidateSession(sessionToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Customer logged out successfully"})
}
