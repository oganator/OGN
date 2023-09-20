package controllers

import (
	"math"
)

// DebtInput -
type DebtInput struct {
	MasterID           int            `json:"MasterID,omitempty"`           //
	Name               string         `json:"Name,omitempty"`               //
	LTV                float64        `json:"LTV,omitempty"`                //
	InterestRate       float64        `json:"InterestRate,omitempty"`       //
	TempRate           float64        `json:"TempRate,omitempty"`           // used for floating rate calcs
	InterestType       string         `json:"InterestType,omitempty"`       // Fixed, Floating
	LoanType           string         `json:"LoanType,omitempty"`           // Interest Only, Amortizing
	LoanBasis          string         `json:"LoanBasis,omitempty"`          // Amount, Market Value, Capex, Revaluation
	LoanStart          Datetype       `json:"LoanStart,omitempty"`          //
	LoanEnd            Datetype       `json:"LoanEnd,omitempty"`            //
	LastIndex          Datetype       `json:"LastIndex,omitempty"`          //
	Amount             float64        `json:"Amount,omitempty"`             //
	FloatBasis         string         `json:"FloatBasis,omitempty"`         // LIBOR, CPI
	Spread             int            `json:"Spread,omitempty"`             // in bps
	AmortizationPeriod int            `json:"AmortizationPeriod,omitempty"` // in months
	COA                IntFloatCOAMap `json:"COA,omitempty"`                // each loan needs to have its own activity tracked, balance, principal repayments
	Active             bool           `json:"Active,omitempty"`             //
	StartEvent         string         `json:"StartEvent,omitempty"`
	EndEvent           string         `json:"EndEvent,omitempty"`
}

// DebtCalc - Ranges over the Loans, and calculates debt COA values for a particular period, including the setup of the loan.
// The resulting 'finalCOA' is the aggregate sum of all loan activity. Each loan will store its own activity in its own .COA field.
func (e *EntityModel) DebtCalc(period Datetype) (finalCOA FloatCOA) {
	for index, loan := range e.DebtInput {
		if loan.Active {
			// get the correct monthly interest rate
			rate := 0.0
			switch loan.InterestType {
			case "Fixed":
				rate = loan.InterestRate / 12
			case "Floating":
				if period.Dateint == Dateadd(loan.LastIndex, 12).Dateint {
					e.DebtInput[index].TempRate = (e.Growth[loan.FloatBasis][period.Dateint] / e.Growth[loan.FloatBasis][loan.LastIndex.Dateint]) - 1
					e.DebtInput[index].LastIndex = period
				}
				rate = (loan.TempRate/12 + float64(loan.Spread)/120000)
			}
			innerCOA := FloatCOA{}
			// Initial loan setup
			if period.Dateint == Dateadd(loan.LoanStart, -1).Dateint {
				innerCOA.LoanBalance = loan.Amount
				innerCOA.LoanProceeds = loan.Amount
				innerCOA.Debt = loan.Amount
			}
			// periodic loan activity
			if period.Dateint >= loan.LoanStart.Dateint && period.Dateint < loan.LoanEnd.Dateint {
				// interest only vs amortizing loans
				switch loan.LoanType {
				case "Interest Only":
					innerCOA.InterestExpense = rate * loan.COA[Dateadd(e.StartDate, -1).Dateint].LoanBalance
					// TODO - include manual principal repayments
					innerCOA.LoanBalance = loan.COA[Dateadd(period, -1).Dateint].LoanBalance + loan.COA[period.Dateint].PrincipalRepayment
					innerCOA.Debt = -(innerCOA.InterestExpense + innerCOA.PrincipalRepayment)
				case "Amortizing":
					priorLoanBalance := loan.COA[Dateadd(period, -1).Dateint].LoanBalance
					innerCOA.InterestExpense = priorLoanBalance * rate
					one_plus_I_to_the_N := math.Pow(1+rate, float64(loan.AmortizationPeriod*12))
					debt := (loan.Amount * rate * one_plus_I_to_the_N) / (one_plus_I_to_the_N - 1)
					innerCOA.Debt = -debt
					innerCOA.PrincipalRepayment = debt - (priorLoanBalance * rate)
					innerCOA.LoanBalance = priorLoanBalance - innerCOA.PrincipalRepayment
				}
			}
			// Loan termination
			if period.Dateint == loan.LoanEnd.Dateint {
				switch loan.LoanType {
				case "Interest Only":
					innerCOA.InterestExpense = rate * loan.COA[Dateadd(e.StartDate, -1).Dateint].LoanBalance
					innerCOA.PrincipalRepayment = loan.Amount
					innerCOA.Debt = -(innerCOA.InterestExpense + innerCOA.PrincipalRepayment)
				case "Amortizing":
					priorLoanBalance := loan.COA[Dateadd(period, -1).Dateint].LoanBalance
					innerCOA.InterestExpense = priorLoanBalance * rate
					one_plus_I_to_the_N := math.Pow(1+rate, float64(loan.AmortizationPeriod*12))
					// innerCOA.PrincipalRepayment = debt - (priorLoanBalance * rate)
					innerCOA.PrincipalRepayment = priorLoanBalance
					debt := (loan.Amount*rate*one_plus_I_to_the_N)/(one_plus_I_to_the_N-1) + innerCOA.PrincipalRepayment
					innerCOA.Debt = -debt
					innerCOA.LoanBalance = priorLoanBalance - innerCOA.PrincipalRepayment
				}
			}
			loan.COA[period.Dateint] = innerCOA
			finalCOA.Add(innerCOA)
		}
	}
	return finalCOA
}

// sets up each loan so it doesnt blow up - runs Dateadd to dates, makes a IntFloatCOAmap for coa, and sets up SensitivityInput.Debt
func (e *EntityModel) DebtSetup() {
	e.SensitivityInput.Debt = make([]DebtInput, len(e.DebtInput))
	for i := range e.DebtInput {
		e.DebtInput[i].TempRate = e.DebtInput[i].InterestRate
		e.DebtInput[i].COA = make(IntFloatCOAMap)
		e.DebtInput[i].LoanStart = Dateadd(e.DebtInput[i].LoanStart, 0)
		e.DebtInput[i].LoanEnd = Dateadd(e.DebtInput[i].LoanEnd, 0)
		e.DebtInput[i].LastIndex = Dateadd(e.DebtInput[i].LoanStart, 0)
		e.SensitivityInput.Debt[i] = DebtInput{Spread: 0}
	}
}
