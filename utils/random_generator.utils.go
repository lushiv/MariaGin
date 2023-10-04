package common_utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateUniqueObjectKey generates a unique object key (filename).
func GenerateUniqueObjectKey() string {
	// Get the current timestamp in milliseconds
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Generate a random string of 6 characters
	randomString := generateRandomString(6)

	// Combine the timestamp and random string to create the object key
	objectKey := fmt.Sprintf("%d_%s", timestamp, randomString)

	return objectKey
}

// generateRandomString generates a random string of a specified length.
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
