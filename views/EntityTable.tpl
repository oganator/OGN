{{define "EntityTable"}}
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col">Name</th>
                    <th scope="col">Parent</th>
                    <th scope="col">Level</th>
                    <th scope="col">Yield</th>
                    <th scope="col">OpEx</th>
                    <th scope="col">Debt</th>
                    <th scope="col">Metrics</th>
                    <th scope="col">Start Date</th>
                    <th scope="col">Sales Date</th>
                    <th scope="col">Frequency</th>
                    <th scope="col">Type</th>
                    <th scope="col">Year Span</th>
                </tr>
            </thead>
            <tbody>
                {{range .}}
                <tr>
                    <td> 
                        <div class="post-content">
                            <form id="query_form" class="form-horizontal form-well" role="form" action="/ViewEntity" method="post">
                                <div>
                                    <button type="submit" class="btn" id="name" name="name" style="background-color: #006A4D; color:#FFFFFF;" Value="{{.Name}}">{{.Name}}</button>
                                </div>
                            </form>
                        </div>
                    </td>
                    <td style="border-color:006A4D;">{{if .Parent}} {{.Parent.Name}} {{end}}</td>
                    <td style="border-color:006A4D;">{{.Level}}</td>
                    <td style="border-color:006A4D;">{{.Model.Valuation.Yield}}</td>
                    <td style="border-color:006A4D;">{{.Model.OpEx}}</td>
                    <td style="border-color:006A4D;">{{.Model.DebtInput}}</td>
                    <td style="border-color:006A4D;">{{.Metrics}}</td>
                    <td style="border-color:006A4D;">{{.Model.StartDate.Value.Year}}/{{.Model.StartDate.Value.Month}}</td>
                    <td style="border-color:006A4D;">{{.Model.SalesDate.Value.Year}}/{{.Model.SalesDate.Value.Month}}</td>
                    <td style="border-color:006A4D;">{{.Model.Settings.Frequency}}</td>
                    <td style="border-color:006A4D;">{{.Model.Settings.Type}}</td>
                    <td style="border-color:006A4D;">{{.Model.Settings.YearSpan}}</td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
{{end}}