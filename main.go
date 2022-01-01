package main

import (
	"runtime/debug"
	"sync"

	_ "github.com/Oganator/OGN/routers"

	ogn "github.com/Oganator/OGN/controllers"

	beego "github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Init() {
	ogn.ReadXLSX()
	ogn.PopulateModels()
	// parent assignment
	for i, v := range ogn.EntityMap {
		v.Entity.Parent = ogn.EntityMap[ogn.EntityDataStore[i].Parent].Entity
		v.Entity.PopulateChildEntities()
	}
	for _, v := range ogn.FundsList {
		ogn.EntityMap[v].Entity.CalculateFund()
	}
	for _, v := range ogn.ModelsList {
		ogn.EntityMap[v].Entity.MonteCarlo("Internal")
	}
	for _, v := range ogn.FundsList {
		if ogn.EntityMap[v].Entity.MasterID == 0 {
			continue
		}
		ogn.EntityMap[v].Entity.FundMonteCarlo()
	}
	ogn.SimCounter = ogn.SimIDType{
		Mutex: &sync.Mutex{},
		ID:    0,
	}
}

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	// runtime.GOMAXPROCS(1)
	Init()
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
