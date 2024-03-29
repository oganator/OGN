package controllers

import "sync"

// Entity -
type Entity struct {
	MasterID        int             `json:"MasterID,omitempty"`        //
	Name            string          `json:"Name,omitempty"`            //
	Models          []*EntityModel  `json:"Models,omitempty"`          //
	AcquisitionDate Datetype        `json:"AcquisitionDate,omitempty"` //
	DispositionDate Datetype        `json:"DispositionDate,omitempty"` //
	EntityType      string          `json:"EntityType,omitempty"`      // Asset, Fund...etc
	Mandate         Diversification `json:"Mandate,omitempty"`         //
	Allocation      Diversification `json:"Allocation,omitempty"`      //
	Actuals         IntFloatCOAMap  `json:"Actuals,omitempty"`         // Contains monthly values, as well as yearly values (rolling or calendar) since inception, no forecasted data
}

// EntityModel - Used to model an Asset, Fund, Property, Sub-Portfolio...etc
type EntityModel struct {
	Mutex                  *sync.Mutex          `json:"-"`                                //
	MasterID               int                  `json:"MasterID,omitempty"`               //
	Entity                 *Entity              `json:"-"`                                // Physical entity on which this model is based
	EntityID               int                  `json:"EntityID,omitempty"`               //
	EntityData             EntityModelData      `json:"EntityData,omitempty"`             // used for Azure MC
	Name                   string               `json:"Name,omitempty"`                   //
	Version                string               `json:"Version,omitempty"`                //
	ChildEntityModels      map[int]*EntityModel `json:"ChildEntities,omitempty"`          // MasterID as key. created by ModelCreate().
	ChildUnitModels        map[int]*UnitModel   `json:"ChildUnits,omitempty"`             // MasterID as key. created by ModelCreate().
	ChildUnitsMC           map[int]UnitModel    `json:"ChildUnitsMC,omitempty"`           // Used only for Azure MC. Needed to actually store the unit data when sent over to azure, otherwise azure only receives a memory address, no actual data
	Metrics                Metrics[float64]     `json:"Metrics,omitempty"`                //
	ParentID               int                  `json:"ParentID,omitempty"`               //
	Parent                 *EntityModel         `json:"-"`                                //
	StartDate              Datetype             `json:"StartDate,omitempty"`              // used for cash flow calculations
	HoldPeriod             int                  `json:"HoldPeriod,omitempty"`             //
	SalesDate              Datetype             `json:"SalesDate,omitempty"`              // used for cash flow calculations
	EndDate                Datetype             `json:"EndDate,omitempty"`                // pushed out further than sales date due to next buyers analysis
	GrowthInput            map[string]HModel    `json:"GrowthInput,omitempty"`            // where the master data of GrowthInput is stored. the float value is a yearly growth number. Populated from Growth Items tab
	Growth                 StringIntFloatMap    `json:"Growth,omitempty"`                 // where the result of the growth is stored (ERV and CPI). Created by Model.Create()
	DebtInput              []DebtInput          `json:"DebtInput,omitempty"`              // NOT a ptr because that would recreate the debt on each level. N/A for units.
	CostInput              IntCostInputMap      `json:"Capex,omitempty"`                  // yearly input for costs.
	GLA                    UnitModel            `json:"GLA,omitempty"`                    // general lease assumptions
	SensitivityInput       SensitivityInput     `json:"SensitivityInput,omitempty"`       // input for the sensitivity analysis
	SensitivityOutputArray [][][]float64        `json:"SensitivityOutputArray,omitempty"` // Keys are formatted as [variable1/variable2]. Each 'Metrics' struct stores the data in [][]float64 instead of float64
	SensitivityOutputKeys  map[string]int       `json:"SensitivityOutputKeys,omitempty"`  // Keys are formatted as [variable1/variable2]. Each 'Metrics' struct stores the data in [][]float64 instead of float64
	SensitivityOutputItems map[int]string       `json:"SensitivityOutputItems,omitempty"` // Keys are formatted as [variable1/variable2]. Each 'Metrics' struct stores the data in [][]float64 instead of float64
	MC                     bool                 `json:"MC,omitempty"`                     //
	MCSetup                MCSetup              `json:"MCSetup,omitempty"`                // used to store sigmas of the variables - the mean values are all stored elsewhere
	MCSlice                []*EntityModel       `json:"MCSlice,omitempty"`                //
	MCResultSlice          MCResultSlice        `json:"MCResultSlice,omitempty"`          //
	MCResults              MCResults            `json:"MCResults,omitempty"`              //
	FactorAnalysis         []FactorIndependant  `json:"FactorAnalysis,omitempty"`         //
	Tax                    Tax                  `json:"Tax,omitempty"`                    //
	COA                    IntFloatCOAMap       `json:"COA,omitempty"`                    // Contains monthly values, as well as yearly values (rolling or calendar) up to the sales date
	Valuation              Valuation            `json:"Valuation,omitempty"`              //
	TableHeader            HeaderType           `json:"TableHeader,omitempty"`            // Years, Months...etc
	Table                  []TableJSON          `json:"Table,omitempty"`                  //
	RetoolTable            []interface{}        `json:"RetoolTable,omitempty"`            //
	Strategy               string               `json:"Strategy,omitempty"`               //
	UOM                    string               `json:"UOM,omitempty"`                    //
	BalloonPercent         float64              `json:"BalloonPercent,omitempty"`         //
}

