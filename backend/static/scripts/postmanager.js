import { PostService } from "./postsservice.js";
import { formatTimeAgo } from "./timestamps.js";
import { CommentManager } from "./comment.js";
import {
	postLikeState,
	postDislikeState,
	POSTS,
	COMMENTS,
	recyclebinState,
} from "./data.js";
import { getUserData } from "./authmiddleware.js";
import { PostModalManager } from "./postModalManager.js";
import { toast } from "./toast.js";

const postsContainers = document.querySelectorAll("#postsContainer");

const commentManager = new CommentManager();

class PostManager {
	constructor() {
		this.likeState = postLikeState;
		this.dislikeState = postDislikeState;
		this.postService = new PostService();
		this.postModalManager = new PostModalManager();

		// DOM Elements
		this.form = document.getElementById("createPostForm");
		this.videoLink = document.getElementById("videoLink");
	}
}

PostManager.prototype.createPostHTML = function (post) {
	if (!post) return;

	const isLiked =
		this.likeState.posts[post.post_id]?.likedBy.has("current-user");
	const isDisliked =
		this.dislikeState.posts[post.post_id]?.dislikedBy.has("current-user");

	// Check if the current page is "/dashboard"
	const isDashboard = window.location.pathname === "/dashboard";

	// Determine if the delete and edit buttons should be shown
	const showPostActions = isDashboard && !recyclebinState.RECYCLEBIN;

	return `
      <article class="post-card" data-post-id="${post.post_id}">
	  	${
				post.post_image
					? `
			<div class="post-img">
			<img src="${post.post_image}" alt="Post Image"/>
			</div>`
					: ""
			}
			</div>

			${
				showPostActions
					? `
						<div class="post-user-actions">
							<button class="edit-button" id="postEditBtn" data-post-id="${post.post_id}">
								<i data-lucide="edit"></i>
							</button>
							<button class="delete-button" id="postDeleteBtn" data-post-id="${post.post_id}">
								<i data-lucide="trash-2"></i>
							</button>
						</div>`
					: ""
			}			

        </div>
        <div class="flex items-start justify-between">
          <div>
			<div class="post-categories">
				${post.post_category
					.split(" ") // Split multiple categories
					.map(
						(category) =>
							`<span class="post-category">${category.trim()}</span>`
					)
					.join("")}
			</div>
            <h3 class="post-title">${post.post_title}</h3>
            <p class="post-excerpt">${post.post_content}</p>
        <div class="post-footer">
		${
			isDashboard
				? `<div class="post-actions">
            <div>
              <i data-lucide="thumbs-up"></i>
              <span class="likes-count">${
								this.likeState.posts[post.post_id]?.count || 0
							}</span>
            </div>
			<div>
			<i data-lucide="thumbs-down"></i>
			 <span class="likes-count">${
					this.dislikeState.posts[post.post_id]?.count || 0
				}</span>
			</div>
            <div>
              <i data-lucide="message-square"></i>
              <span class="comments-count">${post.post_comments}</span>
            </div>
          </div>`
				: `<div class="post-actions">
            <button class="post-action-button like-button ${
							isLiked ? "liked text-blue-600" : ""
						}" data-post-id="${post.post_id}">
              <i data-lucide="thumbs-up"></i>
              <span class="likes-count">${
								this.likeState.posts[post.post_id]?.count || 0
							}</span>
            </button>
			<button class="post-action-button dislike-button ${
				isDisliked ? "disliked text-red-600" : ""
			}" data-post-id="${post.post_id}">
			<i data-lucide="thumbs-down"></i>
			 <span class="likes-count">${
					this.dislikeState.posts[post.post_id]?.count || 0
				}</span>
			</button>
            <button class="post-action-button comment-toggle" data-post-id="${
							post.post_id
						}">
              <i data-lucide="message-square"></i>
              <span class="comments-count">${post.post_comments}</span>
            </button>
          </div>`
		}
          
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
        </div>
      </article>
    `;
};

