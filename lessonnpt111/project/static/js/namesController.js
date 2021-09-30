var app = angular.module('myApp', ['ngAnimate']); //Needed for the animation section

//Set a custom delimiter for templates
app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});



//Used for lists and objects
app.controller('myCtrl', function($scope, $http) {
  /*
  $http.get("/testGetTheSecond").then(function(response){
    $scope.myWelcome = response.data;
    $scope.jsonString = String(response.data);
    $scope.DataGotten = JSON.parse($scope.jsonString);
    console.log("The succ or fail is: " + $scope.DataGotten.SuccOrFail);
    for (var i=0; i < $scope.DataGotten.AllData.ThePerson.length; i++){
      console.log("Here is the person name here: " + String($scope.DataGotten.AllData.ThePersons.Name));
    }
    console.log("Here is the special string: " + $scope.DataGotten.AllData.SpecialString);
  });
  */

  //Second http example
  $scope.hasCompleted = false; // Do not show data until http gets back with data
  $scope.ThePersons = [];
  $scope.PersonNames = [];
  $http({
    method: 'GET',
    url: '/testGetTheSecond'
  }).then(function successCallback(response) {
    // this callback will be called asynchronously
    // when the response is available
    console.log(response.data);
    console.log(response.data.ResultMsg);
    for (var i = 0; i < response.data.AllData.ThePerson.length; i++){
      console.log("This is person: " + response.data.AllData.ThePerson[i].Name);
      $scope.ThePersons.push(response.data.AllData.ThePerson[i]);
      $scope.PersonNames.push(response.data.AllData.ThePerson[i].Name);
    }

    $scope.hasCompleted = true; //Data load complete, we can show data in template
  }, function errorCallback(response) {
    // called asynchronously if an error occurs
    // or server returns response with an error status.
    console.log("You got a fucking error! " + String(response));
  });

});