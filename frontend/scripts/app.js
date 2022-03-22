// Mobile Nav Bar Menu

const menu = document.querySelector('#mobile-menu');
const menuLinks = document.querySelector('.navbar__menu');

menu.addEventListener('click', function() {
    menu.classList.toggle('is-active');
    menuLinks.classList.toggle('active');
});


// Homepage Slider

const slides = document.querySelectorAll('.slide');
const btns = document.querySelectorAll('.btn');
let currentSlide = 1;

// Manual Navigation For Slider on Homepage

const manualNav = function(manual){
    slides.forEach((slide) => {
    slide.classList.remove('active');

    btns.forEach((btn) => {
        btn.classList.remove('active');
    });
    });

    slides[manual].classList.add('active');
    btns[manual].classList.add('active');
};

btns.forEach((btn, i) => {
    btn.addEventListener("click", () => {
    manualNav(i);
    currentSlide = i;
    });
});


// Image Toggle for Individual Product Page
function clickMe(smallImg) {

    const fullImg = document.getElementById("imagebox");
    fullImg.src = smallImg.src;

}


//prevent auto submit of form

form.addEventListener('submit', e => {
    e.preventDefault();

    validate();
});


// Form Validation


function validate() {
    const name = document.getElementById("username").value;
    if(name == ""){
    alert("Please enter your name.");
    return false;
    }

    const email = document.getElementById("email").value;
    if(email == ""){
    alert("Please enter your email.");
    return false;
    } 

    const phone = document.getElementById("phone").value;
    if(phone == ""){
    alert("Please enter your phone number.");
    return false;
    }
} 