{{define "CFTable"}}
    <div class="container-fluid shadow-lg rounded" style="width: 95%; overflow-x:scroll;">
        <table class="table table-hover rounded tableFixHead">
            <thead>
                <tr>
                <th scope="col"></th>
                {{$months := .TableHeader.Monthly}}
                {{$years := .TableHeader.Yearly}}
                {{range $years}} {{if .Bool}}
                    {{$year := .}}
                    {{range $months}} {{if .Bool}}
                        {{if eq .Year $year.Year}}
                        <th scope="col" hidden name="{{.Year}}" style="background-color:white;color:#006A4D;">{{.MonthName}}</th> 
                        {{end}} 
                    {{end}} {{end}}
                    <th scope="col"><button class="btn" onclick="chooseYear('{{.Year}}')">{{.Year}}</button></th>
                {{end}} {{end}}
                </tr>
            </thead>
            <tbody>
                {{range .Table}}
                {{$table := .}}
                    <tr>
                        <th style="min-width: 250px; position: sticky; Left: 0px; background-color:white;color:#006A4D;" scope="row">{{.COA}}</th>
                        {{range $years}} {{if .Bool}}
                            {{$year := .}}
                            {{range $months}} {{if .Bool}}
                                {{if eq .Year $year.Year}}
                                <td hidden name="{{.Year}}" style="background-color:#006A4D;color:white;">{{index $table.Value .Dateint}}</td>
                                {{end}} 
                            {{end}} {{end}}
                            <td>{{index $table.Value $year.Year}}</td>
                        {{end}} {{end}}
                    </tr>
                {{end}}
            </tbody>
        </table>
    </div>
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
{{end}}
