[[define "EntitySettings"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%" id="settingsTable">
        <form class="form-horizontal form-well" role="form" >
            <div class="tabs-wrapper">
                <ul class="nav nav-tabs tabPinned" role="tablist">
                    <li class="nav-item">
                        [[if eq .tab "settings"]]<a class="nav-link waves-light active" data-toggle="tab" href="#settings" role="tab" ng-click="updateSettingsTab('settings')">Settings</a>[[end]]
                        [[if ne .tab "settings"]]<a class="nav-link waves-light" data-toggle="tab" href="#settings" role="tab" ng-click="updateSettingsTab('settings')">Settings</a>[[end]]
                    </li>
                    <li class="nav-item">
                        [[if eq .tab "leasing"]]<a class="nav-link waves-light active" data-toggle="tab" href="#leasing" role="tab" ng-click="updateSettingsTab('leasing')">Leasing</a>[[end]]
                        [[if ne .tab "leasing"]]<a class="nav-link waves-light" data-toggle="tab" href="#leasing" role="tab" ng-click="updateSettingsTab('leasing')">Leasing</a>[[end]]
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
            
            <table class="grid" cellspacing="0">
                <tbody>
                    <tr>
                        <td colspan="2">
                            <div>
                                <div class="form-group col-md-4 d-flex align-items-end">
                                    <button ng-click="updateEntity('[[.entity.Name]]')" class="btn secondary-bg">Submit</button>
                                </div>
                            </div>
                            <div class="form-group col-sm-2">
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

        </form>
    </div>
    <br>
    <br>
[[end]]

[[define "EntitySettingsDefault"]]

    <div class="form-row">
        <input hidden id="name" name="name" ng-model="entity" value="[[.entity.Name]]">
        <div hidden class="form-group col-sm-1">
            <label hidden for="portfolio">Asset</label>
            <input hidden type="text" class="form-control" id="portfolio" name="portfolio" value="[[.entity.Name]]" >
        </div>
<!--    <div class="form-group col-sm-1">
            <label for="holdperiod">Hold Period</label>
            <input type="text" class="form-control" id="holdperiod" name="holdperiod" value="[[.entity.HoldPeriod]]" data-toggle="tooltip" data-placement="top" title="Years">
        </div>
-->
        <div class="form-group col-sm-1">
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
        <div class="form-group col-sm-1">
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
        <div class="form-group col-sm-1">
            <label for="settings_fees">Fees (bps)</label>
            <input type="text" class="form-control" id="settings_fees" name="fees" Value="[[.entity.Fees.PercentOfGAV]]" data-toggle="tooltip" data-placement="top" title="Percent of Gross Market Value">
        </div>
        <div class="form-group col-sm-1">
            <label for="rate">Strategy</label>
            <select type="text" class="form-control" id="strategy" name="strategy" value="[[.entity.Strategy]]">
                <option hidden>[[.entity.Strategy]]</option>
                <option>Standard</option>
                <option>Balloon</option>
                <option>Pure Discount</option>
            </select>
        </div>
        [[if eq .entity.Strategy "Standard"]]
        <div class="form-group col-sm-1" >
            <label for="ltv">LTV</label>
            <input type="text" class="form-control" id="settings_ltv" name="settings_ltv" Value="{{[[.entity.DebtInput.LTV]] * 100 | number:2}}">
        </div>
        <div class="form-group col-sm-1" >
            <label for="rate">Loan Rate</label>
            <input type="text" class="form-control" id="settings_rate" name="settings_rate" Value="{{[[.entity.DebtInput.InterestRate]] * 100 | number:2}}">
        </div>
        [[end]]
        [[if eq .entity.Strategy "Balloon"]]
        <div class="form-group col-sm-1" >
            <label for="discount">Discount Rate</label>
            <input type="text" class="form-control" id="settings_discount" name="settings_discount" Value="{{[[.entity.GLA.DiscountRate]] * 100 | number:2}}">
        </div>
        <div class="form-group col-sm-1" >
            <label for="rate">Rent to Sell %</label>
            <input type="text" class="form-control" id="settings_soldrent" name="settings_soldrent" Value="{{[[.entity.GLA.PercentSoldRent]] * 100 | number:2}}">
        </div>
        <div class="form-group col-sm-1" >
            <label for="balpercent">Balloon Percent</label>
            <input type="text" class="form-control" id="settings_balpercent" name="settings_balpercent" Value="{{[[.entity.BalloonPercent]] * 100 | number:2}}">
        </div>
        [[end]]
    </div>
    <div class="form-row">

        <div class="form-group col-sm-1">
            <label for="opex">Operating Expenses</label>
            <input type="text" class="form-control" id="settings_opex" name="settings_opex" Value="{{[[.entity.OpEx.PercentOfTRI]] * 100 | number:2}}">
        </div>
    </div>
    <div class="form-row">
        <div class="form-group col-sm-1">
            <input type="text" class="form-control" id="settings_opexsigma" name="settings_opexsigma" Value="{{[[.entity.MCSetup.OpEx]] * 100 | number:2}}">
        </div>
    </div>
[[end]]

[[define "EntitySettingsERVCPI"]]
    <div class="form-row">
        [[$erv := index .entity.GrowthInput "ERV"]]
        <div class="form-group col-sm-1">
            <label for="ervshorttermrate">ERV Short Rate</label>
            <input type="text" class="form-control" id="ervshorttermrate" name="ervshorttermrate" Value="{{[[$erv.ShortTermRate]] * 100 | number:2}}">
            </div>
        <div class="form-group col-sm-1">
            <label for="ervshorttermperiod">ERV Short Period</label>
            <input type="text" class="form-control" id="ervshorttermperiod" name="ervshorttermperiod" Value="[[$erv.ShortTermPeriod]]" data-toggle="tooltip" data-placement="top" title="in Months">
        </div>
        <div class="form-group col-sm-1">
            <label for="ervtransitionperiod">ERV Transition Period</label>
            <input type="text" class="form-control" id="ervtransitionperiod" name="ervtransitionperiod" Value="[[$erv.TransitionPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="ervlongtermrate">ERV Long Rate</label>
            <input type="text" class="form-control" id="ervlongtermrate" name="ervlongtermrate" Value="{{[[$erv.LongTermRate]] * 100 | number:2}}">
        </div>
        [[$cpi := index .entity.GrowthInput "CPI"]]
        <div class="form-group col-sm-1">
            <label for="cpishorttermrate">CPI Short Rate</label>
            <input type="text" class="form-control" id="cpishorttermrate" name="cpishorttermrate" Value="{{[[$cpi.ShortTermRate]] * 100 | number:2}}">
            </div>
        <div class="form-group col-sm-1">
            <label for="cpishorttermperiod">CPI Short Period</label>
            <input type="text" class="form-control" id="cpishorttermperiod" name="cpishorttermperiod" Value="[[$cpi.ShortTermPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="cpitransitionperiod">CPI Transition Period</label>
            <input type="text" class="form-control" id="cpitransitionperiod" name="cpitransitionperiod" Value="[[$cpi.TransitionPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="cpilongtermrate">CPI Long Rate</label>
            <input type="text" class="form-control" id="cpilongtermrate" name="cpilongtermrate" Value="{{[[$cpi.LongTermRate]] * 100 | number:2}}">
        </div>
    </div>
    <div class="form-row">
        <div class="form-group col-sm-1">
            <label for="ervshorttermratesigma"></label>
            <input type="text" class="form-control" id="ervshorttermratesigma" name="ervshorttermratesigma" Value="{{[[.entity.MCSetup.ERV.ShortTermRate]] * 100 | number:2}}">
            </div>
        <div class="form-group col-sm-1">
            <label for="ervshorttermperiodsigma"></label>
            <input type="text" class="form-control" id="ervshorttermperiodsigma" name="ervshorttermperiodsigma" Value="[[.entity.MCSetup.ERV.ShortTermPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="ervtransitionperiodsigma"></label>
            <input type="text" class="form-control" id="ervtransitionperiodsigma" name="ervtransitionperiodsigma" Value="[[.entity.MCSetup.ERV.TransitionPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="ervlongtermratesigma"></label>
            <input type="text" class="form-control" id="ervlongtermratesigma" name="ervlongtermratesigma" Value="{{[[.entity.MCSetup.ERV.LongTermRate]] * 100 | number:2}}">
        </div>
        <div class="form-group col-sm-1">
            <label for="cpishorttermratesigma"></label>
            <input type="text" class="form-control" id="cpishorttermratesigma" name="cpishorttermratesigma" Value="{{[[.entity.MCSetup.CPI.ShortTermRate]] * 100 | number:2}}">
            </div>
        <div class="form-group col-sm-1">
            <label for="cpishorttermperiodsigma"></label>
            <input type="text" class="form-control" id="cpishorttermperiodsigma" name="cpishorttermperiodsigma" Value="[[.entity.MCSetup.CPI.ShortTermPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="cpitransitionperiodsigma"></label>
            <input type="text" class="form-control" id="cpitransitionperiodsigma" name="cpitransitionperiodsigma" Value="[[.entity.MCSetup.CPI.TransitionPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="cpilongtermratesigma"></label>
            <input type="text" class="form-control" id="cpilongtermratesigma" name="cpilongtermratesigma" Value="{{[[.entity.MCSetup.CPI.LongTermRate]] * 100 | number:2}}">
        </div>
    </div>
    <br>
    <br>
[[end]]

[[define "EntitySettingsTax"]]
    <div class="form-column">
        <div class="form-row">
            <div class="form-group col-sm-1">
                <label for="rett">RETT</label>
                <input type="text" class="form-control" id="rett" name="rett" Value="{{[[.entity.Tax.RETT]] * 100 | number:2}}">
                </div>
            <div class="form-group col-sm-1">
                <label for="landvalue">Land Value</label>
                <input type="text" class="form-control" id="landvalue" name="landvalue" Value="{{[[.entity.Tax.LandValue]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="minvalue">Min Depreciable Value</label>
                <input type="text" class="form-control" id="minvalue" name="minvalue" Value="{{[[.entity.Tax.MinValue]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="usableperiod">Depreciation Period</label>
                <input type="text" class="form-control" id="usableperiod" name="usableperiod" Value="[[.entity.Tax.UsablePeriod]]">
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-sm-1">
                <label for="vat">VAT</label>
                <input type="text" class="form-control" id="vat" name="vat" Value="{{[[.entity.Tax.VAT]] * 100 | number:2}}">
                </div>
            <div class="form-group col-sm-1">
                <label for="carrybackyrs">Carry Back Years</label>
                <input type="text" class="form-control" id="carrybackyrs" name="carrybackyrs" Value="[[.entity.Tax.CarryBackYrs]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="carryforwardyrs">Carry Forward Years</label>
                <input type="text" class="form-control" id="carryforwardyrs" name="carryforwardyrs" Value="[[.entity.Tax.CarryForwardYrs]]">
            </div>
        </div>
    </div>
[[end]]

[[define "EntitySettingsLeasing"]]
    <div class="form-column">
        <div class="form-row">
            <div class="form-group col-sm-1">
                <label for="void">Void</label>
                <input type="text" class="form-control" id="void" name="void" Value="[[.entity.GLA.Void]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="duration">Extension Duration</label>
                <input type="text" class="form-control" id="duration" name="duration" Value="[[.entity.GLA.EXTDuration]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="rentrevision">Rent Revision ERV</label>
                <input type="text" class="form-control" id="rentrevision" name="rentrevision" Value="{{[[.entity.GLA.RentRevisionERV]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="probability">Probability</label>
                <input type="text" class="form-control" id="probability" name="probability" Value="{{[[.entity.GLA.Probability]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="incentivemonths">Incentive Months</label>
                <input type="text" class="form-control" id="incentivemonths" name="incentivemonths" Value="[[.entity.GLA.RentIncentives.Duration]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="incentivepercent">Incentive Percent</label>
                <input type="text" class="form-control" id="incentivepercent" name="incentivepercent" Value="{{[[.entity.GLA.RentIncentives.PercentOfContractRent]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="fitoutcosts">Fit Out Costs/sqft</label>
                <input type="text" class="form-control" id="fitoutcosts" name="fitoutcosts" Value="[[.entity.GLA.FitOutCosts.AmountPerTotalArea]]">
            </div>                        
            <div class="form-group col-sm-1">
                <label for="hazard">Hazard Rate</label>
                <input type="text" class="form-control" id="hazard" name="hazard" Value="{{[[.entity.GLA.Default.Hazard]] * 100 | number:2}}">
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="voidsigma" name="voidsigma" Value="[[.entity.MCSetup.Void]]">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="probabilitysigma" name="probabilitysigma" Value="{{[[.entity.MCSetup.Probability]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>            
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="hazardsigma" name="hazardsigma" Value="{{[[.entity.MCSetup.Hazard]] * 100 | number:2}}">
            </div>
        </div>
    </div>
[[end]]

[[define "EntitySettingsValuation"]]
    <div class="form-column">
        <div class="form-row">
            <div class="form-group col-sm-1">
                <label for="valuationmethod">Method</label>
                <select type="text" class="form-control" id="valuationmethod" value="[[.entity.Valuation.Method]]">
                    <option hidden>[[.entity.Valuation.Method]]</option>
                    <option>DCF</option>
                    <option>DirectCap</option>
                    <option>German</option>
                    <option>VPV</option>
                </select>
            </div>
            <div class="form-group col-sm-1">
                <label for="entryyield">Entry Yield</label>
                <input type="text" class="form-control" id="entryyield" name="entryyield" value="{{[[.entity.Valuation.EntryYield]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="exityield">Exit Yield</label>
                <input type="text" class="form-control" id="exityield" name="exityield" value="{{[[.entity.Valuation.ExitYield]] * 100 | number:2}}">
            </div>
            <div class="form-group col-sm-1">
                <label for="yieldshift">Yield Shift (bps)</label>
                <input type="text" class="form-control" id="yieldshift" name="yieldshift" value="[[.entity.Valuation.YieldShift]]">
            </div>
            [[if eq .entity.Valuation.Method "DCF"]]
            <div class="form-group col-sm-1">
                <label for="entitydiscountrate">Discount Rate</label>
                <input type="text" class="form-control" id="entitydiscountrate" name="entitydiscountrate" value="{{[[.entity.Valuation.DiscountRate]] * 100 | number:2}}">
            </div>
            [[end]]
        </div>
        <div class="form-row">
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="yieldshiftsigma" name="yieldshiftsigma" value="[[.entity.MCSetup.YieldShift]]">
            </div>    
        </div>
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
                <div class="form-group col-sm-1">
                    <label for="sims">Simulations</label>
                    <input type="text" class="form-control" id="sims" name="sims" Value="[[.entity.MCSetup.Sims]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="settings_irr" name="settings_irr" Value="IRR: [[.entity.Metrics.IRR.NetLeveredAfterTax]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="settings_em" name="settings_em" Value="EM: [[.entity.Metrics.EM.NetLeveredAfterTax]]">
                </div>
                [[if eq .entity.Strategy "Balloon"]]
                <div class="form-group col-sm-1 offset-sm-5">
                    <input readonly type="text" class="form-control" id="settings_ytm" name="settings_ytm" Value="YTM: [[.entity.Metrics.BondHolder.YTM]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="settings_dur" name="settings_dur" Value="DUR: [[.entity.Metrics.BondHolder.Duration]]">
                </div>
                <div class="form-group col-sm-1">
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
        <div hidden class="form-group col-sm-1">
            <label for="portfolio">Asset</label>
            <input type="text" class="form-control" id="portfolio" name="portfolio" value="[[.entity.Name]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="startdate">Start Date</label>
            <select disabled type="text" class="form-control" id="startmonth" value="[[.entity.StartDate.MonthName]]">
                <option hidden>[[.entity.StartDate.MonthName]]</option>
            </select>
            <input disabled type="number" class="form-control" id="startyear" value="[[.entity.StartDate.Year]]" >
        </div>
        <div class="form-group col-sm-1">
            <label for="salesdate">Sales Date</label>
            <select disabled type="text" class="form-control" id="salesmonth" value="[[.entity.SalesDate.MonthName]]">
                <option hidden>[[.entity.SalesDate.MonthName]]</option>
            </select>
            <input disabled type="number" class="form-control" id="salesyear" value="[[.entity.SalesDate.Year]]" >
        </div>
        <div class="form-group col-sm-1">
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