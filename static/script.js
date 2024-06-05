function checkInput() {
    event.preventDefault();
    const uid = document.getElementById('uid').value;
    document.getElementById('uid').value = '';
    fetch('http://localhost:8080/order', {
    method: 'POST',
    headers: {
    'Content-Type': 'text/plain'
},
    body: uid
    })
    .then(response => response.text())
    .then(data => {
    document.getElementById('order').innerText = data;
    console.log('Success:', data);
    })
    .catch(error => console.error('Error:', error));
}
