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
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" ng-click="getRequest('irr')">IRR</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('em')">Equity Multiple</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('ytm')">YTM</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('duration')">Duration</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="details-tab" data-toggle="tab" href="#details" role="tab" aria-controls="details" aria-selected="false" ng-click="getRequest('details')">Details</a>
				</li>
			</ul>
		</div>
		<div class="container-fluid shadow-lg rounded" style="width: 95%; overflow-x:scroll;">
			<div id="cftable">[[template "MCDetails" .entity]]</div>
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
	[[template "footer" .]]
</body>