package controllers

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

	_ "modernc.org/sqlite"
)

// ReadDB - Initial read of the SQLite db into Memory Objects.go
func ReadDB() {
	ReadEntites(DB)
	CreateEntityModels(DB, 0)
	DBAssociations(DB)
	CreateUnitModels(DB, 0, 0)
}

// ReadEntites - Reads entity table and marshals into EntityMap
func ReadEntites(db *sql.DB) {
	entityrows, err1 := db.Query(`
		select 
			entity.masterID, 
			entity.name, 
			entity.acquisition_month, 
			entity.acquisition_year, 
			entity.disposition_month,
			entity.disposition_year, 
			entity.entity_type
		from entity
	`)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer entityrows.Close()
	for entityrows.Next() {
		tempEntity := Entity{}
		err := entityrows.Scan(
			&tempEntity.MasterID,
			&tempEntity.Name,
			&tempEntity.AcquisitionDate.Month,
			&tempEntity.AcquisitionDate.Year,
			&tempEntity.DispositionDate.Month,
			&tempEntity.DispositionDate.Year,
			&tempEntity.EntityType,
		)
		if err != nil {
			fmt.Println(err)
		}
		tempEntity.AcquisitionDate.Add(0)
		tempEntity.DispositionDate.Add(0)
		tempEntity.Models = make([]*EntityModel, 0)
		switch tempEntity.EntityType {
		case "Fund":
			FundsList[tempEntity.Name] = tempEntity.MasterID
			EntitiesList[tempEntity.Name] = tempEntity.MasterID
		case "Asset":
			AssetsList[tempEntity.Name] = tempEntity.MasterID
			EntitiesList[tempEntity.Name] = tempEntity.MasterID
		}
		EntityMap[tempEntity.MasterID] = &tempEntity
	}
}

