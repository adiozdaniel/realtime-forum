import { PostService } from "./postsservice.js";
import { formatTimeAgo } from "./timestamps.js";
// Import comment functions
import { handleCommentSubmit, loadComments } from "./comment.js";
import { API_ENDPOINTS, SAMPLE_POSTS, SAMPLE_COMMENTS } from "./data.js";

// DOM Elements
const postsContainer = document.querySelector("#postsContainer");

// PostManager class encapsulates post and comment management
class PostManager {
	constructor() {
		this.likeState = {
			posts: {},
			comments: {},
		};
		this.postService = new PostService();
	}

	// Create post HTML
	createPostHTML(post) {
		const isLiked =
			this.likeState.posts[post.post_id]?.likedBy.has("current-user");
		return `
      <article class="post-card" data-post-id="${post.post_id}">
        <div class="flex items-start justify-between">
          <div>
            <span class="post-category">${post.post_category}</span>
            <h3 class="post-title">${post.post_title}</h3>
            <p class="post-excerpt">${post.post_content}</p>
          </div>
        </div>
        <div class="post-footer">
          <div class="post-actions">
            <button class="post-action-button like-button ${
							isLiked ? "liked text-blue-600" : ""
						}" data-post-id="${post.post_id}">
              <i data-lucide="thumbs-up"></i>
              <span class="likes-count">${
								this.likeState.posts[post.post_id]?.count || 0
							}</span>
            </button>
            <button class="post-action-button comment-toggle" data-post-id="${
							post.post_id
						}">
              <i data-lucide="message-square"></i>
              <span class="comments-count">${post.post_comments}</span>
            </button>
          </div>
          <div class="post-meta">
            <span>by ${post.post_author}</span>
            <span>•</span>
            <span>${post.post_timeAgo}</span>
          </div>
        </div>
        <div class="comments-section hidden" id="comments-${post.post_id}">
          <div class="comments-container"></div>
          <form class="comment-form" data-post-id="${post.post_id}">
            <textarea placeholder="Write your comment..." class="comment-input"></textarea>
            <button type="submit" class="comment-submit">Post Comment</button>
          </form>
        </div>
      </article>
    `;
	}

	// Toggle comments section visibility
	toggleComments(e) {
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
	renderPosts(posts = SAMPLE_POSTS) {
		postsContainer.innerHTML = posts
			.map((post) => this.createPostHTML(post))
			.join("");
		lucide.createIcons();

		this.attachPostEventListeners();
	}

	// Attach event listeners to post buttons
	attachPostEventListeners() {
		document.querySelectorAll(".like-button").forEach((button) => {
			button.addEventListener("click", this.handleLike.bind(this));
		});
		document.querySelectorAll(".comment-toggle").forEach((button) => {
			button.addEventListener("click", this.toggleComments.bind(this));
		});
		document.querySelectorAll(".comment-form").forEach((form) => {
			form.addEventListener("submit", handleCommentSubmit);
		});
	}

	// Handle like button click for both posts and comments
	async handleLike(e) {
		const button = e.currentTarget.closest(".like-button");
		if (!button) return;

		const postId = button.getAttribute("data-post-id");
		if (!postId) return;

		const isComment = button.hasAttribute("data-comment-id");
		const commentId = isComment ? button.getAttribute("data-comment-id") : null;

		// Retrieve or initialize like state
		const stateRef = isComment
			? (this.likeState.comments[postId] ??= {})
			: (this.likeState.posts[postId] ??= { count: 0, likedBy: new Set() });

		const likeData = isComment
			? (stateRef[commentId] ??= { count: 0, likedBy: new Set() })
			: stateRef;

		// Ensure user is logged in
		const currentUser = this.userData();
		if (!currentUser?.user_id) {
			alert("Please login to like the post");
			window.location.href = "/auth";
			return;
		}

		const postData = { post_id: postId, user_id: currentUser.user_id };
		// const isLiked = likeData.likedBy.has(currentUser.user_id);

		// Call API to toggle like
		const res = await this.postService.likePost(postData);
		if (res.error) {
			alert(res.message);
			return;
		}

		let isLiked = false;

		if (res.data) {
			// Like was added
			likeData.count++;
			likeData.likedBy.add(currentUser.user_id);
			button.classList.add("liked", "text-blue-600");
			isLiked = true;
		} else if (res.data === null) {
			// Like was removed
			likeData.count = Math.max(0, likeData.count - 1);
			likeData.likedBy.delete(currentUser.user_id);
			button.classList.remove("liked", "text-blue-600");
			isLiked = false;
		}

		// Update UI
		const likesCount = button.querySelector(".likes-count");
		if (likesCount) {
			likesCount.textContent = likeData.count;
		}

		button.classList.add("like-animation");
		setTimeout(() => button.classList.remove("like-animation"), 300);
	}

	userData() {
		try {
			const user = localStorage.getItem("userdata");
			return user ? JSON.parse(user) : null;
		} catch (error) {
			console.error("Error retrieving user data:", error);
			return null;
		}
	}

	// Initialize the application
	async init() {
		const posts = await this.postService.fetchPosts();
		const postList = Array.isArray(posts) ? posts : posts.data;

		postList.forEach((post) => SAMPLE_POSTS.push(post));

		// fetchComments();

		SAMPLE_POSTS.forEach((post) => {
			post.post_timeAgo = formatTimeAgo(post.created_at);
			post.post_likes = post?.post_likes || post.likes?.length;
			post.post_comments = post?.post_comments || post.comments?.length || 0;

			if (post.post_hasComments) {
				SAMPLE_COMMENTS[post.post_id] = post.comments;
			}

			post.post_likes = this.likeState.posts[post.post_id] = {
				count: post?.post_likes || 0,
				likedBy: new Set(),
			};
		});

		console.log(SAMPLE_POSTS);
		console.log(SAMPLE_COMMENTS);

		if (postsContainer) this.renderPosts();
	}
}

const postManager = new PostManager();
// Initialize
document.addEventListener("DOMContentLoaded", () => {
	// const postManager = new PostManager();
	postManager.init();
});

export { postManager, SAMPLE_POSTS };
