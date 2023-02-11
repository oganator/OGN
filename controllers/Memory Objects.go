package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// DATASTORES
// EntityStore - Entitymodel MasterID as key
var EntityDataStore = map[int]*EntityModelData{}

// UnitStore -
var UnitStore = map[int]UnitModelData{}

// GrowthItemsStore -
var GrowthItemsStore = make(map[int]map[string]float64)

//////////////////////////////////////////////////////////////////////////////////
// PHYSICAL ENTITES

var EntityMap = map[int]*Entity{}
var EntitiesList = make(map[string]int)
var FundsList = make(map[string]int)
var AssetsList = make(map[string]int)

//////////////////////////////////////////////////////////////////////////////////
// ENTITY MODELS

type EntityMutexMap map[int]EntityMutex
type EntityMutex struct {
	Mutex       *sync.Mutex  `json:"-"`
	EntityModel *EntityModel `json:"EntityModel,omitempty"`
}

// EntityModelsMap - MasterID as key
var EntityModelsMap = EntityMutexMap{}

// EnitiesList - all entities, funds and assets
var EntityModelsList = make(map[string]int)

// ModelsList - Just the Assets, no funds
var AssetModelsList = make(map[string]int)

// FundsList -
var FundModelsList = make(map[string]int)

///////////////////////////////////////////////

// Units -
var Units = map[int]UnitModel{}

// Key -
var Key = 1

var SimCounter SimIDType

// var AzureURL = "http://localhost:8081/Function"
var AzureURL = "http://localhost:7071/api/OGNTrigger"

var Monthly = false

var Compute = "Internal" // Internal or Azure

var InitVersion = ""

type SimIDType struct {
	Mutex *sync.Mutex
	ID    int
}

func StructPrint(temp interface{}) {
	empJSON, err := json.MarshalIndent(temp, "", "	")
	if err != nil {
		fmt.Println("StructPrint Error: ", err)
	}

	//get current function
	counter, _, _, success := runtime.Caller(1)
	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	funcname := runtime.FuncForPC(counter).Name()

	fmt.Printf(funcname+"%s\n", string(empJSON))
}

var CFTableCOA = BoolCOA{
	MarketValue:             true,
	TotalERV:                true,
	OccupiedERV:             false,
	VacantERV:               false,
	TopSlice:                false,
	TotalArea:               false,
	OccupiedArea:            false,
	VacantArea:              false,
	PassingRent:             true,
	Indexation:              true,
	TheoreticalRentalIncome: true,
	BPUplift:                false,
	Vacancy:                 true,
	ContractRent:            true,
	RentFree:                false,
	TurnoverRent:            false,
	MallRent:                false,
	ParkingIncome:           false,
	OtherIncome:             false,
	OperatingIncome:         true,
	OperatingExpenses:       true,
	NetOperatingIncome:      true,
	AcqDispProperty:         true,
	AcqDispCosts:            false,
	LoanProceeds:            false,
	InterestExpense:         true,
	LoanBalance:             false,
	Debt:                    false,
	Tax:                     true,
	TaxableIncome:           false,
	TaxableIncomeCarryBack:  false,
	DTA:                     false,
	Depreciation:            false,
	Capex:                   true,
	Fees:                    true,
	NetCashFlow:             true,
	CashBalance:             true,
	BondIncome:              false,
	BondExpense:             false,
}
