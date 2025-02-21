import { AuthService } from "./authservice.js";
import { getUserData } from "./authmiddleware.js";
import { STATE } from "./data.js";

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
		window.location.href = "/auth"
		return;
	}

	console.log(userData);
	// this.state.profilePic = userData.image;
	// this.state.bio = userData.bio;

	if (userData.data) {
		this.state.profilePic = userData.data.user_info?.image;
		this.state.posts = userData.data.posts;
		this.state.userComments = userData.data.comments;
		this.state.activities = userData.data.recent_activity;
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
	};
	this.elements.bioText.textContent = this.state.bio || "Hey there! I'm on forum.";
	this.elements.profileImage.src = this.state.profilePic;
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
		state.posts = this.state.posts.filter(post => post.id !== postId);
		renderPosts();
		updateStats();
	};
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
	this.elements.sidebarItems.forEach(item => {
		item.classList.toggle('active', item.dataset.view === this.state.currentView);
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
	this.elements.darkModeToggle.innerHTML = `<i data-lucide="${this.state.darkMode ? 'sun' : 'moon'}"></i>`;
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
			`Image size is too large.Please upload an image less than ${maxFileSize / 1024 / 1024
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
	document.getElementById("postsCount").textContent = this.state.posts.length || 0;
	document.getElementById("commentsCount").textContent = this.state.userComments.length || 0;
	document.getElementById("likesCount").textContent = this.state.likes?.length || 0;
	document.getElementById("dislikeCount").textContent = this.state.dislikes?.length || 0;
	document.getElementById("repliesCount").textContent = this.state.replies?.length || 0;
};

ProfileDashboard.prototype.renderActivities = function () {
	document.getElementById("activityList").innerHTML = this.state.activities?.map((activity) => `
		<div class="activity-item">
			<i data-lucide="clock"></i>
			<span>${activity.content}</span> - 
			<span>${activity.timestamp}</span>
		</div>`)
		.join(" ");
	lucide.createIcons();
};

ProfileDashboard.prototype.renderPosts = function () {
	document.getElementById("postsList").innerHTML = this.state.posts
		.map((post) => `
		<div class="post-item">
                <p>${post.content}</p>
                <div class="post-actions">
                    <span><i data-lucide="thumbs-up"></i> ${post.likes}</span>
                    <span><i data-lucide="message-square"></i> ${post.comments}</span>
                    <span>${post.timestamp}</span>
                </div>
                <button class="delete-button" onclick="deletePost(${post.id})">
                    <i data-lucide="trash-2"></i>
                </button>
            </div>`)
		.join(" ");
	lucide.createIcons();
};

ProfileDashboard.prototype.renderComments = function () {
	document.getElementById("commentsList").innerHTML = this.state.userComments
		.map((comment) => `
		<div class="comment-item">
                <h3>Re: ${comment.postTitle}</h3>
                <p>${comment.content}</p>
                <div class="post-actions">
                    <span><i data-lucide="thumbs-up"></i> ${comment.likes}</span>
                    <span>${comment.timestamp}</span>
                </div>
            </div>
		`)
		.join(" ");
	lucide.createIcons();
};

const dashboard = new ProfileDashboard();
document.addEventListener("DOMContentLoaded", () => dashboard.init());
