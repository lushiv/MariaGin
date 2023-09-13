// handlers.go

package restaurants

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// @Summary Get a list of restaurants
// @Description Get a list of all restaurants
// @Tags CRUD API Sample: Restaurant
// @Accept json
// @Produce json
// @Success 200 {array} Restaurant
// @Router /restaurants [get]
func GetRestaurants(c *gin.Context) {
	// Fetch and return a list of restaurants from the database.
	restaurants, err := FetchRestaurantsFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

// @Summary Add a new restaurant
// @Description Add a new restaurant to the database
// @Tags CRUD API Sample: Restaurant
// @Accept json
// @Produce json
// @Param restaurant body Restaurant true "Restaurant object to add"
// @Success 201 {object} Restaurant
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /restaurants [post]
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

// @Summary Update a restaurant by ID
// @Description Update an existing restaurant in the database by its ID
// @Tags CRUD API Sample: Restaurant
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID to update"
// @Param restaurant body Restaurant true "Updated restaurant object"
// @Success 200 {object} Restaurant
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /restaurants/{id} [put]
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

// @Summary Delete a restaurant by ID
// @Description Delete a restaurant from the database by its ID
// @Tags CRUD API Sample: Restaurant
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID to delete"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /restaurants/{id} [delete]
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
