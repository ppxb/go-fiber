package router

import v1 "github.com/ppxb/go-fiber/app/api/v1"

func (r Router) Upload() {
	router := r.ops.group.Group("/asset")
	router.POST("/import", v1.ImportExcel)
}
