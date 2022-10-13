package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/app/api"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/middleware"
)

func Register(ctx context.Context) *gin.Engine {
	r := gin.New()

	r.Use(
		middleware.Cors(),
	)

	apiGroup := r.Group(global.Conf.Server.UrlPrefix)
	apiGroup.GET("/health", api.Health)

	return r
}
