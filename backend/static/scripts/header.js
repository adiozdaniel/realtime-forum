import { API_ENDPOINTS } from "./data.js";
// import { PostManager, POSTS } from "./index.js";
// import { sidebar } from "./sidebar.js";

// const postManager = new PostManager()

import { getUserData } from "./authmiddleware.js";

class Header {
	constructor() {
		this.endpoints = API_ENDPOINTS;

		// DOM Elements
		this.menuToggleBtn = document.querySelector("#menuToggle");
		this.searchInput = document.querySelector("#searchInput");
		this.darkModeToggle = document.querySelector("#darkModeToggle");
		this.authButton = document.querySelector(".sign-in-button");
		this.profileImage = document.querySelector("#userProfileImage");
	}
}

// Toggle mobile menu
Header.prototype.toggleMobileMenu = function () {
	// const isVisible = sidebar.style.display === "block";
	// sidebar.style.display = isVisible ? "none" : "block";
};

// Search functionality
Header.prototype.handleSearch = (e) => {
	// const searchTerm = e.target.value.toLowerCase();
	// const filteredPosts = POSTS.filter(
	// 	(post) =>
	// 		post.title.toLowerCase().includes(searchTerm) ||
	// 		post.excerpt.toLowerCase().includes(searchTerm)
	// );
	// postManager.renderPosts(filteredPosts);
};

// Toggle dark mode
Header.prototype.toggleDarkMode = function () {
	document.body.classList.toggle("dark-mode");
	localStorage.setItem(
		"darkMode",
		document.body.classList.contains("dark-mode")
	);
};

Header.prototype.signOutUser = async function () {
	try {
		let response = await fetch(this.endpoints.logout, {
			method: "POST",
			credentials: "include",
		});

		if (response.error) {
			console.log();
		}

		if (response.ok) {
			console.log("User signed out successfully.");
			this.authButton.textContent = "Sign In";
		}

		return response.status === 200;
	} catch (error) {
		console.error("Error signing out:", error);
	}
};

Header.prototype.handleAuth = async function () {
	if (window.location.pathname === "/auth") {
		if (this.authButton.textContent === "Sign In") return;
	}

	if (this.authButton.textContent === "Sign Out") this.signOutUser();

	if (window.location.pathname !== "/auth") {
		if (this.authButton.textContent === "Sign In")
			window.location.href = "/auth";
	}
};

Header.prototype.handleUserChange = function (userdata) {
	// Update profile image
	this.profileImage.src = userdata?.image || "/static/profiles/avatar.jpg";

	// Update auth button text
	this.authButton.textContent = userdata ? "Sign Out" : "Sign In";

	// Redirect if needed
	if (window.location.pathname === "/auth" && userdata) {
		window.location.href = "/";
	}
};

// Initialize function
Header.prototype.init = async function () {
	const userdata = await getUserData();
	this.handleUserChange(userdata);

	// Set the profile image with a fallback in case of error
	this.profileImage.src = userdata?.image;
	this.profileImage.onerror = () => {
		this.profileImage.src = "/static/profiles/avatar.jpg";
	};

	this.authButton.textContent = userdata ? "Sign Out" : "Sign In";

	// Automatically log out if on /auth
	if (window.location.pathname === "/auth")
		if (userdata) window.location.href = "/";

	// Event listeners
	this.menuToggleBtn?.addEventListener("click", this.toggleMobileMenu);
	this.searchInput?.addEventListener("input", this.handleSearch);
	this.darkModeToggle?.addEventListener("click", this.toggleDarkMode);
	this.authButton?.addEventListener("click", this.handleAuth.bind(this));
	// Check for saved dark mode preference
	const savedDarkMode = localStorage.getItem("darkMode") === "true";
	if (savedDarkMode) {
		document.body.classList.add("dark-mode");
	}
};

// Start the application
document.addEventListener("DOMContentLoaded", () => {
	setTimeout(() => {
		const header = new Header();
		header.init();
	}, 500);
});
