package main

import (
	_ "OGN/routers"
	"runtime/debug"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	// runtime.GOMAXPROCS(1)
	debug.SetGCPercent(2000)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST"},
		// AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// browser.OpenURL("http://localhost:8080/")
	beego.Run()
}
