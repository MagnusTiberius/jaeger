

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


jaegerApp.controller('CarouselCtrlr', ['$scope','$http', function($scope,$http) {
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

  $scope.model = {
      name: "",
      comments: "",
      id: ""
  };

  $scope.files = [];

  $scope.uploadurl;

  //listen for the file selected event
  $scope.$on("fileSelected", function (event, args) {
      $scope.$apply(function () {            
          //add the file object to the scope's files collection
          $scope.files.push(args.file);
      });
  });  
  
  $scope.getRandomNum = function(){
    return Math.floor((Math.random()*9999999)+1);
  } 
 
  $scope.upload2 = function(event) {
    debugger
    alert(event.target.id);
  };

  //the save method
  $scope.upload = function(event,obj) {
      //debugger;
      //alert($scope.requesturi);
      $http({
          method: 'POST',
          url: $scope.requesturi,
          //IMPORTANT!!! You might think this should be set to 'multipart/form-data' 
          // but this is not true because when we are sending up files the request 
          // needs to include a 'boundary' parameter which identifies the boundary 
          // name between parts in this multi-part request and setting the Content-type 
          // manually will not set this boundary parameter. For whatever reason, 
          // setting the Content-type to 'false' will force the request to automatically
          // populate the headers properly including the boundary parameter.
          headers: {    'Content-Type': undefined },
          //This method will allow us to change how the data is sent up to the server
          // for which we'll need to encapsulate the model data in 'FormData'
          transformRequest: function (data) {
              //debugger;
              var formData = new FormData();
              //need to convert our json object to a string version of json otherwise
              // the browser will do a 'toString()' on the object which will result 
              // in the value '[Object object]' on the server.
              data.model.id = obj.item.id;
              formData.append("model", angular.toJson(data.model));
              formData.append("obj", obj);
              //now add all of the assigned files
              for (var i = 0; i < data.files.length; i++) {
                  //add each file to the form data and iteratively name them
                  //debugger;
                  //formData.append("file" + i, data.files[i]);
                  formData.append("file" , data.files[i]);
              }
              return formData;
          },
          //Create an object that contains the model and files which will be transformed
          // in the above transformRequest method
          data: { model: $scope.model, files: $scope.files }
      }).
      success(function (data, status, headers, config) {
          //debugger;
          //$('#'+imgid).src = "/blob/blobKey=" + data.blobKey;
          //imageid = "#"+config.data.model.id + "image";
          //imgObject = $(imageid);
          //imgObject.attr('src','http://localhost:8080/images/1/myImage.png' );
          for (var i=0; i<$scope.list.items.length; i++) {
              if ($scope.list.items[i].id == config.data.model.id) {
                $scope.list.items[i].image = '/blob/?blobKey=' + data.blobKey;
              }
          }
          $("#fileup").type = "text";
          $("#fileup").type = "file";
          $scope.files = [];
          alert("success!" + data.blobKey);
      }).
      error(function (data, status, headers, config) {
          debugger;
          $("#fileup").type = "text";
          $("#fileup").type = "file";
          $scope.files = [];
          alert("failed!");
      });
  };  


  $scope.delete = function(item) {
    debugger
    var index=$scope.list.items.indexOf(item);
    $scope.list.items.splice(index,1);
  };
  $scope.add = function(event) {
    var newitm = new function() {
            this.description = "Description";
            this.id = Math.floor((Math.random()*9999999)+1);
            this.heading = "New Heading " + this.id;
            this.image = "http://placehold.it/350x250";
        }    
    $scope.list.items.push(newitm);
  };
  
}]);

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


jaegerApp.directive('fileUpload', function () {
    return {
        scope: true,        //create a new scope
        link: function (scope, el, attrs) {
            el.bind('change', function (event) {
                var files = event.target.files;
                //iterate files since 'multiple' may be specified on the element
                for (var i = 0;i<files.length;i++) {
                    //emit event upward
                    scope.$emit("fileSelected", { file: files[i] });
                }                                       
            });
        }
    };
});