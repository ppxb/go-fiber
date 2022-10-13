package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
	"strings"
)

func Cors(options ...func(*CorsOptions)) gin.HandlerFunc {
	ops := getCorsOptions(nil)
	for _, f := range options {
		f(ops)
	}

	return func(c *gin.Context) {
		// wait to complete tracer
		method := c.Request.Method
		methods := strings.Split(ops.method, ",")
		if !utils.Contains(methods, method) {
			c.Status(http.StatusMethodNotAllowed)
			c.Abort()
			return
		}
		c.Header("Access-Control-Allow-Origin", ops.origin)
		c.Header("Access-Control-Allow-Headers", ops.header)
		c.Header("Access-Control-Allow-Methods", ops.method)
		c.Header("Access-Control-Expose-Headers", ops.expose)
		c.Header("Access-Control-Allow-Credentials", ops.credential)

		if method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
