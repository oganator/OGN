[[define "Capex"]]
[[range .CostInput]]
    <div class="form-column">
        <div class="d-flex flex-row" name="costInputRow" id="costInput_[[.MasterID]]">
            <div class="form-group " hidden>
                <label>MasterID</label>
                <input type="text" class="form-control" name="costInputMasterID" id="costInput_[[.MasterID]]_masterID" Value="[[.MasterID]]">
            </div>
            <div class="form-group " hidden>
                <label>Type</label>
                <input type="text" class="form-control" name="costInputType" id="costInput_[[.MasterID]]_type" Value="[[.Type]]">
            </div>
            <div class="form-group ">
                <label>Name</label>
                <input type="text" class="form-control" name="costInputName" id="costInput_[[.MasterID]]_name" Value="[[.Name]]">
            </div>
            <div class="form-group ">
                <label>Amount</label>
                <input type="text" class="form-control" name="costInputAmount" id="costInput_[[.MasterID]]_amount" Value="[[.Amount]]">
            </div>
            <div class="form-group ">
                <label>COA Item Basis</label>
                <select type="text" class="form-control" name="costInputCOAItemBasis" id="costInput_[[.MasterID]]_coaItemBasis" Value="[[.COAItemBasis]]">
                    <option hidden>[[.COAItemBasis]]</option>
                    [[template "coaOptions"]]
                </select>
            </div>
            <div class="form-group ">
                <label>COA Item Target</label>
                <select type="text" class="form-control" name="costInputCOAItemTarget" id="costInput_[[.MasterID]]_coaItemTarget" Value="[[.COAItemTarget]]">
                    <option hidden>[[.COAItemTarget]]</option>
                    [[template "coaOptions"]]
                </select>
            </div>
            <div class="form-group ">
                <label>Growth Item</label>
                <select type="text" class="form-control" name="costInputGrowthItem" id="costInput_[[.MasterID]]_growthItem" Value="[[.GrowthItem]]">
                    <option hidden>[[.GrowthItem]]</option>
                    <option>None</option>
                    [[range $k,$v := .GrowthItemOptions]]
                    <option>[[$k]]</option>
                    [[end]]
                </select>
            </div>
            <div class="form-group " [[if eq .Type "Capex"]] hidden [[end]]>
                <label>Duration</label>
                <input type="text" class="form-control" name="costInputDuration" id="costInput_[[.MasterID]]_duration" Value="[[.Duration]]">
            </div>
            <div class="form-group " [[ if eq .Type "Rent Incentives"]] hidden [[end]] [[ if eq .Type "Fit Out Costs"]] hidden [[end]]>
                <label class="d-flex justify-content-center">Start Date</label>
                <div class="d-flex flex-row">
                    <select type="text" class="form-control" name="costInputStartMonth" id="costInput_[[.MasterID]]_startMonth" Value="[[.Start.MonthName]]">
                        <option hidden>[[.Start.MonthName]]</option>
                        [[template "MonthOptions"]]
                    </select>
                    <input type="text" class="form-control" name="costInputStartYear" id="costInput_[[.MasterID]]_startYear" Value="[[.Start.Year]]"></input>
                </div>
            </div>
            <div class="form-group " [[ if eq .Type "Rent Incentives"]] hidden [[end]] [[ if eq .Type "Fit Out Costs"]] hidden [[end]]>
                <label class="d-flex justify-content-center">End Date</label>
                <div class="d-flex flex-row">
                    <select type="text" class="form-control" name="costInputEndMonth" id="costInput_[[.MasterID]]_endMonth" Value="[[.End.MonthName]]">
                        <option hidden>[[.End.MonthName]]</option>
                        [[template "MonthOptions"]]
                    </select>
                    <input type="text" class="form-control" name="costInputEndYear" id="costInput_[[.MasterID]]_endYear" Value="[[.End.Year]]"></input>
                </div>
            </div>
        </div>
    </div>
    [[end]]
    <br><br>
    <div class="d-flex justify-content-center">
        <button  class="btn secondary-bg" ng-click="addCost()">Add</button>
    </div>
[[end]]

[[template "Capex" .]]