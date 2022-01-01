package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/astaxie/beego"
)

// MonteCarlo -
func (e *Entity) AzureMonteCarlo() {
	duration := e.MCDataObjectsCreate(1)
	temp := CreateShellEntity(e)
	AzureChannel := make(chan MCResultSlice)
	go e.AzureSimReceive(AzureChannel, duration) //receives http responses from Azure function
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ee *Entity) {
			defer wg.Done()
			AzureSimSend(&temp, EntityDataStore[e.MasterID], AzureChannel)
		}(e)
	}
	wg.Wait()
	e.MCCalc(duration)
}

func AzureSimSend(e *Entity, tempdata *EntityData, ch chan MCResultSlice) {
	e.EntityData = *tempdata
	postBody, err := json.Marshal(e)
	if err != nil {
		fmt.Println("AzureSimSend Error 1: ", err)
	}
	responseBody := bytes.NewBuffer(postBody)
	resp, err2 := http.Post(AzureURL, "application/json", responseBody)
	if err2 != nil {
		fmt.Println("AzureSimSend Error 2: ", err2)
	}
	defer resp.Body.Close()
	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("AzureSimSend Error 3: ", err3)
	}
	tempresults := MCResultSlice{}
	err4 := json.Unmarshal(body, &tempresults)
	if err4 != nil {
		fmt.Println("AzureSimSend Error 4: ", err4)
	}
	ch <- tempresults
}

func (e *Entity) AzureSimReceive(ch chan MCResultSlice, duration int) {
	instance := 0
	eIndex := 0
	vSize := e.MCSetup.Sims / 100
	for v := range ch {
		instance++
		for vIndex := 0; vIndex < vSize; vIndex++ {
			growth := map[string]HModel{
				"CPI": {EndingIndex: v.CPI[vIndex]},
				"ERV": {EndingIndex: v.ERV[vIndex]},
			}
			e.MCSlice[eIndex] = &Entity{
				Metrics: Metrics{
					IRR: ReturnType{
						GrossUnleveredBeforeTax: 0,
						NetLeveredAfterTax:      v.IRR[vIndex],
					},
					EM: ReturnType{
						GrossUnleveredBeforeTax: 0,
						NetLeveredAfterTax:      v.EM[vIndex],
					},
					BondHolder: BondReturnType{
						Duration: v.Duration[vIndex],
						YTM:      v.YTM[vIndex],
						YTMDUR:   v.YTMDUR[vIndex],
					},
				},
				ParentID:    e.MasterID,
				Parent:      e,
				StartDate:   Datetype{},
				HoldPeriod:  0,
				SalesDate:   Datetype{},
				EndDate:     Datetype{},
				GrowthInput: growth,
				DebtInput:   DebtInput{},
				OpEx:        CostInput{PercentOfTRI: v.OpEx[vIndex]},
				Fees:        CostInput{},
				Capex:       map[int]CostInput{},
				GLA: Unit{
					Probability: v.Probability[vIndex],
					Void:        int(v.Void[vIndex]),
					Default: Default{
						NumberOfDefaults: int(v.NumberOfDefaults[vIndex]),
						Hazard:           v.Hazard[vIndex],
					},
				},
				MC:             false,
				MCSetup:        MCSetup{},
				MCSlice:        []*Entity{},
				MCResultSlice:  MCResultSlice{},
				MCResults:      MCResults{},
				FactorAnalysis: []FactorIndependant{},
				Tax:            Tax{},
				COA:            map[int]FloatCOA{},
				Valuation: Valuation{
					YieldShift: v.YieldShift[vIndex],
				},
				TableHeader:    HeaderType{},
				Table:          []TableJSON{},
				Strategy:       "",
				UOM:            "",
				BalloonPercent: 0,
			}
			e.MCResults.EndCash.Mean = v.EndCash[vIndex]
			e.MCResults.EndNCF.Mean = v.EndNCF[vIndex]
			e.MCResults.EndMarketValue.Mean = v.EndMarketValue[vIndex]
			e.MCResultSlice.EndCash[eIndex] = v.EndCash[vIndex]
			e.MCResultSlice.EndNCF[eIndex] = v.EndNCF[vIndex]
			e.MCResultSlice.EndMarketValue[eIndex] = v.EndMarketValue[vIndex]
			e.MCResultSlice.IRR[eIndex] = v.IRR[vIndex]
			e.MCResultSlice.EM[eIndex] = v.EM[vIndex]
			e.MCResultSlice.Void[eIndex] = v.Void[vIndex]
			e.MCResultSlice.Probability[eIndex] = v.Probability[vIndex]
			e.MCResultSlice.NumberOfDefaults[eIndex] = v.NumberOfDefaults[vIndex]
			e.MCResultSlice.OpEx[eIndex] = v.OpEx[vIndex]
			e.MCResultSlice.CPI[eIndex] = v.CPI[vIndex]
			e.MCResultSlice.ERV[eIndex] = v.ERV[vIndex]
			e.MCResultSlice.Hazard[eIndex] = v.Hazard[vIndex]
			e.MCResultSlice.YieldShift[eIndex] = v.YieldShift[vIndex]
			if e.Strategy != "Standard" {
				e.MCResultSlice.Duration[eIndex] = v.Duration[vIndex]
				e.MCResultSlice.YTMDUR[eIndex] = v.YTMDUR[vIndex]
				e.MCResultSlice.YTM[eIndex] = v.YTM[vIndex]
			}
			for i := 0; i < duration; i++ {
				e.MCResultSlice.CashBalance[i][eIndex] = v.CashBalance[i][vIndex]
				e.MCResultSlice.NCF[i][eIndex] = v.NCF[i][vIndex]
				e.MCResultSlice.MarketValue[i][eIndex] = v.MarketValue[i][vIndex]
			}
			eIndex++
		} // range vIndex
	} // range channel
} // AzureSimReceive

