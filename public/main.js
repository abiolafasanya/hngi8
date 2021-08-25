 let scroll = new SmoothScroll('a[href*="#"]', {
    speed: 800
 });

 const msgAlert = document.querySelector('.success')

 setTimeout(() => {
  msgAlert.classList.add("hidden");
 }, 6000)

 let skillSection = document.querySelector('#skills-section')
 let progressBar = document.querySelectorAll('.skill-progress');

 const showProgressAnimation = () => {
   progressBar.forEach(progress => {
      progress.classList.add('skill-animation')
   })
 }

 const hideProgressAnimation = () => {
   progressBar.forEach(progress => {
      progress.classList.remove('skill-animation')
   })
 }


 window.addEventListener('scroll', () => {
   const sectionPos = skillSection.getBoundingClientRect().top
   const screenPos = window.innerHeight

   if(sectionPos < screenPos){
      console.log('show progress bar')
      showProgressAnimation()
   }
   else {
      console.log('hide progress bar')
      hideProgressAnimation()
   }
 })

 