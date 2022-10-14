package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/app/api"
	"github.com/ppxb/go-fiber/app/models"
	swagger "github.com/ppxb/go-fiber/docs"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/middleware"
	"github.com/ppxb/go-fiber/pkg/req"
	"github.com/ppxb/go-fiber/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	ops Options
}

func Register(ctx context.Context) *gin.Engine {
	r := gin.New()

	r.Use(
		middleware.Cors(),
	)

	apiGroup := r.Group(global.Conf.Server.UrlPrefix)
	apiGroup.GET("/health", api.Health)

	jwtOps := []func(*middleware.JwtOptions){
		middleware.WithJwtRealm(global.Conf.Jwt.Realm),
		middleware.WithJwtKey(global.Conf.Jwt.Key),
		middleware.WithJwtTimeout(global.Conf.Jwt.Timeout),
		middleware.WithJwtMaxRefresh(global.Conf.Jwt.MaxRefresh),
		middleware.WithJwtLoginPwdCheck(func(c *gin.Context, r req.LoginCheck) (models.SysUser, error) {
			db := service.New(c)
			user, err := db.LoginCheck(r)
			return user, err
		}),
	}

	v1Group := apiGroup.Group(global.Conf.Server.ApiVersion)

	swagger.SwaggerInfo.Version = global.Conf.Server.ApiVersion
	swagger.SwaggerInfo.BasePath = v1Group.BasePath()
	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DocExpansion("none"),
		),
	)

	nr := NewRouter(
		WithGroup(v1Group),
		WithJwt(true),
		withJwtOps(jwtOps...),
	)

	nr.Base()

	return r
}

func NewRouter(options ...func(*Options)) *Router {
	ops := getOptions(nil)
	for _, f := range options {
		f(ops)
	}

	if ops.group == nil {
		panic("router group is empty")
	}

	r := &Router{
		ops: *ops,
	}
	return r
}
