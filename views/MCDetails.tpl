[[define "MCDetails"]]
<div>
    <table class="table table-hover tableFixHead rounded" data-toggle="table" data-pagination="true">
        <thead>
            <tr>
                <th scope="col" data-sortable="true">Index</th>
                <th scope="col" data-sortable="true">IRR</th>
                <th scope="col" data-sortable="true">YTM</th>
                <th scope="col" data-sortable="true">Duration</th>
                <th scope="col" data-sortable="true">Void</th>
                <th scope="col" data-sortable="true">Extension Duration</th>
                <th scope="col" data-sortable="true">Hazard Rate</th>
                <th scope="col" data-sortable="true">OpEx(%CR)</th>
                <th scope="col" data-sortable="true">CPI Growth</th>
                <th scope="col" data-sortable="true">ERV Growth</th>
                <th scope="col" data-sortable="true">Yield Shift</th>
                <th scope="col" data-sortable="true">Ending Cash</th>
                <th scope="col" data-sortable="true">Ending NCF</th>
            </tr>
        </thead>
        <tbody>
            [[range $index, $value := .MCSlice]]
            <tr>
                <td><button type="button" class="btn" style="background-color: #006A4D; color:white" href="#mccf" data-toggle="modal" data-target="#mccf" ng-click='viewCFIndex("[[$index]]")'>View CF [[$index]]</button></td>
                <td>[[printf "%.4f" $value.Metrics.IRR.NetLeveredAfterTax]]</td>
                <td>[[printf "%.4f" $value.Metrics.BondHolder.YTM]]</td>
                <td>[[printf "%.4f" $value.Metrics.BondHolder.Duration]]</td>
                <td>[[$value.GLA.Void]]</td>
                <td>[[$value.GLA.EXTDuration]]</td>
                <td>[[printf "%.3f" $value.GLA.Default.Hazard]]</td>
                <td>[[printf "%.2f" $value.OpEx.PercentOfTRI]]</td>
                [[range $value.GrowthInput]]
                <td>[[printf "%.3f" .]]</td>
                [[end]]
                <td>[[printf "%.2f" $value.Valuation.YieldShift]]</td>
                <td>[[printf "%.2f" $value.MCResults.EndCash.Mean]]</td>
                <td>[[printf "%.2f" $value.MCResults.EndNCF.Mean]]</td>
            </tr>
            [[end]]
        </tbody>
    </table>
    <div class="modal" id="mccf" tabindex="-1" role="dialog" aria-labelledby="mccf" aria-hidden="true">
        <div class="modal-content container-fluid shadow-lg rounded" style="width: 95%; overflow-x:scroll;">
            <div ng-bind-html = data></div>
        </div>
    </div>
</div>
[[end]]
[[template "MCDetails" .]]
