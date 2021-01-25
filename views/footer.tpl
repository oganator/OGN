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
    <script src="static/js/app.js"></script>
    <script src="static/js/misc.js"></script>
    <script>
    
    var ognApp = angular.module('ognApp', []);

    // testController
    ognApp.controller('testController', ['$scope', '$http', '$sce', '$compile', function($scope, $http, $sce, $compile) {
        $scope.entity = [[.entity.Name]];

        // getRequest
        $scope.getRequest = function(route) {
            var cftable = angular.element( document.querySelector( '#cftable' ) );
            cftable.remove();
            var url = ""
            if (route == 'cf'){
                url = "http://localhost:8080/CFTable"
            }
            if (route == 'endcash'){
                url = "http://localhost:8080/MCEndingCash"
            }
            if (route == 'endncf'){
                url = "http://localhost:8080/MCEndingNCF"
            }
            if (route == 'irr'){
                url = "http://localhost:8080/MCIRR"
            }
            if (route == 'em'){
                url = "http://localhost:8080/MCEM"
            }
            if (route == 'ytm'){
                url = "http://localhost:8080/MCYTM"
            } 
            if (route == 'duration'){
                url = "http://localhost:8080/MCDuration"
            } 
            if (route == 'details'){
                url = "http://localhost:8080/MCDetails"
            } 
            $http.get(url).then(
                function successCallback(response) {
                    $scope.response = $sce.trustAsHtml(response.data);
                },
                function errorCallback(response) {
                    console.log("Unable to perform get request");
                }
            );
        };//getRequest

        // viewCFIndex
        $scope.viewCFIndex = function(index){
            debugger;
            var body = "?name="+$scope.entity+"&index="+index
            console.log($scope.entity, index);
            $http.post("http://localhost:8080/MCIndex"+body).then(
                function successCallback(response) {
                    $scope.data = $sce.trustAsHtml(response.data);
                },
                function errorCallback(response) {
                    console.log("POST-ing of data failed");
                }
            );
        };// viewCFIndex
       
    }]);

    //directive 
    ognApp.directive('bindHtmlCompile', ['$compile', '$sce', function ($compile, $sce) {
        return {
            restrict: 'A',
            link: function (scope, element, attrs) {
                scope.$watch(function () {
                    return scope.$eval(attrs.bindHtmlCompile);
                }, function (value) {
                    debugger;
                    var compileScope = scope;
                    if (attrs.bindHtmlScope) {
                        compileScope = scope.$eval(attrs.bindHtmlScope);
                    }
                    var html = $compile($sce.trustAsHtml(value))(compileScope);
                    element.html(html);
                });
            }
        };
    }]); //bindHtmlCompile

    </script>

</html>
[[end]]