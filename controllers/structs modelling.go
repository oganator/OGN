package controllers

// RentSchedule -
type RentSchedule struct {
	EXTNumber         int        `json:"EXTNumber,omitempty"`          //
	StartDate         Datetype   `json:"StartDate,omitempty"`          //
	VacancyEnd        Datetype   `json:"VacancyEnd,omitempty"`         //
	VacancyAmount     float64    `json:"VacancyAmount,omitempty"`      //
	RentIncentivesEnd Datetype   `json:"RenewIncentivesEnd,omitempty"` //
	DefaultDate       Datetype   `json:"DefaultDate"`                  //
	EndDate           Datetype   `json:"EndDate,omitempty"`            //
	OriginalEndDate   Datetype   `json:"OriginalEndDate,omitempty"`    //
	RenewRent         float64    `json:"RenewRent,omitempty"`          //
	RotateRent        float64    `json:"RotateRent,omitempty"`         //
	PassingRent       float64    `json:"PassingRent,omitempty"`        //
	EndContractRent   float64    `json:"EndContractRent,omitempty"`    //
	RentRevisionERV   float64    `json:"RentRevisionERV,omitempty"`    //
	Probability       float64    `json:"Probability,omitempty"`        //
	ProbabilitySim    float64    `json:"ProbabilitySim,omitempty"`     //
	RenewIndex        Indexation `json:"RenewIndex,omitempty"`         //
	RotateIndex       Indexation `json:"RotateIndex,omitempty"`        //
	ParentUnit        *Unit      `json:"-"`                            //
}

// Indexation -
type Indexation struct {
	IndexNumber int      `json:"IndexNumber,omitempty"`
	StartDate   Datetype `json:"StartDate,omitempty"`
	EndDate     Datetype `json:"EndDate,omitempty"`
	Amount      float64  `json:"Amount,omitempty"`
	Final       float64  `json:"Final,omitempty"`
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
	Area   float64
	Amount float64
}

// CostInput -
type CostInput struct {
	Amount                float64
	AmountPerTotalArea    float64
	AmountPerOccupiedArea float64
	AmountPerVacantArea   float64
	// GrowthRateofType      []GrowthRateInput
	PercentOfERV          float64
	PercentOfTRI          float64
	PercentOfContractRent float64
	PercentOfNAV          float64
	PercentOfGAV          float64
}
