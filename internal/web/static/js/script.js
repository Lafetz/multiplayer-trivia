

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
// function openJoinModal() {
//     document.querySelector('.dialog-join').show();
//   }
//   function closeJoinModal() {
//     document.querySelector('.dialog-join').hide();
//   }  
