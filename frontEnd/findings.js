let currentPage = 1;
const itemsPerPage = 3;
const container = document.getElementById("findings-container")

document.addEventListener("DOMContentLoaded", () => {
    const queryCity = localStorage.getItem('queryCity')
    const heading = document.getElementById('findind-city')
    heading.innerText = "Trainings in " + queryCity
    loadContainer()
})

function showEvents(data) {
    if (data === null) {
        const newElement = document.createElement("h2")
        newElement.innerText = 'Failed to load events, try again later'

        container.appendChild(newElement)
    }

}

async function getData(url) {
    try {
        const response = await fetch(url)
        const data = response.json()
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

async function loadContainer() {
    container.innerHTML = ''

    const start = (currentPage - 1) * itemsPerPage
    const baseUrl = "https://fit-bond.com"
    const params = {
        city: localStorage.getItem('queryCity'),
        start: start,
        length: itemsPerPage
    }

    const queryString = URLSearchParams(params).toString()
    const url = "{baseUrl}?{queryString}"

    showEvents(await getData(url))
}