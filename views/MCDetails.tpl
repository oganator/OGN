[[define "MCDetails"]]

<div>
	<table class="table table-hover tableFixHead rounded">
		<thead>
			<tr>
				<th scope="col" data-sortable="true">View CF</th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("irr")'>IRR</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("em")'>Equity Multiple</button></th>
				[[if (eq .Strategy "Balloon")]]
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("ytm")'>YTM</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("duration")'>Duration</button></th>
				[[end]]
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("void")'>Void</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("prob")'>Probability</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("hazard")'>Number of Defaults</button></th>
				
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("cpi")'>CPI Ending Index</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("erv")'>ERV Ending Index</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("hazard")'>Hazard Rate</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("yieldshift")'>Yield Shift</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("endcash")'>Ending Cash</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("endncf")'>Ending NCF</button></th>
				<th scope="col" data-sortable="true"><button type="button" class="btn" ng-click='sortMCDetails("endmarketvalue")'>Ending Market Value</button></th>
			</tr>
		</thead>
		<tbody>
			[[range $index, $value := .MCSlice]]
			<tr>
				<td><button type="button" class="btn" href="#mccf" data-toggle="modal" data-target="#mccf" ng-click='viewCFIndex("[[$index]]")'>View CF</button></td>
				<td>[[printf "%.2f" $value.Metrics.IRR.NetLeveredAfterTax]]</td>
				<td>[[printf "%.2f" $value.Metrics.EM.NetLeveredAfterTax]]</td>
				[[if (eq .Strategy "Balloon")]]
				<td>[[printf "%.2f" $value.Metrics.BondHolder.YTM]]</td>
				<td>[[printf "%.2f" $value.Metrics.BondHolder.Duration]]</td>
				[[end]]
				<td>[[$value.GLA.Void]]</td>
				<td>{{[[$value.GLA.Probability]] *100 | number:2}}</td>
				<td>{{[[$value.GLA.Default.NumberOfDefaults]] | number:0}}</td>
				
				[[range $value.GrowthInput]]
				<td>{{[[.EndingIndex]] | number:4}}</td>
				[[end]]
				<td>{{[[$value.GLA.Default.Hazard]] | number:4}}</td>
				<td>[[printf "%.2f" $value.Valuation.YieldShift]]</td>
				<td>{{[[$value.MCResults.EndCash.Mean]] | number:0}}</td>
				<td>{{[[$value.MCResults.EndNCF.Mean]] | number:0}}</td>
				<td>{{[[$value.MCResults.EndMarketValue.Mean]] | number:0}}</td>
			</tr>
			[[end]]
		</tbody>
	</table>
	<div class="d-flex flex-row-reverse">
		<button type="button" class="btn col-sm-1 align-items-end" ng-click="nextMCDetails('+')" ng-hide="hideNext()">Next Page</button>
		<button type="button" class="btn col-sm-1 align-items-end" ng-click="nextMCDetails('-')" ng-hide="hidePrevious()">Previous Page</button>
		page {{mcdetailspage}} of {{mcdetailspagestotal}} &nbsp 
	</div>
	<br>
	<div class="modal" id="mccf" tabindex="-1" role="dialog" aria-labelledby="mccf" aria-hidden="true">
		<div class="modal-content container-fluid" style="width: 95%; overflow-x:scroll;">
			<div bind-html-compile = data></div>
			<!--<div bind-html-compile = unittable></div>-->
		</div>
	</div>
</div>
[[end]]
[[template "MCDetails" .]]
