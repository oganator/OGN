package controllers

import (
	"encoding/json"
	"fmt"
	"sync"
)

// EntityStore - Entity MasterID as key
var EntityDataStore = map[int]*EntityData{}

// UnitStore -
var UnitStore = map[int]UnitData{}

// GrowthItemsStore -
var GrowthItemsStore = make(map[int]map[string]float64)

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

var SimCounter SimIDType

// var AzureURL = "http://localhost:8081/Function"
var AzureURL = "http://localhost:7071/api/OGNTrigger"

var Monthly = false

var Compute = "Internal" // Internal or Azure

type SimIDType struct {
	Mutex *sync.Mutex
	ID    int
}

func StructPrint(name string, temp interface{}) {
	empJSON, err := json.MarshalIndent(temp, "", "	")
	if err != nil {
		fmt.Println("StructPrint Error: ", err)
	}
	fmt.Printf(name+"%s\n", string(empJSON))
}
