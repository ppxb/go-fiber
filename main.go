package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/ppxb/go-fiber/pkg/req"
	"github.com/ppxb/go-fiber/pkg/utils"
	"github.com/xuri/excelize/v2"
	"reflect"
)

var ctx = context.Background()

//go:embed conf
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

	//defer func() {
	//	if err := recover(); err != nil {
	//		log.WithContext(ctx).WithError(errors.Errorf("%v", err)).Error("server run failed, stack: %s", string(debug.Stack()))
	//	}
	//}()
	//
	//_, file, _, _ := runtime.Caller(0)
	//global.RuntimeRoot = strings.TrimSuffix(file, "main.go")
	//
	//initialize.Config(ctx, conf)
	//initialize.Mysql(ctx)
	//
	//listen.Http(
	//	listen.WithHttpCtx(ctx),
	//	listen.WithHttpPort(global.Conf.Server.Port),
	//	listen.WithHttpHandler(router.Register(ctx)),
	//	listen.WithHttpExit(func() {
	//		// pass
	//	}),
	//)
	f, err := excelize.OpenFile("./asset/资产清单信息批量导入模板.xlsx")
	if err != nil {
		fmt.Println(nil)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var req req.CreateAssetDto
	rows, err := f.GetRows("Sheet1")
	for i, r := range rows {
		if i < 3 {
			continue
		}
		v := reflect.ValueOf(&req).Elem()
		for i, value := range r {
			switch v.Field(i).Type().String() {
			case "float64":
				v.Field(i).SetFloat(utils.Str2Float64(value))
			case "int":
				v.Field(i).SetInt(int64(utils.Str2Int(value)))
			case "time.Time":
				v.Field(i).Set(reflect.ValueOf(utils.Str2Time(value)))
			default:
				v.Field(i).SetString(value)
			}
		}
		fmt.Printf("%+v\n", req)
	}
}
