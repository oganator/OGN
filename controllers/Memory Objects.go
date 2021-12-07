package controllers

import "sync"

// EntityStore -
var EntityDataStore = map[int]*EntityData{}

// UnitStore -
var UnitStore = map[int]UnitData{}

// // GrowthItemsRaw -
// var GrowthItemsRaw = GDSlice{}

//GrowthItemsStore -
var GrowthItemsStore = make(map[int]map[string]float64)

// // EntityAssociations -
// var EntityAssociations = make(map[int][]int)

// // UnitAssociations -
// var UnitAssociations = make(map[int][]int)

// Entities - MasterID as key
var EntityMap = EntityMutexMap{}

type EntityMutexMap map[int]EntityMutex

type EntityMutex struct {
	Mutex  *sync.Mutex
	Entity *Entity
}

// EnitiesList - all entities, funds and assets
var EntitiesList = make(map[string]int)

// ModelsList - Just the Assets, no funds
var ModelsList = make(map[string]int)

// FundsList -
var FundsList = make(map[string]int)

// Units -
var Units = map[int]Unit{}

// Key -
var Key = 1

func init() {
	ReadXLSX()
	PopulateModels()
	// parent assignment
	for i, v := range EntityMap {
		v.Entity.Parent = EntityMap[EntityDataStore[i].Parent].Entity
		v.Entity.PopulateChildEntities()
	}
	for _, v := range FundsList {
		EntityMap[v].Entity.CalculateFund()
	}
	for _, v := range ModelsList {
		EntityMap[v].Entity.MonteCarlo()
	}
	for _, v := range FundsList {
		if EntityMap[v].Entity.MasterID == 0 {
			continue
		}
		EntityMap[v].Entity.FundMonteCarlo()
	}
}

var BaseURL = "" //"http://localhost:8080/"

// var BaseURL = "https://oganica.azurewebsites.net/"

var Monthly = false

var Compute = "Internal" // Internal or Azure
