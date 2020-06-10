package router

import (
	"gf-app/app/api/index"
	"gf-app/app/api/list"
	"gf-app/app/api/login"
	"gf-app/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/login", new(login.Controller))
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth)
			group.ALL("/index", new(index.Controller))
			group.ALL("/list", new(list.Controller))
		})
	})
}
