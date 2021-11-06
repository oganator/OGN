[[define "UnitTable"]]
    [[if ne .ParentID 0]]
        <div class="container-fluid shadow-lg rounded" style="width: 95%">
            <table class="table table-hover tableFixHead rounded">
                <thead>
                    <tr>
                        <th scope="col">Unit Name</th>
                        <th scope="col">Tenant</th>
                        <th scope="col">Status</th>
                        <th scope="col">Lease End</th>
                        <th scope="col">ERV Area</th>
                        <th scope="col">ERV Amount</th>
                        <th scope="col">Passing Rent</th>
                        <th scope="col">Probability</th>
                        <th scope="col"></th>
                        <th scope="col"></th>
                    </tr>
                </thead>
                <tbody>
                    [[$name := .Name]]
                    [[range .ChildUnits]]
                    <tr>
                        <td style="border-color:006A4D;">[[.Name]]</td>
                        <td style="border-color:006A4D;">[[.Tenant]]</td>
                        <td style="border-color:006A4D;">[[.UnitStatus]]</td>
                        <td style="border-color:006A4D;">[[.LeaseExpiryDate.MonthName]] [[.LeaseExpiryDate.Year]]</td>
                        <td style="border-color:006A4D;">{{[[printf "%.0f" .ERVArea]] | number:0}}</td>
                        <td style="border-color:006A4D;">{{[[printf "%.0f" .ERVAmount]] | number:0}}</td>
                        <td style="border-color:006A4D;">{{[[printf "%.0f" .PassingRent]] | number:0}}</td>
                        <td style="border-color:006A4D;">{{[[.Probability]] *100 | number:2}}%</td>
                        <td style="border-color:006A4D;">
                            <button href="#unitcf" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#unitcf" ng-click="getUnitCF('[[.MasterID]]')">Cash Flow</button>
                        </td>
                        <td>
                            <button href="#rentschedule" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#rentschedule" ng-click="getRentSchedule('[[.MasterID]]'[[if .Parent.MC]],[[.Parent.MasterID]][[end]])">Rent Schedule [[if .Parent.MC]][[.Parent.MasterID]][[end]]</button>
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
		<div class="container-fluid">
			[[template "AddChildUnitModal" .Name]]
		</div>
    [[end]]
[[end]]

[[template "UnitTable" .entity]]
