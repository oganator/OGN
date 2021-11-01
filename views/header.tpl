[[define "header"]]
    <!DOCTYPE html>
    <html lang="en" ng-app="ognApp">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
        <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.6.3/css/all.css" integrity="sha384-UHRtZLI+pbxtHCWp1t77Bi1L4ZtiqrqD80Kn4Z8NTSRyMA2Fd33n5dQ8lWUE00s/" crossorigin="anonymous">
        <link rel="stylesheet" href="https://unpkg.com/bootstrap-table@1.18.1/dist/bootstrap-table.min.css">
        <link rel="stylesheet" type="text/css" href="static/css/misc.css">
        <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
    </head>
    <body>
        <nav class="first-second navbar navbar-expand-lg navbar-dark shadow-lg">
            <a class="navbar-brand" style="color:white;"> Ogan Capital &nbsp &nbsp &nbsp</a>
             <form id="query_form" class="form-horizontal form-well" role="form" action="/" method="get">
                <div>
                    <button type="submit" class="btn">Home</button>
                </div>
            </form>
        </nav>
    </body>
[[end]]
[[template "footer"]]