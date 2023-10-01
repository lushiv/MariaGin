package restaurants

import (
	restaurants_utils "go-gin-api-boilerplate/routes/v1/restaurants/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	if err := restaurants_utils.DeleteRestaurantFromDB(restaurantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete restaurant"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
