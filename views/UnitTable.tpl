[[define "UnitTable"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col"></th>
                    <th scope="col">Status</th>
                    <th scope="col">% Sold Rent</th>
                    <th scope="col">Lease End</th>
                    <th scope="col">ERV Area</th>
                    <th scope="col">ERV Amount</th>
                    <th scope="col">Passing Rent</th>
                    <th scope="col">Probability</th>
                    <th scope="col">Rent Schedule</th>
                </tr>
            </thead>
            <tbody>
                [[$name := .Name]]
                [[range .ChildUnits]]
                <tr>
                    <td style="border-color:006A4D;">
                        <div class="text-center">
                            <button href="#unitcf" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#unitcf" ng-click="getUnitCF('[[.MasterID]]')">View Cash Flow</button>
                        </div>
                    </td>
                    <td style="border-color:006A4D;">[[.UnitStatus]]</td>
                    <td style="border-color:006A4D;">[[.PercentSoldRent]]</td>
                    <td style="border-color:006A4D;">[[.LeaseExpiryDate.Year]]/[[.LeaseExpiryDate.Month]]</td>
                    <td style="border-color:006A4D;">{{[[printf "%.0f" .ERVArea]] | number:0}}</td>
                    <td style="border-color:006A4D;">{{[[printf "%.0f" .ERVAmount]] | number:0}}</td>
                    <td style="border-color:006A4D;">{{[[printf "%.0f" .PassingRent]] | number:0}}</td>
                    <td style="border-color:006A4D;">{{[[.Probability]] *100 | number:2}}%</td>
                    <td style="border-color:006A4D;">
                        <div class="text-center">
                            <button href="#rentschedule" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#rentschedule" ng-click="getRentSchedule('[[.MasterID]]')">View</button>
                        </div>
                    </td>
                </tr>
                [[end]]
            </tbody>
        </table>
    </div>
    <div class="modal" id="rentschedule" tabindex="-1" role="dialog" aria-labelledby="rentschedule" aria-hidden="true">
		<div class="modal-content modal-content-centered container-fluid" style="width: 80%;">
			<div bind-html-compile = rentschedule></div>
		</div>
	</div>
    <div class="modal" id="unitcf" tabindex="-1" role="dialog" aria-labelledby="unitcf" aria-hidden="true">
		<div class="modal-content modal-content-centered container-fluid" style="width: 80%;">
			<div bind-html-compile = unitcf></div>
		</div>
	</div>
[[end]]

