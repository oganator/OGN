[[define "rentschedule"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col"></th>
                    [[range .]]
                    <th scope="col">Extension [[.EXTNumber]]
                    [[end]]
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th scope="row">Start Date</th>
                    [[range .]]
                    <td>[[.StartDate.Month]]/[[.StartDate.Year]]</td>
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">End Date Vacancy</th>
                    [[range .]]
                    <td>[[.VacancyEnd.Month]]/[[.VacancyEnd.Year]]</td>
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">Default Date</th>
                    [[range .]]
                    <td>[[.DefaultDate.Month]]/[[.DefaultDate.Year]]</td>
                    [[end]]
                </tr>					
                <tr>
                    <th scope="row">Lease End Date</th>
                    [[range .]]
                    <td>[[.EndDate.Month]]/[[.EndDate.Year]]</td>
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">Renewal Rent</th>
                    [[range .]]
                    <td>[[printf "%.0f" .RenewRent]]</td>
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">Rotation Rent</th>
                    [[range .]]
                    <td>[[printf "%.0f" .RotateRent]]</td>
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">Ending Contract Rent</th>
                    [[range .]]
                    <td>[[printf "%.0f" .EndContractRent]]</td>
                    [[end]]
                </tr>
                <tr>
                    <th scope="row">Base Rent</th>
                    [[range .]]
                    <td>[[printf "%.0f" .CFRent]]</td>
                    [[end]]
                </tr>                
            </tbody>
        </table>
    </div>
[[end]]

