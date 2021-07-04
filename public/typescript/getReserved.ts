var url = "/reserved/list";
var app = angular.module('myAppRVS', []);
app.controller('myCtrlRVS', function($scope, $http) {
  $http.get(url)
  .then(function(response: { data: any; }) { 
    //Push data to table with angularJS
    $scope.flights = response.data;

    //Push data into select
    var flightList = response.data;
    if(flightList[0].airlinesID != "-")
    {
      var x = <HTMLSelectElement>document.getElementById("flightChoice");
      for(let i =0; i < flightList.length; i++)
      {
          var option = document.createElement("option") as HTMLOptionElement;
          option.text = flightList[i].airlinesID;
          let jsonValue: string;
          jsonValue = flightList[i].airlinesID + "@" + flightList[i].airlineName + "@" + flightList[i].flightNumber + "@" + flightList[i].flyingTo + "@" + flightList[i].bookingDate + "@" + flightList[i].bookingTime + "@" +flightList[i].ticketAmount;
          option.value = jsonValue;
          if(x != null)
              x.options.add(option);
      }
    }
    else 
    {
      let submitButt = <HTMLSelectElement>document.getElementById("reservedSubmit");
      submitButt.innerHTML = "Error you have not made a reservation of any tickets yet";
      submitButt.disabled = true;
      let deleteButt = <HTMLSelectElement>document.getElementById("reservedDelete");
      deleteButt.innerHTML = "Error you have not made a reservation of any tickets yet";
      deleteButt.disabled = true;
    }
  });
});
function submitOrDelete()
{
  let checkbox = <HTMLSelectElement>document.getElementById("forDelete");
  console.log(checkbox.value);
  console.log(checkbox.value == "Submit");
  if(checkbox.value == "Submit")
  {
    checkbox.value = "Delete";
  }
  else if(checkbox.value == "Delete")
  {
    checkbox.value = "Submit";
  }
  displayTextBox()
}
function displayTextBox()
{
  //get user choice
  let x = <HTMLSelectElement>document.getElementById("flightChoice");
  //get div to add input rows in
  let divFor = <HTMLSelectElement>document.getElementById("inputBySeats");
  let creditcard = <HTMLSelectElement>document.getElementById("cardNum");
  //get submit button and delete button
  let submitButt = <HTMLSelectElement>document.getElementById("reservedSubmit");
  let deleteButt = <HTMLSelectElement>document.getElementById("reservedDelete");
  let checkbox = <HTMLSelectElement>document.getElementById("forDelete");

  //split value in select to get number of rows neeeded
  let split = x.value.split("@");
  let numberOf = parseInt(split[split.length -1]);

  if(x.value != "none" && checkbox.value == "Submit")
  {
    //remove generated input
    divFor.innerHTML = "";

    //Create input element base on number of seats
    for(let i = 0; i < numberOf; i++)
    {
      divFor.appendChild(document.createTextNode("Ticket " + (i+1) + " Passport Number  "));
      let input = document.createElement("input");
      input.type = "text";
      input.name = "passNum" + i;
      input.id = "passNum" + i;
      input.className = "form-control";
      input.required = true;
      divFor.appendChild(input);

      divFor.appendChild(document.createTextNode("Ticket " + (i+1) + " Customer Name  "));
      let input2 = document.createElement("input");
      input2.type = "text";
      input2.name = "name" + i;
      input2.id = "name" + i;
      input2.className = "form-control";
      input2.required = true;
      divFor.appendChild(input2);
      // Append a line break 
      divFor.appendChild(document.createElement("br"));
    }
    //Enable submit button
    submitButt.disabled = false;
    deleteButt.disabled = true;
    creditcard.disabled = false;
  } 
  else if (x.value != "none" && checkbox.value == "Delete")
  {
    divFor.innerHTML = "";
    submitButt.disabled = true;
    deleteButt.disabled = false;
    creditcard.disabled = true;
  } 
  else 
  {
    //remove generated table if user chose none
    divFor.innerHTML = "";
    submitButt.disabled = true;
    deleteButt.disabled = true;
    creditcard.disabled = true;
    //disabled submitbutton
  }
}

function formTypeValue()
{
  let x = <HTMLSelectElement>document.getElementById("flightChoice");
  //split value in select to get number of rows neeeded
  let split = x.value.split("@");
  let numberOf = parseInt(split[split.length -1]);
  let value = "";
  let value2 = "";
  for(let i = 0; i < numberOf; i++)
  {
    let inputSelector = <HTMLSelectElement>document.getElementById("passNum" + i);
    let inputSelector2 = <HTMLSelectElement>document.getElementById("name" + i);
    if(i == 0)
    {
      value += inputSelector.value;
      value2 += inputSelector2.value;
    } 
    else
    {
      value = value + "@" + inputSelector.value;
      value2 = value2 + "@" + inputSelector2.value;
    }
  }
  let hiddenPPNum = <HTMLSelectElement>document.getElementById("passportValues");
  let hiddenName = <HTMLSelectElement>document.getElementById("nameValues");
  hiddenPPNum.value = value;
  hiddenName.value = value2;
}

function updateFormType()
{
  let hiddenValue = <HTMLSelectElement>document.getElementById("formType");
  hiddenValue.value = "Delete";
}