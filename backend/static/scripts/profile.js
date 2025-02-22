import { AuthService } from "./authservice.js";
import { STATE } from "./data.js";
import { formatTimeAgo } from "./timestamps.js";

function toTitleCase(str) {
	return str.replace(/\b\w/g, (char) => char.toUpperCase());
}

class ProfileDashboard {
	constructor() {
		this.authService = new AuthService();
		this.state = STATE;
	}
}

ProfileDashboard.prototype.init = async function () {
	const userData = await this.authService.userDashboard();

	if (userData.error) {
		alert(userData.message);
		window.location.href = "/auth";
		return;
	}

	console.log(userData);
	// this.state.profilePic = userData.image;
	// this.state.bio = userData.bio;

	if (userData.data) {
		this.state.profilePic = userData.data.user_info?.image;
		this.state.posts = userData.data.posts || [];
		this.state.userComments = userData.data.comments || [];
		this.state.activities = userData.data.activities || [];
		this.state.likes = userData.data.likes;
		this.state.dislikes = userData.data.dislikes;
		this.state.replies = userData.data.replies;
		this.state.bio = userData.data.user_info?.bio;
		this.state.username = userData.data.user_info?.user_name;
	}

	this.cacheElements();
	this.setupEventListeners();
	this.updateTheme();
	this.updateStats();
	this.renderActivities();
	this.renderPosts();
	this.renderComments();
	this.updateActiveSection();
};

ProfileDashboard.prototype.cacheElements = function () {
	this.elements = {
		userName: document.getElementById("username"),
		profileImage: document.getElementById("profileImage"),
		headerImage: document.getElementById("userProfileImage"),
		imageUpload: document.getElementById("imageUpload"),
		bioText: document.getElementById("bioText"),
		editBioButton: document.getElementById("editBioButton"),
		darkModeToggle: document.getElementById("darkModeToggle"),
		sections: {
			overview: document.getElementById("overviewSection"),
			posts: document.getElementById("postsSection"),
			comments: document.getElementById("commentsSection"),
			settings: document.getElementById("settingsSection"),
		},
		sidebarItems: document.querySelectorAll(".sidebar-item"),
		modal: {
			container: document.getElementById("postModal"),
			content: document.getElementById("modalPostContent"),
			closeBtn: document.querySelector(".close-modal"),
			viewFullBtn: document.getElementById("viewFullPost"),
		},
	};
	this.elements.bioText.textContent =
		this.state.bio || "Hey there! I'm on forum.";
	this.elements.profileImage.src =
		this.state.profilePic ||
		"data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%239ca3af'%3E%3Cpath d='M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08 1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z'/%3E%3C/svg%3E";
	this.elements.userName.textContent = toTitleCase(this.state.username);
};

ProfileDashboard.prototype.setupEventListeners = function () {
	this.elements.darkModeToggle.addEventListener(
		"click",
		this.toggleDarkMode.bind(this)
	);
	this.elements.imageUpload.addEventListener(
		"change",
		this.handleImageUpload.bind(this)
	);
	this.elements.editBioButton.addEventListener(
		"click",
		this.editBio.bind(this)
	);
	this.elements.sidebarItems.forEach((item) =>
		item.addEventListener("click", () => this.switchView(item.dataset.view))
	);

	//temporary handler for deleting posts
	window.deletePost = (postId) => {
		this.state.posts = this.state.posts.filter((post) => post.id !== postId);
		this.renderPosts();
		this.updateStats();
	};

	// Modal event listeners
	this.elements.modal.closeBtn?.addEventListener("click", () =>
		this.closeModal()
	);
	this.elements.modal.container?.addEventListener("click", (e) => {
		if (e.target === this.elements.modal.container) {
			this.closeModal();
		}
	});

	// Handle ESC key
	document.addEventListener("keydown", (e) => {
		if (
			e.key === "Escape" &&
			!this.elements.modal.container.classList.contains("hidden")
		) {
			this.closeModal();
		}
	});
};

ProfileDashboard.prototype.switchView = function (view) {
	this.state.currentView = view;
	this.updateActiveSection();
};

ProfileDashboard.prototype.updateActiveSection = function () {
	Object.values(this.elements.sections).forEach((section) =>
		section.classList.add("hidden")
	);
	this.elements.sections[this.state.currentView].classList.remove("hidden");
	this.elements.sidebarItems.forEach((item) => {
		item.classList.toggle(
			"active",
			item.dataset.view === this.state.currentView
		);
	});
};

