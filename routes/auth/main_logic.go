package auth

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
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

	// Create a UserInsertData struct from the request.
	userData := TblUsers{
		UUID:       uuid.New().String(),
		Email:      req.Email,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Password:   req.Password,
		ProfilePic: req.ProfilePic,
		Gender:     req.Gender,
	}

	// Insert the user into the database.
	if err := InsertUserIntoDB(userData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// @Summary Log in a customer
// @Description Log in a customer and generate a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func LoginCustomer(c *gin.Context) {
	// Implement customer login logic here
	// Verify customer credentials, generate and return a JWT token upon successful login
	token := "your_generated_jwt_token"
	c.JSON(http.StatusOK, LoginResponse{"Customer logged in successfully", token})
}

// @Summary Log out a customer
// @Description Log out a customer
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} RegisterResponse
// @Router /auth/logout [post]
func LogoutCustomer(c *gin.Context) {
	// Implement customer logout logic here (optional)
	c.JSON(http.StatusOK, RegisterResponse{"Customer logged out successfully"})
}
