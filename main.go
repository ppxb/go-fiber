package main

import (
	"context"
	"github.com/ppxb/go-fiber/initialize"
	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/listen"
	"github.com/ppxb/go-fiber/pkg/router"
	"runtime"
	"strings"
)

var ctx = context.Background()

// @title fiber eam app API
// @version 1.0
// @description A simple Eam system written by golang.
// @license.name MIT
// @license.url https://github.com/ppxb/go-fiber/blob/master/LICENSE
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	_, file, _, _ := runtime.Caller(0)
	global.RuntimeRoot = strings.TrimSuffix(file, "main.go")

	initialize.Config(ctx)
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
