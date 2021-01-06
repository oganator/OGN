package controllers

// Yield -
type Yield struct {
	Entry float64
	Shift float64 //bps per year
	Exit  float64
}

// DebtInput -
type DebtInput struct {
	LTV          float64
	InterestRate float64
}

// Metrics -
type Metrics struct {
	IRR        ReturnType
	EM         ReturnType
	CoC        ReturnType
	TWR        ReturnType
	BondHolder BondReturnType
}

// ReturnType -
type ReturnType struct {
	GrossUnleveredBeforeTax float64
	NetLeveredAfterTax      float64
}

// BondReturnType -
type BondReturnType struct {
	Duration float64
	YTM      float64
}

// Default -
type Default struct {
	Hazard           float64
	NumberOfDefaults int
}

// StringIntFloatMap -
type StringIntFloatMap map[string]map[int]float64

// StringIntFloatMapPtr -
type StringIntFloatMapPtr map[string]IntFloatMap

// IntFloatMap -
type IntFloatMap map[int]float64

// StringFloatMap -
type StringFloatMap map[string]float64

// IntFloatCOAMap -
type IntFloatCOAMap map[int]FloatCOA

// Settings -
type Settings struct {
	Frequency       string // Monthly, Quarterly
	Type            string // Portfolio, JV, SPV, Asset, Conventional, Ground Lease, Income Strip, Retail, Residential, Renewable, Office
	DefaultScenario string // Re-let, Sale
	YearSpan        string // Fiscal, Calendar
}
