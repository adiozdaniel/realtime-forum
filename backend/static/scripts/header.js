import  {renderPosts, SAMPLE_POSTS} from "./index.js";
// DOM Elements 
const searchInput = document.querySelector('#searchInput'); 
const darkModeToggle = document.querySelector('#darkModeToggle'); 

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
    // Event listeners 
    searchInput?.addEventListener('input', handleSearch); 
    darkModeToggle?.addEventListener('click', toggleDarkMode); 
    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; 
    if (savedDarkMode) { document.body.classList.add('dark-mode'); } 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);