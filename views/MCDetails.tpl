[[define "MCDetails"]]
<div>
	<table class="table table-hover tableFixHead rounded">
		<thead>
			<tr>
				<th scope="col" data-sortable="true">View CF</th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("irr")'>IRR</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("em")'>Equity Multiple</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("ytm")'>YTM</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("duration")'>Duration</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("void")'>Void</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("extdur")'>EXT Duration</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("hazard")'>Hazard Rate</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("opex")'>OpEx %TRI</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("cpi")'>CPI Growth</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("erv")'>ERV Growth</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("yieldshift")'>Yield Shift</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("endcash")'>Ending Cash</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn border border-white" style="background-color: #006A4D; color:white;" ng-click='sortMCDetails("endncf")'>Ending NCF</button></th>
			</tr>
		</thead>
		<tbody>
			[[range $index, $value := .MCSlice]]
			<tr>
				<td><button type="button" class="btn" style="background-color: #006A4D; color:white" href="#mccf" data-toggle="modal" data-target="#mccf" ng-click='viewCFIndex("[[$index]]")'>View CF</button></td>
				<td>[[printf "%.2f" $value.Metrics.IRR.NetLeveredAfterTax]]</td>
				<td>[[printf "%.2f" $value.Metrics.EM.NetLeveredAfterTax]]</td>
				<td>[[printf "%.2f" $value.Metrics.BondHolder.YTM]]</td>
				<td>[[printf "%.2f" $value.Metrics.BondHolder.Duration]]</td>
				<td>[[$value.GLA.Void]]</td>
				<td>[[$value.GLA.EXTDuration]]</td>
				<td>{{[[$value.GLA.Default.Hazard]] *100 | number:2}}</td>
				<td>{{[[$value.OpEx.PercentOfTRI]] *100 | number:2}}</td>
				[[range $value.GrowthInput]]
				<td>{{[[.]] *100 | number:2}}</td>
				[[end]]
				<td>[[printf "%.2f" $value.Valuation.YieldShift]]</td>
				<td>{{[[$value.MCResults.EndCash.Mean]] | number:0}}</td>
				<td>{{[[$value.MCResults.EndNCF.Mean]] | number:0}}</td>
			</tr>
			[[end]]
		</tbody>
	</table>
	<div class="d-flex flex-row-reverse">
		<button type="button" class="btn col-sm-1 align-items-end" style="background-color: #006A4D; color:white;">Next Page</button>
	</div>
	<br>
	<div class="modal" id="mccf" tabindex="-1" role="dialog" aria-labelledby="mccf" aria-hidden="true">
		<div class="modal-content container-fluid shadow-lg rounded" style="width: 95%; overflow-x:scroll;">
			<div ng-bind-html = data></div>
		</div>
	</div>
</div>
[[end]]
[[template "MCDetails" .]]
