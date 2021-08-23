 const msgAlert = document.querySelector('.success')

 msgAlert.setTimeout(() => {
  msgAlert.classList.add("hidden");
 }, 6000)

let scroll = new SmoothScroll('.navbar a[href*="#"]');
