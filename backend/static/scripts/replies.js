import { formatTimeAgo } from "./timestamps.js";
import { getUserData } from "./authmiddleware.js";
import { commentLikeState, REPLIES } from "./data.js";
import { CommentService } from "./commentservice.js";
import { toast } from "./toast.js";

class ReplyManager {
	constructor() {
		this.commentService = new CommentService();
		this.replyCall = 0;
	}
}

ReplyManager.prototype.createReplyHTML = function (reply) {
	const replyState = (commentLikeState.replies[reply.reply_id] ??= {
		count: 0,
		likedBy: new Set(),
	});
	const isLiked = replyState.likedBy.has("current-user");
	return `
        <div class="reply" data-reply-id="${reply.reply_id}">
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
            <div class="reply-actions">
              <button class="reply-action-button like-button ${
								isLiked ? "liked text-blue-600" : ""
							}" data-comment-id="${reply.comment_id}" data-reply-id="${
		reply.reply_id
	}">
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

ReplyManager.prototype.showReplyForm = async function (e) {
	e.stopPropagation();

	const button = e.target.closest(".reply-button");
	if (!button) return;

	const postId = button.getAttribute("data-post-id");
	const commentId = button.getAttribute("data-comment-id");

	const repliesContainer = button
		.closest(".comment")
		.querySelector(".replies-container");
	if (!repliesContainer) return;

	const userData = await getUserData();
	if (!userData) {
		toast.createToast("error", "Please login to reply");
		alert("Please login to reply");
		window.location.href = "/auth";
		return;
	}

	const existingReplyForm = repliesContainer.querySelector(".reply-form");
	if (existingReplyForm) {
		existingReplyForm.remove();
		return;
	}

	const replyFormHTML = `
        <form class="reply-form" data-comment-id="${commentId}" data-post-id="${postId}">
            <textarea placeholder="Write your reply..." class="reply-input"></textarea>
            <button type="submit" class="reply-submit">Reply</button>
        </form>`;

	repliesContainer.insertAdjacentHTML("afterbegin", replyFormHTML);

	const replyForm = repliesContainer.querySelector(".reply-form");
	if (replyForm) {
		replyForm.addEventListener("submit", (e) => this.handleReplySubmit(e));
	}
};

ReplyManager.prototype.handleReplySubmit = async function (e) {
	e.preventDefault();
	const userData = await getUserData();
	if (!userData) {
		toast.createToast("error", "Please login to reply");
		alert("Please login to reply");
		window.location.href = "/auth";
		return;
	}

	const form = e.target.closest(".reply-form");
	if (!form) return;

	const commentId = form.getAttribute("data-comment-id");
	const replyText = form.querySelector(".reply-input").value.trim();

	if (!replyText) {
		toast.createToast("warning", "Reply cannot be empty!");
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
		toast.createToast("error", res.message);
		return;
	}

	REPLIES[commentId].push(res.data);

	const commentElement = document.querySelector(
		`.comment[data-comment-id="${commentId}"]`
	);
	if (commentElement) {
		const replyHTML = this.createReplyHTML(res.data);
		const repliesContainer = commentElement.querySelector(".replies-container");
		if (repliesContainer) {
			repliesContainer.insertAdjacentHTML("beforeend", replyHTML);
		} else {
			commentElement.insertAdjacentHTML(
				"beforeend",
				`<div class="replies-container data-comment-id="${commentId}">${replyHTML}</div>`
			);
		}
	}

	form.remove();
	lucide.createIcons();
};

ReplyManager.prototype.handleReplyLikes = async function (e) {
	e.preventDefault();

	const currentUser = await getUserData();
	if (!currentUser?.user_id) {
		alert("Please log in to like comments.");
		window.location.href = "/auth";
		return;
	}

	const likeButton = e.target.closest(".like-button");
	if (!likeButton) return;

	const replyId = likeButton.dataset.replyId;
	const commentId = likeButton.dataset.commentId;

	if (!commentId || !replyId) return;

	const likeData = (commentLikeState.replies[replyId] ??= {
		count: 0,
		likedBy: new Set(),
	});

	const replyLikeData = {
		reply_id: replyId,
		user_id: currentUser.user_id,
		comment_id: commentId,
	};

	// Send request to backend
	const res = await this.commentService.likeReply(replyLikeData);

	if (res.error) {
		alert(res.message);
		return;
	}

	console.log(res);

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

export { ReplyManager };
