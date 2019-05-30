
document.querySelector("#staffSubmit").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#emp").classList.remove("border-danger")
    document.querySelector("#empError").style.display = "none"
    document.querySelector("#em").classList.remove("border-danger")
    document.querySelector("#emError").style.display = "none"
    document.querySelector("#phone").classList.remove("border-danger")
    document.querySelector("#phoneError").style.display = "none"
    document.querySelector("#mypass").classList.remove("border-danger")
    document.querySelector("#mypassError").style.display = "none"
    document.querySelector("#comypass").classList.remove("border-danger")
    document.querySelector("#comypassError").style.display = "none"

    axios.post("/staff_signup", {
        emp: document.querySelector("#emp").value,
        em: document.querySelector("#em").value,
        phone: document.querySelector("#phone").value,
        mypass: document.querySelector("#mypass").value,
        comypass: document.querySelector("#comypass").value

    })
        .then(data => {
            window.location.href = data.data.Resp
            alert("Successful account registration.To continue using iReferral log in")
        })
        .catch(err => {
            var resp = err.response.data
            
            if (resp.StaffEmpId == "invalid") {
                alert("Employee id is already registered")
                document.querySelector("#emp").classList.add("border-danger")
                document.querySelector("#empError").style.display = "block"
            }

            if (resp.StaffEmpId == "incorrect") {
                alert("Employee id does not exist")
                document.querySelector("#emp").classList.add("border-danger")
                document.querySelector("#empError1").style.display = "block"
            }

            if (resp.StaffPassword == "incorrect") {
                alert("Password must contain atleast 8 characters composed of both capital,small letters and digits")
                document.querySelector("#mypass").classList.add("border-danger")
                document.querySelector("#mypassError").style.display = "block"
                document.querySelector("#comypass").classList.add("border-danger")
            }

            if (resp.StaffPassword == "invalid" && resp.StaffCopass == "invalid") {
                alert("password and confirm password does not match")
                document.querySelector("#mypass").classList.add("border-danger")
                document.querySelector("#comypass").classList.add("border-danger")
                document.querySelector("#comypassError").style.display = "block"

            }

            if (resp.StaffEmail == "wrong") {
                alert("Invalid email address")
                document.querySelector("#em").classList.add("border-danger")
                document.querySelector("#emError").style.display = "block"
            }

            if (resp.StaffPhoneNo == "wrong") {
                alert("Invalid phone number")
                document.querySelector("#phone").classList.add("border-danger")
                document.querySelector("#phoneError").style.display = "block"
            }
        })

}
