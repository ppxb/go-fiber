package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ppxb/go-fiber/app/models"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/req"
	"github.com/ppxb/go-fiber/pkg/response"
	"github.com/ppxb/go-fiber/pkg/service"
	"github.com/ppxb/go-fiber/pkg/utils"
	"github.com/xuri/excelize/v2"
	"reflect"
)

func ImportExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.WithContext(c).WithError(errors.Errorf("%v", err)).Error("upload file fail")
	}

	f, err := excelize.OpenReader(file)
	if err != nil {
		log.WithContext(c).WithError(errors.Errorf("%v", err)).Error("open file fail")
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var req req.CreateAssetDto
	var asset models.SysAsset
	var assets []models.SysAsset
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
		_ = c.ShouldBindJSON(&req)
		utils.Struct2StructByJson(req, &asset)
		assets = append(assets, asset)
	}
	db := service.New(c)
	err = db.Q.Db.Create(&assets).Error
	if err != nil {
		err = errors.Errorf(response.AssetImportErrorMsg)
		response.SuccessWithMsg(c, response.AssetImportErrorMsg)
	} else {
		response.SuccessWithMsg(c, "数据导入成功")
	}
}

func DownloadTemplate(c *gin.Context) {
	c.File("./asset/asset_export_template.xlsx")
}
