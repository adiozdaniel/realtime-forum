import { postManager } from "./postmanager.js";

const sidebar = document.querySelector("#sidebar");
const allCategoriesBtn = document.querySelector("#allCategories");
const categoryDropdown = document.querySelector("#categoryDropdown");
const profileBtn = document.querySelector("#profile");
class Sidebar {
	constructor() {}
}

Sidebar.prototype.createCategoryDropdown = function () {
	// Ensure the dropdown container exists
	if (!categoryDropdown) return;

	// Hardcoded categories
	const categories = ["tutorial", "discussion", "guide", "question"];

	// Populate the category dropdown with checkboxes
	categoryDropdown.innerHTML = categories
		.map(
			(category) =>
				`<label><input type="checkbox" class="category-checkbox" value="${category}"> ${category}</label>`
		)
		.join("");

	// Add "All Posts" option at the top
	const allCategoriesLabel = document.createElement("label");
	allCategoriesLabel.innerHTML = `<input type="checkbox" class="category-checkbox" value="all" checked> All Posts`;

	// Ensure "All Posts" appears at the top
	categoryDropdown.insertBefore(allCategoriesLabel, categoryDropdown.firstChild);
};

Sidebar.prototype.filterPosts = function () {
	const selectedCategories = Array.from(
		document.querySelectorAll(".category-checkbox:checked")
	).map((checkbox) => checkbox.value);

	postManager.searchPosts("", selectedCategories);
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
		categoryDropdown.addEventListener("change", (e) => {
			if (e.target.classList.contains("category-checkbox")) {
				const checkbox = e.target;
				const checkboxes = document.querySelectorAll(".category-checkbox");
				if (checkbox.value === "all") {
					checkboxes.forEach((cb) => {
						if (cb !== checkbox) cb.checked = false;
					});
				} else {
					document.querySelector('.category-checkbox[value="all"]').checked = false;
				}
				this.filterPosts();
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

export {sidebar};
