import { handleLike } from "./index.js";

const likeState = {
    comments: {}
};

const SAMPLE_COMMENTS = {
    "01": [
        {
            id: 1,
            author: "walter otieno",
            content: "Great tutorial! The step-by-step approach really helped me understand the concepts.",
            timeAgo: "1h ago",
            likes: 5,
            replies: [
                {
                    id: 101,
                    author: "james ochieng",
                    content: "Totally agree! The examples were very clear.",
                    timeAgo: "45m ago",
                    likes: 2
                }
            ]
        },
        {
            id: 2,
            author: "martin shikuku",
            content: "Would love to see a follow-up tutorial on authentication with Go",
            timeAgo: "45m ago",
            likes: 3,
            replies: []
        }
    ],
    "02": [
        {
            id: 3,
            author: "thagruok owino",
            content: "Versioning is crucial for API design. Good point about maintaining backwards compatibility.",
            timeAgo: "2h ago",
            likes: 7,
            replies: []
        }
    ],
    "03": [
        {
            id: 4,
            author: "grace neema",
            content: "The section about lazy loading really improved my site's performance. Thanks!",
            timeAgo: "30m ago",
            likes: 4,
            replies: []
        }
    ]
};

//  HTML for a single reply
function createReplyHTML(reply) {
    return `
        <div class="reply" data-reply-id="${reply.id}">
            <div class="comment-header">
                <span class="comment-author">${reply.author}</span>
                <span class="comment-time">${reply.timeAgo}</span>
            </div>
            <p class="comment-content">${reply.content}</p>
            <div class="comment-actions">
                <span class="comment-action">
                    <i data-lucide="thumbs-up"></i>
                    <span class="likes-count">${reply.likes}</span>
                </span>
            </div>
        </div>`;
}

//  HTML for replies container
function createRepliesHTML(replies) {
    if (!replies || replies.length === 0) return '';
    return `
        <div class="replies-container">
            ${replies.map(reply => createReplyHTML(reply)).join('')}
        </div>`;
}

// // comment HTML with replies
// function createCommentHTML(comment) {
//     return `
//         <div class="comment" data-comment-id="${comment.id}">
//             <div class="comment-header">
//                 <span class="comment-author">${comment.author}</span>
//                 <span class="comment-time">${comment.timeAgo}</span>
//             </div>
//             <p class="comment-content">${comment.content}</p>
//             <div class="comment-actions">
//                 <button class="comment-action like-button" data-comment-id="${comment.id}">
//                     <i data-lucide="thumbs-up"></i>
//                     <span class="likes-count">${comment.likes}</span>
//                 </button>
//                 <span class="comment-action reply-button">
//                     <i data-lucide="reply"></i>
//                     Reply
//                 </span>
//             </div>
//             ${createRepliesHTML(comment.replies)}
//             <div class="reply-form hidden">
//                 <textarea class="reply-input" placeholder="Write your reply..."></textarea>
//                 <div class="reply-form-actions">
//                     <button type="button" class="reply-submit">Reply</button>
//                     <button type="button" class="reply-cancel">Cancel</button>
//                 </div>
//             </div>
//         </div>`;
// }

// Initialize like state from SAMPLE_COMMENTS
Object.keys(SAMPLE_COMMENTS).forEach(postId => {
    SAMPLE_COMMENTS[postId].forEach(comment => {
        likeState.comments[comment.id] = {
            count: comment.likes,
            likedBy: new Set() // Track users who liked
        };

        comment.replies.forEach(reply => {
            likeState.comments[reply.id] = {
                count: reply.likes,
                likedBy: new Set()
            };
        });
    });
});

