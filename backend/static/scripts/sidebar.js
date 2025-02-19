import { PostManager, POSTS } from "./postmanager.js";
const sidebar = document.querySelector("#sidebar");
const allCategoriesBtn = document.querySelector("#allCategories");
const categoryDropdown = document.querySelector("#categoryDropdown");
const profileBtn = document.querySelector("#profile");

const postManager = new PostManager();

class Sidebar {
	constructor() {}

	// Handle window resize
	handleResize() {
		if (sidebar) {
			if (window.innerWidth >= 768) {
				sidebar.classList.add("visible");
				sidebar.classList.remove("hidden");
			} else {
				sidebar.classList.add("hidden");
				sidebar.classList.remove("visible");
			}
		}
	}

	createCategoryDropdown() {
		const categories = [...new Set(POSTS.map((post) => post.category))];
		if (categoryDropdown) {
			categoryDropdown.innerHTML = categories
				.map(
					(category) =>
						` <label> <input type="checkbox" class="category-checkbox" value="${category}"> ${category} </label> `
				)
				.join("");
		}

		// Add "All Categories" option
		const allCategoriesLabel = document.createElement("label");
		allCategoriesLabel.innerHTML = ` <input type="checkbox" class="category-checkbox" value="all" checked> All Posts `;

		if (categoryDropdown) {
			categoryDropdown.insertBefore(
				allCategoriesLabel,
				categoryDropdown.firstChild
			);
		}
	}

	// Filter posts based on selected categories
	filterPosts() {
		const selectedCategories = Array.from(
			document.querySelectorAll(".category-checkbox:checked")
		).map((checkbox) => checkbox.value);
		const isAllSelected = selectedCategories.includes("all");
		const filteredPosts = isAllSelected
			? POSTS
			: POSTS.filter((post) => selectedCategories.includes(post.category));
		postManager.renderPosts(filteredPosts);
	}

	// Initialize function
	init() {
		this.handleResize();
		this.createCategoryDropdown();

		// Toggle dropdown
		if (allCategoriesBtn) {
			allCategoriesBtn.addEventListener("click", (e) => {
				e.stopPropagation();
				categoryDropdown.classList.toggle("hidden");
			});
		}

		// Toggle profile
		if (profileBtn) {
			profileBtn.addEventListener("click", (e) => {
				e.stopPropagation();
				window.location.href = "/profile";
			});
		}

		// Handle category selection inside Sidebar class
		if (categoryDropdown) {
			categoryDropdown.addEventListener("change", (e) => {
				if (e.target.classList.contains("category-checkbox")) {
					const checkbox = e.target;
					const checkboxes = document.querySelectorAll(".category-checkbox");
					if (checkbox.value === "all") {
						checkboxes.forEach((cb) => {
							if (cb !== checkbox) cb.checked = false;
						});
					} else {
						const allCategoriesCheckbox = document.querySelector(
							'.category-checkbox[value="all"]'
						);
						allCategoriesCheckbox.checked = false;
					}
					this.filterPosts();
				}
			});
		}

		// Close dropdown when clicking outside
		document.addEventListener("click", (e) => {
			if (
				categoryDropdown &&
				!categoryDropdown.contains(e.target) &&
				e.target !== allCategoriesBtn
			) {
				categoryDropdown.classList.add("hidden");
			}
		});
	}
}

// Start the application
document.addEventListener("DOMContentLoaded", () => {
	const sideBar = new Sidebar();
	sideBar.init();
});

export { sidebar };
