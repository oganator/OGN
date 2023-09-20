[[define "sensitivity"]]
<div class="container-fluid" style="width: 95%;">
    <div class="row">
        <div class="container-fluid" style="width: 50%;">
        <h3>Sensitivity Setup</h3>
        [[$entity := .entity]]
        [[$sens := .entity.SensitivityInput]]
        <table class="table" id="sensitivityInputTable">
            <thead>
                <tr>
                    <th scope="col">Variable</th>
                    <th scope="col">Model Value</th>
                    <th scope="col">Increment</th>
                    <th scope="col">Vertical</th>
                    <th scope="col">Horizontal</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th scope="row">Void</th>
                    <td>[[.entity.GLA.Void]]</td>
                    <td><input name="Void" id="sensRow[[$sens.Add]]" value=[[.entity.SensitivityInput.Void]]></input></td>
                    <td><input name="Void" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical "Void"]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Void" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal "Void"]] checked [[end]] class="form-control"></input></td>
                </tr>
                <tr>
                    <th scope="row">Probability</th>
                    <td>{{ [[.entity.GLA.Probability]] * 100 | number:2 }} %</td>
                    <td><input name="Renewal Probability" id="sensRow[[$sens.Add]]" value="{{ [[.entity.SensitivityInput.Probability]] * 100 | number:2 }}"></input></td>
                    <td><input name="Renewal Probability" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical "Renewal Probability"]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Renewal Probability" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal "Renewal Probability"]] checked [[end]] class="form-control"></input></td>
                </tr>
                [[range $key, $value := .entity.GrowthInput]]
                [[ $growth := index $entity.SensitivityInput.Growth $key]]
                <tr>
                    <th scope="row">[[$key]] Short term rate</th>
                    <td>{{ [[$value.ShortTermRate]] * 100 | number:2 }} %</td>
                    <td><input name="[[$key]] Short term rate" id="sensRow[[$sens.Add]]" value="{{ [[$growth.ShortTermRate]] * 100 | number:2 }}"></input></td>
                    <td><input name="[[$key]] Short term rate" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical (print $key " Short term rate") ]] checked [[end]] class="form-control"></input></td>
                    <td><input name="[[$key]] Short term rate" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal (print $key " Short term rate") ]] checked [[end]] class="form-control"></input></td>
                </tr>
                <tr>
                    <th scope="row">[[$key]] Long term rate</th>
                    <td>{{ [[$value.LongTermRate]] * 100 | number:2 }} %</td>
                    <td><input name="[[$key]] Long term rate" id="sensRow[[$sens.Add]]" value="{{ [[$growth.LongTermRate]] * 100 | number:2 }}"></input></td>
                    <td><input name="[[$key]] Long term rate" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical (print $key " Long term rate") ]] checked [[end]] class="form-control"></input></td>
                    <td><input name="[[$key]] Long term rate" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal (print $key " Long term rate") ]] checked [[end]] class="form-control"></input></td>
                </tr>
                [[end]]
                [[range $key, $value := .entity.CostInput]]
                [[ $cost := index $entity.SensitivityInput.EntityCosts $key ]]
                <tr>
                    <th scope="row">[[$value.Name]]</th>
                    <td>[[$value.Amount]]</td>
                    <td><input name="Entity [[$value.Name]]" id="sensRow[[$sens.Add]]" value=[[$cost.Amount]]></input></td>
                    <td><input name="Entity [[$value.Name]]" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical $value.Name ]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Entity [[$value.Name]]" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal $value.Name ]] checked [[end]] class="form-control"></input></td>
                </tr>
                [[end]]
                [[range $key, $value := .entity.GLA.CostInput]]
                [[ $cost := index $entity.SensitivityInput.UnitCosts $key ]]
                <tr>
                    <th scope="row">[[$value.Name]]</th>
                    <td>[[$value.Amount]]</td>
                    <td><input name="Unit [[$value.Name]]" id="sensRow[[$sens.Add]]" value=[[$cost.Amount]]></input></td>
                    <td><input name="Unit [[$value.Name]]" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical $value.Name ]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Unit [[$value.Name]]" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal $value.Name ]] checked [[end]] class="form-control"></input></td>
                </tr>
                [[end]]
                [[range $key, $value := .entity.DebtInput]]
                [[ $debt := index $entity.SensitivityInput.Debt $key]]
                <tr>
                    <th scope="row">Spread: [[$value.Name]]</th>
                    <td> [[$value.Spread]]</td>
                    <td><input name="[[$value.Name]]" id="sensRow[[$sens.Add]]" value="[[$debt.Spread]]"></input></td>
                    <td><input name="[[$value.Name]]" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical $value.Name ]] checked [[end]] class="form-control"></input></td>
                    <td><input name="[[$value.Name]]" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal $value.Name ]] checked [[end]] class="form-control"></input></td>
                </tr>
                [[end]]
                <tr>
                    <th scope="row">Yield Shift</th>
                    <td>[[.entity.Valuation.YieldShift]]</td>
                    <td><input name="Yield Shift" id="sensRow[[$sens.Add]]" value=[[.entity.SensitivityInput.Valuation.YieldShift]]></input></td>
                    <td><input name="Yield Shift" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical "Yield Shift"]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Yield Shift" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal "Yield Shift"]] checked [[end]] class="form-control"></input></td>
                </tr>
                <tr>
                    <th scope="row">Discount Rate</th>
                    <td>{{ [[.entity.Valuation.DiscountRate]] * 100 | number:2 }}</td>
                    <td><input name="Discount Rate" id="sensRow[[$sens.Add]]" value="{{ [[.entity.SensitivityInput.Valuation.DiscountRate]] * 100 | number:2 }}"></input></td>
                    <td><input name="Discount Rate" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical "Discount Rate"]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Discount Rate" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal "Discount Rate"]] checked [[end]] class="form-control"></input></td>
                </tr>
                <tr>
                    <th scope="row">Hold Period</th>
                    <td> [[.entity.HoldPeriod]] </td>
                    <td><input name="Hold Period" id="sensRow[[$sens.Add]]" value="[[.entity.SensitivityInput.HoldPeriod]]"></input></td>
                    <td><input name="Hold Period" type="checkbox" id="sensRowVertical_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Vertical "Hold Period"]] checked [[end]] class="form-control"></input></td>
                    <td><input name="Hold Period" type="checkbox" id="sensRowHorizontal_[[$sens.Count]]" [[if eq $entity.SensitivityInput.Horizontal "Hold Period"]] checked [[end]] class="form-control"></input></td>
                </tr>
                <div hidden>[[.entity.SensitivityInput.Reset]]</div>
            </tbody>
        </table>
        </div>
        <div class="container-fluid" style="width: 50%;"> 
            <div class="row">
                <h3 class="align-self-center" style="rotate: -90deg; ">[[.entity.SensitivityInput.Vertical]]</h3>
                <div class="col">
                    <h3 class="d-flex justify-content-center">[[.entity.SensitivityInput.Horizontal]]</h3>
                    <table class="table tableFixHead" id="sensitivityGrid">
                        <tbody>
                        [[range .grid]]
                            [[$row := .]]
                            <tr>
                            [[range $row]]
                                <td>{{[[.]] | number: 3}}</td>
                            [[end]]
                            </tr>
                        [[end]]
                        </tbody>
                    </table>
                </div>
            </div>
            <div>
                <input id="iterations" value="[[.entity.SensitivityInput.Iterations]]"></input>
                <button ng-click="postSensitivity()" class="btn secondary-bg">Submit</button>
            </div>
        </div>
    </div>
</div>
[[end]]

[[template "sensitivity" .]]

