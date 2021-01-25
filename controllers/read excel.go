package controllers

import (
	"sort"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/360EntSecGroup-Skylar/excelize"
)

// EntityData -
type EntityData struct {
	MasterID            int
	Name                string
	Parent              int // MasterID
	StartMonth          int
	StartYear           int
	SalesMonth          int
	SalesYear           int
	EndMonth            int
	EndYear             int
	CPIGrowth           float64
	ERVGrowth           float64
	EntryYield          float64
	YieldShift          float64
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
}

// UnitData -
type UnitData struct {
	MasterID          int
	Name              string
	ParentMasterID    int
	LeaseStartMonth   int
	LeaseStartYear    int
	LeaseEndMonth     int
	LeaseEndYear      int
	UnitStatus        string
	Tenant            string
	PassingRent       float64
	Probability       float64
	RentRevisionERV   float64
	EXTDuration       int
	IndexFreq         int
	IndexType         string
	IndexStartMonth   int
	IncentivesMonths  int
	IncentivesPercent float64
	Void              int
	DiscountRate      float64
	ERVArea           float64
	ERVAmount         float64
	PercentSoldRent   float64
}

// GrowthData -
type GrowthData struct {
	EntityMasterID int
	Item           string
	Amount         float64
}

// ReadXLSX - Reads Data.xlsx and populates data stores (Entity, Unit and GrowthItems)
func ReadXLSX() {
	xlsx, _ := excelize.OpenFile("./models/Data.xlsx")

	// entityfile, _ := os.Open("./models/Entities.csv")
	// ENTITIES
	entities := xlsx.GetRows("Entities")
	// entities, _ := csv.NewReader(entityfile).ReadAll()
	for i, row := range entities {
		if i < 2 {
			continue
		}
		tempentity := EntityData{}
		tempentity.MasterID, _ = strconv.Atoi(row[1])
		tempentity.Name = row[2]
		tempentity.Parent, _ = strconv.Atoi(row[3])
		tempentity.StartMonth, _ = strconv.Atoi(row[4])
		tempentity.StartYear, _ = strconv.Atoi(row[5])
		tempentity.SalesMonth, _ = strconv.Atoi(row[6])
		tempentity.SalesYear, _ = strconv.Atoi(row[7])
		tempentity.EndMonth, _ = strconv.Atoi(row[8])
		tempentity.EndYear, _ = strconv.Atoi(row[9])
		tempentity.CPIGrowth, _ = strconv.ParseFloat(row[10], 64)
		tempentity.ERVGrowth, _ = strconv.ParseFloat(row[11], 64)
		tempentity.EntryYield, _ = strconv.ParseFloat(row[12], 64)
		tempentity.YieldShift, _ = strconv.ParseFloat(row[13], 64)
		tempentity.ExitYield, _ = strconv.ParseFloat(row[14], 64)
		tempentity.HoldPeriod, _ = strconv.Atoi(row[15])
		tempentity.RETT, _ = strconv.ParseFloat(row[16], 64)
		tempentity.VAT, _ = strconv.ParseFloat(row[17], 64)
		tempentity.WOZpercent, _ = strconv.ParseFloat(row[18], 64)
		tempentity.DeprPeriod, _ = strconv.Atoi(row[19])
		tempentity.Landvalue, _ = strconv.ParseFloat(row[20], 64)
		tempentity.CarryBackYrs, _ = strconv.Atoi(row[24])
		tempentity.CarryForwardYrs, _ = strconv.Atoi(row[25])
		tempentity.LTV, _ = strconv.ParseFloat(row[21], 64)
		tempentity.LoanRate, _ = strconv.ParseFloat(row[22], 64)
		tempentity.OpExpercent, _ = strconv.ParseFloat(row[23], 64)
		tempentity.percentIncometosell, _ = strconv.ParseFloat(row[24], 64)
		tempentity.YearsIncometosell, _ = strconv.Atoi(row[25])
		tempentity.DiscountRate, _ = strconv.ParseFloat(row[26], 64)
		tempentity.Strategy = row[27]
		tempentity.Fees, _ = strconv.ParseFloat(row[28], 64)
		EntityStore[tempentity.MasterID] = &tempentity
	}
	// UNITS
	units := xlsx.GetRows("Units")
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
		tempunit.IncentivesMonths, _ = strconv.Atoi(row[17])
		tempunit.IncentivesPercent, _ = strconv.ParseFloat(row[18], 64)
		tempunit.Void, _ = strconv.Atoi(row[19])
		tempunit.DiscountRate, _ = strconv.ParseFloat(row[20], 64)
		tempunit.ERVArea, _ = strconv.ParseFloat(row[21], 64)
		tempunit.ERVAmount, _ = strconv.ParseFloat(row[22], 64)
		tempunit.PercentSoldRent, _ = strconv.ParseFloat(row[23], 64)
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
	sort.Sort(GrowthItemsRaw)
	for _, v := range EntityStore {
		GrowthItemsStore[v.MasterID] = make(map[string]float64)
		for _, vv := range GrowthItemsRaw {
			if vv.EntityMasterID == v.MasterID {
				GrowthItemsStore[v.MasterID][vv.Item] = vv.Amount
			}
		}
	}
}

// Associations -
func Associations() {
	for _, v := range EntityStore {
		EntityAssociations[v.Parent] = append(EntityAssociations[v.Parent], v.MasterID)
	}
	for _, v := range UnitStore {
		UnitAssociations[v.ParentMasterID] = append(UnitAssociations[v.ParentMasterID], v.MasterID)
	}
}

// GDSlice -
type GDSlice []GrowthData

// Len -
func (gds GDSlice) Len() int {
	return len(gds)
}

// Less - return whether the element with index i should sort before the element with index j
func (gds GDSlice) Less(i, j int) bool {
	if gds[i].EntityMasterID < gds[j].EntityMasterID {
		return true
	}
	return false
}

// Swap -
func (gds GDSlice) Swap(i, j int) {
	gds[i], gds[j] = gds[j], gds[i]
}

// 5373d780-2e0f-447d-89e3-2f9c4fd6e388  2021-01-09T13:21:34+00:00  2M28S     gs://propmodel-271202_cloudbuild/source/1610198475.813473-3a0ab15b6b024d7c818df175788723f5.tgz  gcr.io/propmodel-271202/ogn (+1 more)  SUCCESS
