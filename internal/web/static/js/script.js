//create new game
const dialog = document.querySelector('.dialog-create');
const createBtn = document.querySelector('#dialog-create-open');
const closeCreate =document.querySelector('#dialog-create-close');

createBtn.addEventListener('click', () => dialog.show());
closeCreate.addEventListener('click', () => dialog.hide());
//join game
document.addEventListener("htmx:wsOpen", function() {
    dialog.hide()

});
document.body.addEventListener('htmx:configRequest', (evt) => {
    console.log("yeyeyey")
    if (evt.detail.triggeringEvent.type === 'foo') {
      evt.detail.headers['someheader'] = 'foo_value'
    }
    if (evt.detail.triggeringEvent.type === 'bar') {
      evt.detail.headers['someheader'] = 'bar_value'
    }
  })
const startBtn=document.querySelector('#keep');
console.log(startBtn)
document.addEventListener("htmx:wsBeforeSend",event=>{
    console.log(event)
})
document.addEventListener('htmx:afterRequest', function(evt) {
    console.log('hhsew')
    if(evt.detail.xhr.status == 404){
        /* Notify the user of a 404 Not Found response */
        return alert("Error: Could Not Find Resource");
    } 
    if (evt.detail.successful != true) {
        /* Notify of an unexpected error, & print error to console */
        alert("Unexpected Error");
        return console.error(evt);
    }
    if (evt.detail.target.id == 'info-div') {
        /* Execute code on the target of the HTMX request, which will
        be either the hx-target attribute if set, or the triggering 
        element itself if not set. */
        let infoDiv = document.getElementById('info-div');
        infoDiv.style.backgroundColor = '#000000';  // black background
        infoDiv.style.color = '#FFFFFF';  // white text
    }
});
document.addEventListener('htmx:configRequest', function(event) {
    console.log('idk what is happening',event)
    // Check if the request is targeting a specific element or using a specific class
    if (event.detail.elt.classList.contains('clickable')) {
        // Modify the data payload before sending it over WebSocket
        event.detail.parameters.message = "Modified message: " + event.detail.parameters.message;
        // Send the modified data over WebSocket
        ws.send(JSON.stringify(event.detail.parameters));
    }
});
startBtn.addEventListener('htmx:wsBeforeSend', function(event) {
    console.log('idk what is hassssppening',event)
    // Check if the request is targeting a specific element or using a specific class
    if (event.detail.elt.classList.contains('clickable')) {
        // Modify the data payload before sending it over WebSocket
        event.detail.parameters.message = "Modified message: " + event.detail.parameters.message;
        // Send the modified data over WebSocket
        ws.send(JSON.stringify(event.detail.parameters));
    }
});
document.querySelector('#hello').addEventListener("htmx:wsAfterMessage",(x)=>{
console.log("workd")
})
startBtn.addEventListener("htmx:wsAfterMessage",(x)=>{
    console.log("woasasasrkd")
    })