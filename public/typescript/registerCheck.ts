function passwordCheck() {
    let pw = <HTMLSelectElement>document.getElementById("Password");
    let confirmpw = <HTMLSelectElement>document.getElementById("cPassword");
    let reSubmitButt = <HTMLSelectElement>document.getElementById("regSubmit");
    let msg = <HTMLSelectElement>document.getElementById("passwordMsg");

    if(pw.value == confirmpw.value && pw.value != "" && confirmpw.value !="")
    {
        reSubmitButt.disabled = false;
        msg.innerText = "Passwords matching"
        msg.style.color = "green";
    }
    else if (pw.value == "" && confirmpw.value == "")
    {
        reSubmitButt.disabled = true;
        msg.innerText = ""
    }
    else {
        reSubmitButt.disabled = true;
        msg.innerText = "Passwords not matching"
        msg.style.color = "red";
    }
}