// CreateEntityModels - Reads entity_model table, and marshals into EntityModelsMap.
// Also creates GrowthInput, ChildEntityModels, ChildUnitModels and ChildUnitsMC maps
// Creates Mutex in EntityModelsMap and adds EntityModels to FundsList and AssetsList
// If entitymodel == 0, there is no 'where' clause to the sql query and all models are pulled,
// otherwise entitymodel will be used in the where clause
func CreateEntityModels(db *sql.DB, entitymodel int) {
	query := `		
		select 
			entity_model.masterID,
			entity_model.name,
			entity_model.entityID,
			entity_model.version,
			entity_model.start_month,
			entity_model.start_year,
			entity_model.sales_month,
			entity_model.sales_year,
			entity_model.cpi_short_rate,
			entity_model.cpi_short_rate_sigma,
			entity_model.cpi_short_period,
			entity_model.cpi_short_period_sigma,
			entity_model.cpi_transition,
			entity_model.cpi_transition_sigma,
			entity_model.cpi_long_rate,
			entity_model.cpi_long_rate_sigma, 
			entity_model.erv_short_rate, 
			entity_model.erv_short_rate_sigma, 
			entity_model.erv_short_period, 
			entity_model.erv_short_period_sigma, 
			entity_model.erv_transtion, 
			entity_model.erv_transition_sigma, 
			entity_model.erv_long_rate, 
			entity_model.erv_long_rate_sigma, 
			entity_model.entry_yield, 
			entity_model.yield_shift, 
			entity_model.yield_shift_sigma, 
			entity_model.rett, 
			entity_model.vat, 
			entity_model.woz, 
			entity_model.depr_period, 
			entity_model.land_value, 
			entity_model.carryback_yrs, 
			entity_model.carryforward_yrs, 
		/*		entity_model.opex_percent, 
			entity_model.opex_sigma, 
		*/		entity_model.strategy, 
		/*		entity_model.fees, 
		*/		entity_model.balloon,
			entity_model.uom,
			entity_model.sims,
			entity_model.valuation_method,
			entity_model.discount_rate,
			entity_model.acq_price,
			entity.entity_type,
			unit_models.incomeToSell,
			unit_models.masterID,
			unit_models.void,
			unit_models.void_sigma,
			unit_models.ext_dur,
			unit_models.rent_revision_erv,
			unit_models.probability,
			unit_models.probability_sigma,
			unit_models.hazard,
		/*		unit_models.incentives_months,
			unit_models.incentives_percent,
			unit_models.fitout_costs,
		*/		unit_models.discount_rate
		from entity_model
		left outer join entity on entity_model.entityID = entity.masterID
		left outer join unit_models on entity_model.masterID = unit_models.entity_model_ID
		where unit_models.unit_name like "%GLA%"
	`
	if entitymodel != 0 {
		query = query + `where entity_model.masterID = ` + fmt.Sprint(entitymodel)
	}
	entityModelRows, err1 := db.Query(query)
	if err1 != nil {
		fmt.Println("CreateEntityModels: ", err1)
	}
	defer entityModelRows.Close()
	for entityModelRows.Next() {
		tempModel := &EntityModel{}
		cpi := HModel{}
		erv := HModel{}
		entityType := ""
		err := entityModelRows.Scan(
			&tempModel.MasterID,
			&tempModel.Name,
			&tempModel.EntityID,
			&tempModel.Version,
			&tempModel.StartDate.Month,
			&tempModel.StartDate.Year,
			&tempModel.SalesDate.Month,
			&tempModel.SalesDate.Year,
			&cpi.ShortTermRate,
			&tempModel.MCSetup.CPI.ShortTermRate,
			&cpi.ShortTermPeriod,
			&tempModel.MCSetup.CPI.ShortTermPeriod,
			&cpi.TransitionPeriod,
			&tempModel.MCSetup.CPI.TransitionPeriod,
			&cpi.LongTermRate,
			&tempModel.MCSetup.CPI.LongTermRate,
			&erv.ShortTermRate,
			&tempModel.MCSetup.ERV.ShortTermRate,
			&erv.ShortTermPeriod,
			&tempModel.MCSetup.ERV.ShortTermPeriod,
			&erv.TransitionPeriod,
			&tempModel.MCSetup.ERV.TransitionPeriod,
			&erv.LongTermRate,
			&tempModel.MCSetup.ERV.LongTermRate,
			&tempModel.Valuation.EntryYield,
			&tempModel.Valuation.YieldShift,
			&tempModel.MCSetup.YieldShift,
			&tempModel.Tax.RETT,
			&tempModel.Tax.VAT,
			&tempModel.Tax.MinValue,
			&tempModel.Tax.UsablePeriod,
			&tempModel.Tax.LandValue,
			&tempModel.Tax.CarryBackYrs,
			&tempModel.Tax.CarryForwardYrs,
			// &tempModel.OpEx.PercentOfTRI,
			// &tempModel.MCSetup.OpEx,
			&tempModel.Strategy,
			// &tempModel.Fees.PercentOfGAV,
			&tempModel.BalloonPercent,
			&tempModel.UOM,
			&tempModel.MCSetup.Sims,
			&tempModel.Valuation.Method,
			&tempModel.Valuation.DiscountRate,
			&tempModel.Valuation.AcqPrice,
			&entityType,
			&tempModel.GLA.PercentSoldRent,
			&tempModel.GLA.MasterID,
			&tempModel.GLA.Void,
			&tempModel.MCSetup.Void,
			&tempModel.GLA.EXTDuration,
			&tempModel.GLA.RentRevisionERV,
			&tempModel.GLA.Probability,
			&tempModel.MCSetup.Probability,
			&tempModel.GLA.Default.Hazard,
			// &tempModel.GLA.RentIncentives.Duration,
			// &tempModel.GLA.RentIncentives.PercentOfContractRent,
			// &tempModel.GLA.FitOutCosts.AmountPerTotalArea,
			&tempModel.GLA.DiscountRate,
		)
		if err != nil {
			fmt.Println("CreateEntityModels.Rows: ", err)
		}
		tempModel.Mutex = &sync.Mutex{}
		tempModel.StartDate.Add(0)
		tempModel.SalesDate.Add(0)
		tempModel.EndDate = Dateadd(tempModel.SalesDate, 120)
		tempModel.Entity = EntityMap[tempModel.EntityID]
		tempModel.GrowthInput = make(map[string]HModel, 0)
		tempModel.GrowthInput["CPI"] = cpi
		tempModel.GrowthInput["ERV"] = erv
		tempModel.ChildEntityModels = map[int]*EntityModel{}
		tempModel.ChildUnitModels = map[int]*UnitModel{}
		tempModel.ChildUnitsMC = make(map[int]UnitModel)
		// tempModel.HoldPeriod = dateintdiff(tempModel.SalesDate.Dateint, tempModel.StartDate.Dateint) / 12
		tempModel.EndDate = Dateadd(tempModel.SalesDate, 60)
		tempModel.Growth = map[string]map[int]float64{}
		tempModel.CostInput = make(map[int]CostInput)
		tempModel.CostInput = CreateCosts(DB, tempModel.MasterID, "entity")
		tempModel.GLA.CostInput = make(map[int]CostInput)
		tempModel.GLA.CostInput = CreateCosts(DB, tempModel.GLA.MasterID, "unit")
		tempModel.Tax.DTA = map[int]float64{}
		tempModel.COA = map[int]FloatCOA{}
		tempModel.DebtInput = CreateLoans(DB, tempModel.MasterID)
		EntityModelsMap[tempModel.MasterID] = EntityMutex{
			Mutex:       &sync.Mutex{},
			EntityModel: tempModel,
		}
		switch entityType {
		case "Fund":
			FundModelsList[tempModel.Name] = tempModel.MasterID
			EntityModelsList[tempModel.Name] = tempModel.MasterID
			tempModel.Entity.Models = append(tempModel.Entity.Models, tempModel)
		case "Asset":
			AssetModelsList[tempModel.Name] = tempModel.MasterID
			EntityModelsList[tempModel.Name] = tempModel.MasterID
			tempModel.Entity.Models = append(tempModel.Entity.Models, tempModel)
		}
	}
}