// Helper function to create comment HTML with like button 
function createCommentHTML(comment, postId) { 
    const commentState = likeState.comments[postId]?.[comment.id]; 
    const isLiked = commentState?.likedBy.has('current-user'); 
    return ` 
    <div class="comment" data-comment-id="${comment.id}"> 
        <div class="comment-content"> 
            <div class="comment-author">${comment.author}</div> 
            <div class="comment-text">${comment.content}</div> 
        </div>
        <div class="comment-footer">
            <div class="comment-actions"> 
                <button class="comment-action-button like-button ${isLiked ? 'liked text-blue-600' : ''}" data-post-id="${postId}" data-comment-id="${comment.id}"> 
                    <i data-lucide="thumbs-up"></i> 
                    <span class="likes-count">${commentState?.count || 0}</span> 
                </button> 
                
            </div>
            <div class="comment-meta">
                <span class="comment-time">${comment.timeAgo}</span> 
            </div>
        </div>
            <!-- reply section will be added here -->
    </div>`; 
    }
// Handle reply button click
function handleReplyClick(e) {
    const replyButton = e.target.closest('.reply-button');
    if (!replyButton) return;

    document.querySelectorAll('.reply-form').forEach(form => form.classList.add('hidden'));

    const comment = replyButton.closest('.comment');
    const replyForm = comment.querySelector('.reply-form');
    replyForm.classList.toggle('hidden');
    replyForm.querySelector('.reply-input').focus();
}


// Handle reply submission
function handleReplySubmit(e) {
    const submitButton = e.target.closest('.reply-submit');
    if (!submitButton) return;

    const replyForm = submitButton.closest('.reply-form');
    const comment = replyForm.closest('.comment');
    const commentId = parseInt(comment.dataset.commentId);
    const postId = comment.closest('.comments-section').id.replace('comments-', '');
    const replyInput = replyForm.querySelector('.reply-input');
    const content = replyInput.value.trim();

    if (!content) return;

    const newReply = {
        id: Date.now(),
        author: "Current User",
        content: content,
        timeAgo: "Just now",
        likes: 0
    };

    // Add reply to the comment data
    const commentIndex = SAMPLE_COMMENTS[postId].findIndex(c => c.id === commentId);
    if (commentIndex === -1) return;

    SAMPLE_COMMENTS[postId][commentIndex].replies.push(newReply);

    // Append new reply to the existing replies section
    const repliesContainer = comment.querySelector('.replies-container');
    if (!repliesContainer) {
        const newRepliesContainer = document.createElement('div');
        newRepliesContainer.classList.add('replies-container');
        comment.appendChild(newRepliesContainer);
        newRepliesContainer.innerHTML = createReplyHTML(newReply);
    } else {
        repliesContainer.innerHTML += createReplyHTML(newReply);
    }

    // Reset input field and hide form
    replyForm.classList.add('hidden');
    replyInput.value = '';
}


// Handle reply cancellation
function handleReplyCancel(e) {
    const cancelButton = e.target.closest('.reply-cancel');
    if (!cancelButton) return;
    
    const replyForm = cancelButton.closest('.reply-form');
    replyForm.classList.add('hidden');
    replyForm.querySelector('.reply-input').value = '';
}

// Handle comment submission
function handleCommentSubmit(e) {
    e.preventDefault();
    const postId = e.target.dataset.postId;
    const input = e.target.querySelector('.comment-input');
    const content = input.value.trim();
    
    if (!content) return;
    
    const newComment = {
        id: Date.now(),
        author: "Current User",
        content: content,
        timeAgo: "Just now",
        likes: 0,
        replies: []
    };
    
    SAMPLE_COMMENTS[postId] = SAMPLE_COMMENTS[postId] || [];
    SAMPLE_COMMENTS[postId].unshift(newComment);
    loadComments(postId);
    input.value = '';
}

// Load comments with replies
function loadComments(postId) {
    const commentsSection = document.querySelector(`#comments-${postId}`);
    if (!commentsSection) return;

    const commentsContainer = commentsSection.querySelector('.comments-container');
    const comments = SAMPLE_COMMENTS[postId] || [];

    commentsContainer.innerHTML = comments.map(comment => createCommentHTML(comment)).join('');

    lucide.createIcons(); // Ensure icons render after loading comments
}

// Add like button event listener
document.addEventListener('click', (e) => {
    document.querySelectorAll('.like-button').forEach(button => {
        button.addEventListener('click', handleLike);
    });
});

// Export the functions
export { loadComments, handleCommentSubmit };