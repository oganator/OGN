[[define "EntitySettings"]]
<div class="d-flex flex-row">
    <div class="container-fluid shadow-lg rounded d-flex flex-column" style="width: 55%" id="settingsTable">
        <div class="p-2">
            <div class="tabs-wrapper d-flex flex-row">
                <ul class="nav nav-tabs tabPinned d-flex align-content-start flex-wrap" role="tablist">
                    <li class="nav-item">
                        [[if eq .tab "settings"]]<a class="nav-link waves-light active" data-toggle="tab" href="#settings" role="tab" ng-click="updateSettingsTab('settings')">Settings</a>[[end]]
                        [[if ne .tab "settings"]]<a class="nav-link waves-light" data-toggle="tab" href="#settings" role="tab" ng-click="updateSettingsTab('settings')">Settings</a>[[end]]
                    </li>
                    <li class="nav-item">
                        [[if eq .tab "leasing"]]<a class="nav-link waves-light active" data-toggle="tab" href="#leasing" role="tab" ng-click="updateSettingsTab('leasing')">Leasing</a>[[end]]
                        [[if ne .tab "leasing"]]<a class="nav-link waves-light" data-toggle="tab" href="#leasing" role="tab" ng-click="updateSettingsTab('leasing')">Leasing</a>[[end]]
                    </li>
                    <li class="nav-item">
                        [[if eq .tab "capex"]]<a class="nav-link waves-light active" data-toggle="tab" href="#capex" role="tab" ng-click="updateSettingsTab('capex')">Capex</a>[[end]]
                        [[if ne .tab "capex"]]<a class="nav-link waves-light" data-toggle="tab" href="#capex" role="tab" ng-click="updateSettingsTab('capex')">Capex</a>[[end]]
                    </li>                    
                    <li class="nav-item">
                        [[if eq .tab "financing"]]<a class="nav-link waves-light active" data-toggle="tab" href="#financing" role="tab" ng-click="updateSettingsTab('financing')">Financing</a>[[end]]
                        [[if ne .tab "financing"]]<a class="nav-link waves-light" data-toggle="tab" href="#financing" role="tab" ng-click="updateSettingsTab('financing')">Financing</a>[[end]]
                    </li>
                    <li class="nav-item">
                        [[if eq .tab "ervcpi"]]<a class="nav-link waves-light active" data-toggle="tab" href="#ervcpi" role="tab" ng-click="updateSettingsTab('ervcpi')">ERV/CPI Growth</a>[[end]]
                        [[if ne .tab "ervcpi"]]<a class="nav-link waves-light" data-toggle="tab" href="#ervcpi" role="tab" ng-click="updateSettingsTab('ervcpi')">ERV/CPI Growth</a>[[end]]
                    </li>
                    <li class="nav-item">
                        [[if eq .tab "valuation"]]<a class="nav-link waves-light active" data-toggle="tab" href="#val" role="tab" ng-click="updateSettingsTab('valuation')">Valuation</a>[[end]]
                        [[if ne .tab "valuation"]]<a class="nav-link waves-light" data-toggle="tab" href="#val" role="tab" ng-click="updateSettingsTab('valuation')">Valuation</a>[[end]]
                    </li>
                    <li class="nav-item">
                        [[if eq .tab "tax"]]<a class="nav-link waves-light active" data-toggle="tab" href="#tax" role="tab" ng-click="updateSettingsTab('tax')">Tax</a>[[end]]
                        [[if ne .tab "tax"]]<a class="nav-link waves-light" data-toggle="tab" href="#tax" role="tab" ng-click="updateSettingsTab('tax')">Tax</a>[[end]]
                    </li>
                </ul>
            </div>
            <div class="tab-content">
                [[if eq .tab "settings"]]<div class="tab-pane fade active show" id="settings" role="tabpanel">[[end]]
                [[if ne .tab "settings"]]<div class="tab-pane fade" id="settings" role="tabpanel">[[end]]
                    [[template "EntitySettingsDefault" .]]
                </div>
                [[if eq .tab "leasing"]]<div class="tab-pane fade active show" id="leasing" role="tabpanel">[[end]]
                [[if ne .tab "leasing"]]<div class="tab-pane fade" id="leasing" role="tabpanel">[[end]]
                    [[template "EntitySettingsLeasing" .]]
                </div>
                [[if eq .tab "capex"]]<div class="tab-pane fade active show" id="capex" role="tabpanel">[[end]]
                [[if ne .tab "capex"]]<div class="tab-pane fade" id="capex" role="tabpanel">[[end]]
                    [[template "Capex" .entity]]
                </div>                
                [[if eq .tab "financing"]]<div class="tab-pane fade active show" id="financing" role="tabpanel">[[end]]
                [[if ne .tab "financing"]]<div class="tab-pane fade" id="financing" role="tabpanel">[[end]]
                    [[template "EntitySettingsFinancing" .]]
                </div>
                [[if eq .tab "ervcpi"]]<div class="tab-pane fade active show" id="ervcpi" role="tabpanel">[[end]]
                [[if ne .tab "ervcpi"]]<div class="tab-pane fade" id="ervcpi" role="tabpanel">[[end]]
                    [[template "EntitySettingsERVCPI" .]]
                </div>
                [[if eq .tab "valuation"]]<div class="tab-pane fade active show" id="val" role="tabpanel">[[end]]
                [[if ne .tab "valuation"]]<div class="tab-pane  fade" id="val" role="tabpanel">[[end]]
                    [[template "EntitySettingsValuation" .]]
                </div>
                [[if eq .tab "tax"]]<div class="tab-pane fade active show" id="tax" role="tabpanel">[[end]]
                [[if ne .tab "tax"]]<div class="tab-pane fade" id="tax" role="tabpanel">[[end]]
                    [[template "EntitySettingsTax" .]]
                </div>
            </div>
        </div>
        <div class="mt-auto p-2">
            <table class="grid d-flex " cellspacing="0">
                <tbody>
                    <tr>
                        <td colspan="2">
                            <div class="col-md-4 d-flex align-items-end">
                                <div class="form-group ">
                                    <button ng-click="updateEntity('[[.entity.Name]]')" class="btn secondary-bg">Submit</button>
                                </div>
                            </div>
                            <div class="form-group row-sm-2">
                                <label for="sims">Simulations</label>
                                <input type="text" class="form-control" id="sims" name="sims" Value="[[.entity.MCSetup.Sims]]">
                            </div>
                        </td>
                        <td colspan="10">
                            <table class="table tableFixHead rounded">
                                <thead>
                                    <tr>
                                        [[if eq .entity.Strategy "Standard"]]
                                        <th id="settings_irr" iscope="col" style="width: 12.5%">IRR</th>
                                        <th id="settings_em" scope="col" style="width: 12.5%">EM</th>
                                        [[end]]
                                        [[if eq .entity.Strategy "Balloon"]]
                                        <th id="settings_ytm" scope="col" style="width: 12.5%">YTM</th>
                                        <th id="settings_dur" scope="col" style="width: 12.5%">DUR</th>
                                        <th id="settings_ytmdur" scope="col" style="width: 12.5%">YTM/DUR</th>
                                        [[end]]
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                    [[if eq .entity.Strategy "Standard"]]
                                        <td id="settings_tr_irr">{{[[.entity.Metrics.IRR.NetLeveredAfterTax]] | number:2}}</td>
                                        <td id="settings_tr_em">{{[[.entity.Metrics.EM.NetLeveredAfterTax]] | number:2}}</td>
                                        [[end]]
                                        [[if eq .entity.Strategy "Balloon"]]
                                        <td id="settings_tr_ytm">{{[[.entity.Metrics.BondHolder.YTM]] | number:2}}</td>
                                        <td id="settings_tr_dur">{{[[.entity.Metrics.BondHolder.Duration]] | number:2}}</td>
                                        <td id="settings_tr_ytmdur">{{[[.entity.Metrics.BondHolder.YTMDUR]] | number:2}}</td>
                                        [[end]]
                                    </tr>
                                </tbody>
                            </table>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <div class="container-fluid shadow-lg rounded" style="width: 35%;" id="settingsGraph">
       <div bind-html-compile = settingsChart></div>
    </div>
