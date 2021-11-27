package controllers

import (
	"fmt"
	"strconv"
	"sync"

	xl "github.com/xuri/excelize/v2"
)

var XLSX, _ = xl.OpenFile("./models/Data.xlsx")

// EntityData -
type EntityData struct {
	Mutex               *sync.Mutex
	MasterID            int
	Name                string
	Parent              int // MasterID
	StartMonth          int
	StartYear           int
	SalesMonth          int
	SalesYear           int
	EndMonth            int
	EndYear             int
	CPIGrowth           HModel
	CPISigma            HModel
	ERVGrowth           HModel
	ERVSigma            HModel
	EntryYield          float64
	YieldShift          float64
	YieldShiftSigma     float64
	ExitYield           float64
	HoldPeriod          int
	RETT                float64
	VAT                 float64
	WOZpercent          float64
	DeprPeriod          int
	Landvalue           float64
	CarryBackYrs        int
	CarryForwardYrs     int
	LTV                 float64
	LoanRate            float64
	OpExpercent         float64
	Fees                float64
	percentIncometosell float64
	YearsIncometosell   int
	DiscountRate        float64
	Strategy            string
	BalloonPercent      float64
	GLA                 Unit
	VoidSigma           float64
	ProbabilitySigma    float64
	OpExSigma           float64
	Hazard              float64
}

// UnitData -
type UnitData struct {
	MasterID        int
	Name            string
	ParentMasterID  int
	LeaseStartMonth int
	LeaseStartYear  int
	LeaseEndMonth   int
	LeaseEndYear    int
	UnitStatus      string
	Tenant          string
	PassingRent     float64
	Probability     float64
	RentRevisionERV float64
	EXTDuration     int
	IndexFreq       int
	IndexType       string
	IndexStartMonth int
	Void            int
	DiscountRate    float64
	ERVArea         float64
	ERVAmount       float64
	PercentSoldRent float64
}

// GrowthData -
type GrowthData struct {
	EntityMasterID int
	Item           string
	Amount         float64
}

