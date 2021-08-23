 let scroll = new SmoothScroll('.navbar a[href*="#"]', {
    speed: 800
 });

 const msgAlert = document.querySelector('.success')

 setTimeout(() => {
  msgAlert.classList.add("hidden");
 }, 6000)
