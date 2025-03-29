function openFillModel(modal, payload) {
  document.getElementById("edit-device-input-id").value = payload.id;
  document.getElementById("edit-device-input-name").value = payload.name;
  document.getElementById("edit-device-input-mac").value = payload.mac;
  document.getElementById("edit-device-input-port").value = payload.port;
  openModal(modal);
}

function openModal(id) {
  document.getElementById(id).style.display = "flex";
}

function closeModal(id) {
  document.getElementById(id).style.display = "none";
}

window.addEventListener("click", function(event) {
  const modals = document.querySelectorAll(".modal");
  modals.forEach(modal => {
    if (event.target === modal) modal.style.display = "none";
  });
});
