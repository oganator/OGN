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
	CreateUnitModels(DB, 0, 0)
	DBAssociations(DB)
}

// ReadEntites - Reads entity table and marshals into EntityMap
func ReadEntites(db *sql.DB) {
	entityrows, _ := db.Query(`
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
		entity_model.ltv, 
		entity_model.loan_rate, 
		entity_model.opex_percent, 
		entity_model.opex_sigma, 
		entity_model.strategy, 
		entity_model.fees, 
		entity_model.balloon,
		entity_model.uom,
		entity_model.sims,
		entity_model.valuation_method,
		entity_model.discount_rate,
		entity_model.acq_price,
		entity.entity_type,
		unit_models.incomeToSell,
		unit_models.void,
		unit_models.void_sigma,
		unit_models.ext_dur,
		unit_models.rent_revision_erv,
		unit_models.probability,
		unit_models.probability_sigma,
		unit_models.hazard,
		unit_models.incentives_months,
		unit_models.incentives_percent,
		unit_models.fitout_costs,
		unit_models.discount_rate
	from entity_model
	join entity on entity_model.entityID = entity.masterID
	join unit_models on entity_model.masterID = unit_models.entity_model_ID
	where unit_models.unit_name like "%GLA%"
	`
	if entitymodel != 0 {
		query = query + `where entity_model.masterID = ` + fmt.Sprint(entitymodel)
	}
	entityModelRows, _ := db.Query(query)
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
			&tempModel.DebtInput.LTV,
			&tempModel.DebtInput.InterestRate,
			&tempModel.OpEx.PercentOfTRI,
			&tempModel.MCSetup.OpEx,
			&tempModel.Strategy,
			&tempModel.Fees.PercentOfGAV,
			&tempModel.BalloonPercent,
			&tempModel.UOM,
			&tempModel.MCSetup.Sims,
			&tempModel.Valuation.Method,
			&tempModel.Valuation.DiscountRate,
			&tempModel.Valuation.AcqPrice,
			&entityType,
			&tempModel.GLA.PercentSoldRent,
			&tempModel.GLA.Void,
			&tempModel.MCSetup.Void,
			&tempModel.GLA.EXTDuration,
			&tempModel.GLA.RentRevisionERV,
			&tempModel.GLA.Probability,
			&tempModel.MCSetup.Probability,
			&tempModel.GLA.Default.Hazard,
			&tempModel.GLA.RentIncentives.Duration,
			&tempModel.GLA.RentIncentives.PercentOfContractRent,
			&tempModel.GLA.FitOutCosts.AmountPerTotalArea,
			&tempModel.GLA.DiscountRate,
		)
		if err != nil {
			fmt.Println(err)
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
		tempModel.Capex = map[int]CostInput{}
		tempModel.Tax.DTA = map[int]float64{}
		tempModel.COA = map[int]FloatCOA{}
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
		"unit_models"."incentives_months", 
		"unit_models"."incentives_percent", 
		"unit_models"."fitout_costs", 
		"unit_models"."discount_rate"
	from unit_models 
	where unit_models.unit_name not like "%GLA%"
	`
	if entitymodel != 0 {
		query = query + `and entity_model_ID = ` + fmt.Sprint(entitymodel)
	}
	if unitmodel != 0 {
		query = query + `and masterID = ` + fmt.Sprint(entitymodel)
	}
	unitmodelrows, _ := db.Query(query)
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
			&tempUnitModel.RentIncentives.Duration,
			&tempUnitModel.RentIncentives.PercentOfContractRent,
			&tempUnitModel.FitOutCosts.AmountPerTotalArea,
			&tempUnitModel.DiscountRate,
		)
		if err != nil {
			fmt.Println(err)
		}
		parent := EntityModelsMap[entityModelInt].EntityModel
		tempUnitModel.Parent = parent
		switch tempUnitModel.UnitStatus {
		case "Vacant":
			tempUnitModel.LeaseStartDate.Month = parent.StartDate.Month
			tempUnitModel.LeaseStartDate.Year = parent.StartDate.Year
			tempUnitModel.LeaseStartDate.Add(-1)
			tempUnitModel.LeaseExpiryDate = Dateadd(parent.StartDate, parent.GLA.EXTDuration)
			tempUnitModel.PassingRent = tempUnitModel.ERVArea * tempUnitModel.ERVAmount
		case "Occupied":
			tempUnitModel.LeaseStartDate.Add(0)
			tempUnitModel.LeaseExpiryDate.Add(0)
		}
		// Set up values based on parent entity model
		// 		TODO - add these fields to the unit_models table in the db, and have logic to first
		// 		put in the parents, then the unit specific values so they override
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
		if tempUnitModel.RentIncentives.Duration == -1 {
			tempUnitModel.RentIncentives.Duration = parent.GLA.RentIncentives.Duration
		}
		if tempUnitModel.RentIncentives.PercentOfContractRent == -1 {
			tempUnitModel.RentIncentives.PercentOfContractRent = parent.GLA.RentIncentives.PercentOfContractRent
		}
		if tempUnitModel.Void == -1 {
			tempUnitModel.Void = parent.GLA.Void
		}
		if tempUnitModel.FitOutCosts.AmountPerTotalArea == -1 {
			tempUnitModel.FitOutCosts.AmountPerTotalArea = parent.GLA.FitOutCosts.AmountPerTotalArea
		}
		if tempUnitModel.DiscountRate == -1 {
			tempUnitModel.DiscountRate = parent.GLA.DiscountRate
		}
		// create datetypes
		tempUnitModel.LeaseStartDate.Add(0)
		tempUnitModel.LeaseExpiryDate.Add(0)
		// assign to relevant maps
		Units[tempUnitModel.MasterID] = tempUnitModel
		parent.ChildUnitModels[tempUnitModel.MasterID] = &tempUnitModel
	}
}

// DBAssociations - Reads from entity_model_associations table, and then adds relevant entries to EntityModelsMap
func DBAssociations(db *sql.DB) {
	associations, _ := db.Query(`
	select parent, child, ownership from entity_model_associations
	`)
	for associations.Next() {
		parent := 0
		child := 0
		ownership := 0.0
		err := associations.Scan(
			&parent,
			&child,
			&ownership,
		)
		if err != nil {
			fmt.Println(err)
		}
		EntityModelsMap[parent].EntityModel.ChildEntityModels[child] = EntityModelsMap[child].EntityModel
		EntityModelsMap[child].EntityModel.Parent = EntityModelsMap[parent].EntityModel
		EntityModelsMap[child].EntityModel.ParentID = parent
	}
}

// UpdateEntityModel -
func (e *EntityModel) UpdateEntityModel() {
	go e.WriteDBEntityModel(DB)
	CreateUnitModels(DB, e.MasterID, 0)
	e.EntityModelCalc(false, "Internal")
}

// WriteDBEntityModel - writes changes to the entity_model table for a specific entity model
func (e *EntityModel) WriteDBEntityModel(db *sql.DB) {
	query := "update entity_model set start_month = " + fmt.Sprint(e.StartDate.Month) + " , start_year = " + fmt.Sprint(e.StartDate.Year) + " , sales_month = " + fmt.Sprint(e.SalesDate.Month) + " , sales_year = " + fmt.Sprint(e.SalesDate.Year) + " , entry_yield = " + fmt.Sprint(e.Valuation.EntryYield) + " , yield_shift = " + fmt.Sprint(e.Valuation.YieldShift) + " , yield_shift_sigma = " + fmt.Sprint(e.MCSetup.YieldShift) + " , rett = " + fmt.Sprint(e.Tax.RETT) + " , vat = " + fmt.Sprint(e.Tax.VAT) + " , woz = " + fmt.Sprint(e.Tax.MinValue) + " , depr_period = " + fmt.Sprint(e.Tax.UsablePeriod) + " , land_value = " + fmt.Sprint(e.Tax.LandValue) + " , carryback_yrs = " + fmt.Sprint(e.Tax.CarryBackYrs) + " , carryforward_yrs = " + fmt.Sprint(e.Tax.CarryForwardYrs) + " , ltv = " + fmt.Sprint(e.DebtInput.LTV) + " , loan_rate = " + fmt.Sprint(e.DebtInput.InterestRate) + " , opex_percent = " + fmt.Sprint(e.OpEx.PercentOfTRI) + " , opex_sigma = " + fmt.Sprint(e.MCSetup.OpEx) + " , strategy = '" + fmt.Sprint(e.Strategy) + "' , fees = " + fmt.Sprint(e.Fees.PercentOfGAV) + " , balloon = " + fmt.Sprint(e.BalloonPercent) + " , sims = " + fmt.Sprint(e.MCSetup.Sims) + " where masterID = " + fmt.Sprint(e.MasterID)
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
}

// WriteDBUnitModelSingleValue - updates a particular value for a specific unit, then updates the units map
func WriteDBUnitModelSingleValue(unit int, field string, value string) {
	unitString := strconv.Itoa(unit)
	query := "update unit_models set " + field + " = " + value + " where masterID = " + unitString
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	CreateUnitModels(DB, 0, unit)
}
