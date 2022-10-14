package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/pkg/middleware"
)

type Options struct {
	group  *gin.RouterGroup
	jwt    bool
	jwtOps []func(options *middleware.JwtOptions)
}

func WithGroup(group *gin.RouterGroup) func(*Options) {
	return func(options *Options) {
		getOptions(options).group = group
	}
}

func WithJwt(flag bool) func(*Options) {
	return func(options *Options) {
		getOptions(options).jwt = flag
	}
}

func withJwtOps(ops ...func(*middleware.JwtOptions)) func(*Options) {
	return func(options *Options) {
		getOptions(options).jwtOps = append(getOptions(options).jwtOps, ops...)
	}
}

func getOptions(options *Options) *Options {
	if options == nil {
		return &Options{
			jwt: true,
		}
	}
	return options
}
