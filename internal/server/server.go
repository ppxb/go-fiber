package server

import (
	"context"
	"github.com/ppxb/go-fiber/config"
	"github.com/ppxb/go-fiber/pkg/logger"
	"github.com/ppxb/go-fiber/pkg/router"
	"os"
	"strings"
)

type options struct {
	Mode    string
	Version string
}

type Option func(*options)

func SetMode(s string) Option {
	return func(o *options) {
		o.Mode = s
	}
}

func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func Run(ctx context.Context, opts ...Option) error {
	var ops options
	for _, f := range opts {
		f(&ops)
	}

	logger.WithContext(ctx).Infof("[Server] Start server with Ver(%s.%s) on PID %d ...", ops.Version, strings.ToUpper(ops.Mode), os.Getpid())

	config.Init(ctx, ops.Mode)

	Http(
		SetHttpCtx(ctx),
		SetHttpPort(config.Conf.Server.Port),
		SetHttpHandler(router.Register(ctx)),
	)

	return nil
}
