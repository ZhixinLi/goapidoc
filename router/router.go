package router

import (
	"gf-app/app/api/index"
	"gf-app/app/api/list"
	"gf-app/app/service/middleware"
	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()

	s.BindMiddleware("/*", middleware.Auth)

	s.BindObject("/index", new(index.Controller))
	s.BindObject("/list", new(list.Controller))
}
