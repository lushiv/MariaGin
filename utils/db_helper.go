package common_utils

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// Initialize sets the database connection.
func Initialize(database *sql.DB) {
	db = database
}

func IsSessionDeleted(sessionToken string, userID int) bool {
	query := "SELECT COUNT(*) FROM login_session WHERE token = ? AND user_id = ? AND deleted = 1"
	fmt.Println(sessionToken, userID)
	var count int
	err := db.QueryRow(query, sessionToken, userID).Scan(&count)
	if err != nil {
		fmt.Printf("Error fetching records from login_session table: %v", err)
		return false
	}
	return count > 0
}
