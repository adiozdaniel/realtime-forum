import { CommentService } from "./commentservice.js";
import { formatTimeAgo } from "./timestamps.js";
import { getUserData } from "./authmiddleware.js";
import { SAMPLE_COMMENTS, commentLikeState } from "./data.js";

class CommentManager {
	constructor() {
		this.likeState = commentLikeState;
		this.commentService = new CommentService();
	}
}

CommentManager.prototype.createReplyHTML = function (reply) {
	const replyState = (this.likeState.comments[reply.id] ??= {
		count: 0,
		likedBy: new Set(),
	});
	const isLiked = replyState.likedBy.has("current-user");
	return `
            <div class="reply" data-reply-id="${reply.id}">
			<div class="comment-content"> 
                <div class="profile-image">
                    <img src="${reply.author_img}" 
                         onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
                </div>
                <div>
                    <div class="comment-author">${reply.user_name}</div> 
                    <div class="comment-text">${reply.content}</div> 
                </div>
            </div>
                <div class="comment-footer">
                    <div class="comment-actions">
                        <button class="comment-action-button like-button ${isLiked ? "liked text-blue-600" : ""
		}" 
                            data-comment-id="${reply.id}">
                            <i data-lucide="thumbs-up"></i>
                            <span class="likes-count">${replyState.count}</span>
                        </button>
                    </div>
                    <div class="comment-meta">
                        <span class="reply-time">${formatTimeAgo(reply.updated_at)}</span>
                    </div>
                </div>
            </div>`;
};

CommentManager.prototype.showReplyForm = async function (e) {
	const button = e.target.closest(".reply-button");
	if (!button) return;

	const postId = button.getAttribute("data-post-id");
	const commentId = button.getAttribute("data-comment-id");

	const commentElement = document.querySelector(
		`.comment[data-comment-id="${commentId}"]`
	);
	if (!commentElement) return;

	const userData = await getUserData();
	if (!userData) {
		alert("Please login to comment");
		window.location.href = "/auth";
		return;
	}

	// Check if reply form already exists
	const existingReplyForm = commentElement.querySelector(".reply-form");

	if (existingReplyForm) {
		existingReplyForm.remove();
	} else {
		const replyFormHTML = `
            <form class="reply-form" data-comment-id="${commentId}" data-post-id="${postId}">
                <textarea placeholder="Write your reply..." class="reply-input"></textarea>
                <button type="submit" class="reply-submit">Reply</button>
            </form>`;

		commentElement.insertAdjacentHTML("beforeend", replyFormHTML);

		// Attach event listener to the new reply form if needed
		const replyForm = commentElement.querySelector(".reply-form");
		if (replyForm) {
			replyForm.addEventListener("submit", (e) => this.handleReplySubmit(e));
		}
	}
};

CommentManager.prototype.handleReplySubmit = async function (e) {
	e.preventDefault();

	const userData = await getUserData();
	if (!userData) {
		alert("Please login to comment");
		window.location.href = "/auth";
		return;
	}

	const form = e.target.closest(".reply-form");
	if (!form) return;

	const postId = form.getAttribute("data-post-id");
	const commentId = form.getAttribute("data-comment-id");
	const replyText = form.querySelector(".reply-input").value.trim();

	if (!replyText) {
		alert("Reply cannot be empty!");
		return;
	}

	const replyData = {
		comment_id: commentId,
		user_id: userData.user_id,
		user_name: userData.user_name,
		content: replyText,
	};

	const res = await this.commentService.createReply(replyData);
	if (res.error) {
		alert(res.message);
		return;
	}

	// Find the correct comment and insert the reply
	const commentElement = document.querySelector(
		`.comment[data-comment-id="${commentId}"]`
	);
	if (commentElement) {
		console.log(res.data);
		const replyHTML = this.createReplyHTML(res.data);
		const repliesContainer = commentElement.querySelector(".replies-container");
		if (repliesContainer) {
			repliesContainer.insertAdjacentHTML("beforeend", replyHTML);
		} else {
			commentElement.insertAdjacentHTML(
				"beforeend",
				`<div class="replies-container">${replyHTML}</div>`
			);
		}
	}

	// Update local data
	SAMPLE_COMMENTS[postId].forEach((comment) => {
		if (comment.comment_id === commentId) {
			comment.replies.push(res.data);
		}
	});

	form.remove();
	lucide.createIcons();
};

CommentManager.prototype.createCommentHTML = function (comment, postId) {
	const isLiked =
		this.likeState.comments[comment.comment_id]?.likedBy.has("current-user");

	// Generate replies HTML
	const repliesHTML = comment.replies
		.map((reply) => this.createReplyHTML(reply))
		.join("");

	return `
        <div class="comment" data-comment-id="${comment.comment_id}"> 
            <div class="comment-content"> 
                <div class="profile-image">
                    <img src="${comment.author_img}" 
                         onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
                </div>
                <div>
                    <div class="comment-author">${comment.user_name}</div> 
                    <div class="comment-text">${comment.comment}</div> 
                </div>
            </div>
            <div class="comment-footer">
                <div class="comment-actions"> 
                    <button class="comment-action-button like-button ${
											isLiked ? "liked text-blue-600" : ""
										}" 
                        data-post-id="${postId}" data-comment-id="${
		comment.comment_id
	}"> 
                        <i data-lucide="thumbs-up"></i> 
                        <span class="likes-count">${
													this.likeState.comments[comment.comment_id]?.count ||
													0
												}</span> 
                    </button>
                    <button 
											class="comment-action-button reply-button"
											data-post-id="${postId}"
											data-comment-id="${comment.comment_id}"
											>
                        <i data-lucide="reply"></i> 
                        <span>Reply</span>
                    </button>  
                </div>
                <div class="comment-meta">
                    <span class="comment-time">${formatTimeAgo(
											comment.created_at
										)}</span> 
                </div>
            </div>
            <div class="replies-container">${repliesHTML}</div>
        </div>`;
};

