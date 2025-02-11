import { renderPosts, SAMPLE_POSTS } from "./index.js";
const sidebar = document.querySelector('#sidebar'); 
const allCategoriesBtn = document.querySelector('#allCategories');
const categoryDropdown = document.querySelector('#categoryDropdown')
const profileBtn = document.querySelector('#profile');

// Handle window resize
function handleResize() {
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

function createCategoryDropdown() {
	const categories = [...new Set(SAMPLE_POSTS.map((post) => post.category))];
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

// Toggle dropdown
allCategoriesBtn.addEventListener('click', (e) => {
    e.stopPropagation();
    categoryDropdown.classList.toggle('hidden');
});

// Toggle profile
profileBtn.addEventListener('click', (e) => {
    e.stopPropagation();
    window.location.href = "http://localhost:4000/api/profile"
});

// Close dropdown when clicking outside
document.addEventListener("click", (e) => {
	if (categoryDropdown) {
		if (!categoryDropdown.contains(e.target) && e.target !== allCategoriesBtn) {
			categoryDropdown.classList.add("hidden");
		}
	}
});

// Handle category selection
if (categoryDropdown) {
	categoryDropdown.addEventListener("change", (e) => {
		if (e.target.classList.contains("category-checkbox")) {
			const checkbox = e.target;
			const checkboxes = document.querySelectorAll(".category-checkbox");
			if (checkbox.value === "all") {
				// If "All Categories" is selected, uncheck others
				checkboxes.forEach((cb) => {
					if (cb !== checkbox) cb.checked = false;
				});
			} else {
				// If a specific category is selected, uncheck "All Categories"
				const allCategoriesCheckbox = document.querySelector(
					'.category-checkbox[value="all"]'
				);
				allCategoriesCheckbox.checked = false;
			}
			filterPosts();
		}
	});
}

// Filter posts based on selected categories
function filterPosts() {
	const selectedCategories = Array.from(
		document.querySelectorAll(".category-checkbox:checked")
	).map((checkbox) => checkbox.value);
	const isAllSelected = selectedCategories.includes("all");
	const filteredPosts = isAllSelected
		? SAMPLE_POSTS
		: SAMPLE_POSTS.filter((post) => selectedCategories.includes(post.category));
	renderPosts(filteredPosts);
}

// Initialize function
function init() {
	// Initial resize check
	handleResize();

	// categories dropdown
	createCategoryDropdown();
}

// Start the application
document.addEventListener("DOMContentLoaded", init);

export { handleResize, sidebar };
