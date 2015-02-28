

var jaegerApp = angular.module('jaegerApp', []);

jaegerApp.controller('PhoneListCtrl', function ($scope) {
  $scope.phones = [
    {'name': 'Nexus S',
     'snippet': 'Fast just got faster with Nexus S.'},
    {'name': 'Motorola XOOM™ with Wi-Fi',
     'snippet': 'The Next, Next Generation tablet.'},
    {'name': 'MOTOROLA XOOM™',
     'snippet': 'The Next, Next Generation tablet.'}
  ];
});

jaegerApp.controller('itemController', function ($scope) {

	$scope.upload = function() {alert("aaaa");};
	$scope.delete = function() {alert("delete item?");};
  /*
	$scope.slideItemMessage = "shdfjhsdfjk hsjkfh sdjkfhsdjkfh jksdfh jksdhfjk sdhf jksdhf jksdhf jksdhf jksdh fjksdhfjk fasdfsdf";
  */
});


jaegerApp.controller('CarouselCtrlr', function($scope) {
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
  
  $scope.getRandomNum = function(){
    return Math.floor((Math.random()*9999999)+1);
  } 
 
  $scope.upload = function(event) {
    debugger
    alert(event.target.id);
  };
  $scope.delete = function(event) {
    debugger
    alert(event.target.id);
  };
  $scope.add = function(event) {
    //debugger
    //alert(event.target.id);
    var newitm = new function() {
            this.heading = "New Heading";
            this.description = "Description";
            this.id = Math.floor((Math.random()*9999999)+1);
            this.image = "http://placehold.it/350x250";
        }    
    $scope.list.items.push(newitm);
    
  };
  
});

jaegerApp.directive('contenteditable', function () {
      return {
          restrict: 'A', // only activate on element attribute
          require: '?ngModel', // get a hold of NgModelController
          link: function (scope, element, attrs, ngModel) {
              if (!ngModel) return; // do nothing if no ng-model

              // Specify how UI should be updated
              ngModel.$render = function () {
                  element.html(ngModel.$viewValue || '');
              };

              // Listen for change events to enable binding
              element.on('blur keyup change', function () {
                  scope.$apply(readViewText);
              });

              // No need to initialize, AngularJS will initialize the text based on ng-model attribute

              // Write data to the model
              function readViewText() {
                  var html = element.html();
                  // When we clear the content editable the browser leaves a <br> behind
                  // If strip-br attribute is provided then we strip this out
                  if (attrs.stripBr && html == '<br>') {
                      html = '';
                  }
                  ngModel.$setViewValue(html);
              }
          }
      };
  });