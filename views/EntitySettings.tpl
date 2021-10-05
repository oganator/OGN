[[template "header"]]

[[define "EntitySettings"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <form class="form-horizontal form-well" role="form" method="post" action="/">
            <div class="tabs-wrapper">
                <ul class="nav nav-tabs tabPinned" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link waves-light active" data-toggle="tab" href="#settings" role="tab">Settings</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link waves-light" data-toggle="tab" href="#ervcpi" role="tab">ERV/CPI Growth</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link waves-light" data-toggle="tab" href="#val" role="tab">Valuation</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link waves-light" data-toggle="tab" href="#tax" role="tab">Tax</a>
                    </li>
                </ul>
            </div>
            <div class="tab-content">
                <div class="tab-pane fade in show active" id="settings" role="tabpanel">
                    [[template "EntitySettingsDefault" .]]
                </div>
                <div class="tab-pane fade" id="ervcpi" role="tabpanel">
                    [[template "EntitySettingsERVCPI" .]]
                </div>
                <div class="tab-pane fade" id="val" role="tabpanel">
                </div>
                <div class="tab-pane fade" id="tax" role="tabpanel">
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
                <div class="form-group col-sm-1 offset-sm-6" ng-show="strategy != 'Standard' ">
                    <input readonly type="text" class="form-control" id="ytm" name="ytm" Value="YTM: [[.entity.Metrics.BondHolder.YTM]]">
                </div>
                <div class="form-group col-sm-1" ng-show="strategy != 'Standard' ">
                    <input readonly type="text" class="form-control" id="dur" name="dur" Value="DUR: [[.entity.Metrics.BondHolder.Duration]]">
                </div>
                <div class="form-group col-sm-1" ng-show="strategy != 'Standard' ">
                    <input readonly type="text" class="form-control" id="ytmdur" name="ytmdur" Value="YTM/DUR: ">
                </div>
            </div>
        </form>
    </div>
    <br>
    <br>
    <script>
        window.onload = ytmDur()
        function convertPercent(){
            document.getElementById('entryyield').value = document.getElementById('entryyield').value*100;
            document.getElementById('exityield').value = document.getElementById('exityield').value*100;
            document.getElementById('discount').value = document.getElementById('discount').value*100;
            document.getElementById('soldrent').value = document.getElementById('soldrent').value*100;
            document.getElementById('erv').value = document.getElementById('erv').value*100;
            document.getElementById('cpi').value = document.getElementById('cpi').value*100;
            document.getElementById('opex').value = document.getElementById('opex').value*100;
            document.getElementById('balpercent').value = document.getElementById('balpercent').value*100;
            ytmDur();
        }
        function ytmDur(){
            temp = Math.round([[.entity.Metrics.BondHolder.YTM]] / [[.entity.Metrics.BondHolder.Duration]]*1000)/1000
            document.getElementById('ytmdur').value = document.getElementById('ytmdur').value.concat(temp);
        }
    </script>
[[end]]

[[define "EntitySettingsDefault"]]                  
    <div class="form-row">
        <input hidden id="name" name="name" ng-model="entity" value="[[.entity.Name]]">
        <div class="form-group col-sm-1">
            <label for="holdperiod">Hold Period</label>
            <input type="text" class="form-control" id="holdperiod" name="holdperiod" value="[[.entity.HoldPeriod]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="entryyield">Entry Yield</label>
            <input type="text" class="form-control" id="entryyield" name="entryyield" value="[[.entity.Valuation.EntryYield]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="exityield">Exit Yield</label>
            <input type="text" class="form-control" id="exityield" name="exityield" value="[[.entity.Valuation.ExitYield]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="fees">Fees (bps)</label>
            <input type="text" class="form-control" id="fees" name="fees" Value="[[.entity.Fees.PercentOfGAV]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="rate">Strategy</label>
            <select type="text" class="form-control" id="strategy" name="strategy" placeholder="[[.entity.Strategy]]" ng-model="strategy">
                <option hidden>[[.entity.Strategy]]</option>
                <option>Standard</option>
                <option>Balloon</option>
                <option>Pure Discount</option>
            </select>
        </div>
        <div class="form-group col-sm-1" ng-show="strategy == 'Standard' ">
            <label for="ltv">LTV</label>
            <input type="text" class="form-control" id="ltv" name="ltv" Value="[[.entity.DebtInput.LTV]]">
        </div>
        <div class="form-group col-sm-1" ng-show="strategy == 'Standard' ">
            <label for="rate">Loan Rate</label>
            <input type="text" class="form-control" id="rate" name="rate" Value="[[.entity.DebtInput.InterestRate]]">
        </div>
        <div class="form-group col-sm-1" ng-show="strategy != 'Standard' ">
            <label for="discount">Discount Rate</label>
            <input type="text" class="form-control" id="discount" name="discount" Value="[[.entity.GLA.DiscountRate]]">
        </div>
        <div class="form-group col-sm-1" ng-show="strategy != 'Standard' ">
            <label for="rate">Rent to Sell %</label>
            <input type="text" class="form-control" id="soldrent" name="soldrent" Value="[[.entity.GLA.PercentSoldRent]]">
        </div>
        <div class="form-group col-sm-1" ng-show="strategy != 'Standard' ">
            <label for="balpercent">Balloon Percent</label>
            <input type="text" class="form-control" id="balpercent" name="balpercent" Value="[[.entity.BalloonPercent]]" >
        </div>
    </div>
    <div class="form-row">
        
        <div class="form-group col-sm-1">
            <label for="yieldshift">Yield Shift (bps)</label>
            <input type="text" class="form-control" id="yieldshift" name="yieldshift" value="[[.entity.Valuation.YieldShift]]">
        </div>
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
            <label for="opex">Operating Expenses</label>
            <input type="text" class="form-control" id="opex" name="opex" Value="[[.entity.OpEx.PercentOfTRI]]">
        </div>
        <div class="form-group col-sm-1">
            <label for="hazard">Hazard Rate</label>
            <input type="text" class="form-control" id="hazard" name="hazard" Value="[[.entity.GLA.Default.Hazard]]">
        </div>
    </div>
    <div class="form-row">
        <div class="form-group col-sm-1">
            <input type="text" class="form-control" id="yieldshiftsigma" name="yieldshiftsigma" value="[[.entity.MCSetup.YieldShift]]">
        </div>
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
            <input type="text" class="form-control" id="opexsigma" name="opexsigma" Value="[[.entity.MCSetup.OpEx]]">
        </div>
        <div class="form-group col-sm-1">
            <input type="text" class="form-control" id="hazardsigma" name="hazardsigma" Value="[[.entity.MCSetup.Hazard]]">
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