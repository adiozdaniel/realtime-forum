import { formatTimeAgo } from "./timestamps.js";
import { getUserData } from "./authmiddleware.js";
import { commentLikeState, REPLIES } from "./data.js";
import { CommentService } from "./commentservice.js";

class ReplyManager {
	constructor() {
		this.likeState = commentLikeState;
		this.commentService = new CommentService();
	}
}

ReplyManager.prototype.createReplyHTML = function (reply) {
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
              <button class="comment-action-button like-button ${
								isLiked ? "liked text-blue-600" : ""
							}" data-comment-id="${reply.id}">
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
		const replyForm = commentElement.querySelector(".reply-form");
		if (replyForm) {
			replyForm.addEventListener("submit", (e) => this.handleReplySubmit(e));
		}
	}
};

ReplyManager.prototype.handleReplySubmit = async function (e) {
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
				`<div class="replies-container">${replyHTML}</div>`
			);
		}
	}

	form.remove();
	lucide.createIcons();
};

export { ReplyManager };
