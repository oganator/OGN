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
var ognApp = angular.module('ognApp', []);
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

		// getModels - gets the fund or asset models of a specified entity (not childEntityModels).
		$scope.getModels = function(entity) {
			$http.post("?entity="+entity).then(
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
		$scope.getRequest = function(route) {
			var url = $scope.baseURL + "MCTabs?tab="
			var entity = "&name="+$scope.entity
			$scope.tableTab = route;
			if (route == 'details'){
				$scope.post($scope.baseURL + "MCDetails?name="+$scope.entity+"&page=1");
				$scope.mcdetailspage = 1;
			} else {
				$scope.post(url+route+entity);
			}
		}; // /getRequest

		// getSettings - gets settings tab
		$scope.getSettings = async function(entity){	
			var url = $scope.baseURL + "GetSettings?entity="+entity+"&tab="+$scope.settingsTab;
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
			}else {
				await $scope.getRequest($scope.tableTab);
			}
		}; // /changeEntityModel

		// getUnitTable - index is used for selecting the MC simulation, and getting its unit table
		$scope.getUnitTable = function(index){
			$scope.tableTab = "units";
			$http.post($scope.baseURL + "ViewUnitTable?name="+$scope.entity+"&index="+index).then(
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
		$scope.get = function(url){
			$http.get(url).then(
				function successCallback(response) {
					$scope.response = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("get failed");
				}
			);
		}; // /get

		// post
		$scope.post = function(url){
			$http.post(url).then(
				function successCallback(response) {
					$scope.response = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("post failed");
				}
			);
		}; // /post

		// viewCFIndex
		$scope.viewCFIndex = function(index){
			$scope.data = '';
			var body = "?name="+$scope.entity+"&index="+index*$scope.mcdetailspage;
			$http.post($scope.baseURL + "MCIndex"+body).then(
				function successCallback(response) {
					$scope.data = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("viewCFIndex failed");
				}
			);
		}; // /viewCFIndex
	   
		// sortMCDetails
		$scope.sortMCDetails = function(order){
			if ($scope.mcdetailsorder == order){
				order = order+"-r";
			}
			$scope.mcdetailsorder = order;
			$scope.post($scope.baseURL + "MCDetails?name="+$scope.entity+"&page=1"+"&order="+order);
			$scope.mcdetailspage = 1;
		}; // /sortMCDetails

		// nextMCDetails
		$scope.nextMCDetails = function(direction){
			if (direction === '+'){
			$scope.mcdetailspage++;
			} else if (direction === '-'){
				$scope.mcdetailspage--;
			}
			$scope.post($scope.baseURL + "MCDetails?name="+$scope.entity+"&page="+$scope.mcdetailspage);
		}; // /nextMCDetails

		// getRentSchedule
		$scope.getRentSchedule = function(unit, index){
			var url = $scope.baseURL + "ViewRentSchedule?unit="+unit+"&name="+$scope.entity+"&index="+index
			if (typeof index == 'undefined'){
				url = $scope.baseURL + "ViewRentSchedule?unit="+unit+"&name="+$scope.entity
			}
			$http.post(url).then(
				function successCallback(response) {
					$scope.rentschedule = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getRentSchedule failed");
				}
			);
		}; // /getRentSchedule

		// getUnitCF
		$scope.getUnitCF = function(unit){
			$http.post($scope.baseURL + "ViewUnitCF?unit="+unit+"&name="+$scope.entity).then(
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
		var init = function () {
			$scope.getSettings([[.entity.Name]]);
		};// /init

		// updateEntity - calls the Post method for ViewEntity2 controller
		$scope.updateEntity = async function(entity){
			var startmonth = document.getElementById("startmonth").value;
			var startyear = document.getElementById("startyear").value;
			var salesmonth = document.getElementById("salesmonth").value;
			var salesyear = document.getElementById("salesyear").value;
			var entryyield = (document.getElementById("entryyield")) ? document.getElementById("entryyield").value : "0";
			var ltv = (document.getElementById("settings_ltv")) ? document.getElementById("settings_ltv").value : "0";
			var rate = (document.getElementById("settings_rate")) ? document.getElementById("settings_rate").value : "0";
			var discount = (document.getElementById("settings_discount")) ? document.getElementById("settings_discount").value : "0";
			var soldrent = (document.getElementById("settings_soldrent")) ? document.getElementById("settings_soldrent").value : "0";
			var strategy = (document.getElementById("strategy")) ? document.getElementById("strategy").value : "0";
			var balpercent = (document.getElementById("settings_balpercent")) ? document.getElementById("settings_balpercent").value : "0";
			var ervshorttermrate = (document.getElementById("ervshorttermrate")) ? document.getElementById("ervshorttermrate").value : "0";
			var ervshorttermratesigma = (document.getElementById("ervshorttermratesigma")) ? document.getElementById("ervshorttermratesigma").value : "0";
			var ervshorttermperiod = (document.getElementById("ervshorttermperiod")) ? document.getElementById("ervshorttermperiod").value : "0";
			var ervshorttermperiodsigma = (document.getElementById("ervshorttermperiodsigma")) ? document.getElementById("ervshorttermperiodsigma").value : "0";
			var ervtransitionperiod = (document.getElementById("ervtransitionperiod")) ? document.getElementById("ervtransitionperiod").value : "0";
			var ervtransitionperiodsigma = (document.getElementById("ervtransitionperiodsigma")) ? document.getElementById("ervtransitionperiodsigma").value : "0";
			var ervlongtermrate = (document.getElementById("ervlongtermrate")) ? document.getElementById("ervlongtermrate").value : "0";
			var ervlongtermratesigma = (document.getElementById("ervlongtermratesigma")) ? document.getElementById("ervlongtermratesigma").value : "0";
			var cpishorttermrate = (document.getElementById("cpishorttermrate")) ? document.getElementById("cpishorttermrate").value : "0";
			var cpishorttermratesigma = (document.getElementById("cpishorttermratesigma")) ? document.getElementById("cpishorttermratesigma").value : "0";
			var cpishorttermperiod = (document.getElementById("cpishorttermperiod")) ? document.getElementById("cpishorttermperiod").value : "0";
			var cpishorttermperiodsigma = (document.getElementById("cpishorttermperiodsigma")) ? document.getElementById("cpishorttermperiodsigma").value : "0";
			var cpitransitionperiod = (document.getElementById("cpitransitionperiod")) ? document.getElementById("cpitransitionperiod").value : "0";
			var cpilongtermrate = (document.getElementById("cpilongtermrate")) ? document.getElementById("cpilongtermrate").value : "0";
			var yieldshift = (document.getElementById("yieldshift")) ? document.getElementById("yieldshift").value : "0";
			var valuationmethod = (document.getElementById("valuationmethod")) ? document.getElementById("valuationmethod").value : "0";
			var discountrate = (document.getElementById("entitydiscountrate")) ? document.getElementById("entitydiscountrate").value : "0";
			var void2 = (document.getElementById("void")) ? document.getElementById("void").value : "0";
			var duration = (document.getElementById("duration")) ? document.getElementById("duration").value : "0";
			var rentrevision = (document.getElementById("rentrevision")) ? document.getElementById("rentrevision").value : "0";
			var probability = (document.getElementById("probability")) ? document.getElementById("probability").value : "0";
			var incentivemonths = (document.getElementById("incentivemonths")) ? document.getElementById("incentivemonths").value : "0";
			var incentivepercent = (document.getElementById("incentivepercent")) ? document.getElementById("incentivepercent").value : "0";
			var fitoutcosts = (document.getElementById("fitoutcosts")) ? document.getElementById("fitoutcosts").value : "0";
			var opex = (document.getElementById("settings_opex")) ? document.getElementById("settings_opex").value : "0";
			var fees = (document.getElementById("settings_fees")) ? document.getElementById("settings_fees").value : "0";
			var hazard = (document.getElementById("hazard")) ? document.getElementById("hazard").value : "0";
			var rett = (document.getElementById("rett")) ? document.getElementById("rett").value : "0";
			var landvalue = (document.getElementById("landvalue")) ? document.getElementById("landvalue").value : "0";
			var minvalue = (document.getElementById("minvalue")) ? document.getElementById("minvalue").value : "0";
			var usableperiod = (document.getElementById("usableperiod")) ? document.getElementById("usableperiod").value : "0";
			var vat = (document.getElementById("vat")) ? document.getElementById("vat").value : "0";
			var carrybackyrs = (document.getElementById("carrybackyrs")) ? document.getElementById("carrybackyrs").value : "0";
			var carryforwardyrs = (document.getElementById("carryforwardyrs")) ? document.getElementById("carryforwardyrs").value : "0";
			var sims = (document.getElementById("sims")) ? document.getElementById("sims").value : "0";
			var a = "&";
			var e = "=";
			var params = a+"startmonth"+e+startmonth+a+"startyear"+e+startyear+a+"salesmonth"+e+salesmonth+a+"salesyear"+e+salesyear+a+"entryyield"+e+entryyield+a+"ltv"+e+ltv+a+"rate"+e+rate+a+"discount"+e+discount+a+"soldrent"+e+soldrent+a+"strategy"+e+strategy+a+"balpercent"+e+balpercent+a+"ervshorttermrate"+e+ervshorttermrate+a+"ervshorttermperiod"+e+ervshorttermperiod+a+"ervtransitionperiod"+e+ervtransitionperiod+a+"ervlongtermrate"+e+ervlongtermrate+a+"cpishorttermrate"+e+cpishorttermrate+a+"cpishorttermperiod"+e+cpishorttermperiod+a+"cpitransitionperiod"+e+cpitransitionperiod+a+"cpilongtermrate"+e+cpilongtermrate+a+"yieldshift"+e+yieldshift+a+"discountrate"+e+discountrate+a+"valuationmethod"+e+valuationmethod+a+"void"+e+void2+a+"duration"+e+duration+a+"rentrevision"+e+rentrevision+a+"probability"+e+probability+a+"incentivemonths"+e+incentivemonths+a+"incentivepercent"+e+incentivepercent+a+"fitoutcosts"+e+fitoutcosts+a+"opex"+e+opex+a+"fees"+e+fees+a+"hazard"+e+hazard+a+"rett"+e+rett+a+"landvalue"+e+landvalue+a+"minvalue"+e+minvalue+a+"usableperiod"+e+usableperiod+a+"vat"+e+vat+a+"carrybackyrs"+e+carrybackyrs+a+"carryforwardyrs"+e+carryforwardyrs+a+"sims"+e+sims
			await $http.post($scope.baseURL + "ViewEntity2?name="+$scope.entity+params);
			if ($scope.tableTab == "units"){
				await $scope.getUnitTable(-1);
			}else {
				await $scope.getRequest($scope.tableTab);
			}
			await $scope.getSettings($scope.entity);
		
		}; // /updateEntity

		$scope.updateUnit = async function(unit,field){
			var value = document.getElementById(`unit${unit}.${field}`).value;
			value = value.replace(',','')
			await $http.post(`UpdateUnit?unit=${unit}&field=${field}&value="${value}"`);
			await $scope.getUnitTable(-1);
			await $scope.getSettings($scope.entity);
		
		};

		init();

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
						var compileScope = scope;
						if (attrs.bindHtmlScope) {
							compileScope = scope.$eval(attrs.bindHtmlScope);
						}
						var html = $compile(value.toString())(compileScope);
						element.html(html);
					}
				});
			}
		};
	}]); // /bindHtmlCompile directive

</script>
</html>
[[end]]