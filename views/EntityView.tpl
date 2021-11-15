[[template "header" .]]
<body ng-controller="assetViewController">
	<br>
	<br>
	<div class="container-fluid" style="width: 95%;">
        <ul class="nav nav-pills" id="SettingsSPA">
            [[if .entity.ChildEntities]]
				<li class="nav-item">
					<a class="nav-link active" id="[[.entity.Name]]-tab" data-toggle="tab" href="#[[.entity.Name]]" role="tab" aria-controls="[[.entity.Name]]" aria-selected="true" ng-click="getSettings('[[.entity.Name]]')">[[.entity.Name]]</a>
				</li>
                [[range .entity.ChildEntities]]
                    <li class="nav-item">
                        <a class="nav-link" id="[[.Name]]-tab" data-toggle="tab" href="#[[.Name]]" role="tab" aria-controls="[[.Name]]" aria-selected="true" ng-click="getSettings('[[.Name]]')">[[.Name]]</a>
                    </li>
                [[end]]
            [[end]]
            [[if not .entity.ChildEntities]]
				<li class="nav-item">
					<a class="nav-link" id="[[.entity.Parent.Name]]-tab" data-toggle="tab" href="#[[.entity.Parent.Name]]" role="tab" aria-controls="[[.entity.Parent.Name]]" aria-selected="true" ng-click="getSettings('[[.entity.Parent.Name]]')">[[.entity.Parent.Name]]</a>
				</li>
				[[$entityname := .entity.Name]]
                [[range .entity.Parent.ChildEntities]]
					[[if eq .Name $entityname]]
						<li class="nav-item">
							<a class="nav-link active" id="[[.Name]]-tab" data-toggle="tab" href="#[[.Name]]" role="tab" aria-controls="[[.Name]]" aria-selected="true" ng-click="getSettings('[[.Name]]')">[[.Name]]</a>
						</li>[[end]]
					[[if ne .Name $entityname]]
						<li class="nav-item">
							<a class="nav-link" id="[[.Name]]-tab" data-toggle="tab" href="#[[.Name]]" role="tab" aria-controls="[[.Name]]" aria-selected="true" ng-click="getSettings('[[.Name]]')">[[.Name]]</a>
						</li>[[end]]
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
					<a class="nav-link" id="endcash-tab" data-toggle="tab" href="#endcash" role="tab" aria-controls="endcash" aria-selected="false" ng-click="getRequest('endcash')">Cash Balance</a>
				</li>
<!--			<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('cashbalance')">Cash Balance 3d</a>
				</li>
-->				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('endncf')">Net Cash Flow</a>
				</li>
<!--			<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('ncf')">Net Cash Flow 3d</a>
				</li>
-->				<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('endmarketvalue')">Market Value</a>
				</li>
<!--			<li class="nav-item">
					<a class="nav-link" id="summary-tab" data-toggle="tab" href="#summary" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('marketvalue')">Market Value 3d</a>
				</li>
-->				<li class="nav-item">
					<a class="nav-link" id="irrtabletab" data-toggle="tab" href="#irrtabletab" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('irr')">IRR</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="emtabletab" data-toggle="tab" href="#emtabletab" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('em')">Equity Multiple</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="details-tab" data-toggle="tab" href="#details" role="tab" aria-controls="details" aria-selected="false" ng-click="getRequest('details')">Details</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="ytmtabletab" data-toggle="tab" href="#ytmtabletab" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('ytm')">YTM</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="durationtabletab" data-toggle="tab" href="#durationtabletab" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('duration')">Duration</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" id="ytmdurtabletab" data-toggle="tab" href="#ytmdurtabletab" role="tab" aria-controls="summary" aria-selected="false" ng-click="getRequest('ytmdur')">YTM/DUR</a>
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
	[[template "footer" .]]
</body>
