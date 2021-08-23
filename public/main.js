 let scroll = new SmoothScroll('a[href*="#"]', {
    speed: 800
 });

 const msgAlert = document.querySelector('.success')

 setTimeout(() => {
  msgAlert.classList.add("hidden");
 }, 6000)
