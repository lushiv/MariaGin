package auth

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// Initialize sets the database connection.
func Initialize(database *sql.DB) {
	db = database
}

// InsertUserIntoDB inserts user data into the 'users' table.
func InsertUserIntoDB(user TblUsers) error {
	_, err := db.Exec("INSERT INTO users (uuid, email, first_name, gender, last_name, middle_name, password, phone, profile_pic) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)",
		user.UUID, user.Email, user.FirstName, user.Gender, user.LastName, user.MiddleName, user.Password, user.Phone, user.ProfilePic)
	if err != nil {
		fmt.Printf("Error inserting user into 'users' table: %v", err)
		return err
	}
	return nil
}

// Check if the email already exists in the database.
func emailExists(email string) bool {
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
func phoneNumberExists(phone string) bool {
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