ProfileDashboard.prototype.toggleDarkMode = function () {
	this.state.darkMode = !this.state.darkMode;
	localStorage.setItem("darkMode", this.state.darkMode);
	this.updateTheme();
};

ProfileDashboard.prototype.updateTheme = function () {
	document.body.setAttribute(
		"data-theme",
		this.state.darkMode ? "dark" : "light"
	);
	this.elements.darkModeToggle.innerHTML = `<i data-lucide="${
		this.state.darkMode ? "sun" : "moon"
	}"></i>`;
	lucide.createIcons();
};

ProfileDashboard.prototype.editBio = function () {
	const newBio = prompt("Edit your bio:", this.state.bio);
	if (newBio) {
		this.state.bio = newBio;
		localStorage.setItem("userBio", newBio);
		this.elements.bioText.textContent = newBio;
	}
};

ProfileDashboard.prototype.handleImageUpload = async function (e) {
	const file = e.target.files[0];

	if (!file) return;

	// Validate file types and size
	const ALLOWED_TYPES = ["image/jpeg", "image/png", "image/gif"];
	const maxFileSize = 5 * 1024 * 1024; // 5MB

	if (!ALLOWED_TYPES.includes(file.type)) {
		alert("Invalid file type. Please upload a JPEG, PNG, or GIF image.");
		this.elements.imageUpload.value = ""; // Clear the input
		return;
	}

	if (file.size > maxFileSize) {
		alert(
			`Image size is too large.Please upload an image less than ${
				maxFileSize / 1024 / 1024
			} MB.`
		);
		this.elements.imageUpload.value = ""; // Clear the input
		return;
	}

	// Read the file and display it immediately
	const reader = new FileReader();
	reader.onloadend = () => {
		// Update the profile picture immediately
		this.state.profilePic = reader.result;
		this.elements.profileImage.src = reader.result;
	};
	reader.readAsDataURL(file);

	// Upload the file to the server
	const formData = new FormData();
	formData.append("image", file);

	try {
		const user = await this.authService.uploadProfilePic(formData);

		if (user.error) {
			alert(user.message);
		} else if (user.data !== null) {
			// Update the profile picture URL with the server's response
			this.state.profilePic = user.data.image;
			this.elements.profileImage.src =
				user.data.image || "/static/profiles/avatar.jpg";
			this.elements.headerImage.src = user.data.image;
			alert("Profile picture updated successfully!");
		}
	} catch (error) {
		console.error("Error uploading image:", error);
		alert("Failed to upload image. Please try again.");
	}
};

ProfileDashboard.prototype.updateStats = function () {
	document.getElementById("postsCount").textContent =
		this.state.posts?.length || 0;
	document.getElementById("commentsCount").textContent =
		this.state.userComments?.length || 0;
	document.getElementById("likesCount").textContent =
		this.state.likes?.length || 0;
	document.getElementById("dislikeCount").textContent =
		this.state.dislikes?.length || 0;
	document.getElementById("repliesCount").textContent =
		this.state.replies?.length || 0;
};

ProfileDashboard.prototype.renderActivities = function () {
	document.getElementById("activityList").innerHTML = this.state.activities
		?.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
		.slice(0, 6)
		.map(
			(activity) => `
		<div class="activity-item">
			<i data-lucide="clock"></i>
			<span>${activity.activity_data}</span> - 
			<span>${formatTimeAgo(activity.created_at)}</span>
		</div>`
		)
		.join(" ");
	lucide.createIcons();
};

