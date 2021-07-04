"use strict";
function passwordCheck() {
    var pw = document.getElementById("Password");
    var confirmpw = document.getElementById("cPassword");
    var reSubmitButt = document.getElementById("regSubmit");
    var msg = document.getElementById("passwordMsg");
    if (pw.value == confirmpw.value && pw.value != "" && confirmpw.value != "") {
        reSubmitButt.disabled = false;
        msg.innerText = "Passwords matching";
        msg.style.color = "green";
    }
    else if (pw.value == "" && confirmpw.value == "") {
        reSubmitButt.disabled = true;
        msg.innerText = "";
    }
    else {
        reSubmitButt.disabled = true;
        msg.innerText = "Passwords not matching";
        msg.style.color = "red";
    }
}
