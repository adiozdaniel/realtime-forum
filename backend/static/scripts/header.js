import  {renderPosts, SAMPLE_POSTS} from "./index.js";
import {  sidebar } from "./sidebar.js";
// DOM Elements 
const menuToggleBtn = document.querySelector('#menuToggle');
const searchInput = document.querySelector('#searchInput'); 
const darkModeToggle = document.querySelector('#darkModeToggle'); 

// Toggle mobile menu
function toggleMobileMenu() {
    const isVisible = sidebar.style.display === 'block';
    sidebar.style.display = isVisible ? 'none' : 'block';
}

// Search functionality 
function handleSearch(e) { 
    const searchTerm = e.target.value.toLowerCase(); 
    const filteredPosts = SAMPLE_POSTS.filter(post => 
        post.title.toLowerCase().includes(searchTerm) || 
        post.excerpt.toLowerCase().includes(searchTerm) 
    ); 
    renderPosts(filteredPosts); 
} 

// Toggle dark mode 
function toggleDarkMode() { 
    document.body.classList.toggle('dark-mode'); localStorage.setItem('darkMode', document.body.classList.contains('dark-mode')); 
}

// Initialize function 
function init() {
    // handleResize();
    // Event listeners 
    menuToggleBtn?.addEventListener('click', toggleMobileMenu);
    searchInput?.addEventListener('input', handleSearch); 
    darkModeToggle?.addEventListener('click', toggleDarkMode); 
    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; 
    if (savedDarkMode) { document.body.classList.add('dark-mode'); } 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);