PostManager.prototype.toggleComments = function (e) {
	if (window.location.pathname !== "/") return;

	const commentButton = e.target.closest(".comment-toggle");
	if (!commentButton) return;
	const postId = commentButton.dataset.postId;
	const commentsSection = document.querySelector(`#comments-${postId}`);
	if (commentsSection.classList.contains("hidden")) {
		commentManager.loadComments(postId);
	}
	commentsSection.classList.toggle("hidden");
};

PostManager.prototype.renderPosts = function (posts) {
	if (posts.length === 0) {
		postsContainers.forEach((container) => {
			if (!container.classList.contains("hidden")) {
				container.innerHTML = `<div>Uh-oh! There are no posts yet</div>`;
			}
		});
		return;
	}

	posts.forEach((post) => {
		if (post) {
			post.post_timeAgo = formatTimeAgo(post.created_at);
			post.post_likes = post.likes?.length || 0;
			post.post_dislikes = post.dislikes?.length || 0;
			post.post_comments = post.comments?.length || 0;

			if (post.post_hasComments) {
				COMMENTS[post.post_id] = post.comments;
			}

			post.post_likes = this.likeState.posts[post.post_id] = {
				count: post?.post_likes || 0,
				likedBy: new Set(),
			};

			post.post_dislikes = this.dislikeState.posts[post.post_id] = {
				count: post?.post_dislikes || 0,
				dislikedBy: new Set(),
			};
		}
	});

	const postsHTML = posts.map((post) => this.createPostHTML(post)).join("");
	postsContainers.forEach((container) => {
		if (!container.classList.contains("hidden")) {
			container.innerHTML = postsHTML;
		}

		const postDeleteBtn = container.querySelector("#postDeleteBtn");
		const postEditBtn = container.querySelector("#postEditBtn");
		postDeleteBtn?.addEventListener("click", (e) => this.handlePostDelete(e));
		postEditBtn?.addEventListener("click", (e) => this.handlePostEdit(e));
	});

	this.attachPostEventListeners();
};

PostManager.prototype.handlePostEdit = function (e) {
	e.stopPropagation();

	const button = e.currentTarget.closest("#postEditBtn");
	if (!button) return;
	const postId = button.getAttribute("data-post-id");

	const post = POSTS.find((post) => post.post_id === postId);

	this.postModalManager.openModal(post);
};

PostManager.prototype.handlePostDelete = async function (e) {
	e.preventDefault();
	e.stopPropagation();

	const button = e.currentTarget.closest("#postDeleteBtn");
	if (!button) return;
	const postId = button.getAttribute("data-post-id");

	const deletePost = confirm("Delete this post?");
	if (!deletePost) return;

	const postData = {
		post_id: postId,
	};

	const res = await this.postService.deletePost(postData);
	if (res.error) {
		toast.createToast("error", res.message);
		return;
	}

	const postIndex = POSTS.findIndex((post) => post.post_id === postId);
	if (postIndex !== -1) {
		POSTS.splice(postIndex, 1);
	}

	toast.createToast("success", "Post deleted successfully!");
};

PostManager.prototype.handleSubmit = async function (e) {
	e.preventDefault();

	let formData = {};

	if (recyclebinState.TEMP_DATA !== null) {
		formData = {
			PostTitle: document.getElementById("postTitle").value,
			PostCategory: Array.from(
				document.querySelectorAll('input[name="postCategory"]:checked')
			)
				.map((checkbox) => checkbox.value)
				.join(" "),
			PostContent: document.getElementById("postContent").value,
			PostVideo: this.videoLink.value || null,
			PostImage: recyclebinState.TEMP_DATA.img || null,
			PostID: recyclebinState.TEMP_DATA.post_id || null,
		};
	} else {
		formData = {
			PostTitle: document.getElementById("postTitle").value,
			PostCategory: Array.from(
				document.querySelectorAll('input[name="postCategory"]:checked')
			)
				.map((checkbox) => checkbox.value)
				.join(" "),
			PostContent: document.getElementById("postContent").value,
			PostVideo: this.videoLink.value || null,
		};
	}

	const res = await this.postService.createPost(formData);
	if (res.error) {
		toast.createToast("error", res.message);
		return;
	}

	if (res.data) {
		toast.createToast(
			"success",
			res.data.post_title + " created successfully!"
		);
		this.postModalManager.closeModal();
		POSTS.unshift(res.data);
		this.renderPosts(POSTS);
	}
};

