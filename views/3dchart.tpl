[[define "ribbon"]]
<div class="row">
	<div id="myDiv2"></div>
</div>
<script>
	var layout = {
		showlegend: false,
		autosize: true,
		width: 2200,
		height: 1500,
		scene: {
			xaxis: {title: 'Probability'},
			yaxis: {title: 'Time'},
			zaxis: {title: '$'}
		}
	};
	Plotly.newPlot('myDiv2', [[.]], layout);
</script>
[[end]]

[[template "ribbon" .data]]