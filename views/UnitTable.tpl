[[define "UnitTable"]]
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col">Unit Name</th>
                    <th scope="col">Tenant</th>
                    <th scope="col">Status</th>
                    <th scope="col">Lease Start</th>
                    <th scope="col">Lease End</th>
                    <th scope="col">ERV Area</th>
                    <th scope="col">ERV Amount</th>
                    <th scope="col">Passing Rent</th>
                    <th scope="col">Probability</th>
                    <th scope="col">Void</th>
                    <th scope="col">Lease <br> Duration</th>
                    <th scope="col">Rent<br>Revision<br> to ERV</th>
                    <th scope="col"></th>
                    <th scope="col"></th>
                </tr>
            </thead>
            <tbody>
                [[$name := .Name]]
                [[range .ChildUnitModels]]
                <tr>
                    <td style="border-color:006A4D;" width="6%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'unit_name')" type="text" class="form-control" id="unit[[.MasterID]].unit_name" value="[[.Name]]"></input></td>
                    <td style="border-color:006A4D;" width="9%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'tenant')" type="text" class="form-control" id="unit[[.MasterID]].tenant" value="[[.Tenant]]"></input></td>
                    <td style="border-color:006A4D;" width="7%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'unit_status')" type="text" class="form-control" id="unit[[.MasterID]].unit_status" value="[[.UnitStatus]]"></input></td>
                    <td style="border-color:006A4D;" width="7%">
                        <div class="d-flex flex-row ">
                            <input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'lease_start_month')" type="text" class="form-control" id="unit[[.MasterID]].lease_start_month" value="[[.LeaseStartDate.Month]]"></input> &nbsp
                            <input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'lease_start_year')"  type="text" class="form-control" id="unit[[.MasterID]].lease_start_year" value="[[.LeaseStartDate.Year]]"></input>
                        </div>
                    </td>
                    <td style="border-color:006A4D;" width="7%">
                        <div class="d-flex flex-row ">
                            <input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'lease_end_month')" type="text" class="form-control" id="unit[[.MasterID]].lease_end_month" value="[[.LeaseExpiryDate.Month]]"></input> &nbsp
                            <input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'lease_end_year')"  type="text" class="form-control" id="unit[[.MasterID]].lease_end_year" value="[[.LeaseExpiryDate.Year]]"></input>
                        </div>
                    </td>
                    <td style="border-color:006A4D;" width="5%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'erv_area')" type="text" class="form-control" id="unit[[.MasterID]].erv_area" value="{{[[ .ERVArea]] | number:0}}"></input></td>
                    <td style="border-color:006A4D;" width="5%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'erv_amount')" type="text" class="form-control" id="unit[[.MasterID]].erv_amount" value="{{[[.ERVAmount]] | number:0}}"></input></td>
                    <td style="border-color:006A4D;" width="9%">
                        <div class="d-flex flex-row ">
                            <input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'passing_rent')" type="text" class="form-control" id="unit[[.MasterID]].passing_rent" value="{{[[.PassingRent]] | number:0}}"></input>
                            <input disabled type="text" class="form-control" id="unit[[.MasterID]].passing_rent_monthly" value="{{[[.PassingRent]] / 12 | number:0}}"></input>
                        </div>                    
                    </td>
                    <td style="border-color:006A4D;" width="5%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'probability')" type="text" class="form-control" id="unit[[.MasterID]].probability" value="{{[[.Probability]]  | number:2}}" [[if ne .Probability .Parent.GLA.Probability]] style="color: red;" [[end]]></input></td>
                    <td style="border-color:006A4D;" width="4%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'void')" type="text" class="form-control" id="unit[[.MasterID]].void" value="{{[[.Void]] | number:0}}" [[if ne .Void .Parent.GLA.Void]] style="color: red;" [[end]]></input></td>
                    <td style="border-color:006A4D;" width="4%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'ext_dur')" type="text" class="form-control" id="unit[[.MasterID]].ext_dur" value="{{[[.EXTDuration]] | number:0}}" [[if ne .EXTDuration .Parent.GLA.EXTDuration]] style="color: red;" [[end]]></input></td>
                    <td style="border-color:006A4D;" width="5%"><input ng-keydown="$event.keyCode === 13 && updateUnit([[.MasterID]],'rent_revision_erv')" type="text" class="form-control" id="unit[[.MasterID]].rent_revision_erv" value="{{[[.RentRevisionERV]]  | number:2}}" [[if ne .RentRevisionERV .Parent.GLA.RentRevisionERV]] style="color: red;" [[end]]></input></td>
                    <td style="border-color:006A4D;" width="15%">
                        <div class="d-flex flex-row">
                            <button href="#unitcf" class="btn btn-default btn-rounded" data-toggle="modal" data-target="#unitcf" ng-click="getUnitCF('[[.MasterID]]')">Cash Flow</button>
                            <button href="#rentschedule" class="btn btn-default btn-rounded" data-toggle="modal" data-target="#rentschedule" ng-click="getRentSchedule('[[.MasterID]]'[[if .Parent.MC]],[[.Parent.MasterID]][[end]])">Rent Schedule [[if .Parent.MC]][[.Parent.MasterID]][[end]]</button>
                        </div>
                    </td>
                </tr>
                [[end]]
            </tbody>
        </table>
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
[[template "UnitTable" .entity]]
