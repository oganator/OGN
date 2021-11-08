[[define "EntitySettings"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%" id="settingsTable">
        <form class="form-horizontal form-well" role="form" method="post" action="/ViewEntity">
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
            <div class="form-row"> 
                <div>
                    <div class="form-group col-md-4 d-flex align-items-end">
                        <button type="submit" class="btn secondary-bg">Submit</button>
                    </div>
                </div>
                <div class="form-group col-sm-1">
                    <label for="fees">Simulations</label>
                    <input type="text" class="form-control" id="sims" name="sims" Value="[[.entity.MCSetup.Sims]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="irr" name="irr" Value="IRR: [[.entity.Metrics.IRR.NetLeveredAfterTax]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="em" name="em" Value="EM: [[.entity.Metrics.EM.NetLeveredAfterTax]]">
                </div>
                <div class="form-group col-sm-1 offset-sm-5">
                    <input readonly type="text" class="form-control" id="ytm" name="ytm" Value="YTM: [[.entity.Metrics.BondHolder.YTM]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="dur" name="dur" Value="DUR: [[.entity.Metrics.BondHolder.Duration]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="ytmdur" name="ytmdur" Value="YTM/DUR: {{[[.entity.Metrics.BondHolder.YTMDUR]] | number:4}}">
                </div>
            </div>
        </form>
    </div>
    <br>
    <br>
    <script>
        window.onload = strategyChange();
        
        function strategyChange() {
            var input = document.getElementById("strategy");
            if (input.value == "Standard"){
                document.getElementById("ytm").style.visibility = 'hidden';
                document.getElementById("dur").style.visibility = 'hidden';
                document.getElementById("ytmdur").style.visibility = 'hidden';
                document.getElementById("ltv").style.visibility = 'visible';
                document.getElementById("rate").style.visibility = 'visible';
                document.getElementById("discount").style.visibility = 'hidden';
                document.getElementById("soldrent").style.visibility = 'hidden';
                document.getElementById("balpercent").style.visibility = 'hidden';
                document.getElementById("durationtabletab").style.visibility = 'hidden';
                document.getElementById("ytmtabletab").style.visibility = 'hidden';
                document.getElementById("ytmdurtabletab").style.visibility = 'hidden';
                document.getElementById("irrtabletab").style.visibility = 'visible';
                document.getElementById("emtabletab").style.visibility = 'visible';
            }else{
                document.getElementById("ytm").style.visibility = 'visible';
                document.getElementById("dur").style.visibility = 'visible';
                document.getElementById("ytmdur").style.visibility = 'visible';
                document.getElementById("ltv").style.visibility = 'hidden';
                document.getElementById("rate").style.visibility = 'hidden';
                document.getElementById("discount").style.visibility = 'visible';
                document.getElementById("soldrent").style.visibility = 'visible';
                document.getElementById("balpercent").style.visibility = 'visible';
                document.getElementById("durationtabletab").style.visibility = 'visible';
                document.getElementById("ytmtabletab").style.visibility = 'visible';
                document.getElementById("ytmdurtabletab").style.visibility = 'visible';
                document.getElementById("irrtabletab").style.visibility = 'hidden';
                document.getElementById("emtabletab").style.visibility = 'hidden';
            }
        };
    </script>
[[end]]

[[define "EntitySettingsDefault"]]
    <div class="form-row">
        <input hidden id="name" name="name" ng-model="entity" value="[[.entity.Name]]">
        <div class="form-group col-sm-1">
            <label for="portfolio">Asset</label>
            <input type="text" class="form-control" id="portfolio" name="portfolio" value="[[.entity.Name]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="holdperiod">Hold Period</label>
            <input type="text" class="form-control" id="holdperiod" name="holdperiod" value="[[.entity.HoldPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="fees">Fees (bps)</label>
            <input type="text" class="form-control" id="fees" name="fees" Value="[[.entity.Fees.PercentOfGAV]]">
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
        <div class="form-group col-sm-1" id="ltv">
            <label for="ltv">LTV</label>
            <input type="text" class="form-control" name="ltv" Value="[[.entity.DebtInput.LTV]]">
        </div>
        <div class="form-group col-sm-1" id="rate">
            <label for="rate">Loan Rate</label>
            <input type="text" class="form-control" name="rate" Value="[[.entity.DebtInput.InterestRate]]">
        </div>
        <div class="form-group col-sm-1" id="discount">
            <label for="discount">Discount Rate</label>
            <input type="text" class="form-control" name="discount" Value="[[.entity.GLA.DiscountRate]]">
        </div>
        <div class="form-group col-sm-1" id="soldrent">
            <label for="rate">Rent to Sell %</label>
            <input type="text" class="form-control" name="soldrent" Value="[[.entity.GLA.PercentSoldRent]]">
        </div>
        <div class="form-group col-sm-1" id="balpercent">
            <label for="balpercent">Balloon Percent</label>
            <input type="text" class="form-control" name="balpercent" Value="[[.entity.BalloonPercent]]" >
        </div>
    </div>
    <div class="form-row">

        <div class="form-group col-sm-1">
            <label for="opex">Operating Expenses</label>
            <input type="text" class="form-control" id="opex" name="opex" Value="[[.entity.OpEx.PercentOfTRI]]">
        </div>
    </div>
    <div class="form-row">
        <div class="form-group col-sm-1">
            <input type="text" class="form-control" id="opexsigma" name="opexsigma" Value="[[.entity.MCSetup.OpEx]]">
        </div>
    </div>
[[end]]

[[define "EntitySettingsERVCPI"]]
    <div class="form-row">
        [[$erv := index .entity.GrowthInput "ERV"]]
        <div class="form-group col-sm-1">
            <label for="ervshorttermrate">ERV Short Rate</label>
            <input type="text" class="form-control" id="ervshorttermrate" name="ervshorttermrate" Value='[[$erv.ShortTermRate]]'>
            </div>
        <div class="form-group col-sm-1">
            <label for="ervshorttermperiod">ERV Short Period</label>
            <input type="text" class="form-control" id="ervshorttermperiod" name="ervshorttermperiod" Value='[[$erv.ShortTermPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="ervtransitionperiod">ERV Transition Period</label>
            <input type="text" class="form-control" id="ervtransitionperiod" name="ervtransitionperiod" Value='[[$erv.TransitionPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="ervlongtermrate">ERV Long Rate</label>
            <input type="text" class="form-control" id="ervlongtermrate" name="ervlongtermrate" Value='[[$erv.LongTermRate]]'>
        </div>
        [[$cpi := index .entity.GrowthInput "CPI"]]
        <div class="form-group col-sm-1">
            <label for="cpishorttermrate">CPI Short Rate</label>
            <input type="text" class="form-control" id="cpishorttermrate" name="cpishorttermrate" Value='[[$cpi.ShortTermRate]]'>
            </div>
        <div class="form-group col-sm-1">
            <label for="cpishorttermperiod">CPI Short Period</label>
            <input type="text" class="form-control" id="cpishorttermperiod" name="cpishorttermperiod" Value='[[$cpi.ShortTermPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="cpitransitionperiod">CPI Transition Period</label>
            <input type="text" class="form-control" id="cpitransitionperiod" name="cpitransitionperiod" Value='[[$cpi.TransitionPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="cpilongtermrate">CPI Long Rate</label>
            <input type="text" class="form-control" id="cpilongtermrate" name="cpilongtermrate" Value='[[$cpi.LongTermRate]]'>
        </div>
    </div>
    <div class="form-row">
        <div class="form-group col-sm-1">
            <label for="ervshorttermratesigma"></label>
            <input type="text" class="form-control" id="ervshorttermratesigma" name="ervshorttermratesigma" Value='[[.entity.MCSetup.ERV.ShortTermRate]]'>
            </div>
        <div class="form-group col-sm-1">
            <label for="ervshorttermperiodsigma"></label>
            <input type="text" class="form-control" id="ervshorttermperiodsigma" name="ervshorttermperiodsigma" Value='[[.entity.MCSetup.ERV.ShortTermPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="ervtransitionperiodsigma"></label>
            <input type="text" class="form-control" id="ervtransitionperiodsigma" name="ervtransitionperiodsigma" Value='[[.entity.MCSetup.ERV.TransitionPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="ervlongtermratesigma"></label>
            <input type="text" class="form-control" id="ervlongtermratesigma" name="ervlongtermratesigma" Value='[[.entity.MCSetup.ERV.LongTermRate]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="cpishorttermratesigma"></label>
            <input type="text" class="form-control" id="cpishorttermratesigma" name="cpishorttermratesigma" Value='[[.entity.MCSetup.CPI.ShortTermRate]]'>
            </div>
        <div class="form-group col-sm-1">
            <label for="cpishorttermperiodsigma"></label>
            <input type="text" class="form-control" id="cpishorttermperiodsigma" name="cpishorttermperiodsigma" Value='[[.entity.MCSetup.CPI.ShortTermPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="cpitransitionperiodsigma"></label>
            <input type="text" class="form-control" id="cpitransitionperiodsigma" name="cpitransitionperiodsigma" Value='[[.entity.MCSetup.CPI.TransitionPeriod]]'>
        </div>
        <div class="form-group col-sm-1">
            <label for="cpilongtermratesigma"></label>
            <input type="text" class="form-control" id="cpilongtermratesigma" name="cpilongtermratesigma" Value='[[.entity.MCSetup.CPI.LongTermRate]]'>
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
                <input type="text" class="form-control" id="rett" name="rett" Value='[[.entity.Tax.RETT]]'>
                </div>
            <div class="form-group col-sm-1">
                <label for="landvalue">Land Value</label>
                <input type="text" class="form-control" id="landvalue" name="landvalue" Value='[[.entity.Tax.LandValue]]'>
            </div>
            <div class="form-group col-sm-1">
                <label for="minvalue">Min Depreciable Value</label>
                <input type="text" class="form-control" id="minvalue" name="minvalue" Value='[[.entity.Tax.MinValue]]'>
            </div>
            <div class="form-group col-sm-1">
                <label for="usableperiod">Depreciation Period</label>
                <input type="text" class="form-control" id="usableperiod" name="usableperiod" Value='[[.entity.Tax.UsablePeriod]]'>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-sm-1">
                <label for="vat">VAT</label>
                <input type="text" class="form-control" id="vat" name="vat" Value='[[.entity.Tax.VAT]]'>
                </div>
            <div class="form-group col-sm-1">
                <label for="carrybackyrs">Carry Back Years</label>
                <input type="text" class="form-control" id="carrybackyrs" name="carrybackyrs" Value='[[.entity.Tax.CarryBackYrs]]'>
            </div>
            <div class="form-group col-sm-1">
                <label for="carryforwardyrs">Carry Forward Years</label>
                <input type="text" class="form-control" id="carryforwardyrs" name="carryforwardyrs" Value='[[.entity.Tax.CarryForwardYrs]]'>
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
                <input type="text" class="form-control" id="rentrevision" name="rentrevision" Value="[[.entity.GLA.RentRevisionERV]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="probability">Probability</label>
                <input type="text" class="form-control" id="probability" name="probability" Value="[[.entity.GLA.Probability]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="incentivemonths">Incentive Months</label>
                <input type="text" class="form-control" id="incentivemonths" name="incentivemonths" Value="[[.entity.GLA.RentIncentives.Duration]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="incentivepercent">Incentive Percent</label>
                <input type="text" class="form-control" id="incentivepercent" name="incentivepercent" Value="[[.entity.GLA.RentIncentives.PercentOfContractRent]]">
            </div>            
            <div class="form-group col-sm-1">
                <label for="hazard">Hazard Rate</label>
                <input type="text" class="form-control" id="hazard" name="hazard" Value="[[.entity.GLA.Default.Hazard]]">
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
                <input type="text" class="form-control" id="probabilitysigma" name="probabilitysigma" Value="[[.entity.MCSetup.Probability]]">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="" name="" Value="">
            </div>
            <div class="form-group col-sm-1">
                <input type="text" class="form-control" id="hazardsigma" name="hazardsigma" Value="[[.entity.MCSetup.Hazard]]">
            </div>
        </div>
    </div>
[[end]]

[[define "EntitySettingsValuation"]]
    <div class="form-column">
        <div class="form-row">
            <div class="form-group col-sm-1">
                <label for="entryyield">Entry Yield</label>
                <input type="text" class="form-control" id="entryyield" name="entryyield" value="[[.entity.Valuation.EntryYield]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="exityield">Exit Yield</label>
                <input type="text" class="form-control" id="exityield" name="exityield" value="[[.entity.Valuation.ExitYield]]">
            </div>
            <div class="form-group col-sm-1">
                <label for="yieldshift">Yield Shift (bps)</label>
                <input type="text" class="form-control" id="yieldshift" name="yieldshift" value="[[.entity.Valuation.YieldShift]]">
            </div>
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
        <form class="form-horizontal form-well" role="form" method="post" action="/Fund">
            <div class="tabs-wrapper">
                <ul class="nav nav-tabs tabPinned" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link waves-light active" data-toggle="tab" href="#fundsettings" role="tab">Settings</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link waves-light" data-toggle="tab" href="#fundleasing" role="tab">Leasing</a>
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
                        <button type="submit" class="btn secondary-bg">Submit</button>
                    </div>
                </div>
                <div class="form-group col-sm-1">
                    <label for="fees">Simulations</label>
                    <input type="text" class="form-control" id="sims" name="sims" Value="[[.entity.MCSetup.Sims]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="irr" name="irr" Value="IRR: [[.entity.Metrics.IRR.NetLeveredAfterTax]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="em" name="em" Value="EM: [[.entity.Metrics.EM.NetLeveredAfterTax]]">
                </div>
                [[if eq .entity.Strategy "Balloon"]]
                <div class="form-group col-sm-1 offset-sm-5">
                    <input readonly type="text" class="form-control" id="ytm" name="ytm" Value="YTM: [[.entity.Metrics.BondHolder.YTM]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="dur" name="dur" Value="DUR: [[.entity.Metrics.BondHolder.Duration]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="ytmdur" name="ytmdur" Value="YTM/DUR: ">
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
        <div class="form-group col-sm-1">
            <label for="portfolio">Asset</label>
            <input type="text" class="form-control" id="portfolio" name="portfolio" value="[[.entity.Name]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="holdperiod">Hold Period</label>
            <input type="text" class="form-control" id="fundholdperiod" name="holdperiod" value="[[.entity.HoldPeriod]]">
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