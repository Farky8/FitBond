
function makeJson(requestdata) {
    const requestObj = {
        "heading": requestdata.get('heading'),
        "about": requestdata.get('about'),
        "capacity": requestdata.get('capacity'),
        "price": requestdata.get('price')
    }

    const requestJson = JSON.stringify(requestObj);
    return requestJson;
}

async function sendJson(requestJson) {
    const apiUrl = ".../home/create" // TODOO after gcp
    try {
        const response = await fetch(apiUrl, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: requestJson,
        })
        if (response.status === 201) {
            return 'Training successfully POSTED!';
        } else {
            return 'Invalid data given.';
        }
    } catch {
        return 'Could not reach the server, try again later.'
    }

}

const form = document.querySelector('#create-info')

form.onsubmit('submit', event => {
    event.preventDefault();  // the page will not refresh and redirect
    const requestdata = new FormData(event.target);

    const requestJson = makeJson(requestdata);

    alert(sendJson(requestJson));
})