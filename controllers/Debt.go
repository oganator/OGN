package controllers

import (
	"math"
)

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
