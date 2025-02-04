// Sample comments data
const SAMPLE_COMMENTS = {
    "01": [
        {
            id: 1,
            author: "Shayo victor",
            content: "Great tutorial! The step-by-step approach really helped me understand the concepts.",
            timeAgo: "1h ago",
            likes: 5
        },
        {
            id: 2,
            author: "victor shayo",
            content: "Would love to see a follow-up tutorial on authentication with Go ",
            timeAgo: "45m ago",
            likes: 3
        }
    ],
    "02": [
        {
            id: 3,
            author: "thagruok owino",
            content: "Versioning is crucial for API design. Good point about maintaining backwards compatibility.",
            timeAgo: "2h ago",
            likes: 7
        }
    ],
    "03": [
        {
            id: 4,
            author: "kanyo kanyo",
            content: "The section about lazy loading really improved my site's performance. Thanks!",
            timeAgo: "30m ago",
            likes: 4
        }
    ]
};

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

// Load comments for a post
function loadComments(postId) {
    const commentsSection = document.querySelector(`#comments-${postId}`);
    if (!commentsSection) return;

    const commentsContainer = commentsSection.querySelector('.comments-container');
    const comments = SAMPLE_COMMENTS[postId] || [];
    
    commentsContainer.innerHTML = comments.map(comment => createCommentHTML(comment)).join('');
    lucide.createIcons();
}

// Handle comment submission
function handleCommentSubmit(e) {
    e.preventDefault();
    const postId = e.target.dataset.postId; // Get post ID from the form
    const input = e.target.querySelector('.comment-input');
    const content = input.value.trim();
    
    if (!content) return;
    
    const newComment = {
        id: Date.now(),
        author: "Current User",
        content: content,
        timeAgo: "Just now",
        likes: 0
    };
    
    SAMPLE_COMMENTS[postId] = SAMPLE_COMMENTS[postId] || [];
    SAMPLE_COMMENTS[postId].unshift(newComment); // Add new comment to the sample comments
    loadComments(postId); // Reload comments for the post
    input.value = '';
}

// Export functions for use in index.js
export { loadComments, handleCommentSubmit };