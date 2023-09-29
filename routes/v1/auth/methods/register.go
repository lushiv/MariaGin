package auth

import (
	"net/http"

	auth_schemas "go-gin-api-boilerplate/routes/v1/auth/schemas"
	auth_utils "go-gin-api-boilerplate/routes/v1/auth/utils"
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
	var req auth_schemas.RegisterUserReq

	// Bind the request body to RegisterUserReq struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is valid.
	if !auth_utils.ValidateEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Check if the email already exists in the database.
	if auth_utils.EmailExists(req.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the phone already exists in the database.
	if auth_utils.PhoneNumberExists(req.Email) {
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
	userData := auth_schemas.TblUsers{
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
	if err := auth_utils.InsertUserIntoDB(userData); err != nil {
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
	if err := auth_utils.InsertValidateToken(1, jwtToken, otpCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert validation token"})
		return
	}
	// Return the response including the JWT token
	response := auth_schemas.RegisterResponse{
		Message: "User registered successfully",
		Token:   jwtToken,
	}

	c.JSON(http.StatusCreated, response)

}
