const sidebar = document.querySelector('#sidebar'); 
const menuToggleBtn = document.querySelector('#menuToggle'); 

// Toggle mobile menu 
function toggleMobileMenu() { 
    const isVisible = sidebar.style.display === 'block'; sidebar.style.display = isVisible ? 'none' : 'block'; 
}

// Handle window resize 
function handleResize() { 
    if (window.innerWidth >= 1024) { sidebar.style.display = 'block'; } else { sidebar.style.display = 'none'; } 
}


// Initialize function 
function init() {
    lucide.createIcons();

    // Event listeners 
    menuToggleBtn?.addEventListener('click', toggleMobileMenu);
    // Initial resize check 
    handleResize(); 
}

// Start the application 
document.addEventListener('DOMContentLoaded', init);