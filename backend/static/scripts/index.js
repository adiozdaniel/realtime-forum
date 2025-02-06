// DOM Elements 
const menuToggleBtn = document.querySelector('#menuToggle'); 
const sidebar = document.querySelector('#sidebar'); 
const allCategoriesBtn = document.querySelector('#allCategories');
const categoryDropdown = document.querySelector('#categoryDropdown')
const darkModeToggle = document.querySelector('#darkModeToggle'); 
const postsContainer = document.querySelector('#postsContainer'); 
const searchInput = document.querySelector('#searchInput'); 

// Sample Data 
const SAMPLE_POSTS = [ 
    { 
        postId: "01",
        title: "Getting Started with Go and Angular", 
        author: "Jane Cooper", 
        category: "Tutorial", 
        likes: 42, 
        comments: 12, 
        excerpt: "Learn how to build a modern web application using Go for the backend and Angular for the frontend...", 
        timeAgo: "2h ago" 
    }, 
    { 
        postId: "02",
        title: "Best Practices for API Design", 
        author: "John Smith", 
        category: "Discussion", 
        likes: 28, 
        comments: 8, 
        excerpt: "Let's discuss the best practices for designing RESTful APIs that are both scalable and maintainable...", 
        timeAgo: "4h ago" 

    }, 
    { 
        postId: "03",
        title: "Web Performance Optimization Tips", 
        author: "Alice Johnson", 
        category: "Guide", 
        likes: 35, 
        comments: 15, 
        excerpt: "Essential tips and tricks for optimizing your web application's performance...", 
        timeAgo: "6h ago" 
    } 
]; 

lucide.createIcons();

// Post Template 
function createPostHTML(post) { 
    return ` 
        <article class="post-card"> 
            <div class="flex items-start justify-between"> 
                <div> 
                    <span class="post-category">${post.category}</span> 
                    <h3 class="post-title">${post.title}</h3> 
                    <p class="post-excerpt">${post.excerpt}</p> 
                </div> 
            </div> 
            <div class="post-footer"> 
                <div class="post-actions"> 
                    <button class="post-action-button like-button" data-post-title="${post.title}"> 
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"> 
                            <path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path> 
                        </svg> 
                        <span class="likes-count">${post.likes}</span> 
                    </button> 
                    <button class="post-action-button"> 
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"> 
                            <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path> 
                        </svg> 
                        <span>${post.comments}</span> 
                    </button> 
                </div> 
                
                <div class="post-meta"> 
                    <span>by ${post.author}</span> 
                    <span>•</span> 
                    <span>${post.timeAgo}</span> 
                </div> 
            </div> 
        </article> `; 
    }
    
// Render all posts 
function renderPosts(posts = SAMPLE_POSTS) { 
    postsContainer.innerHTML = posts.map(post => createPostHTML(post)).join(''); attachPostEventListeners(); 
}

// Attach event listeners to post buttons 
function attachPostEventListeners() { 
    document.querySelectorAll('.like-button').forEach(button => { 
        button.addEventListener('click', handleLike); }); 
}

// Toggle mobile menu 
function toggleMobileMenu() { 
    const isVisible = sidebar.style.display === 'block'; sidebar.style.display = isVisible ? 'none' : 'block'; 
} 

// Handle window resize 
function handleResize() { 
    if (window.innerWidth >= 1024) { sidebar.style.display = 'block'; 

    } else { 
        sidebar.style.display = 'none'; 
    } 
}

// Toggle dark mode 
function toggleDarkMode() { 
    document.body.classList.toggle('dark-mode'); 
    localStorage.setItem('darkMode', document.body.classList.contains('dark-mode')); 
}

// Search functionality 
function handleSearch(e) { 
    const searchTerm = e.target.value.toLowerCase(); 
    const filteredPosts = SAMPLE_POSTS.filter(post => 
        post.title.toLowerCase().includes(searchTerm) || 
        post.excerpt.toLowerCase().includes(searchTerm) ); 
        renderPosts(filteredPosts); 
} 

