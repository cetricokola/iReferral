
document.querySelector("#patSubmit").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#first").classList.remove("border-danger")
    document.querySelector("#firstError").style.display = "none"
    document.querySelector("#last").classList.remove("border-danger")
    document.querySelector("#lastError").style.display = "none"
    document.querySelector("#hdm1").classList.remove("border-danger")
    document.querySelector("#hdm1Error").style.display = "none"
    document.querySelector("#phone1").classList.remove("border-danger")
    document.querySelector("#phone1Error").style.display = "none"
    document.querySelector("#dob").classList.remove("border-danger")
    document.querySelector("#dobError").style.display = "none"
    document.querySelector("#pass").classList.remove("border-danger")
    document.querySelector("#passError").style.display = "none"
    document.querySelector("#copass").classList.remove("border-danger")
    document.querySelector("#passError1").style.display = "none"

    axios.post("/patient_signup", {
        huduma: document.querySelector("#hdm1").value,
        first: document.querySelector("#first").value,
        last: document.querySelector("#last").value,
        phone: document.querySelector("#phone1").value,
        dob: document.querySelector("#dob").value,
        sex: document.querySelector("#sex").value,
        pass: document.querySelector("#pass").value,
        copass: document.querySelector("#copass").value

    })
        .then(data => {
            window.location.href = data.data.Res
            alert("Successful account registration.To continue using iReferral log in")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.PatFirstName =="empty"){
                alert("Fill in empty field")
                document.querySelector("#first").classList.add("border-danger")
                document.querySelector("#firstError").style.display = "block"
            }
            if (resp.PatLastName =="empty"){
                alert("Fill in empty field")
                document.querySelector("#last").classList.add("border-danger")
                document.querySelector("#lastError").style.display = "block"
            }
            if (resp.PatDoB =="empty"){
                alert("Fill in empty field")
                document.querySelector("#dob").classList.add("border-danger")
                document.querySelector("#dobError").style.display = "block"
            }

            if (resp.PatHudumaNo == "invalid") {
                alert("Huduma number already used")
                document.querySelector("#hdm1").classList.add("border-danger")
                document.querySelector("#hdm1Error").style.display = "block"
            }
            if (resp.PatPatPassword == "incorrect") {
                alert("Password must contain atleast 8 characters composed of both capital,small letters and digits")
                document.querySelector("#pass").classList.add("border-danger")
                document.querySelector("#passError").style.display = "block"
                document.querySelector("#copass").classList.add("border-danger")
            }

            if (resp.PatPatPassword == "invalid" && resp.Copass == "invalid") {
                alert("password and confirm password does not match")
                document.querySelector("#pass").classList.add("border-danger")
                document.querySelector("#copass").classList.add("border-danger")
                document.querySelector("#passError1").style.display = "block"

            }

            if (resp.PatHudumaNo == "wrong") {
                alert("Invalid huduma number")
                document.querySelector("#hdm1").classList.add("border-danger")
                document.querySelector("#hdmError1").style.display = "block"
            }

            if (resp.PatPhoneNo == "wrong") {
                alert("Invalid phone number")
                document.querySelector("#phone1").classList.add("border-danger")
                document.querySelector("#phone1Error").style.display = "block"
            }
        })

}
