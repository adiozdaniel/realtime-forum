// DOM Elements 
const darkModeToggle = document.querySelector('#darkModeToggle'); 

lucide.createIcons();

// Handle window resize 
function handleResize() { 
    if (window.innerWidth >= 1024) { sidebar.style.display = 'block'; } else { sidebar.style.display = 'none'; } 
}

// Toggle dark mode 
function toggleDarkMode() { 
    document.body.classList.toggle('dark-mode'); localStorage.setItem('darkMode', document.body.classList.contains('dark-mode')); 
}


// Initialize function 
function init() {
    lucide.createIcons();
    // Event listeners 
    darkModeToggle?.addEventListener('click', toggleDarkMode); window.addEventListener('resize', handleResize); searchInput?.addEventListener('input', handleSearch); 
    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; if (savedDarkMode) { document.body.classList.add('dark-mode'); } 
    // Initial resize check 
    handleResize(); 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);