[[define "EntityTable"]]
    <div class="container-fluid shadow-lg rounded" style="width: 95%">
        <table class="table table-hover tableFixHead rounded">
            <thead>
                <tr>
                    <th scope="col">Name</th>
                    <th scope="col">Parent</th>
                    <th scope="col">Yield</th>
                    <th scope="col">OpEx</th>
                    <th scope="col">Debt</th>
                    <th scope="col">IRR</th>
                    <th scope="col">Equity Multiple</th>
                    <th scope="col">Start Date</th>
                    <th scope="col">Sales Date</th>
                </tr>
            </thead>
            <tbody>
                [[range .]]
                <tr>
                    <td> 
                        <div class="post-content">
                            [[if (gt .ParentID 0)]] 
                            <form id="query_form" class="form-horizontal form-well" role="form" action="/ViewEntity" method="get">
                                <div>
                                    <button type="submit" class="btn" id="name" name="name"  Value="[[.Name]]">[[.Name]]</button>
                                </div>
                            </form>[[end]]
                            [[if (eq .ParentID 0)]] 
                            <form id="query_form" class="form-horizontal form-well" role="form" action="/Fund" method="get">
                                <div>
                                    <button type="submit" class="btn" id="name" name="name"  Value="[[.Name]]">[[.Name]]</button>
                                </div>
                            </form>[[end]]                            
                        </div>
                    </td>
                    <td style="border-color:006A4D;">[[if .Parent]] [[.Parent.Name]] [[end]]</td>
                    <td style="border-color:006A4D;">[[.Valuation.EntryYield]]</td>
                    <td style="border-color:006A4D;">[[.OpEx]]</td>
                    <td style="border-color:006A4D;">[[.DebtInput]]</td>
                    <td style="border-color:006A4D;">[[.Metrics.IRR.NetLeveredAfterTax]]</td>
                    <td style="border-color:006A4D;">[[.Metrics.EM.NetLeveredAfterTax]]</td>
                    <td style="border-color:006A4D;">[[.StartDate.Year]]/[[.StartDate.Month]]</td>
                    <td style="border-color:006A4D;">[[.SalesDate.Year]]/[[.SalesDate.Month]]</td>
                </tr>
                [[end]]
            </tbody>
        </table>
    </div>
[[end]]