CommentManager.prototype.loadComments = function (postId) {
	const commentsSection = document.querySelector(`#comments-${postId}`);
	if (!commentsSection) return;

	// Fetch comments including replies
	const comments = SAMPLE_COMMENTS[postId] || [];
	const commentsHTML = comments
		.map((comment) => this.createCommentHTML(comment, postId))
		.join("");

	// Ensure only one comment form per post
	const formHTML = `
        <form class="comment-form" data-post-id="${postId}">
            <textarea placeholder="Write your comment..." class="comment-input"></textarea>
            <button type="submit" class="comment-submit">Post Comment</button>
        </form>`;

	// Set innerHTML with comments and form
	commentsSection.innerHTML = commentsHTML + formHTML;

	// Attach event listener to form
	const form = commentsSection.querySelector(".comment-form");
	if (form) {
		form.addEventListener("submit", (e) => this.handleCommentSubmit(e));
	}

	// Attach event listeners for likes and replies
	document.querySelectorAll(".comments-section").forEach((section) => {
		section.addEventListener("click", (event) => {
			const likeButton = event.target.closest(".like-button");
			const replyButton = event.target.closest(".reply-button");

			if (likeButton && likeButton.dataset.commentId) {
				this.handleCommentLikes(event);
			}

			if (replyButton && replyButton.dataset.commentId) {
				this.showReplyForm(event);
			}
		});
	});

	lucide.createIcons();
};

CommentManager.prototype.handleCommentSubmit = async function (e) {
	e.preventDefault();
	const postId = e.target.dataset.postId;
	const input = e.target.querySelector(".comment-input");
	const content = input.value.trim();
	if (!content) return;

	const userData = await getUserData();
	if (!userData) {
		alert("Please login to comment");
		window.location.href = "/auth";
		return;
	}

	let newComment = {
		post_id: postId,
		comment: content,
		likes: 0,
		replies: [],
	};
	const commentsRes = await this.commentService.createComment(newComment);
	if (commentsRes.error) {
		alert(commentsRes.message);
		return;
	}

	if (!SAMPLE_COMMENTS[postId]) SAMPLE_COMMENTS[postId] = [];

	SAMPLE_COMMENTS[postId].push(commentsRes.data);
	this.loadComments(postId);
	input.value = "";
};

CommentManager.prototype.handleCommentLikes = async function (event) {
	const likeButton = event.target.closest(".like-button");
	if (!likeButton) return;

	const commentId = likeButton.dataset.commentId;
	const postId = likeButton.dataset.postId;

	if (!commentId || !postId) return;

	const likeData = (this.likeState.comments[commentId] ??= {
		count: 0,
		likedBy: new Set(),
	});

	// Get current user (this should be an authenticated user)
	const currentUser = await getUserData();
	if (!currentUser?.user_id) {
		alert("Please log in to like comments.");
		window.location.href = "/auth";
		return;
	}

	const commentData = {
		user_id: currentUser.user_id,
		post_id: postId,
		comment_id: commentId,
	};

	// Send request to backend
	const res = await this.commentService.likeComment(commentData);

	if (res.error) {
		alert(res.message);
		return;
	}

	if (res.data) {
		likeData.count++;
		likeData.likedBy.add(currentUser.user_id);
		likeButton.classList.add("liked", "text-blue-600");
	} else if (res.data === null) {
		likeData.count = Math.max(0, likeData.count - 1);
		likeData.likedBy.delete(currentUser.user_id);
		likeButton.classList.remove("liked", "text-blue-600");
	}

	const likesCount = likeButton.querySelector(".likes-count");

	if (likesCount) likesCount.textContent = likeData.count;

	likeButton.classList.add("like-animation");
	setTimeout(() => likeButton.classList.remove("like-animation"), 300);
};

CommentManager.prototype.attachEventListeners = function () {
	document.querySelectorAll(".comments-section").forEach((section) => {
		const postId = section.id.replace("comments-", "");
		this.loadComments(postId);
	});
};

CommentManager.prototype.initLikeState = function (comments) {
	if (!comments) return;

	const commentList = Array.isArray(comments) ? comments : [];
	if (commentList === null) return;

	Object.keys(comments).forEach((post_id) => {
		comments[post_id]?.forEach((comments) => {
			this.likeState.comments[comments.comment_id] = {
				count: comments.likes?.length || 0,
				likedBy: new Set(),
			};

			// comments.replies?.forEach((reply) => {
			// 	this.likeState.comments[reply.reply_id] = {
			// 		count: reply.likes.length || 0,
			// 		likedBy: new Set(),
			// 	};
			// });
		});
	});
};

const commentManager = new CommentManager();

window.addEventListener("DOMContentLoaded", () => {
	setTimeout(() => {
		commentManager.initLikeState(SAMPLE_COMMENTS);
		commentManager.attachEventListeners();
	}, 300);
});

export { CommentManager };
