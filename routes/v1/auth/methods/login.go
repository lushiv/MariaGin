package auth

import (
	"net/http"
	"strconv"

	auth_schemas "go-gin-api-boilerplate/routes/v1/auth/schemas"
	auth_utils "go-gin-api-boilerplate/routes/v1/auth/utils"
	common_utils "go-gin-api-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Log in a customer
// @Description Log in a customer and generate a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param LoginUserReq body LoginUserReq true "Login request"
// @Success 200 {object} LoginUserResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/login [post]
func LoginCustomer(c *gin.Context) {
	var req auth_schemas.LoginUserReq

	// Bind the request body to LoginUserReq struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email exists in the database.
	userID, err := auth_utils.GetUserIDByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Fetch the user's hashed password from the database.
	hashedPassword := auth_utils.GetUserPassword(req.Email)

	// Verify the provided password against the hashed password.
	if err := common_utils.VerifyPassword(req.Password, hashedPassword); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Generate a JWT token for the logged-in customer.
	userInfo := common_utils.UserInfo{
		ID:   userID,
		Role: "user", // You can set the customer's role as needed.
	}
	jwtToken, err := common_utils.GenerateJWTToken(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	// Convert userID from string to int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		c.Abort()
		return
	}

	// Insert the validation token into the database
	if err := auth_utils.InsertLoginSession(userIDInt, jwtToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert login session"})
		return
	}

	// Set data in Redis
	// Set data in Redis
	err = common_utils.SetKey("loginSession", jwtToken)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set data in Redis"})
		return
	}

	// Return the JWT token in the response.
	response := auth_schemas.LoginUserResponse{
		Message: "Customer logged in successfully",
		Token:   jwtToken,
	}

	c.JSON(http.StatusOK, response)
}
