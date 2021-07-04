"use strict";
//call to display reserved list
var url = "/reserved/list";
var app = angular.module('myAppRVS', []);
app.controller('myCtrlRVS', function ($scope, $http) {
    $http.get(url)
        .then(function (response) {
        console.log(response.data);
        $scope.reserved = response.data;
    });
});
