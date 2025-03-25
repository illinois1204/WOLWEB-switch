function openModal(id) {
  document.getElementById(id).style.display = "flex";
}

function closeModal(id) {
  document.getElementById(id).style.display = "none";
}

// Закрытие при клике вне окна
// window.onclick = function(event) {
//   const modal = document.getElementById("modal");
//   if (event.target === modal) {
//     closeModal("modal");
//   }
// }

window.addEventListener("click", function(event) {
  const modals = document.querySelectorAll(".modal");
  modals.forEach(modal => {
    if (event.target === modal) modal.style.display = "none";
  });
});
