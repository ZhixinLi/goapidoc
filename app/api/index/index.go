package index

import (
	"gf-app/app/service/list"
	"gf-app/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

type ResetpwdReq struct {
	Oldpwd   string `p:"old_pwd" v:"required|min-length:1"`
	Newpwd   string `p:"new_pwd" v:"required|min-length:1|same:check_pwd"`
	Checkpwd string `p:"check_pwd" v:"required|min-length:1"`
}

func (c *Controller) Index(r *ghttp.Request) {
	library.Fetch(r, "index/index")
}

func (c *Controller) Logout(r *ghttp.Request) {
	r.Session.Remove("sys_user")
}

func (c *Controller) Resetpwd(r *ghttp.Request) {
	var req *ResetpwdReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.UpdateWhitelist(r.Session.Get("sys_user").(string), req.Oldpwd, req.Newpwd)
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}
