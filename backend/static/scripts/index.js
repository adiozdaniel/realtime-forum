// DOM Elements 
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
    document.querySelectorAll('.like-button').forEach(button => { button.addEventListener('click', handleLike); }); 
}

// Search functionality 
function handleSearch(e) { 
    const searchTerm = e.target.value.toLowerCase(); const filteredPosts = SAMPLE_POSTS.filter(post => post.title.toLowerCase().includes(searchTerm) || post.excerpt.toLowerCase().includes(searchTerm) ); renderPosts(filteredPosts); 
} 

// Handle like button click 
function handleLike(e) { 
   const button = e.currentTarget; const likesCount = button.querySelector('.likes-count'); const currentLikes = parseInt(likesCount.textContent); likesCount.textContent = currentLikes + 1; button.classList.add('text-blue-600'); 
} 

// Initialize function 
function init() {
    lucide.createIcons();
    // Initial render 
    renderPosts(); 
} 

// Start the application 
document.addEventListener('DOMContentLoaded', init);