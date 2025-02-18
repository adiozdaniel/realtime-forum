import { AuthService } from "./authservice.js";
import { getUserData } from "./authmiddleware.js";

class ProfileDashboard {
	constructor() {
		this.authService = new AuthService();
		this.state = {
			currentView: "overview",
			darkMode: localStorage.getItem("darkMode") === "true",
			profilePic: "",
			bio:
				localStorage.getItem("userBio") ||
				"Hi, I love coding and sharing knowledge with the community!",
			posts: [
				{
					id: 1,
					content: "Just learned about React hooks!",
					comments: 5,
					likes: 12,
					timestamp: "2h ago",
				},
				{
					id: 2,
					content: "Working on a new project using TypeScript",
					comments: 3,
					likes: 8,
					timestamp: "5h ago",
				},
				{
					id: 3,
					content: "Check out my latest blog post about web performance",
					comments: 8,
					likes: 15,
					timestamp: "1d ago",
				},
			],
			userComments: [
				{
					id: 1,
					postTitle: "Introduction to GraphQL",
					content: "Great explanation!",
					likes: 5,
					timestamp: "3h ago",
				},
				{
					id: 2,
					postTitle: "Docker Best Practices",
					content: "Very effective!",
					likes: 3,
					timestamp: "1d ago",
				},
			],
			activities: [
				{ type: "post", content: "Created a new post", timestamp: "2h ago" },
				{
					type: "comment",
					content: "Commented on 'Docker Best Practices'",
					timestamp: "1d ago",
				},
				{
					type: "like",
					content: "Liked 'Introduction to GraphQL'",
					timestamp: "1d ago",
				},
			],
		};
	}
}

ProfileDashboard.prototype.init = async function () {
	const user = await getUserData();

	if (user) this.userData = user;

	if (!user) window.location.href = "/auth";

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
	this.elements.bioText.textContent = this.state.bio;
	this.elements.profileImage.src = this.userData?.image;
	this.elements.userName.textContent = this.userData?.user_name;
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
	formData.append("profileImage", file);

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
	document.getElementById("postsCount").textContent = this.state.posts.length;
	document.getElementById("commentsCount").textContent =
		this.state.userComments.length;
	document.getElementById("likesCount").textContent = this.state.posts.reduce(
		(acc, post) => acc + post.likes,
		0
	);
};

ProfileDashboard.prototype.renderActivities = function () {
	document.getElementById("activityList").innerHTML = this.state.activities
		.map((activity) => `<div>${activity.content} - ${activity.timestamp}</div>`)
		.join(" ");
};

ProfileDashboard.prototype.renderPosts = function () {
	document.getElementById("postsList").innerHTML = this.state.posts
		.map((post) => `<div>${post.content} - ${post.timestamp}</div>`)
		.join(" ");
};

ProfileDashboard.prototype.renderComments = function () {
	document.getElementById("commentsList").innerHTML = this.state.userComments
		.map((comment) => `<div>${comment.content} - ${comment.timestamp}</div>`)
		.join(" ");
};

const dashboard = new ProfileDashboard();
document.addEventListener("DOMContentLoaded", () => dashboard.init());
