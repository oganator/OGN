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
	Relative    bool              // if true, the table is formatted to show relative values up/down from their observed value in the model instead of their actual values
	Iterations  int               // number of simulations above the model value. The total number of simulations with be double this value.
}

type SensitivityItems struct {
	name           string                                   //
	setter         func(input float64, e *EntityModel)      // sets the value of the variable. initial value from *e, plus input
	value          float64                                  //
	originalValue  interface{}                              //
	returnFloatVal func(s SensitivityItems) float64         // returns the current value as a float64
	reset          func(s SensitivityItems, e *EntityModel) //
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
		firstValHeader[v.name] = v.returnFloatVal(v)
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
			outerItem.setter(outerItem.value*float64(-e.SensitivityInput.Iterations), e)
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
					irr, _ = rowCalc(sensitivityItems, e, innerItem.name, outerItem.returnFloatVal(outerItem))
					outerItem.setter(outerItem.value, e)
				}
				e.SensitivityOutputArray[arrayIndex][column] = irr[innerItem.name]
				column++
			}
			arrayIndex++
			outerItem.reset(outerItem, e)
		}
	}
	for i, v := range e.SensitivityOutputArray {
		fmt.Println(e.SensitivityOutputItems[i])
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
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
			v.reset(v, e)
			IRR[v.name] = rowIRR
			EM[v.name] = rowEM
			e.UpdateEntityModel(false)
		} else if item == "header" {
			val := v.returnFloatVal(v) - v.value*float64(e.SensitivityInput.Iterations)
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
		sensitivityItems = append(sensitivityItems, SensitivityItems{name: "Void",
			setter:         func(input float64, e *EntityModel) { e.GLA.Void = e.GLA.Void + int(input) },
			value:          float64(e.SensitivityInput.Void),
			originalValue:  e.GLA.Void,
			returnFloatVal: func(s SensitivityItems) float64 { return float64(e.GLA.Void) },
			reset:          func(s SensitivityItems, e *EntityModel) { e.GLA.Void = s.originalValue.(int) },
		})
	}
	if e.SensitivityInput.Probability != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{name: "Renewal Probability",
			setter:         func(input float64, e *EntityModel) { e.GLA.Probability = e.GLA.Probability + input },
			value:          e.SensitivityInput.Probability,
			originalValue:  e.GLA.Probability,
			returnFloatVal: func(s SensitivityItems) float64 { return e.GLA.Probability },
			reset:          func(s SensitivityItems, e *EntityModel) { e.GLA.Probability = s.originalValue.(float64) },
		})
	}
	for i := range e.SensitivityInput.Growth {
		if e.SensitivityInput.Growth[i].ShortTermRate != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{name: fmt.Sprint(i, ": Short term rate"),
				setter: func(input float64, e *EntityModel) {
					temp := e.GrowthInput[i]
					temp.ShortTermRate = math.Round((temp.ShortTermRate+input)*10000) / 10000
					e.GrowthInput[i] = temp
				},
				value:          e.SensitivityInput.Growth[i].ShortTermRate,
				originalValue:  e.GrowthInput[i].ShortTermRate,
				returnFloatVal: func(s SensitivityItems) float64 { return e.GrowthInput[i].ShortTermRate },
				reset: func(s SensitivityItems, e *EntityModel) {
					temp := e.GrowthInput[i]
					temp.ShortTermRate = s.originalValue.(float64)
					e.GrowthInput[i] = temp
				},
			})
		}
		if e.SensitivityInput.Growth[i].LongTermRate != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{name: fmt.Sprint(i, ": Long term rate"),
				setter: func(input float64, e *EntityModel) {
					temp := e.GrowthInput[i]
					temp.LongTermRate = temp.LongTermRate + input
					e.GrowthInput[i] = temp
				},
				value:          e.SensitivityInput.Growth[i].LongTermRate,
				originalValue:  e.GrowthInput[i].LongTermRate,
				returnFloatVal: func(s SensitivityItems) float64 { return e.GrowthInput[i].LongTermRate },
				reset: func(s SensitivityItems, e *EntityModel) {
					temp := e.GrowthInput[i]
					temp.LongTermRate = s.originalValue.(float64)
					e.GrowthInput[i] = temp
				},
			})
		}
	}
	for i, v := range e.SensitivityInput.EntityCosts {
		if e.SensitivityInput.EntityCosts[i].Amount != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{name: fmt.Sprint("Cost: ", v.MasterID, " ", v.Name),
				setter: func(input float64, e *EntityModel) {
					temp := e.CostInput[i]
					temp.Amount = temp.Amount + input
					e.CostInput[i] = temp
				},
				value:          e.SensitivityInput.EntityCosts[i].Amount,
				originalValue:  e.CostInput[i].Amount,
				returnFloatVal: func(s SensitivityItems) float64 { return e.CostInput[i].Amount },
				reset: func(s SensitivityItems, e *EntityModel) {
					temp := e.CostInput[i]
					temp.Amount = s.originalValue.(float64)
					e.CostInput[i] = temp
				},
			})
		}
	}
	for i, v := range e.SensitivityInput.UnitCosts {
		if e.SensitivityInput.UnitCosts[i].Amount != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{name: fmt.Sprint("Cost: ", v.MasterID, " ", v.Name),
				setter: func(input float64, e *EntityModel) {
					temp := e.GLA.CostInput[i]
					temp.Amount = temp.Amount + input
					e.GLA.CostInput[i] = temp
				},
				value:          e.SensitivityInput.UnitCosts[i].Amount,
				originalValue:  e.GLA.CostInput[i].Amount,
				returnFloatVal: func(s SensitivityItems) float64 { return e.GLA.CostInput[i].Amount },
				reset: func(s SensitivityItems, e *EntityModel) {
					temp := e.GLA.CostInput[i]
					temp.Amount = s.originalValue.(float64)
					e.GLA.CostInput[i] = temp
				},
			})
		}
	}
	for i, v := range e.SensitivityInput.Debt {
		if e.SensitivityInput.Debt[i].Spread != 0 {
			count++
			sensitivityItems = append(sensitivityItems, SensitivityItems{name: fmt.Sprint("Loan: ", v.MasterID, " ", v.Name),
				setter: func(input float64, e *EntityModel) {
					temp := e.DebtInput[i]
					temp.Spread = temp.Spread + int(input)
					e.DebtInput[i] = temp
				},
				value:          float64(e.SensitivityInput.Debt[i].Spread),
				originalValue:  e.DebtInput[i].Spread,
				returnFloatVal: func(s SensitivityItems) float64 { return float64(e.DebtInput[i].Spread) },
				reset: func(s SensitivityItems, e *EntityModel) {
					temp := e.DebtInput[i]
					temp.Spread = s.originalValue.(int)
					e.DebtInput[i] = temp
				},
			})
		}
	}
	if e.SensitivityInput.Valuation.YieldShift != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{name: "Yield Shift",
			setter: func(input float64, e *EntityModel) {
				e.Valuation.YieldShift = math.Round((e.Valuation.YieldShift+input)*10000) / 10000
			},
			value:          e.SensitivityInput.Valuation.YieldShift,
			originalValue:  e.Valuation.YieldShift,
			returnFloatVal: func(s SensitivityItems) float64 { return e.Valuation.YieldShift },
			reset:          func(s SensitivityItems, e *EntityModel) { e.Valuation.YieldShift = s.originalValue.(float64) },
		})
	}
	if e.SensitivityInput.Valuation.DiscountRate != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{name: "Discount Rate",
			setter: func(input float64, e *EntityModel) {
				e.Valuation.DiscountRate = math.Round((e.Valuation.DiscountRate+input)*10000) / 10000
			},
			value:          e.SensitivityInput.Valuation.DiscountRate,
			originalValue:  e.Valuation.DiscountRate,
			returnFloatVal: func(s SensitivityItems) float64 { return e.Valuation.DiscountRate },
			reset:          func(s SensitivityItems, e *EntityModel) { e.Valuation.DiscountRate = s.originalValue.(float64) },
		})
	}
	if e.SensitivityInput.HoldPeriod != 0 {
		count++
		sensitivityItems = append(sensitivityItems, SensitivityItems{name: "Holding Period",
			setter: func(input float64, e *EntityModel) {
				e.SalesDate = Dateadd(e.SalesDate, int(input))
				for i := range e.DebtInput {
					if e.DebtInput[i].EndEvent == "Disposition" {
						e.DebtInput[i].LoanEnd = Dateadd(e.DebtInput[i].LoanEnd, int(input))
					}
				}
			},
			value:          float64(e.SensitivityInput.HoldPeriod),
			originalValue:  e.SalesDate,
			returnFloatVal: func(s SensitivityItems) float64 { return float64(e.SalesDate.Dateint) },
			reset: func(s SensitivityItems, e *EntityModel) {
				e.SalesDate = s.originalValue.(Datetype)
				for i := range e.DebtInput {
					e.DebtInput[i].LoanEnd = s.originalValue.(Datetype)
				}
			},
		})
	}
	return sensitivityItems, count
}

