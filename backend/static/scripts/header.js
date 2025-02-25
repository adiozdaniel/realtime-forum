import { API_ENDPOINTS } from "./data.js";
import { postManager } from "./postmanager.js";
import { getUserData } from "./authmiddleware.js";
import { sidebar } from "./sidebar.js";
import { PostService } from "./postsservice.js";
import { PostModalManager } from "./createposts.js";

class Header {
	constructor() {
		this.endpoints = API_ENDPOINTS;
		this.postService = new PostService();
		this.postModalManager = new PostModalManager();
		this.unreadNotifications = null;
		this.newUnread = null;
		this.noticationsCount = 0;
		this.enableNotifications = false;

		// DOM Elements
		this.menuToggleBtn = document.querySelector("#menuToggle");
		this.searchInput = document.querySelector("#searchInput");
		this.authButton = document.querySelector(".sign-in-button");
		this.profileImage = document.querySelector("#userProfileImage");
		this.darkModeToggle = document.querySelector("#darkModeToggle");
		this.notificationButton = document.querySelector("#notificationButton");
		this.notificationDropdown = document.querySelector("#notificationDropdown");
		this.postEditBtn = document.querySelector("#postEditBtn");
		this.postDeleteBtn = document.querySelector("#postDeleteBtn");
	}
}

// Toggle mobile menu
Header.prototype.toggleMobileMenu = function () {
	const isVisible = sidebar.style.display === "block";
	sidebar.style.display = isVisible ? "none" : "block";
};

Header.prototype.handleResize = function () {
	if (!sidebar) {
		return;
	}

	if (window.innerWidth >= 768) {
		sidebar.style.display = "block";
	} else {
		sidebar.style.display = "none";
	}
};

// Search functionality
Header.prototype.handleSearch = (e) => {
	const searchTerm = e.target.value.toLowerCase();
	postManager.searchPosts(searchTerm, []);
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

		if (this.authCheckInterval) {
			clearInterval(this.authCheckInterval);
			this.authCheckInterval = null;
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

// Watch over notifications
Header.prototype.watchNotifications = function () {
	let unreadNotifications = new Set();
	let canCheck = true;

	if (!canCheck) return;

	setInterval(async () => {
		const res = await this.postService.checkNotifications();

		if (res.error) {
			canCheck = false;
			return;
		}

		this.newUnread = res.data?.filter((n) => !n.is_read);
		const newUnreadIds = new Set(this.newUnread?.map((n) => n.notification_id));

		// Update unread count
		if (newUnreadIds.size > unreadNotifications.size) {
			const existingCount = this.notificationButton.querySelector(
				".notification-count"
			);
			if (existingCount) existingCount.remove();

			// create and append new notification count
			const countSpan = document.createElement("span");
			countSpan.className = "notification-count";
			countSpan.textContent = newUnreadIds.size;

			this.notificationButton.appendChild(countSpan);
			this.notificationButton.classList.add("has-notifications");

			this.notificationDropdown.innerHTML = "";
			this.newUnread.forEach((n) => {
				const notifItem = document.createElement("div");
				notifItem.classList.add("notification-item");
				notifItem.textContent = n.message;
				this.notificationDropdown.appendChild(notifItem);
			});

			const audio = new Audio("/static/audio/bell.mp3");
			if (this.enableNotifications)
				audio.play().catch(() => console.error("error ringing bell"));
		} else {
			this.notificationButton.classList.remove("has-notifications");
		}

		unreadNotifications = newUnreadIds;
	}, 5000);
};

// Handle notifications
Header.prototype.handleNotifications = function (e) {
	e.stopPropagation();

	if (!this.newUnread) return;

	if (this.notificationDropdown.style.display === "")
		this.notificationDropdown.style.display = "none"

	this.notificationDropdown.style.display =
		this.notificationDropdown.style.display === "none" ? "block" : "none";
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
	window.addEventListener("resize", this.handleResize);
	this.authButton?.addEventListener("click", this.handleAuth.bind(this));
	this.notificationButton?.addEventListener(
		"click",
		this.handleNotifications.bind(this)
	);
	
	this.postEditBtn?.addEventListener("click", (e) => this.handlePostEdit(e));

	// Check for saved dark mode preference
	const savedDarkMode = localStorage.getItem("darkMode") === "true";
	if (savedDarkMode) {
		document.body.classList.add("dark-mode");
	}

	this.handleResize();

	// Listen for user data changes
	this.authCheckInterval = setInterval(async () => {
		const newUserdata = await getUserData();
		this.handleUserChange(newUserdata);
	}, 2000);

	this.watchNotifications();
};


Header.prototype.handlePostEdit = function (e) {
	console.log("check....");
	e.stopPropagation();

	this.postModalManager.openModal();

	// postModal.openModal();
}
// Start the application
document.addEventListener("DOMContentLoaded", () => {
	setTimeout(() => {
		const header = new Header();
		header.init();

		document.addEventListener(
			"click",
			() => (header.enableNotifications = true)
		);

		document.addEventListener("click", (event) => {
			const header = document.querySelector("#notificationButton");
			const dropdown = document.querySelector("#notificationDropdown");

			if (dropdown && dropdown.style.display === "block") {
				if (
					!header.contains(event.target) &&
					!dropdown.contains(event.target)
				) {
					dropdown.style.display = "none";
				}
			}
		});
	}, 500);
});
