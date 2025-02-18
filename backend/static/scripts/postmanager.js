import { PostService } from "./postsservice.js";
import { formatTimeAgo } from "./timestamps.js";
import { handleCommentSubmit, loadComments } from "./comment.js";
import { postLikeState, SAMPLE_POSTS, SAMPLE_COMMENTS } from "./data.js";
import { getUserData } from "./authmiddleware.js";
import { PostModalManager } from "./createposts.js";

const postsContainer = document.querySelector("#postsContainer");

class PostManager {
	constructor() {
		this.likeState = postLikeState;
		this.postService = new PostService();
	}
}

PostManager.prototype.createPostHTML = function (post) {
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
		   <div class="profile-image">
		  		<img src=${
						post.author_img
					} onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
			</div>
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
};

PostManager.prototype.toggleComments = function (e) {
	const commentButton = e.target.closest(".comment-toggle");
	if (!commentButton) return;
	const postId = commentButton.dataset.postId;
	const commentsSection = document.querySelector(`#comments-${postId}`);
	if (commentsSection.classList.contains("hidden")) {
		loadComments(postId);
	}
	commentsSection.classList.toggle("hidden");
};

PostManager.prototype.renderPosts = function (posts = SAMPLE_POSTS) {
	postsContainer.innerHTML = posts
		.map((post) => this.createPostHTML(post))
		.join("");
	this.attachPostEventListeners();
};

PostManager.prototype.attachPostEventListeners = function () {
	lucide.createIcons();
	document.querySelectorAll(".like-button").forEach((button) => {
		button.addEventListener("click", (e) => this.handlePostLikes(e));
	});
	document.querySelectorAll(".comment-toggle").forEach((button) => {
		button.addEventListener("click", (e) => this.toggleComments(e));
	});
	document.querySelectorAll(".comment-form").forEach((form) => {
		form.addEventListener("submit", (e) => handleCommentSubmit(e));
	});
};

PostManager.prototype.handlePostLikes = async function (e) {
	const button = e.currentTarget.closest(".like-button");
	if (!button) return;
	const postId = button.getAttribute("data-post-id");
	if (!postId) return;
	const likeData = (this.likeState.posts[postId] ??= {
		count: 0,
		likedBy: new Set(),
	});
	const currentUser = await getUserData();
	if (!currentUser?.user_id) {
		alert("Please login to like the post");
		window.location.href = "/auth";
		return;
	}
	const postData = { post_id: postId, user_id: currentUser.user_id };
	const res = await this.postService.likePost(postData);
	if (res.error) {
		alert(res.message);
		return;
	}
	if (res.data) {
		likeData.count++;
		likeData.likedBy.add(currentUser.user_id);
		button.classList.add("liked", "text-blue-600");
	} else if (res.data === null) {
		likeData.count = Math.max(0, likeData.count - 1);
		likeData.likedBy.delete(currentUser.user_id);
		button.classList.remove("liked", "text-blue-600");
	}
	const likesCount = button.querySelector(".likes-count");
	if (likesCount) {
		likesCount.textContent = likeData.count;
	}
	button.classList.add("like-animation");
	setTimeout(() => button.classList.remove("like-animation"), 300);
};

PostManager.prototype.init = async function () {
	const posts = await this.postService.fetchPosts();
	const postList = Array.isArray(posts) ? posts : posts.data;

	if (postList === null) return;

	postList.forEach((post) => SAMPLE_POSTS.push(post));
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
	if (postsContainer) this.renderPosts();
};

const postManager = new PostManager();
document.addEventListener("DOMContentLoaded", () => {
	postManager.init();

	const postModal = new PostModalManager();
	postModal.init();
});

export { postManager, SAMPLE_POSTS };
