document.addEventListener("DOMContentLoaded", () => {
    fillPage('GET')
})

function showEvent(data) {
    if (data === null) {
        const newElement = document.createElement("h2")
        newElement.innerText = 'Failed to load event, try again later'

        container.appendChild(newElement)
    } else {
        const heading = document.getElementById("event-heading")
        heading.innerText = data.heading

        const about = document.getElementById("event-about")
        about.innerText = data.about

        const appCapp = document.getElementById("event-app-capp")
        appCapp.innerText = data.applied.toString() + '/' + data.capacity.toString()

        const price = document.getElementById("event-price")
        price.innerText = data.price.toString() + 'kc'

        const button = document.getElementById("apply-button")
        button.onclick = () => apply()
    }

}

async function getData(url) {
    try {
        const response = await fetch(url)
        const data = await response.json()
        if (response.status === 200) {
            return data
        }
        console.log('Received code=' + response.status + ' from fit-bond api')
        return null
    } catch (error) {
        console.log(error)
        return null
    }
}

async function putData(url) {
    try {
        const response = await fetch(url, {
            method: 'PUT',
        })
        const data = await response.json()
        if (response.status === 200) {
            return data
        }
        console.log('Received code=' + response.status + ' from fit-bond api')
        return null
    } catch (error) {
        console.log(error)
        return null
    }
}


async function fillPage(method) {
    const id = localStorage.getItem('event-id')
    const url = "https://fit-bond-586541250183.europe-central2.run.app/home/search/" + id

    if (method === 'GET') {
        showEvent(await getData(url))
    } else {
        showEvent(await putData(url))
    }
}

async function apply() {
    fillPage('PUT')
}