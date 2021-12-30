package controllers

import "sync"

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

// var BaseURL = "" //"http://localhost:8080/"

var AzureURL = "http://localhost:8081/Function"

var Monthly = false

var Compute = "Azure" // Internal or Azure

type SimIDType struct {
	Mutex *sync.Mutex
	ID    int
}
