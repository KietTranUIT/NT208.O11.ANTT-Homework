let form = document.getElementById("id")

console.log(document.getElementById('id').value)

form.addEventListener('submit', function(event) {
    event.preventDefault();
    let id = document.getElementById('id').value
    let title = document.getElementById('title').value
    let summary = document.getElementById('summary').value
    let image = document.getElementById('image').value
    let price = document.getElementById('price').value
    let number = document.getElementById('number').value
 
    let json = {
       id: id,
       title: title,
       summary: summary,
       image: image,
       price: price,
       number: number
    }
 
    let data = JSON.stringify(json)
    fetch('/create_product', {
       method: 'POST',
       headers: { 'Content-Type': 'application/json'},
       body: data
    })
    .then(response => {
       if(response.status === 200) {
          console.log("Them san pham thanh cong!")
       } else {
          console.log("Them san pham that bai!")
       }
    })
 })