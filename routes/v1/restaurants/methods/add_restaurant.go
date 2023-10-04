package restaurants

import (
	restaurants_schemas "go-gin-api-boilerplate/routes/v1/restaurants/schemas"
	restaurants_utils "go-gin-api-boilerplate/routes/v1/restaurants/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new restaurant
// @Description Add a new restaurant to the database
// @Tags CRUD Examples API
// @Accept json
// @Produce json
// @Param restaurant body AddRestaurantRequest true "AddRestaurantRequest object to add"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/restaurants [post]
func AddRestaurant(c *gin.Context) {
	// Bind the request body to a Restaurant struct.
	var restaurant restaurants_schemas.AddRestaurantRequest
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the restaurant into the database.
	if err := restaurants_utils.InsertRestaurantIntoDB(restaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add restaurant"})
		return
	}
	c.JSON(http.StatusCreated, restaurant)
}
