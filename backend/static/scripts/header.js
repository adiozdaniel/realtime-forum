import { API_ENDPOINTS, userData } from "./data.js";
// import { PostManager, SAMPLE_POSTS } from "./index.js";
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
	// const filteredPosts = SAMPLE_POSTS.filter(
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
		}

		return response.status === 200;
	} catch (error) {
		console.error("Error signing out:", error);
	}
};

// Initialize function
Header.prototype.init = async function () {
	console.log("initializing data");
	// Automatically log out if on /auth
	if (window.location.pathname === "/auth") {
		this.signOutUser();
		this.authButton.textContent = "Sign In";
	}

	const userdata = await getUserData();

	if (userdata) {
		this.profileImage.src = userdata.image;
	}

	this.authButton.textContent = userdata ? "Sign Out" : "Sign In";

	// Event listeners
	this.menuToggleBtn?.addEventListener("click", this.toggleMobileMenu);
	this.searchInput?.addEventListener("input", this.handleSearch);
	this.darkModeToggle?.addEventListener("click", this.toggleDarkMode);
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