ProfileDashboard.prototype.renderPosts = function () {
	if (this.state.posts.length > 0) {
		document.getElementById("postsList").innerHTML = this.state.posts
			.map(
				(post) => `
		<div class="post-item" data-post-id="${post.post_id}" style="cursor: pointer;">
                <article class="post-card" data-post-id="${post.post_id}">
	  	${
				post.post_image
					? `
			<div class="post-img">
			<img src="${post.post_image}" alt="Post Image"/>
			</div>`
					: ""
			}
        <div class="flex items-start justify-between">
          <div>
			<div class="post-categories">
				${post.post_category
					.split(" ") // Split multiple categories
					.map(
						(category) =>
							`<span class="post-category">${category.trim()}</span>`
					)
					.join("")}
			</div>
            <h3 class="post-title">${post.post_title}</h3>
            <p class="post-excerpt">${post.post_content}</p>
          </div>
        </div>
        <div class="post-footer">
          <div class="post-actions">
            <button class="post-action-button like-button" data-post-id="${
							post.post_id
						}">
              <i data-lucide="thumbs-up"></i>
              <span class="likes-count">${post.likes?.length || 0}</span>
            </button>
			<button class="post-action-button dislike-button" data-post-id="${
				post.post_id
			}">
			<i data-lucide="thumbs-down"></i>
			 <span class="likes-count">${post.dislikes?.length || 0}</span>
			</button>
            <button class="post-action-button comment-toggle" data-post-id="${
							post.post_id
						}">
              <i data-lucide="message-square"></i>
              <span class="comments-count">${post.comments?.length || 0}</span>
            </button>
          </div>
          <div class="post-meta">
		   <div class="profile-image">
		  		<img src=${
						post.author_img
					} onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
			</div>
            <span>by ${post.post_author}</span>
            <span>•</span>
            <span>${formatTimeAgo(post.updated_at)}</span>
          </div>
        </div>
        <div class="comments-section hidden" id="comments-${post.post_id}">
          <div class="comments-container"></div>
        </div>
      </article>
            </div>
			`
			)
			.join(" ");

		// Add click handlers to post items
		document.querySelectorAll(".post-item").forEach((post) => {
			post.addEventListener("click", (e) => {
				const postId = post.dataset.postId;
				this.showPostPreview(postId);
			});
		});

		document.querySelectorAll(".post-action-button").forEach((button) => {
			button.addEventListener("click", (e) => e.stopPropagation());
		});
	} else {
		document.getElementById(
			"postsList"
		).innerHTML = `<div class="post-item"><p>There are no posts yet</p></div>`;
	}
	lucide.createIcons();
};

ProfileDashboard.prototype.showPostPreview = function (postId) {
	const post = this.state.posts.find((p) => p.post_id === postId);
	if (!post) return;

	// Populate modal content
	this.elements.modal.content.innerHTML = `
		<div class="modal-post">
			${
				post.post_image
					? `
				<div class="modal-post-img">
					<img src="${post.post_image}" alt="Post Image"/>
				</div>
			`
					: ""
			}
			<div class="modal-post-header">
				<div class="post-categories">
					${post.post_category
						.split(" ")
						.map(
							(category) =>
								`<span class="post-category">${category.trim()}</span>`
						)
						.join("")}
				</div>
				<h2 class="modal-post-title">${post.post_title}</h2>
			</div>
			<div class="modal-post-content">
				<p>${post.post_content}</p>
			</div>
			<div class="modal-post-footer">
				<div class="post-meta">
					<div class="profile-image">
						<img src="${
							post.author_img
						}" onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
					</div>
					<span>by ${post.post_author}</span>
					<span>•</span>
					<span>${formatTimeAgo(post.updated_at)}</span>
				</div>
			</div>
		</div>
	`;

	// Update "View Full Post" button
	this.elements.modal.viewFullBtn.addEventListener("click", () => {
		window.location.href = `/post/${postId}`;
	});

	// Show modal
	this.elements.modal.container.classList.remove("hidden");
};

ProfileDashboard.prototype.closeModal = function () {
	this.elements.modal.container.classList.add("hidden");
	this.elements.modal.content.innerHTML = "";
};

ProfileDashboard.prototype.renderComments = function () {
	if (this.state.userComments.length > 0) {
		document.getElementById("commentsList").innerHTML = this.state.userComments
			.map(
				(comment) => `
		<div class="comment-item">
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
                    <button 
					class="comment-action-button like-button"
					data-comment-id="${comment.comment_id}"
					> 
                        <i data-lucide="thumbs-up"></i> 
                        <span class="likes-count">${
													comment.likes?.length || 0
												}</span> 
                    </button> 
                </div>
                <div class="comment-meta">
                    <span class="comment-time">${formatTimeAgo(
											comment.created_at
										)}</span> 
                </div>
            </div>
        </div>
            </div>
		`
			)
			.join(" ");
	} else {
		document.getElementById(
			"commentsList"
		).innerHTML = `<div class="comment-item"><p>There are no comments yet</p></div>`;
	}
	lucide.createIcons();
};

const dashboard = new ProfileDashboard();
document.addEventListener("DOMContentLoaded", () => dashboard.init());
