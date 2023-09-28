package auth

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

// Restaurant represents a restaurant entity.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Initialize sets the database connection.
func Initialize(database *sql.DB) {
	db = database
}

// API routes.
func RegisterRoutes(router *gin.RouterGroup) {

	router.POST("/register", RegisterUser)
	router.POST("/login", LoginCustomer)
	router.POST("/logout", LogoutCustomer)
}

func InsertUserIntoDB(user User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username, user.Password)
	if err != nil {
		fmt.Printf("Error inserting users: %v", err)
		return err
	}
	return nil
}
