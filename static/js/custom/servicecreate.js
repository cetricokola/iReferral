document.querySelector("#servcreate").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#nam").classList.remove("border-danger")
    document.querySelector("#namerror").style.display = "none"
    document.querySelector("#cod").classList.remove("border-danger")
    document.querySelector("#coderror").style.display = "none"
    document.querySelector("#coderror1").style.display = "none"
    document.querySelector("#dept").classList.remove("border-danger")
    document.querySelector("#depterror").style.display = "none"


    axios.post("/create_service", {
        name: document.querySelector("#nam").value,
        code: document.querySelector("#cod").value,
        dep: document.querySelector("#dept").value
    })
        .then(data => {
            window.location.href = data.data.ServCreate
            alert("Service successfully created")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.ServCode == "invalid") {
                alert("Service code already used")
                document.querySelector("#code").classList.add("border-danger")
                document.querySelector("#coderror").style.display = "block"
            }

            if (resp.ServCode == "incorrect") {
                alert("Incorrect service code")
                document.querySelector("#cod").classList.add("border-danger")
                document.querySelector("#coderror1").style.display = "block"
            }

            if (resp.ServDept == "noselection") {
                alert("You have not selected the department name")
                document.querySelector("#dept").classList.add("border-danger")
                document.querySelector("#depterror").style.display = "block"
            }

            if (resp.ServName == "invalid") {
                alert("Empty name field.Provide the value")
                document.querySelector("#nam").classList.add("border-danger")
                document.querySelector("#namerror").style.display = "block"
            }

        })

}