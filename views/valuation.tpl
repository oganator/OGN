[[define "valuation"]]
    <br>
    <div class="container-fluid">
        <div class="row">
            <div class="col-sm-2">
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">Next 12 Months Rent</label>
                        <input type ="text" readonly class="form-control" id="rent" name="rent" value="[[printf "%.0f" .Model.Valuation.First12MonthsRent]]">
                    </div>
                </div>
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">Entry Yield</label>
                        <input type ="text" class="form-control" id="entryyield" name="entryyield" value="[[.Model.Valuation.Yield.Entry]]">
                    </div>
                </div>
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">Gross Acquisition</label>
                        <input type ="text" readonly class="form-control" id="gacq" name="gacq" value="">
                    </div>
                </div>
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">Net Acquisition</label>
                        <input type ="text" readonly class="form-control" id="nacq" name="nacq" value="">
                    </div>
                </div>
            </div>
            <div class="form-group col-sm-2">
                <table class="table table-hover rounded tableFixHead" id="fees">
                    <thead>
                        <tr>
                            <th scope="col-sm-1">Fee</th>
                            <th scope="col-sm-1">Cost</th>
                            <th scope="col-sm-1">Percent</th>
                        </tr>
                    </thead>
                    <tbody>
                        
                    </tbody>
                </table>
            </div>
            <div class="form-group col-sm-2">
                <table class="table table-hover rounded tableFixHead" id="capincome">
                    <thead>
                        <tr>
                            <th scope="col-sm-1">COA Item</th>
                            <th scope="col-sm-1">Capitalization %</th>
                            <th scope="col-sm-1">Deduction %</th>
                        </tr>
                    </thead>
                    <tbody>
                    [[$val := .Model]]
                    [[range .Model.Table]]
                    [[if ne .COA ""]]
                    [[if ne .COA "Market Value"]]
                    [[if ne .COA "Total Area"]]
                        <tr>
                            <td>[[.COA]]</td>
                            [[if eq .COA "Total ERV"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.TotalERV]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.TotalERV]]></td>
                            [[end]]
                            [[if eq .COA "Occupied ERV"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.OccupiedERV]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.OccupiedERV]]></td>
                                [[end]]
                            [[if eq .COA "Vacant ERV"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.VacantERV]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.VacantERV]]></td>
                            [[end]]
                            [[if eq .COA "Top Slice"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.TopSlice]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.TopSlice]]></td>
                                [[end]]
                            [[if eq .COA "Passing Rent"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.PassingRent]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.PassingRent]]></td>
                                [[end]]
                            [[if eq .COA "Indexation"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.Indexation]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.Indexation]]></td>
                                [[end]]
                            [[if eq .COA "BPUplift"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.BPUplift]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.BPUplift]]></td>
                                [[end]]
                            [[if eq .COA "Theoretical Rental Income"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.TheoreticalRentalIncome]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.TheoreticalRentalIncome]]></td>
                                [[end]]
                            [[if eq .COA "Vacancy"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.Vacancy]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.Vacancy]]></td>
                                [[end]]
                            [[if eq .COA "Contract Rent"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.ContractRent]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.ContractRent]]></td>
                                [[end]]
                            [[if eq .COA "Rent Free"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.RentFree]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.RentFree]]></td>
                                [[end]]
                            [[if eq .COA "Operating Income"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.OperatingIncome]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.OperatingIncome]]></td>
                                [[end]]
                            [[if eq .COA "Capex"]]
                                <td><input type ="text" id="[[.COA]]" name="Cap [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeCapSetup.Capex]]></td>
                                <td><input type ="text" id="[[.COA]]" name="Deduction [[.COA]]" class="form-control" value=[[$val.Valuation.IncomeDeduction.Capex]]></td>
                                [[end]]
                        </tr>
                    [[end]]
                    [[end]]
                    [[end]]
                    [[end]]
                    </tbody>
                </table>
            </div>
            <div class="col-sm-2">
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">Exit Yield</label>
                        <input type ="text" class="form-control" id="exityield" name="exityield" value="[[.Model.Valuation.Yield.Exit]]">
                    </div>
                </div>
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">PCs</label>
                        <input type ="text" readonly class="form-control" id="gacq" name="gacq" value="">
                    </div>
                </div>
                <div class="form-group">
                    <div class="form-group">
                        <label for="name" class="bmd-label-floating">Net Acquisition</label>
                        <input type ="text" readonly class="form-control" id="nacq" name="nacq" value="">
                    </div>
                </div>
            </div>            
        </div>
    </div>
[[end]]
