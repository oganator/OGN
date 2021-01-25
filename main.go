package main

import (
	_ "OGN/routers"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST"},
		// AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
