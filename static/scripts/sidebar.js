import { PostManager } from "./postmanager.js";

const sidebar = document.querySelector("#sidebar");
const allCategoriesBtn = document.querySelector("#allCategories");
const categoryDropdown = document.querySelector("#categoryDropdown");
const profileBtn = document.querySelector("#profile");
class Sidebar {
	constructor() {
		this.postManager = new PostManager();
	}
}

Sidebar.prototype.createCategoryDropdown = function () {
	// Ensure the dropdown container exists
	if (!categoryDropdown) return;

	// Hardcoded categories
	const categories = ["tutorial", "discussion", "guide", "question"];

	// Populate the category dropdown with pill buttons instead of checkboxes
	categoryDropdown.innerHTML = categories
		.map(
			(category) =>
				`<button class="category-pill" data-category="${category}">${category}</button>`
		)
		.join("");

	// Add "All Posts" option at the top
	const allCategoriesButton = document.createElement("button");
	allCategoriesButton.className = "category-pill active";
	allCategoriesButton.dataset.category = "all";
	allCategoriesButton.textContent = "All Posts";

	// Ensure "All Posts" appears at the top
	categoryDropdown.insertBefore(
		allCategoriesButton,
		categoryDropdown.firstChild
	);
};

Sidebar.prototype.filterPosts = function (selectedCategory) {
	// Convert single category to array for compatibility with existing code
	const selectedCategories =
		selectedCategory === "all" ? ["all"] : [selectedCategory];

	this.postManager.searchPosts("", selectedCategories);
};

Sidebar.prototype.init = function () {
	this.createCategoryDropdown();

	if (allCategoriesBtn) {
		allCategoriesBtn.addEventListener("click", function (e) {
			e.stopPropagation();
			categoryDropdown.classList.toggle("hidden");
		});
	}

	if (profileBtn) {
		profileBtn.addEventListener("click", function (e) {
			e.stopPropagation();
			window.location.href = "/profile";
		});
	}

	if (categoryDropdown) {
		categoryDropdown.addEventListener("click", (e) => {
			// Handle pill button clicks
			if (e.target.classList.contains("category-pill")) {
				const pillButtons = document.querySelectorAll(".category-pill");

				// Remove active class from all buttons
				pillButtons.forEach((pill) => pill.classList.remove("active"));

				// Add active class to clicked button
				e.target.classList.add("active");

				// Get selected category and filter posts
				const selectedCategory = e.target.dataset.category;
				this.filterPosts(selectedCategory);
			}
		});
	}

	document.addEventListener("click", function (e) {
		if (
			categoryDropdown &&
			!categoryDropdown.contains(e.target) &&
			e.target !== allCategoriesBtn
		) {
			categoryDropdown.classList.add("hidden");
		}
	});
};

// Start the application
document.addEventListener("DOMContentLoaded", function () {
	const sideBar = new Sidebar();
	sideBar.init();
});

export { sidebar };