</div>
    <br>
    <br>
[[end]]

[[define "EntitySettingsDefault"]]

    <div class="d-flex flex-row">
        <input hidden id="name" name="name" ng-model="entity" value="[[.entity.Name]]">
        <div hidden class="form-group ">
            <label hidden for="portfolio">Asset</label>
            <input hidden type="text" class="form-control" id="portfolio" name="portfolio" value="[[.entity.Name]]" >
        </div>
<!--    <div class="form-group ">
            <label for="holdperiod">Hold Period</label>
            <input type="text" class="form-control" id="holdperiod" name="holdperiod" value="[[.entity.HoldPeriod]]" data-toggle="tooltip" data-placement="top" title="Years">
        </div>
-->
        <div class="form-group ">
            <label for="startdate">Start Date</label>
            <select type="text" class="form-control" id="startmonth" value="[[.entity.StartDate.MonthName]]">
                <option hidden>[[.entity.StartDate.MonthName]]</option>
                <option>Jan</option>
                <option>Feb</option>
                <option>Mar</option>
                <option>Apr</option>
                <option>May</option>
                <option>Jun</option>
                <option>Jul</option>
                <option>Aug</option>
                <option>Sep</option>
                <option>Oct</option>
                <option>Nov</option>
                <option>Dec</option>
            </select>
            <input type="number" class="form-control" id="startyear" value="[[.entity.StartDate.Year]]" >
        </div>
        <div class="form-group ">
            <label for="salesdate">Sales Date</label>
            <select type="text" class="form-control" id="salesmonth" value="[[.entity.SalesDate.MonthName]]">
                <option hidden>[[.entity.SalesDate.MonthName]]</option>
                <option>Jan</option>
                <option>Feb</option>
                <option>Mar</option>
                <option>Apr</option>
                <option>May</option>
                <option>Jun</option>
                <option>Jul</option>
                <option>Aug</option>
                <option>Sep</option>
                <option>Oct</option>
                <option>Nov</option>
                <option>Dec</option>
            </select>
            <input type="number" class="form-control" id="salesyear" value="[[.entity.SalesDate.Year]]" >
        </div>
        <div class="form-group ">
            <label for="rate">Strategy</label>
            <select type="text" class="form-control" id="strategy" name="strategy" value="[[.entity.Strategy]]">
                <option hidden>[[.entity.Strategy]]</option>
                <option>Standard</option>
                <option>Balloon</option>
                <option>Pure Discount</option>
            </select>
        </div>
        [[if eq .entity.Strategy "Balloon"]]
        <div class="form-group " >
            <label for="discount">Discount Rate</label>
            <input type="text" class="form-control" id="settings_discount" name="settings_discount" Value="{{[[.entity.GLA.DiscountRate]] * 100 | number:2}}">
        </div>
        <div class="form-group " >
            <label for="rate">Rent to Sell %</label>
            <input type="text" class="form-control" id="settings_soldrent" name="settings_soldrent" Value="{{[[.entity.GLA.PercentSoldRent]] * 100 | number:2}}">
        </div>
        <div class="form-group " >
            <label for="balpercent">Balloon Percent</label>
            <input type="text" class="form-control" id="settings_balpercent" name="settings_balpercent" Value="{{[[.entity.BalloonPercent]] * 100 | number:2}}">
        </div>
        [[end]]
    </div>
