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
    } else {
        data.results.forEach(finding => {
            const div = document.createElement('div');
            div.classList.add('found-event')
            div.innerHTML = `
                <h3>${finding.heading}</h3>
                <p>Capacity: ${finding.applied}/${finding.capacity}</p>
                <p>Price: ${finding.price}</p>
            `

            div.onclick = () => {
                localStorage.setItem('event-id', finding.id.toString())
                window.location.href = `https://fit-bond.com/events`
            }

            container.appendChild(div)
        })
        document.getElementById('prev').disabled = currentPage === 1
        document.getElementById('next').disabled = data.length === itemsPerPage
    }

}

function changePage(direction) {
    currentPage += direction
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

async function loadContainer() {
    container.innerHTML = ''

    const start = (currentPage - 1) * itemsPerPage
    const baseUrl = "https://fit-bond-586541250183.europe-central2.run.app/home/search"
    const params = {
        city: localStorage.getItem('queryCity'),
        start: start,
        length: itemsPerPage
    }

    const queryString = new URLSearchParams(params).toString()
    const url = `${baseUrl}/?${queryString}`

    showEvents(await getData(url))
}