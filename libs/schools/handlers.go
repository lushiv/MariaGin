// handlers.go

package restaurants

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRestaurants handles the GET /restaurants route.
func GetRestaurants(c *gin.Context) {
	// Fetch and return a list of restaurants from the database.
	restaurants, err := FetchRestaurantsFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

// AddRestaurant handles the POST /restaurants route.
func AddRestaurant(c *gin.Context) {
	fmt.Println("Calling API :: AddRestaurant")
	// Bind the request body to a Restaurant struct.
	var restaurant Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the restaurant into the database.
	if err := InsertRestaurantIntoDB(restaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add restaurant"})
		return
	}
	c.JSON(http.StatusCreated, restaurant)
}

// UpdateRestaurant handles the PUT /restaurants/:id route.
func UpdateRestaurant(c *gin.Context) {
	// Extract the restaurant ID from the URL parameter.
	restaurantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
		return
	}

	// Bind the request body to an updated Restaurant struct.
	var updatedRestaurant Restaurant
	if err := c.ShouldBindJSON(&updatedRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the restaurant with the given ID in the database.
	if err := UpdateRestaurantInDB(restaurantID, updatedRestaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update restaurant"})
		return
	}
	c.JSON(http.StatusOK, updatedRestaurant)
}

// DeleteRestaurant handles the DELETE /restaurants/:id route.
func DeleteRestaurant(c *gin.Context) {
	// Extract the restaurant ID from the URL parameter.
	restaurantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
		return
	}

	// Delete the restaurant with the given ID from the database.
	if err := DeleteRestaurantFromDB(restaurantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete restaurant"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
