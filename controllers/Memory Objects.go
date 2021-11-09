package controllers

// EntityStore -
var EntityStore = map[int]*EntityData{}

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
var Entities = map[int]*Entity{}

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
	for i, v := range Entities {
		v.Parent = Entities[EntityStore[i].Parent]
		v.PopulateChildEntities()
	}
}
