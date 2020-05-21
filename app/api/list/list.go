package list

import (
	"encoding/json"
	"gf-app/app/model/api_detail"
	"gf-app/app/model/api_group"
	"gf-app/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	pid := r.Get("pid")
	gid := r.Get("gid")
	aid := r.Get("aid")

	//Get default group id
	if gid == nil {
		defaultGroup, _ := api_group.FindOne("pid", pid)
		if defaultGroup == nil {
			gid = 0
		} else {
			gid = defaultGroup.Id
		}
	}

	//get default api id
	if aid == nil {
		aid = 0
	}

	//Group info
	var send []g.Map
	data, _ := api_detail.FindAll("gid", gid)
	for _, v := range data {
		send = append(send, g.Map{
			"id":     v.Id,
			"name":   v.Name,
			"uri":    v.Uri,
			"author": v.Author,
			"time":   time.Unix(v.Time, 0).Format("2006-01-02 15:04:05"),
		})
	}

	//Detail info
	var detail g.Map = make(g.Map)
	detail["return_value"] = "{}"
	if aid != 0 {
		data, _ := api_detail.FindOne("id", aid)
		if data != nil {
			detail["name"] = data.Name
			detail["uri"] = data.Uri
			detail["request_type"] = data.RequestType

			var paramResult map[string]interface{}
			json.Unmarshal([]byte(data.Param), &paramResult)

			var paramMap []g.Map
			for k, v := range paramResult {
				paramMap = append(paramMap, g.Map{"key": k, "value": v})
			}
			detail["param"] = paramMap
			detail["return_value"] = data.ReturnValue
		}
	}

	//Display
	library.Fetch(r, "list/index", g.Map{"table_data": library.GetJson(send), "detail_data": detail, "gid": gid, "aid": aid})
}

func (c *Controller) Modify(r *ghttp.Request) {
	id := r.Get("id")
	param := r.Get("param")
	return_value := r.Get("return_value")

	api_detail.Update(g.Map{"param": param, "return_value": return_value, "update_time": time.Now().Unix()}, g.Map{"id": id})
}
