// API Endpoints
window.API_ENDPOINTS = {
	login: "/api/auth/login",
	register: "/api/auth/register",
	logout: "/api/auth/logout",
	posts: "/api/posts",
	comments: "/api/comments",
};

// Constants
window.CONSTANTS = {
	MIN_PASSWORD_LENGTH: 8,
};

// global variables
window.RESDATA = {
	userData: JSON.parse(localStorage.getItem("res")),
	profileImageElement: document.querySelector(".profile-image img"),
};

// DOM Elements
const postsContainer = document.querySelector("#postsContainer");

// Sample Data
const SAMPLE_POSTS = [
	{
		postId: "01",
		title: "Getting Started with Go and Angular",
		author: "Jane Cooper",
		category: "Tutorial",
		likes: 42,
		comments: 12,
		excerpt:
			"Learn how to build a modern web application using Go for the backend and Angular for the frontend...",
		timeAgo: "2h ago",
		hasComments: true,
	},
	{
		postId: "02",
		title: "Best Practices for API Design",
		author: "John Smith",
		category: "Discussion",
		likes: 28,
		comments: 8,
		excerpt:
			"Let's discuss the best practices for designing RESTful APIs that are both scalable and maintainable...",
		timeAgo: "4h ago",
		hasComments: true,
	},
	{
		postId: "03",
		title: "Web Performance Optimization Tips",
		author: "Alice Johnson",
		category: "Guide",
		likes: 35,
		comments: 15,
		excerpt:
			"Essential tips and tricks for optimizing your web application's performance...",
		timeAgo: "6h ago",
		hasComments: true,
	},
];

// Import comment functions
import { loadComments, handleCommentSubmit } from "./comment.js";

/// Modified createPostHTML function to include liked state
function createPostHTML(post) {
	const isLiked = likeState.posts[post.postId]?.likedBy.has("current-user");
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
                    <button class="post-action-button like-button ${
											isLiked ? "liked text-blue-600" : ""
										}" data-post-id="${post.postId}"> 
                        <i data-lucide="thumbs-up"></i>
                        <span class="likes-count">${
													likeState.posts[post.postId]?.count || 0
												}</span> 
                    </button> 
                    <button class="post-action-button comment-toggle" data-post-id="${
											post.postId
										}">
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
	const commentButton = e.target.closest(".comment-toggle");
	if (!commentButton) return;

	const postId = commentButton.dataset.postId;
	const commentsSection = document.querySelector(`#comments-${postId}`);

	if (commentsSection.classList.contains("hidden")) {
		loadComments(postId);
	}

	commentsSection.classList.toggle("hidden");
}

// Render all posts
function renderPosts(posts = SAMPLE_POSTS) {
	if (postsContainer) {
		postsContainer.innerHTML = posts
			.map((post) => createPostHTML(post))
			.join("");
	}

	lucide.createIcons();
	attachPostEventListeners();
}

// Attach event listeners to post buttons
function attachPostEventListeners() {
	document.querySelectorAll(".like-button").forEach((button) => {
		button.addEventListener("click", handleLike);
	});
	document.querySelectorAll(".comment-toggle").forEach((button) => {
		button.addEventListener("click", toggleComments);
	});
	document.querySelectorAll(".comment-form").forEach((form) => {
		form.addEventListener("submit", handleCommentSubmit);
	});
}

// State management for likes
const likeState = {
	posts: {},
	comments: {},
};

// Initialize like state from SAMPLE_POSTS
SAMPLE_POSTS.forEach((post) => {
	likeState.posts[post.postId] = {
		count: post.likes,
		likedBy: new Set(),
		// Track users who liked (can be expanded with user IDs)
	};
});

// Handle like button click for both posts and comments
function handleLike(e) {
	const button = e.currentTarget.closest(".like-button");
	if (!button) return;
	const isComment = button.hasAttribute("data-comment-id");
	// Get the relevant IDs
	const postId = button.getAttribute("data-post-id");
	if (!postId) return;
	const commentId = isComment ? button.getAttribute("data-comment-id") : null;

	// Get the state reference
	const stateRef = isComment
		? (likeState.comments[postId] = likeState.comments[postId] || {})
		: likeState.posts[postId];

	// Initialize comment like state if needed
	if (isComment && !stateRef[commentId]) {
		stateRef[commentId] = {
			count: parseInt(button.querySelector(".likes-count")?.textContent) || 0,
			likedBy: new Set(),
		};
	}
	const likeData = isComment ? stateRef[commentId] : stateRef; //
	if (!likeData) return;
	// Simulate current user ID (in real app, get from auth)
	const currentUserId = "current-user";
	// Toggle like
	if (likeData.likedBy.has(currentUserId)) {
		likeData.count--;
		likeData.likedBy.delete(currentUserId);
		button.classList.remove("liked", "text-blue-600");
	} else {
		likeData.count++;
		likeData.likedBy.add(currentUserId);
		button.classList.add("liked", "text-blue-600");
	}

	// Update UI
	const likesCount = button.querySelector(".likes-count");
	if (likesCount) {
		likesCount.textContent = likeData.count;
	}
	// Animate the like button
	button.classList.add("like-animation");
	setTimeout(() => button.classList.remove("like-animation"), 300);
	// In a real application, you would send this to your backend
	console.log(
		`${isComment ? "Comment" : "Post"} ${
			isComment ? commentId : postId
		} likes:`,
		{ count: likeData.count, isLiked: likeData.likedBy.has(currentUserId) }
	);
}

// Initialize
function init() {
	// Initial render
	renderPosts();
}

// Start the application
document.addEventListener("DOMContentLoaded", init);

export { renderPosts, SAMPLE_POSTS, handleLike };