// CreateUnitModels - Reads unit_models table and marshals into UnitModels,
// placing them in the Units map as well as each parent assets ChildUnitsModels map
// Also sets vacancy values
// WHAT TO DO WITH UNITS? should they be left as pointers? if so, this makes
// external MC a pain in the ass. If not, I have to tweak all UnitModel methods.
func CreateUnitModels(db *sql.DB, entitymodel int, unitmodel int) {
	query := `
	select 
		"unit_models"."masterID", 
		"unit_models"."unit_name", 
		"unit_models"."entity_model_ID", 
		"unit_models"."lease_start_month", 
		"unit_models"."lease_start_year", 
		"unit_models"."lease_end_month", 
		"unit_models"."lease_end_year", 
		"unit_models"."unit_status", 
		"unit_models"."tenant", 
		"unit_models"."tenant_type", 
		"unit_models"."passing_rent", 
		"unit_models"."erv_area", 
		"unit_models"."erv_amount", 
		"unit_models"."incomeToSell", 
		"unit_models"."void", 
		"unit_models"."void_sigma", 
		"unit_models"."ext_dur", 
		"unit_models"."rent_revision_erv", 
		"unit_models"."probability", 
		"unit_models"."probability_sigma", 
		"unit_models"."hazard", 
/*		"unit_models"."incentives_months", 
		"unit_models"."incentives_percent", 
		"unit_models"."fitout_costs", 
*/		"unit_models"."discount_rate"
	from unit_models 
	where unit_models.unit_name not like "%GLA%" 
	and unit_models.masterID <> 0
	`
	if entitymodel != 0 {
		query = query + `and entity_model_ID = ` + fmt.Sprint(entitymodel)
	}
	if unitmodel != 0 {
		query = query + `and masterID = ` + fmt.Sprint(entitymodel)
	}
	unitmodelrows, err1 := db.Query(query)
	if err1 != nil {
		fmt.Println("CreateUnitModels: ", err1)
	}
	defer unitmodelrows.Close()
	for unitmodelrows.Next() {
		tempUnitModel := UnitModel{}
		entityModelInt := 0
		err := unitmodelrows.Scan(
			&tempUnitModel.MasterID,
			&tempUnitModel.Name,
			&entityModelInt,
			&tempUnitModel.LeaseStartDate.Month,
			&tempUnitModel.LeaseStartDate.Year,
			&tempUnitModel.LeaseExpiryDate.Month,
			&tempUnitModel.LeaseExpiryDate.Year,
			&tempUnitModel.UnitStatus,
			&tempUnitModel.Tenant,
			&tempUnitModel.TenantType,
			&tempUnitModel.PassingRent,
			&tempUnitModel.ERVArea,
			&tempUnitModel.ERVAmount,
			&tempUnitModel.PercentSoldRent,
			&tempUnitModel.Void,
			&tempUnitModel.MCSetup.Void,
			&tempUnitModel.EXTDuration,
			&tempUnitModel.RentRevisionERV,
			&tempUnitModel.Probability,
			&tempUnitModel.MCSetup.Probability,
			&tempUnitModel.Default.Hazard,
			// &tempUnitModel.RentIncentives.Duration,
			// &tempUnitModel.RentIncentives.PercentOfContractRent,
			// &tempUnitModel.FitOutCosts.AmountPerTotalArea,
			&tempUnitModel.DiscountRate,
		)
		if err != nil {
			fmt.Println("CreateUnitModels.Rows: ", err)
		}
		parent := EntityModelsMap[entityModelInt].EntityModel
		tempUnitModel.Parent = parent
		switch tempUnitModel.UnitStatus {
		case "Vacant":
			tempUnitModel.LeaseStartDate.Month = parent.StartDate.Month
			tempUnitModel.LeaseStartDate.Year = parent.StartDate.Year
			tempUnitModel.LeaseStartDate.Add(-1)
			tempUnitModel.PassingRent = tempUnitModel.ERVArea * tempUnitModel.ERVAmount
		}
		// Set up values based on parent entity model
		if tempUnitModel.Probability == -1 {
			tempUnitModel.Probability = parent.GLA.Probability
		}
		if tempUnitModel.PercentSoldRent == -1 {
			tempUnitModel.PercentSoldRent = parent.GLA.PercentSoldRent
		}
		if tempUnitModel.Default.Hazard == -1 {
			tempUnitModel.Default = parent.GLA.Default
		}
		if tempUnitModel.RentRevisionERV == -1 {
			tempUnitModel.RentRevisionERV = parent.GLA.RentRevisionERV
		}
		if tempUnitModel.EXTDuration == -1 {
			tempUnitModel.EXTDuration = parent.GLA.EXTDuration
		}
		tempUnitModel.IndexDetails = parent.GLA.IndexDetails
		// if tempUnitModel.RentIncentives.Duration == -1 {
		// 	tempUnitModel.RentIncentives.Duration = parent.GLA.RentIncentives.Duration
		// }
		// if tempUnitModel.RentIncentives.PercentOfContractRent == -1 {
		// 	tempUnitModel.RentIncentives.PercentOfContractRent = parent.GLA.RentIncentives.PercentOfContractRent
		// }
		if tempUnitModel.Void == -1 {
			tempUnitModel.Void = parent.GLA.Void
		}
		// if tempUnitModel.FitOutCosts.AmountPerTotalArea == -1 {
		// 	tempUnitModel.FitOutCosts.AmountPerTotalArea = parent.GLA.FitOutCosts.AmountPerTotalArea
		// }
		if tempUnitModel.DiscountRate == -1 {
			tempUnitModel.DiscountRate = parent.GLA.DiscountRate
		}
		if tempUnitModel.UnitStatus == "Vacant" {
			tempUnitModel.LeaseExpiryDate = Dateadd(parent.StartDate, tempUnitModel.EXTDuration)
		}
		tempUnitModel.CostInput = make(map[int]CostInput)
		tempUnitModel.CostInput = CreateCosts(DB, tempUnitModel.MasterID, "unit")
		// create datetypes
		tempUnitModel.LeaseStartDate.Add(0)
		tempUnitModel.LeaseExpiryDate.Add(0)
		// assign to units maps for parent asset and fund
		Units[tempUnitModel.MasterID] = tempUnitModel
		parent.ChildUnitModels[tempUnitModel.MasterID] = &tempUnitModel
		parent.Parent.ChildUnitModels[tempUnitModel.MasterID] = &tempUnitModel
	}
}

