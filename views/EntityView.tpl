[[template "header" .]]
<body ng-controller="assetViewController">
	<br>
	<br>
	<div class="container-fluid" style="width: 95%;">
        <ul class="nav nav-pills" id="SettingsSPA">
            <li class="nav-item">
                <a class="nav-link active" id="[[.entity.Name]]-tab" data-toggle="tab" href="#[[.entity.Name]]" role="tab" aria-controls="[[.entity.Name]]" aria-selected="true" ng-click="getSettings('[[.entity.Name]]')">[[.entity.Name]]</a>
            </li>
            [[if .entity.ChildEntities]]
                [[range .entity.ChildEntities]]
                    <li class="nav-item">
                        <a class="nav-link" id="[[.Name]]-tab" data-toggle="tab" href="#[[.Name]]" role="tab" aria-controls="[[.Name]]" aria-selected="true" ng-click="getSettings('[[.Name]]')">[[.Name]]</a>
                    </li>
                [[end]]
            [[end]]
        </ul>
    </div>
    <div bind-html-compile = settingsResponse></div>
    </div>
	[[if .entity.Table]]
		<div class="container-fluid" style="width: 95%;">
			<ul class="nav nav-pills" id="CFtabletabs[[.entity.Name]]">
				<li class="nav-item">
					<a class="nav-link active" id="cf-tab" data-toggle="tab" href="#cf" role="tab" aria-controls="cf" aria-selected="true" ng-click="getRequest('cf')">Cash Flow</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="endcash-tab" data-toggle="tab" href="#endcash" role="tab" aria-controls="endcash" aria-selected="false" ng-click="getRequest('endcash')">Ending Cash</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('cashbalance')">Cash Balance</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('endncf')">Ending NCF</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('ncf')">Net Cash Flow</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('endmarketvalue')">End Market Value</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('marketvalue')">Market Value</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('irr')">IRR</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('em')">Equity Multiple</a>
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
			<div id="cftable"></div>
				<div bind-html-compile = response></div>			
			</div>
		</div>
		<br>
		<br>
	[[end]]
		<div id="unitTable" bind-html-compile = unittable></div>			
	[[if not .entity.ChildEntities]]
		<div class="container-fluid">
			[[template "AddChildUnitModal" .entity.Name]]
		</div>
	[[end]]
	[[template "footer" .]]
</body>
