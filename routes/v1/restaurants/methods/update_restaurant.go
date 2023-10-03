package restaurants

import (
	restaurants_schemas "go-gin-api-boilerplate/routes/v1/restaurants/schemas"
	restaurants_utils "go-gin-api-boilerplate/routes/v1/restaurants/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update a restaurant by ID
// @Description Update an existing restaurant in the database by its ID
// @Tags CRUD Examples API
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID to update"
// @Param restaurant body UpdateRestaurantRequest true "Updated restaurant object"
// @Success 200 {object} CommonResponse
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
	var updatedRestaurant restaurants_schemas.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&updatedRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the restaurant with the given ID in the database.
	if err := restaurants_utils.UpdateRestaurantInDB(restaurantID, updatedRestaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update restaurant"})
		return
	}
	c.JSON(http.StatusOK, updatedRestaurant)
}
