package login

import (
	"gf-app/app/model/whitelist"
	"github.com/gogf/gf/frame/g"
)

func CheckLogin(name string, pwd string) bool {
	res, err := whitelist.FindOne(g.Map{"name": name, "pwd": pwd})
	if err != nil {
		return false
	}

	if res != nil {
		return true
	} else {
		return false
	}
}
