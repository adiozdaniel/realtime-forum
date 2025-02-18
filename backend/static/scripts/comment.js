import { CommentService } from "./commentservice.js";
import { formatTimeAgo } from "./timestamps.js";
import { getUserData } from "./authmiddleware.js";
import { SAMPLE_COMMENTS } from "./data.js";

class CommentManager {
	constructor() {
		this.likeState = { comments: {} };
		this.commentService = new CommentService();
	}
}

CommentManager.prototype.createReplyHTML = function (reply) {
	const replyState = this.likeState.comments[reply.id] || {
		count: 0,
		likedBy: new Set(),
	};
	const isLiked = replyState.likedBy.has("current-user");
	return `
            <div class="reply" data-reply-id="${reply.id}">
                <div class="comment-header">
                    <span class="comment-author">${reply.author}</span>
                </div>
                <p class="comment-content">${reply.content}</p>
                <div class="comment-footer">
                    <div class="comment-actions">
                        <button class="comment-action-button like-button ${
													isLiked ? "liked text-blue-600" : ""
												}" 
                            data-comment-id="${reply.id}">
                            <i data-lucide="thumbs-up"></i>
                            <span class="likes-count">${replyState.count}</span>
                        </button>
                    </div>
                    <div class="comment-meta">
                        <span class="comment-time">${reply.timeAgo}</span>
                    </div>
                </div>
            </div>`;
};

CommentManager.prototype.createCommentHTML = function (comment, postId) {
	const commentState = this.likeState.comments[comment.id] || {
		count: 0,
		likedBy: new Set(),
	};
	const isLiked = commentState.likedBy.has("current-user");
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
                        <span class="likes-count">${commentState.count}</span> 
                    </button>
                    <button class="comment-action-button reply-button">
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
        </div>`;
};

CommentManager.prototype.loadComments = function (postId) {
	const commentsSection = document.querySelector(`#comments-${postId}`);
	if (!commentsSection) return;

	// Generate comment HTML without form
	const comment = SAMPLE_COMMENTS[postId] || [];
	const commentsHTML = comment
		.map((comment) => this.createCommentHTML(comment, postId))
		.join("");

	// Ensure only one form per post
	const formHTML = `
        <form class="comment-form" data-post-id="${postId}">
            <textarea placeholder="Write your comment..." class="comment-input"></textarea>
            <button type="submit" class="comment-submit">Post Comment</button>
        </form>`;

	// Set innerHTML with comments and single form
	commentsSection.innerHTML = commentsHTML + formHTML;

	// Attach event listener to the form
	const form = commentsSection.querySelector(".comment-form");
	if (form) {
		form.addEventListener("submit", (e) => this.handleCommentSubmit(e));
	}
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
	SAMPLE_COMMENTS[postId].push(commentsRes.data);
	this.loadComments(postId);
	input.value = "";
};

CommentManager.prototype.attachEventListeners = function () {
	document.querySelectorAll(".comments-section").forEach((section) => {
		const postId = section.id.replace("comments-", "");
		this.loadComments(postId);
	});
};

CommentManager.prototype.initLikeState = function (sampleComments) {
	Object.keys(sampleComments).forEach((postId) => {
		sampleComments[postId].forEach((comment) => {
			this.likeState.comments[comment.id] = {
				count: comment.likes,
				likedBy: new Set(),
			};
			comment.replies.forEach((reply) => {
				this.likeState.comments[reply.id] = {
					count: reply.likes,
					likedBy: new Set(),
				};
			});
		});
	});
};

const commentManager = new CommentManager();

window.addEventListener("DOMContentLoaded", () => {
	commentManager.initLikeState(SAMPLE_COMMENTS);
	commentManager.attachEventListeners();
});

export { CommentManager };
