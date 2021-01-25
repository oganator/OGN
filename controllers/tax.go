package controllers

import (
	"math"
)

// Tax -
type Tax struct {
	MinValue        float64             // as a percent
	LandValue       float64             // as a percent
	UsablePeriod    int                 // number of years
	RETT            float64             // as a percent
	CIT             map[float64]float64 //map[income]rate
	VAT             float64             //
	CarryBackYrs    int                 //
	CarryForwardYrs int                 //
}

// CIT - Also fees. 15%/21.7% starts in december of the first year, and assesses the tax liability of that year, then inserts the taxes payable into may of the next year. Does not assess capital gain (done in Disposal()). As more items are built out, they need to be added to taxable income, as the inclusion of BPuplift previously skewed the taxable income upwards
func (e *Entity) CIT() {
	startdate := Dateadd(Datetype{Month: 5, Year: e.StartDate.Year}, 0)
	taxyear := e.StartDate.Year - 1
	depr := Depreciation(e)
	for date := startdate; date.Dateint <= e.EndDate.Dateint; date.Add(12) {
		taxespayable := 0.0
		taxableincome := e.COA[taxyear].TheoreticalRentalIncome + e.COA[taxyear].Vacancy + e.COA[taxyear].Capex + e.COA[taxyear].InterestExpense + depr + e.COA[taxyear].Fees
		// capital gain
		if date.Dateint == Dateadd(e.SalesDate, 5).Dateint {
			taxableincome = taxableincome + e.COA[e.SalesDate.Dateint].AcqDispProperty + e.COA[Dateadd(e.StartDate, -1).Dateint].AcqDispProperty
		}
		if taxableincome < 0.0 {
			taxespayable = CarryBack(taxableincome, e, date)
		} else if taxableincome > 200000.0 {
			taxespayable = 200000.0*.15 + (taxableincome-200000)*.217
		} else {
			taxespayable = taxableincome * .15
		}
		temp := FloatCOA{Tax: -taxespayable}
		if date.Dateint == Dateadd(e.SalesDate, 5).Dateint {
			temp.Add(e.COA[e.SalesDate.Dateint])
			e.COA[e.SalesDate.Dateint] = temp
			break
		}
		temp.Add(e.COA[date.Dateint])
		e.COA[date.Dateint] = temp
		taxyear++
	}
}

// Depreciation -
func Depreciation(e *Entity) float64 {
	depr := (-e.Valuation.AcqPrice - (-e.Valuation.AcqPrice * e.Tax.MinValue) - (-e.Valuation.AcqPrice * e.Tax.LandValue)) / float64(e.Tax.UsablePeriod)
	return depr
}

// CarryBack -
func CarryBack(nol float64, e *Entity, date Datetype) (cb float64) {
	tax := 0.0
	for i := 1; i <= e.Tax.CarryBackYrs; i++ {
		tax = tax + e.COA[Dateadd(date, -12*i).Dateint].Tax
	}
	if nol < 1000000 {
		cb = math.Max(nol, tax)
	} else {
		cb = math.Max(nol+1000000, tax) + math.Min(math.Max(nol+1000000, tax+1000000)*.5, 0)
	}
	return cb
}
