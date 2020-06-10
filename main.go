package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
	"time"
)

func main() {
	svr := g.Server()
	svr.SetConfigWithMap(g.Map{
		"SessionMaxAge": time.Minute,
	})
	svr.Run()
}
