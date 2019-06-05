document.querySelector("#sub").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#f").classList.remove("border-danger")
    document.querySelector("#ferror").style.display = "none"
    document.querySelector("#l").classList.remove("border-danger")
    document.querySelector("#lerror").style.display = "none"
    document.querySelector("#id").classList.remove("border-danger")
    document.querySelector("#iderror").style.display = "none"
    

    axios.post("/create_emp", {
        first: document.querySelector("#f").value,
        last: document.querySelector("#l").value,
        nId: document.querySelector("#id").value
         })
       .then(data => {
          window.location.href = data.data.EmpCreate
          alert("Employee successfully created")
       })
       .catch(err => {
          var resp = err.response.data

          if (resp.EmpFirst == "invalid") {
           
             document.querySelector("#f").classList.add("border-danger")
             document.querySelector("#ferror").style.display = "block"
          }

          if (resp.EmpLast == "invalid") {
         
            document.querySelector("#l").classList.add("border-danger")
            document.querySelector("#lerror").style.display = "block"
         }

          if (resp.EmpNid == "invalid") {
           
             document.querySelector("#id").classList.add("border-danger")
             document.querySelector("#iderror1").style.display = "block"
          }

          if (resp.EmpNid == "incorrect") {
         
            document.querySelector("#id").classList.add("border-danger")
            document.querySelector("#iderror").style.display = "block"
         }

         if (resp.EmpNid == "empty") {
          
            document.querySelector("#id").classList.add("border-danger")
            document.querySelector("#iderror2").style.display = "block"
         }
       })

 }