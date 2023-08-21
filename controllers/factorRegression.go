package controllers

import (
	"gonum.org/v1/gonum/stat"
)

type RegressionStats struct {
	Alpha, Beta, Corr, Rsq float64
}

func (ab *RegressionStats) RegressionCalc(x, y *[]float64) {
	ab.Corr = stat.Correlation(*x, *y, nil)
	ab.Alpha, ab.Beta = stat.LinearRegression(*x, *y, nil, false)
	ab.Rsq = stat.RSquared(*x, *y, nil, ab.Alpha, ab.Beta)
}

type FactorIndependant struct {
	Metric                                                                  string
	Void, Probability, NumberOfDefaults, OpEx, CPI, ERV, Hazard, YieldShift RegressionStats
}

func FactorRegression(metric string, mcresults *MCResultSlice, y *[]float64) (fi FactorIndependant) {
	fi.Metric = metric
	fi.Void.RegressionCalc(&mcresults.Void, y)
	fi.Probability.RegressionCalc(&mcresults.Probability, y)
	fi.NumberOfDefaults.RegressionCalc(&mcresults.NumberOfDefaults, y)
	fi.OpEx.RegressionCalc(&mcresults.OpEx, y)
	fi.CPI.RegressionCalc(&mcresults.CPI, y)
	fi.ERV.RegressionCalc(&mcresults.ERV, y)
	fi.Hazard.RegressionCalc(&mcresults.Hazard, y)
	fi.YieldShift.RegressionCalc(&mcresults.YieldShift, y)
	return fi
}

func (e *EntityModel) FactorAnalysisCalc(mcresults *MCResultSlice) {
	if e.Strategy == "Standard" {
		e.FactorAnalysis = make([]FactorIndependant, 5)
		e.FactorAnalysis[0] = FactorRegression("IRR", mcresults, &mcresults.IRR)
		e.FactorAnalysis[1] = FactorRegression("EM", mcresults, &mcresults.EM)
		e.FactorAnalysis[2] = FactorRegression("EndCash", mcresults, &mcresults.EndCash)
		e.FactorAnalysis[3] = FactorRegression("EndNCF", mcresults, &mcresults.EndNCF)
		e.FactorAnalysis[4] = FactorRegression("EndMV", mcresults, &mcresults.EndMarketValue)
	} else {
		e.FactorAnalysis = make([]FactorIndependant, 6)
		e.FactorAnalysis[0] = FactorRegression("YTM", mcresults, &mcresults.YTM)
		e.FactorAnalysis[1] = FactorRegression("Duration", mcresults, &mcresults.Duration)
		e.FactorAnalysis[2] = FactorRegression("YTM/DUR", mcresults, &mcresults.YTMDUR)
		e.FactorAnalysis[3] = FactorRegression("EndCash", mcresults, &mcresults.EndCash)
		e.FactorAnalysis[4] = FactorRegression("EndNCF", mcresults, &mcresults.EndNCF)
		e.FactorAnalysis[5] = FactorRegression("EndMV", mcresults, &mcresults.EndMarketValue)
		irr := e.Metrics.IRR.NetLeveredAfterTax
		if irr > 0.0 {
			e.FactorAnalysis = append(e.FactorAnalysis, FactorRegression("IRR", mcresults, &mcresults.IRR))
		}
		em := e.Metrics.EM.NetLeveredAfterTax
		if em > 0.0 {
			e.FactorAnalysis = append(e.FactorAnalysis, FactorRegression("EM", mcresults, &mcresults.EM))
		}
	}
}
