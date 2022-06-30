package controllers

// Yield -
type Yield struct {
	Entry float64 `json:"Entry,omitempty"`
	Shift float64 `json:"Shift,omitempty"` //bps per year
	Exit  float64 `json:"Exit,omitempty"`
}

// DebtInput -
type DebtInput struct {
	LTV          float64 `json:"LTV,omitempty"`
	InterestRate float64 `json:"InterestRate,omitempty"`
}

// Metrics -
type Metrics struct {
	IRR        ReturnType     `json:"irr,omitempty"`
	EM         ReturnType     `json:"em,omitempty"`
	CoC        ReturnType     `json:"coc,omitempty"`
	TWR        ReturnType     `json:"twr,omitempty"`
	BondHolder BondReturnType `json:"bondholder,omitempty"`
}

// ReturnType -
type ReturnType struct {
	GrossUnleveredBeforeTax float64 `json:"GrossUnleveredBeforeTax,omitempty"`
	NetLeveredAfterTax      float64 `json:"NetLeveredAfterTax,omitempty"`
}

// BondReturnType -
type BondReturnType struct {
	Duration float64 `json:"Duration,omitempty"`
	YTM      float64 `json:"YTM,omitempty"`
	YTMDUR   float64 `json:"YTMDUR,omitempty"`
}

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
