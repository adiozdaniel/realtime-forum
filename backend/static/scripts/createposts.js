document.addEventListener("DOMContentLoaded", () => {
	lucide.createIcons();
	const modal = document.getElementById("modal");
	const createPostBtn = document.getElementById("createPostBtn");
	const cancelBtn = document.getElementById("cancelBtn");
	const postForm = document.getElementById("postForm");
	const title = document.getElementById("modal-content-title").value;
	const user_name = document.getElementById("author").value;
	const category = document.getElementById("category").value;
	const excerpt = document.getElementById("excerpt").value;

	// Open modal
	createPostBtn.addEventListener("click", () => {
		modal.classList.add("open");
	});

	// Close modal when clicking outside
	modal.addEventListener("click", (e) => {
		if (e.target === modal) {
			modal.classList.remove("open");
		}
	});

	// Close modal with cancel button
	cancelBtn.addEventListener("click", () => {
		modal.classList.remove("open");
	});

	// Handle form submission
	postForm.addEventListener("submit", (e) => {
		e.preventDefault();
		const formData = {
			title: title,
			user_name: user_name,
			category: category,
			excerpt: excerpt,
		};

		// Send the data to a server
		console.log("Form submitted:", formData);

		// Reset form and close modal
		postForm.reset();
		modal.classList.remove("open");
	});
});
