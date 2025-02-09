class PostService {
	constructor() {
		this.apiEndpoints = window.API_ENDPOINTS;
	}

	async fetchPosts() {
		console.log(this.apiEndpoints.allposts);
		try {
			const response = await fetch(this.apiEndpoints.allposts);
			if (!response.ok) {
				throw new Error("Failed to fetch posts");
			}
			const posts = await response.json();
			return posts;
		} catch (error) {
			console.error("Error fetching posts:", error);
			return [];
		}
	}

	async createPost(postData) {
		try {
			const response = await fetch(this.apiEndpoints.createpost, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(postData),
			});
			if (!response.ok) {
				throw new Error("Failed to create post");
			}
			const newPost = await response.json();
			return newPost;
		} catch (error) {
			console.error("Error creating post:", error);
			return null;
		}
	}

	async updatePost(postId, postData) {
		try {
			const response = await fetch(
				`${this.apiEndpoints.updatepost}/${postId}`,
				{
					method: "PUT",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify(postData),
				}
			);
			if (!response.ok) {
				throw new Error("Failed to update post");
			}
			const updatedPost = await response.json();
			return updatedPost;
		} catch (error) {
			console.error("Error updating post:", error);
			return null;
		}
	}

	async deletePost(postId) {
		try {
			const response = await fetch(
				`${this.apiEndpoints.deletepost}/${postId}`,
				{
					method: "DELETE",
				}
			);
			if (!response.ok) {
				throw new Error("Failed to delete post");
			}
			return true;
		} catch (error) {
			console.error("Error deleting post:", error);
			return false;
		}
	}

	async likePost(postId) {
		try {
			const response = await fetch(`${this.apiEndpoints.likepost}/${postId}`, {
				method: "POST",
			});
			if (!response.ok) {
				throw new Error("Failed to like post");
			}
			const updatedPost = await response.json();
			return updatedPost;
		} catch (error) {
			console.error("Error liking post:", error);
			return null;
		}
	}

	async dislikePost(postId) {
		try {
			const response = await fetch(
				`${this.apiEndpoints.dislikepost}/${postId}`,
				{
					method: "POST",
				}
			);
			if (!response.ok) {
				throw new Error("Failed to dislike post");
			}
			const updatedPost = await response.json();
			return updatedPost;
		} catch (error) {
			console.error("Error disliking post:", error);
			return null;
		}
	}
}

document.addEventListener("DOMContentLoaded", () => {
	window.postService = new PostService();
});
