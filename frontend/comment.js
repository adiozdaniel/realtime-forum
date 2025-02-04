// Sample comments data
const SAMPLE_COMMENTS = [
    {
        id: 1,
        author: "Sarah Wilson",
        content: "Great tutorial! The step-by-step approach really helped me understand the concepts.",
        timeAgo: "1h ago",
        likes: 5
    },
    {
        id: 2,
        author: "Mike Chen",
        content: "Would love to see a follow-up tutorial on authentication with Go and Angular.",
        timeAgo: "45m ago",
        likes: 3
    },
    {
        id: 3,
        author: "Emma Davis",
        content: "Versioning is crucial for API design. Good point about maintaining backwards compatibility.",
        timeAgo: "2h ago",
        likes: 7
    },
    {
        id: 4,
        author: "Alex Thompson",
        content: "The section about lazy loading really improved my site's performance. Thanks!",
        timeAgo: "30m ago",
        likes: 4
    }
];

// DOM Elements
const commentsContainer = document.querySelector('.comments-container');
const commentForm = document.querySelector('.comment-form');

// Create comment HTML
function createCommentHTML(comment) {
    return `
        <div class="comment" data-comment-id="${comment.id}">
            <div class="comment-header">
                <span class="comment-author">${comment.author}</span>
                <span class="comment-time">${comment.timeAgo}</span>
            </div>
            <p class="comment-content">${comment.content}</p>
            <div class="comment-actions">
                <span class="comment-action">
                    <i data-lucide="heart"></i>
                    <span class="likes-count">${comment.likes}</span>
                </span>
                <span class="comment-action">
                    <i data-lucide="reply"></i>
                    Reply
                </span>
            </div>
        </div>`;
}

// Render all comments
function renderComments() {
    commentsContainer.innerHTML = SAMPLE_COMMENTS.map(comment => 
        createCommentHTML(comment)
    ).join('');
    lucide.createIcons();
}

// Handle comment submission
function handleCommentSubmit(e) {
    e.preventDefault();
    const input = commentForm.querySelector('.comment-input');
    const content = input.value.trim();
    
    if (!content) return;
    
    const newComment = {
        id: Date.now(),
        author: "Current User",
        content: content,
        timeAgo: "Just now",
        likes: 0
    };
    
    SAMPLE_COMMENTS.unshift(newComment);
    renderComments();
    input.value = '';
}

// Handle comment actions
function handleCommentAction(e) {
    const action = e.target.closest('.comment-action');
    if (!action) return;
    
    const comment = action.closest('.comment');
    const commentId = parseInt(comment.dataset.commentId);
    
    if (action.querySelector('.likes-count')) {
        const likesCount = action.querySelector('.likes-count');
        likesCount.textContent = parseInt(likesCount.textContent) + 1;
    }
}

// Initialize
function init() {
    renderComments();
    commentForm.addEventListener('submit', handleCommentSubmit);
    commentsContainer.addEventListener('click', handleCommentAction);
}

// Start the application
document.addEventListener('DOMContentLoaded', init);
