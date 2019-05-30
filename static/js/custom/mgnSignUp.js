
document.querySelector("#mySubmit1").onclick = (e) => {
    e.preventDefault()

   
    document.querySelector("#username").classList.remove("border-danger")
    document.querySelector("#usernameError").style.display = "none"
    document.querySelector("#id").classList.remove("border-danger")
    document.querySelector("#idError").style.display = "none"
    document.querySelector("#email").classList.remove("border-danger")
    document.querySelector("#emailError").style.display = "none"
    document.querySelector("#cpassword").classList.remove("border-danger")
    document.querySelector("#cpasswordError").style.display = "none"
    document.querySelector("#password11").classList.remove("border-danger")
    document.querySelector("#passwordError11").style.display = "none"

    axios.post("/management_signup", {
        username: document.querySelector("#username").value,
        id: document.querySelector("#id").value,
        email: document.querySelector("#email").value,
        password: document.querySelector("#password11").value,
        cpassword: document.querySelector("#cpassword").value
    })
        .then(data => {
            window.location.href = data.data.Feedback
            alert("Successful account registration.To continue using iReferral log in")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.Username == "invalid") {
                alert("Username already taken")
                document.querySelector("#username").classList.add("border-danger")
                document.querySelector("#usernameError").style.display = "block"
            }
            if (resp.Id == "invalid") {
                alert("National id already exists")
                document.querySelector("#id").classList.add("border-danger")
                document.querySelector("#idError").style.display = "block"
            }
            if (resp.Mpassword == "incorrect"){
                alert("Password must contain atleast 8 characters composed of both capital,small letters and digits")
                document.querySelector("#password11").classList.add("border-danger")
                document.querySelector("#passwordError11").style.display = "block"
                document.querySelector("#cpassword").classList.add("border-danger")
                document.querySelector("#cpasswordError2").style.display = "block"
            }

            if ( resp.Mpassword == "invalid" && resp.Cpassword == "invalid") {
                alert("password and confirm password does not match")
                document.querySelector("#password11").classList.add("border-danger")
                document.querySelector("#passwordError11").style.display = "block"
                document.querySelector("#cpassword").classList.add("border-danger")
                document.querySelector("#cpasswordError").style.display = "block"
            }

            if (resp.Email == "invalid"){
                alert("Invalid email address")
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailError").style.display = "block"
            }

            if (resp.Id == "wrong"){
                alert("Invalid national id")
                document.querySelector("#id").classList.add("border-danger")
                document.querySelector("#idError2").style.display = "block"
            }
        })

}
