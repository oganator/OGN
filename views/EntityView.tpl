[[template "header" .]]
<body ng-controller="testController">
	<br>
	<br>
	<div class="container-fluid">
		[[template "EntitySettings" .]]
	</div>
	[[if .entity.Table]]
		<div class="container-fluid" style="width: 95%;">
			<ul class="nav nav-pills" id="myTab">
				<li class="nav-item">
					<a class="nav-link active" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="true" ng-click="getRequest('cf')">Cash Flow</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('endcash')">Ending Cash</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('endncf')">Ending NCF</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('irr')" ng-show="strategy == 'Standard' ">IRR</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('em')" ng-show="strategy == 'Standard' ">Equity Multiple</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('ytm')" ng-show="strategy != 'Standard' ">YTM</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('duration')" ng-show="strategy != 'Standard' ">Duration</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="details-tab" data-toggle="tab" href="#details" role="tab" aria-controls="details" aria-selected="false" ng-click="getRequest('details')">Details</a>
				</li>
			</ul>
		</div>
		<div class="container-fluid shadow-lg rounded" style="width: 95%; overflow-x:scroll;">
			<div id="cftable">[[template "CFTable" .]]</div>
			<div bind-html-compile = response></div>
			</div>
		</div>
		<br>
		<br>
	[[end]]
	[[if .entity.ChildEntities]]
		<div class="container-fluid">
			[[template "EntityTable" .entity.ChildEntities]]
		</div>
	[[end]]
	[[if not .entity.ChildUnits]]
		<div class="container-fluid">
			[[template "AddChildEntityModal" .entity]]
		</div>
	[[end]]
	[[if .entity.ChildUnits]]
		<div class="container-fluid">
			[[template "UnitTable" .entity]]
		</div>
	[[end]]
	[[if not .entity.ChildEntities]]
		<div class="container-fluid">
			[[template "AddChildUnitModal" .entity.Name]]
		</div>
	[[end]]
	<div class="row">
		<div id="myDiv2"></div>
	</div>
<!--	<script>

		Plotly.d3.json('https://raw.githubusercontent.com/plotly/datasets/master/3d-ribbon.json', function(figure){
			var i
			var data = [];
			var length = figure.data.length;
			for (i=0;i<length; i++){
				data[i] = {
					x: figure.data[i].x, 
					y: figure.data[i].y, 
					z: figure.data[i].z,
					name: '',
					colorscale: figure.data[i].colorscale,
					type: 'surface',
					showscale: false
				}
			}
			Plotly.newPlot('myDiv', data, layout);
		});
		-->
	<script>
		var layout = {
			title: 'Ribbon Plot',
			showlegend: false,
			autosize: false,
			width: 2000,
			height: 1000,
			scene: {
				xaxis: {title: 'Probability'},
				yaxis: {title: 'Time'},
				zaxis: {title: '$'}
			}
		};
		var figure = [[.entity.MCResults.NCF]]
		var i
		var data = [];
		var length = figure.length;
		for (i=0;i<length; i++){
			data[i] = {
				x: figure[i].x, 
				y: figure[i].y, 
				z: figure[i].z,
				name: '',
				type: 'surface',
				showscale: false
			}
		}
		Plotly.newPlot('myDiv2', data, layout);
	</script>
	[[template "footer" .]]
</body>