// CreateLoans - em is the MasterID for the EntityModel
func CreateLoans(db *sql.DB, em int) (loanArray []DebtInput) {
	query := `
		select
			"masterID",
			"name",
			"ltv",
			"loan_amount", 
			"fixed_rate",
			"interest_type", 
			"loan_type", 
			"loan_start_month", 
			"loan_start_year", 
			"loan_end_month",
			"loan_end_year", 
			"float_basis", 
			"spread", 
			"amortization_period", 
			"active",
			"loan_basis",
			"start_event",
			"end_event"
			from debt where entity_model_id = 
		`
	loans, err := db.Query(query + fmt.Sprint(em))
	if err != nil {
		fmt.Println("CreateLoan.Query: ", err)
	}
	tempLoan := DebtInput{}
	defer loans.Close()
	for loans.Next() {
		err2 := loans.Scan(
			&tempLoan.MasterID,
			&tempLoan.Name,
			&tempLoan.LTV,
			&tempLoan.Amount,
			&tempLoan.InterestRate,
			&tempLoan.InterestType,
			&tempLoan.LoanType,
			&tempLoan.LoanStart.Month,
			&tempLoan.LoanStart.Year,
			&tempLoan.LoanEnd.Month,
			&tempLoan.LoanEnd.Year,
			&tempLoan.FloatBasis,
			&tempLoan.Spread,
			&tempLoan.AmortizationPeriod,
			&tempLoan.Active,
			&tempLoan.LoanBasis,
			&tempLoan.StartEvent,
			&tempLoan.EndEvent,
		)
		if err2 != nil {
			fmt.Println("CreateLoan.Scan: ", err2)
		}
		// tempLoan.LoanStart.Add(0)
		// tempLoan.LoanEnd.Add(0)
		loanArray = append(loanArray, tempLoan)
	}
	return loanArray
}

