

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



jaegerApp.controller('CarouselCtrlr', ['$scope','$routeParams','$http', function($scope,$routeParams,$http) {
  var json = [

        ]  ;
  
  $scope.list = json;

  $scope.model = {
      name: "",
      comments: "",
      id: ""
  };

  $scope.files = [];

  $scope.uploadurl;

  debugger;


/*
  $scope.getCarousel = function () {
    alert($scope.vehiclekey);
  }



  //$timeout($scope.init);

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

      if ($scope.files.length < 1) {
        alert("there's nothing to upload.");
        return;
      }

      $http.get('/upload/url').
        success(function(data, status, headers, config) {
          // this callback will be called asynchronously
          // when the response is available
          debugger;
          $scope.requesturi = data.uploadurl;
        }).
        error(function(data, status, headers, config) {
          // called asynchronously if an error occurs
          // or server returns response with an error status.
        });          

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
          //alert("success!" + data.blobKey);
      }).
      error(function (data, status, headers, config) {
          debugger;
          $("#fileup").type = "text";
          $("#fileup").type = "file";
          $scope.files = [];
          //alert("failed!");
      });
  };  

*/

  $scope.refresh = function(item) {
    //alert("refresh");
    debugger;
    $http.get('/ws/vehicle/' + $scope.vehiclekey + '/carousel/getall').
      success(function(data, status, headers, config) {
        // this callback will be called asynchronously
        // when the response is available
        debugger;
        if ( data == "null") {
          $scope.list = [];
        } else {
          $scope.list = data;
        }
      }).
      error(function(data, status, headers, config) {
        // called asynchronously if an error occurs
        // or server returns response with an error status.
        debugger;
        $scope.list = [];
      });     
  };

  $scope.add = function(event) {
    $http.get('/ws/vehicle/' + $scope.vehiclekey + '/carousel/allocate').
      success(function(data, status, headers, config) {
        // this callback will be called asynchronously
        // when the response is available
        //alert("allocated");
        $scope.refresh();
      }).
      error(function(data, status, headers, config) {
        // called asynchronously if an error occurs
        // or server returns response with an error status.
        //alert("allocate error");
      });     
      debugger;
  };

  $scope.delete = function(item) {
    debugger
    var index=$scope.list.indexOf(item);
    var item = $scope.list[index];
    //$scope.list.items.splice(index,1);

    $http.get('/ws/vehicle/carousel/' + item.KeyId + '/delete').
      success(function(data, status, headers, config) {
        // this callback will be called asynchronously
        // when the response is available
        //alert("allocated");
        $scope.refresh();
      }).
      error(function(data, status, headers, config) {
        // called asynchronously if an error occurs
        // or server returns response with an error status.
        //alert("allocate error");
      });     
  };
/*

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
  
*/

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

/*
jaegerApp.directive('getCarousel', function () {
    return {
        scope: { method:'&myAction' },
        link: function (scope, el, attrs) {
          debugger;
          var expressionHandler = scope.method();
          var id = attrs.value;
          //var vk = {{vehiclekey}};
          //alert("get car:" + attrs.value);
          //scope.getCarousel();
          $(element).click(function( e, rowid ) {
            expressionHandler(id);
          });
        }
    };
});

*/
/*
http://weblogs.asp.net/dwahlin/creating-custom-angularjs-directives-part-2-isolate-scope
*/
