[[template "header"]]

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