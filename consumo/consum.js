console.log('hola pipe')

const URL = 'http://localhost:8080/transformada'

fetch(URL)
    .then(res => res.json())
    .then(response =>{
        console.log(response)
    })