[[end]]

[[define "EntitySettingsERVCPI"]]
    <div class="d-flex flex-row">
        [[$erv := index .entity.GrowthInput "ERV"]]
        <div class="form-group ">
            <label for="ervshorttermrate">ERV Short Rate</label>
            <input type="text" class="form-control" id="ervshorttermrate" name="ervshorttermrate" Value="{{[[$erv.ShortTermRate]] * 100 | number:2}}"></input>
            <input type="text" class="form-control" id="ervshorttermratesigma" name="ervshorttermratesigma" Value="{{[[.entity.MCSetup.ERV.ShortTermRate]] * 100 | number:2}}"></input>
            </div>
        <div class="form-group ">
            <label for="ervshorttermperiod">ERV Short Period</label>
            <input type="text" class="form-control" id="ervshorttermperiod" name="ervshorttermperiod" Value="[[$erv.ShortTermPeriod]]" data-toggle="tooltip" data-placement="top" title="in Months"></input>
            <input type="text" class="form-control" id="ervshorttermperiodsigma" name="ervshorttermperiodsigma" Value="[[.entity.MCSetup.ERV.ShortTermPeriod]]"></input>
        </div>
        <div class="form-group ">
            <label for="ervtransitionperiod">ERV Transition Period</label>
            <input type="text" class="form-control" id="ervtransitionperiod" name="ervtransitionperiod" Value="[[$erv.TransitionPeriod]]"></input>
            <input type="text" class="form-control" id="ervtransitionperiodsigma" name="ervtransitionperiodsigma" Value="[[.entity.MCSetup.ERV.TransitionPeriod]]"></input>
        </div>
        <div class="form-group ">
            <label for="ervlongtermrate">ERV Long Rate</label>
            <input type="text" class="form-control" id="ervlongtermrate" name="ervlongtermrate" Value="{{[[$erv.LongTermRate]] * 100 | number:2}}"></input>
            <input type="text" class="form-control" id="ervlongtermratesigma" name="ervlongtermratesigma" Value="{{[[.entity.MCSetup.ERV.LongTermRate]] * 100 | number:2}}"></input>            
        </div>
        [[$cpi := index .entity.GrowthInput "CPI"]]
        <div class="form-group ">
            <label for="cpishorttermrate">CPI Short Rate</label>
            <input type="text" class="form-control" id="cpishorttermrate" name="cpishorttermrate" Value="{{[[$cpi.ShortTermRate]] * 100 | number:2}}"></input>
            <input type="text" class="form-control" id="cpishorttermratesigma" name="cpishorttermratesigma" Value="{{[[.entity.MCSetup.CPI.ShortTermRate]] * 100 | number:2}}"></input>
            </div>
        <div class="form-group ">
            <label for="cpishorttermperiod">CPI Short Period</label>
            <input type="text" class="form-control" id="cpishorttermperiod" name="cpishorttermperiod" Value="[[$cpi.ShortTermPeriod]]"></input>
            <input type="text" class="form-control" id="cpishorttermperiodsigma" name="cpishorttermperiodsigma" Value="[[.entity.MCSetup.CPI.ShortTermPeriod]]"></input>
        </div>
        <div class="form-group ">
            <label for="cpitransitionperiod">CPI Transition Period</label>
            <input type="text" class="form-control" id="cpitransitionperiod" name="cpitransitionperiod" Value="[[$cpi.TransitionPeriod]]"></input>
            <input type="text" class="form-control" id="cpitransitionperiodsigma" name="cpitransitionperiodsigma" Value="[[.entity.MCSetup.CPI.TransitionPeriod]]"></input>
        </div>
        <div class="form-group">
            <label for="cpilongtermrate">CPI Long Rate</label>
            <input type="text" class="form-control" id="cpilongtermrate" name="cpilongtermrate" Value="{{[[$cpi.LongTermRate]] * 100 | number:2}}"></input>
            <input type="text" class="form-control" id="cpilongtermratesigma" name="cpilongtermratesigma" Value="{{[[.entity.MCSetup.CPI.LongTermRate]] * 100 | number:2}}"></input>
        </div>
    </div>
    <br>
    <br>
