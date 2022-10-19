package main

import (
	"context"
	"github.com/ppxb/go-fiber/initialize"
	"github.com/ppxb/go-fiber/pkg/global"
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

	//server.Http(
	//	listen.WithHttpCtx(ctx),
	//	listen.WithHttpPort(global.Conf.Server.Port),
	//	server.WithHttpHandler(router.Register(ctx)),
	//	server.WithHttpExit(func() {
	//		// pass
	//	}),
	//)
}
