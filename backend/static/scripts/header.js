import { renderPosts, SAMPLE_POSTS } from "./index.js";
import { sidebar } from "./sidebar.js";
// DOM Elements
const menuToggleBtn = document.querySelector("#menuToggle");
const searchInput = document.querySelector("#searchInput");
const darkModeToggle = document.querySelector("#darkModeToggle");
const authButton = document.querySelector(".sign-in-button");
const profileImageElement = document.querySelector(".profile-image img");
const userData = JSON.parse(localStorage.getItem("res"));

// Toggle mobile menu
function toggleMobileMenu() {
	const isVisible = sidebar.style.display === "block";
	sidebar.style.display = isVisible ? "none" : "block";
}

// Search functionality
function handleSearch(e) {
	const searchTerm = e.target.value.toLowerCase();
	const filteredPosts = SAMPLE_POSTS.filter(
		(post) =>
			post.title.toLowerCase().includes(searchTerm) ||
			post.excerpt.toLowerCase().includes(searchTerm)
	);
	renderPosts(filteredPosts);
}

// Toggle dark mode
function toggleDarkMode() {
	document.body.classList.toggle("dark-mode");
	localStorage.setItem(
		"darkMode",
		document.body.classList.contains("dark-mode")
	);
}
async function signOutUser() {
	try {
		let response = await fetch("/api/auth/logout", {
			method: "POST",
			credentials: "include",
		});

		if (!response.ok) console.log("Not logged in");

		if (response.ok) {
			console.log("User signed out successfully.");
		}

		return response.status === 200;
	} catch (error) {
		console.error("Error signing out:", error);
	}
}

async function isSignedIn() {
	if (!authButton) {
		console.error("Auth button not found.");
		return;
	}

	try {
		let response = await fetch("/api/auth/check", { credentials: "include" });
		if (!response.ok) {
			throw new Error("Not signed in");
		}

		let data = await response.json();

		return data.signedIn;
	} catch (error) {
		return false;
	}
}

// Initialize function
function init() {
	// handleResize();

	authButton.textContent = isSignedIn() ? "Sign Out" : "Sign In";

	// Automatically log out if on /auth
	if (window.location.pathname === "/auth") {
		signOutUser();
		authButton.textContent = "Sign In";
	}

	// Event listeners
	menuToggleBtn?.addEventListener("click", toggleMobileMenu);
	searchInput?.addEventListener("input", handleSearch);
	darkModeToggle?.addEventListener("click", toggleDarkMode);
	// Check for saved dark mode preference
	const savedDarkMode = localStorage.getItem("darkMode") === "true";
	if (savedDarkMode) {
		document.body.classList.add("dark-mode");
	}

	console.log("Profile Image URL:", userData?.image);
	console.log("Profile Image Element:", profileImageElement);

	// Update profile image
	if (userData && userData.image && userData.first_name) {
		profileImageElement.src = userData.image;
		profileImageElement.alt = userData.first_name;
	}
}

// Start the application
document.addEventListener("DOMContentLoaded", init);
