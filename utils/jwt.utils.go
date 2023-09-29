package common_utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWTToken generates a JWT token.
func GenerateJWTToken(userInfo UserInfo) (string, error) {
	// Get the token secret from environment variables
	tokenSecret := os.Getenv("TOKEN_SECRET")
	fmt.Println(tokenSecret)

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = userInfo.ID // Set the user claim to the user ID
	claims["role"] = userInfo.Role
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix() // Token expiration time

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("Failed to sign token: %v", err)
	}

	return tokenString, nil
}

// VerifyToken verifies a JWT token.
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Get the token secret from environment variables
	tokenSecret := os.Getenv("TOKEN_SECRET")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and secret key
		if token.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}

// DecodeJWTToken decodes a JWT token without verifying it.
func DecodeJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}

// UserInfo represents user information.
type UserInfo struct {
	ID   string
	Role string
}
