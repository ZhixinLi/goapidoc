package middleware

import (
	"gf-app/app/model/whitelist"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Auth(r *ghttp.Request) {
	check, err := whitelist.FindOne("ip", r.GetClientIp())
	if err != nil {
		r.Response.WriteStatus(http.StatusBadGateway)
	}

	if check == nil {
		r.Response.WriteStatus(http.StatusForbidden)
	} else {
		r.Middleware.Next()
	}
}