// WriteXLSX -
func WriteXLSXEntities(e *Entity) {
	row, _ := XLSX.SearchSheet("Entities", e.Name)
	rows := row[0]
	rows = rows[1:]
	XLSX.SetCellValue("Entities", "K"+fmt.Sprint(rows), e.GrowthInput["CPI"].ShortTermRate)
	XLSX.SetCellValue("Entities", "L"+fmt.Sprint(rows), e.GrowthInput["CPI"].ShortTermPeriod)
	XLSX.SetCellValue("Entities", "M"+fmt.Sprint(rows), e.GrowthInput["CPI"].TransitionPeriod)
	XLSX.SetCellValue("Entities", "N"+fmt.Sprint(rows), e.GrowthInput["CPI"].LongTermRate)
	XLSX.SetCellValue("Entities", "O"+fmt.Sprint(rows), e.GrowthInput["ERV"].ShortTermRate)
	XLSX.SetCellValue("Entities", "P"+fmt.Sprint(rows), e.GrowthInput["ERV"].ShortTermPeriod)
	XLSX.SetCellValue("Entities", "Q"+fmt.Sprint(rows), e.GrowthInput["ERV"].TransitionPeriod)
	XLSX.SetCellValue("Entities", "R"+fmt.Sprint(rows), e.GrowthInput["ERV"].LongTermRate)
	XLSX.SetCellValue("Entities", "S"+fmt.Sprint(rows), e.Valuation.EntryYield)
	XLSX.SetCellValue("Entities", "T"+fmt.Sprint(rows), e.Valuation.YieldShift)
	XLSX.SetCellValue("Entities", "V"+fmt.Sprint(rows), e.HoldPeriod)
	XLSX.SetCellValue("Entities", "W"+fmt.Sprint(rows), e.Tax.RETT)
	XLSX.SetCellValue("Entities", "X"+fmt.Sprint(rows), e.Tax.VAT)
	XLSX.SetCellValue("Entities", "Y"+fmt.Sprint(rows), e.Tax.MinValue)
	XLSX.SetCellValue("Entities", "Z"+fmt.Sprint(rows), e.Tax.UsablePeriod)
	XLSX.SetCellValue("Entities", "AA"+fmt.Sprint(rows), e.Tax.LandValue)
	XLSX.SetCellValue("Entities", "AB"+fmt.Sprint(rows), e.DebtInput.LTV)
	XLSX.SetCellValue("Entities", "AC"+fmt.Sprint(rows), e.DebtInput.InterestRate)
	XLSX.SetCellValue("Entities", "AD"+fmt.Sprint(rows), e.OpEx.PercentOfTRI)
	XLSX.SetCellValue("Entities", "AE"+fmt.Sprint(rows), e.Tax.CarryBackYrs)
	XLSX.SetCellValue("Entities", "AF"+fmt.Sprint(rows), e.Tax.CarryForwardYrs)
	XLSX.SetCellValue("Entities", "AG"+fmt.Sprint(rows), e.GLA.PercentSoldRent)
	XLSX.SetCellValue("Entities", "AH"+fmt.Sprint(rows), e.Strategy)
	XLSX.SetCellValue("Entities", "AI"+fmt.Sprint(rows), e.Fees.PercentOfGAV)
	XLSX.SetCellValue("Entities", "AJ"+fmt.Sprint(rows), e.GLA.DiscountRate)
	XLSX.SetCellValue("Entities", "AK"+fmt.Sprint(rows), e.GLA.Void)
	XLSX.SetCellValue("Entities", "AL"+fmt.Sprint(rows), e.GLA.EXTDuration)
	XLSX.SetCellValue("Entities", "AM"+fmt.Sprint(rows), e.GLA.RentRevisionERV)
	XLSX.SetCellValue("Entities", "AN"+fmt.Sprint(rows), e.GLA.Probability)
	XLSX.SetCellValue("Entities", "AO"+fmt.Sprint(rows), e.BalloonPercent)
	XLSX.SetCellValue("Entities", "AP"+fmt.Sprint(rows), e.MCSetup.YieldShift)
	XLSX.SetCellValue("Entities", "AQ"+fmt.Sprint(rows), e.MCSetup.Void)
	XLSX.SetCellValue("Entities", "AR"+fmt.Sprint(rows), e.MCSetup.Probability)
	XLSX.SetCellValue("Entities", "AS"+fmt.Sprint(rows), e.MCSetup.OpEx)
	XLSX.SetCellValue("Entities", "AT"+fmt.Sprint(rows), e.MCSetup.CPI.ShortTermRate)
	XLSX.SetCellValue("Entities", "AU"+fmt.Sprint(rows), e.MCSetup.CPI.ShortTermPeriod)
	XLSX.SetCellValue("Entities", "AV"+fmt.Sprint(rows), e.MCSetup.CPI.TransitionPeriod)
	XLSX.SetCellValue("Entities", "AW"+fmt.Sprint(rows), e.MCSetup.CPI.LongTermRate)
	XLSX.SetCellValue("Entities", "AX"+fmt.Sprint(rows), e.MCSetup.ERV.ShortTermRate)
	XLSX.SetCellValue("Entities", "AY"+fmt.Sprint(rows), e.MCSetup.ERV.ShortTermPeriod)
	XLSX.SetCellValue("Entities", "AZ"+fmt.Sprint(rows), e.MCSetup.ERV.TransitionPeriod)
	XLSX.SetCellValue("Entities", "BA"+fmt.Sprint(rows), e.MCSetup.ERV.LongTermRate)
	XLSX.SetCellValue("Entities", "BB"+fmt.Sprint(rows), e.GLA.Default.Hazard)
	XLSX.SetCellValue("Entities", "BC"+fmt.Sprint(rows), e.GLA.RentIncentives.Duration)
	XLSX.SetCellValue("Entities", "BD"+fmt.Sprint(rows), e.GLA.RentIncentives.PercentOfContractRent)
	XLSX.SetCellValue("Entities", "BF"+fmt.Sprint(rows), e.GLA.FitOutCosts.AmountPerTotalArea)
	XLSX.Save()
}

// WriteXLSXUnits -
func (u *UnitData) WriteXLSXUnits() {
	if u.MasterID == 0 {
		u.MasterID = len(UnitStore)
	}
	row := u.MasterID + 2
	XLSX.SetCellValue("Units", "B"+fmt.Sprint(row), u.MasterID)
	XLSX.SetCellValue("Units", "C"+fmt.Sprint(row), u.Name)
	XLSX.SetCellValue("Units", "D"+fmt.Sprint(row), u.ParentMasterID)
	XLSX.SetCellValue("Units", "E"+fmt.Sprint(row), u.LeaseStartMonth)
	XLSX.SetCellValue("Units", "F"+fmt.Sprint(row), u.LeaseStartYear)
	XLSX.SetCellValue("Units", "G"+fmt.Sprint(row), u.LeaseEndMonth)
	XLSX.SetCellValue("Units", "H"+fmt.Sprint(row), u.LeaseEndYear)
	XLSX.SetCellValue("Units", "I"+fmt.Sprint(row), u.UnitStatus)
	XLSX.SetCellValue("Units", "J"+fmt.Sprint(row), u.Tenant)
	XLSX.SetCellValue("Units", "K"+fmt.Sprint(row), u.PassingRent)
	XLSX.SetCellValue("Units", "V"+fmt.Sprint(row), u.ERVArea)
	XLSX.SetCellValue("Units", "W"+fmt.Sprint(row), u.ERVAmount)
	XLSX.Save()
}

