package controllers

import "sync"

// Entity -
type Entity struct {
	Mutex          *sync.Mutex `json:"MCMutex,omitempty"` //
	MasterID       int
	Name           string            `json:"Name,omitempty"`          //
	ChildEntities  map[int]*Entity   `json:"ChildEntities,omitempty"` // MasterID as key. created by ModelCreate().
	ChildUnits     map[int]*Unit     `json:"ChildUnits,omitempty"`    // MasterID as key. created by ModelCreate().
	Metrics        Metrics           `json:"Metrics,omitempty"`       //
	ParentID       int               `json:"ParentID,omitempty"`
	Parent         *Entity           `json:"-"`                         //
	StartDate      Datetype          `json:"StartDate,omitempty"`       // used for cash flow calculations
	HoldPeriod     int               `json:"HoldPeriod,omitempty"`      //
	SalesDate      Datetype          `json:"SalesDate,omitempty"`       // used for cash flow calculations
	EndDate        Datetype          `json:"EndDate,omitempty"`         // pushed out further than sales date due to next buyers analysis
	GrowthInput    map[string]HModel `json:"GrowthInputData,omitempty"` // where the master data of GrowthInput is stored. the float value is a yearly growth number. Populated from Growth Items tab
	Growth         StringIntFloatMap `json:"Growth,omitempty"`          // where the result of the growth is stored (ERV and CPI). Created by Model.Create()
	DebtInput      DebtInput         `json:"Debt,omitempty"`            // NOT a ptr because that would recreate the debt on each level. N/A for units.
	OpEx           CostInput         `json:"OpEx,omitempty"`            // yearly input for costs.
	Fees           CostInput         `json:"Fees,omitempty"`            //
	Capex          map[int]CostInput `json:"Capex,omitempty"`           // yearly input for costs.
	GLA            Unit              `json:"GLA,omitempty"`             //
	MC             bool              `json:"MC,omitempty"`              //
	MCSetup        MCSetup           `json:"MCSetup,omitempty"`         //
	MCSlice        []*Entity         `json:"MCSlice,omitempty"`         //
	MCResultSlice  MCResultSlice     `json:"MCResultSlice,omitempty"`   //
	MCResults      MCResults         `json:"MCResults,omitempty"`       //
	Tax            Tax               `json:"Tax,omitempty"`             //
	COA            IntFloatCOAMap    `json:"COA,omitempty"`             // Contains monthly values, as well as yearly values (rolling or calendar) up to the sales date
	Valuation      Valuation         `json:"Valuation,omitempty"`       //
	TableHeader    HeaderType        `json:"TableHeader,omitempty"`     // Years, Months...etc
	Table          []TableJSON       `json:"Table,omitempty"`           //
	Strategy       string            `json:"Strategy,omitempty"`        //
	BalloonPercent float64
}

// Unit -
type Unit struct {
	MasterID           int
	Name               string         `json:"Name,omitempty"`            //
	LeaseStartDate     Datetype       `json:"LeaseStartDate,omitempty"`  //
	LeaseExpiryDate    Datetype       `json:"LeaseExpiryDate,omitempty"` //
	UnitStatus         string         `json:"UnitStatus,omitempty"`      // vacant or occupied
	Tenant             string         `json:"Tenant,omitempty"`          //
	PassingRent        float64        `json:"PassingRent,omitempty"`     //
	RentSchedule       RentSchedule   `json:"RentSchedule,omitempty"`    // created by Unit.RentScheduleCalc()
	RSStore            []RentSchedule `json:"RSStore,omitempty"`         // only used for reference
	Parent             *Entity        `json:"-"`                         //
	Probability        float64        `json:"Probability,omitempty"`     //
	PercentSoldRent    float64        `json:"PercentSoldRent,omitempty"` //
	DiscountRate       float64        `json:"DiscountRate,omitempty"`    //
	BondProceeds       float64        `json:"BondProceeds,omitempty"`    //
	BondIncome         float64        `json:"BondIncome,omitempty"`      //
	BondIndex          Indexation     `json:"BondIndex,omitempty"`       //
	BondExpense        float64        `json:"BondExpense,omitempty"`     //
	Default            Default        `json:"Default,omitempty"`         //
	RentRevisionERV    float64        `json:"RentRevisionERV,omitempty"` //
	EXTDuration        int            `json:"EXTDuration,omitempty"`     //
	IndexDetails       IndexDetails   `json:"IndexDetails,omitempty"`    //
	RentIncentives     CostInput      `json:"RentIncentives,omitempty"`  //
	Void               int            `json:"Void,omitempty"`            //
	FitOutCosts        CostInput      `json:"FitOutCosts,omitempty"`     // input for costs when the lease expires
	LeasingCommissions CostInput      `json:"LeasingCommissions,omitempty"`
	ERVArea            float64        `json:"ERVArea,omitempty"`   //
	ERVAmount          float64        `json:"ERVAmount,omitempty"` //
	COA                IntFloatCOAMap `json:"COA,omitempty"`       // Contains monthly values, as well as yearly values (rolling or calendar) up to the sales date\
}

// ChildEntities -
type ChildEntities struct {
	Keys  map[string]int
	Array []Entity
}

// ChildUnits -
type ChildUnits struct {
	Keys  map[string]int
	Array []Unit
}

// HeaderType -
type HeaderType struct {
	Monthly   []Datetype
	Quarterly []Datetype
	Yearly    []Datetype `json:"Yearly,omitempty"`
}

// Valuation -
type Valuation struct {
	EntryYield      float64        //
	YieldShift      float64        //
	ExitYield       float64        //
	DiscountRate    float64        //
	AcqPrice        float64        //
	AcqFees         map[string]Fee //
	DispFees        map[string]Fee //
	IncomeCapSetup  FloatCOA       //
	IncomeDeduction FloatCOA       // Deductions to income after it is capped
}

// Fee -
type Fee struct {
	Base    string // Can be Yield or Net Price. If Yield is selected, then the value expressed as a decimal + 1 (4% is 1.04) and multiplied into the yield.
	Percent float64
}
