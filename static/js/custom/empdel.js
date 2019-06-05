document.querySelector("#subdel").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#hdm").classList.remove("border-danger")
   
    
       
    axios.post("/del_emp", {
        empid: document.querySelector("#hdm").value
         })
       .then(data => {
          window.location.href = data.data.EmpDel
          alert("Employee successfully deleted")
       })
       .catch(err => {
          var resp = err.response.data

          if (resp.Empid == "invalid") {
              alert("Employee id does not exist")
             document.querySelector("#hdm").classList.add("border-danger")
            
          }

          if (resp.Empid == "incorrect") {
              alert("Employee id is incorrect")
            document.querySelector("#hdm").classList.add("border-danger")
           
         }

         
         if (resp.Empid == "empty") {
             alert("Employee id field is empty.Enter a value in it")
            document.querySelector("#hdm").classList.add("border-danger")
           
         }

         if (resp.Empid == "notstaff") {
             alert("Employee id entered is not part of your staff")
            document.querySelector("#hdm").classList.add("border-danger")
            
         }
       })

 }