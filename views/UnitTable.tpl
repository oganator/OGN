[[define "UnitTable"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%" >
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col">Unit Name</th>
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
                    <td>
                        <div class="post-content">
                            <form id="query_form" class="form-horizontal form-well" role="form" action="/ViewUnit" method="post">
                                <input type="text" hidden class="form-control" id="parent" name="parent" value="[[$name]]">
                                <div>
                                    <button type="submit" class="btn" id="name" name="name" style="background-color: #006A4D; color:#FFFFFF;" Value="[[.Name]]">[[.Name]]</button>
                                </div>
                            </form>
                        </div>
                    </td>
                    <td style="border-color:006A4D;">[[.UnitStatus]]</td>
                    <td style="border-color:006A4D;">[[.PercentSoldRent]]</td>
                    <td style="border-color:006A4D;">[[.LeaseExpiryDate.Year]]/[[.LeaseExpiryDate.Month]]</td>
                    <td style="border-color:006A4D;">[[printf "%.0f" .ERVArea]]</td>
                    <td style="border-color:006A4D;">[[printf "%.0f" .ERVAmount]]</td>
                    <td style="border-color:006A4D;">[[printf "%.0f" .PassingRent]]</td>
                    <td style="border-color:006A4D;">[[.Probability]]</td>
                    <td style="border-color:006A4D;">
                        <div class="text-center">
                            <a href="#rentschedule" class="btn btn-default btn-rounded my-3" data-toggle="modal" data-target="#rentschedule">View</a>
                        </div>
                    </td>
                </tr>
                <!-- Modal -->
                <div class="modal fade" id="rentschedule" tabindex="-1" role="dialog" aria-labelledby="myrentschedule" aria-hidden="true">
                    <div class="modal-dialog modal-notify modal-success modal-fluid modal-dialog-centered" role="document">
                        <div class="modal-content">
                            [[template "RentSchedule" .]]
                        </div>
                    </div>
                </div>
                [[end]]
            </tbody>
        </table>
    </div>
[[end]]

[[define "RentSchedule"]]
<div class="container-fluid">
    <table class="table table-hover tableFixHead rounded">
        <thead>
            <tr>
                <th scope="col">EXT Number</th>
                <th scope="col">Start Date</th>
                <th scope="col">Vacancy End</th>
                <th scope="col">Rent Incentives End</th>
                <th scope="col">Default Date</th>
                <th scope="col">End Date</th>
                <th scope="col">Original End Date</th>
                <th scope="col">Renew Rent</th>
                <th scope="col">Rotate Rent</th>
                <th scope="col">Passing Rent</th>
                <th scope="col">Ending Contract Rent</th>
                <th scope="col">Rent Revision ERV</th>
                <th scope="col">Probability</th>
            </tr>
        </thead>
        <tbody>
            [[range .RSStore]]
            <tr>
                <td style="border-color:006A4D;">[[.EXTNumber]]</td>
                <td style="border-color:006A4D;">[[.StartDate.MonthName]] [[.StartDate.Year]]</td>
                <td style="border-color:006A4D;">[[.VacancyEnd.MonthName]] [[.VacancyEnd.Year]]</td>
                <td style="border-color:006A4D;">[[.RentIncentivesEnd.MonthName]] [[.RentIncentivesEnd.Year]]</td>
                <td style="border-color:006A4D;">[[.DefaultDate.MonthName]] [[.DefaultDate.Year]]</td>
                <td style="border-color:006A4D;">[[.EndDate.MonthName]] [[.EndDate.Year]]</td>
                <td style="border-color:006A4D;">[[.OriginalEndDate.MonthName]] [[.OriginalEndDate.Year]]</td>
                <td style="border-color:006A4D;">[[.RenewRent]]</td>
                <td style="border-color:006A4D;">[[.RotateRent]]</td>
                <td style="border-color:006A4D;">[[.PassingRent]]</td>
                <td style="border-color:006A4D;">[[.EndContractRent]]</td>
                <td style="border-color:006A4D;">[[.RentRevisionERV]]</td>
                <td style="border-color:006A4D;">[[.Probability]]</td>
            </tr>
            [[end]]
        </tbody>
    </table>
</div>
[[end]]
