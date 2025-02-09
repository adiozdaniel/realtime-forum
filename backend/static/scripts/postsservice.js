// PostService class for handling post-related API requests
class PostService {
	constructor() {
		this.apiEndpoints = window.API_ENDPOINTS;
		this.userData = window.RESDATA.userData;
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
	const formData = new URLSearchParams();
	formData.append("title", postData.title);
	formData.append("content", postData.content);
	formData.append("user_id", this.userData.user_id);

	try {
		const response = await fetch(this.apiEndpoints.createpost, {
			method: "POST",
			body: formData,
		});
		const responseText = await response.text(); // Get raw response text

		let newPost = JSON.parse(responseText); // Parse manually
		if (!response.ok) {
			throw new Error("Failed to create post");
		}
		newPost = await response.json(); // Parse response as JSON
		return newPost;
	} catch (error) {
		console.error("Error creating post:", error);
		return null; // Return null on failure
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
		if (!response.ok) {
			throw new Error("Failed to update post");
		}
		const updatedPost = await response.json(); // Parse response as JSON
		return updatedPost;
	} catch (error) {
		console.error("Error updating post:", error);
		return null; // Return null on failure
	}
};

// Method to delete a post by ID
PostService.prototype.deletePost = async function (postId) {
	try {
		const response = await fetch(`${this.apiEndpoints.deletepost}/${postId}`, {
			method: "DELETE",
		});
		if (!response.ok) {
			throw new Error("Failed to delete post");
		}
		return true; // Return true if deletion is successful
	} catch (error) {
		console.error("Error deleting post:", error);
		return false; // Return false on failure
	}
};

// Method to like a post by ID
PostService.prototype.likePost = async function (postId) {
	try {
		const response = await fetch(`${this.apiEndpoints.likepost}/${postId}`, {
			method: "POST",
		});
		if (!response.ok) {
			throw new Error("Failed to like post");
		}
		const updatedPost = await response.json(); // Parse response as JSON
		return updatedPost;
	} catch (error) {
		console.error("Error liking post:", error);
		return null; // Return null on failure
	}
};

// Method to dislike a post by ID
PostService.prototype.dislikePost = async function (postId) {
	try {
		const response = await fetch(`${this.apiEndpoints.dislikepost}/${postId}`, {
			method: "POST",
		});
		if (!response.ok) {
			throw new Error("Failed to dislike post");
		}
		const updatedPost = await response.json(); // Parse response as JSON
		return updatedPost;
	} catch (error) {
		console.error("Error disliking post:", error);
		return null; // Return null on failure
	}
};

// Initialize the PostService instance when the page loads
document.addEventListener("DOMContentLoaded", () => {
	window.postService = new PostService();
});
