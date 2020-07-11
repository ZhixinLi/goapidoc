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
	var groupArr []g.Map
	var menuArr []g.Map
	var params = make(g.Map)
	var viewdata = make(g.Map)

	name := g.Cfg().GetString("main.name") //Get config value from config.toml.

	projects, _ := project.FindAll()
	user, _ := whitelist.FindOne("name", r.Session.Get("sys_user"))

	for _, v := range projects {
		tmpCount, _ := api_detail.FindCount("pid", v.Id)
		menuArr = append(menuArr, g.Map{"id": v.Id, "name": v.Name, "count": tmpCount})
	}

	//Project menu.
	viewdata["sys_bigMenu"] = menuArr
	viewdata["sys_op"] = user.Op         //Auth: modify.
	viewdata["sys_ip"] = user.Ip         //User's ip.
	viewdata["sys_username"] = user.Name //User's username.
	viewdata["sys_indexname"] = name

	//Get project id form request info.
	pid := r.Get("pid")
	if pid != nil {
		//Project id.
		viewdata["sys_pid"] = pid
		groups, _ := api_group.FindAll("pid", pid)

		for _, v := range groups {
			tmpCount, _ := api_detail.FindCount("gid", v.Id)
			groupArr = append(groupArr, g.Map{"id": v.Id, "name": v.Name, "count": tmpCount, "pid": v.Pid})
		}
		//Project api group list
		viewdata["sys_groupMenu"] = groupArr
	}

	for _, v := range param {
		for key, val := range v {
			viewdata[key] = val
		}
	}

	//Total params.
	params["sys_viewData"] = viewdata
	params["sys_headTitle"] = name
	params["sys_mainTpl"] = uri + ".html" //Input show page.

	//show view page.
	r.Response.WriteTpl("layout/layout.html", params)
}

func Display(r *ghttp.Request, uri string) {
	r.Response.WriteTpl(uri)
}

func SendJson(r *ghttp.Request, send g.Map) {
	r.Response.WriteJsonExit(send)
}

//Transform table data to json string.
func GetJson(send []g.Map) string {
	if b, err := json.Marshal(send); err != nil {
		return ""
	} else {
		return string(b)
	}
}
