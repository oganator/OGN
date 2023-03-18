package controllers

import (
	"reflect"
)

// CapexCalc - p is the renewal probability and is only used for rent incentives.
func (costMap StringCostInputMap) CostMapCalc(date Datetype, growth StringIntFloatMap, coa FloatCOA) (amount FloatCOA) {
	for _, cost := range costMap {
		if date.Dateint >= cost.Start.Dateint && date.Dateint <= cost.End.Dateint {
			amount.Add(CostCalc(cost, date, growth, coa))
		}
	}
	return amount
}

// CostCalc - p is the renewal probability and is only used for rent incentives.
// TODO - set up StartEvent end EndEvent, probably not within this function though
func CostCalc(cost CostInput, date Datetype, growthInput StringIntFloatMap, coa FloatCOA) (amount FloatCOA) {
	cost.Duration = dateintdiff(cost.End.Dateint, cost.Start.Dateint)
	duration := float64(cost.Duration)
	if cost.Type == "Rent Incentives" {
		duration = 1
	}
	coaAmt := 1.0
	if cost.COAItemBasis != "None" {
		coaAmt = reflect.ValueOf(coa).FieldByName(cost.COAItemBasis).Float()
	}
	growth := 1.0
	if cost.GrowthItem != "None" {
		growth = growthInput[cost.GrowthItem][date.Dateint]
	}
	tempAmount := cost.Amount * growth * coaAmt / duration
	reflect.ValueOf(&amount).Elem().FieldByName(cost.COAItemTarget).SetFloat(-tempAmount)
	return amount
}
