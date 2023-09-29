package middleware

import (
	"net/http"
	"strconv"

	common_utils "go-gin-api-boilerplate/utils" // Import your API package

	"github.com/gin-gonic/gin"
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
		claims, err := common_utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
			c.Abort()
			return
		}

		// Access the user claim
		userID, ok := claims["user"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user information"})
			c.Abort()
			return
		}

		// Convert userID from string to int
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		// Check if the session is deleted or expired
		sessionDeleted := common_utils.IsSessionDeleted(tokenString, userIDInt)
		if sessionDeleted {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session is deleted or expired"})
			c.Abort()
			return
		}

		c.Set("user", userID) // Set "user" in the context
		c.Next()
	}
}
