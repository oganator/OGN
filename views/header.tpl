[[define "header"]]
    <!DOCTYPE html>
    <html lang="en" ng-app="ognApp">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=yes">
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
        <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.6.3/css/all.css" integrity="sha384-UHRtZLI+pbxtHCWp1t77Bi1L4ZtiqrqD80Kn4Z8NTSRyMA2Fd33n5dQ8lWUE00s/" crossorigin="anonymous">
        <link rel="stylesheet" href="https://unpkg.com/bootstrap-table@1.18.1/dist/bootstrap-table.min.css">
        <link rel="stylesheet" type="text/css" href="static/css/misc.css">
        <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
    </head>
    <body>
        <nav class="first-second navbar navbar-expand-lg navbar-dark shadow-lg">
            <a class="navbar-brand" style="color:white;">LaSalle Investment Management &nbsp &nbsp &nbsp</a>
            <form class="form-horizontal form-well" role="form" method="post" action="/ChangeEntity">
                <div class="tab-pane fade in show active" id="header" role="tabpanel">
                    <div class="form-row">
                        <div class="form-group col">
                            <select class="form-select" size="1" id="modelname" name="modelname" value="[[.entity.Name]]">
                                <option selected hidden>[[.entity.Name]]</option>
                                [[range $key,$value := .modelslist]]
                                <option>[[$key]]</option>
                                [[end]]
                            </select>
                        </div>
                        <div class="form-group col-md-1 d-flex align-items-end">
                            <button type="submit" class="btn">Submit</button>
                        </div>
                    </div>
                </div>
            </form>
        </nav>
    </body>
[[end]]
[[template "footer"]]