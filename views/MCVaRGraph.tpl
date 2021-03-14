[[define "VaRGraph"]]
<div id="varGraph" style="height:600px" preserveAspectRatio="true">
</div>
    <script>
        var one = {
            x: [[.varp.One.X]],
            y: [[.varp.One.Y]],
            mode: 'lines',
            connectgaps: true,
            name: '1%'
        };

        var five = {
            x: [[.varp.Five.X]],
            y: [[.varp.Five.Y]],
            mode: 'lines',
            connectgaps: true,
            name: '5%'
        };

        var ten = {
            x: [[.varp.Ten.X]],
            y: [[.varp.Ten.Y]],
            mode: 'lines',
            connectgaps: true,
            name: '10%'
        };

        var twentyfive = {
            x: [[.varp.TwentyFive.X]],
            y: [[.varp.TwentyFive.Y]],
            mode: 'lines',
            connectgaps: true,
            name: '25%'
        };

        var fifty = {
            x: [[.varp.Fifty.X]],
            y: [[.varp.Fifty.Y]],
            mode: 'lines',
            connectgaps: true,
            name: '50%'
        };

        var data = [one, five, ten, twentyfive, fifty];

        var layout = {
            showlegend: true
        };

        Plotly.newPlot('varGraph', data, layout,{displayModeBar: true});
    </script>
[[end]]