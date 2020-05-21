package library

import (
	"encoding/json"
	"gf-app/app/model/api_detail"
	"gf-app/app/model/api_group"
	"gf-app/app/model/project"
	"gf-app/app/model/whitelist"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Fetch(r *ghttp.Request, uri string, param ...g.Map) {
	var params = make(g.Map)
	params["headTitle"] = g.Cfg().GetString("main.name")
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
			groupArr = append(groupArr, g.Map{"id": v.Id, "name": v.Name, "count": tmpCount, "pid": v.Pid})
		}
		viewdata["groupMenu"] = groupArr //Project api group list
	}

	for _, v := range param {
		for key, val := range v {
			viewdata[key] = val
		}
	}
	params["viewdata"] = viewdata //View data

	//show view page
	r.Response.WriteTpl("layout/layout.html", params)
}

func SendJson(r *ghttp.Request, send []g.Map) {
	if len(send) == 0 {
		r.Response.WriteJson("")
	} else {
		r.Response.WriteJson(send)
	}
}

func GetJson(send []g.Map) string {
	if b, err := json.Marshal(send); err != nil {
		return ""
	} else {
		return string(b)
	}
}
