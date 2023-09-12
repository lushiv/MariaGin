// db/database.go
package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection and returns it.
func InitDB() (*sql.DB, error) {
	// Construct the database connection string
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Initialize the database connection
	var err error
	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Check database connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
