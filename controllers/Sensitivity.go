package controllers

import (
	"fmt"
	"math"
	"strconv"

	beego "github.com/astaxie/beego"
)

// contains the increments used for sensitivity analysis
type SensitivityInput struct {
	Void        int               //
	Probability float64           //
	Growth      map[string]HModel // Short and long rates
	EntityCosts map[int]CostInput // Amount
	UnitCosts   map[int]CostInput // based on the GLA from the EntityModel
	Debt        []DebtInput       // spread
	Valuation   Valuation         // discount rate and yield shift
	HoldPeriod  int               // value in months
	Iterations  int               // number of simulations above the model value. The total number of simulations with be double this value.
	Vertical    string            //
	Horizontal  string            //
	Count       int               // used in the template to increment the id of the sensitivity table rows
}

// used in the template to increment the id of the sensitivity table rows
func (s *SensitivityInput) Add() int {
	s.Count++
	return s.Count
}

// used in the template to reset Count
func (s *SensitivityInput) Reset() int {
	s.Count = 0
	return s.Count
}

type SensitivityItems struct {
	name           string                                            //
	setter         func(input float64, e *EntityModel)               // sets the value of the variable. initial value from *e, plus input
	value          float64                                           //
	originalValue  interface{}                                       //
	returnFloatVal func(s *SensitivityItems, e *EntityModel) float64 // returns the current value as a float64
	reset          func(s *SensitivityItems, e *EntityModel)         //
}

// SensitivityAnalysisCalc -
func (e *EntityModel) SensitivityAnalysisCalc() {
	sensitivityItems, count := e.CreateSensitivityItems()
	e.SensitivityOutputArray = make([][][]float64, int(math.Pow(float64(count), 2))-count)
	e.SensitivityOutputKeys = make(map[string]int)
	e.SensitivityOutputItems = make(map[int]string)
	firstValHeader := make(map[string]float64)
	// create correctly sized output array for IRR and EM. Necessary to do this in a separate step because if not done before calculations,
	// then one grid will be calculated before the inverse is created, and will result in nil pointer dereference
	index := 0
	for _, v := range sensitivityItems {
		firstValHeader[v.name] = v.returnFloatVal(&v, e)
		for _, vv := range sensitivityItems {
			if vv.name == v.name {
				continue
			}
			e.SensitivityOutputArray[index] = make([][]float64, e.SensitivityInput.Iterations*2+2)
			e.SensitivityOutputKeys[fmt.Sprint(v.name, "/", vv.name)] = index
			e.SensitivityOutputItems[index] = fmt.Sprint(v.name, "/", vv.name)
			index++
		}
	}
	// combination metric calcs. first checks if the opposite combination exists, if so exits, otherwise calculates grid. Also inserts single metric calcs into grid
	// Creates [v1/v2] as well as [v2/v1]
	singleVarIRR, _ := rowCalc(sensitivityItems, e, "all", 0)
	headerRow, _ := rowCalc(sensitivityItems, e, "header", 0)
	arrayIndex := 0 // used to assign a grid in SensitivityOutputArray
	for i, outerItem := range sensitivityItems {
		for ii, innerItem := range sensitivityItems {
			if outerItem.name == innerItem.name {
				continue
			}
			// 'i < ii' is in place to prevent calculating the inverse grid
			if i > ii {
				inverseKey := e.SensitivityOutputKeys[fmt.Sprint(innerItem.name, "/", outerItem.name)]
				currentKey := e.SensitivityOutputKeys[fmt.Sprint(outerItem.name, "/", innerItem.name)]
				for cIndex, cValue := range e.SensitivityOutputArray[inverseKey] {
					for rIndex, rValue := range cValue {
						if e.SensitivityOutputArray[currentKey][rIndex] == nil {
							e.SensitivityOutputArray[currentKey][rIndex] = make([]float64, e.SensitivityInput.Iterations*2+2)
						}
						e.SensitivityOutputArray[currentKey][rIndex][cIndex] = rValue
					}
				}
				arrayIndex++
				continue
			}
			outerItem.setter(outerItem.value*float64(-e.SensitivityInput.Iterations), e)
			column := 0
			for value := -e.SensitivityInput.Iterations - 1; value <= e.SensitivityInput.Iterations; value++ {
				irr := make(map[string][]float64)
				switch value {
				case -e.SensitivityInput.Iterations - 1:
					irr[innerItem.name] = headerRow[innerItem.name]
				case 0:
					irr[innerItem.name] = singleVarIRR[innerItem.name]
					outerItem.setter(outerItem.value, e)
				default:
					irr, _ = rowCalc(sensitivityItems, e, innerItem.name, outerItem.returnFloatVal(&outerItem, e))
					outerItem.setter(outerItem.value, e)
				}
				e.SensitivityOutputArray[arrayIndex][column] = irr[innerItem.name]
				column++
			}
			arrayIndex++
			outerItem.reset(&outerItem, e)
		}
	}
	e.UpdateEntityModel(false)
}

