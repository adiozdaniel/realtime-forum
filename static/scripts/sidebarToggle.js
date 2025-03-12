document.addEventListener("DOMContentLoaded", function () {
    const menuToggle = document.querySelector("[data-menu-toggle]");
    const sidebar = document.querySelector("[data-sidebar]");

    if (menuToggle && sidebar) {
        menuToggle.addEventListener("click", () => {
            sidebar.classList.toggle("active");
        });
    }
});