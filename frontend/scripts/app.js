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
    if(name==""){
    alert("Please enter your name.");
    return false;
    }

    const email = document.getElementById("email").value;
    if(email==""){
    alert("Please enter your email.");
    return false;
    }else{
    const re = /^(?:[a-z0-9!#$%&amp;'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&amp;'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$/;
    const x=re.test(email);
    if(x){
    }else{
    alert("Please enter a valid email.");
    return false;
    } 
    } 

    const phone = document.getElementById("phone").value;
    if(phone == ""){
    alert("Please enter your phone number.");
    return false;
    }
} 