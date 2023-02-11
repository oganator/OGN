package controllers

import (
	"fmt"
	"strconv"
	"sync"

	xl "github.com/xuri/excelize/v2"
)

var XLSX, _ = xl.OpenFile("./models/Data.xlsx")

// EntityData -
type EntityModelData struct {
	Mutex               *sync.Mutex
	Sims                int
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
	GLA                 UnitModel
	VoidSigma           float64
	ProbabilitySigma    float64
	OpExSigma           float64
	Hazard              float64
	EntityID            int
	Version             string
}

// UnitData -
type UnitModelData struct {
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
func WriteXLSXEntities(e *EntityModel) {
	row, _ := XLSX.SearchSheet("EntityModels", e.Name)
	rows := row[0]
	rows = rows[1:]
	XLSX.SetCellValue("EntityModels", "K"+fmt.Sprint(rows), e.GrowthInput["CPI"].ShortTermRate)
	XLSX.SetCellValue("EntityModels", "L"+fmt.Sprint(rows), e.GrowthInput["CPI"].ShortTermPeriod)
	XLSX.SetCellValue("EntityModels", "M"+fmt.Sprint(rows), e.GrowthInput["CPI"].TransitionPeriod)
	XLSX.SetCellValue("EntityModels", "N"+fmt.Sprint(rows), e.GrowthInput["CPI"].LongTermRate)
	XLSX.SetCellValue("EntityModels", "O"+fmt.Sprint(rows), e.GrowthInput["ERV"].ShortTermRate)
	XLSX.SetCellValue("EntityModels", "P"+fmt.Sprint(rows), e.GrowthInput["ERV"].ShortTermPeriod)
	XLSX.SetCellValue("EntityModels", "Q"+fmt.Sprint(rows), e.GrowthInput["ERV"].TransitionPeriod)
	XLSX.SetCellValue("EntityModels", "R"+fmt.Sprint(rows), e.GrowthInput["ERV"].LongTermRate)
	XLSX.SetCellValue("EntityModels", "S"+fmt.Sprint(rows), e.Valuation.EntryYield)
	XLSX.SetCellValue("EntityModels", "T"+fmt.Sprint(rows), e.Valuation.YieldShift)
	XLSX.SetCellValue("EntityModels", "V"+fmt.Sprint(rows), e.HoldPeriod)
	XLSX.SetCellValue("EntityModels", "W"+fmt.Sprint(rows), e.Tax.RETT)
	XLSX.SetCellValue("EntityModels", "X"+fmt.Sprint(rows), e.Tax.VAT)
	XLSX.SetCellValue("EntityModels", "Y"+fmt.Sprint(rows), e.Tax.MinValue)
	XLSX.SetCellValue("EntityModels", "Z"+fmt.Sprint(rows), e.Tax.UsablePeriod)
	XLSX.SetCellValue("EntityModels", "AA"+fmt.Sprint(rows), e.Tax.LandValue)
	XLSX.SetCellValue("EntityModels", "AB"+fmt.Sprint(rows), e.DebtInput.LTV)
	XLSX.SetCellValue("EntityModels", "AC"+fmt.Sprint(rows), e.DebtInput.InterestRate)
	XLSX.SetCellValue("EntityModels", "AD"+fmt.Sprint(rows), e.OpEx.PercentOfTRI)
	XLSX.SetCellValue("EntityModels", "AE"+fmt.Sprint(rows), e.Tax.CarryBackYrs)
	XLSX.SetCellValue("EntityModels", "AF"+fmt.Sprint(rows), e.Tax.CarryForwardYrs)
	XLSX.SetCellValue("EntityModels", "AG"+fmt.Sprint(rows), e.GLA.PercentSoldRent)
	XLSX.SetCellValue("EntityModels", "AH"+fmt.Sprint(rows), e.Strategy)
	XLSX.SetCellValue("EntityModels", "AI"+fmt.Sprint(rows), e.Fees.PercentOfGAV)
	XLSX.SetCellValue("EntityModels", "AJ"+fmt.Sprint(rows), e.GLA.DiscountRate)
	XLSX.SetCellValue("EntityModels", "AK"+fmt.Sprint(rows), e.GLA.Void)
	XLSX.SetCellValue("EntityModels", "AL"+fmt.Sprint(rows), e.GLA.EXTDuration)
	XLSX.SetCellValue("EntityModels", "AM"+fmt.Sprint(rows), e.GLA.RentRevisionERV)
	XLSX.SetCellValue("EntityModels", "AN"+fmt.Sprint(rows), e.GLA.Probability)
	XLSX.SetCellValue("EntityModels", "AO"+fmt.Sprint(rows), e.BalloonPercent)
	XLSX.SetCellValue("EntityModels", "AP"+fmt.Sprint(rows), e.MCSetup.YieldShift)
	XLSX.SetCellValue("EntityModels", "AQ"+fmt.Sprint(rows), e.MCSetup.Void)
	XLSX.SetCellValue("EntityModels", "AR"+fmt.Sprint(rows), e.MCSetup.Probability)
	XLSX.SetCellValue("EntityModels", "AS"+fmt.Sprint(rows), e.MCSetup.OpEx)
	XLSX.SetCellValue("EntityModels", "AT"+fmt.Sprint(rows), e.MCSetup.CPI.ShortTermRate)
	XLSX.SetCellValue("EntityModels", "AU"+fmt.Sprint(rows), e.MCSetup.CPI.ShortTermPeriod)
	XLSX.SetCellValue("EntityModels", "AV"+fmt.Sprint(rows), e.MCSetup.CPI.TransitionPeriod)
	XLSX.SetCellValue("EntityModels", "AW"+fmt.Sprint(rows), e.MCSetup.CPI.LongTermRate)
	XLSX.SetCellValue("EntityModels", "AX"+fmt.Sprint(rows), e.MCSetup.ERV.ShortTermRate)
	XLSX.SetCellValue("EntityModels", "AY"+fmt.Sprint(rows), e.MCSetup.ERV.ShortTermPeriod)
	XLSX.SetCellValue("EntityModels", "AZ"+fmt.Sprint(rows), e.MCSetup.ERV.TransitionPeriod)
	XLSX.SetCellValue("EntityModels", "BA"+fmt.Sprint(rows), e.MCSetup.ERV.LongTermRate)
	XLSX.SetCellValue("EntityModels", "BB"+fmt.Sprint(rows), e.GLA.Default.Hazard)
	XLSX.SetCellValue("EntityModels", "BC"+fmt.Sprint(rows), e.GLA.RentIncentives.Duration)
	XLSX.SetCellValue("EntityModels", "BD"+fmt.Sprint(rows), e.GLA.RentIncentives.PercentOfContractRent)
	XLSX.SetCellValue("EntityModels", "BF"+fmt.Sprint(rows), e.GLA.FitOutCosts.AmountPerTotalArea)
	XLSX.SetCellValue("EntityModels", "BG"+fmt.Sprint(rows), e.MCSetup.Sims)
	XLSX.SetCellValue("EntityModels", "BH"+fmt.Sprint(rows), e.EntityID)
	XLSX.SetCellValue("EntityModels", "BI"+fmt.Sprint(rows), e.Version)
	XLSX.Save()
}

// WriteXLSXUnits -
func (u *UnitModelData) WriteXLSXUnits() {
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
	entityModels, _ := XLSX.GetRows("EntityModels")
	for i, row := range entityModels {
		if i < 2 {
			continue
		}
		tempentitymodel := EntityModelData{}
		tempentitymodel.Mutex = &sync.Mutex{}
		tempentitymodel.MasterID, _ = strconv.Atoi(row[1])
		tempentitymodel.Name = row[2]
		tempentitymodel.Parent, _ = strconv.Atoi(row[3])
		tempentitymodel.StartMonth, _ = strconv.Atoi(row[4])
		tempentitymodel.StartYear, _ = strconv.Atoi(row[5])
		tempentitymodel.SalesMonth, _ = strconv.Atoi(row[6])
		tempentitymodel.SalesYear, _ = strconv.Atoi(row[7])
		tempentitymodel.EndMonth, _ = strconv.Atoi(row[8])
		tempentitymodel.EndYear, _ = strconv.Atoi(row[9])
		tempentitymodel.CPIGrowth.ShortTermRate, _ = strconv.ParseFloat(row[10], 64)
		tempentitymodel.CPIGrowth.ShortTermPeriod, _ = strconv.Atoi(row[11])
		tempentitymodel.CPIGrowth.TransitionPeriod, _ = strconv.Atoi(row[12])
		tempentitymodel.CPIGrowth.LongTermRate, _ = strconv.ParseFloat(row[13], 64)
		tempentitymodel.ERVGrowth.ShortTermRate, _ = strconv.ParseFloat(row[14], 64)
		tempentitymodel.ERVGrowth.ShortTermPeriod, _ = strconv.Atoi(row[15])
		tempentitymodel.ERVGrowth.TransitionPeriod, _ = strconv.Atoi(row[16])
		tempentitymodel.ERVGrowth.LongTermRate, _ = strconv.ParseFloat(row[17], 64)
		tempentitymodel.EntryYield, _ = strconv.ParseFloat(row[18], 64)
		tempentitymodel.YieldShift, _ = strconv.ParseFloat(row[19], 64)
		tempentitymodel.ExitYield, _ = strconv.ParseFloat(row[20], 64)
		tempentitymodel.HoldPeriod, _ = strconv.Atoi(row[21])
		tempentitymodel.RETT, _ = strconv.ParseFloat(row[22], 64)
		tempentitymodel.VAT, _ = strconv.ParseFloat(row[23], 64)
		tempentitymodel.WOZpercent, _ = strconv.ParseFloat(row[24], 64)
		tempentitymodel.DeprPeriod, _ = strconv.Atoi(row[25])
		tempentitymodel.Landvalue, _ = strconv.ParseFloat(row[26], 64)
		tempentitymodel.CarryBackYrs, _ = strconv.Atoi(row[30])
		tempentitymodel.CarryForwardYrs, _ = strconv.Atoi(row[31])
		tempentitymodel.LTV, _ = strconv.ParseFloat(row[27], 64)
		tempentitymodel.LoanRate, _ = strconv.ParseFloat(row[28], 64)
		tempentitymodel.OpExpercent, _ = strconv.ParseFloat(row[29], 64)
		tempentitymodel.percentIncometosell, _ = strconv.ParseFloat(row[32], 64)
		tempentitymodel.DiscountRate, _ = strconv.ParseFloat(row[35], 64)
		tempentitymodel.Strategy = row[33]
		tempentitymodel.Fees, _ = strconv.ParseFloat(row[34], 64)
		tempentitymodel.GLA.DiscountRate, _ = strconv.ParseFloat(row[35], 64)
		tempentitymodel.GLA.Void, _ = strconv.Atoi(row[36])
		tempentitymodel.GLA.EXTDuration, _ = strconv.Atoi(row[37])
		tempentitymodel.GLA.RentRevisionERV, _ = strconv.ParseFloat(row[38], 64)
		tempentitymodel.GLA.Probability, _ = strconv.ParseFloat(row[39], 64)
		tempentitymodel.BalloonPercent, _ = strconv.ParseFloat(row[40], 64)
		tempentitymodel.YieldShiftSigma, _ = strconv.ParseFloat(row[41], 64)
		tempentitymodel.VoidSigma, _ = strconv.ParseFloat(row[42], 64)
		tempentitymodel.ProbabilitySigma, _ = strconv.ParseFloat(row[43], 64)
		tempentitymodel.OpExSigma, _ = strconv.ParseFloat(row[44], 64)
		tempentitymodel.CPISigma.ShortTermRate, _ = strconv.ParseFloat(row[45], 64)
		tempentitymodel.CPISigma.ShortTermPeriod, _ = strconv.Atoi(row[46])
		tempentitymodel.CPISigma.TransitionPeriod, _ = strconv.Atoi(row[47])
		tempentitymodel.CPISigma.LongTermRate, _ = strconv.ParseFloat(row[48], 64)
		tempentitymodel.ERVSigma.ShortTermRate, _ = strconv.ParseFloat(row[49], 64)
		tempentitymodel.ERVSigma.ShortTermPeriod, _ = strconv.Atoi(row[50])
		tempentitymodel.ERVSigma.TransitionPeriod, _ = strconv.Atoi(row[51])
		tempentitymodel.ERVSigma.LongTermRate, _ = strconv.ParseFloat(row[52], 64)
		tempentitymodel.GLA.Default.Hazard, _ = strconv.ParseFloat(row[53], 64)
		tempentitymodel.GLA.RentIncentives.Duration, _ = strconv.Atoi(row[54])
		tempentitymodel.GLA.RentIncentives.PercentOfContractRent, _ = strconv.ParseFloat(row[55], 64)
		tempentitymodel.GLA.FitOutCosts.AmountPerTotalArea, _ = strconv.ParseFloat(row[57], 64)
		tempentitymodel.Sims, _ = strconv.Atoi(row[58])
		tempentitymodel.EntityID, _ = strconv.Atoi(row[59])
		tempentitymodel.Version = row[60]
		EntityDataStore[tempentitymodel.MasterID] = &tempentitymodel
	}

	// UNITS
	unitModels, _ := XLSX.GetRows("UnitModels")
	// unitfile, _ := os.Open("./models/Units.csv")
	// units, _ := csv.NewReader(unitfile).ReadAll()
	for i, row := range unitModels {
		if i < 2 {
			continue
		}
		tempunit := UnitModelData{}
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

	// ENTITIES
	entities, _ := XLSX.GetRows("Entities")
	for i, row := range entities {
		if i < 2 {
			continue
		}
		tempentity := Entity{}
		tempentity.MasterID, _ = strconv.Atoi(row[1])
		tempentity.Name = row[2]
		// tempentity.ParentID, _ = strconv.Atoi(row[2])
		tempentity.AcquisitionDate.Month, _ = strconv.Atoi(row[4])
		tempentity.AcquisitionDate.Year, _ = strconv.Atoi(row[5])
		Dateadd(tempentity.AcquisitionDate, 0)
		tempentity.DispositionDate.Month, _ = strconv.Atoi(row[6])
		tempentity.DispositionDate.Year, _ = strconv.Atoi(row[7])
		Dateadd(tempentity.DispositionDate, 0)
		tempentity.EntityType = row[8]
		EntityMap[tempentity.MasterID] = &tempentity
		switch tempentity.EntityType {
		case "Fund":
			FundsList[tempentity.Name] = tempentity.MasterID
		case "Asset":
			AssetsList[tempentity.Name] = tempentity.MasterID
		}
	}
}