// calculates and returns []float64 for each entry in SensitivityItems
func rowCalc(sensitivityItems []SensitivityItems, e *EntityModel, item string, firstIndexVal float64) (IRR map[string][]float64, EM map[string][]float64) {
	IRR = make(map[string][]float64)
	EM = make(map[string][]float64)
	for _, v := range sensitivityItems {
		if item == v.name || item == "all" {
			rowIRR := make([]float64, e.SensitivityInput.Iterations*2+2)
			rowEM := make([]float64, e.SensitivityInput.Iterations*2+2)
			index := 0
			v.setter(v.value*float64(-e.SensitivityInput.Iterations), e)
			for ii := -e.SensitivityInput.Iterations - 1; ii <= e.SensitivityInput.Iterations; ii++ {
				if ii == -e.SensitivityInput.Iterations-1 {
					if item == "all" {
						// rowIRR[index] = v.returnFloatVal(v)
						// rowEM[index] = v.returnFloatVal(v)
					} else {
						rowIRR[index] = firstIndexVal
						rowEM[index] = firstIndexVal
					}
					index++
					continue
				}
				e.UpdateEntityModel(false)
				rowIRR[index] = e.Metrics.IRR.NetLeveredAfterTax
				rowEM[index] = e.Metrics.EM.NetLeveredAfterTax
				index++
				v.setter(v.value, e)
			}
			v.reset(&v, e)
			IRR[v.name] = rowIRR
			EM[v.name] = rowEM
			e.UpdateEntityModel(false)
		} else if item == "header" {
			val := v.returnFloatVal(&v, e) - v.value*float64(e.SensitivityInput.Iterations)
			rowIRR := make([]float64, e.SensitivityInput.Iterations*2+2)
			index := 1
			for i := -e.SensitivityInput.Iterations; i <= e.SensitivityInput.Iterations; i++ {
				rowIRR[index] = val
				val = val + v.value
				index++
			}
			IRR[v.name] = rowIRR
		}
	}
	return IRR, EM
}