// CreateCapex - model is the MasterID for the EntityModel OR UnitModel, modelType is either "entity" or "unit"
func CreateCosts(db *sql.DB, masterID int, modelType string) (capexMap map[int]CostInput) {
	capexMap = make(map[int]CostInput)
	query := `
		select
			"masterID",
			"name",
			"type",
			"name",
			"amount",
			"amount_sigma",
			"coa_item_basis",
			"coa_item_target", 
			"duration", 
			"duration_sigma",
			"start_month", 
			"start_year",
			"start_event",
			"end_month", 
			"end_year", 
			"end_event",
			"growth_item"
		from cost where
		`
	modelTypeQuery := ""
	switch modelType {
	case "entity":
		modelTypeQuery = ` entity_modelID = ` + fmt.Sprint(masterID)
	case "unit":
		modelTypeQuery = ` unit_modelID = ` + fmt.Sprint(masterID)
	}
	tempCost := CostInput{}
	name := ""
	costs, err := db.Query(query + modelTypeQuery)
	if err != nil {
		fmt.Println("CreateCosts.Query: ", err)
		fmt.Println(query + modelTypeQuery)
	}
	defer costs.Close()
	for costs.Next() {
		err2 := costs.Scan(
			&tempCost.MasterID,
			&tempCost.Name,
			&tempCost.Type,
			&name,
			&tempCost.Amount,
			&tempCost.AmountSigma,
			&tempCost.COAItemBasis,
			&tempCost.COAItemTarget,
			&tempCost.Duration,
			&tempCost.DurationSigma,
			&tempCost.Start.Month,
			&tempCost.Start.Year,
			&tempCost.StartEvent,
			&tempCost.End.Month,
			&tempCost.End.Year,
			&tempCost.EndEvent,
			&tempCost.GrowthItem,
		)
		if err2 != nil {
			fmt.Println("CreateCosts.Scan: ", err2)
		}
		tempCost.Start.Add(0)
		tempCost.End.Add(0)
		capexMap[tempCost.MasterID] = tempCost
	}
	return capexMap
}

