package login

import (
	"gf-app/app/service/login"
	"gf-app/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

type LoginReq struct {
	Name string `p:"name" v:"required|min-length:1"`
	Pwd  string `p:"pwd" v:"required|min-length:1"`
}

func (c *Controller) Index(r *ghttp.Request) {
	userinfo := r.Session.Get("sys_user")
	if userinfo != nil {
		r.Response.RedirectTo("/index")
	}

	library.Display(r, "login/index.html")
}

func (c *Controller) Dologin(r *ghttp.Request) {
	var req *LoginReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := login.CheckLogin(req.Name, req.Pwd)
	if res {
		r.Session.Set("sys_user", req.Name)
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}
