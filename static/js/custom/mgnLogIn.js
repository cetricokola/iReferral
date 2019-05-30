
   document.querySelector("#mySubmit").onclick = (e) => {
      e.preventDefault()

      document.querySelector("#nationalId").classList.remove("border-danger")
      document.querySelector("#nationalIdError").style.display = "none"
      document.querySelector("#password").classList.remove("border-danger")
      document.querySelector("#passwordError").style.display = "none"

      axios.post("/management_login", {
        nationalId: document.querySelector("#nationalId").value,
         password: document.querySelector("#password").value
      })
         .then(data => {
            window.location.href = data.data.Response
            alert("Successful log in")
         })
         .catch(err => {
            var resp = err.response.data

            if (resp.NationalId == "invalid") {
               alert("Incorrect National id")
               document.querySelector("#nationalId").classList.add("border-danger")
               document.querySelector("#nationalIdError").style.display = "block"
            }

            if (resp.MgnPassword == "invalid") {
               alert("Incorrect password")
               document.querySelector("#password").classList.add("border-danger")
               document.querySelector("#passwordError").style.display = "block"
            }
         })

   }
