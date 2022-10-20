package router

import (
	"github.com/ppxb/go-fiber/api/v1"
)

func (r Router) Asset() {
	router := r.ops.group.Group("/asset")
	router.POST("/import", v1.ImportExcel)
	router.GET("/template", v1.DownloadTemplate)
}
