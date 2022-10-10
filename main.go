package main

import (
	"context"
	"embed"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/initialize"
	"runtime"
	"strings"

	"github.com/ppxb/go-fiber/pkg/global"
	"github.com/ppxb/go-fiber/pkg/log"
	"runtime/debug"
)

var ctx = context.Background()

var conf embed.FS

// @title evetion-eam-app API
// @version 1.0
// @description A simple Eam system written by golang.
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).WithError(errors.Errorf("%v", err)).Error("server run failed, stack: %s", global.ProName, string(debug.Stack()))
		}
	}()

	_, file, _, _ := runtime.Caller(0)
	global.RuntimeRoot = strings.TrimSuffix(file, "main.go")

	initialize.Config(ctx, conf)
	initialize.Mysql()
}
