package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Time struct {
	CreatedAt time.Time `json:"createdAt" swaggertype:"string" example:"2019-01-01 00:00:00"` // create time
	UpdatedAt time.Time `json:"updatedAt" swaggertype:"string" example:"2019-01-01 00:00:00"` // update time
}

type Base struct {
	Id uint `json:"id"`
	Time
}

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

func SuccessWithData(c *gin.Context, data interface{}) {
	Result(c, Ok, ErrMsg[Ok], data)
}
