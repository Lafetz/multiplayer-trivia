//create new game

//
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
   // text: `I ask you... ${e.detail.question}`
  }).then(function(result) {
    if(result.isConfirmed) e.detail.issueRequest(true) // use true to skip window.confirm
  })
})
function openJoinModal() {
    document.querySelector('.dialog-join').show();
  }
  function closeJoinModal() {
    document.querySelector('.dialog-join').hide();
  }  //hx-confirm="Some confirm text here"