import { API_ENDPOINTS } from "./data.js";
import { getUserData } from "./authmiddleware.js";

// CommentService class for handling comment-related API requests
export class CommentService {
	constructor() {
		this.apiEndpoints = API_ENDPOINTS;
	}
}

// Method to fetch all comments for a post
CommentService.prototype.listCommentsByPost = async function (postID) {
	try {
		const response = await fetch(this.apiEndpoints.listcommbypost, {
			method: "POST",
			body: JSON.stringify({ post_id: postID }),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		// console.error("Error fetching comments:", error);
		return []; // Return empty array on failure
	}
};

// Method to create a new comment
CommentService.prototype.createComment = async function (commentData) {
	const userData = await getUserData();

	if (!userData) {
		return {
			error: true,
			message: "You need to login to post a comment!",
		};
	}

	if (!commentData?.comment) {
		return { error: true, message: "Please provide content for the comment!" };
	}

	const formData = {
		post_id: commentData.post_id,
		user_id: userData.user_id,
		user_name: userData.user_name,
		author_img: userData.image,
		parent_comment_id: commentData.parent_comment_id,
		comment: commentData.comment,
	};

	try {
		const response = await fetch(this.apiEndpoints.createcomment, {
			method: "POST",
			body: JSON.stringify(formData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error creating comment:", error);
		return {
			error: true,
			message: "Failed to create comment. Please try again.",
		};
	}
};

// Method to update a comment by ID
CommentService.prototype.updateComment = async function (commentData) {
	try {
		const response = await fetch(this.apiEndpoints.updatecomment, {
			method: "POST",
			body: JSON.stringify(commentData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		return {
			error: true,
			message: "Failed to edit comment. Please try again.",
		};
	}
};

// Method to delete a comment by ID
CommentService.prototype.deleteComment = async function (commentData) {
	try {
		const response = await fetch(this.apiEndpoints.deletecomment, {
			method: "POST",
			body: JSON.stringify(commentData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error deleting comment:", error);
		return {
			error: true,
			message: "Failed to delete comment. Please try again.",
		};
	}
};

// Method to like a comment by ID
CommentService.prototype.likeComment = async function (commentData) {
	const userData = await getUserData();

	if (!userData) {
		return {
			error: true,
			message: "You need to login to post a comment!",
		};
	}

	try {
		const response = await fetch(this.apiEndpoints.likecomment, {
			method: "POST",
			body: JSON.stringify(commentData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		return {
			error: true,
			data: null,
			message: error,
		};
	}
};

// Method to dislike a comment by ID
CommentService.prototype.dislikeComment = async function (commentData) {
	try {
		const response = await fetch(this.apiEndpoints.dislikecomment, {
			method: "POST",
			body: JSON.stringify(commentData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		console.error("Error disliking comment:", error);
		return null;
	}
};

// Method to post a reply
CommentService.prototype.createReply = async function (replyData) {
	const userData = await getUserData();

	if (!userData) {
		return {
			error: true,
			message: "You need to login to post a comment!",
		};
	}

	try {
		const response = await fetch(this.apiEndpoints.createReply, {
			method: "POST",
			body: JSON.stringify(replyData),
			headers: { "Content-Type": "application/json" },
		});

		return await response.json();
	} catch (error) {
		return {
			error: true,
			message: "Failed to create reply. Please try again.",
		};
	}
};
