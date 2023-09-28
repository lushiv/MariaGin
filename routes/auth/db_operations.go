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
