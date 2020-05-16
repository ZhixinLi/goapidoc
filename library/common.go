package library

import (
	"gf-app/app/model/whitelist"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Fetch(r *ghttp.Request, uri string, param ...g.Map) {
	var params = make(g.Map)
	for _, v := range param {
		for key, val := range v {
			params[key] = val
		}
	}

	//嵌入页面
	params["mainTpl"] = uri + ".html"

	var menu_arr []g.Map
	menu_arr = append(menu_arr, g.Map{"id": 1, "controller": "test", "name": "测试"})

	var viewdata = make(g.Map)
	//项目菜单
	viewdata["bigMenu"] = menu_arr

	user, _ := whitelist.FindOne("ip", r.GetClientIp())
	viewdata["op"] = user.Op
	viewdata["ip"] = user.Ip
	viewdata["username"] = user.Name
	viewdata["indexname"]=g.Cfg().GetString("main.name")

	//视图数据
	params["viewdata"] = viewdata

	//显示页面
	r.Response.WriteTpl("layout/layout.html", params)
}
