// CommentService class for handling comment-related API requests
export class CommentService {
	constructor() {
		this.apiEndpoints = window.API_ENDPOINTS;
		this.userData = window.RESDATA?.userData;
	}
}

// Method to fetch all comments for a post
CommentService.prototype.listCommentsByPost = async function (postID) {
	try {
		const response = await fetch(this.apiEndpoints.listCommentsByPost, {
			method: "POST",
			body: JSON.stringify({ post_id: postID }),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error fetching comments:", error);
		return []; // Return empty array on failure
	}
};

// Method to create a new comment
CommentService.prototype.createComment = async function (commentData) {
	if (!this.userData?.user_id) {
		return { error: true, message: "You need to login to comment on this post!" };
	}

	if (!commentData?.content) {
		return { error: true, message: "Please provide content for the comment!" };
	}

	const formData = {
		post_id: commentData.post_id,
		user_id: this.userData.user_id,
		user_name: this.userData.user_name,
		parent_comment_id: commentData.parent_comment_id,
		content: commentData.content,
	};

	try {
		const response = await fetch(this.apiEndpoints.createComment, {
			method: "POST",
			body: JSON.stringify(formData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error creating comment:", error);
		return { error: true, message: "Failed to create comment. Please try again." };
	}
};

// Method to delete a comment by ID
CommentService.prototype.deleteComment = async function (commentID) {
	try {
		const response = await fetch(`${this.apiEndpoints.deleteComment}/${commentID}`, {
			method: "DELETE",
		});

		return response.ok; // Returns true if deletion is successful
	} catch (error) {
		console.error("Error deleting comment:", error);
		return false;
	}
};

// Method to like a comment by ID
CommentService.prototype.likeComment = async function (commentID) {
	try {
		const response = await fetch(this.apiEndpoints.likeComment, {
			method: "POST",
			body: JSON.stringify({ comment_id: commentID }),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error liking comment:", error);
		return null;
	}
};

// Method to dislike a comment by ID
CommentService.prototype.dislikeComment = async function (commentID) {
	try {
		const response = await fetch(this.apiEndpoints.dislikeComment, {
			method: "POST",
			body: JSON.stringify({ comment_id: commentID }),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error disliking comment:", error);
		return null;
	}
};
