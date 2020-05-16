package list

import (
	"gf-app/library"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	library.Fetch(r, "list/index")
}
