package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Auth(r *ghttp.Request) {
	//check, err := whitelist.FindOne("ip", r.GetClientIp())
	//if err != nil {
	//	r.Response.WriteStatus(http.StatusBadGateway)
	//}
	//
	//if check == nil {
	//	r.Response.WriteStatus(http.StatusForbidden)
	//} else {
	//	r.Middleware.Next()
	//}

	userinfo := r.Session.Get("user")
	fmt.Println(userinfo)
	if userinfo != nil {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
