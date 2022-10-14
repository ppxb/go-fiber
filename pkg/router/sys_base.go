package router

import "github.com/ppxb/go-fiber/pkg/middleware"

func (r Router) Base() {
	router := r.ops.group.Group("/base")
	router.POST("/login", middleware.JwtLogin(r.ops.jwtOps...))
}
