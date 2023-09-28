package common_utils

import (
	"time"

	"github.com/pquerna/otp/totp"
)

// GenerateOTP generates a Time-Based One-Time Password (TOTP).
func GenerateOTP(secret string) (string, error) {
	secretBytes := []byte(secret)
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "YourAppName",
		AccountName: "user@example.com",
		Secret:      secretBytes,
		// You can customize the TOTP options as needed.
	})
	if err != nil {
		return "", err
	}

	otpCode, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return "", err
	}

	return otpCode, nil
}
