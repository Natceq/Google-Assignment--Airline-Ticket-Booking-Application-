//call to display booking list
var url = "/booked";
var app = angular.module('myApp', []);
app.controller('myCtrl', function($scope, $http) {
  $http.get(url)
  .then(function(response: { data: any; }) { 
    console.log(response.data);
    $scope.flights = response.data;
  });

  url = "/ticketlist"
  $http.get(url)
  .then(function(response: { data: any; }) { 
    console.log(response.data);
    $scope.tickets = response.data;
  });
});