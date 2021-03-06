package list

import (
	"gf-app/app/model/api_detail"
	"gf-app/app/service/list"
	"gf-app/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

type UpdateReq struct {
	Id          int    `p:"id" v:"required|min:1"`
	Uri         string `p:"uri" v:"required|min-length:1"`
	Name        string `p:"name" v:"required|min-length:1"`
	RequestType string `p:"request_type" v:"required|min-length:1"`
	Param       string `p:"param" v:"required|json"`
	ReturnValue string `p:"return_value" v:"required|json"`
}

type DeleteReq struct {
	Id int `p:"id" v:"required|min:1"`
}

type AddgroupReq struct {
	Pid  int    `p:"pid" v:"required|min:1"`
	Name string `p:"name" v:"required|min-length:1"`
}

type AddprojectReq struct {
	Name string `p:"name" v:"required|min-length:1"`
}

type UpdateprojectReq struct {
	Id   int    `p:"id" v:"required|min:1"`
	Name string `p:"name" v:"required|min-length:1"`
}

type UpdategroupReq struct {
	Id   int    `p:"id" v:"required|min:1"`
	Name string `p:"name" v:"required|min-length:1"`
}

func (c *Controller) Index(r *ghttp.Request) {
	pid := r.Get("pid")
	gid := r.Get("gid")
	aid := r.Get("aid")

	library.Fetch(r, "list/index", list.GetAssignData(pid, gid, aid))
}

func (c *Controller) Add(r *ghttp.Request) {
	var req *api_detail.Entity
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.AddDetail(req, r.Session.Get("sys_user").(string))
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}

func (c *Controller) Update(r *ghttp.Request) {
	var req *UpdateReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.UpdateDetial(req.Id, req.Uri, req.Name, req.RequestType, req.Param, req.ReturnValue)
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}

func (c *Controller) Delete(r *ghttp.Request) {
	var req *DeleteReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.DeleteDetail(req.Id)
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}

func (c *Controller) Addgroup(r *ghttp.Request) {
	var req *AddgroupReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.AddGroup(req.Pid, req.Name)
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}

func (c *Controller) Addproject(r *ghttp.Request) {
	var req *AddprojectReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.AddProject(req.Name)
	if res != 0 {
		library.SendJson(r, g.Map{"status": 1, "id": res})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}

func (c *Controller) Updateproject(r *ghttp.Request) {
	var req *UpdateprojectReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.UpdateProject(req.Id, req.Name)
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}

func (c *Controller) Updategroup(r *ghttp.Request) {
	var req *UpdategroupReq
	if err := r.Parse(&req); err != nil {
		library.SendJson(r, g.Map{"status": -1})
	}

	res := list.UpdateGroup(req.Id, req.Name)
	if res {
		library.SendJson(r, g.Map{"status": 1})
	} else {
		library.SendJson(r, g.Map{"status": 0})
	}
}
