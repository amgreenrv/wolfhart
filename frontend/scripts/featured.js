//Get the products from the database

let featuredContainer = document.querySelector(".featured");

async function fetchFeatured() {
    
    const response = await fetch(`http://localhost:8080/products`);
    const data = await response.json();

    for (i = 0; i < data.length; i++) {
        featuredContainer.innerHTML += `
        <div class="related__card">
        <a href="./products.html">
        <h3 class="related__card--subhead">${data[i].category}</h3>
        <img src="${data[i].img}></a>
    </div>`;
    }
    return data;
}

fetchFeatured();

// Filter script Below

