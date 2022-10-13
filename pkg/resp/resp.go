package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Resp struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Resp{
		Code:      code,
		Data:      data,
		Msg:       msg,
		Timestamp: time.Now().Unix(),
	})
}

func Success(c *gin.Context) {
	Result(c, 200, "success", map[string]interface{}{})
}

func SuccessWithData(c *gin.Context, data interface{}) {
	Result(c, 200, "success", data)
}

func SuccessWithMsg(c *gin.Context, msg string) {
	Result(c, 200, msg, map[string]interface{}{})
}

func Fail(c *gin.Context, code int) {
	Result(c, code, "fail", map[string]interface{}{})
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(c, 500, msg, map[string]interface{}{})
}
