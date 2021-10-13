    <table class="table tableFixHead rounded">
        <thead>
            <tr>
                <th scope="col" style="width: 12.5%">Mean</th>
                <th scope="col" style="width: 12.5%">Standard Deviation</th>
                <th scope="col" style="width: 12.5%">Skew</th>
                <th scope="col" style="width: 12.5%">Excess Kurtosis</th>
                <th scope="col" style="width: 12.5%">Regression</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>{{[[.data.Mean]] | number:2}}</td>
                <td>{{[[.data.StDev]] | number:2}}</td>
                <td>{{[[.data.Skew]] | number:2}}</td>
                <td>{{[[.data.Kurtosis]] | number:2}}</td>
                <td>{{[[.data.LRalpha]] | number:2}} + {{[[.data.LRbeta]] | number:2}}x</td>
            </tr>
        </tbody>
    </table>
    <table class="table tableFixHead rounded">
        <thead>
            <tr>
                <th scope="col" style="width: 12.5%">P1</th>
                <th scope="col" style="width: 12.5%">P5</th>
                <th scope="col" style="width: 12.5%">P10</th>
                <th scope="col" style="width: 12.5%">P25</th>
                <th scope="col" style="width: 12.5%">P50</th>
                <th scope="col" style="width: 12.5%">P75</th>
                <th scope="col" style="width: 12.5%">P90</th>
                <th scope="col" style="width: 12.5%">P95</th>
                <th scope="col" style="width: 12.5%">P99</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>{{[[.data.P1]] | number:2}}</td>
                <td>{{[[.data.P5]] | number:2}}</td>
                <td>{{[[.data.P10]] | number:2}}</td>
                <td>{{[[.data.P25]] | number:2}}</td>
                <td>{{[[.data.P50]] | number:2}}</td>
                <td>{{[[.data.P75]] | number:2}}</td>
                <td>{{[[.data.P90]] | number:2}}</td>
                <td>{{[[.data.P95]] | number:2}}</td>
                <td>{{[[.data.P99]] | number:2}}</td>
            </tr>
        </tbody>
    </table>

    <div class="row">
        <div class="container-fluid" style="width: 45%;">
            <canvas id="myChart" class="chart chart-bar"></canvas>
        </div>
        [[if .varp]]
            <div class="container-fluid" style="width: 45%;">
                [[template "VaRGraph" .]]
            </div>  
        [[end]]
    </div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.3/Chart.bundle.min.js"></script>
    <script>
        var ctx = document.getElementById('myChart').getContext('2d');
        var keys = [[.data.Hist.Keys]];
        var myChart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: keys,
                datasets: [{
                    label: '',
                    data: [[.data.Hist.Vals]],
                    backgroundColor: "rgba(0,36,93,.5)", 
                    borderWidth: 0
                }]
            },
            options: {
                scales: {
                    yAxes: [{
                        ticks: {
                            beginAtZero: true
                        }
                    }]
                }
            }
        });
    </script>
