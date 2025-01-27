
const form = document.querySelector('#find-info')

form.addEventListener('submit', event => {
    event.preventDefault();

    const requestdata = new FormData(event.target);

    const queryCity = requestdata.get('find-city')
    localStorage.setItem('queryCity', queryCity)
    console.log('saved queryCity = ' + localStorage.getItem('queryCity'))
    window.location.href = "findings.html"
})