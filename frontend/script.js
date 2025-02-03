// DOM Elements 
const menuToggleBtn = document.querySelector('#menuToggle'); 
const sidebar = document.querySelector('#sidebar'); 
const darkModeToggle = document.querySelector('#darkModeToggle'); 
const postsContainer = document.querySelector('#postsContainer'); 
const searchInput = document.querySelector('#searchInput'); 

// Sample Data 
const SAMPLE_POSTS = [ 
    { 
        title: "Getting Started with Go and Angular", 
        author: "Jane Cooper", 
        category: "Tutorial", 
        likes: 42, 
        comments: 12, 
        excerpt: "Learn how to build a modern web application using Go for the backend and Angular for the frontend...", 
        timeAgo: "2h ago" 
    }, 
    { 
        title: "Best Practices for API Design", 
        author: "John Smith", 
        category: "Discussion", 
        likes: 28, 
        comments: 8, 
        excerpt: "Let's discuss the best practices for designing RESTful APIs that are both scalable and maintainable...", 
        timeAgo: "4h ago" 

    }, 
    { 
        title: "Web Performance Optimization Tips", 
        author: "Alice Johnson", 
        category: "Guide", 
        likes: 35, 
        comments: 15, 
        excerpt: "Essential tips and tricks for optimizing your web application's performance...", 
        timeAgo: "6h ago" 
    } 
]; 