PostManager.prototype.attachPostEventListeners = function () {
	lucide.createIcons();
	document.querySelectorAll(".like-button").forEach((button) => {
		button.addEventListener("click", (e) => this.handlePostLikes(e));
	});
	document.querySelectorAll(".dislike-button").forEach((button) => {
		button.addEventListener("click", (e) => this.handlePostDisLikes(e));
	});
	document.querySelectorAll(".comment-toggle").forEach((button) => {
		button.addEventListener("click", (e) => this.toggleComments(e));
	});
	document
		.getElementById("createPostForm")
		?.addEventListener("submit", (e) => this.handleSubmit(e));
};

PostManager.prototype.handlePostLikes = async function (e) {
	e.stopPropagation();

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

PostManager.prototype.handlePostDisLikes = async function (e) {
	e.stopPropagation();

	const button = e.currentTarget.closest(".dislike-button");
	if (!button) return;
	const postId = button.getAttribute("data-post-id");
	if (!postId) return;

	const dislikeData = (this.dislikeState.posts[postId] ??= {
		count: 0,
		dislikedBy: new Set(),
	});

	const currentUser = await getUserData();
	if (!currentUser?.user_id) {
		alert("Please login to dislike the post");
		window.location.href = "/auth";
		return;
	}

	const postData = { post_id: postId, user_id: currentUser.user_id };
	const res = await this.postService.dislikePost(postData);

	if (res.error) {
		alert(res.message);
		return;
	}

	if (res.data) {
		dislikeData.count++;
		dislikeData.dislikedBy.add(currentUser.user_id);
		button.classList.add("disliked", "text-red-600");
	} else if (res.data === null) {
		dislikeData.count = Math.max(0, dislikeData.count - 1);
		dislikeData.dislikedBy.delete(currentUser.user_id);
		button.classList.remove("disliked", "text-red-600");
	}

	const dislikesCount = button.querySelector(".likes-count");
	if (dislikesCount) {
		dislikesCount.textContent = dislikeData.count;
	}

	button.classList.add("dislike-animation");
	setTimeout(() => button.classList.remove("dislike-animation"), 300);
};

PostManager.prototype.searchPosts = function (
	searchTerm = "",
	selectedCategories = []
) {
	const term = searchTerm.toLowerCase();
	const isAllSelected = selectedCategories.includes("all");

	// Filter posts based on search term
	if (searchTerm !== "") {
		const filteredBySearch = POSTS.filter(
			(post) =>
				!term ||
				post.post_title.toLowerCase().includes(term) ||
				post.post_category.toLowerCase().includes(term)
		);

		this.renderPosts(filteredBySearch);
		return;
	} else if (selectedCategories) {
		// Filter posts based on categories
		const filteredByCategory = isAllSelected
			? POSTS
			: POSTS.filter((post) =>
					selectedCategories.some((category) =>
						post.post_category.toLowerCase().includes(category.toLowerCase())
					)
			  );

		// Merge results without duplicates
		this.renderPosts(filteredByCategory);
		return;
	}
	this.renderPosts(POSTS);
};

PostManager.prototype.init = async function () {
	const posts = await this.postService.fetchPosts();
	this.postList = Array.isArray(posts) ? posts : posts.data;

	if (this.postList === null) return;

	this.postList.forEach((post) => POSTS.unshift(post));
	this.renderPosts(POSTS);
};

document.addEventListener("DOMContentLoaded", () => {
	const postManager = new PostManager();
	postManager.init();
	// postModal.init();
	postManager.postModalManager.init();
});

export { PostManager };
