import { AuthService } from "./authservice.js";
import { STATE } from "./data.js";
import { postManager } from "./postmanager.js";
import { formatTimeAgo } from "./timestamps.js";
import { toast } from "./toast.js";

function toTitleCase(str) {
	return str.replace(/\b\w/g, (char) => char.toUpperCase());
}

class ProfileDashboard {
	constructor() {
		this.authService = new AuthService();
		this.state = STATE;
		this.userData = null;
	}
}

ProfileDashboard.prototype.init = async function () {
	const userData = await this.authService.userDashboard();

	if (userData.error) {
		alert(userData.message);
		window.location.href = "/auth";
		return;
	}

	console.log(userData.data);

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

		this.userData = userData.data.user_info;
	}

	this.cacheElements();
	this.setupEventListeners();
	this.updateTheme();
	this.updateStats();
	this.renderActivities();
	postManager.renderPosts(this.state.posts);
	this.renderComments();
	this.updateActiveSection();
};

ProfileDashboard.prototype.createCommentHTML = function (comment) {
	return `
        <div class="comment-item" data-comment-id="${comment.comment_id}">
			<div class="comment-user-actions">
				<button class="edit-button" id="editCommentBtn" data-comment-id="${comment.comment_id}">
              		<i data-lucide="edit"></i>
            	</button>
				<button class="delete-button" id="deleteCommentBtn">
                    <i data-lucide="trash-2"></i>
            	</button>
			</div>
            <div class="comment-content"> 
                <div class="profile-image">
                    <img src="${comment.author_img}" 
                         onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
                </div>
                <div>
                    <div class="comment-author">${comment.user_name}</div> 
                    <div class="comment-text">${comment.comment}</div> 
                </div>
				<div>|</div>
				<div class="profile-image">
                    <img src="${comment.post_author_img}" 
                         onerror="this.onerror=null;this.src='/static/profiles/avatar.jpg';"/>
                </div>
				<div>
                    <div class="comment-author">${comment.post_title}</div> 
                    <div class="comment-text"> by ${comment.post_author}</div> 
                </div>
            </div>
            <div class="comment-footer">
                <div class="comment-actions"> 
                    <button class="comment-action-button like-button data-comment-id="${comment.comment_id}"> 
                        <i data-lucide="thumbs-up"></i> 
                        <span class="likes-count">${comment.likes?.length || 0}</span> 
                    </button>
					 <button class="comment-action-button dislike-button data-comment-id="${comment.comment_id}"> 
                        <i data-lucide="thumbs-down"></i> 
                        <span class="likes-count">${comment.dislikes?.length || 0}</span> 
                    </button>
                </div>
                <div class="comment-meta">
                    <span class="comment-time">${formatTimeAgo(comment.created_at)}</span> 
                </div>
            </div>
        </div>`;
};

ProfileDashboard.prototype.renderComments = function () {
	this.commentsList = document.getElementById("commentsList");

	let comments = ``;

	this.state.userComments?.forEach((comment) => {
		comments += this.createCommentHTML(comment);
	});

	this.commentsList.innerHTML = comments;
	lucide.createIcons();
}

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
};

ProfileDashboard.prototype.switchView = function (view) {
	this.state.currentView = view;
	this.updateActiveSection();
};

ProfileDashboard.prototype.updateActiveSection = function () {
	// Hide all sections
	Object.values(this.elements.sections).forEach((section) =>
		section.classList.add("hidden")
	);
	
	this.elements.sections[this.state.currentView].classList.remove("hidden");
	this.elements.sidebarItems.forEach((item) => item.classList.remove("active"));

	const activeItem = Array.from(this.elements.sidebarItems).find(
		(item) => item.dataset.view === this.state.currentView
	);
	if (activeItem) activeItem.classList.add("active");
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

ProfileDashboard.prototype.editBio = async function () {
	const newBio = prompt("Edit your bio:", this.state.bio);
	if (newBio) {
		const formData = {
			user_id: this.userData.user_id,
			bio: newBio,
		};

		const res = await this.authService.editBio(formData);

		if (res.error) {
			toast.createToast("error", res.message);
			return;
		}

		if (res.data) {
			toast.createToast("success", res.message || "Bio updated!");
			this.state.bio = res.data.bio;
			this.elements.bioText.textContent = res.data.bio;
		}
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
		console.log(user.data);

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

const dashboard = new ProfileDashboard();
document.addEventListener("DOMContentLoaded", () => dashboard.init());
