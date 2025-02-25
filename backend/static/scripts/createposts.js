import { POSTS, TEMP_DATA } from "./data.js";
import { postManager } from "./postmanager.js";
import { PostService } from "./postsservice.js";

class PostModalManager {
	constructor() {
		this.modal = document.getElementById("createPostModal");
		this.createPostBtn = document.getElementById("createPostBtn");
		this.cancelBtn = document.getElementById("cancelPost");
		this.form = document.getElementById("createPostForm");
		this.imageUpload = document.getElementById("imageUpload");
		this.imagePreviewContainer = document.getElementById(
			"imagePreviewContainer"
		);
		this.imagePreview = document.getElementById("imagePreview");
		this.removeImage = document.getElementById("removeImage");
		this.videoLink = document.getElementById("videoLink");
		this.videoPreviewContainer = document.getElementById(
			"videoPreviewContainer"
		);
		this.videoPreview = document.getElementById("videoPreview");
		this.removeVideo = document.getElementById("removeVideo");
		this.mediaPreview = document.getElementById("mediaPreview");
		this.uploadError = document.getElementById("uploadError");
		this.MAX_FILE_SIZE = 20 * 1024 * 1024;
		this.ALLOWED_TYPES = ["image/jpeg", "image/png", "image/gif"];
		this.postService = new PostService();
		this.posts = postManager;
	}
}

PostModalManager.prototype.init = function () {
	if (!this.createPostBtn || !this.cancelBtn || !this.modal) return;
	this.createPostBtn.addEventListener("click", this.openModal.bind(this));
	this.cancelBtn.addEventListener("click", this.closeModal.bind(this));
	this.modal.addEventListener("click", (e) => {
		if (e.target === this.modal) this.closeModal();
	});
	this.imageUpload.addEventListener(
		"change",
		this.handleImageUpload.bind(this)
	);
	this.videoLink.addEventListener("input", this.handleVideoLink.bind(this));
	this.removeImage.addEventListener(
		"click",
		this.removeImagePreview.bind(this)
	);
	this.removeVideo.addEventListener(
		"click",
		this.removeVideoPreview.bind(this)
	);
	this.form.addEventListener("submit", this.handleSubmit.bind(this));
};

PostModalManager.prototype.openModal = function (post) {
	if (window.location.pathname === "/dashboard") {
		this.form["postTitle"].value = post.post_title;
		this.form["postContent"].value = post.post_content;
		this.form["postId"].value = post.post_id;
		this.form["createdAt"].value = post.created_at;
		this.mediaPreview.classList.remove("hidden");
		this.imagePreviewContainer.classList.remove("hidden");
		this.imagePreview.src = post.post_image;

		const categories = post.post_category.split(" ");

		categories.forEach((category) => {
			const checkbox = document.querySelector(
				`input[id="${category.toLowerCase()}"]`
			);
			if (checkbox) checkbox.checked = true;
		});
	}

	this.modal.classList.remove("hidden");
};

PostModalManager.prototype.closeModal = function () {
	this.modal.classList.add("hidden");
	this.form.reset();
	this.imagePreview.src = "";
	this.imagePreviewContainer.classList.add("hidden");
	this.videoPreviewContainer.classList.add("hidden");
	this.mediaPreview.classList.add("hidden");
	this.uploadError.classList.add("hidden");
	this.uploadError.textContent = "";
};

