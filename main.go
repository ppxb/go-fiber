package main

import (
	"context"
	"embed"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/initialize"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/listen"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/router"
	"runtime"
	"runtime/debug"
	"strings"
)

var ctx = context.Background()

//go:embed configs
var conf embed.FS

// @title fiber eam app API
// @version 1.0
// @description A simple Eam system written by golang.
// @license.name MIT
// @license.url https://github.com/ppxb/go-fiber/blob/master/LICENSE
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).WithError(errors.Errorf("%v", err)).Error("server run failed, stack: %s", string(debug.Stack()))
		}
	}()

	_, file, _, _ := runtime.Caller(0)
	global.RuntimeRoot = strings.TrimSuffix(file, "main.go")

	initialize.Config(ctx, conf)
	initialize.Mysql(ctx)

	listen.Http(
		listen.WithHttpCtx(ctx),
		listen.WithHttpPort(global.Conf.Server.Port),
		listen.WithHttpHandler(router.Register(ctx)),
		listen.WithHttpExit(func() {
			// pass
		}),
	)
}
