package common_utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHashPassword generates a hashed password using bcrypt.
func GenerateHashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword verifies a plain password against a hashed password using bcrypt.
func VerifyPassword(plainPassword string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
}