// Returns array of SensitivityItems, which can then be looped over to calculate each sensitivity grid.
// Count is the number of variables being analyzed
// void and probability will not work properly until there is a way to assign the GLA down to units, currently this is done in the db query
// for holding period, it modifies the LoanEnd date for each "Disposition" loan, but not others
func (e *EntityModel) CreateSensitivityItems() (sensitivityItems []SensitivityItems, count int) {
	if e.SensitivityInput.Void != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{
			name:           "Void",
			setter:         func(input float64, e *EntityModel) { e.GLA.Void = e.GLA.Void + int(input) },
			value:          float64(e.SensitivityInput.Void),
			originalValue:  e.GLA.Void,
			returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return float64(ee.GLA.Void) },
			reset:          func(s *SensitivityItems, ee *EntityModel) { ee.GLA.Void = s.originalValue.(int) },
		})
	}
	if e.SensitivityInput.Probability != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{
			name:           "Renewal Probability",
			setter:         func(input float64, ee *EntityModel) { ee.GLA.Probability = ee.GLA.Probability + input },
			value:          e.SensitivityInput.Probability,
			originalValue:  e.GLA.Probability,
			returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.GLA.Probability },
			reset:          func(s *SensitivityItems, ee *EntityModel) { ee.GLA.Probability = s.originalValue.(float64) },
		})
	}
	for i := range e.SensitivityInput.Growth {
		if e.SensitivityInput.Growth[i].ShortTermRate != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{
				name: fmt.Sprint(i, " Short term rate"),
				setter: func(input float64, ee *EntityModel) {
					temp := ee.GrowthInput[i]
					temp.ShortTermRate = math.Round((temp.ShortTermRate+input)*10000) / 10000
					ee.GrowthInput[i] = temp
				},
				value:          e.SensitivityInput.Growth[i].ShortTermRate,
				originalValue:  e.GrowthInput[i].ShortTermRate,
				returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.GrowthInput[i].ShortTermRate },
				reset: func(s *SensitivityItems, ee *EntityModel) {
					temp := ee.GrowthInput[i]
					temp.ShortTermRate = s.originalValue.(float64)
					ee.GrowthInput[i] = temp
				},
			})
		}
		if e.SensitivityInput.Growth[i].LongTermRate != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{
				name: fmt.Sprint(i, " Long term rate"),
				setter: func(input float64, ee *EntityModel) {
					temp := ee.GrowthInput[i]
					temp.LongTermRate = temp.LongTermRate + input
					ee.GrowthInput[i] = temp
				},
				value:          e.SensitivityInput.Growth[i].LongTermRate,
				originalValue:  e.GrowthInput[i].LongTermRate,
				returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.GrowthInput[i].LongTermRate },
				reset: func(s *SensitivityItems, ee *EntityModel) {
					temp := ee.GrowthInput[i]
					temp.LongTermRate = s.originalValue.(float64)
					ee.GrowthInput[i] = temp
				},
			})
		}
	}
	for i, v := range e.SensitivityInput.EntityCosts {
		if e.SensitivityInput.EntityCosts[i].Amount != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{
				name: fmt.Sprint("Entity ", v.Name),
				setter: func(input float64, ee *EntityModel) {
					temp := ee.CostInput[i]
					temp.Amount = temp.Amount + input
					ee.CostInput[i] = temp
				},
				value:          e.SensitivityInput.EntityCosts[i].Amount,
				originalValue:  e.CostInput[i].Amount,
				returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.CostInput[i].Amount },
				reset: func(s *SensitivityItems, ee *EntityModel) {
					temp := ee.CostInput[i]
					temp.Amount = s.originalValue.(float64)
					ee.CostInput[i] = temp
				},
			})
		}
	}
	for i, v := range e.SensitivityInput.UnitCosts {
		if e.SensitivityInput.UnitCosts[i].Amount != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{
				name: fmt.Sprint("Unit ", v.Name),
				setter: func(input float64, ee *EntityModel) {
					temp := ee.GLA.CostInput[i]
					temp.Amount = temp.Amount + input
					ee.GLA.CostInput[i] = temp
				},
				value:          e.SensitivityInput.UnitCosts[i].Amount,
				originalValue:  e.GLA.CostInput[i].Amount,
				returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.GLA.CostInput[i].Amount },
				reset: func(s *SensitivityItems, ee *EntityModel) {
					temp := ee.GLA.CostInput[i]
					temp.Amount = s.originalValue.(float64)
					ee.GLA.CostInput[i] = temp
				},
			})
		}
	}
	for i, v := range e.SensitivityInput.Debt {
		if e.SensitivityInput.Debt[i].Spread != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{
				name: fmt.Sprint(v.Name),
				setter: func(input float64, ee *EntityModel) {
					temp := ee.DebtInput[i]
					temp.Spread = temp.Spread + int(input)
					ee.DebtInput[i] = temp
				},
				value:          float64(e.SensitivityInput.Debt[i].Spread),
				originalValue:  e.DebtInput[i].Spread,
				returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return float64(ee.DebtInput[i].Spread) },
				reset: func(s *SensitivityItems, ee *EntityModel) {
					temp := ee.DebtInput[i]
					temp.Spread = s.originalValue.(int)
					ee.DebtInput[i] = temp
				},
			})
		}
	}
	if e.SensitivityInput.Valuation.YieldShift != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{
			name: "Yield Shift",
			setter: func(input float64, ee *EntityModel) {
				ee.Valuation.YieldShift = math.Round((ee.Valuation.YieldShift+input)*10000) / 10000
			},
			value:          e.SensitivityInput.Valuation.YieldShift,
			originalValue:  e.Valuation.YieldShift,
			returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.Valuation.YieldShift },
			reset:          func(s *SensitivityItems, ee *EntityModel) { ee.Valuation.YieldShift = s.originalValue.(float64) },
		})
	}
	if e.SensitivityInput.Valuation.DiscountRate != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{
			name: "Discount Rate",
			setter: func(input float64, ee *EntityModel) {
				e.Valuation.DiscountRate = math.Round((ee.Valuation.DiscountRate+input)*10000) / 10000
			},
			value:          e.SensitivityInput.Valuation.DiscountRate,
			originalValue:  e.Valuation.DiscountRate,
			returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return ee.Valuation.DiscountRate },
			reset:          func(s *SensitivityItems, ee *EntityModel) { ee.Valuation.DiscountRate = s.originalValue.(float64) },
		})
	}
	if e.SensitivityInput.HoldPeriod != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{
			name: "Holding Period",
			setter: func(input float64, ee *EntityModel) {
				e.SalesDate = Dateadd(ee.SalesDate, int(input))
				for i := range ee.DebtInput {
					if ee.DebtInput[i].EndEvent == "Disposition" {
						ee.DebtInput[i].LoanEnd = Dateadd(ee.DebtInput[i].LoanEnd, int(input))
					}
				}
			},
			value:          float64(e.SensitivityInput.HoldPeriod),
			originalValue:  e.SalesDate,
			returnFloatVal: func(s *SensitivityItems, ee *EntityModel) float64 { return float64(ee.SalesDate.Dateint) },
			reset: func(s *SensitivityItems, ee *EntityModel) {
				ee.SalesDate = s.originalValue.(Datetype)
				for i := range e.DebtInput {
					ee.DebtInput[i].LoanEnd = s.originalValue.(Datetype)
				}
			},
		})
	}
	return sensitivityItems, count
}

