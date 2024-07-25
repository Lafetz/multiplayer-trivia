

document.addEventListener("htmx:confirm", function(e) { // confirm before sending
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

function countdownTimer(initialTime) { // timer
  return {
      timeRemaining: 0,

      startCountdown() {
          this.timeRemaining = initialTime;
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
// toast
function toastComponent() {
  return {
      visible: false,
      msg: "Hello, this is a toast message!",

      isVisible() {
          return this.visible;
      },

      message() {
          return this.msg;
      },

      showToast() {
          this.visible = true;
          setTimeout(() => {
              this.visible = false;
          }, 3000); 
      }
  };
}
function fillAndSubmitForm(username, password) {
    document.getElementById('email').value = username;
    document.getElementById('password').value = password;
    document.getElementById('signin-form').requestSubmit();
}