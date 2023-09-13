package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// @Summary Register a new customer
// @Description Register a new customer and generate a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} RegisterResponse
// @Router /auth/register [post]
func RegisterCustomer(c *gin.Context) {
	// Implement customer registration logic here
	// Use your crypto helper to hash the password
	// Save customer data in the database
	// Generate and return a JWT token upon successful registration
	c.JSON(http.StatusOK, RegisterResponse{"Customer registered successfully"})
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