// Get -
func (c *SensitivityController) Get() {
	temp := make(map[interface{}]interface{})
	key := EntityModelsList[GetStringSensitivity(c, "name")]
	if EntityModelsMap[key].EntityModel.SensitivityOutputArray != nil {
		vert := EntityModelsMap[key].EntityModel.SensitivityInput.Vertical
		hori := EntityModelsMap[key].EntityModel.SensitivityInput.Horizontal
		sKey := EntityModelsMap[key].EntityModel.SensitivityOutputKeys[fmt.Sprint(vert, "/", hori)]
		temp["grid"] = EntityModelsMap[key].EntityModel.SensitivityOutputArray[sKey]
	}
	temp["entity"] = EntityModelsMap[key].EntityModel
	c.TplName = "sensitivity.tpl"
	c.Data = temp
}

// Post -
func (c *SensitivityController) Post() {
	calc, _ := strconv.ParseBool(GetStringSensitivity(c, "calc"))
	key := EntityModelsList[GetStringSensitivity(c, "entityModel")]
	iterations := GetIntSensitivity(c, "iterations")
	if calc {
		growth := make(map[string]HModel)
		for i := range EntityModelsMap[key].EntityModel.GrowthInput {
			temp := HModel{
				ShortTermRate: GetFloatSensitivity(c, fmt.Sprint(i, " Short term rate")) / 100,
				LongTermRate:  GetFloatSensitivity(c, fmt.Sprint(i, " Long term rate")) / 100,
			}
			growth[i] = temp
		}
		entityCosts := make(map[int]CostInput)
		for i, v := range EntityModelsMap[key].EntityModel.CostInput {
			temp := CostInput{
				Amount: GetFloatSensitivity(c, fmt.Sprint("Entity ", v.Name)),
			}
			entityCosts[i] = temp
		}
		unitCosts := make(map[int]CostInput)
		for i, v := range EntityModelsMap[key].EntityModel.GLA.CostInput {
			temp := CostInput{
				Amount: GetFloatSensitivity(c, fmt.Sprint("Unit ", v.Name)),
			}
			unitCosts[i] = temp
		}
		debt := make([]DebtInput, len(EntityModelsMap[key].EntityModel.DebtInput))
		for i := range EntityModelsMap[key].EntityModel.DebtInput {
			debt[i].Spread = GetIntSensitivity(c, fmt.Sprint(debt[i].Name))
		}
		EntityModelsMap[key].EntityModel.SensitivityInput = SensitivityInput{
			Void:        GetIntSensitivity(c, "Void"),
			Probability: GetFloatSensitivity(c, "Renewal Probability") / 100,
			Growth:      growth,
			EntityCosts: entityCosts,
			UnitCosts:   unitCosts,
			Debt:        debt,
			Valuation:   Valuation{DiscountRate: GetFloatSensitivity(c, "Discount Rate") / 100, YieldShift: GetFloatSensitivity(c, "Yield Shift")},
			HoldPeriod:  GetIntSensitivity(c, "Hold Period"),
			Iterations:  iterations,
			Vertical:    GetStringSensitivity(c, "vertical"),
			Horizontal:  GetStringSensitivity(c, "horizontal"),
		}
		EntityModelsMap[key].EntityModel.SensitivityAnalysisCalc()
	}
	sKey := EntityModelsMap[key].EntityModel.SensitivityOutputKeys[fmt.Sprint(EntityModelsMap[key].EntityModel.SensitivityInput.Vertical, "/", EntityModelsMap[key].EntityModel.SensitivityInput.Horizontal)]
	temp := make(map[interface{}]interface{})
	temp["grid"] = EntityModelsMap[key].EntityModel.SensitivityOutputArray[sKey]
	temp["entity"] = EntityModelsMap[key].EntityModel
	c.TplName = "sensitivity.tpl"
	c.Data = temp
}

// SensitivityController -
type SensitivityController struct {
	beego.Controller
}

// GetStringSensitivity -
func GetStringSensitivity(c *SensitivityController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetFloatSensitivity -
func GetFloatSensitivity(c *SensitivityController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// GetIntSensitivity -
func GetIntSensitivity(c *SensitivityController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}
