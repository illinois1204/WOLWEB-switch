function toggleMenu(button) {
  const dropdown = button.parentElement;
  dropdown.classList.toggle("open");

  // Закрыть другие открытые меню
  document.querySelectorAll('.dropdown').forEach(d => {
    if (d !== dropdown) d.classList.remove('open');
  });
}

// Закрытие меню при клике вне его
document.addEventListener("click", function (event) {
  if (!event.target.closest(".dropdown")) {
    document.querySelectorAll('.dropdown').forEach(d => d.classList.remove('open'));
  }
});
