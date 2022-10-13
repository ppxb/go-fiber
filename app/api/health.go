package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/pkg/resp"
)

func Health(c *gin.Context) {
	resp.SuccessWithData(c, map[string]string{
		"STATUS": "UP",
	})
}
