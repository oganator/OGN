{{define "header"}}
<!DOCTYPE html>
<html lang="en">
  <head>
	<meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    <link rel="stylesheet" type="text/css" href="static/css/misc.css">
    <link href="https://unpkg.com/bootstrap-table@1.16.0/dist/bootstrap-table.min.css" rel="stylesheet">
    <link href="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.3/Chart.min.css" rel="stylesheet">
  </head>
</html>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark shadow-lg" style="background-color: #006A4D;">
        <a class="navbar-brand" href="#">OGAN CAPITAL &nbsp &nbsp &nbsp</a>
        <form class="form-horizontal form-well" style="color: #006A4D;" role="form" method="post" action="/ChangeEntity">
            <div class="tab-pane fade in show active" id="header" role="tabpanel">
                <div class="form-row">
                    <div class="form-group col">
                        <label for="head" class="bmd-label-floating">Name</label>
                        <select type="text" class="form-control" id="modelname" name="modelname" value="{{.entity.Name}}">
                            <option>{{.entity.Name}}</option>
                            {{range $key,$value := .modelslist}}
                            <option>{{$key}}</option>
                            {{end}}
                        </select>
                    </div>
                    <div class="form-group col-md-4 d-flex align-items-end">
                        <button type="submit" class="btn" style="background-color: #006A4D; color:white">Submit</button>
                    </div>
                </div>
            </div>
        </form>
    </nav>
</body>
{{end}}
{{template "footer"}}