package controllers

// EntityStore -
var EntityStore = map[int]*EntityData{}

// UnitStore -
var UnitStore = map[int]UnitData{}

// // GrowthItemsRaw -
// var GrowthItemsRaw = GDSlice{}

//GrowthItemsStore -
var GrowthItemsStore = make(map[int]map[string]float64)

// EntityAssociations -
var EntityAssociations = make(map[int][]int)

// UnitAssociations -
var UnitAssociations = make(map[int][]int)

// Models - MasterID as key
var Models = map[int]*Entity{}

// ModelsList -
var ModelsList = make(map[string]int)

// Units -
var Units = map[int]Unit{}

// Key -
var Key = 1

func init() {
	ReadXLSX()
	Associations()
	PopulateModels()
}
