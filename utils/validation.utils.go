package common_utils

import (
	"strings"
)

// ValidateEmail validates an email address.
func ValidateEmail(email string) bool {
	// Implement your email validation logic here.
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
