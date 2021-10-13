[[define "RentSchedule"]]
<div class="container-fluid">
    <table class="table table-hover tableFixHead rounded">
        <thead>
            <tr>
                <th scope="col">EXT Number</th>
                <th scope="col">Start Date</th>
                <th scope="col">Vacancy End</th>
                <th scope="col">Renewal Rent Incentives End</th>
                <th scope="col">Rotation Rent Incentives End</th>
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
        [[range .]]
            <tr>
                <td style="border-color:006A4D;">[[.EXTNumber]]</td>
                <td style="border-color:006A4D;">[[.StartDate.MonthName]] [[.StartDate.Year]]</td>
                <td style="border-color:006A4D;">[[.VacancyEnd.MonthName]] [[.VacancyEnd.Year]]</td>
                <td style="border-color:006A4D;">[[.RentIncentivesEndRenew.MonthName]] [[.RentIncentivesEndRenew.Year]]</td>
                <td style="border-color:006A4D;">[[.RentIncentivesEndRotate.MonthName]] [[.RentIncentivesEndRotate.Year]]</td>
                <td style="border-color:006A4D;">[[.DefaultDate.MonthName]] [[.DefaultDate.Year]]</td>
                <td style="border-color:006A4D;">[[.EndDate.MonthName]] [[.EndDate.Year]]</td>
                <td style="border-color:006A4D;">[[.OriginalEndDate.MonthName]] [[.OriginalEndDate.Year]]</td>
                <td style="border-color:006A4D;">{{[[.RenewRent]] *12 | number:0}}</td>
                <td style="border-color:006A4D;">{{[[.RotateRent]] *12 | number:0}}</td>
                <td style="border-color:006A4D;">{{[[.PassingRent]] *12 | number:0}}</td>
                <td style="border-color:006A4D;">{{[[.EndContractRent]] *12 | number:0}}</td>
                <td style="border-color:006A4D;">{{[[.RentRevisionERV]] *100 | number:2}}%</td>
                <td style="border-color:006A4D;">{{[[.Probability]] *100 | number:2}}%</td>
            </tr>
        [[end]]
        </tbody>
    </table>
</div>
[[end]]
[[template "RentSchedule" .data]]