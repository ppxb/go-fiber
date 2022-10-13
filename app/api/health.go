package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/pkg/resp"
)

// Health
// @Tags Server
// @Description Get server health status
// @Router /health [GET]
func Health(c *gin.Context) {
	resp.SuccessWithData(c, map[string]string{
		"STATUS": "UP",
	})
}
