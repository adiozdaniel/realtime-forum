// DOM Elements 
const darkModeToggle = document.querySelector('#darkModeToggle'); 

lucide.createIcons();


// Toggle dark mode 
function toggleDarkMode() { 
    document.body.classList.toggle('dark-mode'); localStorage.setItem('darkMode', document.body.classList.contains('dark-mode')); 
}


// Initialize function 
function init() {
    lucide.createIcons();
    // Event listeners 
    darkModeToggle?.addEventListener('click', toggleDarkMode); 
    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; 
    if (savedDarkMode) { document.body.classList.add('dark-mode'); } 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);