package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

func Auth(r *ghttp.Request) {
	//Check auth through IP.
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

	//Check auth through login session.
	userinfo := r.Session.Get("sys_user")
	if userinfo != nil {
		r.Middleware.Next()
	} else {
		//Session was expired,redirecting to login page.
		//r.Response.WriteStatus(http.StatusForbidden)
		r.Response.RedirectTo("/login")
	}
}
