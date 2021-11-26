[[define "FactorAnalysis"]]
<div>
	<table class="table table-hover tableFixHead rounded">
		<thead>
			<tr>
				<th scope="col" data-sortable="true">Correlation</th>
				<th scope="col" data-sortable="true">Void</th>
				<th scope="col" data-sortable="true">Probability</th>
				<th scope="col" data-sortable="true">NumberOfDefaults</th>
				<th scope="col" data-sortable="true">OpEx</th>
				<th scope="col" data-sortable="true">CPI</th>
				<th scope="col" data-sortable="true">ERV</th>
				<th scope="col" data-sortable="true">Hazard</th>
				<th scope="col" data-sortable="true">YieldShift</th>
			</tr>
		</thead>
		<tbody>
		[[range $index, $value := .data.FactorAnalysis]]
			<tr>
                <td>[[$value.Metric]]</td>
				<td>{{[[$value.Void.Corr]] | number:2 }}</td>
				<td>{{[[$value.Probability.Corr]] | number:2 }}</td>
				<td>{{[[$value.NumberOfDefaults.Corr]] | number:2 }}</td>
				<td>{{[[$value.OpEx.Corr]] | number:2 }}</td>
				<td>{{[[$value.CPI.Corr]] | number:2 }}</td>
				<td>{{[[$value.ERV.Corr]] | number:2 }}</td>
				<td>{{[[$value.Hazard.Corr]] | number:2 }}</td>
				<td>{{[[$value.YieldShift.Corr]] | number:2 }}</td>
			</tr>
		[[end]]
	</table>
	<br>
	<table class="table table-hover tableFixHead rounded">
		<thead>
			<tr>
				<th scope="col" data-sortable="true">R Squared</th>
				<th scope="col" data-sortable="true">Void</th>
				<th scope="col" data-sortable="true">Probability</th>
				<th scope="col" data-sortable="true">NumberOfDefaults</th>
				<th scope="col" data-sortable="true">OpEx</th>
				<th scope="col" data-sortable="true">CPI</th>
				<th scope="col" data-sortable="true">ERV</th>
				<th scope="col" data-sortable="true">Hazard</th>
				<th scope="col" data-sortable="true">YieldShift</th>
			</tr>
		</thead>
		<tbody>
		[[range $index, $value := .data.FactorAnalysis]]
			<tr>
                <td>[[$value.Metric]]</td>
				<td>{{[[$value.Void.Rsq]] | number:2 }}</td>
				<td>{{[[$value.Probability.Rsq]] | number:2 }}</td>
				<td>{{[[$value.NumberOfDefaults.Rsq]] | number:2 }}</td>
				<td>{{[[$value.OpEx.Rsq]] | number:2 }}</td>
				<td>{{[[$value.CPI.Rsq]] | number:2 }}</td>
				<td>{{[[$value.ERV.Rsq]] | number:2 }}</td>
				<td>{{[[$value.Hazard.Rsq]] | number:2 }}</td>
				<td>{{[[$value.YieldShift.Rsq]] | number:2 }}</td>
			</tr>
		[[end]]
	</table>	
<!--[[template "MCResultsTable" .]]-->
</div>
[[end]]

[[define "MCResultsTable"]]
	<table class="table table-hover tableFixHead rounded">
		<thead>
			<tr>
				<th scope="col" data-sortable="true">Index</th>
				<th scope="col" data-sortable="true">IRR</th>
				<th scope="col" data-sortable="true">EM</th>
				<th scope="col" data-sortable="true">Void</th>
				<th scope="col" data-sortable="true">Probability</th>
				<th scope="col" data-sortable="true">NumberOfDefaults</th>
				<th scope="col" data-sortable="true">OpEx</th>
				<th scope="col" data-sortable="true">CPI</th>
				<th scope="col" data-sortable="true">ERV</th>
				<th scope="col" data-sortable="true">Hazard</th>
				<th scope="col" data-sortable="true">YieldShift</th>
				<th scope="col" data-sortable="true">End Cash</th>				
				<th scope="col" data-sortable="true">End NCF</th>
				<th scope="col" data-sortable="true">End MV</th>				
				
			</tr>
		</thead>
		<tbody>
		[[$IRR := .data.MCResultSlice.IRR]]
		[[$EM := .data.MCResultSlice.EM]]
		[[$Void := .data.MCResultSlice.Void]]
		[[$Prob := .data.MCResultSlice.Probability]]
		[[$Defaults := .data.MCResultSlice.NumberOfDefaults]]
		[[$OpEx := .data.MCResultSlice.OpEx]]
		[[$CPI := .data.MCResultSlice.CPI]]
		[[$ERV := .data.MCResultSlice.ERV]]
		[[$Hazard := .data.MCResultSlice.Hazard]]
		[[$YieldShift := .data.MCResultSlice.YieldShift]]
		[[$EndNCF := .data.MCResultSlice.EndNCF]]
		[[$EndMV := .data.MCResultSlice.EndMarketValue]]
		[[range $index, $value := .data.MCResultSlice.EndCash]]
			<tr>
                <td>[[$index]]</td>
				<td>{{[[index $IRR $index]] | number:4}}</td>
				<td>{{[[index $EM $index]] | number:4}}</td>				
				<td>{{[[index $Void $index]] | number:0}}</td>
				<td>{{[[index $Prob $index]] | number:4}}</td>
				<td>{{[[index $Defaults $index]] | number:0}}</td>
				<td>{{[[index $OpEx $index]] | number:4}}</td>
				<td>{{[[index $CPI $index]] | number:4}}</td>
				<td>{{[[index $ERV $index]] | number:4}}</td>
				<td>{{[[index $Hazard $index]] | number:4}}</td>
				<td>{{[[index $YieldShift $index]] | number:4}}</td>
				<td>{{[[$value]] | number:0 }}</td>
				<td>{{[[index $EndNCF $index]] | number:0}}</td>
				<td>{{[[index $EndMV $index]] | number:0}}</td>
			</tr>
		[[end]]
		</tbody>
	</table>
[[end]]

[[template "FactorAnalysis" .]] 