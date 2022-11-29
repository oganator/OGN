package controllers

import (
	"math"
)

// Tax -
type Tax struct {
	MinValue        float64             `json:"MinValue,omitempty"`        // as a percent
	LandValue       float64             `json:"LandValue,omitempty"`       // as a percent
	UsablePeriod    int                 `json:"UsablePeriod,omitempty"`    // number of years
	RETT            float64             `json:"RETT,omitempty"`            // as a percent
	CIT             map[float64]float64 `json:"CIT,omitempty"`             //map[income]rate
	VAT             float64             `json:"VAT,omitempty"`             //
	CarryBackYrs    int                 `json:"CarryBackYrs,omitempty"`    //
	CarryForwardYrs int                 `json:"CarryForwardYrs,omitempty"` //
	DTA             IntFloatMap         `json:"DTA,omitempty"`             //
}

// CIT - starts in december of the second year, and assesses the tax liability of that year, then inserts the taxes payable into may of the next year.
// Does not assess capital gain - done in Disposal().
// As more items are built out, they need to be added to taxable income, as the inclusion of BPuplift previously skewed the taxable income upwards
func (e *EntityModel) CIT() {
	startdate := Dateadd(Datetype{Month: 5, Year: e.StartDate.Year}, 0)
	taxyear := e.StartDate.Year - 1
	depr := Depreciation(e)
	for date := Dateadd(startdate, 0); date.Dateint <= e.EndDate.Dateint; date.Add(12) {
		taxespayable := 0.0
		// taxableincome := e.COA[taxyear].TheoreticalRentalIncome + e.COA[taxyear].Vacancy + e.COA[taxyear].Capex + e.COA[taxyear].InterestExpense + depr + e.COA[taxyear].Fees
		// fmt.Println(date.Dateint, e.COA[taxyear].NetOperatingIncome, e.COA[taxyear].Capex, e.COA[taxyear].InterestExpense, depr, e.COA[taxyear].Fees)
		taxableincome := e.COA[taxyear].NetOperatingIncome + e.COA[taxyear].Capex + e.COA[taxyear].InterestExpense + depr //+ e.COA[taxyear].Fees
		// capital gain
		if date.Dateint == Dateadd(e.SalesDate, 5).Dateint {
			taxableincome = taxableincome + e.COA[e.SalesDate.Dateint].AcqDispProperty + e.COA[Dateadd(e.StartDate, -1).Dateint].AcqDispProperty
		}
		dta := 0.0
		taxableincomecarryback := 0.0
		if taxableincome < 0.0 {
			if date != startdate {
				taxespayable = CarryBack(taxableincome, e, date)
				e.Tax.DTA[date.Dateint] = math.Min(taxableincome+taxespayable, 0.0) // store DTA
				//
				temp := e.COA[date.Dateint]
				temp.DTA = e.Tax.DTA[date.Dateint]
				e.COA[date.Dateint] = temp
				//
				dta = -math.Min(taxableincome-taxespayable, 0.0)
			}
			taxableincomecarryback = 0.0
		} else {
			taxableincome = CarryForward(taxableincome, e, date)
			taxespayable = taxableincome * 0.21
			taxableincomecarryback = taxableincome
		}
		// } else if taxableincome > 395000.0 {
		// 	taxespayable = 395000.0 * 0.15 + (taxableincome-395000) * 0.217
		// } else {
		// 	taxespayable = taxableincome * 0.15
		// }
		if date == startdate {
			taxableincome = 0.0
		}
		temp := FloatCOA{Tax: -taxespayable, Depreciation: depr, TaxableIncome: taxableincome, TaxableIncomeCarryBack: taxableincomecarryback, DTA: dta}
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
func Depreciation(e *EntityModel) float64 {
	// return (-e.Valuation.AcqPrice - (-e.Valuation.AcqPrice * e.Tax.MinValue) - (-e.Valuation.AcqPrice * e.Tax.LandValue)) / float64(e.Tax.UsablePeriod)
	return (-e.Valuation.AcqPrice * (1 - e.Tax.MinValue) * (1 - e.Tax.LandValue)) / float64(e.Tax.UsablePeriod)
}

func CarryBack(startincome float64, e *EntityModel, date Datetype) (refund float64) {
	// fmt.Println("tax year: ", date.Year)
	for prevdate := Dateadd(date, e.Tax.CarryBackYrs*-12); prevdate.Dateint < date.Dateint; prevdate.Add(12) {
		if e.COA[prevdate.Dateint].TaxableIncomeCarryBack <= 0.0 {
			// fmt.Println("break")
			continue
		}
		updatedtaxableincome := math.Max(0.0, startincome+e.COA[prevdate.Dateint].TaxableIncomeCarryBack)
		updatedtaxespayable := updatedtaxableincome * .21
		refund = refund - e.COA[prevdate.Dateint].Tax - updatedtaxespayable
		startincome = math.Min(startincome+e.COA[prevdate.Dateint].TaxableIncome, 0.0)
		remainingincome := math.Max(0.0, e.COA[prevdate.Dateint].TaxableIncome+startincome)
		//
		temp := e.COA[prevdate.Dateint]
		temp.TaxableIncomeCarryBack = -remainingincome
		e.COA[prevdate.Dateint] = temp
		//
		// fmt.Println("     cf date: ", prevdate.Dateint, "     updatedtaxableincome: ", updatedtaxableincome, "     updatedtaxespayable: ", updatedtaxespayable, "     refund: ", refund, "     remainingincome: ", remainingincome, "     taxincomecarryback: ", e.COA[prevdate.Dateint].TaxableIncomeCarryBack)
		if startincome == 0.0 {
			break
		}
	}
	return -refund
}

func CarryForward(startincome float64, e *EntityModel, date Datetype) (resultincome float64) {
	// fmt.Println("tax year: ", date.Year)
	resultincome = startincome
	tempincome := resultincome
	for prevdate := Dateadd(date, e.Tax.CarryForwardYrs*-12); prevdate.Dateint < date.Dateint; prevdate.Add(12) {
		startdta := e.Tax.DTA[prevdate.Dateint]
		resultincome = math.Max(resultincome+e.Tax.DTA[prevdate.Dateint], 0.0)
		e.Tax.DTA[prevdate.Dateint] = math.Min(0.0, tempincome+startdta)
		//
		temp := e.COA[prevdate.Dateint]
		temp.DTA = e.Tax.DTA[prevdate.Dateint]
		e.COA[prevdate.Dateint] = temp
		//
		// fmt.Println("     cf date: ", prevdate.Dateint, "     resultincome: ", resultincome, "     StartDTA: ", startdta, "     EndDTA: ", e.Tax.DTA[prevdate.Dateint])
		if resultincome == 0.0 {
			// fmt.Println("break")
			break
		}
		tempincome = resultincome
	}
	return resultincome
}
