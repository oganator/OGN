package controllers

// Default -
type Default struct {
	Hazard           float64  `json:"Hazard,omitempty"`
	DefaultEnd       Datetype `json:"DefaultEnd,omitempty"`
	NumberOfDefaults int      `json:"NumberOfDefaults,omitempty"`
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