// YearlyRates - currently not called as the HModel is used instead
func YearlyRates(db *sql.DB, loanID int) (result IntFloatMap) {
	result = make(IntFloatMap)
	query := `
		select
			period,
			rate
		from yearly_rates
		where loanID = 
	`
	rates, err := db.Query(query + fmt.Sprint(loanID))
	if err != nil {
		fmt.Println("YearlyRates.Query: ", err)
	}
	defer rates.Close()
	for rates.Next() {
		period := 0
		value := 0.0
		err2 := rates.Scan(
			&period,
			&value,
		)
		if err2 != nil {
			fmt.Println("YearlyRates.Scan: ", err2)
		}
		result[period] = value
	}
	return result
}

// DBAssociations - Reads from entity_model_associations table, and then adds relevant entries to EntityModelsMap
func DBAssociations(db *sql.DB) {
	associations, err1 := db.Query(`
	select masterID, parent, child, ownership from entity_model_associations
	`)
	if err1 != nil {
		fmt.Println("DBAssociations.Query: ", err1)
	}
	defer associations.Close()
	for associations.Next() {
		parent := 0
		child := 0
		ownership := 0.0
		masterID := 0
		err := associations.Scan(
			&masterID,
			&parent,
			&child,
			&ownership,
		)
		if err != nil {
			fmt.Println("DBAssociations.Scan", err)
		}
		EntityModelsMap[parent].EntityModel.ChildEntityModels[child] = EntityModelsMap[child].EntityModel
		EntityModelsMap[child].EntityModel.Parent = EntityModelsMap[parent].EntityModel
		EntityModelsMap[child].EntityModel.ParentID = parent
	}
}

// UpdateEntityModel - dbWrite determines if the EntityModel is written back to the db.
// UnitModels are always queried from the db, and then EntityModelCalc is executed
func (e *EntityModel) UpdateEntityModel(dbWrite bool) {
	if dbWrite {
		e.WriteDBEntityModel(DB)
	}
	CreateUnitModels(DB, e.MasterID, 0)
	e.EntityModelCalc(false, "Internal")
}

