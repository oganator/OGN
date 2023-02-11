package controllers

type Industry struct {
	Sector    interface{} `json:"Sector,omitempty"`    // Office, Industrial, Retail or Residential
	SubSector interface{} `json:"SubSector,omitempty"` //
}

type Geography struct {
	Region  interface{} `json:"Region,omitempty"`  //
	Country interface{} `json:"Country,omitempty"` //
	MSA     interface{} `json:"MSA,omitempty"`     //
}

type Diversification struct {
	Industry  Industry  `json:"Industry,omitempty"`  //
	Geography Geography `json:"Geography,omitempty"` //
}

type DetailsInput struct {
	Name           bool
	IRR            bool
	EM             bool
	StartDate      bool
	HoldPeriod     bool
	SalesDate      bool
	Growth         bool
	Debt           bool
	FactorAnalysis bool
	Valuation      bool
	Strategy       bool
	Mandate        bool
	Allocation     bool
}

type KeyValue struct {
	Key   string      `json:"Key,omitempty"`   //
	Value interface{} `json:"Value,omitempty"` //
}

// ReturnDetails -
func (e *EntityModel) ModelDetails(input DetailsInput) (final []KeyValue) {
	// final = make([]KeyValue, 11)
	if input.Name {
		temp := KeyValue{
			Key:   "Name",
			Value: e.Name,
		}
		final = append(final, temp)
	}
	if input.IRR {
		temp := KeyValue{
			Key:   "IRR",
			Value: e.Metrics.IRR.NetLeveredAfterTax,
		}
		final = append(final, temp)
	}
	if input.EM {
		temp := KeyValue{
			Key:   "EM",
			Value: e.Metrics.EM.NetLeveredAfterTax,
		}
		final = append(final, temp)
	}
	if input.StartDate {
		temp := KeyValue{
			Key:   "Start Date",
			Value: e.StartDate,
		}
		final = append(final, temp)
	}
	if input.HoldPeriod {
		temp := KeyValue{
			Key:   "Hold Period",
			Value: e.HoldPeriod,
		}
		final = append(final, temp)
	}
	if input.SalesDate {
		temp := KeyValue{
			Key:   "Sales Date",
			Value: e.SalesDate,
		}
		final = append(final, temp)
	}
	if input.Growth {
		temp := KeyValue{
			Key:   "Growth",
			Value: e.GrowthInput,
		}
		final = append(final, temp)
	}
	if input.Debt {
		temp := KeyValue{
			Key:   "Debt",
			Value: e.DebtInput,
		}
		final = append(final, temp)
	}
	if input.FactorAnalysis {
		temp := KeyValue{
			Key:   "Factor Analysis",
			Value: e.FactorAnalysis,
		}
		final = append(final, temp)
	}
	if input.Valuation {
		temp := KeyValue{
			Key:   "Valuation",
			Value: e.Valuation,
		}
		final = append(final, temp)
	}
	if input.Strategy {
		temp := KeyValue{
			Key:   "Strategy",
			Value: e.Strategy,
		}
		final = append(final, temp)
	}
	if input.Mandate {
		temp := KeyValue{
			Key:   "Mandate",
			Value: e.Entity.Mandate,
		}
		final = append(final, temp)
	}
	if input.Allocation {
		temp := KeyValue{
			Key:   "Allocation",
			Value: e.Entity.Allocation,
		}
		final = append(final, temp)
	}
	return final
}
