//Search name

function searchFunc() {
    const items = document.querySelectorAll(".featured__card")
    // let itemName = document.querySelectorAll(".featured__card--subhead")
    let input = document.querySelector("#search")
    let inputValue = input.value.toLowerCase()
  
    for (i = 0; i < items.length; i++) {
      let itemValue = items[i].textContent
      if (itemValue.toLowerCase().indexOf(inputValue) > -1){
        items[i].style.display = "grid";
      } else {
        items[i].style.display = "none";
      }
    }
}



//Sort Price

function sortPrice(){
    let sortSelect = document.querySelector(".priceSort").value

    if (sortSelect === "Highest") {
        itemArray.sort((a,b) => (a.price < b.price ? 1 : -1))
    } else if (sortSelect === "Lowest") {
        itemArray.sort((a,b) => (a.price > b.price ? 1 : -1))
    }

    itemContainer.innerHTML = ""

    for(i = 0; i < itemArray.length; i++){
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
        </div>`};
}