// ReadXLSX - Reads Data.xlsx and populates data stores (Entity, Unit and GrowthItems)
func ReadXLSX() {
	// ENTITIES
	entities, _ := XLSX.GetRows("Entities")
	for i, row := range entities {
		if i < 2 {
			continue
		}
		tempentity := EntityData{}
		tempentity.Mutex = &sync.Mutex{}
		tempentity.MasterID, _ = strconv.Atoi(row[1])
		tempentity.Name = row[2]
		tempentity.Parent, _ = strconv.Atoi(row[3])
		tempentity.StartMonth, _ = strconv.Atoi(row[4])
		tempentity.StartYear, _ = strconv.Atoi(row[5])
		tempentity.SalesMonth, _ = strconv.Atoi(row[6])
		tempentity.SalesYear, _ = strconv.Atoi(row[7])
		tempentity.EndMonth, _ = strconv.Atoi(row[8])
		tempentity.EndYear, _ = strconv.Atoi(row[9])
		tempentity.CPIGrowth.ShortTermRate, _ = strconv.ParseFloat(row[10], 64)
		tempentity.CPIGrowth.ShortTermPeriod, _ = strconv.Atoi(row[11])
		tempentity.CPIGrowth.TransitionPeriod, _ = strconv.Atoi(row[12])
		tempentity.CPIGrowth.LongTermRate, _ = strconv.ParseFloat(row[13], 64)
		tempentity.ERVGrowth.ShortTermRate, _ = strconv.ParseFloat(row[14], 64)
		tempentity.ERVGrowth.ShortTermPeriod, _ = strconv.Atoi(row[15])
		tempentity.ERVGrowth.TransitionPeriod, _ = strconv.Atoi(row[16])
		tempentity.ERVGrowth.LongTermRate, _ = strconv.ParseFloat(row[17], 64)
		tempentity.EntryYield, _ = strconv.ParseFloat(row[18], 64)
		tempentity.YieldShift, _ = strconv.ParseFloat(row[19], 64)
		tempentity.ExitYield, _ = strconv.ParseFloat(row[20], 64)
		tempentity.HoldPeriod, _ = strconv.Atoi(row[21])
		tempentity.RETT, _ = strconv.ParseFloat(row[22], 64)
		tempentity.VAT, _ = strconv.ParseFloat(row[23], 64)
		tempentity.WOZpercent, _ = strconv.ParseFloat(row[24], 64)
		tempentity.DeprPeriod, _ = strconv.Atoi(row[25])
		tempentity.Landvalue, _ = strconv.ParseFloat(row[26], 64)
		tempentity.CarryBackYrs, _ = strconv.Atoi(row[30])
		tempentity.CarryForwardYrs, _ = strconv.Atoi(row[31])
		tempentity.LTV, _ = strconv.ParseFloat(row[27], 64)
		tempentity.LoanRate, _ = strconv.ParseFloat(row[28], 64)
		tempentity.OpExpercent, _ = strconv.ParseFloat(row[29], 64)
		tempentity.percentIncometosell, _ = strconv.ParseFloat(row[32], 64)
		tempentity.DiscountRate, _ = strconv.ParseFloat(row[35], 64)
		tempentity.Strategy = row[33]
		tempentity.Fees, _ = strconv.ParseFloat(row[34], 64)
		tempentity.GLA.DiscountRate, _ = strconv.ParseFloat(row[35], 64)
		tempentity.GLA.Void, _ = strconv.Atoi(row[36])
		tempentity.GLA.EXTDuration, _ = strconv.Atoi(row[37])
		tempentity.GLA.RentRevisionERV, _ = strconv.ParseFloat(row[38], 64)
		tempentity.GLA.Probability, _ = strconv.ParseFloat(row[39], 64)
		tempentity.BalloonPercent, _ = strconv.ParseFloat(row[40], 64)
		tempentity.YieldShiftSigma, _ = strconv.ParseFloat(row[41], 64)
		tempentity.VoidSigma, _ = strconv.ParseFloat(row[42], 64)
		tempentity.ProbabilitySigma, _ = strconv.ParseFloat(row[43], 64)
		tempentity.OpExSigma, _ = strconv.ParseFloat(row[44], 64)
		tempentity.CPISigma.ShortTermRate, _ = strconv.ParseFloat(row[45], 64)
		tempentity.CPISigma.ShortTermPeriod, _ = strconv.Atoi(row[46])
		tempentity.CPISigma.TransitionPeriod, _ = strconv.Atoi(row[47])
		tempentity.CPISigma.LongTermRate, _ = strconv.ParseFloat(row[48], 64)
		tempentity.ERVSigma.ShortTermRate, _ = strconv.ParseFloat(row[49], 64)
		tempentity.ERVSigma.ShortTermPeriod, _ = strconv.Atoi(row[50])
		tempentity.ERVSigma.TransitionPeriod, _ = strconv.Atoi(row[51])
		tempentity.ERVSigma.LongTermRate, _ = strconv.ParseFloat(row[52], 64)
		tempentity.GLA.Default.Hazard, _ = strconv.ParseFloat(row[53], 64)
		tempentity.GLA.RentIncentives.Duration, _ = strconv.Atoi(row[54])
		tempentity.GLA.RentIncentives.PercentOfContractRent, _ = strconv.ParseFloat(row[55], 64)
		tempentity.GLA.FitOutCosts.AmountPerTotalArea, _ = strconv.ParseFloat(row[57], 64)
		EntityDataStore[tempentity.MasterID] = &tempentity
	}

	// UNITS
	units, _ := XLSX.GetRows("Units")
	// unitfile, _ := os.Open("./models/Units.csv")
	// units, _ := csv.NewReader(unitfile).ReadAll()
	for i, row := range units {
		if i < 2 {
			continue
		}
		tempunit := UnitData{}
		tempunit.MasterID, _ = strconv.Atoi(row[1])
		tempunit.Name = row[2]
		tempunit.ParentMasterID, _ = strconv.Atoi(row[3])
		tempunit.LeaseStartMonth, _ = strconv.Atoi(row[4])
		tempunit.LeaseStartYear, _ = strconv.Atoi(row[5])
		tempunit.LeaseEndMonth, _ = strconv.Atoi(row[6])
		tempunit.LeaseEndYear, _ = strconv.Atoi(row[7])
		tempunit.UnitStatus = row[8]
		tempunit.Tenant = row[9]
		tempunit.PassingRent, _ = strconv.ParseFloat(row[10], 64)
		tempunit.Probability, _ = strconv.ParseFloat(row[11], 64)
		tempunit.RentRevisionERV, _ = strconv.ParseFloat(row[12], 64)
		tempunit.EXTDuration, _ = strconv.Atoi(row[13])
		tempunit.IndexFreq, _ = strconv.Atoi(row[14])
		tempunit.IndexType = row[15]
		tempunit.IndexStartMonth, _ = strconv.Atoi(row[16])
		tempunit.Void, _ = strconv.Atoi(row[19])
		tempunit.DiscountRate, _ = strconv.ParseFloat(row[20], 64)
		tempunit.ERVArea, _ = strconv.ParseFloat(row[21], 64)
		tempunit.ERVAmount, _ = strconv.ParseFloat(row[22], 64)
		// tempunit.PercentSoldRent, _ = strconv.ParseFloat(row[23], 64)
		UnitStore[tempunit.MasterID] = tempunit
	}

	// GROWTHITEMSRAW
	// growth := xlsx.GetRows("Growth Items")
	// for i, row := range growth {
	// 	if i < 2 {
	// 		continue
	// 	}
	// 	tempgrowth := GrowthData{}
	// 	tempgrowth.EntityMasterID, _ = strconv.Atoi(row[1])
	// 	tempgrowth.Item = row[2]
	// 	tempgrowth.Amount, _ = strconv.ParseFloat(row[3], 64)
	// 	GrowthItemsRaw = append(GrowthItemsRaw, tempgrowth)
	// }

	//GROWTHITEMSSTORE
	// sort.Sort(GrowthItemsRaw)
	// for _, v := range EntityStore {
	// 	GrowthItemsStore[v.MasterID] = make(map[string]float64)
	// 	for _, vv := range GrowthItemsRaw {
	// 		if vv.EntityMasterID == v.MasterID {
	// 			GrowthItemsStore[v.MasterID][vv.Item] = vv.Amount
	// 		}
	// 	}
	// }
	// for _, v := range units {
	// 	fmt.Println(v)
	// }

	// for i, v := range UnitStore {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(UnitStore[0])
	// fmt.Println(UnitStore[2])
}

// Associations -
// func Associations() {
// 	for _, v := range EntityStore {
// 		EntityAssociations[v.Parent] = append(EntityAssociations[v.Parent], v.MasterID)
// 	}
// 	fmt.Println(EntityAssociations)
// 	for _, v := range UnitStore {
// 		UnitAssociations[v.ParentMasterID] = append(UnitAssociations[v.ParentMasterID], v.MasterID)
// 	}
// }

// // GDSlice -
// type GDSlice []GrowthData

// // Len -
// func (gds GDSlice) Len() int {
// 	return len(gds)
// }

// // Less - return whether the element with index i should sort before the element with index j
// func (gds GDSlice) Less(i, j int) bool {
// 	if gds[i].EntityMasterID < gds[j].EntityMasterID {
// 		return true
// 	}
// 	return false
// }

// // Swap -
// func (gds GDSlice) Swap(i, j int) {
// 	gds[i], gds[j] = gds[j], gds[i]
// }
