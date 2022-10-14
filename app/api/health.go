package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/pkg/response"
)

// Health actuator api
func Health(c *gin.Context) {
	response.SuccessWithData(c, map[string]string{
		"status": "UP",
	})
}
