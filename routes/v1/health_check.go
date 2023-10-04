// healthcheck.go
package v1_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary HealthCheck
// @Description get the status of server.
// @Tags Health Check
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /abcd/health-check [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}
