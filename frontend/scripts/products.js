//Get the products from the database

let itemContainer = document.querySelector(".inventory");

//pull data from database
async function fetchProducts() {
    console.log("function was hit");
    const response = await fetch(`http://localhost:8080/products`);
    const data = await response.json();
    console.log(data);
    for (i = 0; i < data.length; i++) {
            itemContainer.innerHTML +=
            `<div class="featured__card"><div class="category">${data[i].category}</div>
            <h3 class="featured__card--subhead">${data[i].prodname}</h3>
                <div class="featured__card--cardimg">
                    <a href="./detail.html"><img src="${data[i].img}" alt="h1"></a>
                </div>
                <div class="featured__card--description">
                <p class="featured__card--price">Price: ${data[i].price}</p>
                <p>${data[i].proddesc}</p>
                <a href="./detail.html" class="featured__card--button">Buy Now</a>
                </div>
            </div>`;
    }
    return data;
}

fetchProducts();

// Filter script Below