package controllers

// RentSchedule -
type RentSchedule struct {
	EXTNumber               int        `json:"EXTNumber,omitempty"`               //
	StartDate               Datetype   `json:"StartDate,omitempty"`               //
	VacancyEnd              Datetype   `json:"VacancyEnd,omitempty"`              //
	VacancyAmount           float64    `json:"VacancyAmount,omitempty"`           //
	RentIncentivesEndRenew  Datetype   `json:"RentIncentivesEndRenew,omitempty"`  //
	RentIncentivesEndRotate Datetype   `json:"RentIncentivesEndRotate,omitempty"` //
	DefaultDate             Datetype   `json:"DefaultDate,omitempty"`             //
	EndDate                 Datetype   `json:"EndDate,omitempty"`                 //
	OriginalEndDate         Datetype   `json:"OriginalEndDate,omitempty"`         //
	RenewRent               float64    `json:"RenewRent,omitempty"`               //
	RotateRent              float64    `json:"RotateRent,omitempty"`              //
	PassingRent             float64    `json:"PassingRent,omitempty"`             //
	EndContractRent         float64    `json:"EndContractRent,omitempty"`         //
	RentRevisionERV         float64    `json:"RentRevisionERV,omitempty"`         //
	Probability             float64    `json:"Probability,omitempty"`             //
	ProbabilitySim          float64    `json:"ProbabilitySim,omitempty"`          // either a 1 or 0 based on random sample
	RenewIndex              Indexation `json:"RenewIndex,omitempty"`              //
	RotateIndex             Indexation `json:"RotateIndex,omitempty"`             //
	ParentUnit              *UnitModel `json:"-"`                                 //
}

// Indexation -
type Indexation struct {
	IndexNumber  int           `json:"IndexNumber,omitempty"`
	StartDate    Datetype      `json:"StartDate,omitempty"`
	EndDate      Datetype      `json:"EndDate,omitempty"`
	Amount       float64       `json:"Amount,omitempty"`
	Final        float64       `json:"Final,omitempty"`
	RentSchedule *RentSchedule `json:"-"`
}

// IndexDetails -
type IndexDetails struct {
	Frequency   int    `json:"Frequency,omitempty"` //# of years
	Type        string `json:"Type,omitempty"`      // selection from parents GrowthInput
	StartMonth  int    `json:"StartMonth,omitempty"`
	Anniversary string `json:"Anniversary,omitempty"`
}

// InitialGrowth -
type InitialGrowth struct {
	Area   float64 `json:"Area,omitempty"`
	Amount float64 `json:"Amount,omitempty"`
}

// CostInput -
type CostInput struct {
	Amount                float64 `json:"Amount,omitempty"`
	AmountPerTotalArea    float64 `json:"AmountPerTotalArea,omitempty"`
	AmountPerOccupiedArea float64 `json:"AmountPerOccupiedArea,omitempty"`
	AmountPerVacantArea   float64 `json:"AmountPerVacantArea,omitempty"`
	PercentOfERV          float64 `json:"PercentOfERV,omitempty"`
	PercentOfTRI          float64 `json:"PercentOfTRI,omitempty"`
	PercentOfContractRent float64 `json:"PercentOfContractRent,omitempty"`
	PercentOfNAV          float64 `json:"PercentOfNAV,omitempty"`
	PercentOfGAV          float64 `json:"PercentOfGAV,omitempty"`
	IsCapitalized         bool    `json:"IsCapitalized,omitempty"`
	Duration              int     `json:"Duration,omitempty"`  // number of months
	IsIndexed             bool    `json:"IsIndexed,omitempty"` // if true, the cost grows with CPI
}
