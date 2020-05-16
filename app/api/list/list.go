package list

import (
	"gf-app/app/model/api_detail"
	"gf-app/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	library.Fetch(r, "list/index")
}

func (c *Controller) Data(r *ghttp.Request) {
	gid := r.Get("gid")

	var send []g.Map
	if gid == nil {
		r.Response.WriteJson(send)
	}

	data, _ := api_detail.FindAll("gid", gid)

	for _, v := range data {
		send = append(send, g.Map{
			"id":     v.Id,
			"name":   v.Name,
			"uri":    v.Uri,
			"author": v.Author,
			"time":   v.Time,
		})
	}
	r.Response.WriteJson(send)
}
