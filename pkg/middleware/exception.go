package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/response"
	"net/http"
	"runtime/debug"
	"time"
)

func Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			e := errors.Errorf("%v", err)
			log.WithContext(c).WithError(e).Error("runtime exception, stack: %s", string(debug.Stack()))
			rp := response.Resp{
				Code:      response.InternalServerError,
				Data:      map[string]interface{}{},
				Msg:       response.ErrMsg[response.InternalServerError],
				Timestamp: time.Now().Unix(),
			}
			c.JSON(http.StatusOK, rp)
			c.Abort()
			return
		}
	}()
	c.Next()
}
