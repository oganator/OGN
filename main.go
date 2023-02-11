package main

import (
	// "runtime/debug"

	_ "OGN/routers"
	"database/sql"
	"fmt"
	"sync"

	ogn "OGN/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Init() {
	ogn.InitVersion = "v1"
	ogn.ReadXLSX()
	ogn.PopulateModels()
	// parent assignment
	for i, v := range ogn.EntityModelsMap {
		v.EntityModel.Parent = ogn.EntityModelsMap[ogn.EntityDataStore[i].Parent].EntityModel
		v.EntityModel.PopulateChildEntities()
	}
	for _, v := range ogn.FundModelsList {
		ogn.EntityModelsMap[v].EntityModel.CalculateFund()
	}
	for _, v := range ogn.AssetModelsList {
		ogn.EntityModelsMap[v].EntityModel.MonteCarlo("Internal")
	}
	for _, v := range ogn.FundModelsList {
		if ogn.EntityModelsMap[v].EntityModel.MasterID == 0 {
			continue
		}
		// ogn.EntityModelsMap[v].EntityModel.FundMonteCarlo()
	}
	ogn.SimCounter = ogn.SimIDType{
		Mutex: &sync.Mutex{},
		ID:    0,
	}
	for _, v := range ogn.EntityMap {
		fmt.Println(v.Models)
	}
}

// Init2 - made to work with SQLite
func Init2() {
	ogn.InitVersion = "v2"
	ogn.ReadDB()
	for _, v := range ogn.AssetModelsList {
		ogn.EntityModelsMap[v].EntityModel.EntityModelCalc(false, "Internal")
	}
	for _, v := range ogn.FundModelsList {
		ogn.EntityModelsMap[v].EntityModel.UpdateFundModel()
	}
}

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	// runtime.GOMAXPROCS(1)
	// Init()
	ogn.DB, _ = sql.Open("sqlite", "./models/ogndata.db")
	defer ogn.DB.Close()
	Init2()
	// debug.SetGCPercent(2000)
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
