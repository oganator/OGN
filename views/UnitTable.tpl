{{define "UnitTable"}}
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
                </tr>
            </thead>
            <tbody>
                {{$name := .Name}}
                {{range .ChildUnits}}
                <tr>
                    <td>
                        <div class="post-content">
                            <form id="query_form" class="form-horizontal form-well" role="form" action="/ViewUnit" method="post">
                                <input type="text" hidden class="form-control" id="parent" name="parent" value="{{$name}}">
                                <div>
                                    <button type="submit" class="btn" id="name" name="name" style="background-color: #006A4D; color:#FFFFFF;" Value="{{.Name}}">{{.Name}}</button>
                                </div>
                            </form>
                        </div>
                    </td>
                    <td style="border-color:006A4D;">{{.UnitStatus}}</td>
                    <td style="border-color:006A4D;">{{.PercentSoldRent}}</td>
                    <td style="border-color:006A4D;">{{.LeaseExpiryDate.Year}}/{{.LeaseExpiryDate.Month}}</td>
                    <td style="border-color:006A4D;">{{printf "%.0f" .ERVArea}}</td>
                    <td style="border-color:006A4D;">{{printf "%.0f" .ERVAmount}}</td>
                    <td style="border-color:006A4D;">{{printf "%.0f" .PassingRent}}</td>
                    <td style="border-color:006A4D;">{{.Probability}}</td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
{{end}}
