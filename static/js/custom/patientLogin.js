document.querySelector("#submit1").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#huduma1").classList.remove("border-danger")
    document.querySelector("#hudumaError1").style.display = "none"
    document.querySelector("#password1").classList.remove("border-danger")
    document.querySelector("#passwordError1").style.display = "none"

    axios.post("/patient_login", {
       huduma: document.querySelector("#huduma1").value,
       password: document.querySelector("#password1").value
    })
       .then(data => {
          window.location.href = data.data.Messages
          alert("Successful log in")
       })
       .catch(err => {
          var resp = err.response.data

          if (resp.PatientHuduma == "invalid") {
             alert("Incorrect huduma number")
             document.querySelector("#huduma1").classList.add("border-danger")
             document.querySelector("#hudumaError1").style.display = "block"
          }

          if (resp.PatientPassword == "invalid") {
             alert("Incorrect password")
             document.querySelector("#password1").classList.add("border-danger")
             document.querySelector("#passwordError1").style.display = "block"
          }
       })

 }