// Post -
func (c *SensitivityController) Post() {
	calc, _ := strconv.ParseBool(GetStringSensitivity(c, "calc"))
	key := EntityModelsList[GetStringSensitivity(c, "name")]
	if calc {
		cpi := HModel{ShortTermRate: GetFloatSensitivity(c, "cpiShortRate"), LongTermRate: GetFloatSensitivity(c, "cpiLongRate")}
		iterations := GetIntSensitivity(c, "iterations")
		growthMap := map[string]HModel{}
		growthMap["CPI"] = cpi
		EntityModelsMap[key].EntityModel.SensitivityInput = SensitivityInput{
			Iterations: iterations,
			// HoldPeriod:  GetIntSensitivity(c, "holdPeriod"),
			Void:        GetIntSensitivity(c, "void"),
			Probability: GetFloatSensitivity(c, "prob"),
			Valuation:   Valuation{DiscountRate: GetFloatSensitivity(c, "discountRate"), YieldShift: GetFloatSensitivity(c, "yieldShift")},
			Growth:      growthMap,
		}
		EntityModelsMap[key].EntityModel.SensitivityAnalysisCalc()
	}
	var1 := GetStringSensitivity(c, "var1")
	var2 := GetStringSensitivity(c, "var2")
	sKey := EntityModelsMap[key].EntityModel.SensitivityOutputKeys[fmt.Sprint(var1, "/", var2)]
	temp := make(map[interface{}]interface{})
	temp["grid"] = EntityModelsMap[key].EntityModel.SensitivityOutputArray[sKey]
	c.TplName = "sensitivity.tpl"
	c.Data = temp
	// c.Data["json"] = EntityModelsMap[key].EntityModel.SensitivityOutputArray[sKey]
	// c.ServeJSON()
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
