
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
    <div class="container-fluid" style="width: 50%;">
        <canvas id="myChart" class="chart chart-bar"></canvas>
    </div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.3/Chart.bundle.min.js"></script>
    <script>
        function alphaCheck(muDiv, sigmaDiv){
            var mu = document.getElementById(muDiv).value;
            var sigma = document.getElementById(sigmaDiv).value;
            var alpha = ((1-mu)/Math.pow(sigma,2) - 1/mu) * Math.pow(mu, 2);
            if (alpha <= 0) {
                var mtwo = Math.pow(mu,2);
                var mthree = Math.pow(mu,3);
                var x = (mtwo-mthree);
                var alphaZero = Math.pow(x/mu,.5)-.00001;
                document.getElementById(sigmaDiv).value = alphaZero;
            }
        }

        function sharpe(){
            var rfr = document.getElementById('riskfree').value;
            var mean = [[.data.Mean]];
            var stdev = [[.data.StDev]];
            var sr = (mean - rfr)/stdev;
            document.getElementById('sharpe').innerHTML = sr.toFixed(2);
        }

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
