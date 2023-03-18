[[define "Chart"]]
<div id="coaChart" style="height:600px;" preserveAspectRatio="true">

</div>
<script>

    [[range .coaArray]]
    var [[.Index]] = {
        x: [[.X]],
        y: [[.Y]],
        name: [[.Name]],
        mode: 'lines',
        connectgaps: true,
    };
    [[end]]
    var layout = {
        showlegend: true
    };

//    finalData = [ [[range .coaArray]][[.Index]],[[end]] ]
    finalData = [ [[range .coaArray]][[.Index]],[[end]] ]
    Plotly.newPlot('coaChart', finalData, layout,{displayModeBar: true});


</script>
[[end]]

[[template "Chart" .]]