// WriteDBEntityModel - writes changes to the entity_model table, unit_model table and debt table for a specific entity model
func (e *EntityModel) WriteDBEntityModel(db *sql.DB) {
	queryEntityModel := "update entity_model set start_month = " + fmt.Sprint(e.StartDate.Month) + " , start_year = " + fmt.Sprint(e.StartDate.Year) + " , sales_month = " + fmt.Sprint(e.SalesDate.Month) + " , sales_year = " + fmt.Sprint(e.SalesDate.Year) + " , entry_yield = " + fmt.Sprint(e.Valuation.EntryYield) + " , acq_price = " + fmt.Sprint(e.Valuation.AcqPrice) + " , valuation_method = '" + fmt.Sprint(e.Valuation.Method) + "' , yield_shift = " + fmt.Sprint(e.Valuation.YieldShift) + " , yield_shift_sigma = " + fmt.Sprint(e.MCSetup.YieldShift) + " , rett = " + fmt.Sprint(e.Tax.RETT) + " , vat = " + fmt.Sprint(e.Tax.VAT) + " , woz = " + fmt.Sprint(e.Tax.MinValue) + " , depr_period = " + fmt.Sprint(e.Tax.UsablePeriod) + " , land_value = " + fmt.Sprint(e.Tax.LandValue) + " , carryback_yrs = " + fmt.Sprint(e.Tax.CarryBackYrs) + " , carryforward_yrs = " + fmt.Sprint(e.Tax.CarryForwardYrs) + " , strategy = '" + fmt.Sprint(e.Strategy) + "' , balloon = " + fmt.Sprint(e.BalloonPercent) + " , sims = " + fmt.Sprint(e.MCSetup.Sims) + " where masterID = " + fmt.Sprint(e.MasterID) // + " , opex_percent = " + fmt.Sprint(e.OpEx.PercentOfTRI) + " , opex_sigma = " + fmt.Sprint(e.MCSetup.OpEx) + "' , fees = " + fmt.Sprint(e.Fees.PercentOfGAV)
	_, err := DB.Exec(queryEntityModel)
	if err != nil {
		fmt.Println("WriteDBEntityModel.queryEntityModel: ", err)
		fmt.Println(queryEntityModel)
	}
	queryUnitModel := "update unit_models set incomeToSell = " + fmt.Sprint(e.GLA.PercentSoldRent) + " , void = " + fmt.Sprint(e.GLA.Void) + " , void_sigma = " + fmt.Sprint(e.MCSetup.Void) + " , ext_dur = " + fmt.Sprint(e.GLA.EXTDuration) + " , rent_revision_erv = " + fmt.Sprint(e.GLA.RentRevisionERV) + " , probability = " + fmt.Sprint(e.GLA.Probability) + " , probability_sigma = " + fmt.Sprint(e.MCSetup.Probability) + " , hazard = " + fmt.Sprint(e.GLA.Default.Hazard) + " , discount_rate = " + fmt.Sprint(e.GLA.DiscountRate) + ` where unit_name like "%GLA%" and entity_model_ID = ` + fmt.Sprint(e.MasterID) // + " , incentives_percent = " + fmt.Sprint(e.GLA.RentIncentives.PercentOfContractRent) + " , fitout_costs = " + fmt.Sprint(e.GLA.FitOutCosts.AmountPerTotalArea)+ " , incentives_months = " + fmt.Sprint(e.GLA.RentIncentives.Duration)
	_, err2 := DB.Exec(queryUnitModel)
	if err2 != nil {
		fmt.Println("WriteDBEntityModel.queryUnitModel: ", err2)
		fmt.Println(queryUnitModel)
	}
	for _, loan := range e.DebtInput {
		queryDebt := "update debt set ltv = " + fmt.Sprint(loan.LTV) + " , name = " + fmt.Sprint("'", loan.Name, "'") + " , loan_amount = " + fmt.Sprint(loan.Amount) + " , loan_type = " + fmt.Sprint("'", loan.LoanType, "'") + " , fixed_rate = " + fmt.Sprint(loan.InterestRate) + " , interest_type = " + fmt.Sprint("'", loan.InterestType, "'") + " , loan_start_month = " + fmt.Sprint(loan.LoanStart.Month) + " , loan_start_year = " + fmt.Sprint(loan.LoanStart.Year) + " , loan_end_month = " + fmt.Sprint(loan.LoanEnd.Month) + " , loan_end_year = " + fmt.Sprint(loan.LoanEnd.Year) + " , float_basis = " + fmt.Sprint("'", loan.FloatBasis, "'") + " , spread = " + fmt.Sprint(loan.Spread) + " , active = " + fmt.Sprint("'", loan.Active, "'") + " , amortization_period = " + fmt.Sprint(loan.AmortizationPeriod) + " , start_event = " + fmt.Sprint("'", loan.StartEvent, "'") + " , end_event = " + fmt.Sprint("'", loan.EndEvent, "'") + " where masterID = " + fmt.Sprint(loan.MasterID)
		_, err3 := DB.Exec(queryDebt)
		if err3 != nil {
			fmt.Println("WriteDBEntityModel.queryDebt: ", err3)
			fmt.Println(queryDebt)
		}
	}
	for _, cost := range e.CostInput {
		queryCost := "update cost set type = " + fmt.Sprint("'", cost.Type, "'") + " , name = " + fmt.Sprint("'", cost.Name, "'") + " , amount = " + fmt.Sprint("'", cost.Amount, "'") + " , amount_sigma = " + fmt.Sprint(cost.AmountSigma) + " , coa_item_basis = " + fmt.Sprint("'", cost.COAItemBasis, "'") + " , coa_item_target = " + fmt.Sprint("'", cost.COAItemTarget, "'") + " , duration = " + fmt.Sprint(cost.Duration) + " , duration_sigma = " + fmt.Sprint(cost.DurationSigma) + " , start_month = " + fmt.Sprint(cost.Start.Month) + " , start_year = " + fmt.Sprint("'", cost.Start.Year, "'") + " , start_event = " + fmt.Sprint("'", cost.StartEvent, "'") + " , end_month = " + fmt.Sprint("'", cost.End.Month, "'") + " , end_year = " + fmt.Sprint(cost.End.Year) + " , end_event = " + fmt.Sprint("'", cost.EndEvent, "'") + " , growth_item = " + fmt.Sprint("'", cost.GrowthItem, "'") + " where masterID = " + fmt.Sprint(cost.MasterID)
		_, err3 := DB.Exec(queryCost)
		if err3 != nil {
			fmt.Println("WriteDBEntityModel.queryCost: ", err3)
			fmt.Println(queryCost)
		}
	}
}

