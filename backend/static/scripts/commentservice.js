import { API_ENDPOINTS, userData, SAMPLE_POSTS } from "./data.js";

// CommentService class for handling comment-related API requests
export class CommentService {
	constructor() {
		this.apiEndpoints = API_ENDPOINTS;
		this.userData = {
			userDetails: () => localStorage.getItem("userdata"),
			data: function () {
				try {
					passedData = JSON.parse(this.userDetails());
					return passedData?.data || null;
				} catch (error) {
					console.log("error", error);
					return null;
				}
			},
		};
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
	const local = localStorage.getItem("userdata");
	const userData = JSON.parse(local);

	if (!userData) {
		console.log("local storage", JSON.parse(localStorage.getItem("userdata")));
		return {
			error: true,
			message: "You need to login to post a comment!",
		};
	}

	if (!commentData?.content) {
		return { error: true, message: "Please provide content for the comment!" };
	}

	const formData = {
		post_id: commentData.post_id,
		user_id: userData.user_id,
		user_name: userData.user_name,
		author_img: userData.image,
		parent_comment_id: commentData.parent_comment_id,
		comment: commentData.content,
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

// Method to delete a comment by ID
CommentService.prototype.deleteComment = async function (commentID) {
	try {
		const response = await fetch(
			`${this.apiEndpoints.deleteComment}/${commentID}`,
			{
				method: "DELETE",
			}
		);

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
