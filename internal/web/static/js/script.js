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