[[define "footer"]]
</body>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js" integrity="sha512-bLT0Qm9VnAYZDflyKcBaQ2gg0hSYNQrJ8RilYldYQ1FxQYoCLtUjuuRuZo+fjqhx/qtq/1itJ0C2ejDxltZVFg==" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
<script src="https://unpkg.com/bootstrap-table@1.18.1/dist/bootstrap-table.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.3/Chart.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.3/Chart.bundle.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.6.9/angular.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.6.9/angular-resource.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.6.9/angular-route.min.js"></script>
<script src="http://angular-ui.github.io/bootstrap/ui-bootstrap-tpls-0.12.1.min.js" data-require="ui-bootstrap@*" data-semver="0.12.1"></script>
<script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jspdf/2.4.0/jspdf.umd.min.js"></script>
<script src="static/js/app.js"></script>
<script src="static/js/misc.js"></script>
<script>
let ognApp = angular.module('ognApp', []);
	// testController
	ognApp.controller('assetViewController', ['$scope', '$http', '$sce', function($scope, $http, $sce) {
		$scope.baseURL = ""//[[.baseURL]];
		$scope.entity = [[.entity.Name]];
		$scope.settingsTab = "settings";
		$scope.tableTab = "cf";
		$scope.mcdetailspage = 1;
		$scope.mcdetailspagestotal = [[.entity.MCSetup.Sims]]/10;
		$scope.mcdetailsorder = "";
		$scope.path = [[.path]];
		$scope.strategy = [[.entity.Strategy]];
		$scope.coaSelection = [];
		$scope.monthly = [[.Monthly]];
		$scope.azureURL = [[.AzureURL]];
		$scope.compute = [[.Compute]];
		$scope.mcActive = [[.MCActive]];
		$scope.debug = [[.Debug]];

		// getModels - gets the fund or asset models of a specified entity (not childEntityModels).
		$scope.getModels = async function(entity) {
			await $http.post("?entity="+entity).then(
				function successCallback(response) {
					$scope.entityModelTableResponse = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("post failed");
				}
			);
		};

		// viewEntityModel - starting from the home page, calls get method for API ViewEntity 
		$scope.viewEntityModel = async function(entityModel){
			document.getElementById("body-Home").remove();
			await $http.get("/ViewEntity?name="+entityModel).then(
				function successCallback(response) {
					$scope.viewEntityResponse = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("post failed");
				}
			);
			await $scope.getSettings(entityModel);
			await $scope.getRequest('cf');
		};

		// getRequest - used to update the table (cf, irr...etc)
		$scope.getRequest = async function(route) {
			let url = $scope.baseURL + "MCTabs?tab="
			let entity = "&name="+$scope.entity
			$scope.tableTab = route;
			if (route == 'details'){
				await $scope.post($scope.baseURL + "MCDetails?name="+$scope.entity+"&page=1");
				$scope.mcdetailspage = 1;
			} else {
				await $scope.post(url+route+entity+"&coa="+$scope.coaSelection);
			}
			await $scope.getChart();
		}; // /getRequest

		// getSettings - gets settings tab
		$scope.getSettings = async function(entity){	
			let url = $scope.baseURL + "GetSettings?entity="+entity+"&tab="+$scope.settingsTab;
			$scope.entity = entity;
			await $http.post(url).then(
				function successCallback(response) {
					$scope.settingsResponse = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getSettings failed");
				}
			);
		}; // /getSettings

		// changeEntityModel - used by the entityModel tabs above the settings table
		$scope.changeEntityModel = async function(entityModel) {
			await $scope.getSettings(entityModel);
			if ($scope.tableTab == "units"){
				await $scope.getUnitTable(-1);
				await $scope.getChart();
			}else if ($scope.tableTab == "sensitivity"){
				await $scope.getSensitivity();
				await $scope.getChart();
			}else {
				await $scope.getRequest($scope.tableTab);
			}
		}; // /changeEntityModel

		// getUnitTable - index is used for selecting the MC simulation, and getting its unit table
		$scope.getUnitTable = async function(index){
			$scope.tableTab = "units";
			await $http.post($scope.baseURL + "ViewUnitTable?name="+$scope.entity+"&index="+index).then(
				function successCallback(response) {
					// $scope.unittable = $sce.trustAsHtml(response.data);
					$scope.response = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getUnitTable failed");
				}
			);
		}; // /getUnitTable

		// updateSettingsTab
		$scope.updateSettingsTab = function(tab){
			$scope.settingsTab = tab;
		}; // /updateSettingsTab

		// get
		$scope.get = async function(url){
			await $http.get(url).then(
				function successCallback(response) {
					$scope.response = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("get failed");
				}
			);
		}; // /get

		// post
		$scope.post = async function(url){
			await $http.post(url).then(
				function successCallback(response) {
					$scope.response = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("post failed");
				}
			);
		}; // /post

		// viewCFIndex
		$scope.viewCFIndex = async function(index){
			$scope.data = '';
			let body = "?name="+$scope.entity+"&index="+index*$scope.mcdetailspage;
			await $http.post($scope.baseURL + "MCIndex"+body).then(
				function successCallback(response) {
					$scope.data = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("viewCFIndex failed");
				}
			);
		}; // /viewCFIndex
	   
		// sortMCDetails
		$scope.sortMCDetails = async function(order){
			if ($scope.mcdetailsorder == order){
				order = order+"-r";
			}
			$scope.mcdetailsorder = order;
			await $scope.post($scope.baseURL + "MCDetails?name="+$scope.entity+"&page=1"+"&order="+order);
			$scope.mcdetailspage = 1;
		}; // /sortMCDetails

		// nextMCDetails
		$scope.nextMCDetails = async function(direction){
			if (direction === '+'){
			$scope.mcdetailspage++;
			} else if (direction === '-'){
				$scope.mcdetailspage--;
			}
			await $scope.post($scope.baseURL + "MCDetails?name="+$scope.entity+"&page="+$scope.mcdetailspage);
		}; // /nextMCDetails

		// getRentSchedule
		$scope.getRentSchedule = async function(unit, index){
			let url = $scope.baseURL + "ViewRentSchedule?unit="+unit+"&name="+$scope.entity+"&index="+index
			if (typeof index == 'undefined'){
				url = $scope.baseURL + "ViewRentSchedule?unit="+unit+"&name="+$scope.entity
			}
			await $http.post(url).then(
				function successCallback(response) {
					$scope.rentschedule = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getRentSchedule failed");
				}
			);
		}; // /getRentSchedule

		// getUnitCF
		$scope.getUnitCF = async function(unit){
			await $http.post($scope.baseURL + "ViewUnitCF?unit="+unit+"&name="+$scope.entity).then(
				function successCallback(response) {
					$scope.unitcf = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getUnitCF failed");
				}
			);
		}; // /getUnitCF

		// hideNext
		$scope.hideNext = function(){
			return $scope.mcdetailspage == $scope.mcdetailspagestotal;
		} // /hideNext

		// hidePrevious
		$scope.hidePrevious = function(){
			return $scope.mcdetailspage == 1;
		} // /hidePrevious

		// init
		let init = function () {
			$scope.getSettings([[.entity.Name]]);
			console.log($scope);
		};// /init

		// updateEntity - calls the Post method for ViewEntity2 controller
		$scope.updateEntity = async function(entity){
			let a = "&";
			let e = "=";
			let paramString = ""
			paramString = paramString.concat(a,"startmonth",e,document.getElementById("startmonth").value)
			paramString = paramString.concat(a,"startyear",e,document.getElementById("startyear").value)
			paramString = paramString.concat(a,"salesmonth",e,document.getElementById("salesmonth").value)
			paramString = paramString.concat(a,"salesyear",e,document.getElementById("salesyear").value)
			paramString = paramString.concat(a,"entryyield",e,(document.getElementById("entryyield")) ? document.getElementById("entryyield").value : "0")
			paramString = paramString.concat(a,"ltv",e,(document.getElementById("settings_ltv")) ? document.getElementById("settings_ltv").value : "0")
			paramString = paramString.concat(a,"rate",e,(document.getElementById("settings_rate")) ? document.getElementById("settings_rate").value : "0")
			paramString = paramString.concat(a,"discount",e,(document.getElementById("settings_discount")) ? document.getElementById("settings_discount").value : "0")
			paramString = paramString.concat(a,"soldrent",e,(document.getElementById("settings_soldrent")) ? document.getElementById("settings_soldrent").value : "0")
			paramString = paramString.concat(a,"strategy",e,(document.getElementById("strategy")) ? document.getElementById("strategy").value : "0")
			paramString = paramString.concat(a,"balpercent",e,(document.getElementById("settings_balpercent")) ? document.getElementById("settings_balpercent").value : "0")
			paramString = paramString.concat(a,"ervshorttermrate",e,(document.getElementById("ervshorttermrate")) ? document.getElementById("ervshorttermrate").value : "0")
			paramString = paramString.concat(a,"ervshorttermratesigma",e,(document.getElementById("ervshorttermratesigma")) ? document.getElementById("ervshorttermratesigma").value : "0")
			paramString = paramString.concat(a,"ervshorttermperiod",e,(document.getElementById("ervshorttermperiod")) ? document.getElementById("ervshorttermperiod").value : "0")
			paramString = paramString.concat(a,"ervshorttermperiodsigma",e,(document.getElementById("ervshorttermperiodsigma")) ? document.getElementById("ervshorttermperiodsigma").value : "0")
			paramString = paramString.concat(a,"ervtransitionperiod",e,(document.getElementById("ervtransitionperiod")) ? document.getElementById("ervtransitionperiod").value : "0")
			paramString = paramString.concat(a,"ervtransitionperiodsigma",e,(document.getElementById("ervtransitionperiodsigma")) ? document.getElementById("ervtransitionperiodsigma").value : "0")
			paramString = paramString.concat(a,"ervlongtermrate",e,(document.getElementById("ervlongtermrate")) ? document.getElementById("ervlongtermrate").value : "0")
			paramString = paramString.concat(a,"ervlongtermratesigma",e,(document.getElementById("ervlongtermratesigma")) ? document.getElementById("ervlongtermratesigma").value : "0")
			paramString = paramString.concat(a,"cpishorttermrate",e,(document.getElementById("cpishorttermrate")) ? document.getElementById("cpishorttermrate").value : "0")
			paramString = paramString.concat(a,"cpishorttermratesigma",e,(document.getElementById("cpishorttermratesigma")) ? document.getElementById("cpishorttermratesigma").value : "0")
			paramString = paramString.concat(a,"cpishorttermperiod",e,(document.getElementById("cpishorttermperiod")) ? document.getElementById("cpishorttermperiod").value : "0")
			paramString = paramString.concat(a,"cpishorttermperiodsigma",e,(document.getElementById("cpishorttermperiodsigma")) ? document.getElementById("cpishorttermperiodsigma").value : "0")
			paramString = paramString.concat(a,"cpitransitionperiod",e,(document.getElementById("cpitransitionperiod")) ? document.getElementById("cpitransitionperiod").value : "0")
			paramString = paramString.concat(a,"cpitransitionperiodsigma",e,(document.getElementById("cpitransitionperiodsigma")) ? document.getElementById("cpitransitionperiodsigma").value : "0")
			paramString = paramString.concat(a,"cpilongtermrate",e,(document.getElementById("cpilongtermrate")) ? document.getElementById("cpilongtermrate").value : "0")
			paramString = paramString.concat(a,"cpilongtermratesigma",e,(document.getElementById("cpilongtermratesigma")) ? document.getElementById("cpilongtermratesigma").value : "0")
			paramString = paramString.concat(a,"yieldshift",e,(document.getElementById("yieldshift")) ? document.getElementById("yieldshift").value : "0")
			paramString = paramString.concat(a,"valuationmethod",e,(document.getElementById("valuationmethod")) ? document.getElementById("valuationmethod").value : "0")
			paramString = paramString.concat(a,"discountrate",e,(document.getElementById("entitydiscountrate")) ? document.getElementById("entitydiscountrate").value : "0")
			paramString = paramString.concat(a,"acqprice",e,(document.getElementById("settingsacqprice")) ? document.getElementById("settingsacqprice").value.replace(/,/g,'') : "0")
			paramString = paramString.concat(a,"void",e,(document.getElementById("void")) ? document.getElementById("void").value : "0")
			paramString = paramString.concat(a,"voidsigma",e,(document.getElementById("voidsigma")) ? document.getElementById("voidsigma").value : "0")
			paramString = paramString.concat(a,"duration",e,(document.getElementById("duration")) ? document.getElementById("duration").value : "0")
			paramString = paramString.concat(a,"rentrevision",e,(document.getElementById("rentrevision")) ? document.getElementById("rentrevision").value : "0")
			paramString = paramString.concat(a,"probability",e,(document.getElementById("probability")) ? document.getElementById("probability").value : "0")
			paramString = paramString.concat(a,"probabilitysigma",e,(document.getElementById("probabilitysigma")) ? document.getElementById("probabilitysigma").value : "0")
			paramString = paramString.concat(a,"incentivemonths",e,(document.getElementById("incentivemonths")) ? document.getElementById("incentivemonths").value : "0")
			paramString = paramString.concat(a,"incentivepercent",e,(document.getElementById("incentivepercent")) ? document.getElementById("incentivepercent").value : "0")
			paramString = paramString.concat(a,"fitoutcosts",e,(document.getElementById("fitoutcosts")) ? document.getElementById("fitoutcosts").value : "0")
			paramString = paramString.concat(a,"opex",e,(document.getElementById("settings_opex")) ? document.getElementById("settings_opex").value : "0")
			paramString = paramString.concat(a,"fees",e,(document.getElementById("settings_fees")) ? document.getElementById("settings_fees").value : "0")
			paramString = paramString.concat(a,"hazard",e,(document.getElementById("hazard")) ? document.getElementById("hazard").value : "0")
			paramString = paramString.concat(a,"rett",e,(document.getElementById("rett")) ? document.getElementById("rett").value : "0")
			paramString = paramString.concat(a,"landvalue",e,(document.getElementById("landvalue")) ? document.getElementById("landvalue").value : "0")
			paramString = paramString.concat(a,"minvalue",e,(document.getElementById("minvalue")) ? document.getElementById("minvalue").value : "0")
			paramString = paramString.concat(a,"usableperiod",e,(document.getElementById("usableperiod")) ? document.getElementById("usableperiod").value : "0")
			paramString = paramString.concat(a,"vat",e,(document.getElementById("vat")) ? document.getElementById("vat").value : "0")
			paramString = paramString.concat(a,"carrybackyrs",e,(document.getElementById("carrybackyrs")) ? document.getElementById("carrybackyrs").value : "0")
			paramString = paramString.concat(a,"carryforwardyrs",e,(document.getElementById("carryforwardyrs")) ? document.getElementById("carryforwardyrs").value : "0")
			paramString = paramString.concat(a,"sims",e,(document.getElementById("sims")) ? document.getElementById("sims").value : "0")
			let loans = $scope.getLoans();	
			let costs = $scope.getCosts();		
			await $http.post($scope.baseURL + "ViewEntity2?name="+$scope.entity+paramString+loans+costs);
			if ($scope.tableTab == "units"){
				await $scope.getUnitTable(-1);
			}else {
				await $scope.getRequest($scope.tableTab);
			}
			await $scope.getSettings($scope.entity);
		}; // /updateEntity

		$scope.getLoans = function() {
			let table = document.getElementById('loansTable');
			let loanString = "";
			let a = "&";
			let e = "=";
			for (let r = 1, n = table.rows.length; r < n; r++) {
				let loan = table.rows[r].id.replace('row_loan_','')
				loanString = loanString.concat(a,"loanAmount",loan,e,document.getElementById("loan_amount_" + loan).value.replace(/,/g,''))
				loanString = loanString.concat(a,"interestRate",loan,e,document.getElementById("interestrate_" + loan).value/100)
				loanString = loanString.concat(a,"interestType",loan,e,document.getElementById("interest_type_" + loan).value)
				loanString = loanString.concat(a,"loanType",loan,e,document.getElementById("loan_type_" + loan).value)
				loanString = loanString.concat(a,"loanName",loan,e,document.getElementById("loan_name_" + loan).value)
				loanString = loanString.concat(a,"loanStartMonth",loan,e,document.getElementById("loan_start_month_" + loan).value)
				loanString = loanString.concat(a,"loanStartYear",loan,e,document.getElementById("loan_start_year_" + loan).value)
				loanString = loanString.concat(a,"loanEndMonth",loan,e,document.getElementById("loan_end_month_" + loan).value)
				loanString = loanString.concat(a,"loanEndYear",loan,e,document.getElementById("loan_end_year_" + loan).value)
				loanString = loanString.concat(a,"floatBasis",loan,e,document.getElementById("float_basis_" + loan).value)
				loanString = loanString.concat(a,"spread",loan,e,document.getElementById("spread_" + loan).value)
				loanString = loanString.concat(a,"amortizationPeriod",loan,e,document.getElementById("amortization_period_" + loan).value)
				loanString = loanString.concat(a,"active",loan,e,document.getElementById("active_" + loan).checked)
			}
			return loanString
		};

		$scope.getCosts = function() {
			let table = document.getElementsByName('costInputRow');
			let costString = "";
			let a = "&";
			let e = "=";
			for (let r = 0, n = table.length; r < n; r++){
				let rowID = table[r].id
				costString = costString.concat(a,`${rowID}_masterID=`,document.getElementById(`${rowID}_masterID`).value)
				costString = costString.concat(a,`${rowID}_type=`,document.getElementById(`${rowID}_type`).value)
				costString = costString.concat(a,`${rowID}_name=`,document.getElementById(`${rowID}_name`).value) 
				costString = costString.concat(a,`${rowID}_amount=`,document.getElementById(`${rowID}_amount`).value)
				costString = costString.concat(a,`${rowID}_coaItemBasis=`,document.getElementById(`${rowID}_coaItemBasis`).value.trim())
				costString = costString.concat(a,`${rowID}_coaItemTarget=`,document.getElementById(`${rowID}_coaItemTarget`).value.trim())
				costString = costString.concat(a,`${rowID}_growthItem=`,document.getElementById(`${rowID}_growthItem`).value.trim())
				costString = costString.concat(a,`${rowID}_duration=`,document.getElementById(`${rowID}_duration`).value)
				costString = costString.concat(a,`${rowID}_startMonth=`,document.getElementById(`${rowID}_startMonth`).value)
				costString = costString.concat(a,`${rowID}_startYear=`,document.getElementById(`${rowID}_startYear`).value)
				costString = costString.concat(a,`${rowID}_endMonth=`,document.getElementById(`${rowID}_endMonth`).value)
				costString = costString.concat(a,`${rowID}_endYear=`,document.getElementById(`${rowID}_endYear`).value) 
			}
			return costString
		};

		$scope.addCost = async function() {
			await $http.post(`/AddCost?entityModel=${$scope.entity}`);
			await $scope.getSettings($scope.entity);			
		};

		$scope.updateUnit = async function(unit,field){
			let value = document.getElementById(`unit${unit}.${field}`).value;
			value = value.replace(/,/g,'')
			await $http.post(`UpdateUnit?unit=${unit}&field=${field}&value="${value}"`);
			await $scope.getUnitTable(-1);
			await $scope.getSettings($scope.entity);	
			await $scope.getChart();	
		};

		$scope.addUnit = async function(){
			let params = "AddChildUnit?";
			let a = "&";
			let e = "=";
			params = params.concat("parent=",document.getElementById("new_unit_parent").value)
			params = params.concat(a,"tenant",e,document.getElementById("new_unit_tenant").value)
			params = params.concat(a,"unitname",e,document.getElementById("new_unit_unitname").value)
			params = params.concat(a,"rent",e,document.getElementById("new_unit_rent").value)
			params = params.concat(a,"startmonth",e,document.getElementById("new_unit_startmonth").value)
			params = params.concat(a,"startyear",e,document.getElementById("new_unit_startyear").value)
			params = params.concat(a,"expirymonth",e,document.getElementById("new_unit_expirymonth").value)
			params = params.concat(a,"expiryyear",e,document.getElementById("new_unit_expiryyear").value)
			params = params.concat(a,"amount",e,document.getElementById("new_unit_amount").value)
			params = params.concat(a,"area",e,document.getElementById("new_unit_area").value)
            document.querySelector('.jw-modal.open').classList.remove('open');
            document.body.classList.remove('jw-modal-open');
			await $http.post(params);
			await $scope.getUnitTable(-1);
			await $scope.getSettings($scope.entity);
		};

		$scope.selectForChart = async function(coa) {
			if ($scope.coaSelection.includes(coa)){
				$scope.coaSelection = $scope.coaSelection.filter(e => e !== coa);
				document.getElementById(`${coa} row`).classList.remove("fifth-bg");
			} else {
				$scope.coaSelection.push(coa)
				document.getElementById(`${coa} row`).classList.add("fifth-bg");
			}
			await $scope.getChart();
		}

		$scope.getChart = async function(){
			await $http.post("/Chart?coa="+$scope.coaSelection+"&entityModel="+$scope.entity).then(
				function successCallback(response) {
					$scope.settingsChart = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("selectForChart.post failed");
				}
			);
		}

		$scope.getSensitivity = async function(){
			$scope.tableTab = "sensitivity";
			await $http.get("/Sensitivity?name="+$scope.entity).then(
				function successCallback(response) {
					$scope.response = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getSensitivity failed");
				}
			);
		}

		$scope.postSensitivity = async function(){
			let params = "Sensitivity?calc=true&entityModel="+$scope.entity;
			let a = "&";
			let e = "=";
			let vertical = ""
			let horizontal = ""
			let iterations = document.getElementById("iterations").value
			if (iterations != 0) {
				params = params.concat(a,"iterations",e,iterations)
				let table = document.getElementById("sensitivityInputTable");
				for (let i = 1; i <= table.rows.length - 1; i ++){
					let v = document.getElementById(`sensRowVertical_${i}`);
					if (v.checked){
						vertical = v.name
					}
					let h = document.getElementById(`sensRowHorizontal_${i}`);
					if (h.checked){
						horizontal = h.name
					}
					let input = document.getElementById(`sensRow${i}`)
					console.log(input.name, input.value);
					params = params.concat(a,input.name, e, input.value)
				}
				params = params.concat(a,"vertical",e,vertical)
				params = params.concat(a,"horizontal",e,horizontal)
				await $http.post(params).then(
					function successCallback(response) {
						$scope.response = $sce.trustAsHtml(response.data);
					},
					function errorCallback(response) {
						console.log("getSensitivity failed");
					}
				);
			}
		}

		$scope.updateChartCOASelections = function(){
			for (coa in $scope.coaSelection){
				document.getElementById(`${coa} row`).classList.add("fifth-bg");
			}
		}

		init();

		$scope.updateAppSettings = async function(){
			await $http.post(`AppSettings?monthly=${$scope.monthly}&compute=${$scope.compute}&azureurl="${$scope.azureURL}"&mcActive=${$scope.mcActive}&debug=${$scope.debug}`);
		}

	}]); // /testController

	//bindHtmlCompile directive 
	ognApp.directive('bindHtmlCompile', ['$compile', function ($compile) {
		return {
			restrict: 'A',
			link: function (scope, element, attrs) {
				scope.$watch(function () {
					return scope.$eval(attrs.bindHtmlCompile);
				}, function (value) {
					if (!!value) {
						let compileScope = scope;
						if (attrs.bindHtmlScope) {
							compileScope = scope.$eval(attrs.bindHtmlScope);
						}
						let html = $compile(value.toString())(compileScope);
						element.html(html);
					}
				});
			}
		};
	}]); // /bindHtmlCompile directive

</script>
</html>
[[end]]

