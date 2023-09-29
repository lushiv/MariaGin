// Step 1: Create an authentication middleware in middleware/auth_middleware.go

package middleware

import (
	"fmt"
	"net/http"

	common_utils "go-gin-api-boilerplate/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the request header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Verify the JWT token
		token, err := common_utils.VerifyToken(tokenString)
		fmt.Println("token:::", token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract customer information from the token and set it in the context
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println("claims:::", claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		fmt.Println("kskk")

		// Set "user" and "role" claims in the context
		c.Set("user", token["user"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
