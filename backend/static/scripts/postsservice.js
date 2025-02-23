import { getUserData } from "./authmiddleware.js";
import { API_ENDPOINTS } from "./data.js";

// PostService class for handling post-related API requests
export class PostService {
	constructor() {
		this.apiEndpoints = API_ENDPOINTS;
	}
}

// Method to fetch all posts
PostService.prototype.fetchPosts = async function () {
	try {
		const response = await fetch(this.apiEndpoints.allposts);
		if (!response.ok) {
			throw new Error("Failed to fetch posts");
		}
		const posts = await response.json();
		return posts;
	} catch (error) {
		console.error("Error fetching posts:", error);
		return []; // Return empty array on failure
	}
};

// Method to create a new post
PostService.prototype.createPost = async function (postData) {
	const userData = await getUserData();
	if (!userData) {
		return {
			error: true,
			message: "You need to login to create a post!",
		};
	}

	if (!postData?.PostTitle) {
		return {
			error: true,
			message: "Please provide a title for the post!",
		};
	}

	if (!postData?.PostContent) {
		return {
			error: true,
			message: "Please provide a content for the post!",
		};
	}

	if (!postData?.PostCategory) {
		return {
			error: true,
			message: "Please select a category for the post!",
		};
	}

	const formData = {
		post_title: postData.PostTitle,
		post_content: postData.PostContent,
		user_id: userData.user_id,
		post_category: postData.PostCategory,
		post_author: userData.user_name,
		post_image: postData.PostImage,
		post_id: postData.PostID,
		post_video: postData.PostVideo,
	};

	try {
		const response = await fetch(this.apiEndpoints.createpost, {
			method: "POST",
			body: JSON.stringify(formData),
			headers: {
				"Content-Type": "application/json",
			},
		});

		const newPost = await response.json();

		return newPost;
	} catch (error) {
		if (error) {
			return {
				error: true,
				message: "You need to login to create a post!",
			};
		}
	}
};

// Method to update an existing post
PostService.prototype.updatePost = async function (postId, postData) {
	try {
		const response = await fetch(`${this.apiEndpoints.updatepost}/${postId}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(postData), // Convert post data to JSON
		});

		const updatedPost = await response.json(); // Parse response as JSON
		return updatedPost;
	} catch (error) {
		return null; // Return null on failure
	}
};

// Method to delete a post by ID
PostService.prototype.deletePost = async function (postId) {
	try {
		const response = await fetch(`${this.apiEndpoints.deletepost}/${postId}`, {
			method: "DELETE",
		});

		return true; // Return true if deletion is successful
	} catch (error) {
		return false; // Return false on failure
	}
};

// Method to like a post by ID
PostService.prototype.likePost = async function (postData) {
	if (!postData.user_id) {
		return {
			error: true,
			message: "You need to login to like the post!",
		};
	}

	try {
		const response = await fetch(this.apiEndpoints.likepost, {
			method: "POST",
			body: JSON.stringify(postData),
			headers: {
				"Content-Type": "application/json",
			},
		});

		const updatedPost = await response.json();
		return updatedPost;
	} catch (error) {
		return null; // Return null on failure
	}
};

// Method to dislike a post by ID
PostService.prototype.dislikePost = async function (postData) {
	if (!postData.user_id) {
		return {
			error: true,
			message: "You need to login to dislike the post!",
		};
	}

	try {
		const response = await fetch(this.apiEndpoints.dislikepost, {
			method: "POST",
			body: JSON.stringify(postData),
			headers: {
				"Content-Type": "application/json",
			},
		});

		const updatedPost = await response.json();
		return updatedPost;
	} catch (error) {
		console.error("Error disliking post:", error);
		return null; // Return null on failure
	}
};

// Method to upload a profile picture
PostService.prototype.uploadPostImg = async function (formData) {
	try {
		const response = await fetch(this.apiEndpoints.uploadPostImg, {
			method: "POST",
			body: formData,
		});

		return response.json();
	} catch (error) {
		return error;
	}
};

// Method to check notifications
PostService.prototype.checkNotifications = async function () {
	try {
		const response = await fetch(this.apiEndpoints.checkNotifications, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		});

		return response.json();
	} catch (error) {
		return error;
	}
};
