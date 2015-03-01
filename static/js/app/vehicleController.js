

jaegerApp.controller('VehicleCtrlr', ['$scope','$http', function($scope,$http) {

  var json = {
    "items":
        [
                {
                   "id":"11111111111",
                   "heading":"heading1",
                   "description":"description of heading1",
                   "image":"http://placehold.it/350x250"
                },
                {
                   "id":"222222222222",
                   "heading":"heading2",
                   "description":"description of heading2",
                   "image":"http://placehold.it/350x250"
                },
                {
                   "id":"3333333333333",
                   "heading":"heading3",
                   "description":"description of heading3",
                   "image":"http://placehold.it/350x250"
                }

        ]
  };
  
  $scope.list = json;

  $http.get('/ws/user/bgonza/vehicles/getall').
    success(function(data, status, headers, config) {
      // this callback will be called asynchronously
      // when the response is available
      debugger;
      $scope.list = data;
    }).
    error(function(data, status, headers, config) {
      // called asynchronously if an error occurs
      // or server returns response with an error status.
    });   

  $scope.add = function(event) {
  	alert("Add Vehicle");
  };	

}]);