[[end]]

[[define "EntitySettingsTax"]]
    <div class="form-column">
        <div class="d-flex flex-row">
            <div class="form-group ">
                <label for="rett">RETT</label>
                <input type="text" class="form-control" id="rett" name="rett" Value="{{[[.entity.Tax.RETT]] * 100 | number:2}}">
                </div>
            <div class="form-group ">
                <label for="landvalue">Land Value</label>
                <input type="text" class="form-control" id="landvalue" name="landvalue" Value="{{[[.entity.Tax.LandValue]] * 100 | number:2}}">
            </div>
            <div class="form-group ">
                <label for="minvalue">Min Depreciable Value</label>
                <input type="text" class="form-control" id="minvalue" name="minvalue" Value="{{[[.entity.Tax.MinValue]] * 100 | number:2}}">
            </div>
            <div class="form-group ">
                <label for="usableperiod">Depreciation Period</label>
                <input type="text" class="form-control" id="usableperiod" name="usableperiod" Value="[[.entity.Tax.UsablePeriod]]">
            </div>
        </div>
        <div class="d-flex flex-row">
            <div class="form-group ">
                <label for="vat">VAT</label>
                <input type="text" class="form-control" id="vat" name="vat" Value="{{[[.entity.Tax.VAT]] * 100 | number:2}}">
                </div>
            <div class="form-group ">
                <label for="carrybackyrs">Carry Back Years</label>
                <input type="text" class="form-control" id="carrybackyrs" name="carrybackyrs" Value="[[.entity.Tax.CarryBackYrs]]">
            </div>
            <div class="form-group ">
                <label for="carryforwardyrs">Carry Forward Years</label>
                <input type="text" class="form-control" id="carryforwardyrs" name="carryforwardyrs" Value="[[.entity.Tax.CarryForwardYrs]]">
            </div>
        </div>
    </div>
