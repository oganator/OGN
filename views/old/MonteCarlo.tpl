
    <table class="table tableFixHead rounded">
        <thead>
            <tr>
                <th scope="col" style="width: 12.5%">Mean</th>
                <th scope="col" style="width: 12.5%">Standard Deviation</th>
                <th scope="col" style="width: 12.5%">Skew</th>
                <th scope="col" style="width: 12.5%">Excess Kurtosis</th>
                <th scope="col" style="width: 12.5%">Regression</th>
                <th scope="col" style="width: 12.5%">Risk Free Rate</th>
                <th scope="col" style="width: 12.5%">Sharpe Ratio</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>[[printf "%.2f" .Mean]]</td>
                <td>[[printf "%.2f" .StDev]]</td>
                <td>[[printf "%.2f" .Skew]]</td>
                <td>[[printf "%.2f" .Kurtosis]]</td>
                <td>[[printf "%.2f" .LRalpha]] + [[printf "%.2f" .LRbeta]]x</td>
                <td>
                    <div class="form-group col-xs-1">
                        <input onchange="sharpe()" type="text" class="form-control border-top-0 border-right-0 border-left-0" id="riskfree" name="riskfree">
                    </div>
                </td>
                <td id="sharpe"></td>
            </tr>
        </tbody>
    </table>
    <canvas id="myChart"></canvas>
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
            var mean = [[.Mean]];
            var stdev = [[.StDev]];
            var sr = (mean - rfr)/stdev;
            document.getElementById('sharpe').innerHTML = sr.toFixed(2);
        }

        var ctx = document.getElementById('myChart').getContext('2d');
        var keys = [[.Hist.Keys]];
        var myChart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: keys,
                datasets: [{
                    label: 'IRR',
                    data: [[.Hist.Vals]],
                    borderWidth: 1
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


[[define "MonteCarloDetails2"]]
    <table class="table table-hover tableFixHead rounded" data-toggle="table" data-search="true" data-pagination="true">
        <thead>
            <tr>
                <th scope="col" data-sortable="true">Index</th>
                <th scope="col" data-sortable="true">IRR</th>
                <th scope="col" data-sortable="true">YTM</th>
                <th scope="col" data-sortable="true">Duration</th>
                <th scope="col" data-sortable="true">Void</th>
                <th scope="col" data-sortable="true">Extension Duration</th>
                <th scope="col" data-sortable="true">Hazard Rate</th>
                <th scope="col" data-sortable="true">OpEx(%CR)</th>
                <th scope="col" data-sortable="true">CPI Growth</th>
                <th scope="col" data-sortable="true">ERV Growth</th>
                <th scope="col" data-sortable="true">Yield Shift</th>
            </tr>
        </thead>
        <tbody>
            [[range $index, $value := .]]
            <tr>
                <td>[[$index]]
                <td>[[printf "%.4f" $value.Metrics.IRR.NetLeveredAfterTax]]</td>
                <td>[[printf "%.4f" $value.Metrics.BondHolder.YTM]]</td>
                <td>[[printf "%.4f" $value.Metrics.BondHolder.Duration]]</td>
                <td>[[$value.GLA.Void]]</td>
                <td>[[$value.GLA.EXTDuration]]</td>
                <td>[[printf "%.2f" $value.GLA.Default.Hazard]]</td>
                <td>[[printf "%.2f" $value.OpEx.PercentOfTRI]]</td>
                [[range $value.GrowthInput]]
                <td>[[printf "%.3f" .]]</td>
                [[end]]
                <td>[[printf "%.2f" $value.Valuation.YieldShift]]</td>
            </tr>
            [[end]]
        </tbody>
    </table>
[[end]]
