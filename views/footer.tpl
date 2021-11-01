[[define "footer"]]
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
		$scope.entity = [[.entity.Name]];
		$scope.settingsTab = "settings";
		$scope.tableTab = "cf";
		$scope.mcdetailspage = 1;
		$scope.mcdetailspagestotal = [[.entity.MCSetup.Sims]]/10;
		$scope.mcdetailsorder = "";
		$scope.strategy = [[.entity.Strategy]];

		// getRequest - used to update the table (cf, irr...etc)
		$scope.getRequest = function(route) {
			var url = "http://localhost:8080/MCTabs?tab="
			var entity = "&name="+$scope.entity
			$scope.tableTab = route;
			if (route == 'details'){
				$scope.post("http://localhost:8080/MCDetails?name="+$scope.entity+"&page=1");
				$scope.mcdetailspage = 1;
			} else {
				$scope.post(url+route+entity);
			}
		}; // /getRequest

		// getSettings
		$scope.getSettings = function(entity){
			var url = "http://localhost:8080/GetSettings?entity="+entity
			$scope.entity = entity;
			
			$http.post(url).then(
				function successCallback(response) {
					$scope.settingsResponse = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getSettings failed");
				}
			);
			$scope.getRequest($scope.tableTab);
			$scope.getUnitTable(-1);
		}; // /getSettings

		// updateSettingsTab
		$scope.updateSettingsTab = function(tab){
			$scope.settingsTab = tab;
		};

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
			//$scope.getUnitTable(index);
			$scope.data = '';
			var body = "?name="+$scope.entity+"&index="+index*$scope.mcdetailspage;
			$http.post("http://localhost:8080/MCIndex"+body).then(
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
			$scope.post("http://localhost:8080/MCDetails?name="+$scope.entity+"&page=1"+"&order="+order);
			$scope.mcdetailspage = 1;
		}; // /sortMCDetails

		// nextMCDetails
		$scope.nextMCDetails = function(direction){
			if (direction === '+'){
			$scope.mcdetailspage++;
			} else if (direction === '-'){
				$scope.mcdetailspage--;
			}
			$scope.post("http://localhost:8080/MCDetails?name="+$scope.entity+"&page="+$scope.mcdetailspage);
		}; // /nextMCDetails

		// getRentSchedule
		$scope.getRentSchedule = function(unit, index){
			var url = "http://localhost:8080/ViewRentSchedule?unit="+unit+"&name="+$scope.entity+"&index="+index
			if (typeof index == 'undefined'){
				url = "http://localhost:8080/ViewRentSchedule?unit="+unit+"&name="+$scope.entity
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
			$http.post("http://localhost:8080/ViewUnitCF?unit="+unit+"&name="+$scope.entity).then(
				function successCallback(response) {
					$scope.unitcf = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getUnitCF failed");
				}
			);
		}; // /getUnitCF

		// getUnitTable
		$scope.getUnitTable = function(index){
			$http.post("http://localhost:8080/ViewUnitTable?name="+$scope.entity+"&index="+index).then(
				function successCallback(response) {
					$scope.unittable = $sce.trustAsHtml(response.data);
				},
				function errorCallback(response) {
					console.log("getUnitTable failed");
				}
			);
		}; // /getUnitTable

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