// Handle image upload
PostModalManager.prototype.handleImageUpload = async function (e) {
	const file = e.target.files[0];
	this.uploadError.textContent = "";
	this.uploadError.classList.add("hidden");

	if (!file) return;

	if (!this.ALLOWED_TYPES.includes(file.type)) {
		this.showUploadError(
			"Invalid file type. Please upload a JPEG, PNG, or GIF image."
		);
		this.imageUpload.value = "";
		return;
	}

	if (file.size > this.MAX_FILE_SIZE) {
		this.showUploadError("File size exceeds 20MB limit.");
		this.imageUpload.value = "";
		return;
	}

	const reader = new FileReader();
	reader.onload = (e) => {
		this.imagePreview.src = e.target.result;
		this.imagePreviewContainer.classList.remove("hidden");
		this.mediaPreview.classList.remove("hidden");
		this.videoLink.value = "";
		this.videoPreviewContainer.classList.add("hidden");
	};
	reader.onerror = () => {
		this.showUploadError("Error reading file. Please try again.");
		this.imageUpload.value = "";
	};
	reader.readAsDataURL(file);

	const formData = new FormData();
	formData.append("image", file);

	try {
		const imgRes = await this.postService.uploadPostImg(formData);

		if (imgRes.error) alert(imgRes.message);

		if (imgRes.data !== null) TEMP_DATA = imgRes.data;
	} catch (error) {
		this.showUploadError("Error uploading image. Please try again.");
		this.imageUpload.value = "";
	}
};

PostModalManager.prototype.handleVideoLink = function (e) {
	const url = e.target.value.trim();
	if (this.isValidVideoUrl(url)) {
		const embedUrl = this.getEmbedUrl(url);
		if (embedUrl) {
			this.videoPreview.innerHTML = `<iframe width="100%" height="250" src="${embedUrl}" frameborder="0" allowfullscreen></iframe>`;
			this.videoPreviewContainer.classList.remove("hidden");
			this.mediaPreview.classList.remove("hidden");
			this.imageUpload.value = "";
			this.imagePreviewContainer.classList.add("hidden");
		}
	} else {
		this.videoPreviewContainer.classList.add("hidden");
	}
};

PostModalManager.prototype.removeImagePreview = function () {
	this.imageUpload.value = "";
	this.imagePreview.src = "";
	this.imagePreviewContainer.classList.add("hidden");
	if (this.videoPreviewContainer.classList.contains("hidden")) {
		this.mediaPreview.classList.add("hidden");
	}
	TEMP_DATA = null;
};

PostModalManager.prototype.removeVideoPreview = function () {
	this.videoLink.value = "";
	this.videoPreviewContainer.classList.add("hidden");
	if (this.imagePreviewContainer.classList.contains("hidden")) {
		this.mediaPreview.classList.add("hidden");
	}
};

PostModalManager.prototype.handleSubmit = async function (e) {
	e.preventDefault();

	const formData = {
		PostTitle: document.getElementById("postTitle").value,
		PostCategory: Array.from(
			document.querySelectorAll('input[name="postCategory"]:checked')
		)
			.map((checkbox) => checkbox.value)
			.join(" "),
		PostContent: document.getElementById("postContent").value,
		PostImage: TEMP_DATA?.img || null,
		PostID: TEMP_DATA?.post_id || null,
		PostVideo: this.videoLink.value || null,
	};

	try {
		const res = await this.postService.createPost(formData);
		if (!res.error) {
			this.closeModal();
		}
		POSTS.unshift(res.data);

		this.posts.renderPosts();
	} catch (error) {
		console.error("Error creating post:", error);
		this.showUploadError("Error creating post. Please try again.");
	}
};

PostModalManager.prototype.showUploadError = function (message) {
	this.uploadError.textContent = message;
	this.uploadError.classList.remove("hidden");
	this.uploadError.classList.remove("error-shake");
	void this.uploadError.offsetWidth;
	this.uploadError.classList.add("error-shake");
};

PostModalManager.prototype.isValidVideoUrl = function (url) {
	return /^(https?:\/\/)?(www\.)?(youtube\.com|youtu\.be|vimeo\.com)\/.+/.test(
		url
	);
};

PostModalManager.prototype.getEmbedUrl = function (url) {
	const youtubeRegex =
		/(?:youtube\.com\/(?:watch\?v=|embed\/|shorts\/)|youtu\.be\/|youtube\.com\/live\/)([a-zA-Z0-9_-]+)/;
	const vimeoRegex = /vimeo\.com\/(\d+)/;
	let match = url.match(youtubeRegex);
	if (match) return `https://www.youtube.com/embed/${match[1]}`;
	match = url.match(vimeoRegex);
	if (match) return `https://player.vimeo.com/video/${match[1]}`;
	return null;
};

export { PostModalManager };
