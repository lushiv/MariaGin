package auth

import (
	"net/mail"
	"strings"
)

// ValidateEmail checks if the provided email is a valid email address.
func ValidateEmail(email string) bool {
	// Parse the email address.
	_, err := mail.ParseAddress(email)

	// Check if there was no error in parsing and the email address contains "@".
	return err == nil && strings.Contains(email, "@")
}
