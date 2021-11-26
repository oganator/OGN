[[define "CFTable"]]
        <table class="table table-hover rounded tableFixHead">
            <thead>
                <tr>
                <th scope="col"></th>
                [[$name := .entity.Name]]
                [[$months := .entity.TableHeader.Monthly]]
                [[$years := .entity.TableHeader.Yearly]]
                [[range $years]] 
                    [[if .Bool]]
                        [[$year := .]]
                        [[range $months]] 
                            [[if .Bool]]
                                [[if eq .Year $year.Year]]
                                    <th scope="col" hidden name="[[$name]][[.Year]]">[[.MonthName]]</th> 
                                [[end]] 
                            [[end]] 
                        [[end]]
                        <th scope="col"><button class="btn" onclick="chooseYear('[[$name]][[.Year]]')">[[.Year]]</button></th>
                    [[end]] 
                [[end]]
                </tr>
            </thead>
            <tbody>
                [[range .entity.Table]]
                [[$table := .]]
                    <tr>
                        <th style="min-width: 250px; position: sticky; Left: 0px;" scope="row">[[.COA]]</th>
                        [[range $years]] 
                            [[if .Bool]]
                                [[$year := .]]
                                [[range $months]] 
                                    [[if .Bool]]
                                        [[if eq .Year $year.Year]]
                                            <td hidden name="[[$name]][[.Year]]" class="second-bg">[[index $table.Value .Dateint]]</td>
                                        [[end]] 
                                    [[end]] 
                                [[end]]
                                <td>[[index $table.Value $year.Year]]</td>
                            [[end]]
                        [[end]]
                    </tr>
                [[end]]
            </tbody>
        </table>
    <script>
        function chooseYear(year) {
            var tds = document.getElementsByName(year);
            var i
            for (i = 0; i < tds.length; i++){
                if (tds[i].hidden == false) {
                    tds[i].hidden = true;
                } else {
                    tds[i].hidden = false;
                }
            }
        }
    </script>
[[end]]

[[template "CFTable" .]]