// UnitModel -
type UnitModel struct {
	// Mutex              *sync.Mutex    `json:"-"`                            //
	MasterID        int             `json:"MasterID,omitempty"`        //
	Name            string          `json:"Name,omitempty"`            //
	LeaseStartDate  Datetype        `json:"LeaseStartDate,omitempty"`  //
	LeaseExpiryDate Datetype        `json:"LeaseExpiryDate,omitempty"` //
	UnitStatus      string          `json:"UnitStatus,omitempty"`      // vacant or occupied
	Tenant          string          `json:"Tenant,omitempty"`          //
	TenantType      string          `json:"TenantType,omitempty"`      // Commercial, Residential
	PassingRent     float64         `json:"PassingRent,omitempty"`     //
	RentSchedule    RentSchedule    `json:"RentSchedule,omitempty"`    // created by Unit.RentScheduleCalc()
	RSStore         []RentSchedule  `json:"RSStore,omitempty"`         // only used for reference
	Parent          *EntityModel    `json:"-"`                         //
	Probability     float64         `json:"Probability,omitempty"`     //
	PercentSoldRent float64         `json:"PercentSoldRent,omitempty"` //
	DiscountRate    float64         `json:"DiscountRate,omitempty"`    //
	BondProceeds    float64         `json:"BondProceeds,omitempty"`    //
	BondIncome      float64         `json:"BondIncome,omitempty"`      //
	BondIndex       Indexation      `json:"BondIndex,omitempty"`       //
	BondExpense     float64         `json:"BondExpense,omitempty"`     //
	Default         Default         `json:"Default,omitempty"`         //
	RentRevisionERV float64         `json:"RentRevisionERV,omitempty"` //
	EXTDuration     int             `json:"EXTDuration,omitempty"`     //
	IndexDetails    IndexDetails    `json:"IndexDetails,omitempty"`    //
	Void            int             `json:"Void,omitempty"`            //
	ERVArea         float64         `json:"ERVArea,omitempty"`         //
	ERVAmount       float64         `json:"ERVAmount,omitempty"`       //
	COA             IntFloatCOAMap  `json:"COA,omitempty"`             // Contains monthly values, as well as yearly values (rolling or calendar) up to the sales date\
	CostInput       IntCostInputMap `json:"Capex,omitempty"`           //
	MCSetup         MCSetup         `json:"MCSetup,omitempty"`         // used to store sigmas of the variables - the mean values are all stored elsewhere
	Override        *UnitModel      `json:"Override,omitempty"`        //
}

// HeaderType -
type HeaderType struct {
	Monthly   []Datetype `json:"Monthly,omitempty"`   //
	Quarterly []Datetype `json:"Quarterly,omitempty"` //
	Yearly    []Datetype `json:"Yearly,omitempty"`    //
}

// Valuation -
type Valuation struct {
	Method          string          `json:"Method,omitempty"`          // DirectCap, DCF, German, Manual, VacantPossesion
	EntryYield      float64         `json:"EntryYield,omitempty"`      //
	YieldShift      float64         `json:"YieldShift,omitempty"`      //
	ExitYield       float64         `json:"ExitYield,omitempty"`       //
	DiscountRate    float64         `json:"DiscountRate,omitempty"`    //
	AcqPrice        float64         `json:"AcqPrice,omitempty"`        //
	Fees            IntCostInputMap `json:"Fees,omitempty"`            //
	IncomeCapSetup  FloatCOA        `json:"IncomeCapSetup,omitempty"`  //
	IncomeDeduction FloatCOA        `json:"IncomeDeduction,omitempty"` // Deductions to income after it is capped
}
