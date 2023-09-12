package common_utils

import (
	"strings"
)

// ValidateEmail validates an email address.
func ValidateEmail(email string) bool {
	// Implement your email validation logic here.
	// Return true if the email is valid, otherwise return false.
	// You can use regular expressions or other validation methods.
	// For simplicity, let's assume all emails are valid in this example.
	return true
}

// InsertValidation validates the request body for insertion.
func InsertValidation(reqBody map[string]interface{}) (bool, string) {
	title, titleExists := reqBody["title"]
	content, contentExists := reqBody["content"]

	if !titleExists || title == nil || strings.TrimSpace(title.(string)) == "" {
		return false, "title is required."
	}

	if !contentExists || content == nil || strings.TrimSpace(content.(string)) == "" {
		return false, "content is required."
	}

	return true, ""
}
