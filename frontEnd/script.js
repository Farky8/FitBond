
function makeJson(requestdata) {
    const requestObj = {
        "heading": requestdata.get('heading'),
        "about": requestdata.get('about'),
        "city": requestdata.get('city'),
        "capacity": Number(requestdata.get('capacity')),
        "price": Number(requestdata.get('price'))
    }

    const requestJson = JSON.stringify(requestObj);
    return requestJson;
}

async function sendJson(requestJson) {
    const apiUrl = "https://fit-bond-586541250183.europe-central2.run.app/home/create" // TODOO after gcp
    try {
        const response = await fetch(apiUrl, {
            method: "POST",
            mode: "no-cors",
            headers: {
                "Content-Type": "application/json",
            },
            body: requestJson,
        })
        if (response.status === 201) {
            return 'Training successfully POSTED!';
        } else {
            return 'Error (code:' + response.status + '): ' + response.text();
        }
    } catch {
        return 'Could not reach the server, try again later.'
    }

}

const form = document.querySelector('#create-info')

form.addEventListener('submit', event => {
    event.preventDefault();  // the page will not refresh and redirect
    const requestdata = new FormData(event.target);

    const requestJson = makeJson(requestdata);

    sendJson(requestJson)
	.then( msg => alert(msg) )
	.catch( msg => alert(msg) )
})