[[end]]

[[define "EntitySettingsLeasing"]]
    <div class="form-column">
        <div class="d-flex flex-row">
            <div class="form-group ">
                <label for="void">Void</label>
                <input type="text" class="form-control" id="void" name="void" Value="[[.entity.GLA.Void]]"></input>
                <input type="text" class="form-control" id="voidsigma" name="voidsigma" Value="[[.entity.MCSetup.Void]]"></input>
            </div>
            <div class="form-group ">
                <label for="duration">Extension Duration</label>
                <input type="text" class="form-control" id="duration" name="duration" Value="[[.entity.GLA.EXTDuration]]"></input>
            </div>
            <div class="form-group ">
                <label for="rentrevision">Rent Revision ERV</label>
                <input type="text" class="form-control" id="rentrevision" name="rentrevision" Value="{{[[.entity.GLA.RentRevisionERV]] * 100 | number:2}}"></input>
            </div>
            <div class="form-group ">
                <label for="probability">Probability</label>
                <input type="text" class="form-control" id="probability" name="probability" Value="{{[[.entity.GLA.Probability]] * 100 | number:2}}"></input>
                <input type="text" class="form-control" id="probabilitysigma" name="probabilitysigma" Value="{{[[.entity.MCSetup.Probability]] * 100 | number:2}}"></input>
            </div>
            <div class="form-group ">
                <label for="hazard">Hazard Rate</label>
                <input type="text" class="form-control" id="hazard" name="hazard" Value="{{[[.entity.GLA.Default.Hazard]] * 100 | number:2}}"></input>
                <input type="text" class="form-control" id="hazardsigma" name="hazardsigma" Value="{{[[.entity.MCSetup.Hazard]] * 100 | number:2}}"></input>
            </div>
        </div>
    </div>
    <br><br>
    <div class="column">
        [[template "Capex" .entity.GLA]]
    </div>
[[end]]

[[define "EntitySettingsFinancing"]]
    <table class="table table-hover tableFixHead rounded" id="loansTable">
        <thead>
            <tr>
                <th hidden scope="col">masterID</th>
                <th scope="col">Amount</th>
                <th scope="col">Rate</th>
                <th scope="col">Interest Type</th>
                <th scope="col">Loan Type</th>
                <th scope="col">Start</th>
                <th scope="col">End</th>
                <th scope="col">Floating Rate Basis</th>
                <th scope="col">Spread</th>
                <th scope="col">Amortization Period</th>
                <th scope="col">Active</th>
            </tr>
        </thead>
        <tbody>
            [[$growthItems := .entity.GrowthInput]]
            [[range .entity.DebtInput]]
            <tr name="row_loan" id="row_loan_[[.MasterID]]">
                <td hidden style="border-color:006A4D;" width="7%"><input type="text" class="form-control" id="loan_masterID_[[.MasterID]]" value="[[.MasterID]]"></input></td>
                <td style="border-color:006A4D;" width="7%"><input type="text" class="form-control" id="loan_amount_[[.MasterID]]" value="{{[[.Amount]] | number:0}}"></input></td>                    
                <td style="border-color:006A4D;" width="7%"><input type="text" class="form-control" id="interestrate_[[.MasterID]]" value="{{[[.InterestRate]] * 100 | number:2 }}"></input></td>
                <td style="border-color:006A4D;" width="7%">
                    <select type="text" class="form-control" id="interest_type_[[.MasterID]]" value="[[.InterestType]]">
                        <option hidden>[[.InterestType]]</option>
                        <option>Fixed</option>
                        <option>Floating</option>
                    </select>
                </td>
                <td style="border-color:006A4D;" width="10%">
                    <select type="text" class="form-control" id="loan_type_[[.MasterID]]" value="[[.LoanType]]">
                        <option hidden>[[.LoanType]]</option>
                        <option>Amortizing</option>
                        <option>Interest Only</option>
                    </select>
                </td>
