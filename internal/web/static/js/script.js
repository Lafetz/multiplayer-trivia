

document.addEventListener("htmx:confirm", function(e) {
  e.preventDefault()
  if (!e.target.hasAttribute('hx-confirm')) {
    e.detail.issueRequest(true);
    return;
} 
  Swal.fire({
    title: e.detail.question,
    showCancelButton: true,
cancelButtonText: "cancel",  
confirmButtonText: "Confirm",
reverseButtons: true,
confirmButtonColor: '#0284c7',
   
  }).then(function(result) {
    if(result.isConfirmed) e.detail.issueRequest(true) 
  })
})
///
//
function countdownTimer(initialTime) {
  return {
      timeRemaining: 0,

      startCountdown() {
          this.timeRemaining = initialTime;
          // Decrement time remaining every second until it reaches 0
          const countdownInterval = setInterval(() => {
              if (this.timeRemaining > 0) {
                  this.timeRemaining--;
              } else {
                  clearInterval(countdownInterval);
              }
          }, 1000);
      },

      formatTimeRemaining() {
          if (this.timeRemaining < 0) {
              return '0 seconds';
          } else {
              return `${this.timeRemaining} seconds`;
          }
      }
  };
}