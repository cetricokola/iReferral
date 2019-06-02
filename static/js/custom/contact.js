document.querySelector("#send").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#name").classList.remove("border-danger")
    document.querySelector("#nameerror").style.display = "none"
    document.querySelector("#email").classList.remove("border-danger")
    document.querySelector("#emailerror").style.display = "none"
    document.querySelector("#subject").classList.remove("border-danger")
    document.querySelector("#subjecterror").style.display = "none"
    document.querySelector("#message").classList.remove("border-danger")
    document.querySelector("#messageerror").style.display = "none"

    axios.post("/contact_us", {
        name: document.querySelector("#name").value,
        email: document.querySelector("#email").value,
        subject: document.querySelector("#subject").value,
        message: document.querySelector("#message").value
    })
        .then(data => {
            window.location.href = data.data.Contact
            alert("Email successful send")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.ContactEmail == "invalid") {
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailerror").style.display = "block"
            }

            if (resp.ContactEmail == "incorrect") {
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailerror2").style.display = "block"
            }

            if (resp.ContactName == "invalid") {
                document.querySelector("#name").classList.add("border-danger")
                document.querySelector("#nameerror").style.display = "block"
            }
            if (resp.ContactMessage == "invalid") {
                document.querySelector("#message").classList.add("border-danger")
                document.querySelector("#messageerror").style.display = "block"
            }

            if (resp.ContactSubject == "invalid") {
                alert("You cannot send an email without a subject")
                document.querySelector("#subject").classList.add("border-danger")
                document.querySelector("#subjecterror").style.display = "block"
            }
        })

}