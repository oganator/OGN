package main

import (
	_ "OGN/routers"
	"fmt"
	"math/rand"
	"time"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST"},
		// AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	end := time.Date(2021, 11, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println("This Trial ends: ", end)
	if time.Now().Before(end) {
		beego.Run()
	}
}