// // Category Dropdown Handling
// function createCategoryDropdown() {
//     const existingDropdown = document.querySelector('#categoryDropdown');
//     if (existingDropdown) return; // Prevent duplicate dropdowns

//     const categories = [...new Set(SAMPLE_POSTS.map(post => post.category))]; // Extract unique categories

//     const dropdown = document.createElement('div');
//     dropdown.classList.add('category-dropdown');
//     dropdown.classList.add('hidden');
//     // dropdown.style.display = 'none'; // Initially hidden

//     categories.forEach(category => {
//         const label = document.createElement('label');
//         label.innerHTML = `<input type="checkbox" value="${category}" class="category-checkbox"> ${category}`;
//         dropdown.appendChild(label);
//     });

//     allCategories.parentElement.appendChild(dropdown); // Append dropdown to the parent of #allCategories
// }

function createCategoryDropdown() { 
    const categories = [...new Set(SAMPLE_POSTS.map(post => post.category))]; 
    categoryDropdown.innerHTML = categories.map(
        category => ` <label> <input type="checkbox" class="category-checkbox" value="${category}"> ${category} </label> `).join(''); 
        
        // Add "All Categories" option 
        const allCategoriesLabel = document.createElement('label'); 
        allCategoriesLabel.innerHTML = ` <input type="checkbox" class="category-checkbox" value="all" checked> All Posts `; 
        categoryDropdown.insertBefore(allCategoriesLabel, categoryDropdown.firstChild); 
} 

// Toggle dropdown 
allCategoriesBtn.addEventListener('click', (e) => { 
    e.stopPropagation(); 
    categoryDropdown.classList.toggle('hidden'); 
}); 

// Close dropdown when clicking outside 
document.addEventListener('click', (e) => { 
    if (!categoryDropdown.contains(e.target) && e.target !== allCategoriesBtn) { 
        categoryDropdown.classList.add('hidden'); 
    } 
}); 

// Handle category selection 
categoryDropdown.addEventListener('change', (e) => { 
    if (e.target.classList.contains('category-checkbox')) { 
        const checkbox = e.target; 
        const checkboxes = document.querySelectorAll('.category-checkbox'); 
        if (checkbox.value === 'all') { 
            // If "All Categories" is selected, uncheck others 
            checkboxes.forEach(cb => { 
                if (cb !== checkbox) cb.checked = false; 
                }); 
        } else { 
            // If a specific category is selected, uncheck "All Categories" 
            const allCategoriesCheckbox = document.querySelector('.category-checkbox[value="all"]'); 
            allCategoriesCheckbox.checked = false; 
        } 
        filterPosts(); 
    } 
});

// Filter posts based on selected categories 
function filterPosts() { 
    const selectedCategories = Array.from(document.querySelectorAll('.category-checkbox:checked')) 
    .map(checkbox => checkbox.value); 
    const isAllSelected = selectedCategories.includes('all'); 
    const filteredPosts = isAllSelected ? SAMPLE_POSTS : SAMPLE_POSTS.filter(post => selectedCategories.includes(post.category)); 
    renderPosts(filteredPosts); 
}

// Handle like button click 
function handleLike(e) { 
   const button = e.currentTarget; const likesCount = button.querySelector('.likes-count'); 
   const currentLikes = parseInt(likesCount.textContent); 
   likesCount.textContent = currentLikes + 1; 
   button.classList.add('text-blue-600'); 
} 

// Initialize function 
function init() {
    lucide.createIcons();

    // Initial render 
    renderPosts(); 

    // Event listeners 
    menuToggleBtn?.addEventListener('click', toggleMobileMenu); 
    darkModeToggle?.addEventListener('click', toggleDarkMode); 
    window.addEventListener('resize', handleResize); 
    searchInput?.addEventListener('input', handleSearch); 

    // categories dropdown
    createCategoryDropdown();

    // Check for saved dark mode preference 
    const savedDarkMode = localStorage.getItem('darkMode') === 'true'; 
    if (savedDarkMode) { 
        document.body.classList.add('dark-mode'); 
    }

    // Initial resize check 
    handleResize(); 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);