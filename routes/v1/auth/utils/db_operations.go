package auth

import (
	"database/sql"
	"fmt"
	auth "go-gin-api-boilerplate/routes/v1/auth/schemas"
	"time"

	"github.com/google/uuid"
)

var db *sql.DB

// Initialize sets the database connection.
func Initialize(database *sql.DB) {
	db = database
}

// InsertUserIntoDB inserts user data into the 'users' table.
func InsertUserIntoDB(user auth.TblUsers) error {
	_, err := db.Exec("INSERT INTO users (uuid, email, first_name, gender, last_name, middle_name, password, phone, profile_pic) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)",
		user.UUID, user.Email, user.FirstName, user.Gender, user.LastName, user.MiddleName, user.Password, user.Phone, user.ProfilePic)
	if err != nil {
		fmt.Printf("Error inserting user into 'users' table: %v", err)
		return err
	}
	return nil
}

// Check if the email already exists in the database.
func EmailExists(email string) bool {
	// Perform a database query to check if the email exists.
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		// Handle the error (e.g., log it or return false).
		return false
	}
	return count > 0
}

// Check if the phone number already exists in the database.
func PhoneNumberExists(phone string) bool {
	// Perform a database query to check if the phone number exists.
	query := "SELECT COUNT(*) FROM users WHERE phone = ?"
	var count int
	err := db.QueryRow(query, phone).Scan(&count)
	if err != nil {
		// Handle the error (e.g., log it or return false).
		return false
	}
	return count > 0
}

// Define a constant for token expiry duration in minutes
const TokenExpiryMinutes = 60

// InsertValidateTokenIntoDB inserts a validation token into the database.
func InsertValidateTokenIntoDB(data auth.TblValidateToken) error {
	// Replace this with your actual database insertion logic.
	// Assume db is a database connection that you have established.

	// Prepare the SQL query
	query := `
		INSERT INTO validate_token (uuid, user_id, token, used, expiry_time)
		VALUES (?, ?, ?, ?, ?)
	`

	// Execute the SQL query with the provided data
	_, err := db.Exec(query,
		data.UUID, data.UserID, data.Token, data.Used, data.ExpiryTime)

	if err != nil {
		fmt.Printf("Error inserting user into 'validate_token' table: %v", err)
		return err
	}

	return nil
}

// InsertValidateToken inserts a validation token into the database.
func InsertValidateToken(userID int, token string, otp string) error {
	// Generate a new UUID for the token
	uuid := uuid.New().String()

	// Calculate the token expiry time (you can customize this as needed)
	expiryTime := time.Now().Add(time.Minute * time.Duration(TokenExpiryMinutes))

	// Prepare the data for insertion
	data := auth.TblValidateToken{
		UUID:   uuid,
		UserID: userID,
		Token:  token,
		Used:   0,
		// OTP:      otp, // You can include OTP if needed
		ExpiryTime: expiryTime, // Assign the expiry time directly as time.Time
	}

	// Insert the data into the database
	if err := InsertValidateTokenIntoDB(data); err != nil {
		return err
	}

	return nil
}

// InsertValidateTokenIntoDB inserts a validation token into the database.
func InsertLoginSessionIntoDB(data auth.TblLoginSession) error {
	// Prepare the SQL query
	query := `
		INSERT INTO login_session (uuid, user_id, token, session_expiry_timestamp)
		VALUES (?, ?, ?, ?)
	`
	// Execute the SQL query with the provided data
	_, err := db.Exec(query,
		data.UUID, data.UserID, data.Token, data.SessionExpiryTimestamp)

	if err != nil {
		fmt.Printf("Error inserting user into 'login_session' table: %v", err)
		return err
	}

	return nil
}

// InsertValidateToken inserts a validation token into the database.
func InsertLoginSession(userID int, token string) error {
	// Generate a new UUID for the token
	uuid := uuid.New().String()

	// Calculate the token expiry time (you can customize this as needed)
	expiryTime := time.Now().Add(time.Minute * time.Duration(TokenExpiryMinutes))

	// Prepare the data for insertion
	data := auth.TblLoginSession{
		UUID:                   uuid,
		UserID:                 userID,
		Token:                  token,
		SessionExpiryTimestamp: expiryTime,
	}

	// Insert the data into the database
	if err := InsertLoginSessionIntoDB(data); err != nil {
		return err
	}

	return nil
}

// Check if a user with the given email exists in the database.
func CheckUser(email string) bool {
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// Fetch the user's ID by email from the database.
func GetUserIDByEmail(email string) (string, error) {
	var userID string
	query := "SELECT id FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		// Handle the error (e.g., log it or return an empty string).
		return "", err
	}
	return userID, nil
}

// Fetch the user's hashed password from the database.
func GetUserPassword(email string) string {
	var hashedPassword string
	query := "SELECT password FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		// Handle the error (e.g., log it or return an empty string).
		return ""
	}
	return hashedPassword
}

// InvalidateSession updates the session to mark it as deleted
func InvalidateSession(sessionToken string) error {
	query := "UPDATE login_session SET deleted = 1 WHERE token = ?"
	_, err := db.Exec(query, sessionToken)
	return err
}

func IsSessionDeleted(sessionToken string, userID int) bool {
	query := "SELECT COUNT(*) FROM login_session WHERE token = ? AND user_id = ? AND deleted = 1"
	fmt.Println(sessionToken, userID)
	var count int
	err := db.QueryRow(query, sessionToken, userID).Scan(&count)
	if err != nil {
		// Handle the error (e.g., log it or return false).
		return false
	}
	return count > 0
}
