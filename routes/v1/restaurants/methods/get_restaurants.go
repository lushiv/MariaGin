package restaurants

import (
	restaurants_utils "go-gin-api-boilerplate/routes/v1/restaurants/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get a list of restaurants
// @Description Get a list of all restaurants
// @Tags CRUD Examples API
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} GetRestaurantsResponse
// @Router /api/v1/restaurants [get]
func GetRestaurants(c *gin.Context) {
	// Fetch and return a list of restaurants from the database.
	restaurants, err := restaurants_utils.FetchRestaurantsFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}