func CreateShellEntity(e *Entity) Entity {
	temp := Entity{
		Mutex:         &sync.Mutex{},
		MasterID:      e.MasterID,
		Name:          e.Name,
		ChildEntities: map[int]*Entity{},
		ChildUnits:    map[int]*Unit{},
		Metrics:       Metrics{},
		ParentID:      e.ParentID,
		Parent:        e.Parent,
		StartDate:     Dateadd(e.StartDate, 0),
		HoldPeriod:    e.HoldPeriod,
		SalesDate:     Dateadd(e.SalesDate, 0),
		EndDate:       Dateadd(e.EndDate, 0),
		GrowthInput:   e.GrowthInput,
		Growth:        map[string]map[int]float64{},
		DebtInput:     e.DebtInput,
		OpEx:          e.OpEx,
		Fees:          e.Fees,
		Capex:         e.Capex,
		GLA:           e.GLA,
		MC:            true,
		MCSetup:       e.MCSetup,
		// MCSlice:        []*Entity{},
		MCResultSlice:  MCResultSlice{},
		MCResults:      MCResults{},
		FactorAnalysis: []FactorIndependant{},
		Tax:            e.Tax,
		COA:            map[int]FloatCOA{},
		Valuation:      e.Valuation,
		TableHeader:    HeaderType{},
		Table:          []TableJSON{},
		Strategy:       e.Strategy,
		UOM:            "",
		BalloonPercent: e.BalloonPercent,
	}
	temp.PopulateUnits()
	return temp
}

//////////////////////////////////////////////////////////////////////////////////
// AZURE

type FunctionController struct {
	beego.Controller
}

func (c *FunctionController) Post() {
	tempentity := Entity{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &tempentity)
	if err != nil {
		fmt.Println("Azure Function Controller Error: ", err)
	}
	tempentity.MCSetup.Sims = tempentity.MCSetup.Sims / 100
	tempentity.MonteCarlo("Azure")
	SimCounter.Mutex.Lock()
	tempentity.MCResultSlice.SimNumber = SimCounter.ID
	SimCounter.ID++
	SimCounter.Mutex.Unlock()
	c.Data["json"] = tempentity.MCResultSlice
	c.ServeJSON()
	tempentity = Entity{}
	go debug.FreeOSMemory()
}
