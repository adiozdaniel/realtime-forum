// DOM Elements
const menuToggleBtn = document.querySelector('#menuToggle');
const sidebar = document.querySelector('#sidebar');
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
        timeAgo: "2h ago",
        hasComments: true
    },
    {
        postId: "02",
        title: "Best Practices for API Design",
        author: "John Smith",
        category: "Discussion",
        likes: 28,
        comments: 8,
        excerpt: "Let's discuss the best practices for designing RESTful APIs that are both scalable and maintainable...",
        timeAgo: "4h ago",
        hasComments: true
    },
    {
        postId: "03",
        title: "Web Performance Optimization Tips",
        author: "Alice Johnson",
        category: "Guide",
        likes: 35,
        comments: 15,
        excerpt: "Essential tips and tricks for optimizing your web application's performance...",
        timeAgo: "6h ago",
        hasComments: true
    }
];

// Import comment functions
import { loadComments, handleCommentSubmit } from './comment.js'; // Adjust the path as necessary

// Post Template
function createPostHTML(post) {
    return `
        <article class="post-card" data-post-id="${post.postId}">
            <div class="flex items-start justify-between">
                <div>
                    <span class="post-category">${post.category}</span>
                    <h3 class="post-title">${post.title}</h3>
                    <p class="post-excerpt">${post.excerpt}</p>
                </div>
            </div>
            <div class="post-footer">
                <div class="post-actions">
                    <button class="post-action-button like-button" data-post-id="${post.postId}">
                        <i data-lucide="thumbs-up"></i>
                        <span class="likes-count">${post.likes}</span>
                    </button>
                    <button class="post-action-button comment-toggle" data-post-id="${post.postId}">
                        <i data-lucide="message-square"></i>
                        <span class="comments-count">${post.comments}</span>
                    </button>
                </div>
                <div class="post-meta">
                    <span>by ${post.author}</span>
                    <span>•</span>
                    <span>${post.timeAgo}</span>
                </div>
            </div>
            <div class="comments-section hidden" id="comments-${post.postId}">
                <div class="comments-container">
                    <!-- Comments will be inserted here -->
                </div>
                <form class="comment-form" data-post-id="${post.postId}">
                    <textarea placeholder="Write your comment..." class="comment-input"></textarea>
                    <button type="submit" class="comment-submit">Post Comment</button>
                </form>
            </div>
        </article>`;
}

// Toggle comments section
function toggleComments(e) {
    const commentButton = e.target.closest('.comment-toggle');
    if (!commentButton) return;

    const postId = commentButton.dataset.postId;
    const commentsSection = document.querySelector(`#comments-${postId}`);
    
    if (commentsSection.classList.contains('hidden')) {
        loadComments(postId); // Load comments when showing
    }
    
    commentsSection.classList.toggle('hidden');
}

// Render all posts
function renderPosts(posts = SAMPLE_POSTS) {
    postsContainer.innerHTML = posts.map(post => createPostHTML(post)).join('');
    lucide.createIcons();
    attachPostEventListeners();
}

// Attach event listeners to post buttons
function attachPostEventListeners() {
    document.querySelectorAll('.like-button').forEach(button => {
        button.addEventListener('click', handleLike);
    });
    document.querySelectorAll('.comment-toggle').forEach(button => {
        button.addEventListener('click', toggleComments);
    });
    document.querySelectorAll('.comment-form').forEach(form => {
        form.addEventListener('submit', handleCommentSubmit); // Use the imported function
    });
}

// Handle like button click
function handleLike(e) {
    const button = e.currentTarget;
    const likesCount = button.querySelector('.likes-count');
    const currentLikes = parseInt(likesCount.textContent);
    likesCount.textContent = currentLikes + 1;
    button.classList.add('text-blue-600');
}

// Toggle mobile menu
function toggleMobileMenu() {
    const isVisible = sidebar.style.display === 'block';
    sidebar.style.display = isVisible ? 'none' : 'block';
}

// Handle window resize
function handleResize() {
    if (window.innerWidth >= 1024) {
        sidebar.style.display = 'block';
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
        post.excerpt.toLowerCase().includes(searchTerm)
    );
    renderPosts(filteredPosts);
}

// Initialize
function init() {
    renderPosts();
    
    // Event listeners
    menuToggleBtn?.addEventListener('click', toggleMobileMenu);
    darkModeToggle?.addEventListener('click', toggleDarkMode);
    window.addEventListener('resize', handleResize);
    searchInput?.addEventListener('input', handleSearch);
    
    // Check for saved dark mode preference
    const savedDarkMode = localStorage.getItem('darkMode') === 'true';
    if (savedDarkMode) {
        document.body.classList.add('dark-mode');
    }
    
    // Initial resize check
    handleResize();
    
    // Initialize Lucide icons
    lucide.createIcons();
}

// Start the application
document.addEventListener('DOMContentLoaded', init);