// WriteDBUnitModelSingleValue - updates a particular value for a specific unit, then updates the units map
func WriteDBUnitModelSingleValue(unit int, field string, value string) {
	unitString := strconv.Itoa(unit)
	query := "update unit_models set " + field + " = " + value + " where masterID = " + unitString
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println("WriteDBUnitModelSingleValue: ", err)
		fmt.Println(query)
	}
	CreateUnitModels(DB, 0, unit)
}

// creates a new unit model
func (u *UnitModel) WriteDBUnitModel() {
	query :=
		`insert into unit_models (
			"unit_name", 
			"entity_model_ID", 
			"lease_start_month", 
			"lease_start_year", 
			"lease_end_month", 
			"lease_end_year", 
			"unit_status", 
			"tenant", 
			"tenant_type
			"passing_rent", 
			"erv_area", 
			"erv_amount")
			values(`
	s := string("' , '")
	query = query + fmt.Sprint(" '", u.Name, s, u.Parent.MasterID, s, u.LeaseStartDate.Month, s, u.LeaseStartDate.Year, s, u.LeaseExpiryDate.Month, s, u.LeaseExpiryDate.Year, s, u.UnitStatus, s, u.Tenant, s, u.TenantType, s, u.PassingRent, s, u.ERVArea, s, u.ERVAmount, "' );")
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println("WriteDBUnitModel: ", err)
		fmt.Println(query)
	}
}

// Adds a new Cost to the DB, and returns that cost back to the caller
func AddCostInput(e, u int) CostInput {
	query :=
		`insert into cost (
		"entity_modelID", 
		"unit_modelID")
		values(`
	s := string("' , '")
	query = query + fmt.Sprint(" '", e, s, u, "' );")
	cost, err := DB.Query(query)
	if err != nil {
		fmt.Println("AddCostInput: ", err)
		fmt.Println(query)
	}
	defer cost.Close()
	resultCost := CostInput{}
	for cost.Next() {
		err := cost.Scan(
			&resultCost.MasterID,
			&resultCost.Type,
			&resultCost.Name,
			&resultCost.Amount,
			&resultCost.AmountSigma,
			&resultCost.COAItemBasis,
			&resultCost.COAItemTarget,
			&resultCost.Duration,
			&resultCost.DurationSigma,
			&resultCost.Start.Month,
			&resultCost.Start.Year,
			&resultCost.StartEvent,
			&resultCost.End.Month,
			&resultCost.End.Year,
			&resultCost.EndEvent,
			&resultCost.GrowthItem,
		)
		if err != nil {
			fmt.Println("AddCostInput.Scan", err)
		}
	}
	return resultCost
}
