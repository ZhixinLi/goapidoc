package library

import (
	"gf-app/app/model/api_detail"
	"gf-app/app/model/api_group"
	"gf-app/app/model/project"
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

	params["mainTpl"] = uri + ".html" //Input show page

	projects, _ := project.FindAll()
	var menuArr []g.Map
	for _, v := range projects {
		tmpCount, _ := api_detail.FindCount("pid", v.Id)
		menuArr = append(menuArr, g.Map{"id": v.Id, "name": v.Name, "count": tmpCount})
	}

	var viewdata = make(g.Map)
	viewdata["bigMenu"] = menuArr //Project menu

	user, _ := whitelist.FindOne("ip", r.GetClientIp())
	viewdata["op"] = user.Op         //Auth: modify
	viewdata["ip"] = user.Ip         //User's ip
	viewdata["username"] = user.Name //User's username

	viewdata["indexname"] = g.Cfg().GetString("main.name") //Get config value from config.toml

	pid := r.Get("pid") //Get project id form request info
	if pid != nil {
		viewdata["pid"] = pid //Project id
		groups, _ := api_group.FindAll("pid", pid)
		var groupArr []g.Map
		for _, v := range groups {
			tmpCount, _ := api_detail.FindCount("gid", v.Id)
			groupArr = append(groupArr, g.Map{"id": v.Id, "name": v.Name, "count": tmpCount})
		}
		viewdata["groupMenu"] = groupArr //Project api group list
	}

	params["viewdata"] = viewdata //View data

	//显示页面
	r.Response.WriteTpl("layout/layout.html", params)
}