<!--                <td style="border-color:006A4D;" width="8%">
                    <select type="text" class="form-control" id="loan_basis_[[.MasterID]]" value="[[.LoanBasis]]">
                        <option hidden>[[.LoanBasis]]</option>
                        <option>Market Value</option>
                        <option>Capex</option>
                        <option>Amount</option>
                        <option>Revaluation</option>
                    </select> -->
                </td>
                <td style="border-color:006A4D;" width="11%">
                    <div class="d-flex flex-row ">
                        <input type="number" class="form-control" id="loan_start_month_[[.MasterID]]" value="[[.LoanStart.Month]]"></input>
                        <input type="number" class="form-control" id="loan_start_year_[[.MasterID]]" value="[[.LoanStart.Year]]"></input>
                    </div>
                </td>
                <td style="border-color:006A4D;" width="11%">
                    <div class="d-flex flex-row ">
                        <input type="number" class="form-control" id="loan_end_month_[[.MasterID]]" value="[[.LoanEnd.Month]]"></input>
                        <input type="number" class="form-control" id="loan_end_year_[[.MasterID]]" value="[[.LoanEnd.Year]]"></input>
                    </div>
                </td>
                <td style="border-color:006A4D;" width="7%"> 
                    <select type="text" class="form-control" id="float_basis_[[.MasterID]]" value="[[.FloatBasis]]">
                        <option hidden>[[.FloatBasis]]</option>
                        [[range $k,$v := $growthItems]]
                        <option>[[$k]]</option>
                        [[end]]
                    </select>
                </td>
                <td style="border-color:006A4D;" width="7%"><input type="number" id="spread_[[.MasterID]]" class="form-control" value="[[.Spread]]"></input></td>
                <td style="border-color:006A4D;" width="5%"><input type="number" id="amortization_period_[[.MasterID]]" class="form-control" value="[[.AmortizationPeriod]]"></input></td>
                <td style="border-color:006A4D;" width="7%"><input type="checkbox" [[if .Active]] checked [[end]] id="active_[[.MasterID]]" class="form-control"></input></td>
            </tr>
            [[end]]
        </tbody>
    </table>
[[end]]

[[define "EntitySettingsValuation"]]
        <div class="form-row">
            <div class="form-group ">
                <label for="entryyield">Acquisition Price</label>
                <input type="text" class="form-control" id="settingsacqprice" name="settingsacqprice" value="{{[[.entity.Valuation.AcqPrice]] | number:0}}">
            </div>
            <div class="form-group ">
                <label for="valuationmethod">Method</label>
                <select type="text" class="form-control" id="valuationmethod" value="[[.entity.Valuation.Method]]">
                    <option hidden>[[.entity.Valuation.Method]]</option>
                    <option>DCF</option>
                    <option>DirectCap</option>
                    <option>German</option>
                    <option>VPV</option>
                </select>
            </div>
            <div class="form-group ">
                <label for="entryyield">Entry Yield</label>
                <input type="text" class="form-control" id="entryyield" name="entryyield" value="{{[[.entity.Valuation.EntryYield]] * 100 | number:2}}">
            </div>
            <div class="form-group ">
                <label for="exityield">Exit Yield</label>
                <input type="text" class="form-control" id="exityield" name="exityield" value="{{[[.entity.Valuation.ExitYield]] * 100 | number:2}}">
            </div>
            <div class="form-group ">
                <label for="yieldshift">Yield Shift (bps)</label>
                <input type="text" class="form-control" id="yieldshift" name="yieldshift" value="[[.entity.Valuation.YieldShift]]"></input>
                <input type="text" class="form-control" id="yieldshiftsigma" name="yieldshiftsigma" value="[[.entity.MCSetup.YieldShift]]"></input>   
            </div>
            [[if eq .entity.Valuation.Method "DCF"]]
            <div class="form-group ">
                <label for="entitydiscountrate">Discount Rate</label>
                <input type="text" class="form-control" id="entitydiscountrate" name="entitydiscountrate" value="{{[[.entity.Valuation.DiscountRate]] * 100 | number:2}}">
            </div>
            [[end]]
        </div>
[[end]]

