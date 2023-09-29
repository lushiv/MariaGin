package auth

import (
	"fmt"
	"net/http"
	"strconv"

	common_utils "go-gin-api-boilerplate/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Register a new customer
// @Description Register a new customer and generate a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param registerRequest body RegisterUserReq true "Registration request"
// @Success 200 {object} RegisterResponse
// @Router /auth/register [post]
// RegisterUser handles user registration.
func RegisterUser(c *gin.Context) {
	var req RegisterUserReq

	// Bind the request body to RegisterUserReq struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is valid.
	if !ValidateEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Check if the email already exists in the database.
	if emailExists(req.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the phone already exists in the database.
	if phoneNumberExists(req.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Phone number already exists"})
		return
	}

	// Generate an OTP (replace 'yourSecretOTP' with your actual OTP secret).
	hashedPassword, err := common_utils.GenerateHashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	// Create a UserInsertData struct from the request.
	userData := TblUsers{
		UUID:       uuid.New().String(),
		Email:      req.Email,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Password:   hashedPassword,
		ProfilePic: req.ProfilePic,
		Gender:     req.Gender,
	}

	// Insert the user into the database.
	if err := InsertUserIntoDB(userData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Generate a JWT token for the registered user.
	userInfo := common_utils.UserInfo{
		ID:   userData.UUID,
		Role: "user", // You can set the user's role as needed.
	}
	jwtToken, err := common_utils.GenerateJWTToken(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	// Generate an OTP (replace 'yourSecretOTP' with your actual OTP secret).
	otpCode, err := common_utils.GenerateOTP("JBSWY3DPEHPK3PXP")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	// Insert the validation token into the database
	if err := InsertValidateToken(1, jwtToken, otpCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert validation token"})
		return
	}
	// Return the response including the JWT token
	response := RegisterResponse{
		Message: "User registered successfully",
		Token:   jwtToken,
	}

	c.JSON(http.StatusCreated, response)

}

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
	var req LoginUserReq

	// Bind the request body to LoginUserReq struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email exists in the database.
	userID, err := getUserIDByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Fetch the user's hashed password from the database.
	hashedPassword := getUserPassword(req.Email)

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
	if err := InsertLoginSession(userIDInt, jwtToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert login session"})
		return
	}

	// Return the JWT token in the response.
	response := LoginUserResponse{
		Message: "Customer logged in successfully",
		Token:   jwtToken,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Log out a customer
// @Description Log out a customer
// @Tags Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} commonResponse
// @Router /auth/logout [post]
func LogoutCustomer(c *gin.Context) {
	fmt.Println("LogoutCustomer")
	// Use the AuthMiddleware to verify authentication
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract user information from the userID
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user information"})
		return
	}

	fmt.Println(userIDInt)
	// Get the session token from the JWT token
	sessionToken := c.GetHeader("Authorization")

	// Invalidate the session by marking it as deleted in the database
	err := invalidateSession(sessionToken, userIDInt)
	if err != false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Customer logged out successfully"})
}
