package controllers

// ByIRR -
type ByIRR []*EntityModel

func (a ByIRR) Len() int { return len(a) }
func (a ByIRR) Less(i, j int) bool {
	return a[i].Metrics.IRR.NetLeveredAfterTax < a[j].Metrics.IRR.NetLeveredAfterTax
}
func (a ByIRR) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByIRRr []*EntityModel

func (a ByIRRr) Len() int { return len(a) }
func (a ByIRRr) Less(i, j int) bool {
	return a[i].Metrics.IRR.NetLeveredAfterTax > a[j].Metrics.IRR.NetLeveredAfterTax
}
func (a ByIRRr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByEM -
type ByEM []*EntityModel

func (a ByEM) Len() int { return len(a) }
func (a ByEM) Less(i, j int) bool {
	return a[i].Metrics.EM.NetLeveredAfterTax < a[j].Metrics.EM.NetLeveredAfterTax
}
func (a ByEM) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByEMr []*EntityModel

func (a ByEMr) Len() int { return len(a) }
func (a ByEMr) Less(i, j int) bool {
	return a[i].Metrics.EM.NetLeveredAfterTax > a[j].Metrics.EM.NetLeveredAfterTax
}
func (a ByEMr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByYTM -
type ByYTM []*EntityModel

func (a ByYTM) Len() int { return len(a) }
func (a ByYTM) Less(i, j int) bool {
	return a[i].Metrics.BondHolder.YTM < a[j].Metrics.BondHolder.YTM
}
func (a ByYTM) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByYTMr []*EntityModel

func (a ByYTMr) Len() int { return len(a) }
func (a ByYTMr) Less(i, j int) bool {
	return a[i].Metrics.BondHolder.YTM > a[j].Metrics.BondHolder.YTM
}
func (a ByYTMr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByDuration -
type ByDuration []*EntityModel

func (a ByDuration) Len() int { return len(a) }
func (a ByDuration) Less(i, j int) bool {
	return a[i].Metrics.BondHolder.Duration < a[j].Metrics.BondHolder.Duration
}
func (a ByDuration) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByDurationr []*EntityModel

func (a ByDurationr) Len() int { return len(a) }
func (a ByDurationr) Less(i, j int) bool {
	return a[i].Metrics.BondHolder.Duration > a[j].Metrics.BondHolder.Duration
}
func (a ByDurationr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByVoid -
type ByVoid []*EntityModel

func (a ByVoid) Len() int { return len(a) }
func (a ByVoid) Less(i, j int) bool {
	return a[i].GLA.Void < a[j].GLA.Void
}
func (a ByVoid) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByVoidr []*EntityModel

func (a ByVoidr) Len() int { return len(a) }
func (a ByVoidr) Less(i, j int) bool {
	return a[i].GLA.Void > a[j].GLA.Void
}
func (a ByVoidr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByEXTDUR -
type ByEXTDUR []*EntityModel

func (a ByEXTDUR) Len() int { return len(a) }
func (a ByEXTDUR) Less(i, j int) bool {
	return float64(a[i].GLA.EXTDuration) < float64(a[j].GLA.EXTDuration)
}
func (a ByEXTDUR) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByEXTDURr []*EntityModel

func (a ByEXTDURr) Len() int { return len(a) }
func (a ByEXTDURr) Less(i, j int) bool {
	return float64(a[i].GLA.EXTDuration) > float64(a[j].GLA.EXTDuration)
}
func (a ByEXTDURr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByHazard -
type ByHazard []*EntityModel

func (a ByHazard) Len() int { return len(a) }
func (a ByHazard) Less(i, j int) bool {
	return float64(a[i].GLA.Default.NumberOfDefaults) < float64(a[j].GLA.Default.NumberOfDefaults)
}
func (a ByHazard) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByHazardr []*EntityModel

func (a ByHazardr) Len() int { return len(a) }
func (a ByHazardr) Less(i, j int) bool {
	return float64(a[i].GLA.Default.NumberOfDefaults) > float64(a[j].GLA.Default.NumberOfDefaults)
}
func (a ByHazardr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByOpEx -
type ByOpEx []*EntityModel

func (a ByOpEx) Len() int { return len(a) }
func (a ByOpEx) Less(i, j int) bool {
	return a[i].OpEx.PercentOfTRI < a[j].OpEx.PercentOfTRI
}
func (a ByOpEx) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByOpExr []*EntityModel

func (a ByOpExr) Len() int { return len(a) }
func (a ByOpExr) Less(i, j int) bool {
	return a[i].OpEx.PercentOfTRI > a[j].OpEx.PercentOfTRI
}
func (a ByOpExr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByCPI -
type ByCPI []*EntityModel

func (a ByCPI) Len() int { return len(a) }
func (a ByCPI) Less(i, j int) bool {
	return a[i].GrowthInput["CPI"].EndingIndex < a[j].GrowthInput["CPI"].EndingIndex
}
func (a ByCPI) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByCPIr []*EntityModel

func (a ByCPIr) Len() int { return len(a) }
func (a ByCPIr) Less(i, j int) bool {
	return a[i].GrowthInput["CPI"].EndingIndex > a[j].GrowthInput["CPI"].EndingIndex
}
func (a ByCPIr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByERV -
type ByERV []*EntityModel

func (a ByERV) Len() int { return len(a) }
func (a ByERV) Less(i, j int) bool {
	return a[i].GrowthInput["ERV"].EndingIndex < a[j].GrowthInput["ERV"].EndingIndex
}
func (a ByERV) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByERVr []*EntityModel

func (a ByERVr) Len() int { return len(a) }
func (a ByERVr) Less(i, j int) bool {
	return a[i].GrowthInput["ERV"].EndingIndex > a[j].GrowthInput["ERV"].EndingIndex
}
func (a ByERVr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByYieldShift -
type ByYieldShift []*EntityModel

func (a ByYieldShift) Len() int { return len(a) }
func (a ByYieldShift) Less(i, j int) bool {
	return a[i].Valuation.YieldShift < a[j].Valuation.YieldShift
}
func (a ByYieldShift) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByYieldShiftr []*EntityModel

func (a ByYieldShiftr) Len() int { return len(a) }
func (a ByYieldShiftr) Less(i, j int) bool {
	return a[i].Valuation.YieldShift > a[j].Valuation.YieldShift
}
func (a ByYieldShiftr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByYTM -
type ByEndCash []*EntityModel

func (a ByEndCash) Len() int { return len(a) }
func (a ByEndCash) Less(i, j int) bool {
	return a[i].COA[a[i].SalesDate.Dateint].CashBalance < a[j].COA[a[j].SalesDate.Dateint].CashBalance
}
func (a ByEndCash) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByEndCashr []*EntityModel

func (a ByEndCashr) Len() int { return len(a) }
func (a ByEndCashr) Less(i, j int) bool {
	return a[i].COA[a[i].SalesDate.Dateint].CashBalance > a[j].COA[a[j].SalesDate.Dateint].CashBalance
}
func (a ByEndCashr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByEndNCF -
type ByEndNCF []*EntityModel

func (a ByEndNCF) Len() int { return len(a) }
func (a ByEndNCF) Less(i, j int) bool {
	return a[i].COA[a[i].SalesDate.Dateint].NetCashFlow < a[j].COA[a[j].SalesDate.Dateint].NetCashFlow
}
func (a ByEndNCF) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByEndNCFr []*EntityModel

func (a ByEndNCFr) Len() int { return len(a) }
func (a ByEndNCFr) Less(i, j int) bool {
	return a[i].COA[a[i].SalesDate.Dateint].NetCashFlow > a[j].COA[a[j].SalesDate.Dateint].NetCashFlow
}
func (a ByEndNCFr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByEndMarketValue -
type ByEndMarketValue []*EntityModel

func (a ByEndMarketValue) Len() int { return len(a) }
func (a ByEndMarketValue) Less(i, j int) bool {
	return a[i].COA[a[i].SalesDate.Dateint].MarketValue < a[j].COA[a[j].SalesDate.Dateint].MarketValue
}
func (a ByEndMarketValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByEndMarketValuer []*EntityModel

func (a ByEndMarketValuer) Len() int { return len(a) }
func (a ByEndMarketValuer) Less(i, j int) bool {
	return a[i].COA[a[i].SalesDate.Dateint].MarketValue > a[j].COA[a[j].SalesDate.Dateint].MarketValue
}
func (a ByEndMarketValuer) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByProb -
type ByProb []*EntityModel

func (a ByProb) Len() int { return len(a) }
func (a ByProb) Less(i, j int) bool {
	return a[i].GLA.Probability < a[j].GLA.Probability
}
func (a ByProb) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByProbr []*EntityModel

func (a ByProbr) Len() int { return len(a) }
func (a ByProbr) Less(i, j int) bool {
	return a[i].GLA.Probability > a[j].GLA.Probability
}
func (a ByProbr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
