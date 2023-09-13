package auth

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

// Restaurant represents a restaurant entity.
type Restaurant struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
}

// Initialize sets the database connection.
func Initialize(database *sql.DB) {
	db = database
}

// RegisterRoutes sets up restaurant-related API routes.
func RegisterRoutes(router *gin.RouterGroup) {

	router.POST("/register", RegisterCustomer)
	router.POST("/login", LoginCustomer)
	router.POST("/logout", LogoutCustomer)
}

// FetchRestaurantsFromDB retrieves a list of restaurants from the database.
func FetchRestaurantsFromDB() ([]Restaurant, error) {
	// Define a slice to store the retrieved restaurants.
	var restaurants []Restaurant

	// Query the database to fetch restaurants.
	rows, err := db.Query("SELECT id, name, location, rating FROM restaurants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set and populate the restaurants slice.
	for rows.Next() {
		var restaurant Restaurant
		err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Rating)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func InsertRestaurantIntoDB(restaurant Restaurant) error {
	_, err := db.Exec("INSERT INTO restaurants (name, location, rating) VALUES (?, ?, ?)",
		restaurant.Name, restaurant.Location, restaurant.Rating)
	if err != nil {
		fmt.Printf("Error inserting restaurant: %v", err)
		return err
	}
	return nil
}

// UpdateRestaurantInDB updates an existing restaurant in the database.
func UpdateRestaurantInDB(restaurantID int, updatedRestaurant Restaurant) error {
	// Update the restaurant in the database without the 'description' field.
	_, err := db.Exec("UPDATE restaurants SET name=?, location=?, rating=? WHERE id=?",
		updatedRestaurant.Name, updatedRestaurant.Location, updatedRestaurant.Rating, restaurantID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRestaurantFromDB deletes a restaurant with the given ID from the database.
func DeleteRestaurantFromDB(restaurantID int) error {
	// Delete the restaurant from the database.
	_, err := db.Exec("DELETE FROM restaurants WHERE id=?", restaurantID)
	if err != nil {
		return err
	}
	return nil
}
