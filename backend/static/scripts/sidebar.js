import { postManager } from "./postmanager.js";
const sidebar = document.querySelector("#sidebar");
const allCategoriesBtn = document.querySelector("#allCategories");
const categoryDropdown = document.querySelector("#categoryDropdown");
const profileBtn = document.querySelector("#profile");

class Sidebar {
	constructor() {
		this.postManager = postManager;
	}
}

// Sidebar.prototype.createCategoryDropdown = function () {
// 	const categories = [...new Set(POSTS.map((post) => post.category))];
// 	if (categoryDropdown) {
// 		categoryDropdown.innerHTML = categories
// 			.map(
// 				(category) =>
// 					`<label><input type="checkbox" class="category-checkbox" value="${category}"> ${category}</label>`
// 			)
// 			.join("");
// 	}

// 	const allCategoriesLabel = document.createElement("label");
// 	allCategoriesLabel.innerHTML = `<input type="checkbox" class="category-checkbox" value="all" checked> All Posts`;

// 	if (categoryDropdown) {
// 		categoryDropdown.insertBefore(allCategoriesLabel, categoryDropdown.firstChild);
// 	}
// };

// Sidebar.prototype.filterPosts = function () {
// 	const selectedCategories = Array.from(
// 		document.querySelectorAll(".category-checkbox:checked")
// 	).map((checkbox) => checkbox.value);
// 	const isAllSelected = selectedCategories.includes("all");
// 	const filteredPosts = isAllSelected
// 		? POSTS
// 		: POSTS.filter((post) => selectedCategories.includes(post.category));
// 	postManager.renderPosts(filteredPosts);
// };

Sidebar.prototype.init = function () {
	// this.createCategoryDropdown();

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

export { sidebar };
