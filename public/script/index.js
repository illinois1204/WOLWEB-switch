function toast(message, duration = 3 * 1000) {
    const toastEl = document.getElementById("toast");
    toastEl.textContent = message;
    toastEl.classList.add("show");

    setTimeout(() => {
      toastEl.classList.remove("show");
    }, duration);
}