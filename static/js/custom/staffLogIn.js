document.querySelector("#submit2").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#empid").classList.remove("border-danger")
    document.querySelector("#empiderror").style.display = "none"
    document.querySelector("#password2").classList.remove("border-danger")
    document.querySelector("#passwordError2").style.display = "none"

    axios.post("/staff_login", {
      empid: document.querySelector("#empid").value,
       password: document.querySelector("#password2").value
    })
       .then(data => {
          window.location.href = data.data.Success
          alert("Successful log in")
       })
       .catch(err => {
          var resp = err.response.data

          if (resp.EmpId == "invalid") {
             alert("Incorrect employee id")
             document.querySelector("#empid").classList.add("border-danger")
             document.querySelector("#empiderror").style.display = "block"
          }

          if (resp.StaffPassword == "invalid") {
             alert("Incorrect password")
             document.querySelector("#password2").classList.add("border-danger")
             document.querySelector("#passwordError2").style.display = "block"
          }
       })

 }