[[define "FundSettings"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%" id="settingsTable">
        <form class="form-horizontal form-well" role="form" >
            <div class="tabs-wrapper">
                <ul class="nav nav-tabs tabPinned" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link waves-light active" data-toggle="tab" href="#fundsettings" role="tab">Settings</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link waves-light" data-toggle="tab" href="#fundleasing" role="tab">&nbsp</a>
                    </li>
                </ul>
            </div>
            <div class="tab-content">
                <div class="tab-pane fade in show active" id="fundsettings" role="tabpanel">
                    [[template "FundSettingsDefault" .]]
                </div>
                <div class="tab-pane fade" id="fundleasing" role="tabpanel">
                </div>
            </div>
            <div class="form-row"> 
                <div>
                    <div class="form-group col-md-4 d-flex align-items-end">
                        <button ng-click="updateEntity('[[.entity.Name]]')" class="btn secondary-bg">Submit</button>
                    </div>
                </div>
                <div class="form-group ">
                    <label for="sims">Simulations</label>
                    <input type="text" class="form-control" id="sims" name="sims" Value="[[.entity.MCSetup.Sims]]">
                </div>
                <div class="form-group ">
                    <input readonly type="text" class="form-control" id="settings_irr" name="settings_irr" Value="IRR: [[.entity.Metrics.IRR.NetLeveredAfterTax]]">
                </div>
                <div class="form-group ">
                    <input readonly type="text" class="form-control" id="settings_em" name="settings_em" Value="EM: [[.entity.Metrics.EM.NetLeveredAfterTax]]">
                </div>
                [[if eq .entity.Strategy "Balloon"]]
                <div class="form-group  offset-sm-5">
                    <input readonly type="text" class="form-control" id="settings_ytm" name="settings_ytm" Value="YTM: [[.entity.Metrics.BondHolder.YTM]]">
                </div>
                <div class="form-group ">
                    <input readonly type="text" class="form-control" id="settings_dur" name="settings_dur" Value="DUR: [[.entity.Metrics.BondHolder.Duration]]">
                </div>
                <div class="form-group ">
                    <input readonly type="text" class="form-control" id="settings_ytmdur" name="settings_ytmdur" Value="YTM/DUR: {{[[.entity.Metrics.BondHolder.YTMDUR]] | number:4}}">
                </div>
                [[end]]
            </div>
        </form>
    </div>
    <br>
    <br>
[[end]]

[[define "FundSettingsDefault"]]                  
    <div class="form-row">
        <input hidden id="name" name="name" ng-model="entity" value="[[.entity.Name]]">
        <div hidden class="form-group ">
            <label for="portfolio">Asset</label>
            <input type="text" class="form-control" id="portfolio" name="portfolio" value="[[.entity.Name]]">
        </div>
        <div class="form-group ">
            <label for="startdate">Start Date</label>
            <select disabled type="text" class="form-control" id="startmonth" value="[[.entity.StartDate.MonthName]]">
                <option hidden>[[.entity.StartDate.MonthName]]</option>
            </select>
            <input disabled type="number" class="form-control" id="startyear" value="[[.entity.StartDate.Year]]" >
        </div>
        <div class="form-group ">
            <label for="salesdate">Sales Date</label>
            <select disabled type="text" class="form-control" id="salesmonth" value="[[.entity.SalesDate.MonthName]]">
                <option hidden>[[.entity.SalesDate.MonthName]]</option>
            </select>
            <input disabled type="number" class="form-control" id="salesyear" value="[[.entity.SalesDate.Year]]" >
        </div>
        <div class="form-group ">
            <label for="rate">Strategy</label>
            <select type="text" class="form-control" id="strategy" name="strategy" value="[[.entity.Strategy]]" onchange="strategyChange()">
                <option hidden>[[.entity.Strategy]]</option>
                <option>Standard</option>
                <option>Balloon</option>
                <option>Pure Discount</option>
            </select>
        </div>
    </div>
[[end]]


[[define "Settings"]]
    [[if (gt .entity.ParentID 0)]]
        <div class="container-fluid">
            [[template "EntitySettings" .]]
        </div>
    [[end]]
    [[if (eq .entity.ParentID 0)]]
        <div class="container-fluid">   
            [[template "FundSettings" .]]
        </div>
    [[end]]	
[[end]]

[[template "Settings" .]]