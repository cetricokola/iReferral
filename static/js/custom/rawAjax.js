var huduma = document.getElementById('huduma'),
    service = document.getElementById('service');


var httpRequest = new XMLHttpRequest();

console.log("loaded");

//AddEventListener
service.addEventListener('keypress', makeRequest);

function makeRequest() {
    console.log(event.key);
    setTimeout(() => {
        //Get Values
        var hudumaValue = huduma.value;
        var serviceValue = service.value;

        //Run Ajax Call
        // httpRequest.setRequestHeader('Content-type', 'application/json')

        if (!httpRequest) {
            alert("You have entered wrong details");
            return false;
        }

        httpRequest.onreadystatechange = getHudumaNoAndService;

        //Get Values from DB
        httpRequest.open("GET", '/referpatient', true);
        httpRequest.send(null);

        // console.log(hudumaValue, serviceValue);
    }, 2000);
}

function getHudumaNoAndService() {
    if (httpRequest.readyState === XMLHttpRequest.DONE) {
        if (httpRequest.status === 200) {
            //Enable Submit Button if Ajax is successfull
            var submitButton = document.getElementById('submit');
            submitButton.removeAttribute('disabled');
            alert("Correct Details");
        } else {
            var error = "Huduma Number or Service or the service provided is not available"
            alert(error);
        }
    }
}