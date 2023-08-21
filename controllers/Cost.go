package controllers

import (
	"reflect"
)

// CostInput -
type CostInput struct {
	Name              string            `json:"Name,omitempty"`              //
	MasterID          int               `json:"MasterID,omitempty"`          //
	Type              string            `json:"Type,omitempty"`              // "Capex", "OpEx", "Fees", "Rent Incentives", "Fit Out Costs", "Leasing Commissions"
	Amount            float64           `json:"Amount,omitempty"`            //
	AmountSigma       float64           `json:"AmountSigma,omitempty"`       //
	COAItemBasis      string            `json:"COAItemBasis,omitempty"`      // The basis of the calculation
	COAItemTarget     string            `json:"COAItemTarget,omitempty"`     // Where the result is stored
	Duration          int               `json:"Duration,omitempty"`          // in months
	DurationSigma     int               `json:"DurationSigma,omitempty"`     // in months
	Start             Datetype          `json:"Start,omitempty"`             //
	StartEvent        string            `json:"StartEvent,omitempty"`        // ModelStart, LeaseRenewStart, LeaseRotateStart, LoanAcq
	End               Datetype          `json:"End,omitempty"`               //
	EndEvent          string            `json:"EndEvent,omitempty"`          // ModelEnd, SalesDate, LeaseRenewEnd, LeaseRotateEnd, LoanTermination
	GrowthItem        string            `json:"GrowthItem,omitempty"`        // "None", "CPI", "ERV"...etc
	GrowthItemOptions map[string]HModel `json:"GrowthItemOptions,omitempty"` // used for templates
}

type IntCostInputMap map[int]CostInput

// CapexCalc -
// The `CostMapCalc` function is a method of the `IntCostInputMap` type. It calculates the total cost
// amount for a given date range.
func (costMap IntCostInputMap) CostMapCalc(date Datetype, growth StringIntFloatMap, coa FloatCOA) (amount FloatCOA) {
	for _, cost := range costMap {
		if date.Dateint >= cost.Start.Dateint && date.Dateint <= cost.End.Dateint {
			amount.Add(CostCalc(cost, date, growth, coa))
		}
	}
	return amount
}

// CostCalc -
// TODO - set up StartEvent end EndEvent, probably not within this function though
// TODO - remove use of reflect. Instead, write switch statements. These compiled statements will execute much faster that reflect at runtime
func CostCalc(cost CostInput, date Datetype, growthInput StringIntFloatMap, coa FloatCOA) (amount FloatCOA) {
	cost.Duration = dateintdiff(cost.End.Dateint, cost.Start.Dateint)
	// duration := float64(cost.Duration)
	// TEST FIX: remove duration from calculation - I am seeing that if I have operating expenses set to be based off of total area, it is taking the first monthly value
	// and spreading that over the time horizon of the cost instead of applying the cost to each month

	// duration = 1
	// if cost.Type == "Rent Incentives" {
	// 	// duration = 1
	// }
	if cost.COAItemBasis == "TotalArea" || cost.COAItemBasis == "OccupiedArea" || cost.COAItemBasis == "VacantArea" {
		cost.Amount = cost.Amount / 12
	}
	coaAmt := 1.0
	if cost.COAItemBasis != "None" {
		coaAmt = reflect.ValueOf(coa).FieldByName(cost.COAItemBasis).Float()
	}
	growth := 1.0
	if cost.GrowthItem != "None" {
		growth = growthInput[cost.GrowthItem][date.Dateint]
	}
	tempAmount := cost.Amount * growth * coaAmt // duration
	reflect.ValueOf(&amount).Elem().FieldByName(cost.COAItemTarget).SetFloat(-tempAmount)
	return amount
}

// The `CostMapSetup()` function is a method of the `EntityModel` struct. It is used to set up the cost
// inputs for the entity model and its child unit models.
func (e *EntityModel) CostMapSetup() {
	for i := range e.CostInput {
		temp := e.CostInput[i]
		temp.GrowthItemOptions = e.GrowthInput
		switch temp.StartEvent {
		case "ModelStart":
			temp.Start = e.StartDate
		}
		switch temp.EndEvent {
		case "ModelEnd":
			temp.End = e.EndDate
		case "SalesDate":
			temp.End = e.SalesDate
		}
		e.CostInput[i] = temp
	}
	for i := range e.GLA.CostInput {
		temp := e.GLA.CostInput[i]
		temp.GrowthItemOptions = e.GrowthInput
		switch temp.StartEvent {
		case "ModelStart":
			temp.Start = e.StartDate
		}
		switch temp.EndEvent {
		case "ModelEnd":
			temp.End = e.EndDate
		case "SalesDate":
			temp.End = e.SalesDate
		}
		e.GLA.CostInput[i] = temp
	}
	for i := range e.ChildUnitModels {
		for ii := range e.ChildUnitModels[i].CostInput {
			temp := e.ChildUnitModels[i].CostInput[ii]
			temp.GrowthItemOptions = e.GrowthInput
			switch temp.StartEvent {
			case "ModelStart":
				temp.Start = e.StartDate
			}
			switch temp.EndEvent {
			case "ModelEnd":
				temp.End = e.EndDate
			case "SalesDate":
				temp.End = e.SalesDate
			}
			e.ChildUnitModels[i].CostInput[ii] = temp
		}
	}
}
