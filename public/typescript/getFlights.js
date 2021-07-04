"use strict";
var url = "/booking/list";
var app = angular.module('myApp', []);
app.controller('myCtrl', function ($scope, $http) {
    $http.get(url)
        .then(function (response) {
        console.log(response.data);
        $scope.flights = response.data;
        var flightList = response.data;
        if (flightList[0].airlinesID != "-") {
            var x = document.getElementById("flightChoice");
            for (var i = 0; i < flightList.length; i++) {
                var option = document.createElement("option");
                option.text = flightList[i].airlinesID;
                var jsonValue = void 0;
                jsonValue = flightList[i].airlinesID + "@" + flightList[i].nameAL + "@" + flightList[i].flightNumber + "@" + flightList[i].flyingTo + "@" + flightList[i].flightDate + "@" + flightList[i].flightTime + "@" + flightList[i].numberOfSeat;
                option.value = jsonValue;
                if (x != null)
                    x.options.add(option);
            }
        }
        else {
            //Disable submit button
            var submitButt = document.getElementById("bookingSubmit");
            submitButt.innerHTML = "There is no avaliable flights as of now.";
            submitButt.disabled = true;
            //disable search options
            var inputBox1 = document.getElementById("myInput");
            var inputBox2 = document.getElementById("myInputDate");
            var inputBox3 = document.getElementById("myInputCountry");
            inputBox1.disabled = true;
            inputBox2.disabled = true;
            inputBox3.disabled = true;
        }
    });
});
function searchAirlines() {
    //get html element containing user input
    var selector = document.getElementById("myInput");
    //get html element hidden div
    var div = document.getElementById("Hiddentable");
    var submitButt = document.getElementById("bookingSubmit");
    //var submitButt2 = <HTMLSelectElement>document.getElementById("bookingReserve");
    //get html element for other search options
    if (selector.value != "") {
        //Uppercase user input
        var input = selector.value.toUpperCase();
        var table = document.getElementById("airlinesDisplay");
        var tableRow = table.getElementsByTagName("tr");
        for (var i = 0; i < tableRow.length; i++) {
            var td = tableRow[i].getElementsByTagName("td")[1];
            var td2 = tableRow[i].getElementsByTagName("td")[3];
            var td3 = tableRow[i].getElementsByTagName("td")[4];
            if (td || td2 || td3) {
                var textValue = td.textContent || td.innerText;
                var textValue2 = td2.textContent || td2.innerText;
                var textValue3 = td3.textContent || td3.innerText;
                if (textValue.toUpperCase().indexOf(input) > -1 || textValue2.toUpperCase().indexOf(input) > -1 || textValue3.toUpperCase().indexOf(input) > -1) {
                    tableRow[i].style.display = "";
                }
                else {
                    tableRow[i].style.display = "none";
                }
            }
        }
        div.style.display = "block";
        submitButt.disabled = false;
    }
    else {
        div.style.display = "none";
        submitButt.disabled = true;
    }
}
function searchDateCountry() {
    var selector = document.getElementById("myInputDate");
    var selector2 = document.getElementById("myInputCountry");
    var div = document.getElementById("Hiddentable");
    var inputDate = selector.value.toUpperCase();
    var inputCountry = selector2.value.toUpperCase();
    var table = document.getElementById("airlinesDisplay");
    var tableRow = table.getElementsByTagName("tr");
    var submitButt = document.getElementById("bookingSubmit");
    if (selector.value != "" && selector2.value == "") {
        for (var i = 0; i < tableRow.length; i++) {
            var td = tableRow[i].getElementsByTagName("td")[4];
            if (td) {
                var textValue = td.textContent || td.innerText;
                if (textValue.toUpperCase().indexOf(inputDate) > -1) {
                    tableRow[i].style.display = "";
                }
                else {
                    tableRow[i].style.display = "none";
                }
            }
        }
        div.style.display = "block";
        submitButt.disabled = false;
    }
    else if (selector.value == "" && selector2.value != "") {
        for (var i = 0; i < tableRow.length; i++) {
            var td = tableRow[i].getElementsByTagName("td")[3];
            if (td) {
                var textValue = td.textContent || td.innerText;
                if (textValue.toUpperCase().indexOf(inputCountry) > -1) {
                    tableRow[i].style.display = "";
                }
                else {
                    tableRow[i].style.display = "none";
                }
            }
        }
        div.style.display = "block";
        submitButt.disabled = false;
    }
    else if (selector.value != "" && selector2.value != "") {
        for (var i = 0; i < tableRow.length; i++) {
            var td = tableRow[i].getElementsByTagName("td")[4];
            var td2 = tableRow[i].getElementsByTagName("td")[3];
            if (td && td2) {
                var textValue = td.textContent || td.innerText;
                var textValueNext = td2.textContent || td2.innerText;
                if (textValue.toUpperCase().indexOf(inputDate) > -1 && textValueNext.toUpperCase().indexOf(inputCountry) > -1) {
                    tableRow[i].style.display = "";
                }
                else {
                    tableRow[i].style.display = "none";
                }
            }
        }
        div.style.display = "block";
        submitButt.disabled = false;
    }
    else {
        div.style.display = "none";
        submitButt.disabled = true;
    }
}
