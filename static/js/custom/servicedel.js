document.querySelector("#servdel").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#h").classList.remove("border-danger")
    // document.querySelector("#hdmerror").style.display = "none"
       
    axios.post("/del_service", {
        code: document.querySelector("#h").value
         })
       .then(data => {
          window.location.href = data.data.Servdel
          alert("Service successfully deleted")
       })
       .catch(err => {
          var resp = err.response.data

          if (resp.ServCode == "invalid") {
              alert("Service code does not exist")
             document.querySelector("#h").classList.add("border-danger")
            
          }

          if (resp.ServCode == "incorrect") {
              alert("Service code is incorrect")
            document.querySelector("#h").classList.add("border-danger")
            
         }

         
         if (resp.ServCode == "empty") {
             alert("Service code field is empty.Enter a value in it")
            document.querySelector("#h").classList.add("border-danger")
           
         }

         if (resp.ServCode == "notcode") {
             alert("Service code is not part of your services")
            document.querySelector("#h").classList.add("border-danger")
           
         }
       })

 }