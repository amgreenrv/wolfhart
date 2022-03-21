//Get the products from the database

let itemContainer = document.querySelector(".inventory__container");
let featuredContainer = document.querySelector(".featured__container");

//pull data from database
async function fetchProducts() {
    console.log("fetchProducts was hit");
    const response = await fetch(`http://localhost:8080/products`);
    const data = await response.json();
    console.log(data);
    for (i = 0; i < data.length; i++) {
            itemContainer.innerHTML += `
            <div class="featured__card">
                <div class="featured__category">${data[i].category}</div>
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

async function fetchFeatured() {
    console.log("fetchFeatured was hit");
    const response = await fetch(`http://localhost:8080/featured`);
    const data = await response.json();

    for (i = 0; i < data.length; i++) {
        featuredContainer.innerHTML += `
        <div class="featured__card">
                <h3 class="featured__card--subhead">${data[i].prodname}</h3>
                <div class="featured__card--cardimg">
                    <a href="./detail.html"><img src="${data[i].img}" alt="component image"></a>
                </div>
        </div>`;
    }
    return data;
}

fetchProducts();
fetchFeatured();

// Filter script 

var categoryList = document.querySelectorAll(".fcategory")
var category = document.querySelectorAll(".featured__category")
setTimeout(filterFunc,100)


for (i = 0; i < categoryList.length; i++) {
    categoryList[i].addEventListener("click", filterFunc)
}

function filterFunc() {
    var ip = document.querySelectorAll(".featured__card")
    category = document.querySelectorAll(".featured__category")
    for (i = 0; i < categoryList.length; i++) {
        if (categoryList[i].checked) {
            for (v = 0; v < ip.length; v++) {
                if (categoryList[i].value == "all") {
                    ip[v].style.display = "grid"
            }
            else {
                if (category[v].innerHTML != categoryList[i].value){
                    category[v].parentNode.style.display = "none"
                }
                else {
                    category[v].parentNode.style.display = "grid"
                    }
                }
            }
        }
    }
}
