package list

import (
	"encoding/json"
	"gf-app/app/model/api_detail"
	"gf-app/app/model/api_group"
	"gf-app/app/model/project"
	"gf-app/app/model/whitelist"
	"gf-app/library"
	"github.com/gogf/gf/frame/g"
	"time"
)

//Get assign data
func GetAssignData(pid interface{}, gid interface{}, aid interface{}) g.Map {
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
	detail["name"] = ""
	detail["uri"] = ""
	detail["request_type"] = ""
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

	return g.Map{"table_data": library.GetJson(send), "detail_data": detail, "gid": gid, "aid": aid}
}

//Update detail
func UpdateDetial(id interface{}, uri interface{}, name interface{}, requestType interface{}, param interface{}, returnValue interface{}) bool {
	_, err := api_detail.Update(g.Map{"uri": uri, "name": name, "request_type": requestType, "param": param, "return_value": returnValue, "update_time": time.Now().Unix()}, g.Map{"id": id})
	if err != nil {
		return false
	}
	return true
}

//Delete detail
func DeleteDetail(id interface{}) bool {
	_, err := api_detail.Delete(g.Map{"id": id})
	if err != nil {
		return false
	}
	return true
}

//Add group
func AddGroup(pid interface{}, name interface{}) bool {
	_, err := api_group.Insert(g.Map{"pid": pid, "name": name})
	if err != nil {
		return false
	}
	return true
}

//Add Detail
func AddDetail(detail *api_detail.Entity, ip string) bool {
	user, _ := whitelist.FindOne("ip", ip)
	detail.Author = user.Name
	detail.Time = time.Now().Unix()
	if len(detail.Param) == 0 {
		detail.Param = "{}"
	}
	if len(detail.ReturnValue) == 0 {
		detail.ReturnValue = "{}"
	}
	_, err := api_detail.Insert(detail)
	if err != nil {
		return false
	}
	return true
}

//Add project
func AddProject(name interface{}) int64 {
	res, err := project.Insert(g.Map{"name": name})
	if err != nil {
		return 0
	}
	id, _ := res.LastInsertId()
	return id
}

//Update project
func UpdateProject(id interface{}, name interface{}) bool {
	_, err := project.Update(g.Map{"name": name}, g.Map{"id": id})
	if err != nil {
		return false
	}
	return true
}

//Update group
func UpdateGroup(id interface{}, name interface{}) bool {
	_, err := api_group.Update(g.Map{"name": name}, g.Map{"id": id})
	if err != nil {
		return false
	}
	return true
}
