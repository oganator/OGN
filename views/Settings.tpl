[[template "header"]]

[[define "EntitySettings"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <br>
        <form class="form-horizontal form-well" style="color: #006A4D;" role="form" method="post" action="/">
                <div class="tab-pane fade in show active" id="lease" role="tabpanel">
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
                            <label for="ltv">LTV</label>
                            <input type="text" class="form-control" id="ltv" name="ltv" Value="[[.entity.DebtInput.LTV]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="rate">Loan Rate</label>
                            <input type="text" class="form-control" id="rate" name="rate" Value="[[.entity.DebtInput.InterestRate]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="discount">Discount Rate</label>
                            <input type="text" class="form-control" id="discount" name="discount" Value="[[.entity.GLA.DiscountRate]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="rate">Rent to Sell %</label>
                            <input type="text" class="form-control" id="soldrent" name="soldrent" Value="[[.entity.GLA.PercentSoldRent]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="rate">Strategy</label>
                            <select type="text" class="form-control" id="strategy" name="strategy" placeholder="[[.entity.Strategy]]">
                                <option>[[.entity.Strategy]]</option>
                                <option>Pure Discount</option>
                                <option>Amortized Coupon</option>
                                <option>Balloon</option>
                            </select>
                        </div>
                        [[if eq .entity.Strategy "Balloon"]]
                        <div class="form-group col-sm-1">
                            <label for="balpercent">Balloon Percent</label>
                            <input type="text" class="form-control" id="balpercent" name="balpercent" Value="[[.entity.BalloonPercent]]">
                        </div>
                        [[end]]
                        <div class="form-group col-sm-1">
                            <label for="fees">Fees (bps)</label>
                            <input type="text" class="form-control" id="fees" name="fees" Value="[[.entity.Fees.PercentOfGAV]]">
                        </div>
                    </div>
                    <div class="form-row">
                        <div class="form-group col-sm-1">
                            <label for="erv">ERV Growth</label>
                            <input type="text" class="form-control" id="erv" name="erv" Value="[[index .entity.GrowthInput "ERV"]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="erv">CPI Growth</label>
                            <input type="text" class="form-control" id="cpi" name="cpi" Value="[[index .entity.GrowthInput "CPI"]]">
                        </div>
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
                            <input type="text" class="form-control" id="ervsigma" name="ervsigma" Value="[[.entity.MCSetup.ERV]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <input type="text" class="form-control" id="cpisigma" name="cpisigma" Value="[[.entity.MCSetup.CPI]]">
                        </div>
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
                </div>
            <br>
            <div class="form-row"> 
                <div>
                    <div class="form-group col-md-4 d-flex align-items-end">
                        <button type="submit" class="btn" style="background-color: #006A4D; color:white">Submit</button>
                    </div>
                </div>
                <div class="form-group col-sm-1">
                    <label for="fees">Simulations</label>
                    <input type="text" class="form-control" id="sims" name="sims" Value="[[.entity.MCSetup.Sims]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="irr" name="irr" Value="IRR: [[.entity.Metrics.IRR.NetLeveredAfterTax]]">
                </div>
                <div class="form-group col-sm-1 offset-sm-6">
                    <input readonly type="text" class="form-control" id="ytm" name="ytm" Value="YTM: [[.entity.Metrics.BondHolder.YTM]]">
                </div>
                <div class="form-group col-sm-1">
                    <input readonly type="text" class="form-control" id="dur" name="dur" Value="DUR: [[.entity.Metrics.BondHolder.Duration]]">
                </div>
                <div class="form-group col-sm-1">
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

[[define "UnitSettings"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <br>
        <form class="form-horizontal form-well" style="color: #006A4D;" role="form" method="post" action="/UpdateUnit">
            <div class="tabs-wrapper">
                <ul class="nav nav-tabs tabPinned" role="tablist">
                    <li class="nav-item">
                        <a class="nav-link waves-light active" data-toggle="tab" href="#lease" role="tab">Lease Assumptions</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link waves-light" data-toggle="tab" href="#erv" role="tab">Rent & Indexation</a>
                    </li>
                </ul>
            </div>
            <div class="tab-content">
                <div class="tab-pane fade in show active" id="lease" role="tabpanel">
                    <div class="form-row">
                        <div class="form-group col-sm-1">
                            <label for="parent" class="bmd-label-floating">Parent</label>
                            <input type="text" readonly class="form-control" id="parent" name="parent" value=[[.Parent.Name]]>
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="name" class="bmd-label-floating">Name</label>
                            <input type="text" readonly class="form-control" id="name" name="name" value=[[.Name]]>
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="leasestartmonth" class="bmd-label-floating">Lease Start Month</label>
                            <input type="text" class="form-control" id="leasestartmonth" name="leasestartmonth" value="[[.LeaseStartDate.Month]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="leasestartyear">Lease Start Year</label>
                            <input type="text" class="form-control" id="leasestartyear" name="leasestartyear" value="[[.LeaseStartDate.Year]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="leaseexpirymonth">Lease Expiry Month</label>
                            <input type="text" class="form-control" id="leaseexpirymonth" name="leaseexpirymonth" value="[[.LeaseExpiryDate.Month]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="leaseexpiryyear">Lease Expiry Year</label>
                            <input type="text" class="form-control" id="leaseexpiryyear" name="leaseexpiryyear" value="[[.LeaseExpiryDate.Year]]">
                        </div>
                        <div class="form-group col-sm-1 d-flex align-items-end">
                            <div class="btn-group btn-group-toggle" id="frequency" data-toggle="buttons">
                                <label class="btn btn-secondary" id="monthlylabel">
                                    <input type="radio" name="frequency" id="monthlyinput" value="Monthly">Monthly 
                                </label>
                                <label class="btn btn-secondary" id="quarterlylabel">
                                    <input type="radio" name="frequency" id="quarterlyinput" value="Quarterly">Quarterly
                                </label>
                            </div>
                        </div>
                    </div>
                    <div class="form-row">
                        <div class="form-group col-sm-1">
                            <label for="void">Void</label>
                            <input type="text" class="form-control" id="void" name="void"  Value="[[.Model.Void.Value]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="duration">Extension Duration</label>
                            <input type="text" class="form-control" id="duration" name="duration"  Value="[[.Model.EXTDuration.Value]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="rentrevision">Rent Revision ERV</label>
                            <input type="text" class="form-control" id="rentrevision" name="rentrevision"  Value="[[.Model.RentRevisionERV.Value]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="probability">Probability</label>
                            <input type="text" class="form-control" id="probability" name="probability"  Value="[[.Model.Probability.Value]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="incentivesm">Rent Incentives (m)</label>
                            <input type="text" class="form-control" id="incentivesm" name="incentivesm"  Value="[[.Model.RentIncentives.RenewalMonths.Value]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="incentivesp">Rent Incentives (%)</label>
                            <input type="text" class="form-control" id="incentivesp" name="incentivesp"  Value="[[.Model.RentIncentives.RenewalPercent.Value]]">
                        </div>
                        <div class="form-group col-sm-1">
                            <label for="fitoutcosts">Fit Out Costs per sqft</label>
                            <input type="text" class="form-control" id="fitoutcosts" name="fitoutcosts"  Value="[[.Model.FitOutCosts.AmountPerTotalArea]]">
                        </div>
                    </div>
                </div>
                <div class="tab-pane fade" id="erv" role="tabpanel">
                    <div class="form-row">
                        &nbsp&nbsp&nbsp
                        <div class="col-sm-1">
                            <div class="form-group">
                                <label for="rent">Passing Rent</label>
                                <input type="text" class="form-control" id="rent" name="rent" value="[[printf "%.0f" .PassingRent]]">
                            </div>
                            <div class="form-group">
                                <label for="indexfreq">Indexation Frequency (yrs)</label>
                                <input type="text" class="form-control" id="indexfreq" name="indexfreq" value="[[.Model.IndexDetails.Value.Frequency]]">
                            </div>
                            <div class="form-group">
                                <label for="indextype">Indexation Type</label>
                                <select input type="text" class="form-control" id="indextype" name="indextype">
                                    <option selected> [[.Model.IndexDetails.Value.Type]] </option>
                                    <option> ERV</option>
                                    [[$growth := .Parent.Model.GrowthInputData]]
                                    [[range $key, $value := $growth]]
                                        [[if not $value.IsERV]]<option>[[$key]]</option>[[end]]
                                    [[end]]
                                </select>
                            </div>
                        </div> &nbsp&nbsp&nbsp&nbsp
                        [[$growth := .InitialGrowth]]
                        [[range $key, $value := $growth]]
                        <div class="col-sm-1">
                            <div class="form-group">
                                <label for="[[$key]]">Type</label>
                                <input type="text" readonly class="form-control" id="[[$key]]" name="[[$key]]"  Value="[[$key]]">
                            </div>
                            <div class="form-group">
                                <label for="Area[[$key]]">Area</label>
                                <input type="text" class="form-control" id="[[$key]]Area" name="[[$key]]Area"  Value="[[$value.Area]]">
                            </div>
                            <div class="form-group">
                                <label for="Amount[[$key]]">Amount</label>
                                <input type="text" class="form-control" id="[[$key]]Amount" name="[[$key]]Amount"  Value="[[$value.Amount]]">
                            </div>
                        </div>
                        [[end]]
                    </div>
                </div>
            </div>
            <script>
            window.onload = checkFrequency()
                function checkFrequency(){
                    var freq = [[.Model.Settings.Frequency]];
                    var monthlylabel = document.getElementById('monthlylabel');
                    var monthlyinput = document.getElementById('monthlyinput');
                    var quarterlylabel = document.getElementById('quarterlylabel');
                    var quarterlyinput = document.getElementById('quarterlyinput');
                    if (freq == 'Monthly'){
                        monthlylabel.setAttribute('class','btn btn-secondary active focus');
                        monthlyinput.click();
                    }
                    if (freq == 'Quarterly'){
                        quarterlylabel.setAttribute('class','btn btn-secondary active focus');
                        quarterlyinput.click();
                    }
                }
            </script>
            [[template "footer"]]
            <br>
            <div class="form-row"> 
                <div>
                    <div class="form-group col-md-4 d-flex align-items-end">
                        <button type="submit" class="btn" style="background-color: #006A4D; color:white">Submit</button>
                    </div>
                </div>
                <div> IRR: %
                </div>
            </div>
        </form>
    </div>
    <br>
    <br>
[[end]]

