package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/pkg/response"
)

// Health
// @Tags Server
// @Description Get server health status
// @Router /health [GET]
func Health(c *gin.Context) {
	response.SuccessWithData(c, map[string]string{
		"status": "UP",
	})
}
